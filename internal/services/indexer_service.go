package services

import (
	"log"
	"strconv"

	"github.com/dappnode/lido-events/internal/domain"
)

type IndexerService struct {
	Storage  domain.Indexer // Should be a concrete storage implementation (e.g., FileStorage)
	Notifier domain.Notifier
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
	// Load performance data from the JSON file using LoadPerformance
	performanceData, err := s.Storage.LoadPerformance()
	if err != nil {
		return nil, err
	}

	// Filter performance data based on the start and end epochs
	reportData := make(map[string]domain.Report)
	for epoch, report := range performanceData {
		epochInt, err := strconv.Atoi(epoch)
		if err != nil {
			log.Printf("Invalid epoch format: %v", err) // Log the error
			continue
		}

		if epochInt >= start && epochInt <= end {
			reportData[epoch] = report
		}
	}

	return reportData, nil
}

func (s *IndexerService) GetExitRequests() (map[string]string, error) {
	// Load exit requests from the JSON file
	exitRequests, err := s.Storage.LoadExitRequests()
	if err != nil {
		return nil, err
	}

	return exitRequests, nil
}
