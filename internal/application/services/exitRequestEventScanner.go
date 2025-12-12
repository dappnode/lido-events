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
)

type ExitRequestEventScanner struct {
	storagePort             ports.ExitsStorage
	notifierPort            ports.NotifierPort
	veboPort                ports.VeboPort
	executionPort           ports.ExecutionPort
	beaconchainPort         ports.Beaconchain
	csParametersPort        ports.CsParametersPort
	secondsPerSlot          uint64
	defaultAllowedExitDelay uint64
	exitDelayMultiplier     uint64
	mu                      sync.Mutex
	servicePrefix           string
}

func NewExitRequestEventScanner(storagePort ports.ExitsStorage, notifierPort ports.NotifierPort, veboPort ports.VeboPort, executionPort ports.ExecutionPort, beaconchainPort ports.Beaconchain, csParametersPort ports.CsParametersPort, secondsPerSlot uint64, defaultAllowedExitDelay uint64, exitDelayMultiplier uint64) *ExitRequestEventScanner {
	return &ExitRequestEventScanner{
		storagePort,
		notifierPort,
		veboPort,
		executionPort,
		beaconchainPort,
		csParametersPort,
		secondsPerSlot,
		defaultAllowedExitDelay,
		exitDelayMultiplier,
		sync.Mutex{},
		"ValidatorExitRequestEventScanner",
	}
}

