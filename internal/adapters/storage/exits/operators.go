package exits

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
			ExitsRequests: ExitsRequests{
				Exits:              make(domain.ExitRequests), // Initialize Exits map
				LastProcessedBlock: 0,
			},
		}

		// Save changes to disk
		if err := fs.SaveDatabase(db); err != nil {
			return err
		}
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

// GetValidatorExitRequestLastProcessedBlock retrieves the last processed block for the ValidatorExitRequest event for a specific operator ID.
func (fs *Storage) GetValidatorExitRequestLastProcessedBlock(operatorID *big.Int) (uint64, error) {
	db, err := fs.LoadDatabase()
	if err != nil {
		return 0, err
	}

	opID := operatorID.String()
	operatorData, exists := db.Operators[opID]
	if !exists {
		return 0, fmt.Errorf("operator ID %s not found", opID)
	}

	return operatorData.ExitsRequests.LastProcessedBlock, nil
}

// SaveValidatorExitRequestLastProcessedBlock updates the last processed block for the ValidatorExitRequest event for a specific operator ID.
func (fs *Storage) SaveValidatorExitRequestLastProcessedBlock(operatorID *big.Int, block uint64) error {
	db, err := fs.LoadDatabase()
	if err != nil {
		return err

	}

	opID := operatorID.String()
	operatorData, exists := db.Operators[opID]
	if !exists {
		return fmt.Errorf("operator ID %s not found", opID)
	}

	operatorData.ExitsRequests.LastProcessedBlock = block
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
	if db.Operators == nil {
		db.Operators = make(map[string]OperatorData)
	}

	operatorData, exists := db.Operators[opID]
	if !exists {
		operatorData = OperatorData{
			ExitsRequests: ExitsRequests{
				Exits: make(domain.ExitRequests),
			},
		}
	} else if operatorData.ExitsRequests.Exits == nil {
		operatorData.ExitsRequests.Exits = make(domain.ExitRequests)
	}

	// Add or update the exit request for the given validator index
	operatorData.ExitsRequests.Exits[validatorIndex] = exitRequest
	db.Operators[opID] = operatorData

	return fs.SaveDatabase(db)
}

// GetExitRequests retrieves all exit requests for a specific operator ID.
func (fs *Storage) GetExitRequests(operatorID *big.Int) (domain.ExitRequests, error) {
	db, err := fs.LoadDatabase()
	if err != nil {
		return nil, err
	}

	opID := operatorID.String()
	operatorData, exists := db.Operators[opID]
	if !exists {
		return nil, fmt.Errorf("operator ID %s not found", opID)
	}

	return operatorData.ExitsRequests.Exits, nil
}

// Update ExitRequestStatus updates the status of a specific exit request for a validator index.
func (fs *Storage) UpdateExitRequestStatus(operatorID *big.Int, validatorIndex string, status domain.ValidatorStatus) error {
	db, err := fs.LoadDatabase()
	if err != nil {
		return err
	}

	opID := operatorID.String()
	operatorData, exists := db.Operators[opID]
	if !exists {
		return fmt.Errorf("operator ID %s not found", opID)
	}

	// Check if the exit request exists for the given validatorIndex
	exitRequest, exists := operatorData.ExitsRequests.Exits[validatorIndex]
	if !exists {
		return fmt.Errorf("exit request not found for validator index %s", validatorIndex)
	}

	// Update the status and save back the modified exit request
	exitRequest.Status = status
	operatorData.ExitsRequests.Exits[validatorIndex] = exitRequest
	db.Operators[opID] = operatorData

	return fs.SaveDatabase(db)
}

// DeleteExitRequest removes an exit request for a specific operator ID and validator index.
func (fs *Storage) DeleteExitRequest(operatorID *big.Int, validatorIndex string) error {
	db, err := fs.LoadDatabase()
	if err != nil {
		return err
	}

	opID := operatorID.String()
	operatorData, exists := db.Operators[opID]
	if !exists {
		return fmt.Errorf("operator ID %s not found", opID)
	}

	// Check if the exit request exists for the given validatorIndex
	_, exists = operatorData.ExitsRequests.Exits[validatorIndex]
	if !exists {
		return fmt.Errorf("exit request not found for validator index %s", validatorIndex)
	}

	// Delete the exit request
	delete(operatorData.ExitsRequests.Exits, validatorIndex)

	// Save the updated OperatorData back to the database
	db.Operators[opID] = operatorData

	return fs.SaveDatabase(db)
}
