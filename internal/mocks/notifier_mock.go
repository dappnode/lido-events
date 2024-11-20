package mocks

import (
	"github.com/stretchr/testify/mock"
)

// MockNotifierPort is a mock implementation of the NotifierPort interface using testify/mock
type MockNotifierPort struct {
	mock.Mock
}

// SendNotification simulates sending a notification
func (m *MockNotifierPort) SendNotification(message string) error {
	args := m.Called(message)
	return args.Error(0)
}
