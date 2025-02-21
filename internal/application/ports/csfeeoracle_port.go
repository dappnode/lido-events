package ports

import (
	"context"
	"lido-events/internal/application/domain"
	"math/big"
)

type CsFeeOraclePort interface {
	ScanProcessingStartedEvents(
		ctx context.Context,
		start uint64,
		end *uint64,
		handleProcessingStarted func(*domain.BindingsProcessingStarted, *big.Int) error,
	) error
}
