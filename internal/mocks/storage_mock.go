package mocks

import (
	"lido-events/internal/application/domain"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/mock"
)

// MockStoragePort is a mock implementation of the StoragePort interface using testify/mock
type MockStoragePort struct {
	mock.Mock
}

// Node Operators
func (m *MockStoragePort) GetOperatorIds() ([]*big.Int, error) {
	args := m.Called()
	return args.Get(0).([]*big.Int), args.Error(1)
}

func (m *MockStoragePort) SaveOperatorId(operatorID string) error {
	args := m.Called(operatorID)
	return args.Error(0)
}

func (m *MockStoragePort) RegisterOperatorIdListener() chan []*big.Int {
	args := m.Called()
	return args.Get(0).(chan []*big.Int)
}

func (m *MockStoragePort) DeleteOperator(operatorID string) error {
	args := m.Called(operatorID)
	return args.Error(0)
}

// Reports
func (m *MockStoragePort) SaveReport(operatorID *big.Int, report domain.Report) error {
	args := m.Called(operatorID, report)
	return args.Error(0)
}

func (m *MockStoragePort) GetReports(operatorID *big.Int) (domain.Reports, error) {
	args := m.Called(operatorID)
	return args.Get(0).(domain.Reports), args.Error(1)
}

func (m *MockStoragePort) GetDistributionLogLastProcessedBlock(operatorID *big.Int) (uint64, error) {
	args := m.Called(operatorID)
	return args.Get(0).(uint64), args.Error(1)
}

func (m *MockStoragePort) SaveDistributionLogLastProcessedBlock(operatorID *big.Int, block uint64) error {
	args := m.Called(operatorID, block)
	return args.Error(0)
}

func (m *MockStoragePort) AddPendingHash(operatorID *big.Int, hash string) error {
	args := m.Called(operatorID, hash)
	return args.Error(0)
}

func (m *MockStoragePort) DeletePendingHash(operatorID *big.Int, hash string) error {
	args := m.Called(operatorID, hash)
	return args.Error(0)
}

func (m *MockStoragePort) GetPendingHashes(operatorID *big.Int) ([]string, error) {
	args := m.Called(operatorID)
	return args.Get(0).([]string), args.Error(1)
}

// Exit Requests
func (m *MockStoragePort) GetValidatorExitRequestLastProcessedBlock(operatorID *big.Int) (uint64, error) {
	args := m.Called(operatorID)
	return args.Get(0).(uint64), args.Error(1)
}

func (m *MockStoragePort) SaveValidatorExitRequestLastProcessedBlock(operatorID *big.Int, block uint64) error {
	args := m.Called(operatorID, block)
	return args.Error(0)
}

func (m *MockStoragePort) SaveExitRequest(operatorID *big.Int, validatorIndex string, exitRequest domain.ExitRequest) error {
	args := m.Called(operatorID, validatorIndex, exitRequest)
	return args.Error(0)
}

func (m *MockStoragePort) GetExitRequests(operatorID *big.Int) (domain.ExitRequests, error) {
	args := m.Called(operatorID)
	return args.Get(0).(domain.ExitRequests), args.Error(1)
}

func (m *MockStoragePort) UpdateExitRequestStatus(operatorID *big.Int, validatorIndex string, status domain.ValidatorStatus) error {
	args := m.Called(operatorID, validatorIndex, status)
	return args.Error(0)
}

func (m *MockStoragePort) DeleteExitRequest(operatorID *big.Int, validatorIndex string) error {
	args := m.Called(operatorID, validatorIndex)
	return args.Error(0)
}

// Withdrawals
func (m *MockStoragePort) GetElRewardsStealingPenaltiesReportedLastProcessedBlock() (uint64, error) {
	args := m.Called()
	return args.Get(0).(uint64), args.Error(1)
}

func (m *MockStoragePort) SaveElRewardsStealingPenaltiesReportedLastProcessedBlock(block uint64) error {
	args := m.Called(block)
	return args.Error(0)
}

func (m *MockStoragePort) GetElRewardsStealingPenaltiesReported() ([]domain.CsmoduleELRewardsStealingPenaltyReported, error) {
	args := m.Called()
	return args.Get(0).([]domain.CsmoduleELRewardsStealingPenaltyReported), args.Error(1)
}

