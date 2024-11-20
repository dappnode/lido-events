package mocks

import (
	"context"
	"lido-events/internal/application/domain"

	"github.com/stretchr/testify/mock"
)

// MockCsFeeDistributorImplPort is a mock implementation of the CsFeeDistributorImplPort interface using testify/mock
type MockCsFeeDistributorImplPort struct {
	mock.Mock
}

// ScanDistributionLogUpdatedEvents simulates scanning for DistributionLogUpdated events
func (m *MockCsFeeDistributorImplPort) ScanDistributionLogUpdatedEvents(
	ctx context.Context,
	start uint64,
	end *uint64,
	handleDistributionLogUpdated func(*domain.BindingsDistributionLogUpdated) error,
) error {
	args := m.Called(ctx, start, end, handleDistributionLogUpdated)
	return args.Error(0)
}
