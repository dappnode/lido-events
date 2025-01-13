//go:build integration
// +build integration

package storage_test

import (
	"lido-events/internal/adapters/storage"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSaveCsModuleLastProcessedBlock tests saving the last processed block for CsModule events.
func TestSaveCsModuleLastProcessedBlock(t *testing.T) {
	tmpFile := CreateTempDatabaseFile(t, nil)
	defer os.Remove(tmpFile.Name())

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	block := uint64(150)

	// Save the last processed block
	err := storageAdapter.SaveCsModuletLastProcessedBlock(block)
	assert.NoError(t, err)

	// Verify that the block was saved
	db, err := storageAdapter.LoadDatabase()
	assert.NoError(t, err)
	assert.Equal(t, block, db.Events.CsModule.LastProcessedBlock)
}

// TestGetCsModuleLastProcessedBlock tests retrieving the last processed block for CsModule events.
func TestGetCsModuleLastProcessedBlock(t *testing.T) {
	initialData := &storage.Database{
		Events: storage.Events{
			CsModule: struct {
				LastProcessedBlock uint64 `json:"lastProcessedBlock"`
			}{
				LastProcessedBlock: 150,
			},
		},
	}

	tmpFile := CreateTempDatabaseFile(t, initialData)
	defer os.Remove(tmpFile.Name())

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}

	// Retrieve the last processed block
	block, err := storageAdapter.GetCsModuleLastProcessedBlock()
	assert.NoError(t, err)
	assert.Equal(t, initialData.Events.CsModule.LastProcessedBlock, block)
}
