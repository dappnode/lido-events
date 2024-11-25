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
	GetPendingHashes() ([]string, error)
	DeletePendingHash(hash string) error

	// operator IDs
	GetOperatorIds() ([]*big.Int, error)
	SaveOperatorId(operatorID string) error
	RegisterOperatorIdListener() chan []*big.Int
	DeleteOperator(operatorID string) error

	// reports
	GetReports(operatorID *big.Int) (domain.Reports, error)
	SaveReport(operatorID *big.Int, report domain.Report) error

	// exit requests
	GetExitRequests(string) (domain.ExitRequests, error)
	SaveExitRequests(operatorID *big.Int, requests domain.ExitRequests) error
	SaveExitRequest(operatorID *big.Int, validatorIndex string, exitRequest domain.ExitRequest) error
	UpdateExitRequestStatus(operatorId string, validatorIndex string, status domain.ValidatorStatus) error
	DeleteExitRequest(operatorID string, validatorIndex string) error

	// telegram
	GetTelegramConfig() (domain.TelegramConfig, error)
	SaveTelegramConfig(config domain.TelegramConfig) error
	RegisterTelegramConfigListener() chan domain.TelegramConfig
}
