package domain

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Events types copied from sc bindings codegen autogenerated file
// DO NOT EDIT

type AddressEvents struct {
	LastProcessedBlock				   uint64 									  `json:"lastProcessedBlock"`
	NodeOperatorAdded                 []CsmoduleNodeOperatorAdded                 `json:"nodeOperatorAdded"`
	NodeOperatorManagerAddressChanged []CsmoduleNodeOperatorManagerAddressChanged `json:"nodeOperatorManagerAddressChanged"`
	NodeOperatorRewardAddressChanged  []CsmoduleNodeOperatorRewardAddressChanged  `json:"nodeOperatorRewardAddressChanged"`
}

// CsmoduleNodeOperatorAdded represents a NodeOperatorAdded event raised by the Csmodule contract.
type CsmoduleNodeOperatorAdded struct {
	NodeOperatorId *big.Int
	ManagerAddress common.Address
	RewardAddress  common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// CsmoduleDepositedSigningKeysCountChanged represents a DepositedSigningKeysCountChanged event raised by the Csmodule contract.
type CsmoduleDepositedSigningKeysCountChanged struct {
	NodeOperatorId     *big.Int
	DepositedKeysCount *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// CsmoduleELRewardsStealingPenaltyReported represents a ELRewardsStealingPenaltyReported event raised by the Csmodule contract.
type CsmoduleELRewardsStealingPenaltyReported struct {
	NodeOperatorId    *big.Int
	ProposedBlockHash [32]byte
	StolenAmount      *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// CsmoduleELRewardsStealingPenaltySettled represents a ELRewardsStealingPenaltySettled event raised by the Csmodule contract.
type CsmoduleELRewardsStealingPenaltySettled struct {
	NodeOperatorId *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// CsmoduleELRewardsStealingPenaltyCancelled represents a ELRewardsStealingPenaltyCancelled event raised by the Csmodule contract.
type CsmoduleELRewardsStealingPenaltyCancelled struct {
	NodeOperatorId *big.Int
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// CsmoduleInitialSlashingSubmitted represents a InitialSlashingSubmitted event raised by the Csmodule contract.
type CsmoduleInitialSlashingSubmitted struct {
	NodeOperatorId *big.Int
	KeyIndex       *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// CsmoduleKeyRemovalChargeApplied represents a KeyRemovalChargeApplied event raised by the Csmodule contract.
type CsmoduleKeyRemovalChargeApplied struct {
	NodeOperatorId *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// CsmoduleNodeOperatorManagerAddressChangeProposed represents a NodeOperatorManagerAddressChangeProposed event raised by the Csmodule contract.
type CsmoduleNodeOperatorManagerAddressChangeProposed struct {
	NodeOperatorId     *big.Int
	OldProposedAddress common.Address
	NewProposedAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// CsmoduleNodeOperatorManagerAddressChanged represents a NodeOperatorManagerAddressChanged event raised by the Csmodule contract.
type CsmoduleNodeOperatorManagerAddressChanged struct {
	NodeOperatorId *big.Int
	OldAddress     common.Address
	NewAddress     common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// CsmoduleNodeOperatorRewardAddressChangeProposed represents a NodeOperatorRewardAddressChangeProposed event raised by the Csmodule contract.
type CsmoduleNodeOperatorRewardAddressChangeProposed struct {
	NodeOperatorId     *big.Int
	OldProposedAddress common.Address
	NewProposedAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// CsmoduleNodeOperatorRewardAddressChanged represents a NodeOperatorRewardAddressChanged event raised by the Csmodule contract.
type CsmoduleNodeOperatorRewardAddressChanged struct {
	NodeOperatorId *big.Int
	OldAddress     common.Address
	NewAddress     common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// CsmoduleStuckSigningKeysCountChanged represents a StuckSigningKeysCountChanged event raised by the Csmodule contract.
type CsmoduleStuckSigningKeysCountChanged struct {
	NodeOperatorId *big.Int
	StuckKeysCount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// CsmoduleVettedSigningKeysCountDecreased represents a VettedSigningKeysCountDecreased event raised by the Csmodule contract.
type CsmoduleVettedSigningKeysCountDecreased struct {
	NodeOperatorId *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// CsmoduleWithdrawalSubmitted represents a WithdrawalSubmitted event raised by the Csmodule contract.
type CsmoduleWithdrawalSubmitted struct {
	NodeOperatorId *big.Int
	KeyIndex       *big.Int
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// CsmoduleTotalSigningKeysCountChanged represents a TotalSigningKeysCountChanged event raised by the Csmodule contract.
type CsmoduleTotalSigningKeysCountChanged struct {
	NodeOperatorId *big.Int
	TotalKeysCount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// CsmodulePublicRelease represents a PublicRelease event raised by the Csmodule contract.
type CsmodulePublicRelease struct {
	Raw types.Log // Blockchain specific contextual infos
}
