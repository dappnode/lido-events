package ports

import "lido-events/internal/application/domain"

type StoragePort interface {
	GetLidoReport(start, end string) (domain.Reports, error)
	SaveLidoReport(report domain.Reports) error
	GetExitRequests() (domain.ExitRequests, error)
	SaveExitRequests(requests domain.ExitRequests) error
	GetTelegramConfig() (domain.TelegramConfig, error)
	SaveTelegramConfig(config domain.TelegramConfig) error
	GetOperatorId() (domain.OperatorId, error)
	SaveOperatorId(operatorID domain.OperatorId) error
}
