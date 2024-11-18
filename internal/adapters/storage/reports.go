package storage

import (
	"fmt"
	"lido-events/internal/application/domain"
	"math/big"
)

// SaveReport saves report data for a specific operator ID, using the frame from the report.
func (fs *Storage) SaveReport(operatorID *big.Int, report domain.Report) error {
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
			Reports:      make(domain.Reports),
			ExitRequests: make(domain.ExitRequests),
		}
	}

	// Construct the frame key as "startFrame-endFrame" from report.Frame
	frameKey := fmt.Sprintf("%d-%d", report.Frame[0], report.Frame[1])

	// Update the report data for the specified frame
	operatorData.Reports[frameKey] = report
	db.Operators[opID] = operatorData

	return fs.SaveDatabase(db)
}

// GetReports retrieves all report data for a specific operator ID.
func (fs *Storage) GetReports(operatorID *big.Int) (domain.Reports, error) {
	db, err := fs.LoadDatabase()
	if err != nil {
		return nil, err
	}

	opID := operatorID.String()
	operatorData, exists := db.Operators[opID]
	if !exists {
		return nil, fmt.Errorf("operator ID %s not found", opID)
	}

	// Return all reports for the specified operator ID
	return operatorData.Reports, nil
}
