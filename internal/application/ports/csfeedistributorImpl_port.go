// ports/subscriberPort.go
package ports

import (
	"context"
	"lido-events/internal/application/domain"
)

type CsFeeDistributorImplPort interface {
	ScanDistributionLogUpdatedEvents(ctx context.Context, start uint64, end *uint64, handleDistributionLogUpdated func(*domain.BindingsDistributionLogUpdated) error) error
}
