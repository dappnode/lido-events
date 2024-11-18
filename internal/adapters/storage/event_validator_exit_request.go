// validator_exit_request.go
package storage

// GetValidatorExitRequestLastProcessedBlock retrieves the last processed block for the ValidatorExitRequest event.
func (fs *Storage) GetValidatorExitRequestLastProcessedBlock() (uint64, error) {
	db, err := fs.LoadDatabase()
	if err != nil {
		return 0, err
	}
	return db.Events.ValidatorExitRequest.LastProcessedBlock, nil
}

// SaveValidatorExitRequestLastProcessedBlock updates the last processed block for the ValidatorExitRequest event.
func (fs *Storage) SaveValidatorExitRequestLastProcessedBlock(block uint64) error {
	db, err := fs.LoadDatabase()
	if err != nil {
		return err
	}
	db.Events.ValidatorExitRequest.LastProcessedBlock = block
	return fs.SaveDatabase(db)
}
