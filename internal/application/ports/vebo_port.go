// ports/subscriberPort.go
package ports

import (
	"context"
	"lido-events/internal/application/domain"
)

type VeboPort interface {
	WatchReportSubmittedEvents(context.Context, func(*domain.VeboReportSubmitted) error) error
	ScanVeboValidatorExitRequestEvent(context.Context, uint64, *uint64, func(*domain.VeboValidatorExitRequest) error) error
}
