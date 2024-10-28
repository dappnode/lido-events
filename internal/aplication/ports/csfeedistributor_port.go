// ports/subscriberPort.go
package ports

import (
	"context"
	"lido-events/internal/aplication/domain"
)

type CsFeeDistributorPort interface {
	WatchCsFeeDistributorEvents(ctx context.Context, handlers CsFeeDistributorHandlers) error
}

type CsFeeDistributorHandlers interface {
	HandleDistributionDataUpdated(reportSubmitted *domain.CsfeedistributorDistributionDataUpdated) error
}
