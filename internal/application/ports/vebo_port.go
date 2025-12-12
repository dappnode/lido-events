// ports/subscriberPort.go
package ports

import (
	"context"
	"lido-events/internal/application/domain"
	"math/big"
)

type VeboPort interface {
	ScanVeboValidatorExitRequestEvent(ctx context.Context, operatorId *big.Int, start uint64, end *uint64, handleExitRequestEvent func(*domain.VeboValidatorExitRequest) error) error
}
