package ports

import "context"

type RelaysUsedPort interface {
	GetRelaysUsed(ctx context.Context) ([]string, error)
}
