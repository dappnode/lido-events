package ports

import (
	"context"
	"lido-events/internal/application/domain"
)

type Relays interface {
	GetRelaysUsed(ctx context.Context) ([]string, error)
	GetRelaysAllowList(ctx context.Context) ([]domain.RelayAllowed, error)
}
