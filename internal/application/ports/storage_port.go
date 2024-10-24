package ports

import "lido-events/internal/application/domain"

type StoragePort interface {
	GetLidoReport(start, end string) (domain.Reports, error)
	SaveLidoReport(report domain.Reports) error
	GetExitRequests() (domain.ExitRequest, error)
	SaveExitRequests(requests domain.ExitRequest) error
	GetTelegramConfig() (domain.TelegramConfig, error)
	SaveTelegramConfig(config domain.TelegramConfig) error
	GetOperatorId() (domain.OperatorId, error)
	SaveOperatorId(operatorID domain.OperatorId) error
}
