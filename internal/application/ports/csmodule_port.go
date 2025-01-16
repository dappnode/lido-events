package ports

import (
	"context"
	"lido-events/internal/application/domain"

	"github.com/ethereum/go-ethereum/common"
)

type CsModulePort interface {
	ScanNodeOperatorEvents(ctx context.Context, address common.Address, start uint64, end *uint64, handleNodeOperatorAddedEvent func(*domain.CsmoduleNodeOperatorAdded, common.Address) error, handleNodeOperatorManagerAddressChangedEvent func(*domain.CsmoduleNodeOperatorManagerAddressChanged, common.Address) error, handleNodeOperatorRewardAddressChangedEvent func(*domain.CsmoduleNodeOperatorRewardAddressChanged, common.Address) error) error
	WatchCsModuleEvents(ctx context.Context, handlers CsModuleHandlers) error
	ResubscribeSignal() <-chan struct{}
}

type CsModuleHandlers interface {
	HandleDepositedSigningKeysCountChanged(reportSubmitted *domain.CsmoduleDepositedSigningKeysCountChanged) error
	HandleElRewardsStealingPenaltyReported(reportSubmitted *domain.CsmoduleELRewardsStealingPenaltyReported) error
	HandleElRewardsStealingPenaltySettled(reportSubmitted *domain.CsmoduleELRewardsStealingPenaltySettled) error
	HandleElRewardsStealingPenaltyCancelled(reportSubmitted *domain.CsmoduleELRewardsStealingPenaltyCancelled) error
	HandleInitialSlashingSubmitted(reportSubmitted *domain.CsmoduleInitialSlashingSubmitted) error
	HandleKeyRemovalChargeApplied(reportSubmitted *domain.CsmoduleKeyRemovalChargeApplied) error
	HandleNodeOperatorManagerAddressChangeProposed(reportSubmitted *domain.CsmoduleNodeOperatorManagerAddressChangeProposed) error
	HandleNodeOperatorManagerAddressChanged(reportSubmitted *domain.CsmoduleNodeOperatorManagerAddressChanged) error
	HandleNodeOperatorRewardAddressChangeProposed(reportSubmitted *domain.CsmoduleNodeOperatorRewardAddressChangeProposed) error
	HandleNodeOperatorRewardAddressChanged(reportSubmitted *domain.CsmoduleNodeOperatorRewardAddressChanged) error
	HandleStuckSigningKeysCountChanged(reportSubmitted *domain.CsmoduleStuckSigningKeysCountChanged) error
	HandleVettedSigningKeysCountDecreased(reportSubmitted *domain.CsmoduleVettedSigningKeysCountDecreased) error
	HandleWithdrawalSubmitted(reportSubmitted *domain.CsmoduleWithdrawalSubmitted) error
	HandleTotalSigningKeysCountChanged(reportSubmitted *domain.CsmoduleTotalSigningKeysCountChanged) error
	HandlePublicRelease(reportSubmitted *domain.CsmodulePublicRelease) error
}
