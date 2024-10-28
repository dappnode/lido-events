package services

import (
	"context"
	"fmt"
	"lido-events/internal/aplication/domain"
	"lido-events/internal/aplication/ports"
	"log"
)

type VeboService struct {
	storagePort  ports.StoragePort
	notifierPort ports.NotifierPort
	veboPort     ports.VeboPort
}

func NewVeboService(storagePort ports.StoragePort, notifierPort ports.NotifierPort, veboPort ports.VeboPort) *VeboService {
	return &VeboService{
		storagePort:  storagePort,
		notifierPort: notifierPort,
		veboPort:     veboPort,
	}
}

// WatchVeboEvents subscribes to Ethereum events and handles them.
// It passes to the csModule port a function that processes the log data.
func (vs *VeboService) WatchVeboEvents(ctx context.Context) error {
	return vs.veboPort.WatchVeboEvents(ctx, vs)
}

// Make VeboService implement the VeboHandlers interface by adding the methods
func (vs *VeboService) HandleValidatorExitRequestEvent(reportSubmitted *domain.VeboValidatorExitRequest) error {
	message := fmt.Sprintf("- 🚨 One of the validators requested to exit: %s", reportSubmitted.ValidatorIndex)
	if err := vs.notifierPort.SendNotification(message); err != nil {
		log.Printf("Failed to send notification: %v", err)
		return err
	}

	// TODO: implement exit validator automatically and notify

	// save the exit request
	exitRequests := domain.ExitRequest{
		string(reportSubmitted.ValidatorPubkey): string(reportSubmitted.Raw.BlockNumber),
	}
	if err := vs.storagePort.SaveExitRequests(exitRequests); err != nil {
		log.Printf("Failed to save exit requests: %v", err)
		return err
	}
	return nil
}

func (vs *VeboService) HandleReportSubmittedEvent(reportSubmitted *domain.VeboReportSubmitted) error {
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
