package services

import (
	"context"
	"fmt"
	"lido-events/internal/application/domain"
	"lido-events/internal/application/ports"
	"lido-events/internal/logger"
	"math/big"
	"strings"
	"sync"
	"time"
)

type PendingHashesLoader struct {
	storagePort    ports.ExitsStorage
	notifierPort   ports.NotifierPort
	ipfsPort       ports.IpfsPort
	minGenesisTime uint64
	mu             sync.Mutex
	servicePrefix  string
}

func NewPendingHashesLoader(storagePort ports.ExitsStorage, notifierPort ports.NotifierPort, ipfsPort ports.IpfsPort, minGenesisTime uint64) *PendingHashesLoader {
	return &PendingHashesLoader{
		storagePort,
		notifierPort,
		ipfsPort,
		minGenesisTime,
		sync.Mutex{},
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
	if err := phl.LoadPendingHashes(false); err != nil {
		logger.InfoWithPrefix(phl.servicePrefix, "Error loading pending hashes: %v", err)
	}

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Call the load method periodically
			if err := phl.LoadPendingHashes(false); err != nil {
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

func (phl *PendingHashesLoader) LoadPendingHashes(giveUp bool) error {
	// Lock the mutex
	phl.mu.Lock()
	defer phl.mu.Unlock()

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

	// Iterate through the operator IDs
	for _, operatorID := range operatorIDs {
		logger.DebugWithPrefix(phl.servicePrefix, "Processing operator ID %s", operatorID.String())
		// Get all pending hashes
		pendingHashes, err := phl.storagePort.GetPendingHashes(operatorID)
		if err != nil {
			logger.ErrorWithPrefix(phl.servicePrefix, "Failed to get pending hashes", err)
			return err
		}
		if len(pendingHashes) == 0 {
			logger.InfoWithPrefix(phl.servicePrefix, "No pending hashes found; skipping loading pending hashes")
			continue
		}

		// Process each pending hash
		for _, pendingHash := range pendingHashes {
			logger.DebugWithPrefix(phl.servicePrefix, "Fetching and parsing IPFS data for pending hash %s", pendingHash)

			originalReport, err := phl.ipfsPort.FetchAndParseIpfs(pendingHash)
			if err != nil {
				logger.ErrorWithPrefix(phl.servicePrefix, "Failed to fetch and parse IPFS data for pending hash %s, skipping hash: %v", pendingHash, err)
				// giveUp flag is set to true when we dont want to retry fetching the data of all pending hashes.
				// This is useful if we want this function to finish somewhat quickly and not wait for all pending hashes to be fetched.
				if giveUp {
					logger.DebugWithPrefix(phl.servicePrefix, "Called with 'giveUp' flag set to true, skipping the rest of the pending hashes")
					return nil
				}
				continue
			}

			logger.DebugWithPrefix(phl.servicePrefix, "Processing data for operator ID %s", operatorID.String())

			data, exists := originalReport.Operators[operatorID.String()]
			if !exists {
				logger.WarnWithPrefix(phl.servicePrefix, "Operator ID %s not found in the original report, deleting pending hash and skipping", operatorID.String())
				// Remove the pending hash
				if err := phl.storagePort.DeletePendingHash(operatorID, pendingHash); err != nil {
					logger.ErrorWithPrefix(phl.servicePrefix, "Failed to delete pending hash %s: %v", pendingHash, err)
				}
				continue
			}

			// Check performance and send notifications if needed
			// Only sends the notification if the report is from the last 24 hours. Calculated from genesis timestamp & report epoch
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

			// Remove the pending hash
			if err := phl.storagePort.DeletePendingHash(operatorID, pendingHash); err != nil {
				logger.ErrorWithPrefix(phl.servicePrefix, "Failed to delete pending hash %s: %v", pendingHash, err)
			}
		}
	}

	return nil
}

func (phl *PendingHashesLoader) CheckAndNotifyPerformance(operatorID *big.Int, validators map[string]domain.Validator, originalReport domain.Report) {
	operatorData, exists := originalReport.Operators[operatorID.String()]
	if !exists {
		logger.WarnWithPrefix(phl.servicePrefix, "Operator ID %s not found in the original report, skipping", operatorID.String())
		return
	}

	// Check if the operator is stuck, if so skip performance evaluation and send a notification
	if operatorData.Stuck {
		logger.WarnWithPrefix(phl.servicePrefix, "Operator ID %s is in a 'stuck' state; skipping performance evaluation", operatorID.String())
		// send notification
		message := fmt.Sprintf(
			"- 🚨 Operator ID: %s, from epoch %d to epoch %d, is in a 'stuck' state. Please check the operator's performance.",
			operatorID.String(), originalReport.Frame[0], originalReport.Frame[1],
		)
		if err := phl.notifierPort.SendNotification(message); err != nil {
			logger.ErrorWithPrefix(phl.servicePrefix, "Failed to send stuck state notification: %v", err)
		}
		return
	}

	// Initialize slices to collect messages for bad and good-performing validators
	var badValidators []string
	var goodValidators []string

	// Track if there are any bad-performing validators
	anyBadPerformance := false
	anyGoodPerformance := false

	// Iterate through the validators and check performance
	for validatorID, validator := range operatorData.Validators {
		if validator.Perf.Assigned > 60 { // Only try to notify if the validator has had duties to at least 60 epochs
			performance := float64(validator.Perf.Included) / float64(validator.Perf.Assigned) * 100
			if performance < originalReport.Threshold*100 {
				// Mark that we have at least one bad-performing validator
				anyBadPerformance = true

				// Collect bad-performing validator details
				badValidators = append(badValidators, fmt.Sprintf("| %s | %.2f%% |", validatorID, performance))
			} else if performance >= originalReport.Threshold*100 {
				// Only categorize as good if performance meets or exceeds the threshold
				anyGoodPerformance = true
				goodValidators = append(goodValidators, fmt.Sprintf("| %s | %.2f%% |", validatorID, performance))
			} else {
				logger.WarnWithPrefix(phl.servicePrefix, "Validator %s has an invalid performance value: %.2f%%", validatorID, performance)
			}
		}
	}

	// Calculate the report timestamp from the frame data
	reportTimestamp := phl.minGenesisTime + uint64(originalReport.Frame[1])*384

	// Get the current time
	currentTime := uint64(time.Now().Unix())

	// Calculate the time difference in seconds
	timeDiff := currentTime - reportTimestamp

	// Convert the time difference to days and hours
	days := timeDiff / 86400           // 86400 seconds in a day
	hours := (timeDiff % 86400) / 3600 // 3600 seconds in an hour

	// Format the time ago string
	timeAgo := fmt.Sprintf("%d days and %d hours ago", days, hours)

	// Check if the report is older than 24h (86400 seconds)
	// Only skip if it's older than 24h
	if timeDiff > 86400 {
		logger.DebugWithPrefix(phl.servicePrefix, "Skipping notification for operator ID %s, report (epoch %d) is older than 24h", operatorID.String(), originalReport.Frame[1])
		return
	}

	// If at least one validator had bad performance, send a bad performance notification
	if anyBadPerformance {
		// Create the bad performance notification message
		message := fmt.Sprintf(
			"- 🚨 Operator ID: %s, from epoch %d to epoch %d (%s), the following validators have had bad performance below the threshold of %.2f%%:\n",
			operatorID.String(), originalReport.Frame[0], originalReport.Frame[1], timeAgo, originalReport.Threshold*100, // Multiply by 100 to convert to percentage
		)
		// Add table header
		message += "| Validator ID | Performance (%) |\n"
		message += "|--------------|-----------------|\n"
		// Add bad validators' performance
		message += strings.Join(badValidators, "\n") + "\n"

		// Log the notification message
		logger.InfoWithPrefix(phl.servicePrefix, "Sending bad performance notification for report epoch %d", originalReport.Frame[1])

		// Send the notification
		if err := phl.notifierPort.SendNotification(message); err != nil {
			logger.ErrorWithPrefix(phl.servicePrefix, "Failed to send bad performance notification: %v", err)
		}
	}

	// If at least one validator had good performance, send a good performance notification
	if anyGoodPerformance {
		// Create the good performance notification message
		message := fmt.Sprintf(
			"- ✅ Operator ID: %s, from epoch %d to epoch %d (%s), the following validators performed well (above the threshold of %.2f%%):\n",
			operatorID.String(), originalReport.Frame[0], originalReport.Frame[1], timeAgo, originalReport.Threshold*100, // Multiply by 100 to convert to percentage
		)
		// Add table header
		message += "| Validator ID | Performance (%) |\n"
		message += "|--------------|-----------------|\n"
		// Add good validators' performance
		message += strings.Join(goodValidators, "\n") + "\n"

		// Log the notification message
		logger.InfoWithPrefix(phl.servicePrefix, "Sending good performance notification for report epoch %d", originalReport.Frame[1])

		// Send the notification
		if err := phl.notifierPort.SendNotification(message); err != nil {
			logger.ErrorWithPrefix(phl.servicePrefix, "Failed to send good performance notification: %v", err)
		}
	}
}
