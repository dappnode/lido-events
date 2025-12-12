package mocks

import (
	"lido-events/internal/application/domain"
	"math/big"

	"github.com/stretchr/testify/mock"
)

// MockExitsStorage is a mock implementation of the ExitsStorage interface using testify/mock.
type MockExitsStorage struct {
	mock.Mock
}

// Node Operators
func (m *MockExitsStorage) GetOperatorIds() ([]*big.Int, error) {
	args := m.Called()
	return args.Get(0).([]*big.Int), args.Error(1)
}

func (m *MockExitsStorage) SaveOperatorId(operatorID string) error {
	args := m.Called(operatorID)
	return args.Error(0)
}

func (m *MockExitsStorage) DeleteOperator(operatorID string) error {
	args := m.Called(operatorID)
	return args.Error(0)
}

// Exit Requests
func (m *MockExitsStorage) GetValidatorExitRequestLastProcessedBlock(operatorID *big.Int) (uint64, error) {
	args := m.Called(operatorID)
	return args.Get(0).(uint64), args.Error(1)
}

func (m *MockExitsStorage) SaveValidatorExitRequestLastProcessedBlock(operatorID *big.Int, block uint64) error {
	args := m.Called(operatorID, block)
	return args.Error(0)
}

func (m *MockExitsStorage) SaveExitRequest(operatorID *big.Int, validatorIndex string, exitRequest domain.ExitRequest) error {
	args := m.Called(operatorID, validatorIndex, exitRequest)
	return args.Error(0)
}

func (m *MockExitsStorage) GetExitRequests(operatorID *big.Int) (domain.ExitRequests, error) {
	args := m.Called(operatorID)
	return args.Get(0).(domain.ExitRequests), args.Error(1)
}

func (m *MockExitsStorage) UpdateExitRequestStatus(operatorID *big.Int, validatorIndex string, status domain.ValidatorStatus) error {
	args := m.Called(operatorID, validatorIndex, status)
	return args.Error(0)
}

func (m *MockExitsStorage) DeleteExitRequest(operatorID *big.Int, validatorIndex string) error {
	args := m.Called(operatorID, validatorIndex)
	return args.Error(0)
}
