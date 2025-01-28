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
			DistributionLogsUpdated:            DistributionLogsUpdated{},
			ExitsRequests:                      ExitsRequests{},
			WithdrawalsSubmitted:               WithdrawalsSubmitted{},
			ElRewardsStealingPenaltiesReported: ElRewardsStealingPenaltiesReported{},
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
			DistributionLogsUpdated:            DistributionLogsUpdated{},
			ExitsRequests:                      ExitsRequests{},
			WithdrawalsSubmitted:               WithdrawalsSubmitted{},
			ElRewardsStealingPenaltiesReported: ElRewardsStealingPenaltiesReported{},
		}
	}

	// Construct the frame key as "startFrame-endFrame" from report.Frame
	frameKey := fmt.Sprintf("%d-%d", report.Frame[0], report.Frame[1])

	// Update the report data for the specified frame
	operatorData.DistributionLogsUpdated.Reports[frameKey] = report
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
	return operatorData.DistributionLogsUpdated.Reports, nil
}

// GetDistributionLogLastProcessedBlock retrieves the last processed block for the DistributionLogUpdated event for a specific operator ID.
func (fs *Storage) GetDistributionLogLastProcessedBlock(operatorID *big.Int) (uint64, error) {
	db, err := fs.LoadDatabase()
	if err != nil {
		return 0, err
	}

	opID := operatorID.String()
	operatorData, exists := db.Operators[opID]
	if !exists {
		return 0, fmt.Errorf("operator ID %s not found", opID)
	}

	return operatorData.DistributionLogsUpdated.LastProcessedBlock, nil
}

// SaveDistributionLogLastProcessedBlock updates the last processed block for the DistributionLogUpdated event for a specific operator ID.
func (fs *Storage) SaveDistributionLogLastProcessedBlock(operatorID *big.Int, block uint64) error {
	db, err := fs.LoadDatabase()
	if err != nil {
		return err // Return error if database load fails
	}

	opID := operatorID.String()
	operatorData, exists := db.Operators[opID]
	if !exists {
		return fmt.Errorf("operator ID %s not found", opID) // Return error if operator ID not found
	}

	operatorData.DistributionLogsUpdated.LastProcessedBlock = block
	db.Operators[opID] = operatorData

	return fs.SaveDatabase(db)
}

// AddPendingHash adds a hash to the PendingHashes array in DistributionLogUpdated if it does not already exist.
func (fs *Storage) AddPendingHash(operatorID *big.Int, hash string) error {
	db, err := fs.LoadDatabase()
	if err != nil {
		return err
	}

	opID := operatorID.String()
	operatorData, exists := db.Operators[opID]
	if !exists {
		return fmt.Errorf("operator ID %s not found", opID)
	}

	// Check if hash already exists
	for _, h := range operatorData.DistributionLogsUpdated.PendingHashes {
		if h == hash {
			return nil // Hash already exists, no need to add
		}
	}

	// Add hash to PendingHashes
	operatorData.DistributionLogsUpdated.PendingHashes = append(operatorData.DistributionLogsUpdated.PendingHashes, hash)
	db.Operators[opID] = operatorData

	return fs.SaveDatabase(db)
}

// DeletePendingHash removes a hash from the PendingHashes array in DistributionLogUpdated.
func (fs *Storage) DeletePendingHash(operatorID *big.Int, hash string) error {
	db, err := fs.LoadDatabase()
	if err != nil {
		return err
	}

	opID := operatorID.String()
	operatorData, exists := db.Operators[opID]
	if !exists {
		return fmt.Errorf("operator ID %s not found", opID)
	}

	// Find and remove the hash
	for i, h := range operatorData.DistributionLogsUpdated.PendingHashes {
		if h == hash {
			operatorData.DistributionLogsUpdated.PendingHashes = append(operatorData.DistributionLogsUpdated.PendingHashes[:i], operatorData.DistributionLogsUpdated.PendingHashes[i+1:]...)
			db.Operators[opID] = operatorData
			return fs.SaveDatabase(db)
		}
	}

	return nil // Hash not found; nothing to delete
}

// GetPendingHashes retrieves all pending hashes from the DistributionLogUpdated event.
func (fs *Storage) GetPendingHashes(operatorID *big.Int) ([]string, error) {
	db, err := fs.LoadDatabase()
	if err != nil {
		return nil, err
	}

	opID := operatorID.String()
	operatorData, exists := db.Operators[opID]
	if !exists {
		return nil, fmt.Errorf("operator ID %s not found", opID)
	}

	return operatorData.DistributionLogsUpdated.PendingHashes, nil
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
			DistributionLogsUpdated:            DistributionLogsUpdated{},
			ExitsRequests:                      ExitsRequests{},
			WithdrawalsSubmitted:               WithdrawalsSubmitted{},
			ElRewardsStealingPenaltiesReported: ElRewardsStealingPenaltiesReported{},
		}
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

// GetElRewardsStealingPenaltyReportedLastProcessedBlock retrieves the last processed block for the ELRewardsStealingPenaltyReported event for a specific operator ID.
func (fs *Storage) GetElRewardsStealingPenaltiesReportedLastProcessedBlock(operatorID *big.Int) (uint64, error) {
	db, err := fs.LoadDatabase()
	if err != nil {
		return 0, err
	}

	opID := operatorID.String()
	operatorData, exists := db.Operators[opID]
	if !exists {
		return 0, fmt.Errorf("operator ID %s not found", opID)
	}

	return operatorData.ElRewardsStealingPenaltiesReported.LastProcessedBlock, nil
}

