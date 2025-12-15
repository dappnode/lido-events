package mocks

import (
	"math/big"

	"github.com/stretchr/testify/mock"
)

// MockNotifierPort is a mock implementation of the NotifierPort interface using testify/mock.
type MockNotifierPort struct {
	mock.Mock
}

// SendMissingLogReceiptsNotification simulates sending a missing log receipts notification.
func (m *MockNotifierPort) SendMissingLogReceiptsNotification(message string) error {
	args := m.Called(message)
	return args.Error(0)
}

// SendValidatorExitRequestedNotification simulates sending a validator exit requested notification.
func (m *MockNotifierPort) SendValidatorExitRequestedNotification(message string) error {
	args := m.Called(message)
	return args.Error(0)
}

// SendValidatorFailedExitNotification simulates sending a validator failed exit notification.
func (m *MockNotifierPort) SendValidatorFailedExitNotification(message string) error {
	args := m.Called(message)
	return args.Error(0)
}

// SendValidatorSucceedExitNotification simulates sending a validator succeeded exit notification.
func (m *MockNotifierPort) SendValidatorSucceedExitNotification(message string, validatorIndex *big.Int) error {
	args := m.Called(message, validatorIndex)
	return args.Error(0)
}

// SendBlackListedNotification simulates sending a blacklisted relays notification.
func (m *MockNotifierPort) SendBlackListedNotification(message string) error {
	args := m.Called(message)
	return args.Error(0)
}

// SendMissingRelayNotification simulates sending a missing mandatory relays notification.
func (m *MockNotifierPort) SendMissingRelayNotification(message string) error {
	args := m.Called(message)
	return args.Error(0)
}

// SendNewPerformanceReport simulates sending a new performance report notification.
func (m *MockNotifierPort) SendNewPerformanceReport(message string) error {
	args := m.Called(message)
	return args.Error(0)
}