// ScanExitRequestEventsCron runs a periodic scan for ValidatorExitRequest events
func (vs *ExitRequestEventScanner) ScanExitRequestEventsCron(ctx context.Context, interval time.Duration, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine finishes
	wg.Add(1)       // Increment the WaitGroup counter

	// Run the scan logic immediately
	vs.RunScan(ctx)

	logger.DebugWithPrefix(vs.servicePrefix, "Starting periodic cron scan for ValidatorExitRequest events")

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
func (vs *ExitRequestEventScanner) RunScan(ctx context.Context) error {
	vs.mu.Lock()         // Lock mutex to ensure only one execution at a time
	defer vs.mu.Unlock() // Unlock when function exits

	if err := vs.ensureNodesInSync(); err != nil {
		return err
	}

	if err := vs.scanOperators(ctx); err != nil {
		return err
	}

	return nil
}

func (vs *ExitRequestEventScanner) ensureNodesInSync() error {
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

	return nil
}

// computeAllowedExitDelayBlock computes the block number corresponding to the allowed exit delay window (approximately 4 days ago from now multiplied by 2 to be safe)
func (vs *ExitRequestEventScanner) computeAllowedExitDelayBlock(ctx context.Context) (uint64, error) {
	// get allowed exit delay from csParameters contract
	allowedExitDelay, err := vs.csParametersPort.GetDefaultAllowedExitDelay(ctx)
	if err != nil {
		// continue on error and default to 4 days in seconds
		allowedExitDelay = vs.defaultAllowedExitDelay
		logger.ErrorWithPrefix(vs.servicePrefix, "Error getting default allowed exit delay from csParameters contract: %v. Defaulting to %d seconds", err, allowedExitDelay)
	}

	// Multiply it by 2 to be sure we cover the entire exit delay window
	allowedExitDelaySeconds := allowedExitDelay * vs.exitDelayMultiplier
	logger.InfoWithPrefix(vs.servicePrefix, "Default allowed exit delay multiplied by 2 to safely cover entire exit delay window: %d seconds", allowedExitDelaySeconds)

	// calculate its timestamp slot
	currentTime := uint64(time.Now().Unix())
	allowedExitDelayTimestamp := currentTime - allowedExitDelaySeconds

	// calculate allowedExitDelaySlot based on the timestamp using the genesis
	genesisTime, err := vs.beaconchainPort.GetGenesisTime()
	if err != nil {
		logger.ErrorWithPrefix(vs.servicePrefix, "Error getting genesis time from beaconchain: %v", err)
		return 0, err
	}

	var slotsSinceGenesis uint64
	if allowedExitDelayTimestamp <= genesisTime {
		slotsSinceGenesis = 0
	} else {
		slotsSinceGenesis = (allowedExitDelayTimestamp - genesisTime) / vs.secondsPerSlot
	}

	logger.InfoWithPrefix(vs.servicePrefix, "Allowed exit delay timestamp: %d corresponds to slot: %d", allowedExitDelayTimestamp, slotsSinceGenesis)

	// get the corresponding block number for the allowed exit delay timestamp slot
	blockNumber, err := vs.beaconchainPort.GetBlockNumber(fmt.Sprintf("%d", slotsSinceGenesis))
	if err != nil {
		logger.ErrorWithPrefix(vs.servicePrefix, "Error getting allowed exit delay slot block number: %v", err)
		return 0, err
	}

	// print the allowed exit delay slot and block number
	logger.InfoWithPrefix(vs.servicePrefix, "Allowed exit delay slot: %d, block number: %d", slotsSinceGenesis, blockNumber)

	// ensure block has receipts
	if err := vs.ensureBlockHasReceipts(blockNumber); err != nil {
		return 0, err
	}

	return blockNumber, nil
}

func (vs *ExitRequestEventScanner) ensureBlockHasReceipts(blockNumber uint64) error {
	// Skip if the block does not have tx receipts which means the execution client does not store log receipts
	blockID := fmt.Sprintf("0x%x", blockNumber) // convert block number to hex
	receiptExists, err := vs.executionPort.GetBlockHasReceipts(blockID)
	if err != nil {
		logger.ErrorWithPrefix(vs.servicePrefix, "Error checking if block has receipts: %v", err)
		return err
	}
	if !receiptExists {
		logger.WarnWithPrefix(vs.servicePrefix, "Block %s does not have transaction receipts. This probably means your node does not store log receipts, check out the official documentation of your node and configure the node to store them. Skipping ValidatorExitRequests event scan", blockID)
		// notify the user to switch to an execution client that does store the log receipts
		message := "- 🚨 Your Execution Client appears to be missing log receipt storage. As a result, ValidatorExitRequest events cannot be scanned. To resolve this issue, consider switching to an Execution Client that supports log receipt storage or updating your node configuration to enable this feature"
		if err := vs.notifierPort.SendMissingLogReceiptsNotification(message); err != nil {
			logger.ErrorWithPrefix(vs.servicePrefix, "Error sending notification: %v", err)
		}
		return fmt.Errorf("block %s does not have transaction receipts", blockID)
	}

	return nil
}

func (vs *ExitRequestEventScanner) scanOperators(ctx context.Context) error {
	// Iterate over operator ids
	operatorIDs, err := vs.storagePort.GetOperatorIds()
	if err != nil {
		logger.ErrorWithPrefix(vs.servicePrefix, "Error getting operator ids: %v", err)
		return err
	}

	for _, operatorID := range operatorIDs {
		logger.InfoWithPrefix(vs.servicePrefix, "Scanning ValidatorExitRequest events for operator ID %s", operatorID.String())
		// Retrieve start and end blocks for scanning.
		start, err := vs.storagePort.GetValidatorExitRequestLastProcessedBlock(operatorID)
		if start == 0 || err != nil {
			// If we don't have a last processed block, we approximate by using the
			// "allowed exit delay" window (Lido CSM specifies an exit delay of ~4/5 days).
			// See: https://csm.lido.fi/type/parameters
			start, err = vs.computeAllowedExitDelayBlock(ctx)
			if err != nil {
				logger.ErrorWithPrefix(vs.servicePrefix, "Error computing allowed exit delay block: %v", err)
				return err
			}
			logger.InfoWithPrefix(vs.servicePrefix, "No last processed block found, using allowed exit delay block number %d", start)
		}

		end, err := vs.executionPort.GetMostRecentBlockNumber()
		if err != nil {
			logger.ErrorWithPrefix(vs.servicePrefix, "Failed to get most recent block number, cannot continue with cron execution, waiting for next iteration: %v", err)
			return err
		}

		// Perform the scan
		logger.DebugWithPrefix(vs.servicePrefix, "Scanning ValidatorExitRequest events from block %d to %d for operator ID %s", start, end, operatorID.String())
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
func (vs *ExitRequestEventScanner) HandleValidatorExitRequestEvent(validatorExitEvent *domain.VeboValidatorExitRequest) error {
	logger.DebugWithPrefix(vs.servicePrefix, "Processing ValidatorExitRequest event for node operator id %s and validator index %s", validatorExitEvent.NodeOperatorId, validatorExitEvent.ValidatorIndex)

	pubkeyHex := "0x" + hex.EncodeToString(validatorExitEvent.ValidatorPubkey) // This is exactly the format that signer needs for the pubkey to be when signing the exit message
	validatorStatus, err := vs.beaconchainPort.GetValidatorStatus(validatorExitEvent.ValidatorIndex.String())
	if err != nil {
		return err
	}

	// If validator is active and not exiting (ongoing or slashed), send a notification. We don't send a notification if the validator is already exiting or has already exited
	if validatorStatus == domain.StatusActiveOngoing || validatorStatus == domain.StatusActiveSlashed {
		logger.InfoWithPrefix(vs.servicePrefix, "Validator %s is active and has been requested to exit", validatorExitEvent.ValidatorIndex)
		message := fmt.Sprintf("- 🚨 Your validator with ID: %s has been requested to exit. It will be automatically ejected within the next hour, you will receive a notification when exited", validatorExitEvent.ValidatorIndex)
		err := vs.notifierPort.SendValidatorExitRequestedNotification(message)
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
			logger.WarnWithPrefix(vs.servicePrefix, "Validator %s is still active and requires to exit, it could not be saved and will not be exited automatically, a manual exit is required", validatorExitEvent.ValidatorIndex)
		}
		return err
	}

	logger.DebugWithPrefix(vs.servicePrefix, "Finished processing ValidatorExitRequest event")

	return nil
}
