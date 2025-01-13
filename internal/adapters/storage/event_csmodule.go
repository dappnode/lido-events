package storage

// GetCsModuleLastProcessedBlock retrieves the last processed block for the CsModule events.
func (fs *Storage) GetCsModuleLastProcessedBlock() (uint64, error) {
	db, err := fs.LoadDatabase()
	if err != nil {
		return 0, err
	}
	return db.Events.CsModule.LastProcessedBlock, nil
}

// SaveCsModuletLastProcessedBlock updates the last processed block for the CsModule events.
func (fs *Storage) SaveCsModuletLastProcessedBlock(block uint64) error {
	db, err := fs.LoadDatabase()
	if err != nil {
		return err
	}
	db.Events.CsModule.LastProcessedBlock = block
	return fs.SaveDatabase(db)
}
