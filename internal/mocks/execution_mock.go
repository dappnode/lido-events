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

// IsSyncing simulates checking whether the execution client is currently syncing.
func (m *MockExecutionPort) IsSyncing() (bool, error) {
	args := m.Called()
	return args.Get(0).(bool), args.Error(1)
}

// GetBlockHasReceipts simulates checking whether a given block has receipts
// available in the execution client.
func (m *MockExecutionPort) GetBlockHasReceipts(blockID string) (bool, error) {
	args := m.Called(blockID)
	return args.Get(0).(bool), args.Error(1)
}
