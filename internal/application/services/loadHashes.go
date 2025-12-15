package services

import (
	"context"
	"fmt"
	"lido-events/internal/application/domain"
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
	if err := phl.loadHashes(ctx); err != nil {
		logger.InfoWithPrefix(phl.servicePrefix, "Error loading hashes: %v", err)
	}

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Call the load method periodically
			if err := phl.loadHashes(ctx); err != nil {
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
func (phl *AllHashesLoader) loadHashes(ctx context.Context) error {
	logger.DebugWithPrefix(phl.servicePrefix, "Starting hashes load")

	// 0. Fetch existing log CIDs from performance storage and build a fast lookup set
	existingLogCids, err := phl.performanceStorage.GetUniqueLogCids()
	if err != nil {
		logger.ErrorWithPrefix(phl.servicePrefix, "Error getting existing log CIDs from performance storage: %v", err)
		return err
	}

	existingSet := make(map[string]struct{}, len(existingLogCids))
	for _, cid := range existingLogCids {
		existingSet[cid] = struct{}{}
	}

	// 1. Get all log CIDs from the CsFeeDistributor port.
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
	var newReports []domain.Report

	for _, cid := range logCids {
		// Skip CIDs that are already present in the performance DB
		if _, ok := existingSet[cid]; ok {
			logger.DebugWithPrefix(phl.servicePrefix, "Skipping CID %s: already stored in performance DB", cid)
			continue
		}

		report, err := phl.ipfsPort.FetchAndParseIpfs(cid)
		if err != nil {
			logger.ErrorWithPrefix(phl.servicePrefix, "Failed to fetch/parse IPFS report for CID %s: %v", cid, err)
			continue
		}

		if err := phl.performanceStorage.SaveReport(cid, &report); err != nil {
			logger.ErrorWithPrefix(phl.servicePrefix, "Failed to save performance report for CID %s: %v", cid, err)
			continue
		}

		// keep track of newly stored reports to send a single
		// consolidated notification at the end of the run
		newReports = append(newReports, report)

		// Mark CID as stored to avoid processing duplicates within the same run
		existingSet[cid] = struct{}{}

		logger.DebugWithPrefix(phl.servicePrefix, "Successfully stored performance report for CID %s", cid)
	}

	// 3. Send a single notification if at least one new report was stored.
	if len(newReports) > 0 {
		message := buildPerformanceNotificationMessage(newReports)
		if err := phl.notifierPort.SendNewPerformanceReport(message); err != nil {
			logger.ErrorWithPrefix(phl.servicePrefix, "Failed to notify about new performance reports: %v", err)
		}
	}

	logger.DebugWithPrefix(phl.servicePrefix, "Finished hashes load")
	return nil
}

// buildPerformanceNotificationMessage builds a human-readable summary of how
// validators have performed during the last processed frames.
func buildPerformanceNotificationMessage(reports []domain.Report) string {
	if len(reports) == 0 {
		return "New performance report available."
	}

	var (
		globalValidators              int
		globalUnderThreshold          int
		globalSlashed                 int
		globalPerformanceSum          float64
		firstFrameStart, lastFrameEnd int
	)

	for i, r := range reports {
		if i == 0 {
			firstFrameStart = r.Frame[0]
		}
		lastFrameEnd = r.Frame[1]

		for _, opData := range r.Operators {
			for _, v := range opData.Validators {
				globalValidators++
				globalPerformanceSum += v.Performance
				if v.Slashed {
					globalSlashed++
				}
				if v.Performance < v.Threshold {
					globalUnderThreshold++
				}
			}
		}
	}

	avgPerformance := 0.0
	if globalValidators > 0 {
		avgPerformance = globalPerformanceSum / float64(globalValidators)
	}

	framesText := "frame"
	if len(reports) > 1 {
		framesText = "frames"
	}

	return fmt.Sprintf(
		"New validator performance %s %d-%d: %d validators (avg performance %.4f), %d below threshold, %d slashed.",
		framesText,
		firstFrameStart,
		lastFrameEnd,
		globalValidators,
		avgPerformance,
		globalUnderThreshold,
		globalSlashed,
	)
}
