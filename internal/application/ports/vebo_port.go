// ports/subscriberPort.go
package ports

import (
	"context"
	"lido-events/internal/application/domain"
)

type VeboPort interface {
	WatchReportSubmittedEvents(ctx context.Context, handleReportSubmittedEvent func(*domain.VeboReportSubmitted) error) error
	ScanVeboValidatorExitRequestEvent(ctx context.Context, start uint64, end *uint64, handleExitRequestEvent func(*domain.VeboValidatorExitRequest) error) error
}
