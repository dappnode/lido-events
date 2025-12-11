package ports

import (
	"context"
)

type CsParametersPort interface {
	GetDefaultAllowedExitDelay(ctx context.Context) (uint64, error)
}
