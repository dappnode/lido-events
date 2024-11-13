package storage

import (
	"errors"
	"lido-events/internal/application/domain"
	"math/big"
)

// SaveOperatorId adds a new operator ID to the storage if it doesn’t already exist.
func (fs *Storage) SaveOperatorId(operatorID string) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

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
		if err := fs.SaveDatabase(db); err != nil {
			return err
		}
		fs.notifyOperatorIdListeners() // Notify listeners of the change
	}

	return nil // Operator ID already exists, no changes needed
}

// GetOperatorIds retrieves a list of all operator IDs in storage.
func (fs *Storage) GetOperatorIds() ([]*big.Int, error) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()

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

// RegisterOperatorIdListener registers a channel to receive updates when operator IDs change.
func (fs *Storage) RegisterOperatorIdListener() chan []*big.Int {
	fs.mu.Lock()
	defer fs.mu.Unlock()

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
