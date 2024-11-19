package services

import (
	"context"
	"lido-events/internal/application/domain"
	"lido-events/internal/application/ports"
	"lido-events/internal/logger"
	"sync"
	"time"
)

type PendingHashesLoader struct {
	storagePort   ports.StoragePort
	ipfsPort      ports.IpfsPort
	servicePrefix string
}

func NewPendingHashesLoader(storagePort ports.StoragePort, ipfsPort ports.IpfsPort) *PendingHashesLoader {
	return &PendingHashesLoader{
		storagePort,
		ipfsPort,
		"PendingHashesLoader",
	}
}

// loadPendingHashesCron starts a periodic service to load pending hashes from IPFS
func (phl *PendingHashesLoader) LoadPendingHashesCron(ctx context.Context, interval time.Duration, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine finishes
	wg.Add(1)       // Increment the WaitGroup counter

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Call the load method periodically
			if err := phl.loadPendingHashes(); err != nil {
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

// loadPendingHashes loads all pending hashes from the storage and fetches the corresponding IPFS data
func (phl *PendingHashesLoader) loadPendingHashes() error {
	// Get operator IDs
	operatorIDs, err := phl.storagePort.GetOperatorIds()
	if err != nil {
		logger.ErrorWithPrefix(phl.servicePrefix, "Failed to get operator IDs", err)
		return err
	}
	// Skip if there are no operator IDs
	if len(operatorIDs) == 0 {
		logger.InfoWithPrefix(phl.servicePrefix, "No operator IDs found; skipping loading pending hashes")
		return nil
	}

	// Get all pending hashes from the storage
	pendingHashes, err := phl.storagePort.GetPendingHashes()
	if err != nil {
		logger.ErrorWithPrefix(phl.servicePrefix, "Failed to get pending hashes", err)
		return err
	}
	// Skip if there are no pending hashes
	if len(pendingHashes) == 0 {
		logger.InfoWithPrefix(phl.servicePrefix, "No pending hashes found; skipping loading pending hashes")
		return nil
	}

	// Fetch and parse IPFS data for each pending hash
	for _, pendingHash := range pendingHashes {
		logger.DebugWithPrefix(phl.servicePrefix, "Fetching and parsing IPFS data for pending hash %s", pendingHash)

		originalReport, err := phl.ipfsPort.FetchAndParseIpfs(pendingHash)
		if err != nil {
			logger.ErrorWithPrefix(phl.servicePrefix, "Failed to fetch and parse IPFS data for pending hash %s, skipping hash: %v", pendingHash, err)
			continue
		}

		// Process each operator ID in the original report
		for _, operatorID := range operatorIDs {
			logger.DebugWithPrefix(phl.servicePrefix, "Saving report data for operator ID %s", operatorID.String())

			// Get the data for the operator ID
			data, exists := originalReport.Operators[operatorID.String()]
			if !exists {
				// Skip if the operator ID is not found in the report
				logger.WarnWithPrefix(phl.servicePrefix, "Operator ID %s not found in the original report, skipping", operatorID.String())
				continue
			}

			report := domain.Report{
				Frame:     originalReport.Frame,
				Data:      data,
				Threshold: originalReport.Threshold,
			}

			// Save the report data for the operator ID
			if err := phl.storagePort.SaveReport(operatorID, report); err != nil {
				logger.ErrorWithPrefix(phl.servicePrefix, "Failed to save report data for operator ID %s: %v", operatorID.String(), err)
				continue
			}
		}

		// Remove the pending hash from storage
		if err := phl.storagePort.DeletePendingHash(pendingHash); err != nil {
			logger.ErrorWithPrefix(phl.servicePrefix, "Failed to delete pending hash %s: %v", pendingHash, err)
			continue
		}
	}

	return nil
}
