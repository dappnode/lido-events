package services

import (
	"context"
	"fmt"
	"lido-events/internal/application/domain"
	"lido-events/internal/application/ports"
	"log"
)

type CsModuleService struct {
	notifierPort ports.NotifierPort
	storagePort  ports.StoragePort
	csModulePort ports.CsModulePort
}

func NewCsModuleService(storagePort ports.StoragePort, notifierPort ports.NotifierPort, csModulePort ports.CsModulePort) *CsModuleService {
	return &CsModuleService{
		notifierPort,
		storagePort,
		csModulePort,
	}
}

// WatchCsModuleEvents subscribes to Ethereum events and handles them.
// It passes to the csModule port a function that processes the log data.
func (csms *CsModuleService) WatchCsModuleEvents(ctx context.Context) error {
	return csms.csModulePort.WatchCsModuleEvents(ctx, csms)
}

// Make CsModuleService implement the CsModuleHandlers interface by adding the methods

func (cs *CsModuleService) HandleDepositedSigningKeysCountChanged(depositedSigningKeysCountChanged *domain.CsmoduleDepositedSigningKeysCountChanged) error {
	message := fmt.Sprintf("- 🤩 Node Operator's keys received depositst: %s", depositedSigningKeysCountChanged.DepositedKeysCount)
	if err := cs.notifierPort.SendNotification(message); err != nil {
		log.Printf("Failed to send notification: %v", err)
		return err
	}
	return nil
}

func (cs *CsModuleService) HandleElRewardsStealingPenaltyReported(eLRewardsStealingPenaltyReported *domain.CsmoduleELRewardsStealingPenaltyReported) error {
	message := fmt.Sprintf("- 🚨 Penalty for stealing EL rewards reported: %s", eLRewardsStealingPenaltyReported.StolenAmount)
	if err := cs.notifierPort.SendNotification(message); err != nil {
		log.Printf("Failed to send notification: %v", err)
		return err
	}
	return nil
}

func (cs *CsModuleService) HandleElRewardsStealingPenaltySettled(eLRewardsStealingPenaltySettled *domain.CsmoduleELRewardsStealingPenaltySettled) error {
	message := fmt.Sprintf("- 🚨 EL rewards stealing penalty confirmed and applied: %s", eLRewardsStealingPenaltySettled.NodeOperatorId)
	if err := cs.notifierPort.SendNotification(message); err != nil {
		log.Printf("Failed to send notification: %v", err)
		return err
	}
	return nil
}

func (cs *CsModuleService) HandleElRewardsStealingPenaltyCancelled(eLRewardsStealingPenaltyCancelled *domain.CsmoduleELRewardsStealingPenaltyCancelled) error {
	message := fmt.Sprintf("- 😮‍💨 Cancelled penalty for stealing EL reward: %s", eLRewardsStealingPenaltyCancelled.Amount)
	if err := cs.notifierPort.SendNotification(message); err != nil {
		log.Printf("Failed to send notification: %v", err)
		return err
	}
	return nil
}

func (cs *CsModuleService) HandleInitialSlashingSubmitted(initialSlashingSubmitted *domain.CsmoduleInitialSlashingSubmitted) error {
	message := fmt.Sprintf("- 🚨 Initial slashing submitted for one of the validators: %s", initialSlashingSubmitted.KeyIndex)
	if err := cs.notifierPort.SendNotification(message); err != nil {
		log.Printf("Failed to send notification: %v", err)
		return err
	}
	return nil
}

func (cs *CsModuleService) HandleKeyRemovalChargeApplied(keyRemovalChargeApplied *domain.CsmoduleKeyRemovalChargeApplied) error {
	message := fmt.Sprintf("- 🔑 Applied charge for key removal: %s", keyRemovalChargeApplied.NodeOperatorId)
	if err := cs.notifierPort.SendNotification(message); err != nil {
		log.Printf("Failed to send notification: %v", err)
		return err
	}
	return nil
}

