package ports

import (
	"lido-events/internal/application/domain"
	"math/big"
)

type StoragePort interface {
	// exit requests event
	GetValidatorExitRequestLastProcessedBlock() (uint64, error)
	SaveValidatorExitRequestLastProcessedBlock(block uint64) error

	// distributionLogUpdated event
	GetDistributionLogLastProcessedBlock() (uint64, error)
	SaveDistributionLogLastProcessedBlock(block uint64) error
	AddPendingHash(hash string) error
	DeletePendingHash(hash string) error

	// operator IDs
	GetOperatorIds() ([]*big.Int, error)
	SaveOperatorId(operatorID string) error
	RegisterOperatorIdListener() chan []*big.Int

	// performance
	GetOperatorPerformance(operatorID *big.Int, startEpoch, endEpoch string) (domain.Reports, error)
	SaveOperatorPerformance(operatorID *big.Int, epoch string, report domain.Report) error

	// exit requests
	GetExitRequests(string) (map[string]domain.ExitRequest, error)
	SaveExitRequests(operatorID *big.Int, requests map[string]domain.ExitRequest) error
	SaveExitRequest(operatorID *big.Int, validatorIndex string, exitRequest domain.ExitRequest) error
	UpdateExitRequestStatus(operatorId string, validatorIndex string, status domain.ValidatorStatus) error

	// telegram
	GetTelegramConfig() (domain.TelegramConfig, error)
	SaveTelegramConfig(config domain.TelegramConfig) error
	RegisterTelegramConfigListener() chan domain.TelegramConfig
}
