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

type ValidatorExitRequestEventScanner struct {
	storagePort         ports.StoragePort
	notifierPort        ports.NotifierPort
	veboPort            ports.VeboPort
	executionPort       ports.ExecutionPort
	beaconchainPort     ports.Beaconchain
	veboBlockDeployment uint64
	servicePrefix       string
}

func NewValidatorExitRequestEventScanner(storagePort ports.StoragePort, notifierPort ports.NotifierPort, veboPort ports.VeboPort, executionPort ports.ExecutionPort, beaconchainPort ports.Beaconchain, veboBlockDeployment uint64) *ValidatorExitRequestEventScanner {
	return &ValidatorExitRequestEventScanner{
		storagePort,
		notifierPort,
		veboPort,
		executionPort,
		beaconchainPort,
		veboBlockDeployment,
		"ValidatorExitRequestEventScanner",
	}
}

// ScanValidatorExitRequestEventsCron runs a periodic scan for DistributionLogUpdated events
func (vs *ValidatorExitRequestEventScanner) ScanValidatorExitRequestEventsCron(ctx context.Context, interval time.Duration, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine finishes
	wg.Add(1)       // Increment the WaitGroup counter

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Retrieve start and end blocks for scanning
			start, err := vs.storagePort.GetValidatorExitRequestLastProcessedBlock()
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
				logger.ErrorWithPrefix(vs.servicePrefix, "Failed to get most recent block number, cannot continue wit cron execution, waiting for next iteration: %v", err)
				continue
			}

			// Perform the scan
			if err := vs.veboPort.ScanVeboValidatorExitRequestEvent(ctx, start, &end, vs.HandleValidatorExitRequestEvent); err != nil {
				logger.ErrorWithPrefix(vs.servicePrefix, "Error scanning ValidatorExitRequest events: %v", err)
				continue
			}

			// Save the last processed epoch if successful
			if err := vs.storagePort.SaveValidatorExitRequestLastProcessedBlock(end); err != nil {
				logger.ErrorWithPrefix(vs.servicePrefix, "Error saving last processed block, next cron execution will run from previous processed: %v", err)
				continue
			}
		case <-ctx.Done():
			logger.InfoWithPrefix(vs.servicePrefix, "Stopping ValidatorExitRequest cron scan")
			return
		}
	}
}

// HandleValidatorExitRequestEvent processes a ValidatorExitRequest event
func (vs *ValidatorExitRequestEventScanner) HandleValidatorExitRequestEvent(validatorExitEvent *domain.VeboValidatorExitRequest) error {
	logger.DebugWithPrefix(vs.servicePrefix, "Processing ValidatorExitRequest event for node operatror id %s and validator index %s", validatorExitEvent.NodeOperatorId, validatorExitEvent.ValidatorIndex)

	pubkeyHex := "0x" + hex.EncodeToString(validatorExitEvent.ValidatorPubkey) // This is exactly the format that signer needs for the pubkey to be when signing the exit message
	validatorStatus, err := vs.beaconchainPort.GetValidatorStatus(validatorExitEvent.ValidatorIndex.String())
	if err != nil {
		return err
	}

	if validatorStatus == domain.StatusActiveOngoing {
		logger.InfoWithPrefix(vs.servicePrefix, "Validator %s is active and requires to exit", validatorExitEvent.ValidatorIndex)
		message := fmt.Sprintf("- 🚨 One of the validators requested to exit: %s. It will be automatically ejected within the next hour, you will receive a notification when exited", validatorExitEvent.ValidatorIndex)
		vs.notifierPort.SendNotification(message)
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
