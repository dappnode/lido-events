package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
)

// MockCsParametersPort is a mock implementation of the CsParametersPort
// interface using testify/mock.
type MockCsParametersPort struct {
	mock.Mock
}

// GetDefaultAllowedExitDelay simulates retrieving the defaultAllowedExitDelay
// value from the CSParametersRegistry contract.
func (m *MockCsParametersPort) GetDefaultAllowedExitDelay(ctx context.Context) (uint64, error) {
	args := m.Called(ctx)
	return args.Get(0).(uint64), args.Error(1)
}
