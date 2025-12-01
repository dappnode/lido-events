package services

import (
	"context"
	"lido-events/internal/application/ports"
	"lido-events/internal/logger"
	"sync"
	"time"
)

type AllHashesLoader struct {
	performanceStorage ports.PerformanceStorage
	notifierPort       ports.NotifierPort
	csfeeDistributor   ports.CsFeeDistributorPort
	ipfsPort           ports.IpfsPort
	servicePrefix      string
}

func NewAllHashesLoader(performanceStorage ports.PerformanceStorage, notifierPort ports.NotifierPort, csfeeDistributor ports.CsFeeDistributorPort, ipfsPort ports.IpfsPort) *AllHashesLoader {
	return &AllHashesLoader{
		performanceStorage,
		notifierPort,
		csfeeDistributor,
		ipfsPort,
		"HashesLoader",
	}
}

func (phl *AllHashesLoader) LoadHashesCron(ctx context.Context, interval time.Duration, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine finishes
	wg.Add(1)       // Increment the WaitGroup counter

	// Execute immediately on startup
	if err := phl.loadHashes(); err != nil {
		logger.InfoWithPrefix(phl.servicePrefix, "Error loading hashes: %v", err)
	}

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Call the load method periodically
			if err := phl.loadHashes(); err != nil {
				logger.InfoWithPrefix(phl.servicePrefix, "Error loading hashes: %v", err)
				continue
			}
		case <-ctx.Done():
			// Stop the periodic load if the context is canceled
			logger.Info("Stopping periodic cron for loading CIDs")
			return
		}
	}
}

// loadHashes fetches all log CIDs from the CsFeeDistributor,
// retrieves each report from IPFS, and stores it in the performance DB.
func (phl *AllHashesLoader) loadHashes() error {
	logger.DebugWithPrefix(phl.servicePrefix, "Starting hashes load")

	// 1. Get all log CIDs from the CsFeeDistributor port.
	ctx := context.Background()
	logCids, err := phl.csfeeDistributor.GetAllLogCids(ctx)
	if err != nil {
		logger.ErrorWithPrefix(phl.servicePrefix, "Error getting log CIDs from CsFeeDistributor: %v", err)
		return err
	}

	if len(logCids) == 0 {
		logger.DebugWithPrefix(phl.servicePrefix, "No log CIDs returned from CsFeeDistributor")
		return nil
	}

	// 2. For each CID: fetch and parse from IPFS, then store in performance storage.
	for _, cid := range logCids {
		report, err := phl.ipfsPort.FetchAndParseIpfs(cid)
		if err != nil {
			logger.ErrorWithPrefix(phl.servicePrefix, "Failed to fetch/parse IPFS report for CID %s: %v", cid, err)
			continue
		}

		if err := phl.performanceStorage.SaveReport(cid, &report); err != nil {
			logger.ErrorWithPrefix(phl.servicePrefix, "Failed to save performance report for CID %s: %v", cid, err)
			continue
		}

		logger.DebugWithPrefix(phl.servicePrefix, "Successfully stored performance report for CID %s", cid)
	}

	logger.DebugWithPrefix(phl.servicePrefix, "Finished hashes load")
	return nil
}
