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
	// Error channel to collect errors
	errorChan := make(chan error, 3)

	// Start each watcher in its own goroutine
	startWatcher := func(watchFunc func(context.Context) error, name string) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			logger.InfoWithPrefix(ew.servicePrefix, "Starting watcher for "+name+" events...")
			if err := watchFunc(ctx); err != nil {
				errorChan <- fmt.Errorf("%s watcher error: %w", name, err)
			}
		}()
	}

	startWatcher(ew.watchCsModuleEvents, "CSModule")
	startWatcher(ew.watchCsFeeDistributorEvents, "CsFeeDistributor")
	startWatcher(ew.watchReportSubmittedEvents, "ReportSubmitted")

	// Monitor for errors and context cancellation
	go func() {
		defer close(errorChan)
		for {
			select {
			case err := <-errorChan:
				if err != nil {
					logger.ErrorWithPrefix(ew.servicePrefix, "Error in event watcher: %v", err)
				}
			case <-ctx.Done():
				logger.InfoWithPrefix(ew.servicePrefix, "Context canceled, stopping all event watchers.")
				return
			}
		}
	}()
}

func (ew *EventsWatcher) watchReportSubmittedEvents(ctx context.Context) error {
	return ew.veboPort.WatchReportSubmittedEvents(ctx, ew.HandleReportSubmittedEvent)
}

func (ew *EventsWatcher) watchCsModuleEvents(ctx context.Context) error {
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
	return ew.csFeeDistributorPort.WatchCsFeeDistributorEvents(ctx, ew.HandleDistributionDataUpdated)
}

// Handlers Vebo

func (ew *EventsWatcher) HandleReportSubmittedEvent(reportSubmitted *domain.VeboReportSubmitted) error {
	// TODO: research what does this event mean and what should be done with it and why its emitted so often
	// send the notification message
	// message := fmt.Sprintf("- 📈 New submitted report in slot %s", reportSubmitted.RefSlot)
	// if err := ew.notifierPort.SendNotification(message); err != nil {
	// 	logger.ErrorWithPrefix(ew.servicePrefix, "Error sending reportSubmitted event notification", err)
	// 	return err
	// }

	return nil
}

// Handlers csFeeDistributor

func (ew *EventsWatcher) HandleDistributionDataUpdated(rewardsDistributed *domain.CsfeedistributorDistributionDataUpdated) error {
	// send the notification message
	message := fmt.Sprintf("- 📈 New rewards distributed. Total claimable shares: %s", rewardsDistributed.TotalClaimableShares)
	if err := ew.notifierPort.SendNotification(message); err != nil {
		logger.ErrorWithPrefix(ew.servicePrefix, "Error sending rewardsDistributed event notification", err)
		return err
	}

	return nil
}

// Handlers CSModule

func (ew *EventsWatcher) HandleDepositedSigningKeysCountChanged(depositedSigningKeysCountChanged *domain.CsmoduleDepositedSigningKeysCountChanged) error {
	message := fmt.Sprintf("- 🤩 Node Operator's keys received deposit keys count: %s", depositedSigningKeysCountChanged.DepositedKeysCount)
	if err := ew.notifierPort.SendNotification(message); err != nil {
		logger.ErrorWithPrefix(ew.servicePrefix, "Error sending depositedSigningKeysCountChanged event notification", err)
		return err
	}
	return nil
}

func (ew *EventsWatcher) HandleElRewardsStealingPenaltyReported(eLRewardsStealingPenaltyReported *domain.CsmoduleELRewardsStealingPenaltyReported) error {
	message := fmt.Sprintf("- 🚨 Penalty for stealing EL rewards reported. Stolen amont: %s", eLRewardsStealingPenaltyReported.StolenAmount)
	if err := ew.notifierPort.SendNotification(message); err != nil {
		logger.ErrorWithPrefix(ew.servicePrefix, "Error sending eLRewardsStealingPenaltyReported event notification", err)
		return err
	}
	return nil
}

func (ew *EventsWatcher) HandleElRewardsStealingPenaltySettled(eLRewardsStealingPenaltySettled *domain.CsmoduleELRewardsStealingPenaltySettled) error {
	message := fmt.Sprintf("- 🚨 EL rewards stealing penalty confirmed and applied for Node Operator ID: %s", eLRewardsStealingPenaltySettled.NodeOperatorId)
	if err := ew.notifierPort.SendNotification(message); err != nil {
		logger.ErrorWithPrefix(ew.servicePrefix, "Error sending eLRewardsStealingPenaltySettled event notification", err)
		return err
	}
	return nil
}

