package services

import (
	"lido-events/internal/domain"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
)

type Service struct {
	Storage      domain.IndexerRepository // Core Storage interface for data management
	Notifier     domain.AlertService      // Separate Notifier interface for notifications
	EventMapping map[string]string        // Event descriptions
}

func (s *Service) UpdateTelegramToken(token string) error {
	err := s.Storage.UpdateTelegramToken(token)
	if err != nil {
		return err
	}

	if s.Notifier != nil {
		return s.Notifier.Notify("Telegram token updated.")
	}
	return nil
}

func (s *Service) UpdateOperatorID(operatorID string) error {
	err := s.Storage.UpdateOperatorID(operatorID)
	if err != nil {
		return err
	}

	if s.Notifier != nil {
		return s.Notifier.Notify("Operator ID updated.")
	}
	return nil
}

func (s *Service) GetLidoReport(start, end int) (map[string]domain.Report, error) {
	return s.Storage.GetLidoReport(start, end)
}

func (s *Service) GetExitRequests() (map[string]string, error) {
	return s.Storage.GetExitRequests()
}

// HandleEthereumEvent processes the Ethereum events
func (s *Service) HandleEthereumEvent(contractName string, eventName string, vLog types.Log) {
	// Check if the event has a known description
	if description, found := s.EventMapping[eventName]; found {
		log.Printf("Event from contract %s: %s - %s", contractName, eventName, description)

		// Notify via Telegram
		if s.Notifier != nil {
			err := s.Notifier.Notify(description)
			if err != nil {
				log.Printf("Failed to send notification: %v", err)
			}
		}

		// Store or process event data (based on the event type)
		s.processEvent(contractName, eventName, vLog)

	} else {
		log.Printf("Unknown event received from contract %s: %v", contractName, vLog.Topics[0])
	}
}

// processEvent handles event data storage or further processing
func (s *Service) processEvent(contractName string, eventName string, vLog types.Log) {
	// Example: handle ValidatorExitRequest event
	if eventName == "ValidatorExitRequest" {
		// Add logic to store exit request
		log.Printf("Processing exit request from contract: %s", contractName)
		// e.g., call s.IndexerRepo.SaveExitRequest(...)
	}
}
