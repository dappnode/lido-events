package domain

import "github.com/ethereum/go-ethereum/core/types"

type BindingsDistributionLogUpdated struct {
	LogCid string
	Raw    types.Log // Blockchain specific contextual infos
}
