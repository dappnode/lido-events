package mocks

import (
	"lido-events/internal/application/domain"
	"math/big"

	"github.com/stretchr/testify/mock"
)

// MockStoragePort is a mock implementation of the StoragePort interface using testify/mock
type MockStoragePort struct {
	mock.Mock
}

// GetOperatorIds returns a list of operator IDs
func (m *MockStoragePort) GetOperatorIds() ([]*big.Int, error) {
	args := m.Called()
	return args.Get(0).([]*big.Int), args.Error(1)
}

// SaveOperatorId simulates saving an operator ID
func (m *MockStoragePort) SaveOperatorId(operatorID string) error {
	args := m.Called(operatorID)
	return args.Error(0)
}

// DeleteOperator simulates deleting an operator ID
func (m *MockStoragePort) DeleteOperator(operatorID string) error {
	args := m.Called(operatorID)
	return args.Error(0)
}

// RegisterOperatorIdListener returns a channel for operator ID updates
func (m *MockStoragePort) RegisterOperatorIdListener() chan []*big.Int {
	args := m.Called()
	return args.Get(0).(chan []*big.Int)
}

// GetReports returns a map of reports for an operator ID
func (m *MockStoragePort) GetReports(operatorID *big.Int) (domain.Reports, error) {
	args := m.Called(operatorID)
	return args.Get(0).(domain.Reports), args.Error(1)
}

// SaveReport simulates saving report data for an operator
func (m *MockStoragePort) SaveReport(operatorID *big.Int, report domain.Report) error {
	args := m.Called(operatorID, report)
	return args.Error(0)
}

// GetDistributionLogLastProcessedBlock returns the last processed block for the DistributionLogUpdated event
func (m *MockStoragePort) GetDistributionLogLastProcessedBlock() (uint64, error) {
	args := m.Called()
	return args.Get(0).(uint64), args.Error(1)
}

// SaveDistributionLogLastProcessedBlock simulates saving the last processed block for the DistributionLogUpdated event
func (m *MockStoragePort) SaveDistributionLogLastProcessedBlock(block uint64) error {
	args := m.Called(block)
	return args.Error(0)
}

// GetValidatorExitRequestLastProcessedBlock returns the last processed block for the ValidatorExitRequest event
func (m *MockStoragePort) GetValidatorExitRequestLastProcessedBlock() (uint64, error) {
	args := m.Called()
	return args.Get(0).(uint64), args.Error(1)
}

// SaveValidatorExitRequestLastProcessedBlock simulates saving the last processed block for the ValidatorExitRequest event
func (m *MockStoragePort) SaveValidatorExitRequestLastProcessedBlock(block uint64) error {
	args := m.Called(block)
	return args.Error(0)
}

// GetExitRequests returns exit requests for a given operator
func (m *MockStoragePort) GetExitRequests(operatorID string) (domain.ExitRequests, error) {
	args := m.Called(operatorID)
	return args.Get(0).(domain.ExitRequests), args.Error(1)
}

// SaveExitRequests simulates saving a map of exit requests
func (m *MockStoragePort) SaveExitRequests(operatorID *big.Int, requests domain.ExitRequests) error {
	args := m.Called(operatorID, requests)
	return args.Error(0)
}

// SaveExitRequest simulates saving an individual exit request
func (m *MockStoragePort) SaveExitRequest(operatorID *big.Int, validatorIndex string, exitRequest domain.ExitRequest) error {
	args := m.Called(operatorID, validatorIndex, exitRequest)
	return args.Error(0)
}

// UpdateExitRequestStatus simulates updating the status of an exit request
func (m *MockStoragePort) UpdateExitRequestStatus(operatorID string, validatorIndex string, status domain.ValidatorStatus) error {
	args := m.Called(operatorID, validatorIndex, status)
	return args.Error(0)
}

// DeleteExitRequest simulates deleting an individual exit request for a specific operator ID and validator index.
func (m *MockStoragePort) DeleteExitRequest(operatorID string, validatorIndex string) error {
	args := m.Called(operatorID, validatorIndex)
	return args.Error(0)
}

