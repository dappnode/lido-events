package domain

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

// VeboValidatorExitRequest represents a ValidatorExitRequest event raised by the Vebo contract.
type VeboValidatorExitRequest struct {
	StakingModuleId *big.Int
	NodeOperatorId  *big.Int
	ValidatorIndex  *big.Int
	ValidatorPubkey []byte
	Timestamp       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// VeboReportSubmitted represents a ReportSubmitted event raised by the Vebo contract.
type VeboReportSubmitted struct {
	RefSlot                *big.Int
	Hash                   [32]byte
	ProcessingDeadlineTime *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}
