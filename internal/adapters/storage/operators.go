package storage

import (
	"errors"
	"lido-events/internal/application/domain"
	"math/big"
)

// SaveOperatorId adds a new operator ID to the storage if it doesn’t already exist.
func (fs *Storage) SaveOperatorId(operatorID string) error {
	db, err := fs.LoadDatabase()
	if err != nil {
		return err
	}

	if db.Operators.OperatorDetails == nil {
		db.Operators.OperatorDetails = make(map[string]OperatorDetails)
	}

	// Check if the operator ID already exists
	if _, exists := db.Operators.OperatorDetails[operatorID]; !exists {
		// Initialize an empty OperatorDetails if this is a new operator ID
		db.Operators.OperatorDetails[operatorID] = OperatorDetails{
			Performance:  make(map[string]domain.Report),
			ExitRequests: make(map[string]domain.ExitRequest),
		}
		return fs.SaveDatabase(db)
	}

	return nil // Operator ID already exists, no changes needed
}

// GetOperatorIds retrieves a list of all operator IDs in storage.
func (fs *Storage) GetOperatorIds() ([]*big.Int, error) {
	db, err := fs.LoadDatabase()
	if err != nil {
		return nil, err
	}

	var operatorIDs []*big.Int
	for opID := range db.Operators.OperatorDetails {
		operatorID := new(big.Int)
		operatorID, success := operatorID.SetString(opID, 10)
		if !success {
			return nil, errors.New("failed to convert operator ID to big.Int")
		}
		operatorIDs = append(operatorIDs, operatorID)
	}

	return operatorIDs, nil
}
