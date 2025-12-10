package ports

import (
	"context"
	"math/big"
)

type CsParametersPort interface {
	GetDefaultAllowedExitDelay(ctx context.Context) (*big.Int, error)
}
