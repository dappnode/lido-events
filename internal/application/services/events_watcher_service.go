package services

import (
	"context"
	"fmt"
	"lido-events/internal/application/domain"
	"lido-events/internal/application/ports"
	"lido-events/internal/logger"
	"sync"
)

type EventsWatcher struct {
	veboPort             ports.VeboPort
	csModulePort         ports.CsModulePort
	csFeeDistributorPort ports.CsFeeDistributorPort
	notifierPort         ports.NotifierPort
	servicePrefix        string
}

func NewEventsWatcherService(veboPort ports.VeboPort, csModulePort ports.CsModulePort, csFeeDistributorPort ports.CsFeeDistributorPort, notifierPort ports.NotifierPort) *EventsWatcher {
	return &EventsWatcher{
		veboPort,
		csModulePort,
		csFeeDistributorPort,
		notifierPort,
		"EventsWatcher",
	}
}

func (ew *EventsWatcher) WatchAllEvents(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	logger.InfoWithPrefix(ew.servicePrefix, "Starting to watch all events...")

	// Watch each event type in separate loops
	go func() {
		if err := ew.watchCsModuleEvents(ctx); err != nil {
			logger.ErrorWithPrefix(ew.servicePrefix, fmt.Sprintf("Failed to subscribe to CSModule events: %v", err))
		}
	}()

	go func() {
		if err := ew.watchCsFeeDistributorEvents(ctx); err != nil {
			logger.ErrorWithPrefix(ew.servicePrefix, fmt.Sprintf("Failed to subscribe to CsFeeDistributor events: %v", err))
		}
	}()

	go func() {
		if err := ew.watchReportSubmittedEvents(ctx); err != nil {
			logger.ErrorWithPrefix(ew.servicePrefix, fmt.Sprintf("Failed to subscribe to Vebo events: %v", err))
		}
	}()

	// Block until context is done
	<-ctx.Done()
	logger.InfoWithPrefix(ew.servicePrefix, "Context canceled, stopping all event watchers.")
}

func (ew *EventsWatcher) watchReportSubmittedEvents(ctx context.Context) error {
	logger.InfoWithPrefix(ew.servicePrefix, "Starting to watch Vebo ReportSubmitted events")
	return ew.veboPort.WatchReportSubmittedEvents(ctx, ew.HandleReportSubmittedEvent)
}

func (ew *EventsWatcher) watchCsModuleEvents(ctx context.Context) error {
	logger.InfoWithPrefix(ew.servicePrefix, "Starting to watch CsModule events...")

	// Listen for events in a loop to handle dynamic resubscriptions
	for {
		err := ew.csModulePort.WatchCsModuleEvents(ctx, ew)
		if err != nil {
			logger.ErrorWithPrefix(ew.servicePrefix, fmt.Sprintf("Error watching CsModule events: %v", err))
		}

		select {
		case <-ctx.Done():
			// Stop watching events if the context is canceled
			logger.InfoWithPrefix(ew.servicePrefix, "Context canceled, stopping CsModule events watcher.")
			return nil
		case <-ew.csModulePort.ResubscribeSignal():
			// Operator IDs updated, restart watching events
			logger.InfoWithPrefix(ew.servicePrefix, "Operator IDs updated, restarting CsModule events watcher...")
		}
	}
}

func (ew *EventsWatcher) watchCsFeeDistributorEvents(ctx context.Context) error {
	logger.InfoWithPrefix(ew.servicePrefix, "Starting to watch CsFeeDistributor events")
	return ew.csFeeDistributorPort.WatchCsFeeDistributorEvents(ctx, ew.HandleDistributionDataUpdated)
}

// Handlers Vebo

func (ew *EventsWatcher) HandleReportSubmittedEvent(reportSubmitted *domain.VeboReportSubmitted) error {
	// send the notification message
	message := fmt.Sprintf("- 📈 New submitted report: %s", reportSubmitted.RefSlot)
	if err := ew.notifierPort.SendNotification(message); err != nil {
		logger.ErrorWithPrefix(ew.servicePrefix, "Error sending reportSubmitted event notification", err)
		return err
	}

	return nil
}

// Handlers csFeeDistributor

