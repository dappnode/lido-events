package domain

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

type BindingsProcessingStarted struct {
	RefSlot *big.Int
	Hash    [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}
