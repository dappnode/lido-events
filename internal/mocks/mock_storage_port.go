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

// RegisterOperatorIdListener returns a channel for operator ID updates
func (m *MockStoragePort) RegisterOperatorIdListener() chan []*big.Int {
	args := m.Called()
	return args.Get(0).(chan []*big.Int)
}

// GetOperatorPerformance returns a map of performance reports for an operator ID
func (m *MockStoragePort) GetOperatorPerformance(operatorID *big.Int, startEpoch, endEpoch string) (map[string]domain.Report, error) {
	args := m.Called(operatorID, startEpoch, endEpoch)
	return args.Get(0).(map[string]domain.Report), args.Error(1)
}

// SaveOperatorPerformance simulates saving performance data for an operator
func (m *MockStoragePort) SaveOperatorPerformance(operatorID *big.Int, epoch string, report domain.Report) error {
	args := m.Called(operatorID, epoch, report)
	return args.Error(0)
}

// GetLastProcessedEpoch returns the last processed epoch
func (m *MockStoragePort) GetLastProcessedEpoch() (uint64, error) {
	args := m.Called()
	return args.Get(0).(uint64), args.Error(1)
}

// SaveLastProcessedEpoch simulates saving the last processed epoch
func (m *MockStoragePort) SaveLastProcessedEpoch(epoch uint64) error {
	args := m.Called(epoch)
	return args.Error(0)
}

// GetExitRequests returns exit requests for a given operator
func (m *MockStoragePort) GetExitRequests(operatorID string) (map[string]domain.ExitRequest, error) {
	args := m.Called(operatorID)
	return args.Get(0).(map[string]domain.ExitRequest), args.Error(1)
}

// SaveExitRequests simulates saving a map of exit requests
func (m *MockStoragePort) SaveExitRequests(operatorID *big.Int, requests map[string]domain.ExitRequest) error {
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
