package services

import (
	"context"
	"fmt"
	"lido-events/internal/application/domain"
	"lido-events/internal/application/ports"
	"log"
	"math/big"
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

// WatchReportSubmittedEvents subscribes to Ethereum events and handles them.
// It passes to the csModule port a function that processes the log data.
func (vs *VeboEventsProcessor) WatchReportSubmittedEvents(ctx context.Context) error {
	return vs.veboPort.WatchReportSubmittedEvents(ctx, vs.HandleReportSubmittedEvent)
}

// ScanVeboValidatorExitRequestEventCron starts a periodic scan for ValidatorExitRequest events.
func (vs *VeboEventsProcessor) ScanVeboValidatorExitRequestEventCron(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Get start from last processed epoch finalized from storage
			start, err := vs.storagePort.GetLastProcessedEpoch()
			if err != nil {
				log.Printf("Failed to get last processed epoch: %v", err)
			}
			// 0 is the default value for the first run
			if start == 0 {
				start = vs.veboBlockDeployment
			}

			// Get end from latest finalized epoch from the chain
			end, err := vs.beaconchainPort.GetEpochHeader("finalized")
			if err != nil {
				log.Printf("Failed to get latest finalized epoch: %v", err)
			}

			// Scan the events
			if err := vs.veboPort.ScanVeboValidatorExitRequestEvent(ctx, start, &end, vs.HandleValidatorExitRequestEvent); err != nil {
				log.Printf("Error scanning ValidatorExitRequest events: %v", err)
				break
			}

			// update last processed epoch finalized if no error
			if err := vs.storagePort.SaveLastProcessedEpoch(end); err != nil {
				log.Printf("Failed to save last processed epoch: %v", err)
			}
		case <-ctx.Done():
			// Stop the periodic scan if the context is canceled
			log.Println("Stopping periodic scan for ValidatorExitRequest events")
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
	// send the notification message
	message := fmt.Sprintf("- 📈 New submitted report: %s", reportSubmitted.RefSlot)
	if err := vs.notifierPort.SendNotification(message); err != nil {
		log.Printf("Failed to send notification: %v", err)
		return err
	}

	// TODO: use reportSubmitted.Hash to fetch the report from IPFS

	// save the report
	lidoReport := domain.Report{
		Threshold:  "some-threshold",
		Validators: map[string]domain.ValidatorPerformance{},
	}
	epoch := "123456"
	operatorId := big.NewInt(0)
	if err := vs.storagePort.SaveOperatorPerformance(operatorId, epoch, lidoReport); err != nil {
		log.Printf("Failed to save Lido report: %v", err)
		return err
	}

	return nil
}
