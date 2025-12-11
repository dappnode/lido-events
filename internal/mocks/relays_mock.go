package mocks

import (
	"context"
	"lido-events/internal/application/domain"

	"github.com/stretchr/testify/mock"
)

// MockRelays is a mock implementation of the Relays port using testify/mock.
type MockRelays struct {
	mock.Mock
}

// GetRelaysUsed simulates retrieving the list of relays configured in the
// MEV Boost DNP via Dappmanager.
func (m *MockRelays) GetRelaysUsed(ctx context.Context) ([]string, error) {
	args := m.Called(ctx)
	return args.Get(0).([]string), args.Error(1)
}

// GetRelaysAllowList simulates retrieving the relays allow list from the
// smart contract.
func (m *MockRelays) GetRelaysAllowList(ctx context.Context) ([]domain.RelayAllowed, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.RelayAllowed), args.Error(1)
}
