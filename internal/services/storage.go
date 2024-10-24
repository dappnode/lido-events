package service

import (
	"errors"
	"lido-events/internal/domain/entities"
	"lido-events/internal/domain/repositories"
	"log"
)

type StorageService struct {
	storageService repositories.StorageRepository
}

// NewStorageService creates a new instance of OperatorService
func NewStorageService(storage repositories.StorageRepository) *StorageService {
	return &StorageService{
		storageService: storage,
	}
}

// GetLidoReport retrieves the Lido report for the given range of epochs
func (os *StorageService) GetLidoReport(start, end string) (map[string]entities.Report, error) {
	return os.storageService.GetLidoReport(start, end)
}

// SaveLidoReport saves the Lido report to the repository
func (os *StorageService) SaveLidoReport(report map[string]entities.Report) error {
	err := os.storageService.SaveLidoReport(report)
	if err != nil {
		return err
	}
	log.Printf("Lido report updated")

	return nil
}

// GetExitRequests retrieves the exit requests from the repository
func (os *StorageService) GetExitRequests() (entities.ExitRequest, error) {
	return os.storageService.GetExitRequests()
}

// SetExitRequests saves the exit requests to the repository
func (os *StorageService) SetExitRequests(exitRequests entities.ExitRequest) error {
	err := os.storageService.SaveExitRequests(exitRequests)
	if err != nil {
		return err
	}
	log.Printf("Exit requests updated")

	return nil
}

// GetTelegramConfig retrieves the Telegram configuration from the repository
func (os *StorageService) GetTelegramConfig() (entities.TelegramConfig, error) {
	return os.storageService.GetTelegramConfig()
}

// SetTelegramConfig validates and saves the Telegram configuration, then notifies
func (os *StorageService) SetTelegramConfig(config entities.TelegramConfig) error {
	if config.Token == "" {
		return errors.New("telegram token cannot be empty")
	}
	if config.ChatID == 0 {
		return errors.New("telegram chat id cannot be empty")
	}

	err := os.storageService.SaveTelegramConfig(config)
	if err != nil {
		return err
	}
	log.Printf("Telegram config updated")

	return nil
}

// GetOperatorId retrieves the operator ID from the repository
func (os *StorageService) GetOperatorId() (entities.OperatorId, error) {
	return os.storageService.GetOperatorId()
}

// SetOperatorId validates and sets the operator ID, then notifies
func (os *StorageService) SetOperatorId(operatorID entities.OperatorId) error {
	if operatorID == "" {
		return errors.New("operator id cannot be empty")
	}

	err := os.storageService.SaveOperatorId(operatorID)
	if err != nil {
		return err
	}
	log.Printf("Operator ID set to %s", operatorID)

	return nil
}
