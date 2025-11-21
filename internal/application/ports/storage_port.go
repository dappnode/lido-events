package ports

import (
	"lido-events/internal/application/domain"
	"math/big"
)

type StoragePort interface {
	// node operators
	GetOperatorIds() ([]*big.Int, error)
	SaveOperatorId(operatorID string) error
	DeleteOperator(operatorID string) error
	// reports
	SaveReport(operatorID *big.Int, report domain.Report) error
	GetReports(operatorID *big.Int) (domain.Reports, error)
	GetDistributionLogLastProcessedBlock(operatorID *big.Int) (uint64, error)
	SaveDistributionLogLastProcessedBlock(operatorID *big.Int, block uint64) error
	AddPendingHash(operatorID *big.Int, hash string) error
	DeletePendingHash(operatorID *big.Int, hash string) error
	GetPendingHashes(operatorID *big.Int) ([]string, error)
	// exit requests
	GetValidatorExitRequestLastProcessedBlock(operatorID *big.Int) (uint64, error)
	SaveValidatorExitRequestLastProcessedBlock(operatorID *big.Int, block uint64) error
	SaveExitRequest(operatorID *big.Int, validatorIndex string, exitRequest domain.ExitRequest) error
	GetExitRequests(operatorID *big.Int) (domain.ExitRequests, error)
	UpdateExitRequestStatus(operatorID *big.Int, validatorIndex string, status domain.ValidatorStatus) error
	DeleteExitRequest(operatorID *big.Int, validatorIndex string) error
}
