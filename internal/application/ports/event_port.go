package ports

import (
	"lido-events/internal/application/domain"

	"github.com/ethereum/go-ethereum/core/types"
)

//TODO: this uses a type defined from an external library!
type EventPort interface {
    ProcessEvent(eventName domain.EventName, vLog types.Log) error
}
