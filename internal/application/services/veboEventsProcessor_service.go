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
	storagePort       ports.StoragePort
	notifierPort      ports.NotifierPort
	veboPort          ports.VeboPort
	exitValidatorPort ports.ExitValidator
}

func NewVeboEventsProcessorService(storagePort ports.StoragePort, notifierPort ports.NotifierPort, veboPort ports.VeboPort, exitValidatorPort ports.ExitValidator) *VeboEventsProcessor {
	return &VeboEventsProcessor{
		storagePort,
		notifierPort,
		veboPort,
		exitValidatorPort,
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
			// Call the scan method periodically
			if err := vs.veboPort.ScanVeboValidatorExitRequestEvent(ctx, vs.HandleValidatorExitRequestEvent); err != nil {
				log.Printf("Error scanning ValidatorExitRequest events: %v", err)
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
	message := fmt.Sprintf("- 🚨 One of the validators requested to exit: %s", validatorExitEvent.ValidatorIndex)
	if err := vs.notifierPort.SendNotification(message); err != nil {
		log.Printf("Failed to send notification: %v", err)
		return err
	}

	// validate the exit request
	if err := vs.exitValidatorPort.ExitValidator(string(validatorExitEvent.ValidatorPubkey), validatorExitEvent.ValidatorIndex.String()); err != nil {
		log.Printf("Failed to validate exit request: %v", err)
		return err
	}

	// notify the validator has been exited
	message = fmt.Sprintf("- 🚪 Validator %s has been exited", validatorExitEvent.ValidatorIndex)
	if err := vs.notifierPort.SendNotification(message); err != nil {
		log.Printf("Failed to send notification: %v", err)
		return err
	}

	// save the exit request
	exitRequests := domain.ExitRequests{
		validatorExitEvent.ValidatorIndex.String(): domain.ExitRequest{
			ValidatorPubkey: validatorExitEvent.ValidatorPubkey,
			Raw:             validatorExitEvent.Raw,
		},
	}

	if err := vs.storagePort.SaveExitRequests(exitRequests); err != nil {
		log.Printf("Failed to save exit requests: %v", err)
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
	report := domain.Report{
		Threshold:  "some-threshold",
		Validators: map[string]domain.ValidatorPerformance{},
	}
	epoch := "some-epoch"
	lidoReport := map[string]domain.Report{
		epoch: report,
	}
	if err := vs.storagePort.SaveLidoReport(lidoReport); err != nil {
		log.Printf("Failed to save Lido report: %v", err)
		return err
	}

	return nil
}
