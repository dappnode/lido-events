package mocks

import (
	"github.com/stretchr/testify/mock"
)

// MockExitValidator is a mock implementation of the ExitValidator interface using testify/mock
type MockExitValidator struct {
	mock.Mock
}

// ExitValidator simulates exiting a validator given its public key and index
func (m *MockExitValidator) ExitValidator(pubKey, validatorIndex string) error {
	args := m.Called(pubKey, validatorIndex)
	return args.Error(0)
}
