// services/eventService.go
package services

import (
	"context"
	"errors"
	"lido-events/internal/aplication/domain"
	"lido-events/internal/aplication/ports"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
)

type EventService struct {
	storagePort    ports.StoragePort
	notifierPort   ports.NotifierPort
	subscriberPort ports.SubscriberPort
}

func NewEventService(storagePort ports.StoragePort, notifierPort ports.NotifierPort, subscriberPort ports.SubscriberPort) *EventService {
	return &EventService{
		storagePort:    storagePort,
		notifierPort:   notifierPort,
		subscriberPort: subscriberPort,
	}
}

// SubscribeToEvents subscribes to Ethereum events and handles them.
// It passes to the subscriberPort a function that processes the log data.
func (es *EventService) SubscribeToEvents(ctx context.Context) error {
	return es.subscriberPort.SubscribeToEvents(ctx, es.handleLog)
}

// handleLog processes the log data and saves it.
func (es *EventService) handleLog(logData interface{}) error {
	vLog, ok := logData.(types.Log)
	if !ok {
		log.Printf("Failed to cast log data")
		return nil
	}

	// Process the event name from the log and then process it
	eventName, found := es.getEventNameFromLog(vLog)
	if !found {
		log.Printf("Event name not found")
		return nil
	}

	return es.ProcessEvent(domain.EventName(eventName), vLog)
}

// ProcessEvent processes a specific event by name
func (es *EventService) ProcessEvent(eventName domain.EventName, vLog types.Log) error {
	// Retrieve the message from domain model
	message := domain.NotificationMessageMap[eventName]
	if message == "" {
		log.Printf("No notification message defined for event: %s", eventName)
		return errors.New("no notification message defined")
	}

	// Send the notification
	if err := es.notifierPort.SendNotification(message); err != nil {
		log.Printf("Failed to send notification: %v", err)
		return err
	}
	switch eventName {
	case domain.ValidatorExitRequest:
		exitRequests := domain.ExitRequest{
			"validator1": "exit-request",
		}
		if err := es.storagePort.SaveExitRequests(exitRequests); err != nil {
			log.Printf("Failed to save exit requests: %v", err)
			return err
		}
	case domain.ReportSubmitted:
		report := domain.Report{
			Threshold:  "some-threshold",
			Validators: map[string]domain.ValidatorPerformance{},
		}
		epoch := "some-epoch"
		lidoReport := map[string]domain.Report{
			epoch: report,
		}
		if err := es.storagePort.SaveLidoReport(lidoReport); err != nil {
			log.Printf("Failed to save Lido report: %v", err)
			return err
		}
	}
	return nil
}

func (es *EventService) getEventNameFromLog(vLog types.Log) (string, bool) {
	// Extract the event name from the log topics using your ABI knowledge
	// This can involve parsing the log against the ABI you hold.
	return "EventName", true // Placeholder logic
}
