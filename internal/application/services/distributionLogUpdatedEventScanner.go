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
				log.Printf("Failed to get last processed block: %v", err)
				continue
			}

			if start == 0 {
				start = ds.csFeeDistributorBlockDeployment
			}

			end, err := ds.executionPort.GetMostRecentBlockNumber()
			if err != nil {
				log.Printf("Failed to get latest finalized block: %v", err)
				continue
			}

			// Perform the scan
			if err := ds.csFeeDistributorImplPort.ScanDistributionLogUpdatedEvents(ctx, start, &end, ds.HandleDistributionLogUpdatedEvent); err != nil {
				log.Printf("Error scanning DistributionLogUpdated events: %v", err)
				continue
			}

			// Save the last processed epoch if successful
			if err := ds.storagePort.SaveDistributionLogLastProcessedBlock(end); err != nil {
				log.Printf("Failed to save last processed epoch: %v", err)
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
		log.Printf("Failed to store pending CID: %v", err)
		return err
	}

	message := fmt.Sprintf("- 📦 New distribution log updated: %s", distributionLogUpdated.LogCid)
	if err := ds.notifierPort.SendNotification(message); err != nil {
		log.Printf("Failed to send notification: %v", err)
		return err
	}

	return nil
}
