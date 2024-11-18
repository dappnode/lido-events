//go:build integration
// +build integration

package storage_test

import (
	"lido-events/internal/adapters/storage"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSaveValidatorExitRequestLastProcessedBlock tests saving the last processed block for ValidatorExitRequest.
func TestSaveValidatorExitRequestLastProcessedBlock(t *testing.T) {
	tmpFile := CreateTempDatabaseFile(t, nil)
	defer os.Remove(tmpFile.Name())

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	block := uint64(200)

	// Save the last processed block
	err := storageAdapter.SaveValidatorExitRequestLastProcessedBlock(block)
	assert.NoError(t, err)

	// Verify that the block was saved
	db, err := storageAdapter.LoadDatabase()
	assert.NoError(t, err)
	assert.Equal(t, block, db.Events.ValidatorExitRequest.LastProcessedBlock)
}

// TestGetValidatorExitRequestLastProcessedBlock tests retrieving the last processed block for ValidatorExitRequest.
func TestGetValidatorExitRequestLastProcessedBlock(t *testing.T) {
	initialData := &storage.Database{
		Events: storage.Events{
			ValidatorExitRequest: struct {
				LastProcessedBlock uint64 `json:"lastProcessedBlock"`
			}{
				LastProcessedBlock: 200,
			},
		},
	}

	tmpFile := CreateTempDatabaseFile(t, initialData)
	defer os.Remove(tmpFile.Name())

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}

	// Retrieve the last processed block
	block, err := storageAdapter.GetValidatorExitRequestLastProcessedBlock()
	assert.NoError(t, err)
	assert.Equal(t, initialData.Events.ValidatorExitRequest.LastProcessedBlock, block)
}
