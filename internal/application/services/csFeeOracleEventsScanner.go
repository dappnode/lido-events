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

// CsFeeOracleEventsScanner scans for ProcessingStarted events from the csFeeOracle contract.
type CsFeeOracleEventsScanner struct {
	storagePort                ports.StoragePort
	executionPort              ports.ExecutionPort
	csFeeOraclePort            ports.CsFeeOraclePort
	csFeeOracleBlockDeployment uint64
	mu                         sync.Mutex
	servicePrefix              string
}

// NewCsFeeOracleEventsScanner creates a new instance of the ProcessingStartedEventScanner.
func NewCsFeeOracleEventsScanner(
	storagePort ports.StoragePort,
	executionPort ports.ExecutionPort,
	csFeeOraclePort ports.CsFeeOraclePort,
	csFeeOracleBlockDeployment uint64,
) *CsFeeOracleEventsScanner {
	return &CsFeeOracleEventsScanner{
		storagePort:                storagePort,
		executionPort:              executionPort,
		csFeeOraclePort:            csFeeOraclePort,
		csFeeOracleBlockDeployment: csFeeOracleBlockDeployment,
		servicePrefix:              "ProcessingStartedEventScanner",
	}
}

// ScanProcessingStartedEventsCron runs an immediate scan and then periodically rescans for ProcessingStarted events.
// It signals the first execution via firstExecutionComplete and stops when the context is canceled.
func (ps *CsFeeOracleEventsScanner) ScanProcessingStartedEventsCron(
	ctx context.Context,
	interval time.Duration,
	wg *sync.WaitGroup,
	firstExecutionComplete chan struct{},
) {
	// Decrement WaitGroup when goroutine exits and increment immediately for this goroutine.
	defer wg.Done()
	wg.Add(1)

	// Execute immediately on startup.
	ps.RunScan(ctx)

	logger.DebugWithPrefix(ps.servicePrefix, "First execution complete, starting periodic cron for ProcessingStarted events")
	// Signal that the first execution is complete.
	close(firstExecutionComplete)

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			ps.RunScan(ctx)
		case <-ctx.Done():
			logger.InfoWithPrefix(ps.servicePrefix, "Stopping ProcessingStarted events cron scan")
			return
		}
	}
}

// RunScan contains the logic for scanning ProcessingStarted events.
func (ps *CsFeeOracleEventsScanner) RunScan(ctx context.Context) error {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	// Check if the node is still syncing.
	isSyncing, err := ps.executionPort.IsSyncing()
	if err != nil {
		logger.ErrorWithPrefix(ps.servicePrefix, "Error checking node sync status: %v", err)
		return err
	}
	if isSyncing {
		logger.InfoWithPrefix(ps.servicePrefix, "Node is syncing, skipping ProcessingStarted scan")
		return fmt.Errorf("node is syncing")
	}

	// Get the last processed block from storage. If not found, fall back to the deployment block.
	start, err := ps.storagePort.GetProcessingStartedLastProcessedBlock()
	if err != nil {
		logger.WarnWithPrefix(ps.servicePrefix, "Failed to get last processed block for ProcessingStarted events, using deployment block %d: %v", ps.csFeeOracleBlockDeployment, err)
		start = ps.csFeeOracleBlockDeployment
	}
	if start == 0 {
		logger.InfoWithPrefix(ps.servicePrefix, "No last processed block found for ProcessingStarted events, using deployment block %d", ps.csFeeOracleBlockDeployment)
		start = ps.csFeeOracleBlockDeployment
	}

	// Get the most recent block number from the execution port.
	end, err := ps.executionPort.GetMostRecentBlockNumber()
	if err != nil {
		logger.ErrorWithPrefix(ps.servicePrefix, "Failed to get most recent block number: %v", err)
		return err
	}

	// Scan for ProcessingStarted events in the block range [start, end].
	err = ps.csFeeOraclePort.ScanProcessingStartedEvents(ctx, start, &end, ps.HandleProcessingStartedEvent)
	if err != nil {
		logger.ErrorWithPrefix(ps.servicePrefix, "Error scanning ProcessingStarted events: %v", err)
		return err
	}

	// Save the most recent block number as the last processed block.
	if err := ps.storagePort.SaveProcessingStartedLastProcessedBlock(end); err != nil {
		logger.ErrorWithPrefix(ps.servicePrefix, "Error saving last processed block for ProcessingStarted events: %v", err)
		return err
	}

	return nil
}

// HandleProcessingStartedEvent processes a single ProcessingStarted event.
// It logs the event, stores the pending hash, and sends a notification if the event occurred within the last 24 hours.
func (ps *CsFeeOracleEventsScanner) HandleProcessingStartedEvent(event *domain.BindingsProcessingStarted, blockNumber *big.Int) error {
	logger.DebugWithPrefix(ps.servicePrefix, "Found ProcessingStarted event")

	// Add the event hash to pending processing in storage.
	if err := ps.storagePort.SaveProcessingStartedEvent(*event); err != nil {
		logger.ErrorWithPrefix(ps.servicePrefix, "Error adding pending ProcessingStarted hash: %v", err)
		return err
	}

	return nil
}
