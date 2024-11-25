package services

import (
	"context"
	"fmt"
	"lido-events/internal/application/domain"
	"lido-events/internal/application/ports"
	"lido-events/internal/logger"
	"sync"
	"time"
)

type DistributionLogUpdatedEventScanner struct {
	storagePort                     ports.StoragePort
	notifierPort                    ports.NotifierPort
	executionPort                   ports.ExecutionPort
	csFeeDistributorImplPort        ports.CsFeeDistributorImplPort
	csFeeDistributorBlockDeployment uint64
	servicePrefix                   string
}

func NewDistributionLogUpdatedEventScanner(storagePort ports.StoragePort, notifierPort ports.NotifierPort, executionPort ports.ExecutionPort, csFeeDistributorImplPort ports.CsFeeDistributorImplPort, csFeeDistributorBlockDeployment uint64) *DistributionLogUpdatedEventScanner {
	return &DistributionLogUpdatedEventScanner{
		storagePort,
		notifierPort,
		executionPort,
		csFeeDistributorImplPort,
		csFeeDistributorBlockDeployment,
		"DistributionLogUpdatedEventScanner",
	}
}

// ScanDistributionLogUpdatedEventsCron runs a periodic scan for DistributionLogUpdated events
func (ds *DistributionLogUpdatedEventScanner) ScanDistributionLogUpdatedEventsCron(ctx context.Context, interval time.Duration, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine finishes
	wg.Add(1)       // Increment the WaitGroup counter

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	// Run the scan logic immediately
	ds.runScan(ctx)

	for {
		select {
		case <-ticker.C:
			// Run the scan logic periodically
			ds.runScan(ctx)
		case <-ctx.Done():
			logger.InfoWithPrefix(ds.servicePrefix, "Stopping DistributionLogUpdated cron scan")
			return
		}
	}
}

// runScan contains the execution logic for scanning DistributionLogUpdated events
func (ds *DistributionLogUpdatedEventScanner) runScan(ctx context.Context) {
	// Retrieve start and end blocks for scanning
	start, err := ds.storagePort.GetDistributionLogLastProcessedBlock()
	if err != nil {
		logger.WarnWithPrefix(ds.servicePrefix, "Failed to get last processed block, using deployment block %d: %v", ds.csFeeDistributorBlockDeployment, err)
		start = ds.csFeeDistributorBlockDeployment
	}

	if start == 0 {
		logger.InfoWithPrefix(ds.servicePrefix, "No last processed block found, using deployment block %d", ds.csFeeDistributorBlockDeployment)
		start = ds.csFeeDistributorBlockDeployment
	}

	end, err := ds.executionPort.GetMostRecentBlockNumber()
	if err != nil {
		logger.ErrorWithPrefix(ds.servicePrefix, "Failed to get most recent block number, cannot continue with cron execution, waiting for next iteration: %v", err)
		return
	}

	// Perform the scan
	if err := ds.csFeeDistributorImplPort.ScanDistributionLogUpdatedEvents(ctx, start, &end, ds.HandleDistributionLogUpdatedEvent); err != nil {
		logger.ErrorWithPrefix(ds.servicePrefix, "Error scanning DistributionLogUpdated events: %v", err)
		return
	}

	// Save the last processed block if successful
	if err := ds.storagePort.SaveDistributionLogLastProcessedBlock(end); err != nil {
		logger.ErrorWithPrefix(ds.servicePrefix, "Error saving last processed block, next cron execution will run from previous processed: %v", err)
		return
	}
}

// HandleDistributionLogUpdatedEvent processes a DistributionLogUpdated event
func (ds *DistributionLogUpdatedEventScanner) HandleDistributionLogUpdatedEvent(distributionLogUpdated *domain.BindingsDistributionLogUpdated) error {
	// TODO add message saying go to this dashboard to see the validator performance
	logger.DebugWithPrefix(ds.servicePrefix, "Found DistributionLogUpdated event with log cid: %s", distributionLogUpdated.LogCid)

	if err := ds.storagePort.AddPendingHash(distributionLogUpdated.LogCid); err != nil {
		logger.ErrorWithPrefix(ds.servicePrefix, "Error adding pending hash %s: %v", distributionLogUpdated.LogCid, err)
		return err
	}

	message := fmt.Sprintf("- 📦 New distribution log updated: %s", distributionLogUpdated.LogCid)
	if err := ds.notifierPort.SendNotification(message); err != nil {
		logger.ErrorWithPrefix(ds.servicePrefix, "Error sending distributionLogUpdated notification: %v", err)
		return err
	}

	return nil
}
