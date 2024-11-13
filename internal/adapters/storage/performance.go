package storage

import (
	"fmt"
	"lido-events/internal/application/domain"
	"math/big"
)

// SaveOperatorPerformance saves performance data for a specific operator ID and epoch.
func (fs *Storage) SaveOperatorPerformance(operatorID *big.Int, epoch string, report domain.Report) error {
	db, err := fs.LoadDatabase()
	if err != nil {
		return err
	}

	opID := operatorID.String()
	if db.Operators.OperatorDetails == nil {
		db.Operators.OperatorDetails = make(map[string]OperatorDetails)
	}

	operatorDetails, exists := db.Operators.OperatorDetails[opID]
	if !exists {
		operatorDetails = OperatorDetails{
			Performance:  make(map[string]domain.Report),
			ExitRequests: make(map[string]domain.ExitRequest),
		}
	}

	// Update the performance data for the specified epoch
	operatorDetails.Performance[epoch] = report
	db.Operators.OperatorDetails[opID] = operatorDetails

	return fs.SaveDatabase(db)
}

// GetOperatorPerformance retrieves performance data for a specific operator ID within an epoch range.
func (fs *Storage) GetOperatorPerformance(operatorID *big.Int, startEpoch, endEpoch string) (domain.Reports, error) {
	db, err := fs.LoadDatabase()
	if err != nil {
		return nil, err
	}

	opID := operatorID.String()
	operatorDetails, exists := db.Operators.OperatorDetails[opID]
	if !exists {
		return nil, fmt.Errorf("operator ID %s not found", opID)
	}

	reportData := make(map[string]domain.Report)
	for epoch, report := range operatorDetails.Performance {
		if epoch >= startEpoch && epoch <= endEpoch {
			reportData[epoch] = report
		}
	}
	return reportData, nil
}

// AddPendingHash adds a hash to the pendingHashes array in the global OperatorsData struct.
func (fs *Storage) AddPendingHash(hash string) error {
	db, err := fs.LoadDatabase()
	if err != nil {
		return err
	}

	// Ensure pendingHashes array exists
	if db.Operators.PendingHashes == nil {
		db.Operators.PendingHashes = []string{}
	}

	// Add the new hash if it does not already exist
	for _, h := range db.Operators.PendingHashes {
		if h == hash {
			return fmt.Errorf("hash %s already exists in pendingHashes", hash)
		}
	}

	db.Operators.PendingHashes = append(db.Operators.PendingHashes, hash)
	return fs.SaveDatabase(db)
}

// DeletePendingHash removes a hash from the pendingHashes array in the global OperatorsData struct.
func (fs *Storage) DeletePendingHash(hash string) error {
	db, err := fs.LoadDatabase()
	if err != nil {
		return err
	}

	// Find and remove the hash
	found := false
	for i, h := range db.Operators.PendingHashes {
		if h == hash {
			db.Operators.PendingHashes = append(db.Operators.PendingHashes[:i], db.Operators.PendingHashes[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("hash not found in pendingHashes")
	}

	return fs.SaveDatabase(db)
}
