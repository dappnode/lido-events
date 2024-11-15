package services

import (
	"context"
	"fmt"
	"lido-events/internal/application/domain"
	"lido-events/internal/application/ports"
	"log"
	"time"
)

type EventsScanner struct {
	storagePort              ports.StoragePort
	notifierPort             ports.NotifierPort
	veboPort                 ports.VeboPort
	csFeeDistributorImplPort ports.CsFeeDistributorImplPort
	beaconchainPort          ports.Beaconchain
	veboBlockDeployment      uint64
}

func NewEventsScannerService(storagePort ports.StoragePort, notifierPort ports.NotifierPort, veboPort ports.VeboPort, csFeeDistributorPort ports.CsFeeDistributorImplPort, beaconchainPort ports.Beaconchain, veboBlockDeployment uint64) *EventsScanner {
	return &EventsScanner{
		storagePort,
		notifierPort,
		veboPort,
		csFeeDistributorPort,
		beaconchainPort,
		veboBlockDeployment,
	}
}

// ScanEventsCron starts a unified periodic scan for both ValidatorExitRequest and DistributionLogUpdated events
func (vs *EventsScanner) ScanEventsCron(ctx context.Context, interval time.Duration) {
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

			// Scan for DistributionLogUpdated events
			if err := vs.csFeeDistributorImplPort.ScanDistributionLogUpdatedEvents(ctx, start, &end, vs.HandleDistributionLogUpdatedEvent); err != nil {
				log.Printf("Error scanning DistributionLogUpdated events: %v", err)
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
func (vs *EventsScanner) HandleValidatorExitRequestEvent(validatorExitEvent *domain.VeboValidatorExitRequest) error {
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

func (vs *EventsScanner) HandleDistributionLogUpdatedEvent(distributionLogUpdated *domain.BindingsDistributionLogUpdated) error {
	// Store the CID
	if err := vs.storagePort.AddPendingHash(distributionLogUpdated.LogCid); err != nil {
		log.Printf("Failed to store pending CID: %v", err)
		return err
	}

	// Send notification
	message := fmt.Sprintf("- 📦 New distribution log updated: %s", distributionLogUpdated.LogCid)
	if err := vs.notifierPort.SendNotification(message); err != nil {
		log.Printf("Failed to send notification: %v", err)
		return err
	}

	return nil
}
