package services

import (
	"context"
	"fmt"
	"lido-events/internal/application/domain"
	"lido-events/internal/application/ports"
	"log"
	"time"
)

type ValidatorExitRequestEventScanner struct {
	storagePort         ports.StoragePort
	notifierPort        ports.NotifierPort
	veboPort            ports.VeboPort
	executionPort       ports.ExecutionPort
	beaconchainPort     ports.Beaconchain
	veboBlockDeployment uint64
}

func NewValidatorExitRequestEventScanner(storagePort ports.StoragePort, notifierPort ports.NotifierPort, veboPort ports.VeboPort, executionPort ports.ExecutionPort, beaconchainPort ports.Beaconchain, veboBlockDeployment uint64) *ValidatorExitRequestEventScanner {
	return &ValidatorExitRequestEventScanner{
		storagePort,
		notifierPort,
		veboPort,
		executionPort,
		beaconchainPort,
		veboBlockDeployment,
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
			start, err := vs.storagePort.GetLastProcessedBlock()
			if err != nil {
				log.Printf("Failed to get last processed block: %v", err)
				continue
			}

			end, err := vs.executionPort.GetMostRecentBlockNumber()
			if err != nil {
				log.Printf("Failed to get latest finalized block: %v", err)
				continue
			}

			// Perform the scan
			if err := vs.veboPort.ScanVeboValidatorExitRequestEvent(ctx, start, &end, vs.HandleValidatorExitRequestEvent); err != nil {
				log.Printf("Error scanning ValidatorExitRequest events: %v", err)
				continue
			}

			// Save the last processed epoch if successful
			if err := vs.storagePort.SaveLastProcessedBlock(end); err != nil {
				log.Printf("Failed to save last processed epoch: %v", err)
			}
		case <-ctx.Done():
			log.Println("Stopping DistributionLogUpdated cron scan")
			return
		}
	}
}

// HandleValidatorExitRequestEvent processes a ValidatorExitRequest event
func (vs *ValidatorExitRequestEventScanner) HandleValidatorExitRequestEvent(validatorExitEvent *domain.VeboValidatorExitRequest) error {
	validatorStatus, err := vs.beaconchainPort.GetValidatorStatus(string(validatorExitEvent.ValidatorPubkey))
	if err != nil {
		log.Printf("Failed to get validator status: %v", err)
		return err
	}

	if validatorStatus == domain.StatusActiveOngoing {
		message := fmt.Sprintf("- 🚨 One of the validators requested to exit: %s. It will be automatically ejected within the next hour, you will receive a notification when exited", validatorExitEvent.ValidatorIndex)
		if err := vs.notifierPort.SendNotification(message); err != nil {
			log.Printf("Failed to send notification: %v", err)
			return err
		}
	}

	blockHash := fmt.Sprintf("0x%x", validatorExitEvent.Raw.BlockHash)
	epoch, err := vs.beaconchainPort.GetEpochHeader(blockHash)
	if err != nil {
		log.Printf("Failed to get epoch header: %v", err)
		return err
	}

	exitRequest := domain.ExitRequest{
		Event:  *validatorExitEvent,
		Status: validatorStatus,
	}

	if err := vs.storagePort.SaveExitRequest(validatorExitEvent.NodeOperatorId, validatorExitEvent.ValidatorIndex.String(), exitRequest); err != nil {
		log.Printf("Failed to save exit requests: %v", err)
		return err
	}

	if err := vs.storagePort.SaveLastProcessedBlock(epoch); err != nil {
		log.Printf("Failed to save last processed epoch: %v", err)
		return err
	}

	return nil
}
