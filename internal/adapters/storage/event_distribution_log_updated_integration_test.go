//go:build integration
// +build integration

package storage_test

import (
	"lido-events/internal/adapters/storage"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSaveDistributionLogLastProcessedBlock tests saving the last processed block for DistributionLogUpdated.
func TestSaveDistributionLogLastProcessedBlock(t *testing.T) {
	tmpFile := CreateTempDatabaseFile(t, nil)
	defer os.Remove(tmpFile.Name())

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	block := uint64(100)

	// Save the last processed block
	err := storageAdapter.SaveDistributionLogLastProcessedBlock(block)
	assert.NoError(t, err)

	// Verify that the block was saved
	db, err := storageAdapter.LoadDatabase()
	assert.NoError(t, err)
	assert.Equal(t, block, db.Events.DistributionLogUpdated.LastProcessedBlock)
}

// TestGetDistributionLogLastProcessedBlock tests retrieving the last processed block for DistributionLogUpdated.
func TestGetDistributionLogLastProcessedBlock(t *testing.T) {
	initialData := &storage.Database{
		Events: storage.Events{
			DistributionLogUpdated: struct {
				LastProcessedBlock uint64   `json:"lastProcessedBlock"`
				PendingHashes      []string `json:"pendingHashes"`
			}{
				LastProcessedBlock: 100,
			},
		},
	}

	tmpFile := CreateTempDatabaseFile(t, initialData)
	defer os.Remove(tmpFile.Name())

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}

	// Retrieve the last processed block
	block, err := storageAdapter.GetDistributionLogLastProcessedBlock()
	assert.NoError(t, err)
	assert.Equal(t, initialData.Events.DistributionLogUpdated.LastProcessedBlock, block)
}

// TestAddPendingHash tests adding a pending hash to the DistributionLogUpdated struct.
func TestAddPendingHash(t *testing.T) {
	tmpFile := CreateTempDatabaseFile(t, nil)
	defer os.Remove(tmpFile.Name())

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	hash := "hash1"

	// Add a pending hash
	err := storageAdapter.AddPendingHash(hash)
	assert.NoError(t, err)

	// Verify the hash was added
	db, err := storageAdapter.LoadDatabase()
	assert.NoError(t, err)
	assert.Contains(t, db.Events.DistributionLogUpdated.PendingHashes, hash)
}

// TestDeletePendingHash tests removing a pending hash from the DistributionLogUpdated struct.
func TestDeletePendingHash(t *testing.T) {
	initialData := &storage.Database{
		Events: storage.Events{
			DistributionLogUpdated: struct {
				LastProcessedBlock uint64   `json:"lastProcessedBlock"`
				PendingHashes      []string `json:"pendingHashes"`
			}{
				PendingHashes: []string{"hash1"},
			},
		},
	}

	tmpFile := CreateTempDatabaseFile(t, initialData)
	defer os.Remove(tmpFile.Name())

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	hash := "hash1"

	// Delete the pending hash
	err := storageAdapter.DeletePendingHash(hash)
	assert.NoError(t, err)

	// Verify the hash was removed
	db, err := storageAdapter.LoadDatabase()
	assert.NoError(t, err)
	assert.NotContains(t, db.Events.DistributionLogUpdated.PendingHashes, hash)
}

// TestGetPendingHashes tests retrieving all pending hashes from DistributionLogUpdated.
func TestGetPendingHashes(t *testing.T) {
	initialData := &storage.Database{
		Events: storage.Events{
			DistributionLogUpdated: struct {
				LastProcessedBlock uint64   `json:"lastProcessedBlock"`
				PendingHashes      []string `json:"pendingHashes"`
			}{
				PendingHashes: []string{"hash1", "hash2"},
			},
		},
	}

	tmpFile := CreateTempDatabaseFile(t, initialData)
	defer os.Remove(tmpFile.Name())

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}

	// Retrieve pending hashes
	hashes, err := storageAdapter.GetPendingHashes()
	assert.NoError(t, err)
	assert.Equal(t, initialData.Events.DistributionLogUpdated.PendingHashes, hashes)
}
