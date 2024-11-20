package mocks

import (
	"context"
	"lido-events/internal/application/domain"

	"github.com/stretchr/testify/mock"
)

// MockVeboPort is a mock implementation of the VeboPort interface using testify/mock
type MockVeboPort struct {
	mock.Mock
}

// WatchReportSubmittedEvents simulates watching for ReportSubmitted events
func (m *MockVeboPort) WatchReportSubmittedEvents(
	ctx context.Context,
	handleReportSubmittedEvent func(*domain.VeboReportSubmitted) error,
) error {
	args := m.Called(ctx, handleReportSubmittedEvent)
	return args.Error(0)
}

// ScanVeboValidatorExitRequestEvent simulates scanning for ValidatorExitRequest events
func (m *MockVeboPort) ScanVeboValidatorExitRequestEvent(
	ctx context.Context,
	start uint64,
	end *uint64,
	handleExitRequestEvent func(*domain.VeboValidatorExitRequest) error,
) error {
	args := m.Called(ctx, start, end, handleExitRequestEvent)
	return args.Error(0)
}