func (ew *EventsWatcher) HandleDistributionDataUpdated(rewardsDistributed *domain.CsfeedistributorDistributionDataUpdated) error {
	// send the notification message
	message := fmt.Sprintf("- 📈 New rewards distributed: %s", rewardsDistributed.TotalClaimableShares)
	if err := ew.notifierPort.SendNotification(message); err != nil {
		logger.ErrorWithPrefix(ew.servicePrefix, "Error sending rewardsDistributed event notification", err)
		return err
	}

	return nil
}

// Handlers CSModule

func (ew *EventsWatcher) HandleDepositedSigningKeysCountChanged(depositedSigningKeysCountChanged *domain.CsmoduleDepositedSigningKeysCountChanged) error {
	message := fmt.Sprintf("- 🤩 Node Operator's keys received depositst: %s", depositedSigningKeysCountChanged.DepositedKeysCount)
	if err := ew.notifierPort.SendNotification(message); err != nil {
		logger.ErrorWithPrefix(ew.servicePrefix, "Error sending depositedSigningKeysCountChanged event notification", err)
		return err
	}
	return nil
}

func (ew *EventsWatcher) HandleElRewardsStealingPenaltyReported(eLRewardsStealingPenaltyReported *domain.CsmoduleELRewardsStealingPenaltyReported) error {
	message := fmt.Sprintf("- 🚨 Penalty for stealing EL rewards reported: %s", eLRewardsStealingPenaltyReported.StolenAmount)
	if err := ew.notifierPort.SendNotification(message); err != nil {
		logger.ErrorWithPrefix(ew.servicePrefix, "Error sending eLRewardsStealingPenaltyReported event notification", err)
		return err
	}
	return nil
}

func (ew *EventsWatcher) HandleElRewardsStealingPenaltySettled(eLRewardsStealingPenaltySettled *domain.CsmoduleELRewardsStealingPenaltySettled) error {
	message := fmt.Sprintf("- 🚨 EL rewards stealing penalty confirmed and applied: %s", eLRewardsStealingPenaltySettled.NodeOperatorId)
	if err := ew.notifierPort.SendNotification(message); err != nil {
		logger.ErrorWithPrefix(ew.servicePrefix, "Error sending eLRewardsStealingPenaltySettled event notification", err)
		return err
	}
	return nil
}

func (ew *EventsWatcher) HandleElRewardsStealingPenaltyCancelled(eLRewardsStealingPenaltyCancelled *domain.CsmoduleELRewardsStealingPenaltyCancelled) error {
	message := fmt.Sprintf("- 😮‍💨 Cancelled penalty for stealing EL reward: %s", eLRewardsStealingPenaltyCancelled.Amount)
	if err := ew.notifierPort.SendNotification(message); err != nil {
		logger.ErrorWithPrefix(ew.servicePrefix, "Error sending eLRewardsStealingPenaltyCancelled event notification", err)
		return err
	}
	return nil
}

func (ew *EventsWatcher) HandleInitialSlashingSubmitted(initialSlashingSubmitted *domain.CsmoduleInitialSlashingSubmitted) error {
	message := fmt.Sprintf("- 🚨 Initial slashing submitted for one of the validators: %s", initialSlashingSubmitted.KeyIndex)
	if err := ew.notifierPort.SendNotification(message); err != nil {
		logger.ErrorWithPrefix(ew.servicePrefix, "Error sending initialSlashingSubmitted event notification", err)
		return err
	}
	return nil
}

func (ew *EventsWatcher) HandleKeyRemovalChargeApplied(keyRemovalChargeApplied *domain.CsmoduleKeyRemovalChargeApplied) error {
	message := fmt.Sprintf("- 🔑 Applied charge for key removal: %s", keyRemovalChargeApplied.NodeOperatorId)
	if err := ew.notifierPort.SendNotification(message); err != nil {
		logger.ErrorWithPrefix(ew.servicePrefix, "Error sending keyRemovalChargeApplied event notification", err)
		return err
	}
	return nil
}

func (ew *EventsWatcher) HandleNodeOperatorManagerAddressChangeProposed(nodeOperatorManagerAddressChangeProposed *domain.CsmoduleNodeOperatorManagerAddressChangeProposed) error {
	message := fmt.Sprintf("- ℹ️ New manager address proposed: %s", nodeOperatorManagerAddressChangeProposed.NewProposedAddress)
	if err := ew.notifierPort.SendNotification(message); err != nil {
		logger.ErrorWithPrefix(ew.servicePrefix, "Error sending nodeOperatorManagerAddressChangeProposed event notification", err)
		return err
	}
	return nil
}