func (cs *CsModuleService) HandleNodeOperatorManagerAddressChangeProposed(nodeOperatorManagerAddressChangeProposed *domain.CsmoduleNodeOperatorManagerAddressChangeProposed) error {
	message := fmt.Sprintf("- ℹ️ New manager address proposed: %s", nodeOperatorManagerAddressChangeProposed.NewProposedAddress)
	if err := cs.notifierPort.SendNotification(message); err != nil {
		log.Printf("Failed to send notification: %v", err)
		return err
	}
	return nil
}

func (cs *CsModuleService) HandleNodeOperatorManagerAddressChanged(nodeOperatorManagerAddressChanged *domain.CsmoduleNodeOperatorManagerAddressChanged) error {
	message := fmt.Sprintf("- ✅ Manager address changedt: %s", nodeOperatorManagerAddressChanged.NewAddress)
	if err := cs.notifierPort.SendNotification(message); err != nil {
		log.Printf("Failed to send notification: %v", err)
		return err
	}
	return nil
}

func (cs *CsModuleService) HandleNodeOperatorRewardAddressChangeProposed(nodeOperatorRewardAddressChangeProposed *domain.CsmoduleNodeOperatorRewardAddressChangeProposed) error {
	message := fmt.Sprintf("- ℹ️ New rewards address proposed: %s", nodeOperatorRewardAddressChangeProposed.NewProposedAddress)
	if err := cs.notifierPort.SendNotification(message); err != nil {
		log.Printf("Failed to send notification: %v", err)
		return err
	}

	return nil
}

func (cs *CsModuleService) HandleNodeOperatorRewardAddressChanged(nodeOperatorRewardAddressChanged *domain.CsmoduleNodeOperatorRewardAddressChanged) error {
	message := fmt.Sprintf("- ✅ Rewards address changed: %s", nodeOperatorRewardAddressChanged.NewAddress)
	if err := cs.notifierPort.SendNotification(message); err != nil {
		log.Printf("Failed to send notification: %v", err)
		return err
	}
	return nil
}

func (cs *CsModuleService) HandleStuckSigningKeysCountChanged(stuckSigningKeysCountChanged *domain.CsmoduleStuckSigningKeysCountChanged) error {
	message := fmt.Sprintf("- 🚨 Reported stuck keys that were not exited in time: %s", stuckSigningKeysCountChanged.StuckKeysCount)
	if err := cs.notifierPort.SendNotification(message); err != nil {
		log.Printf("Failed to send notification: %v", err)
		return err
	}
	return nil
}

func (cs *CsModuleService) HandleVettedSigningKeysCountDecreased(vettedSigningKeysCountDecreased *domain.CsmoduleVettedSigningKeysCountDecreased) error {
	message := fmt.Sprintf("- 🚨 Uploaded invalid keys: %s", vettedSigningKeysCountDecreased.NodeOperatorId)
	if err := cs.notifierPort.SendNotification(message); err != nil {
		log.Printf("Failed to send notification: %v", err)
		return err
	}
	return nil
}

func (cs *CsModuleService) HandleWithdrawalSubmitted(withdrawalSubmitted *domain.CsmoduleWithdrawalSubmitted) error {
	message := fmt.Sprintf("- 👀 Key withdrawal information submitted: %s", withdrawalSubmitted.Amount)
	if err := cs.notifierPort.SendNotification(message); err != nil {
		log.Printf("Failed to send notification: %v", err)
		return err
	}
	return nil
}

func (cs *CsModuleService) HandleTotalSigningKeysCountChanged(totalSigningKeysCountChanged *domain.CsmoduleTotalSigningKeysCountChanged) error {
	message := fmt.Sprintf("- 👀 New keys uploaded or removedt: %s", totalSigningKeysCountChanged.TotalKeysCount)
	if err := cs.notifierPort.SendNotification(message); err != nil {
		log.Printf("Failed to send notification: %v", err)
		return err
	}
	return nil
}

func (cs *CsModuleService) HandlePublicRelease(publicRelease *domain.CsmodulePublicRelease) error {
	message := fmt.Sprintf("- 🎉 Public release of CSM!: %s", publicRelease.Raw.TxHash)
	if err := cs.notifierPort.SendNotification(message); err != nil {
		log.Printf("Failed to send notification: %v", err)
		return err
	}
	return nil
}
