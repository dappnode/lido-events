package mocks

import (
	"lido-events/internal/application/domain"

	"github.com/stretchr/testify/mock"
)

// MockPerformanceStorage is a mock implementation of the PerformanceStorage
// interface using testify/mock.
type MockPerformanceStorage struct {
	mock.Mock
}

// SaveReport simulates storing a performance report for a given log CID.
func (m *MockPerformanceStorage) SaveReport(logCid string, report *domain.Report) error {
	args := m.Called(logCid, report)
	return args.Error(0)
}

// GetNoPerformance simulates retrieving all performance reports for a node
// operator ID.
func (m *MockPerformanceStorage) GetNoPerformance(noID string) ([]*domain.Report, error) {
	args := m.Called(noID)
	return args.Get(0).([]*domain.Report), args.Error(1)
}

// GetUniqueLogCids simulates retrieving all distinct log CIDs stored in the
// performance database.
func (m *MockPerformanceStorage) GetUniqueLogCids() ([]string, error) {
	args := m.Called()
	return args.Get(0).([]string), args.Error(1)
}
