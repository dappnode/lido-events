package services

import (
	"errors"
	"lido-events/internal/application/domain"
	"lido-events/internal/application/ports"

	"log"
)

type StorageService struct {
	storagePort ports.StoragePort
	notiferPort ports.NotifierPort
}

// NewStorageService creates a new instance of OperatorService
func NewStorageService(storage ports.StoragePort) *StorageService {
	return &StorageService{
		storagePort: storage,
	}
}

// GetLidoReport retrieves the Lido report for the given range of epochs
func (os *StorageService) GetLidoReport(start, end string) (map[string]domain.Report, error) {
	return os.storagePort.GetLidoReport(start, end)
}

// SaveLidoReport saves the Lido report to the repository
func (os *StorageService) SaveLidoReport(report map[string]domain.Report) error {
	err := os.storagePort.SaveLidoReport(report)
	if err != nil {
		return err
	}
	log.Printf("Lido report updated")

	return nil
}

// GetExitRequests retrieves the exit requests from the repository
func (os *StorageService) GetExitRequests() (domain.ExitRequest, error) {
	return os.storagePort.GetExitRequests()
}

// SetExitRequests saves the exit requests to the repository
func (os *StorageService) SetExitRequests(exitRequests domain.ExitRequest) error {
	err := os.storagePort.SaveExitRequests(exitRequests)
	if err != nil {
		return err
	}
	log.Printf("Exit requests updated")

	return nil
}

// GetTelegramConfig retrieves the Telegram configuration from the repository
func (os *StorageService) GetTelegramConfig() (domain.TelegramConfig, error) {
	return os.storagePort.GetTelegramConfig()
}

// SetTelegramConfig validates and saves the Telegram configuration, then notifies
func (os *StorageService) SetTelegramConfig(config domain.TelegramConfig) error {
	if config.Token == "" {
		return errors.New("telegram token cannot be empty")
	}
	if config.ChatID == 0 {
		return errors.New("telegram chat id cannot be empty")
	}

	err := os.storagePort.SaveTelegramConfig(config)
	if err != nil {
		return err
	}

	err = os.notiferPort.SendNotification("Telegram configuration updated")
	if err != nil {
		return err
	}

	log.Printf("Telegram config updated")

	return nil
}

// GetOperatorId retrieves the operator ID from the repository
func (os *StorageService) GetOperatorId() (domain.OperatorId, error) {
	return os.storagePort.GetOperatorId()
}

// SetOperatorId validates and sets the operator ID, then notifies
func (os *StorageService) SetOperatorId(operatorID domain.OperatorId) error {
	err := os.storagePort.SaveOperatorId(operatorID)
	if err != nil {
		return err
	}

	err = os.notiferPort.SendNotification("Operator ID updated")
	if err != nil {
		return err
	}

	log.Printf("Operator ID updated %v", operatorID)

	return nil
}
