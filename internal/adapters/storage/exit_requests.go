package storage

import (
	"fmt"
	"lido-events/internal/application/domain"
	"math/big"
)

// SaveExitRequests saves multiple exit requests for a specific operator ID.
func (fs *Storage) SaveExitRequests(operatorID *big.Int, requests map[string]domain.ExitRequest) error {
	db, err := fs.LoadDatabase()
	if err != nil {
		return err
	}

	opID := operatorID.String()
	// Initialize OperatorDetails if it doesn't exist
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

	// Add or update each exit request in the operator's ExitRequests map
	for validatorIndex, exitRequest := range requests {
		operatorDetails.ExitRequests[validatorIndex] = exitRequest
	}

	// Save the updated OperatorDetails back to the database
	db.Operators.OperatorDetails[opID] = operatorDetails
	return fs.SaveDatabase(db)
}

// SaveExitRequest saves an individual exit request for a specific operator ID and validator index
func (fs *Storage) SaveExitRequest(operatorID *big.Int, validatorIndex string, exitRequest domain.ExitRequest) error {
	db, err := fs.LoadDatabase()
	if err != nil {
		return err
	}

	opID := operatorID.String()
	// Initialize OperatorDetails if it doesn't exist
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

	// Add or update the exit request for the given validator index
	operatorDetails.ExitRequests[validatorIndex] = exitRequest
	db.Operators.OperatorDetails[opID] = operatorDetails

	return fs.SaveDatabase(db)
}

// GetExitRequests retrieves all exit requests for a specific operator ID.
func (fs *Storage) GetExitRequests(operatorID string) (map[string]domain.ExitRequest, error) {
	db, err := fs.LoadDatabase()
	if err != nil {
		return nil, err
	}

	operatorDetails, exists := db.Operators.OperatorDetails[operatorID]
	if !exists {
		return nil, nil
	}

	return operatorDetails.ExitRequests, nil
}

// UpdateExitRequestStatus updates the status of a specific exit request for a validator index.
func (fs *Storage) UpdateExitRequestStatus(operatorID string, validatorIndex string, status domain.ValidatorStatus) error {
	db, err := fs.LoadDatabase()
	if err != nil {
		return err
	}

	operatorDetails, exists := db.Operators.OperatorDetails[operatorID]
	if !exists {
		return fmt.Errorf("operator ID %s not found", operatorID)
	}

	// Check if the exit request exists for the given validatorIndex
	exitRequest, exists := operatorDetails.ExitRequests[validatorIndex]
	if !exists {
		return fmt.Errorf("exit request not found for validator index %s", validatorIndex)
	}

	// Update the status and save back the modified exit request
	exitRequest.Status = status
	operatorDetails.ExitRequests[validatorIndex] = exitRequest
	db.Operators.OperatorDetails[operatorID] = operatorDetails

	return fs.SaveDatabase(db)
}

// SaveLastProcessedBlock saves the global last processed epoch in operators data.
func (fs *Storage) SaveLastProcessedBlock(epoch uint64) error {
	db, err := fs.LoadDatabase()
	if err != nil {
		return err
	}

	// Update the global last processed epoch in operators data
	db.Operators.LastProcessedBlock = epoch
	return fs.SaveDatabase(db)
}

// GetLastProcessedBlock retrieves the global last processed epoch from operators data.
func (fs *Storage) GetLastProcessedBlock() (uint64, error) {
	db, err := fs.LoadDatabase()
	if err != nil {
		return 0, err
	}

	return db.Operators.LastProcessedBlock, nil
}
