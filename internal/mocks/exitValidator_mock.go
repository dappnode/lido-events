package mocks

import "github.com/stretchr/testify/mock"

// MockExitValidator is a mock implementation of the ExitValidator
// port/interface using testify/mock.
type MockExitValidator struct {
	mock.Mock
}

// ExitValidator simulates orchestrating a voluntary exit for a validator.
func (m *MockExitValidator) ExitValidator(pubKey, validatorIndex string) error {
	args := m.Called(pubKey, validatorIndex)
	return args.Error(0)
}
