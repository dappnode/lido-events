// distribution_log_updated.go
package storage

// GetDistributionLogLastProcessedBlock retrieves the last processed block for the DistributionLogUpdated event.
func (fs *Storage) GetDistributionLogLastProcessedBlock() (uint64, error) {
	db, err := fs.LoadDatabase()
	if err != nil {
		return 0, err
	}
	return db.Events.DistributionLogUpdated.LastProcessedBlock, nil
}

// SaveDistributionLogLastProcessedBlock updates the last processed block for the DistributionLogUpdated event.
func (fs *Storage) SaveDistributionLogLastProcessedBlock(block uint64) error {
	db, err := fs.LoadDatabase()
	if err != nil {
		return err
	}
	db.Events.DistributionLogUpdated.LastProcessedBlock = block
	return fs.SaveDatabase(db)
}

// AddPendingHash adds a hash to the PendingHashes array in DistributionLogUpdated if it does not already exist.
func (fs *Storage) AddPendingHash(hash string) error {
	db, err := fs.LoadDatabase()
	if err != nil {
		return err
	}

	// Check if hash already exists
	for _, h := range db.Events.DistributionLogUpdated.PendingHashes {
		if h == hash {
			return nil // Hash already exists, no need to add
		}
	}

	// Add hash to PendingHashes
	db.Events.DistributionLogUpdated.PendingHashes = append(db.Events.DistributionLogUpdated.PendingHashes, hash)
	return fs.SaveDatabase(db)
}

// DeletePendingHash removes a hash from the PendingHashes array in DistributionLogUpdated.
func (fs *Storage) DeletePendingHash(hash string) error {
	db, err := fs.LoadDatabase()
	if err != nil {
		return err
	}

	// Find and remove the hash
	for i, h := range db.Events.DistributionLogUpdated.PendingHashes {
		if h == hash {
			db.Events.DistributionLogUpdated.PendingHashes = append(db.Events.DistributionLogUpdated.PendingHashes[:i], db.Events.DistributionLogUpdated.PendingHashes[i+1:]...)
			return fs.SaveDatabase(db)
		}
	}

	return nil // Hash not found; nothing to delete
}

// GetPendingHashes retrieves all pending hashes from the DistributionLogUpdated event.
func (fs *Storage) GetPendingHashes() ([]string, error) {
	db, err := fs.LoadDatabase()
	if err != nil {
		return nil, err
	}
	return db.Events.DistributionLogUpdated.PendingHashes, nil
}
