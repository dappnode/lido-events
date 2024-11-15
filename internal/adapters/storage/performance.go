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
	if db.Operators == nil {
		db.Operators = make(map[string]OperatorData)
	}

	operatorData, exists := db.Operators[opID]
	if !exists {
		operatorData = OperatorData{
			Performance:  make(map[string]domain.Report),
			ExitRequests: make(map[string]domain.ExitRequest),
		}
	}

	// Update the performance data for the specified epoch
	operatorData.Performance[epoch] = report
	db.Operators[opID] = operatorData

	return fs.SaveDatabase(db)
}

// GetOperatorPerformance retrieves performance data for a specific operator ID within an epoch range.
func (fs *Storage) GetOperatorPerformance(operatorID *big.Int, startEpoch, endEpoch string) (domain.Reports, error) {
	db, err := fs.LoadDatabase()
	if err != nil {
		return nil, err
	}

	opID := operatorID.String()
	operatorData, exists := db.Operators[opID]
	if !exists {
		return nil, fmt.Errorf("operator ID %s not found", opID)
	}

	reportData := make(map[string]domain.Report)
	for epoch, report := range operatorData.Performance {
		if epoch >= startEpoch && epoch <= endEpoch {
			reportData[epoch] = report
		}
	}
	return reportData, nil
}
