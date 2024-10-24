// ports/subscriberPort.go
package ports

import "context"

type SubscriberPort interface {
	SubscribeToEvents(ctx context.Context, handleLog func(logData interface{}) error) error
}
