//go:build integration
// +build integration

package storage_test

import (
	"lido-events/internal/adapters/storage"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test migration when database exists but has no version field
func TestMigrations_NoVersion(t *testing.T) {
	// Create a simulated old database without the "version" field
	initialData := &storage.Database{
		// Note: **Version is intentionally omitted to simulate an old database**
		Operators: map[string]storage.OperatorData{
			"123": {},
			"456": {},
		},
	}

	// Create a temporary database file with the old structure (no version)
	tmpFile := CreateTempDatabaseFile(t, initialData)
	defer os.Remove(tmpFile.Name()) // Cleanup after test

	// Initialize the storage adapter with the temp file
	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}

	// Run migrations
	err := storage.RunMigrations(storageAdapter)
	assert.NoError(t, err)

	// Reload the database after migration
	db, err := storageAdapter.LoadDatabase()
	assert.NoError(t, err)

	// Check that the migration **added the missing version**
	assert.Equal(t, 1, db.Version, "Database version should be updated to 1")

	// Ensure operator IDs were retained but reset to empty OperatorData
	expectedOperators := map[string]storage.OperatorData{
		"123": {},
		"456": {},
	}
	assert.Equal(t, expectedOperators, db.Operators, "Operators should exist but have empty OperatorData after migration")
}