func (ew *EventsWatcher) HandleElRewardsStealingPenaltyCancelled(eLRewardsStealingPenaltyCancelled *domain.CsmoduleELRewardsStealingPenaltyCancelled) error {
	message := fmt.Sprintf("- 😮‍💨 Cancelled penalty for stealing EL reward. Amount: %s", eLRewardsStealingPenaltyCancelled.Amount)
	if err := ew.notifierPort.SendNotification(message); err != nil {
		logger.ErrorWithPrefix(ew.servicePrefix, "Error sending eLRewardsStealingPenaltyCancelled event notification", err)
		return err
	}
	return nil
}

func (ew *EventsWatcher) HandleInitialSlashingSubmitted(initialSlashingSubmitted *domain.CsmoduleInitialSlashingSubmitted) error {
	message := fmt.Sprintf("- 🚨 Initial slashing submitted for one of the validators. Key index: %s", initialSlashingSubmitted.KeyIndex)
	if err := ew.notifierPort.SendNotification(message); err != nil {
		logger.ErrorWithPrefix(ew.servicePrefix, "Error sending initialSlashingSubmitted event notification", err)
		return err
	}
	return nil
}

func (ew *EventsWatcher) HandleKeyRemovalChargeApplied(keyRemovalChargeApplied *domain.CsmoduleKeyRemovalChargeApplied) error {
	message := fmt.Sprintf("- 🔑 Applied charge for key removal for Node Operator ID: %s", keyRemovalChargeApplied.NodeOperatorId)
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
	message := fmt.Sprintf("- 🚨 Reported stuck keys that were not exited in time. Stuck keys count: %s", stuckSigningKeysCountChanged.StuckKeysCount)
	if err := ew.notifierPort.SendNotification(message); err != nil {
		logger.ErrorWithPrefix(ew.servicePrefix, "Error sending stuckSigningKeysCountChanged event notification", err)
		return err
	}
	return nil
}

func (ew *EventsWatcher) HandleVettedSigningKeysCountDecreased(vettedSigningKeysCountDecreased *domain.CsmoduleVettedSigningKeysCountDecreased) error {
	message := fmt.Sprintf("- 🚨 Uploaded invalid keys for Node Operator ID: %s", vettedSigningKeysCountDecreased.NodeOperatorId)
	if err := ew.notifierPort.SendNotification(message); err != nil {
		logger.ErrorWithPrefix(ew.servicePrefix, "Error sending vettedSigningKeysCountDecreased event notification", err)
		return err
	}
	return nil
}

func (ew *EventsWatcher) HandleWithdrawalSubmitted(withdrawalSubmitted *domain.CsmoduleWithdrawalSubmitted) error {
	message := fmt.Sprintf("- 👀 Key withdrawal information submitted. Amount: %s", withdrawalSubmitted.Amount)
	if err := ew.notifierPort.SendNotification(message); err != nil {
		logger.ErrorWithPrefix(ew.servicePrefix, "Error sending withdrawalSubmitted event notification", err)
		return err
	}
	return nil
}

func (ew *EventsWatcher) HandleTotalSigningKeysCountChanged(totalSigningKeysCountChanged *domain.CsmoduleTotalSigningKeysCountChanged) error {
	message := fmt.Sprintf("- 👀 New keys uploaded or removed. Keys count: %s", totalSigningKeysCountChanged.TotalKeysCount)
	if err := ew.notifierPort.SendNotification(message); err != nil {
		logger.ErrorWithPrefix(ew.servicePrefix, "Error sending totalSigningKeysCountChanged event notification", err)
		return err
	}
	return nil
}

func (ew *EventsWatcher) HandlePublicRelease(publicRelease *domain.CsmodulePublicRelease) error {
	message := fmt.Sprintf("- 🎉 Public release of CSM!. Tx Hash: %s", publicRelease.Raw.TxHash)
	if err := ew.notifierPort.SendNotification(message); err != nil {
		logger.ErrorWithPrefix(ew.servicePrefix, "Error sending publicRelease event notification", err)
		return err
	}
	return nil
}
