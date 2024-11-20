package mocks

import (
	"lido-events/internal/application/domain"

	"github.com/stretchr/testify/mock"
)

// MockBeaconchain is a mock implementation of the Beaconchain interface using testify/mock
type MockBeaconchain struct {
	mock.Mock
}

// GetValidatorStatus simulates retrieving the validator status for a given public key
func (m *MockBeaconchain) GetValidatorStatus(pubkey string) (domain.ValidatorStatus, error) {
	args := m.Called(pubkey)
	return args.Get(0).(domain.ValidatorStatus), args.Error(1)
}

// GetEpochHeader simulates retrieving the epoch header for a given block ID
func (m *MockBeaconchain) GetEpochHeader(blockID string) (uint64, error) {
	args := m.Called(blockID)
	return args.Get(0).(uint64), args.Error(1)
}