// GetTelegramConfig returns Telegram configuration data
func (m *MockStoragePort) GetTelegramConfig() (domain.TelegramConfig, error) {
	args := m.Called()
	return args.Get(0).(domain.TelegramConfig), args.Error(1)
}

// SaveTelegramConfig simulates saving Telegram configuration data
func (m *MockStoragePort) SaveTelegramConfig(config domain.TelegramConfig) error {
	args := m.Called(config)
	return args.Error(0)
}

// RegisterTelegramConfigListener returns a channel for Telegram config updates
func (m *MockStoragePort) RegisterTelegramConfigListener() chan domain.TelegramConfig {
	args := m.Called()
	return args.Get(0).(chan domain.TelegramConfig)
}

// AddPendingHash simulates adding a pending hash to DistributionLogUpdated.PendingHashes
func (m *MockStoragePort) AddPendingHash(hash string) error {
	args := m.Called(hash)
	return args.Error(0)
}

// DeletePendingHash simulates deleting a pending hash from DistributionLogUpdated.PendingHashes
func (m *MockStoragePort) DeletePendingHash(hash string) error {
	args := m.Called(hash)
	return args.Error(0)
}

// GetPendingHashes returns the list of pending hashes from DistributionLogUpdated
func (m *MockStoragePort) GetPendingHashes() ([]string, error) {
	args := m.Called()
	return args.Get(0).([]string), args.Error(1)
}

// GetCsModuleLastProcessedBlock returns the last processed block for CsModule events
func (m *MockStoragePort) GetCsModuleLastProcessedBlock() (uint64, error) {
	args := m.Called()
	return args.Get(0).(uint64), args.Error(1)
}

// SaveCsModuletLastProcessedBlock simulates saving the last processed block for CsModule events
func (m *MockStoragePort) SaveCsModuletLastProcessedBlock(block uint64) error {
	args := m.Called(block)
	return args.Error(0)
}

// GetNodeOperatorAdded returns the NodeOperatorAdded event for a specific operator ID
func (m *MockStoragePort) GetNodeOperatorAdded(operatorID string) ([]domain.CsmoduleNodeOperatorAdded, error) {
	args := m.Called(operatorID)
	return args.Get(0).([]domain.CsmoduleNodeOperatorAdded), args.Error(1)
}

// SetNodeOperatorAdded simulates saving a NodeOperatorAdded event for a specific operator ID
func (m *MockStoragePort) SetNodeOperatorAdded(operatorID string, event domain.CsmoduleNodeOperatorAdded) error {
	args := m.Called(operatorID, event)
	return args.Error(0)
}

// GetNodeOperatorManagerAddressChanged returns the NodeOperatorManagerAddressChanged event for a specific operator ID
func (m *MockStoragePort) GetNodeOperatorManagerAddressChanged(operatorID string) ([]domain.CsmoduleNodeOperatorManagerAddressChanged, error) {
	args := m.Called(operatorID)
	return args.Get(0).([]domain.CsmoduleNodeOperatorManagerAddressChanged), args.Error(1)
}

// SetNodeOperatorManagerAddressChanged simulates saving a NodeOperatorManagerAddressChanged event for a specific operator ID
func (m *MockStoragePort) SetNodeOperatorManagerAddressChanged(operatorID string, event domain.CsmoduleNodeOperatorManagerAddressChanged) error {
	args := m.Called(operatorID, event)
	return args.Error(0)
}

// GetNodeOperatorRewardAddressChanged returns the NodeOperatorRewardAddressChanged event for a specific operator ID
func (m *MockStoragePort) GetNodeOperatorRewardAddressChanged(operatorID string) ([]domain.CsmoduleNodeOperatorRewardAddressChanged, error) {
	args := m.Called(operatorID)
	return args.Get(0).([]domain.CsmoduleNodeOperatorRewardAddressChanged), args.Error(1)
}

// SetNodeOperatorRewardAddressChanged simulates saving a NodeOperatorRewardAddressChanged event for a specific operator ID
func (m *MockStoragePort) SetNodeOperatorRewardAddressChanged(operatorID string, event domain.CsmoduleNodeOperatorRewardAddressChanged) error {
	args := m.Called(operatorID, event)
	return args.Error(0)
}
