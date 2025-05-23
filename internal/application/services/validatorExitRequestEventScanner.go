package services

import (
	"context"
	"encoding/hex"
	"fmt"
	"lido-events/internal/application/domain"
	"lido-events/internal/application/ports"
	"lido-events/internal/logger"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type ValidatorExitRequestEventScanner struct {
	storagePort         ports.StoragePort
	notifierPort        ports.NotifierPort
	veboPort            ports.VeboPort
	executionPort       ports.ExecutionPort
	beaconchainPort     ports.Beaconchain
	veboBlockDeployment uint64
	csModuleTxReceipt   common.Hash
	mu                  sync.Mutex
	servicePrefix       string
}

func NewValidatorExitRequestEventScanner(storagePort ports.StoragePort, notifierPort ports.NotifierPort, veboPort ports.VeboPort, executionPort ports.ExecutionPort, beaconchainPort ports.Beaconchain, veboBlockDeployment uint64, csModuleTxReceipt common.Hash) *ValidatorExitRequestEventScanner {
	return &ValidatorExitRequestEventScanner{
		storagePort,
		notifierPort,
		veboPort,
		executionPort,
		beaconchainPort,
		veboBlockDeployment,
		csModuleTxReceipt,
		sync.Mutex{},
		"ValidatorExitRequestEventScanner",
	}
}

// ScanValidatorExitRequestEventsCron runs a periodic scan for ValidatorExitRequest events
func (vs *ValidatorExitRequestEventScanner) ScanValidatorExitRequestEventsCron(ctx context.Context, interval time.Duration, wg *sync.WaitGroup, firstExecutionComplete chan struct{}) {
	defer wg.Done() // Decrement the counter when the goroutine finishes
	wg.Add(1)       // Increment the WaitGroup counter

	// Run the scan logic immediately
	vs.RunScan(ctx)

	logger.DebugWithPrefix(vs.servicePrefix, "First execution complete, sending signal to start periodic ejector cron for ValidatorExitRequest events")
	// Signal that the first execution is complete
	close(firstExecutionComplete)

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Run the scan logic periodically
			vs.RunScan(ctx)
		case <-ctx.Done():
			logger.InfoWithPrefix(vs.servicePrefix, "Stopping ValidatorExitRequest cron scan")
			return
		}
	}
}

// RunScan contains the execution logic for scanning ValidatorExitRequest events
func (vs *ValidatorExitRequestEventScanner) RunScan(ctx context.Context) error {
	vs.mu.Lock()         // Lock mutex to ensure only one execution at a time
	defer vs.mu.Unlock() // Unlock when function exits

	// Skip if execution or beaconchain nodes are syncing
	executionSyncing, err := vs.executionPort.IsSyncing()
	if err != nil {
		logger.ErrorWithPrefix(vs.servicePrefix, "Error checking if execution node is syncing: %v", err)
		return err
	}
	if executionSyncing {
		logger.InfoWithPrefix(vs.servicePrefix, "Execution node is syncing, skipping ValidatorExitRequest event scan")
		return fmt.Errorf("execution node is syncing")
	}
	beaconchainSyncing, err := vs.beaconchainPort.GetSyncingStatus()
	if err != nil {
		logger.ErrorWithPrefix(vs.servicePrefix, "Error checking if beaconchain node is syncing: %v", err)
		return err
	}
	if beaconchainSyncing {
		logger.InfoWithPrefix(vs.servicePrefix, "Beaconchain node is syncing, skipping ValidatorExitRequest event scan")
		return fmt.Errorf("beaconchain node is syncing")
	}

	// Skip if tx receipt not found. This means that the node does not store log receipts and there are no logs at all
	receiptExists, err := vs.executionPort.GetTransactionReceiptExists(vs.csModuleTxReceipt)
	if err != nil {
		logger.ErrorWithPrefix(vs.servicePrefix, "Error checking if transaction receipt exists: %v", err)
		return err
	}
	if !receiptExists {
		logger.WarnWithPrefix(vs.servicePrefix, "Transaction receipt for csModule deployment not found. This probably means your node does not store log receipts, check out the official documentation of your node and configure the node to store them. Skipping ValidatorExitRequests event scan")
		// notify the user to switch to an execution client that does store the log receipts
		message := "- 🚨 Your Execution Client appears to be missing log receipt storage. As a result, ValidatorExitRequest events cannot be scanned. To resolve this issue, consider switching to an Execution Client that supports log receipt storage or updating your node configuration to enable this feature"
		if err := vs.notifierPort.SendNotification(message); err != nil {
			logger.ErrorWithPrefix(vs.servicePrefix, "Error sending notification: %v", err)
		}
		return fmt.Errorf("transaction receipt for csModule deployment not found")
	}

	// Iterrate over operator ids
	operatorIDs, err := vs.storagePort.GetOperatorIds()
	if err != nil {
		logger.ErrorWithPrefix(vs.servicePrefix, "Error getting operator ids: %v", err)
		return err
	}

	for _, operatorID := range operatorIDs {
		// Retrieve start and end blocks for scanning
		start, err := vs.storagePort.GetValidatorExitRequestLastProcessedBlock(operatorID)
		if err != nil {
			logger.WarnWithPrefix(vs.servicePrefix, "Failed to get last processed block, using deployment block %d: %v", vs.veboBlockDeployment, err)
			start = vs.veboBlockDeployment
		}

		if start == 0 {
			logger.InfoWithPrefix(vs.servicePrefix, "No last processed block found, using deployment block %d", vs.veboBlockDeployment)
			start = vs.veboBlockDeployment
		}

		end, err := vs.executionPort.GetMostRecentBlockNumber()
		if err != nil {
			logger.ErrorWithPrefix(vs.servicePrefix, "Failed to get most recent block number, cannot continue with cron execution, waiting for next iteration: %v", err)
			return err
		}

		// Perform the scan
		if err := vs.veboPort.ScanVeboValidatorExitRequestEvent(ctx, operatorID, start, &end, vs.HandleValidatorExitRequestEvent); err != nil {
			logger.ErrorWithPrefix(vs.servicePrefix, "Error scanning ValidatorExitRequest events: %v", err)
			return err
		}

		// Save the last processed block if successful
		if err := vs.storagePort.SaveValidatorExitRequestLastProcessedBlock(operatorID, end); err != nil {
			logger.ErrorWithPrefix(vs.servicePrefix, "Error saving last processed block, next cron execution will run from previous processed: %v", err)
			return err
		}
	}

	return nil
}

