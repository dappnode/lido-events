package ports

import (
	"lido-events/internal/aplication/domain"

	"github.com/ethereum/go-ethereum/core/types" // TODO: should not import external libraries in domain layer
)

type EventPort interface {
	ProcessEvent(eventName domain.EventName, vLog types.Log) error
}