// SaveElRewardsStealingPenaltiesReportedLastProcessedBlock updates the last processed block for the ELRewardsStealingPenaltyReported event for a specific operator ID.
func (fs *Storage) SaveElRewardsStealingPenaltiesReportedLastProcessedBlock(operatorID *big.Int, block uint64) error {
	db, err := fs.LoadDatabase()
	if err != nil {
		return err
	}

	opID := operatorID.String()
	operatorData, exists := db.Operators[opID]
	if !exists {
		return fmt.Errorf("operator ID %s not found", opID)
	}

	operatorData.ElRewardsStealingPenaltiesReported.LastProcessedBlock = block
	db.Operators[opID] = operatorData

	return fs.SaveDatabase(db)
}

// GetElRewardsStealingPenaltiesReported retrieves all ELRewardsStealingPenaltyReported events for a specific operator ID.
func (fs *Storage) GetElRewardsStealingPenaltiesReported(operatorID *big.Int) ([]domain.CsmoduleELRewardsStealingPenaltyReported, error) {
	db, err := fs.LoadDatabase()
	if err != nil {
		return nil, err
	}

	opID := operatorID.String()
	operatorData, exists := db.Operators[opID]
	if !exists {
		return nil, fmt.Errorf("operator ID %s not found", opID)
	}

	return operatorData.ElRewardsStealingPenaltiesReported.Penalties, nil
}

// SaveElRewardsStealingPenaltyReported saves an ELRewardsStealingPenaltyReported event for a specific operator ID.
func (fs *Storage) SaveElRewardsStealingPenaltyReported(operatorID *big.Int, penalty domain.CsmoduleELRewardsStealingPenaltyReported) error {
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
			DistributionLogsUpdated:            DistributionLogsUpdated{},
			ExitsRequests:                      ExitsRequests{},
			WithdrawalsSubmitted:               WithdrawalsSubmitted{},
			ElRewardsStealingPenaltiesReported: ElRewardsStealingPenaltiesReported{},
		}
	}

	// Append the new penalty to the existing list
	operatorData.ElRewardsStealingPenaltiesReported.Penalties = append(operatorData.ElRewardsStealingPenaltiesReported.Penalties, penalty)
	db.Operators[opID] = operatorData

	return fs.SaveDatabase(db)
}

// GetWithdrawalsSubmittedLastProcessedBlock retrieves the last processed block for the WithdrawalsSubmitted event for a specific operator ID.
func (fs *Storage) GetWithdrawalsSubmittedLastProcessedBlock(operatorID *big.Int) (uint64, error) {
	db, err := fs.LoadDatabase()
	if err != nil {
		return 0, err
	}

	opID := operatorID.String()
	operatorData, exists := db.Operators[opID]
	if !exists {
		return 0, fmt.Errorf("operator ID %s not found", opID)
	}

	return operatorData.WithdrawalsSubmitted.LastProcessedBlock, nil
}

// SaveWithdrawalsSubmittedLastProcessedBlock updates the last processed block for the WithdrawalsSubmitted event for a specific operator ID.
func (fs *Storage) SaveWithdrawalsSubmittedLastProcessedBlock(operatorID *big.Int, block uint64) error {
	db, err := fs.LoadDatabase()
	if err != nil {
		return err
	}

	opID := operatorID.String()
	operatorData, exists := db.Operators[opID]
	if !exists {
		return fmt.Errorf("operator ID %s not found", opID)
	}

	operatorData.WithdrawalsSubmitted.LastProcessedBlock = block
	db.Operators[opID] = operatorData

	return fs.SaveDatabase(db)
}

// GetWithdrawals retrieves all WithdrawalsSubmitted events for a specific operator ID.
func (fs *Storage) GetWithdrawals(operatorID *big.Int) ([]domain.CsmoduleWithdrawalSubmitted, error) {
	db, err := fs.LoadDatabase()
	if err != nil {
		return nil, err
	}

	opID := operatorID.String()
	operatorData, exists := db.Operators[opID]
	if !exists {
		return nil, fmt.Errorf("operator ID %s not found", opID)
	}

	return operatorData.WithdrawalsSubmitted.Withdrawals, nil
}

// SaveWithdrawal saves a WithdrawalSubmitted event for a specific operator ID.
func (fs *Storage) SaveWithdrawal(operatorID *big.Int, withdrawal domain.CsmoduleWithdrawalSubmitted) error {
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
			DistributionLogsUpdated:            DistributionLogsUpdated{},
			ExitsRequests:                      ExitsRequests{},
			WithdrawalsSubmitted:               WithdrawalsSubmitted{},
			ElRewardsStealingPenaltiesReported: ElRewardsStealingPenaltiesReported{},
		}
	}

	// Append the new withdrawal to the existing list
	operatorData.WithdrawalsSubmitted.Withdrawals = append(operatorData.WithdrawalsSubmitted.Withdrawals, withdrawal)
	db.Operators[opID] = operatorData

	return fs.SaveDatabase(db)
}
