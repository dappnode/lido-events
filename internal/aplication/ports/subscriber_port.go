// ports/subscriberPort.go
package ports

import (
	"context"
)

// TODO: consider using a custom DecodedLog: handleLog(logData DecodedLog)
// type DecodedLog struct {
// 	Address     common.Address
// 	Topics      []common.Hash
// 	Data        []byte
// 	BlockNumber uint64
// 	TxHash      common.Hash
// 	EventData   map[string]interface{} // Decoded event fields
// }

type SubscriberPort interface {
	SubscribeToEvents(ctx context.Context, handleLog func(logData interface{}) error) error
}
