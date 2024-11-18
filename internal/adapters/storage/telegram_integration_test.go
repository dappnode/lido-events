//go:build integration
// +build integration

package storage_test

import (
	"lido-events/internal/adapters/storage"
	"lido-events/internal/application/domain"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test for saving a new Telegram configuration
func TestSaveTelegramConfig(t *testing.T) {
	tmpFile := CreateTempDatabaseFile(t, nil) // Helper function to create a temporary database file
	defer os.Remove(tmpFile.Name())           // Clean up

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	telegramConfig := domain.TelegramConfig{
		Token:  "test-token",
		UserID: 123456789,
	}

	// Save the new Telegram configuration
	err := storageAdapter.SaveTelegramConfig(telegramConfig)
	assert.NoError(t, err)

	// Verify that the configuration was saved correctly
	db, err := storageAdapter.LoadDatabase()
	assert.NoError(t, err)
	assert.Equal(t, telegramConfig, db.Telegram)
}

// Test for retrieving the Telegram configuration after it has been saved
func TestGetTelegramConfig(t *testing.T) {
	// Initialize the database with an existing Telegram configuration
	initialData := &storage.Database{
		Telegram: domain.TelegramConfig{
			Token:  "initial-token",
			UserID: 987654321,
		},
	}

	tmpFile := CreateTempDatabaseFile(t, initialData)
	defer os.Remove(tmpFile.Name()) // Clean up

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}

	// Retrieve the Telegram configuration and verify it matches the initial data
	config, err := storageAdapter.GetTelegramConfig()
	assert.NoError(t, err)
	assert.Equal(t, initialData.Telegram, config)
}

// Test for getting a default Telegram configuration if none exists
func TestGetTelegramConfig_Default(t *testing.T) {
	tmpFile := CreateTempDatabaseFile(t, nil) // No initial data, so db file is effectively empty
	defer os.Remove(tmpFile.Name())           // Clean up

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}

	// Retrieve the Telegram configuration, which should be the default (empty values)
	config, err := storageAdapter.GetTelegramConfig()
	assert.NoError(t, err)
	assert.Equal(t, domain.TelegramConfig{}, config)
}

// Test for saving a new Telegram configuration and then retrieving it
func TestSaveAndGetTelegramConfig(t *testing.T) {
	tmpFile := CreateTempDatabaseFile(t, nil)
	defer os.Remove(tmpFile.Name()) // Clean up

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	telegramConfig := domain.TelegramConfig{
		Token:  "final-token",
		UserID: 111222333,
	}

	// Save the configuration
	err := storageAdapter.SaveTelegramConfig(telegramConfig)
	assert.NoError(t, err)

	// Retrieve and verify it matches what was saved
	retrievedConfig, err := storageAdapter.GetTelegramConfig()
	assert.NoError(t, err)
	assert.Equal(t, telegramConfig, retrievedConfig)
}
