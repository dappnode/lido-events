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

type PendingHashesLoader struct {
	storagePort   ports.StoragePort
	notifierPort  ports.NotifierPort
	ipfsPort      ports.IpfsPort
	servicePrefix string
}

func NewPendingHashesLoader(storagePort ports.StoragePort, notifierPort ports.NotifierPort, ipfsPort ports.IpfsPort) *PendingHashesLoader {
	return &PendingHashesLoader{
		storagePort,
		notifierPort,
		ipfsPort,
		"PendingHashesLoader",
	}
}

func (phl *PendingHashesLoader) LoadPendingHashesCron(ctx context.Context, interval time.Duration, wg *sync.WaitGroup, firstExecutionComplete chan struct{}) {
	defer wg.Done() // Decrement the counter when the goroutine finishes
	wg.Add(1)       // Increment the WaitGroup counter

	// Wait for the signal from cron event scanner
	<-firstExecutionComplete
	logger.DebugWithPrefix(phl.servicePrefix, "Signal received, starting periodic cron for loading pending CIDs")

	// Execute immediately on startup
	if err := phl.LoadPendingHashes(); err != nil {
		logger.InfoWithPrefix(phl.servicePrefix, "Error loading pending hashes: %v", err)
	}

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Call the load method periodically
			if err := phl.LoadPendingHashes(); err != nil {
				logger.InfoWithPrefix(phl.servicePrefix, "Error loading pending hashes: %v", err)
				continue
			}
		case <-ctx.Done():
			// Stop the periodic load if the context is canceled
			logger.Info("Stopping periodic cron for loading pending CIDs")
			return
		}
	}
}

func (phl *PendingHashesLoader) LoadPendingHashes() error {
	// Get operator IDs
	operatorIDs, err := phl.storagePort.GetOperatorIds()
	if err != nil {
		logger.ErrorWithPrefix(phl.servicePrefix, "Failed to get operator IDs", err)
		return err
	}
	if len(operatorIDs) == 0 {
		logger.InfoWithPrefix(phl.servicePrefix, "No operator IDs found; skipping loading pending hashes")
		return nil
	}

	// Get all pending hashes
	pendingHashes, err := phl.storagePort.GetPendingHashes()
	if err != nil {
		logger.ErrorWithPrefix(phl.servicePrefix, "Failed to get pending hashes", err)
		return err
	}
	if len(pendingHashes) == 0 {
		logger.InfoWithPrefix(phl.servicePrefix, "No pending hashes found; skipping loading pending hashes")
		return nil
	}

	// Process each pending hash
	for _, pendingHash := range pendingHashes {
		logger.DebugWithPrefix(phl.servicePrefix, "Fetching and parsing IPFS data for pending hash %s", pendingHash)

		originalReport, err := phl.ipfsPort.FetchAndParseIpfs(pendingHash)
		if err != nil {
			logger.ErrorWithPrefix(phl.servicePrefix, "Failed to fetch and parse IPFS data for pending hash %s, skipping hash: %v", pendingHash, err)
			continue
		}

		for _, operatorID := range operatorIDs {
			logger.DebugWithPrefix(phl.servicePrefix, "Processing data for operator ID %s", operatorID.String())

			data, exists := originalReport.Operators[operatorID.String()]
			if !exists {
				logger.WarnWithPrefix(phl.servicePrefix, "Operator ID %s not found in the original report, skipping", operatorID.String())
				continue
			}

			// Check performance and send notifications if needed
			// TODO: only send the notification if the report is from the last 24 hours, consider using the frame to calculate the timestamp
			phl.CheckAndNotifyPerformance(operatorID, data.Validators, originalReport)

			// Save the report
			report := domain.Report{
				Frame:     originalReport.Frame,
				Data:      data,
				Threshold: originalReport.Threshold,
			}
			if err := phl.storagePort.SaveReport(operatorID, report); err != nil {
				logger.ErrorWithPrefix(phl.servicePrefix, "Failed to save report for operator ID %s: %v", operatorID.String(), err)
				continue
			}
		}

		// Remove the pending hash
		if err := phl.storagePort.DeletePendingHash(pendingHash); err != nil {
			logger.ErrorWithPrefix(phl.servicePrefix, "Failed to delete pending hash %s: %v", pendingHash, err)
		}
	}

	return nil
}

func (phl *PendingHashesLoader) CheckAndNotifyPerformance(operatorID *big.Int, validators map[string]domain.Validator, originalReport domain.OriginalReport) {
	for validatorID, validator := range validators {
		if validator.Perf.Assigned > 0 {
			performance := float64(validator.Perf.Included) / float64(validator.Perf.Assigned) * 100
			if performance < originalReport.Threshold {
				message := fmt.Sprintf(
					"- 🚨 Operator ID: %s, Validator: %s has performance %.2f%% below the threshold %.2f%% in frame %v",
					operatorID.String(), validatorID, performance, originalReport.Threshold, originalReport.Frame,
				)
				logger.WarnWithPrefix(phl.servicePrefix, message)

				// Send notification
				if err := phl.notifierPort.SendNotification(message); err != nil {
					logger.ErrorWithPrefix(phl.servicePrefix, "Failed to send notification: %v", err)
				}
			}
		}
	}
}
