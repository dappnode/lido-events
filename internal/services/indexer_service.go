package services

import (
	"log"
	"strconv"

	"lido-events/internal/domain"
)

type IndexerService struct {
	Storage  domain.Indexer  // Core Indexer interface for data management
	Notifier domain.Notifier // Separate Notifier interface for notifications
	Config   domain.Config
}

func (s *IndexerService) UpdateTelegramToken(token string) error {
	// Update the token in the storage
	err := s.Storage.UpdateTelegramToken(token)
	if err != nil {
		return err
	}

	// Notify via Telegram or other means if necessary
	if s.Notifier != nil {
		err = s.Notifier.Notify("Telegram token updated.")
		if err != nil {
			log.Printf("Failed to send notification: %v", err)
		}
	}
	return nil
}

func (s *IndexerService) UpdateOperatorID(operatorID string) error {
	// Update operator ID in storage
	err := s.Storage.UpdateOperatorID(operatorID)
	if err != nil {
		return err
	}

	// Notify if necessary
	if s.Notifier != nil {
		err = s.Notifier.Notify("Operator ID updated.")
		if err != nil {
			log.Printf("Failed to send notification: %v", err)
		}
	}
	return nil
}

func (s *IndexerService) GetLidoReport(start, end int) (map[string]domain.Report, error) {
	// Load all performance data from the storage layer
	performanceData, err := s.Storage.GetLidoReport(start, end)
	if err != nil {
		return nil, err
	}

	// Initialize an empty map to store the filtered reports
	reportData := make(map[string]domain.Report)

	// Assuming performanceData is a map[string]domain.Report
	// Loop through the performance data (epoch: string, report: domain.Report)
	for epoch, report := range performanceData {
		// Convert epoch (which is a string) to an integer
		epochInt, err := strconv.Atoi(epoch)
		if err != nil {
			log.Printf("Invalid epoch format: %v", err) // Log the error and skip this report
			continue
		}

		// Check if the epoch is within the range of start and end
		if epochInt >= start && epochInt <= end {
			// Add the report to the filtered results
			reportData[epoch] = report
		}
	}

	// Return the filtered reports within the epoch range
	return reportData, nil
}

func (s *IndexerService) GetExitRequests() (map[string]string, error) {
	// Load exit requests from the JSON file
	exitRequests, err := s.Storage.GetExitRequests()
	if err != nil {
		return nil, err
	}

	return exitRequests, nil
}
