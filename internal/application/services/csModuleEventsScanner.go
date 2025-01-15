package services

import (
	"context"
	"lido-events/internal/application/domain"
	"lido-events/internal/application/ports"
	"lido-events/internal/logger"
	"sync"
	"time"
)

type CsModuleEventsScanner struct {
	storagePort   ports.StoragePort
	executionPort ports.ExecutionPort
	csModulePort  ports.CsModulePort
	servicePrefix string
}

func NewCsModuleEventsScanner(storagePort ports.StoragePort, executionPort ports.ExecutionPort, csModulePort ports.CsModulePort) *CsModuleEventsScanner {
	return &CsModuleEventsScanner{
		storagePort:   storagePort,
		executionPort: executionPort,
		csModulePort:  csModulePort,
		servicePrefix: "CsModuleEventsScanner",
	}
}

func (cs *CsModuleEventsScanner) ScanCsModuleEventsCron(ctx context.Context, interval time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(1)

	cs.runScan(ctx)

	logger.DebugWithPrefix(cs.servicePrefix, "First execution complete of CsModule events scanner")

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			cs.runScan(ctx)
		case <-ctx.Done():
			logger.InfoWithPrefix(cs.servicePrefix, "Stopping CsModule events cron scan")
			return
		}
	}
}

func (cs *CsModuleEventsScanner) runScan(ctx context.Context) {
	isSyncing, err := cs.executionPort.IsSyncing()
	if err != nil {
		logger.ErrorWithPrefix(cs.servicePrefix, "Error checking if node is syncing: %v", err)
		return
	}

	if isSyncing {
		logger.InfoWithPrefix(cs.servicePrefix, "Node is syncing, skipping CsModule events scan")
		return
	}

	start, err := cs.storagePort.GetCsModuleLastProcessedBlock()
	if err != nil {
		logger.WarnWithPrefix(cs.servicePrefix, "Failed to get last processed block, using deployment block: %v", err)
		start = 0 // Default to block 0 if no value is set
	}

	end, err := cs.executionPort.GetMostRecentBlockNumber()
	if err != nil {
		logger.ErrorWithPrefix(cs.servicePrefix, "Failed to get most recent block number: %v", err)
		return
	}

	if err := cs.csModulePort.ScanNodeOperatorEvents(
		ctx,
		start,
		&end,
		cs.HandleNodeOperatorAddedEvent,
		cs.HandleNodeOperatorManagerAddressChangedEvent,
		cs.HandleNodeOperatorRewardAddressChangedEvent,
	); err != nil {
		logger.ErrorWithPrefix(cs.servicePrefix, "Error scanning NodeOperator events: %v", err)
		return
	}

	if err := cs.storagePort.SaveCsModuletLastProcessedBlock(end); err != nil {
		logger.ErrorWithPrefix(cs.servicePrefix, "Error saving last processed block: %v", err)
		return
	}
}

func (cs *CsModuleEventsScanner) HandleNodeOperatorAddedEvent(event *domain.CsmoduleNodeOperatorAdded) error {
	logger.DebugWithPrefix(cs.servicePrefix, "Handling NodeOperatorAdded event: %+v", event)

	if err := cs.storagePort.SetNodeOperatorAdded(event.NodeOperatorId.String(), *event); err != nil {
		logger.ErrorWithPrefix(cs.servicePrefix, "Error saving NodeOperatorAdded event: %v", err)
		return err
	}
	return nil
}

func (cs *CsModuleEventsScanner) HandleNodeOperatorManagerAddressChangedEvent(event *domain.CsmoduleNodeOperatorManagerAddressChanged) error {
	logger.DebugWithPrefix(cs.servicePrefix, "Handling NodeOperatorManagerAddressChanged event: %+v", event)

	if err := cs.storagePort.SetNodeOperatorManagerAddressChanged(event.NodeOperatorId.String(), *event); err != nil {
		logger.ErrorWithPrefix(cs.servicePrefix, "Error saving NodeOperatorManagerAddressChanged event: %v", err)
		return err
	}
	return nil
}

func (cs *CsModuleEventsScanner) HandleNodeOperatorRewardAddressChangedEvent(event *domain.CsmoduleNodeOperatorRewardAddressChanged) error {
	logger.DebugWithPrefix(cs.servicePrefix, "Handling NodeOperatorRewardAddressChanged event: %+v", event)

	if err := cs.storagePort.SetNodeOperatorRewardAddressChanged(event.NodeOperatorId.String(), *event); err != nil {
		logger.ErrorWithPrefix(cs.servicePrefix, "Error saving NodeOperatorRewardAddressChanged event: %v", err)
		return err
	}
	return nil
}
