package domain

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

// EventName is a custom type to represent event names.
type EventName string

// Define constants for each event name.
const (
	DepositedSigningKeysCountChanged         EventName = "DepositedSigningKeysCountChanged"         // CSModule
	ELRewardsStealingPenaltyReported         EventName = "ELRewardsStealingPenaltyReported"         // CSModule
	ELRewardsStealingPenaltySettled          EventName = "ELRewardsStealingPenaltySettled"          // CSModule
	ELRewardsStealingPenaltyCancelled        EventName = "ELRewardsStealingPenaltyCancelled"        // CSModule
	InitialSlashingSubmitted                 EventName = "InitialSlashingSubmitted"                 // CSModule
	KeyRemovalChargeApplied                  EventName = "KeyRemovalChargeApplied"                  // CSModule
	NodeOperatorManagerAddressChangeProposed EventName = "NodeOperatorManagerAddressChangeProposed" // CSModule
	NodeOperatorManagerAddressChanged        EventName = "NodeOperatorManagerAddressChanged"        // CSModule
	NodeOperatorRewardAddressChangeProposed  EventName = "NodeOperatorRewardAddressChangeProposed"  // CSModule
	NodeOperatorRewardAddressChanged         EventName = "NodeOperatorRewardAddressChanged"         // CSModule
	StuckSigningKeysCountChanged             EventName = "StuckSigningKeysCountChanged"             // CSModule
	VettedSigningKeysCountDecreased          EventName = "VettedSigningKeysCountDecreased"          // CSModule
	WithdrawalSubmitted                      EventName = "WithdrawalSubmitted"                      // CSModule
	TotalSigningKeysCountChanged             EventName = "TotalSigningKeysCountChanged"             // CSModule
	ValidatorExitRequest                     EventName = "ValidatorExitRequest"                     // VEBO
	PublicRelease                            EventName = "PublicRelease"                            // CSModule
	DistributionDataUpdated                  EventName = "DistributionDataUpdated"                  // CSFeeDistributor
	ReportSubmitted                          EventName = "ReportSubmitted"                          // VEBO
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
