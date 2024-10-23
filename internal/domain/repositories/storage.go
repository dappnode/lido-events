package repositories

import "lido-events/internal/domain/entities"

type StorageRepository interface {
	GetLidoReport(start, end string) (entities.Reports, error)
	SaveLidoReport(report entities.Reports) error
	GetExitRequests() (entities.ExitRequest, error)
	SaveExitRequests(requests entities.ExitRequest) error
	GetTelegramConfig() (entities.TelegramConfig, error)
	SaveTelegramConfig(config entities.TelegramConfig) error
	GetOperatorId() (entities.OperatorId, error)
	SaveOperatorId(operatorID entities.OperatorId) error
}
