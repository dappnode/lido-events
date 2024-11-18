package services

import (
	"context"
	"fmt"
	"lido-events/internal/application/domain"
	"lido-events/internal/application/ports"
	"log"
	"os"
	"time"
)

type ValidatorExitRequestEventScanner struct {
	storagePort         ports.StoragePort
	notifierPort        ports.NotifierPort
	veboPort            ports.VeboPort
	executionPort       ports.ExecutionPort
	beaconchainPort     ports.Beaconchain
	veboBlockDeployment uint64
	logger              *log.Logger
}

func NewValidatorExitRequestEventScanner(storagePort ports.StoragePort, notifierPort ports.NotifierPort, veboPort ports.VeboPort, executionPort ports.ExecutionPort, beaconchainPort ports.Beaconchain, veboBlockDeployment uint64) *ValidatorExitRequestEventScanner {
	logger := log.New(os.Stdout, "[ValidatorExitRequestEventScanner] ", log.LstdFlags)

	return &ValidatorExitRequestEventScanner{
		storagePort,
		notifierPort,
		veboPort,
		executionPort,
		beaconchainPort,
		veboBlockDeployment,
		logger,
	}
}

// ScanValidatorExitRequestEventsCron runs a periodic scan for DistributionLogUpdated events
func (vs *ValidatorExitRequestEventScanner) ScanValidatorExitRequestEventsCron(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Retrieve start and end blocks for scanning
			start, err := vs.storagePort.GetValidatorExitRequestLastProcessedBlock()
			if err != nil {
				vs.logger.Printf("Failed to get last processed block, using block deployment: %v", err)
				start = vs.veboBlockDeployment
			}

			if start == 0 {
				vs.logger.Printf("No last processed block found, using deployment block %d", vs.veboBlockDeployment)
				start = vs.veboBlockDeployment
			}

			end, err := vs.executionPort.GetMostRecentBlockNumber()
			if err != nil {
				continue
			}

			// Perform the scan
			if err := vs.veboPort.ScanVeboValidatorExitRequestEvent(ctx, start, &end, vs.HandleValidatorExitRequestEvent); err != nil {
				continue
			}

			// Save the last processed epoch if successful
			if err := vs.storagePort.SaveValidatorExitRequestLastProcessedBlock(end); err != nil {
				continue
			}
		case <-ctx.Done():
			vs.logger.Println("Stopping DistributionLogUpdated cron scan")
			return
		}
	}
}

// HandleValidatorExitRequestEvent processes a ValidatorExitRequest event
func (vs *ValidatorExitRequestEventScanner) HandleValidatorExitRequestEvent(validatorExitEvent *domain.VeboValidatorExitRequest) error {
	vs.logger.Printf("Processing ValidatorExitRequest event: %v", validatorExitEvent)

	validatorStatus, err := vs.beaconchainPort.GetValidatorStatus(string(validatorExitEvent.ValidatorPubkey))
	if err != nil {
		return err
	}

	if validatorStatus == domain.StatusActiveOngoing {
		vs.logger.Printf("Validator %s is still active and requires to exit, it will be automatically exited by the ejector", validatorExitEvent.ValidatorIndex)
		message := fmt.Sprintf("- 🚨 One of the validators requested to exit: %s. It will be automatically ejected within the next hour, you will receive a notification when exited", validatorExitEvent.ValidatorIndex)
		if err := vs.notifierPort.SendNotification(message); err != nil {
			// dont return error here, continue with the process so the request is stored for later usage
		}
	}

	exitRequest := domain.ExitRequest{
		Event:  *validatorExitEvent,
		Status: validatorStatus,
	}

	if err := vs.storagePort.SaveExitRequest(validatorExitEvent.NodeOperatorId, validatorExitEvent.ValidatorIndex.String(), exitRequest); err != nil {
		// if validator is active print attention warning log that manual exit is required
		if validatorStatus == domain.StatusActiveOngoing {
			vs.logger.Printf("ATTENTION: Validator %s is still active and requires to exit, it could not be saved and will not be exited automatically, a manal exit is required", validatorExitEvent.ValidatorIndex)
		}
		return err
	}

	return nil
}
