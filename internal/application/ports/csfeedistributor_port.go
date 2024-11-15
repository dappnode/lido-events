// ports/subscriberPort.go
package ports

import (
	"context"
	"lido-events/internal/application/domain"
)

type CsFeeDistributorPort interface {
	WatchCsFeeDistributorEvents(ctx context.Context, handleDistributionDataUpdated func(reportSubmitted *domain.CsfeedistributorDistributionDataUpdated) error) error
}