func (ew *EventsWatcher) HandleNodeOperatorManagerAddressChanged(nodeOperatorManagerAddressChanged *domain.CsmoduleNodeOperatorManagerAddressChanged) error {
	message := fmt.Sprintf("- ✅ Manager address changedt: %s", nodeOperatorManagerAddressChanged.NewAddress)
	if err := ew.notifierPort.SendNotification(message); err != nil {
		logger.ErrorWithPrefix(ew.servicePrefix, "Error sending nodeOperatorManagerAddressChanged event notification", err)
		return err
	}
	return nil
}

func (ew *EventsWatcher) HandleNodeOperatorRewardAddressChangeProposed(nodeOperatorRewardAddressChangeProposed *domain.CsmoduleNodeOperatorRewardAddressChangeProposed) error {
	message := fmt.Sprintf("- ℹ️ New rewards address proposed: %s", nodeOperatorRewardAddressChangeProposed.NewProposedAddress)
	if err := ew.notifierPort.SendNotification(message); err != nil {
		logger.ErrorWithPrefix(ew.servicePrefix, "Error sending nodeOperatorRewardAddressChangeProposed event notification", err)
		return err
	}

	return nil
}

func (ew *EventsWatcher) HandleNodeOperatorRewardAddressChanged(nodeOperatorRewardAddressChanged *domain.CsmoduleNodeOperatorRewardAddressChanged) error {
	message := fmt.Sprintf("- ✅ Rewards address changed: %s", nodeOperatorRewardAddressChanged.NewAddress)
	if err := ew.notifierPort.SendNotification(message); err != nil {
		logger.ErrorWithPrefix(ew.servicePrefix, "Error sending nodeOperatorRewardAddressChanged event notification", err)
		return err
	}
	return nil
}

func (ew *EventsWatcher) HandleStuckSigningKeysCountChanged(stuckSigningKeysCountChanged *domain.CsmoduleStuckSigningKeysCountChanged) error {
	message := fmt.Sprintf("- 🚨 Reported stuck keys that were not exited in time: %s", stuckSigningKeysCountChanged.StuckKeysCount)
	if err := ew.notifierPort.SendNotification(message); err != nil {
		logger.ErrorWithPrefix(ew.servicePrefix, "Error sending stuckSigningKeysCountChanged event notification", err)
		return err
	}
	return nil
}

func (ew *EventsWatcher) HandleVettedSigningKeysCountDecreased(vettedSigningKeysCountDecreased *domain.CsmoduleVettedSigningKeysCountDecreased) error {
	message := fmt.Sprintf("- 🚨 Uploaded invalid keys: %s", vettedSigningKeysCountDecreased.NodeOperatorId)
	if err := ew.notifierPort.SendNotification(message); err != nil {
		logger.ErrorWithPrefix(ew.servicePrefix, "Error sending vettedSigningKeysCountDecreased event notification", err)
		return err
	}
	return nil
}

func (ew *EventsWatcher) HandleWithdrawalSubmitted(withdrawalSubmitted *domain.CsmoduleWithdrawalSubmitted) error {
	message := fmt.Sprintf("- 👀 Key withdrawal information submitted: %s", withdrawalSubmitted.Amount)
	if err := ew.notifierPort.SendNotification(message); err != nil {
		logger.ErrorWithPrefix(ew.servicePrefix, "Error sending withdrawalSubmitted event notification", err)
		return err
	}
	return nil
}

func (ew *EventsWatcher) HandleTotalSigningKeysCountChanged(totalSigningKeysCountChanged *domain.CsmoduleTotalSigningKeysCountChanged) error {
	message := fmt.Sprintf("- 👀 New keys uploaded or removedt: %s", totalSigningKeysCountChanged.TotalKeysCount)
	if err := ew.notifierPort.SendNotification(message); err != nil {
		logger.ErrorWithPrefix(ew.servicePrefix, "Error sending totalSigningKeysCountChanged event notification", err)
		return err
	}
	return nil
}

func (ew *EventsWatcher) HandlePublicRelease(publicRelease *domain.CsmodulePublicRelease) error {
	message := fmt.Sprintf("- 🎉 Public release of CSM!: %s", publicRelease.Raw.TxHash)
	if err := ew.notifierPort.SendNotification(message); err != nil {
		logger.ErrorWithPrefix(ew.servicePrefix, "Error sending publicRelease event notification", err)
		return err
	}
	return nil
}
