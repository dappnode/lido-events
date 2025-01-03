package storage

import (
	"fmt"
	"lido-events/internal/application/domain"
	"math/big"
)

// SaveExitRequests saves multiple exit requests for a specific operator ID.
func (fs *Storage) SaveExitRequests(operatorID *big.Int, requests domain.ExitRequests) error {
	db, err := fs.LoadDatabase()
	if err != nil {
		return err
	}

	opID := operatorID.String()
	// Initialize OperatorData if it doesn't exist
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

	// Add or update each exit request in the operator's ExitRequests map
	for validatorIndex, exitRequest := range requests {
		operatorData.ExitRequests[validatorIndex] = exitRequest
	}

	// Save the updated OperatorData back to the database
	db.Operators[opID] = operatorData
	return fs.SaveDatabase(db)
}

// SaveExitRequest saves an individual exit request for a specific operator ID and validator index.
func (fs *Storage) SaveExitRequest(operatorID *big.Int, validatorIndex string, exitRequest domain.ExitRequest) error {
	db, err := fs.LoadDatabase()
	if err != nil {
		return err
	}

	opID := operatorID.String()
	// Initialize OperatorData if it doesn't exist
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

	// Add or update the exit request for the given validator index
	operatorData.ExitRequests[validatorIndex] = exitRequest
	db.Operators[opID] = operatorData

	return fs.SaveDatabase(db)
}

// GetExitRequests retrieves all exit requests for a specific operator ID.
func (fs *Storage) GetExitRequests(operatorID string) (domain.ExitRequests, error) {
	db, err := fs.LoadDatabase()
	if err != nil {
		return nil, err
	}

	operatorData, exists := db.Operators[operatorID]
	if !exists {
		return nil, nil
	}

	return operatorData.ExitRequests, nil
}

// UpdateExitRequestStatus updates the status of a specific exit request for a validator index.
func (fs *Storage) UpdateExitRequestStatus(operatorID string, validatorIndex string, status domain.ValidatorStatus) error {
	db, err := fs.LoadDatabase()
	if err != nil {
		return err
	}

	operatorData, exists := db.Operators[operatorID]
	if !exists {
		return fmt.Errorf("operator ID %s not found", operatorID)
	}

	// Check if the exit request exists for the given validatorIndex
	exitRequest, exists := operatorData.ExitRequests[validatorIndex]
	if !exists {
		return fmt.Errorf("exit request not found for validator index %s", validatorIndex)
	}

	// Update the status and save back the modified exit request
	exitRequest.Status = status
	operatorData.ExitRequests[validatorIndex] = exitRequest
	db.Operators[operatorID] = operatorData

	return fs.SaveDatabase(db)
}

// DeleteExitRequest removes an exit request for a specific operator ID and validator index.
func (fs *Storage) DeleteExitRequest(operatorID string, validatorIndex string) error {
	db, err := fs.LoadDatabase()
	if err != nil {
		return err
	}

	operatorData, exists := db.Operators[operatorID]
	if !exists {
		return fmt.Errorf("operator ID %s not found", operatorID)
	}

	// Check if the exit request exists for the given validatorIndex
	_, exists = operatorData.ExitRequests[validatorIndex]
	if !exists {
		return fmt.Errorf("exit request not found for validator index %s", validatorIndex)
	}

	// Delete the exit request
	delete(operatorData.ExitRequests, validatorIndex)

	// Save the updated OperatorData back to the database
	db.Operators[operatorID] = operatorData

	return fs.SaveDatabase(db)
}
