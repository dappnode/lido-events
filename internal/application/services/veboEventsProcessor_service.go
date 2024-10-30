package services

import (
	"context"
	"fmt"
	"lido-events/internal/application/domain"
	"lido-events/internal/application/ports"
	"log"
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

// WatchVeboEvents subscribes to Ethereum events and handles them.
// It passes to the csModule port a function that processes the log data.
func (vs *VeboEventsProcessor) WatchVeboEvents(ctx context.Context) error {
	return vs.veboPort.WatchVeboEvents(ctx, vs)
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
	exitRequests := domain.ExitRequest{
		string(validatorExitEvent.ValidatorPubkey): string(validatorExitEvent.Raw.BlockNumber),
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
