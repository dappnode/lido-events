package ports

import "lido-events/internal/domain"

// Repository defines the methods for persisting and retrieving operator data
type Operator interface {
	SaveTelegramConfig(config domain.TelegramConfig) error
	GetTelegramConfig() (domain.TelegramConfig, error)
	SaveOperatorId(operatorID string) error
	GetOperatorId() (string, error)
	SaveLidoReport(report map[string]domain.Report) error
	GetLidoReport(start, end string) (map[string]domain.Report, error)
	SaveExitRequests(exitRequests map[string]string) error
	GetExitRequests() (map[string]string, error)
}