// HandleValidatorExitRequestEvent processes a ValidatorExitRequest event
func (vs *ValidatorExitRequestEventScanner) HandleValidatorExitRequestEvent(validatorExitEvent *domain.VeboValidatorExitRequest) error {
	logger.DebugWithPrefix(vs.servicePrefix, "Processing ValidatorExitRequest event for node operatror id %s and validator index %s", validatorExitEvent.NodeOperatorId, validatorExitEvent.ValidatorIndex)

	pubkeyHex := "0x" + hex.EncodeToString(validatorExitEvent.ValidatorPubkey) // This is exactly the format that signer needs for the pubkey to be when signing the exit message
	validatorStatus, err := vs.beaconchainPort.GetValidatorStatus(validatorExitEvent.ValidatorIndex.String())
	if err != nil {
		return err
	}

	// If validator is active and not exiting (ongoing or slashed), send a notification. We don't send a notification if the validator is already exiting or has already exited
	if validatorStatus == domain.StatusActiveOngoing || validatorStatus == domain.StatusActiveSlashed {
		logger.InfoWithPrefix(vs.servicePrefix, "Validator %s is active and has been requestes to exit", validatorExitEvent.ValidatorIndex)
		message := fmt.Sprintf("- 🚨 Your validator with ID: %s has been requested to exit. It will be automatically ejected within the next hour, you will receive a notification when exited", validatorExitEvent.ValidatorIndex)
		err := vs.notifierPort.SendNotification(message)
		if err != nil {
			logger.ErrorWithPrefix(vs.servicePrefix, "Error sending notification: %v", err)
		}
	}

	exitRequest := domain.ExitRequest{
		Event:              *validatorExitEvent,
		Status:             validatorStatus,
		ValidatorPubkeyHex: pubkeyHex, // Save the pubkey in hex format for the ejector to use later
	}
	if err := vs.storagePort.SaveExitRequest(validatorExitEvent.NodeOperatorId, validatorExitEvent.ValidatorIndex.String(), exitRequest); err != nil {
		logger.ErrorWithPrefix(vs.servicePrefix, "Error saving exit request for validator %s: %v", validatorExitEvent.ValidatorIndex, err)
		// if validator is active print attention warning log that manual exit is required
		if validatorStatus == domain.StatusActiveOngoing {
			logger.WarnWithPrefix(vs.servicePrefix, "Validator %s is still active and requires to exit, it could not be saved and will not be exited automatically, a manal exit is required", validatorExitEvent.ValidatorIndex)
		}
		return err
	}

	logger.DebugWithPrefix(vs.servicePrefix, "Finished processing ValidatorExitRequest event")

	return nil
}
