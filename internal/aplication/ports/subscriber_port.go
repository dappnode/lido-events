// ports/subscriberPort.go
package ports

import (
	"context"
	"lido-events/internal/aplication/domain"
)

type SubscriberPort interface {
	WatchVeboEvents(ctx context.Context, handlers VeboHandlers) error
	WatchCsModuleEvents(ctx context.Context, handleLog func(logData interface{}) error) error
	WatchCsFeeDistributorEvents(ctx context.Context, handleLog func(logData interface{}) error) error
}

type VeboHandlers interface {
	HandleValidatorExitRequestEvent(reportSubmitted domain.VeboValidatorExitRequest) error
	HandleReportSubmittedEvent(reportSubmitted domain.VeboReportSubmitted) error
}
