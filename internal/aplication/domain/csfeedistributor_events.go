package domain

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

// CsfeedistributorDistributionDataUpdated represents a DistributionDataUpdated event raised by the Csfeedistributor contract.
type CsfeedistributorDistributionDataUpdated struct {
	TotalClaimableShares *big.Int
	TreeRoot             [32]byte
	TreeCid              string
	Raw                  types.Log // Blockchain specific contextual infos
}
