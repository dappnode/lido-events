package service

import (
	"errors"
	"lido-events/internal/domain"
	"lido-events/internal/ports"
	"log"
)

type OperatorService struct {
	operator ports.Operator
}

// NewOperatorService creates a new instance of OperatorService
func NewOperatorService(operator ports.Operator) *OperatorService {
	return &OperatorService{
		operator: operator,
	}
}

// GetLidoReport retrieves the Lido report for the given range of epochs
func (os *OperatorService) GetLidoReport(start, end string) (map[string]domain.Report, error) {
	return os.operator.GetLidoReport(start, end)
}

// SaveLidoReport saves the Lido report to the repository
func (os *OperatorService) SaveLidoReport(report map[string]domain.Report) error {
	err := os.operator.SaveLidoReport(report)
	if err != nil {
		return err
	}
	log.Printf("Lido report updated")

	return nil
}

// GetExitRequests retrieves the exit requests from the repository
func (os *OperatorService) GetExitRequests() (map[string]string, error) {
	return os.operator.GetExitRequests()
}

// SetExitRequests saves the exit requests to the repository
func (os *OperatorService) SetExitRequests(exitRequests map[string]string) error {
	err := os.operator.SaveExitRequests(exitRequests)
	if err != nil {
		return err
	}
	log.Printf("Exit requests updated")

	return nil
}

// GetTelegramConfig retrieves the Telegram configuration from the repository
func (os *OperatorService) GetTelegramConfig() (domain.TelegramConfig, error) {
	return os.operator.GetTelegramConfig()
}

// SetTelegramConfig validates and saves the Telegram configuration, then notifies
func (os *OperatorService) SetTelegramConfig(config domain.TelegramConfig) error {
	if config.Token == "" {
		return errors.New("telegram token cannot be empty")
	}
	if config.ChatID == 0 {
		return errors.New("telegram chat id cannot be empty")
	}

	err := os.operator.SaveTelegramConfig(config)
	if err != nil {
		return err
	}
	log.Printf("Telegram config updated")

	return nil
}

// GetOperatorId retrieves the operator ID from the repository
func (os *OperatorService) GetOperatorId() (string, error) {
	return os.operator.GetOperatorId()
}

// SetOperatorId validates and sets the operator ID, then notifies
func (os *OperatorService) SetOperatorId(operatorID string) error {
	if operatorID == "" {
		return errors.New("operator id cannot be empty")
	}

	err := os.operator.SaveOperatorId(operatorID)
	if err != nil {
		return err
	}
	log.Printf("Operator ID set to %s", operatorID)

	return nil
}
