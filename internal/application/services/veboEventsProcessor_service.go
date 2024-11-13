package services

import (
	"context"
	"fmt"
	"lido-events/internal/application/domain"
	"lido-events/internal/application/ports"
	"log"
	"time"
)

type VeboEventsProcessor struct {
	storagePort         ports.StoragePort
	notifierPort        ports.NotifierPort
	veboPort            ports.VeboPort
	exitValidatorPort   ports.ExitValidator
	beaconchainPort     ports.Beaconchain
	veboBlockDeployment uint64
}

func NewVeboEventsProcessorService(storagePort ports.StoragePort, notifierPort ports.NotifierPort, veboPort ports.VeboPort, exitValidatorPort ports.ExitValidator, beaconchainPort ports.Beaconchain, veboBlockDeployment uint64) *VeboEventsProcessor {
	return &VeboEventsProcessor{
		storagePort,
		notifierPort,
		veboPort,
		exitValidatorPort,
		beaconchainPort,
		veboBlockDeployment,
	}
}

// ScanEventsCron starts a unified periodic scan for both ValidatorExitRequest and ReportSubmitted events.
func (vs *VeboEventsProcessor) ScanEventsCron(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Retrieve the start epoch from the last processed epoch, defaulting to the deployment block
			start, err := vs.storagePort.GetLastProcessedEpoch()
			if err != nil {
				log.Printf("Failed to get last processed epoch: %v", err)
				continue
			}
			if start == 0 {
				start = vs.veboBlockDeployment
			}

			// Retrieve the end epoch as the latest finalized epoch from the beacon chain
			end, err := vs.beaconchainPort.GetEpochHeader("finalized")
			if err != nil {
				log.Printf("Failed to get latest finalized epoch: %v", err)
				continue
			}

			// Scan for ValidatorExitRequest events
			if err := vs.veboPort.ScanVeboValidatorExitRequestEvent(ctx, start, &end, vs.HandleValidatorExitRequestEvent); err != nil {
				log.Printf("Error scanning ValidatorExitRequest events: %v", err)
				continue
			}

			// Scan for ReportSubmitted events
			if err := vs.veboPort.ScanReportSubmittedEvents(ctx, start, &end, vs.HandleReportSubmittedEvent); err != nil {
				log.Printf("Error scanning ReportSubmitted events: %v", err)
				continue
			}

			// Update the last processed epoch if both scans were successful
			if err := vs.storagePort.SaveLastProcessedEpoch(end); err != nil {
				log.Printf("Failed to save last processed epoch: %v", err)
			}
		case <-ctx.Done():
			// Stop the periodic scan if the context is canceled
			log.Println("Stopping unified periodic scan for events")
			return
		}
	}
}

// Make VeboService implement the VeboHandlers interface by adding the methods
func (vs *VeboEventsProcessor) HandleValidatorExitRequestEvent(validatorExitEvent *domain.VeboValidatorExitRequest) error {
	// get validator status
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

	// get last processed epoch finalized from the block hash of the event
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

	// save exit request
	// TODO: verify NodeOperatorId trully exists in the event
	if err := vs.storagePort.SaveExitRequest(validatorExitEvent.NodeOperatorId, validatorExitEvent.ValidatorIndex.String(), exitRequest); err != nil {
		log.Printf("Failed to save exit requests: %v", err)
		return err
	}

	// save epoch
	if err := vs.storagePort.SaveLastProcessedEpoch(epoch); err != nil {
		log.Printf("Failed to save last processed epoch: %v", err)
		return err
	}

	return nil
}

func (vs *VeboEventsProcessor) HandleReportSubmittedEvent(reportSubmitted *domain.VeboReportSubmitted) error {
	// Send the notification message
	message := fmt.Sprintf("- 📈 New submitted report: %s", reportSubmitted.RefSlot)
	if err := vs.notifierPort.SendNotification(message); err != nil {
		log.Printf("Failed to send notification: %v", err)
		return err
	}

	// Store the pending hash
	stringHash := fmt.Sprintf("0x%x", reportSubmitted.Hash)
	if err := vs.storagePort.AddPendingHash(stringHash); err != nil {
		log.Printf("Failed to store pending hash: %v", err)
		return err
	}

	return nil
}
