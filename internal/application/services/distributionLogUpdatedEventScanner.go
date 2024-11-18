package services

import (
	"context"
	"fmt"
	"lido-events/internal/application/domain"
	"lido-events/internal/application/ports"
	"log"
	"time"
)

type DistributionLogUpdatedEventScanner struct {
	storagePort                     ports.StoragePort
	notifierPort                    ports.NotifierPort
	executionPort                   ports.ExecutionPort
	csFeeDistributorImplPort        ports.CsFeeDistributorImplPort
	csFeeDistributorBlockDeployment uint64
}

func NewDistributionLogUpdatedEventScanner(storagePort ports.StoragePort, notifierPort ports.NotifierPort, executionPort ports.ExecutionPort, csFeeDistributorImplPort ports.CsFeeDistributorImplPort, csFeeDistributorBlockDeployment uint64) *DistributionLogUpdatedEventScanner {
	return &DistributionLogUpdatedEventScanner{
		storagePort,
		notifierPort,
		executionPort,
		csFeeDistributorImplPort,
		csFeeDistributorBlockDeployment,
	}
}

// ScanDistributionLogUpdatedEventsCron runs a periodic scan for DistributionLogUpdated events
func (ds *DistributionLogUpdatedEventScanner) ScanDistributionLogUpdatedEventsCron(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Retrieve start and end blocks for scanning
			start, err := ds.storagePort.GetDistributionLogLastProcessedBlock()
			if err != nil {
				log.Printf("Failed to get last processed block, using deployment block: %v", err)
				start = ds.csFeeDistributorBlockDeployment
			}

			if start == 0 {
				log.Printf("No last processed block found, using deployment block %d", ds.csFeeDistributorBlockDeployment)
				start = ds.csFeeDistributorBlockDeployment
			}

			end, err := ds.executionPort.GetMostRecentBlockNumber()
			if err != nil {
				continue
			}

			// Perform the scan
			if err := ds.csFeeDistributorImplPort.ScanDistributionLogUpdatedEvents(ctx, start, &end, ds.HandleDistributionLogUpdatedEvent); err != nil {
				continue
			}

			// Save the last processed epoch if successful
			if err := ds.storagePort.SaveDistributionLogLastProcessedBlock(end); err != nil {
				continue
			}
		case <-ctx.Done():
			log.Println("Stopping DistributionLogUpdated cron scan")
			return
		}
	}
}

// HandleDistributionLogUpdatedEvent processes a DistributionLogUpdated event
func (ds *DistributionLogUpdatedEventScanner) HandleDistributionLogUpdatedEvent(distributionLogUpdated *domain.BindingsDistributionLogUpdated) error {
	// TODO add message saying go to this dashboard to see the validator performance
	log.Printf("Found DistributionLogUpdated event with log cid: %s", distributionLogUpdated.LogCid)

	if err := ds.storagePort.AddPendingHash(distributionLogUpdated.LogCid); err != nil {
		return err
	}

	message := fmt.Sprintf("- 📦 New distribution log updated: %s", distributionLogUpdated.LogCid)
	if err := ds.notifierPort.SendNotification(message); err != nil {
		return err
	}

	return nil
}
