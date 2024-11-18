package ports

import "lido-events/internal/application/domain"

type Beaconchain interface {
	GetValidatorStatus(pubkey string) (domain.ValidatorStatus, error)
	GetEpochHeader(blockID string) (uint64, error)
}
