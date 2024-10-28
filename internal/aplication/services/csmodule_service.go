package services

import (
	"context"
	"lido-events/internal/aplication/domain"
	"lido-events/internal/aplication/ports"
)

type CsModuleService struct {
	notifierPort ports.NotifierPort
	storagePort  ports.StoragePort
	csModule     ports.CsModulePort
}

func NewCsModuleService(storagePort ports.StoragePort, notifierPort ports.NotifierPort, csModulePort ports.CsModulePort) *CsModuleService {
	return &CsModuleService{
		notifierPort: notifierPort,
		storagePort:  storagePort,
		csModule:     csModulePort,
	}
}

// WatchCsModuleEvents subscribes to Ethereum events and handles them.
// It passes to the csModule port a function that processes the log data.
func (csms *CsModuleService) WatchCsModuleEvents(ctx context.Context) error {
	return csms.csModule.WatchCsModuleEvents(ctx, csms)
}

// Make CsModuleService implement the CsModuleHandlers interface by adding the methods

func (cs *CsModuleService) HandleDepositedSigningKeysCountChanged(reportSubmitted *domain.CsmoduleDepositedSigningKeysCountChanged) error {
	return nil
}

func (cs *CsModuleService) HandleElRewardsStealingPenaltyReported(reportSubmitted *domain.CsmoduleELRewardsStealingPenaltyReported) error {
	return nil
}

func (cs *CsModuleService) HandleElRewardsStealingPenaltySettled(reportSubmitted *domain.CsmoduleELRewardsStealingPenaltySettled) error {
	return nil
}

func (cs *CsModuleService) HandleElRewardsStealingPenaltyCancelled(reportSubmitted *domain.CsmoduleELRewardsStealingPenaltyCancelled) error {
	return nil
}

func (cs *CsModuleService) HandleInitialSlashingSubmitted(reportSubmitted *domain.CsmoduleInitialSlashingSubmitted) error {
	return nil
}

func (cs *CsModuleService) HandleKeyRemovalChargeApplied(reportSubmitted *domain.CsmoduleKeyRemovalChargeApplied) error {
	return nil
}

func (cs *CsModuleService) HandleNodeOperatorManagerAddressChangeProposed(reportSubmitted *domain.CsmoduleNodeOperatorManagerAddressChangeProposed) error {
	return nil
}

func (cs *CsModuleService) HandleNodeOperatorManagerAddressChanged(reportSubmitted *domain.CsmoduleNodeOperatorManagerAddressChanged) error {
	return nil
}

func (cs *CsModuleService) HandleNodeOperatorRewardAddressChangeProposed(reportSubmitted *domain.CsmoduleNodeOperatorRewardAddressChangeProposed) error {

	return nil
}

func (cs *CsModuleService) HandleNodeOperatorRewardAddressChanged(reportSubmitted *domain.CsmoduleNodeOperatorRewardAddressChanged) error {
	return nil
}

func (cs *CsModuleService) HandleStuckSigningKeysCountChanged(reportSubmitted *domain.CsmoduleStuckSigningKeysCountChanged) error {
	return nil
}

func (cs *CsModuleService) HandleVettedSigningKeysCountDecreased(reportSubmitted *domain.CsmoduleVettedSigningKeysCountDecreased) error {
	return nil
}

func (cs *CsModuleService) HandleWithdrawalSubmitted(reportSubmitted *domain.CsmoduleWithdrawalSubmitted) error {
	return nil
}

func (cs *CsModuleService) HandleTotalSigningKeysCountChanged(reportSubmitted *domain.CsmoduleTotalSigningKeysCountChanged) error {
	return nil
}

func (cs *CsModuleService) HandlePublicRelease(reportSubmitted *domain.CsmodulePublicRelease) error {
	return nil
}
