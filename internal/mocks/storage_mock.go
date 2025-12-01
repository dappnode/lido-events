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

func (m *MockStoragePort) GetWithdrawalsSubmittedLastProcessedBlock(operatorID *big.Int) (uint64, error) {
	args := m.Called(operatorID)
	return args.Get(0).(uint64), args.Error(1)
}

func (m *MockStoragePort) SaveWithdrawalsSubmittedLastProcessedBlock(operatorID *big.Int, block uint64) error {
	args := m.Called(operatorID, block)
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
