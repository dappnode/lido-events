package services

import (
	"lido-events/internal/domain/entities"

	"github.com/ethereum/go-ethereum/core/types" // TODO: should not import in domain layer
)

type EventProcessor interface {
	ProcessEvent(eventName entities.EventName, vLog types.Log) error
}
