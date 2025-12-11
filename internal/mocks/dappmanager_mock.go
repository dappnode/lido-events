package mocks

import (
	"context"
	"lido-events/internal/application/domain"

	"github.com/stretchr/testify/mock"
)

// MockDappManager is a mock implementation of the DappManager adapter.
// It is intended to be used in tests wherever a component depends on
// fetching notification settings from DappManager.
type MockDappManager struct {
	mock.Mock
}

// GetNotificationsEnabled simulates retrieving the notifications configuration
// from the DappManager API.
func (m *MockDappManager) GetNotificationsEnabled(ctx context.Context) (domain.LidoNotificationsEnabled, error) {
	args := m.Called(ctx)
	return args.Get(0).(domain.LidoNotificationsEnabled), args.Error(1)
}
