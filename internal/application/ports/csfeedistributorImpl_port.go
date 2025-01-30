// ports/subscriberPort.go
package ports

import (
	"context"
	"lido-events/internal/application/domain"
	"math/big"
)

type CsFeeDistributorImplPort interface {
	ScanDistributionLogUpdatedEvents(ctx context.Context, start uint64, end *uint64, operatorId *big.Int, handleDistributionLogUpdated func(event *domain.BindingsDistributionLogUpdated, operatorId *big.Int) error) error
}
