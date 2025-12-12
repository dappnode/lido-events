package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
)

// MockCsFeeDistributor is a mock implementation of the CsFeeDistributorPort
// interface using testify/mock.
type MockCsFeeDistributor struct {
	mock.Mock
}

// GetAllLogCids simulates retrieving all historical log CIDs from the
// CsFeeDistributor contract.
func (m *MockCsFeeDistributor) GetAllLogCids(ctx context.Context) ([]string, error) {
	args := m.Called(ctx)
	return args.Get(0).([]string), args.Error(1)
}