func (m *MockStoragePort) SaveElRewardsStealingPenaltyReported(penalty domain.CsmoduleELRewardsStealingPenaltyReported) error {
	args := m.Called(penalty)
	return args.Error(0)
}

func (m *MockStoragePort) GetWithdrawalsSubmittedLastProcessedBlock(operatorID *big.Int) (uint64, error) {
	args := m.Called(operatorID)
	return args.Get(0).(uint64), args.Error(1)
}

func (m *MockStoragePort) SaveWithdrawalsSubmittedLastProcessedBlock(operatorID *big.Int, block uint64) error {
	args := m.Called(operatorID, block)
	return args.Error(0)
}

func (m *MockStoragePort) GetWithdrawals(operatorID *big.Int) ([]domain.CsmoduleWithdrawalSubmitted, error) {
	args := m.Called(operatorID)
	return args.Get(0).([]domain.CsmoduleWithdrawalSubmitted), args.Error(1)
}

func (m *MockStoragePort) SaveWithdrawal(operatorID *big.Int, withdrawal domain.CsmoduleWithdrawalSubmitted) error {
	args := m.Called(operatorID, withdrawal)
	return args.Error(0)
}

// Addresses
func (m *MockStoragePort) GetAddressLastProcessedBlock(address common.Address) (uint64, error) {
	args := m.Called(address)
	return args.Get(0).(uint64), args.Error(1)
}

func (m *MockStoragePort) SaveAddressLastProcessedBlock(address common.Address, block uint64) error {
	args := m.Called(address, block)
	return args.Error(0)
}

func (m *MockStoragePort) GetAddressEvents(address common.Address) (domain.AddressEvents, error) {
	args := m.Called(address)
	return args.Get(0).(domain.AddressEvents), args.Error(1)
}

func (m *MockStoragePort) SetNodeOperatorAdded(address common.Address, event domain.CsmoduleNodeOperatorAdded) error {
	args := m.Called(address, event)
	return args.Error(0)
}

func (m *MockStoragePort) SetNodeOperatorManagerAddressChanged(address common.Address, event domain.CsmoduleNodeOperatorManagerAddressChanged) error {
	args := m.Called(address, event)
	return args.Error(0)
}

func (m *MockStoragePort) SetNodeOperatorRewardAddressChanged(address common.Address, event domain.CsmoduleNodeOperatorRewardAddressChanged) error {
	args := m.Called(address, event)
	return args.Error(0)
}

func (m *MockStoragePort) SetNodeOperatorRewardAddressChangeProposed(address common.Address, event domain.CsmoduleNodeOperatorRewardAddressChangeProposed) error {
	args := m.Called(address, event)
	return args.Error(0)
}

func (m *MockStoragePort) SetNodeOperatorManagerAddressChangeProposed(address common.Address, event domain.CsmoduleNodeOperatorManagerAddressChangeProposed) error {
	args := m.Called(address, event)
	return args.Error(0)
}

// Telegram
func (m *MockStoragePort) GetTelegramConfig() (domain.TelegramConfig, error) {
	args := m.Called()
	return args.Get(0).(domain.TelegramConfig), args.Error(1)
}

func (m *MockStoragePort) SaveTelegramConfig(config domain.TelegramConfig) error {
	args := m.Called(config)
	return args.Error(0)
}

// ProcessingStarted

func (m *MockStoragePort) GetProcessingStartedLastProcessedBlock() (uint64, error) {
	args := m.Called()
	return args.Get(0).(uint64), args.Error(1)
}

func (m *MockStoragePort) SaveProcessingStartedLastProcessedBlock(block uint64) error {
	args := m.Called(block)
	return args.Error(0)
}

func (m *MockStoragePort) SaveProcessingStartedEvent(event domain.BindingsProcessingStarted) error {
	args := m.Called(event)
	return args.Error(0)
}

func (m *MockStoragePort) GetProcessingStartedEvents() ([]domain.BindingsProcessingStarted, error) {
	args := m.Called()
	return args.Get(0).([]domain.BindingsProcessingStarted), args.Error(1)
}
