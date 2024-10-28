// ports/subscriberPort.go
package ports

import (
	"context"
	"lido-events/internal/application/domain"
)

type VeboPort interface {
	WatchVeboEvents(ctx context.Context, handlers VeboHandlers) error
}

type VeboHandlers interface {
	HandleValidatorExitRequestEvent(reportSubmitted *domain.VeboValidatorExitRequest) error
	HandleReportSubmittedEvent(reportSubmitted *domain.VeboReportSubmitted) error
}
