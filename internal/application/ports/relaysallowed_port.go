package ports

import (
	"context"
	"lido-events/internal/application/domain"
)

type RelaysAllowedPort interface {
	GetRelaysAllowList(ctx context.Context) ([]domain.RelayAllowed, error)
}
