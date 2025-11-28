package services

import (
	"context"
	"fmt"
	"lido-events/internal/application/domain"
	"lido-events/internal/application/ports"
	"lido-events/internal/logger"
	"math/big"
	"sync"
	"time"
)

type DistributionLogUpdatedEventScanner struct {
	storagePort          ports.StoragePort
	notifierPort         ports.NotifierPort
	executionPort        ports.ExecutionPort
	csFeeDistributorPort ports.CsFeeDistributorPort
	mu                   sync.Mutex
	servicePrefix        string
}

func NewDistributionLogUpdatedEventScanner(storagePort ports.StoragePort, notifierPort ports.NotifierPort, executionPort ports.ExecutionPort, csFeeDistributorPort ports.CsFeeDistributorPort) *DistributionLogUpdatedEventScanner {
	return &DistributionLogUpdatedEventScanner{
		storagePort,
		notifierPort,
		executionPort,
		csFeeDistributorPort,
		sync.Mutex{},
		"DistributionLogUpdatedEventScanner",
	}
}

func (ds *DistributionLogUpdatedEventScanner) ScanDistributionLogUpdatedEventsCron(ctx context.Context, interval time.Duration, wg *sync.WaitGroup, firstExecutionComplete chan struct{}) {
	defer wg.Done() // Decrement the counter when the goroutine finishes
	wg.Add(1)       // Increment the WaitGroup counter

	// Execute immediately on startup
	ds.RunScan(ctx)

	logger.DebugWithPrefix(ds.servicePrefix, "First execution complete, sending signal to start periodic cron for loading pending CIDs")
	// Signal that the first execution is complete
	close(firstExecutionComplete)

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Run the scan logic periodically
			ds.RunScan(ctx)
		case <-ctx.Done():
			logger.InfoWithPrefix(ds.servicePrefix, "Stopping DistributionLogUpdated cron scan")
			return
		}
	}
}

// RunScan contains the execution logic for scanning DistributionLogUpdated events
func (ds *DistributionLogUpdatedEventScanner) RunScan(ctx context.Context) error {
	ds.mu.Lock()         // Lock mutex to ensure only one execution at a time
	defer ds.mu.Unlock() // Unlock when function exits

	// Skip if node is syncing
	isSyncing, err := ds.executionPort.IsSyncing()
	if err != nil {
		logger.ErrorWithPrefix(ds.servicePrefix, "Error checking if node is syncing: %v", err)
		return err
	}

	if isSyncing {
		logger.InfoWithPrefix(ds.servicePrefix, "Node is syncing, skipping DistributionLogUpdated scan")
		return fmt.Errorf("node is syncing")
	}

	// Do a for loop over the operator ids
	operatorIds, err := ds.storagePort.GetOperatorIds()
	if err != nil {
		logger.ErrorWithPrefix(ds.servicePrefix, "Error getting operator ids: %v", err)
		return err
	}

	for _, operatorId := range operatorIds {
		// Retrieve start and end blocks for scanning
		start, err := ds.storagePort.GetDistributionLogLastProcessedBlock(operatorId)
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
			return err
		}

		// Perform the scan
		if err := ds.csFeeDistributorImplPort.ScanDistributionLogUpdatedEvents(ctx, start, &end, operatorId, ds.HandleDistributionLogUpdatedEvent); err != nil {
			logger.ErrorWithPrefix(ds.servicePrefix, "Error scanning DistributionLogUpdated events: %v", err)
			return err
		}

		// Save the last processed block if successful
		if err := ds.storagePort.SaveDistributionLogLastProcessedBlock(operatorId, end); err != nil {
			logger.ErrorWithPrefix(ds.servicePrefix, "Error saving last processed block, next cron execution will run from previous processed: %v", err)
			return err
		}
	}

	return nil
}

// HandleDistributionLogUpdatedEvent processes a DistributionLogUpdated event
func (ds *DistributionLogUpdatedEventScanner) HandleDistributionLogUpdatedEvent(distributionLogUpdated *domain.BindingsDistributionLogUpdated, operatorId *big.Int) error {
	logger.DebugWithPrefix(ds.servicePrefix, "Found DistributionLogUpdated event with log cid: %s", distributionLogUpdated.LogCid)

	if err := ds.storagePort.AddPendingHash(operatorId, distributionLogUpdated.LogCid); err != nil {
		logger.ErrorWithPrefix(ds.servicePrefix, "Error adding pending hash %s: %v", distributionLogUpdated.LogCid, err)
		return err
	}

	blockTimestamp, err := ds.executionPort.GetBlockTimestampByNumber(distributionLogUpdated.Raw.BlockNumber)
	if err != nil {
		logger.ErrorWithPrefix(ds.servicePrefix, "Error getting block timestamp for block number %d: %v", distributionLogUpdated.Raw.BlockNumber, err)
		return err
	}

	// If the block timestamp is within the last 24 hours, send a notification
	if time.Now().Unix()-int64(blockTimestamp) < 24*60*60 {
		message := fmt.Sprintf("- 📦 New distribution log updated: %s. This report will be fetched and parsed from an IPFS gateway for later validator performance visualization.", distributionLogUpdated.LogCid)
		if err := ds.notifierPort.SendNotification(message); err != nil {
			logger.ErrorWithPrefix(ds.servicePrefix, "Error sending distributionLogUpdated notification: %v", err)
			return err
		}

	}

	return nil
}
