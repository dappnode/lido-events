package mocks

import (
	"github.com/stretchr/testify/mock"
)

// MockExecutionPort is a mock implementation of the ExecutionPort interface using testify/mock
type MockExecutionPort struct {
	mock.Mock
}

// GetMostRecentBlockNumber simulates retrieving the most recent block number
func (m *MockExecutionPort) GetMostRecentBlockNumber() (uint64, error) {
	args := m.Called()
	return args.Get(0).(uint64), args.Error(1)
}
