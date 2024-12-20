package services

import (
	"context"
	"fmt"
	"lido-events/internal/application/domain"
	"lido-events/internal/application/ports"
	"lido-events/internal/logger"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type DistributionLogUpdatedEventScanner struct {
	storagePort                     ports.StoragePort
	notifierPort                    ports.NotifierPort
	executionPort                   ports.ExecutionPort
	csFeeDistributorImplPort        ports.CsFeeDistributorImplPort
	csFeeDistributorBlockDeployment uint64
	csModuleTxReceipt               common.Hash
	servicePrefix                   string
}

func NewDistributionLogUpdatedEventScanner(storagePort ports.StoragePort, notifierPort ports.NotifierPort, executionPort ports.ExecutionPort, csFeeDistributorImplPort ports.CsFeeDistributorImplPort, csFeeDistributorBlockDeployment uint64, csModuleTxReceipt common.Hash) *DistributionLogUpdatedEventScanner {
	return &DistributionLogUpdatedEventScanner{
		storagePort,
		notifierPort,
		executionPort,
		csFeeDistributorImplPort,
		csFeeDistributorBlockDeployment,
		csModuleTxReceipt,
		"DistributionLogUpdatedEventScanner",
	}
}

func (ds *DistributionLogUpdatedEventScanner) ScanDistributionLogUpdatedEventsCron(ctx context.Context, interval time.Duration, wg *sync.WaitGroup, firstExecutionComplete chan struct{}) {
	defer wg.Done() // Decrement the counter when the goroutine finishes
	wg.Add(1)       // Increment the WaitGroup counter

	// Execute immediately on startup
	ds.runScan(ctx)

	logger.DebugWithPrefix(ds.servicePrefix, "First execution complete, sending signal to start periodic cron for loading pending CIDs")
	// Signal that the first execution is complete
	close(firstExecutionComplete)

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

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
	// Skip if node is syncing
	isSyncing, err := ds.executionPort.IsSyncing()
	if err != nil {
		logger.ErrorWithPrefix(ds.servicePrefix, "Error checking if node is syncing: %v", err)
		return
	}

	if isSyncing {
		logger.InfoWithPrefix(ds.servicePrefix, "Node is syncing, skipping DistributionLogUpdated scan")
		return
	}

	// Skip if tx receipt not found (nil). This means that the node does not store log receipts and there are no logs at all
	receipt, err := ds.executionPort.GetTransactionReceipt(ds.csModuleTxReceipt)
	if err != nil {
		logger.ErrorWithPrefix(ds.servicePrefix, "Error getting transaction receipt for csModule deployment: %v", err)
		return
	}
	if receipt == nil {
		logger.WarnWithPrefix(ds.servicePrefix, "Transaction receipt for csModule deployment not found, skipping ValidatorExitRequest event scan. This means that the node does not store log receipts and there are no logs at all")
		return
	}

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
	logger.DebugWithPrefix(ds.servicePrefix, "Found DistributionLogUpdated event with log cid: %s", distributionLogUpdated.LogCid)

	if err := ds.storagePort.AddPendingHash(distributionLogUpdated.LogCid); err != nil {
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
