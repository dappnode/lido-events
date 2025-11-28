// ports/subscriberPort.go
package ports

import (
	"context"
)

type CsFeeDistributorPort interface {
	GetAllLogCids(ctx context.Context) ([]string, error)
}
