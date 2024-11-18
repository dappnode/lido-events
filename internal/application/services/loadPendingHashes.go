package services

import (
	"context"
	"fmt"
	"lido-events/internal/application/domain"
	"lido-events/internal/application/ports"
	"log"
	"time"
)

type PendingHashesLoader struct {
	storagePort ports.StoragePort
	ipfsPort    ports.IpfsPort
}

func NewPendingHashesLoader(storagePort ports.StoragePort, ipfsPort ports.IpfsPort) *PendingHashesLoader {
	return &PendingHashesLoader{
		storagePort,
		ipfsPort,
	}
}

// LoadPendingHashesCron starts a periodic service to load pending hashes from IPFS
func (phl *PendingHashesLoader) LoadPendingHashesCron(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Call the scan method periodically
			phl.LoadPendingHashes()
		case <-ctx.Done():
			// Stop the periodic scan if the context is canceled
			log.Println("Stopping periodic cron for loading pending CIDs")
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
	// skip if there are no operator IDs
	if len(operatorIDs) == 0 {
		return nil
	}
	// Get all pending hashes from the storage
	pendingHashes, err := phl.storagePort.GetPendingHashes()
	if err != nil {
		return err
	}

	// Fetch and parse IPFS data for each pending hash
	for _, pendingHash := range pendingHashes {
		originalReport, err := phl.ipfsPort.FetchAndParseIpfs(pendingHash)
		if err != nil {
			return err
		}

		// find the operator IDs in the original report
		for _, operatorID := range operatorIDs {
			// get the data for the operator ID
			data, exists := originalReport.Operators[operatorID.String()]
			if !exists {
				// skip if the operator ID is not found
				fmt.Printf("Operator ID %s not found in the original report\n", operatorID.String())
				continue
			}

			// build the report data with the type domain.report
			report := domain.Report{
				Frame:     originalReport.Frame,
				Data:      data,
				Threshold: originalReport.Threshold,
			}

			// save the report data for the operator ID
			err = phl.storagePort.SaveReport(operatorID, report)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
