package ports

import (
	"context"
	"lido-events/internal/application/domain"
)

type CsModulePort interface {
	WatchCsModuleEvents(ctx context.Context, handlers CsModuleHandlers) error
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
