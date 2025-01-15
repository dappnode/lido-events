package storage

import (
	"fmt"
	"lido-events/internal/application/domain"
	"math/big"
)

// SaveOperatorId adds a new operator ID to the storage if it doesn’t already exist.
func (fs *Storage) SaveOperatorId(operatorID string) error {
	db, err := fs.LoadDatabase()
	if err != nil {
		return err
	}

	if db.Operators == nil {
		db.Operators = make(map[string]OperatorData)
	}

	// Check if the operator ID already exists
	if _, exists := db.Operators[operatorID]; !exists {
		// Initialize an empty OperatorData if this is a new operator ID
		db.Operators[operatorID] = OperatorData{
			Reports:      make(domain.Reports),
			ExitRequests: make(domain.ExitRequests),
		}
		if err := fs.SaveDatabase(db); err != nil {
			return err
		}
		fs.notifyOperatorIdListeners() // Notify listeners of the change
	}

	return nil // Operator ID already exists, no changes needed
}

// DeleteOperator removes an operator ID from the storage if it exists.
func (fs *Storage) DeleteOperator(operatorID string) error {
	db, err := fs.LoadDatabase()
	if err != nil {
		return err
	}

	// Check if the operator ID exists
	if _, exists := db.Operators[operatorID]; exists {
		// Remove the operator from the storage
		delete(db.Operators, operatorID)
		if err := fs.SaveDatabase(db); err != nil {
			return err
		}
		fs.notifyOperatorIdListeners() // Notify listeners of the change
		return nil
	}

	// Return an error if the operator ID does not exist
	return fmt.Errorf("operator ID %s not found", operatorID)
}

// GetOperatorIds retrieves a list of all operator IDs in storage.
func (fs *Storage) GetOperatorIds() ([]*big.Int, error) {
	db, err := fs.LoadDatabase()
	if err != nil {
		return nil, err
	}

	var operatorIDs []*big.Int
	for opID := range db.Operators {
		operatorID := new(big.Int)
		operatorID, success := operatorID.SetString(opID, 10)
		if !success {
			return nil, fmt.Errorf("failed to convert operator ID %s to big.Int", opID)
		}
		operatorIDs = append(operatorIDs, operatorID)
	}

	return operatorIDs, nil
}

// RegisterOperatorIdListener registers a channel to receive updates when operator IDs change.
func (fs *Storage) RegisterOperatorIdListener() chan []*big.Int {
	updateChan := make(chan []*big.Int, 1)
	fs.operatorIdListeners = append(fs.operatorIdListeners, updateChan)
	return updateChan
}

// notifyOperatorIdListeners sends updates to all registered listeners of operator ID changes.
func (fs *Storage) notifyOperatorIdListeners() {
	operatorIds, err := fs.GetOperatorIds()
	if err != nil {
		return
	}

	for _, listener := range fs.operatorIdListeners {
		select {
		case listener <- operatorIds:
		default:
			// Ignore if channel is full to prevent blocking
		}
	}
}
