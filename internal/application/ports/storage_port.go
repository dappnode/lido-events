package ports

import (
	"lido-events/internal/application/domain"
	"math/big"
)

type StoragePort interface {
	GetOperatorPerformance(operatorID *big.Int, startEpoch, endEpoch string) (map[string]domain.Report, error)
	SaveOperatorPerformance(operatorID *big.Int, epoch string, report domain.Report) error
	GetExitRequests(string) (map[string]domain.ExitRequest, error)
	SaveExitRequests(operatorID *big.Int, requests map[string]domain.ExitRequest) error
	SaveExitRequest(operatorID *big.Int, validatorIndex string, exitRequest domain.ExitRequest) error
	UpdateExitRequestStatus(operatorId string, validatorIndex string, status domain.ValidatorStatus) error
	GetLastProcessedEpoch() (uint64, error)
	SaveLastProcessedEpoch(epoch uint64) error
	GetTelegramConfig() (domain.TelegramConfig, error)
	SaveTelegramConfig(config domain.TelegramConfig) error
	GetOperatorIds() ([]*big.Int, error)
	SaveOperatorId(operatorID string) error
}
