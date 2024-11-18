package services

import (
	"context"
	"lido-events/internal/application/domain"
	"lido-events/internal/application/ports"
	"log"
	"os"
	"time"
)

type PendingHashesLoader struct {
	storagePort ports.StoragePort
	ipfsPort    ports.IpfsPort
	logger      *log.Logger
}

func NewPendingHashesLoader(storagePort ports.StoragePort, ipfsPort ports.IpfsPort) *PendingHashesLoader {
	logger := log.New(os.Stdout, "[PendingHashesLoader] ", log.LstdFlags)

	return &PendingHashesLoader{
		storagePort,
		ipfsPort,
		logger,
	}
}

// LoadPendingHashesCron starts a periodic service to load pending hashes from IPFS
func (phl *PendingHashesLoader) LoadPendingHashesCron(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Call the load method periodically
			if err := phl.LoadPendingHashes(); err != nil {
				phl.logger.Printf("Error loading pending hashes: %v", err)
			}
		case <-ctx.Done():
			// Stop the periodic load if the context is canceled
			phl.logger.Println("Stopping periodic cron for loading pending CIDs")
			return
		}
	}
}

// LoadPendingHashes loads all pending hashes from the storage and fetches the corresponding IPFS data
func (phl *PendingHashesLoader) LoadPendingHashes() error {
	// Get operator IDs
	operatorIDs, err := phl.storagePort.GetOperatorIds()
	if err != nil {
		return err
	}
	// Skip if there are no operator IDs
	if len(operatorIDs) == 0 {
		phl.logger.Println("No operator IDs found; skipping loading pending hashes")
		return nil
	}

	// Get all pending hashes from the storage
	pendingHashes, err := phl.storagePort.GetPendingHashes()
	if err != nil {
		return err
	}
	// Skip if there are no pending hashes
	if len(pendingHashes) == 0 {
		phl.logger.Println("No pending hashes found; skipping loading pending hashes")
		return nil
	}

	// Fetch and parse IPFS data for each pending hash
	for _, pendingHash := range pendingHashes {
		phl.logger.Printf("Fetching and parsing IPFS data for pending hash %s", pendingHash)

		originalReport, err := phl.ipfsPort.FetchAndParseIpfs(pendingHash)
		if err != nil {
			continue
		}

		// Process each operator ID in the original report
		for _, operatorID := range operatorIDs {
			phl.logger.Printf("Saving report data for operator ID %s", operatorID.String())

			// Get the data for the operator ID
			data, exists := originalReport.Operators[operatorID.String()]
			if !exists {
				// Skip if the operator ID is not found in the report
				phl.logger.Printf("Operator ID %s not found in the original report", operatorID.String())
				continue
			}

			report := domain.Report{
				Frame:     originalReport.Frame,
				Data:      data,
				Threshold: originalReport.Threshold,
			}

			// Save the report data for the operator ID
			if err := phl.storagePort.SaveReport(operatorID, report); err != nil {
				continue
			}
		}

		// Remove the pending hash from storage
		if err := phl.storagePort.DeletePendingHash(pendingHash); err != nil {
			phl.logger.Printf("Failed to remove pending hash %s: %v", pendingHash, err)
			continue
		}
	}

	return nil
}
