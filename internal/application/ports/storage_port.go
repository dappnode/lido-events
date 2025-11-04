package ports

import (
	"lido-events/internal/application/domain"
	"math/big"
)

type StoragePort interface {
	// node operators
	GetOperatorIds() ([]*big.Int, error)
	SaveOperatorId(operatorID string) error
	RegisterOperatorIdListener() chan []*big.Int
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
	// Withdrawals
	GetWithdrawalsSubmittedLastProcessedBlock(operatorID *big.Int) (uint64, error)
	SaveWithdrawalsSubmittedLastProcessedBlock(operatorID *big.Int, block uint64) error
	GetWithdrawals(operatorID *big.Int) ([]domain.CsmoduleWithdrawalSubmitted, error)
	SaveWithdrawal(operatorID *big.Int, withdrawal domain.CsmoduleWithdrawalSubmitted) error

	// Penalties
	GetElRewardsStealingPenaltiesReportedLastProcessedBlock() (uint64, error)
	SaveElRewardsStealingPenaltiesReportedLastProcessedBlock(block uint64) error
	GetElRewardsStealingPenaltiesReported() ([]domain.CsmoduleELRewardsStealingPenaltyReported, error)
	SaveElRewardsStealingPenaltyReported(penalty domain.CsmoduleELRewardsStealingPenaltyReported) error

	// telegram
	GetTelegramConfig() (domain.TelegramConfig, error)
	SaveTelegramConfig(config domain.TelegramConfig) error

	// ProcessingStarted events
	GetProcessingStartedLastProcessedBlock() (uint64, error)
	SaveProcessingStartedLastProcessedBlock(block uint64) error
	GetProcessingStartedEvents() ([]domain.BindingsProcessingStarted, error)
	SaveProcessingStartedEvent(event domain.BindingsProcessingStarted) error
}
