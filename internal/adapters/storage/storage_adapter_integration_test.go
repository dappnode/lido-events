//go:build integration
// +build integration

package storage_test

import (
	"lido-events/internal/adapters/storage"
	"lido-events/internal/application/domain"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
)

// Test for LoadDatabase when the database file is missing
func TestLoadDatabase_MissingFile(t *testing.T) {
	tmpFile := CreateTempDatabaseFile(t, nil)
	defer os.Remove(tmpFile.Name())

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}

	// Load database should initialize a new structure if the file is missing
	db, err := storageAdapter.LoadDatabase()
	assert.NoError(t, err)

	// Check if the database has been initialized with the correct defaults
	assert.Equal(t, domain.TelegramConfig{}, db.Telegram)
	assert.Equal(t, uint64(0), db.Operators.LastProcessedEpoch)
	assert.NotNil(t, db.Operators.OperatorDetails)
	assert.Empty(t, db.Operators.OperatorDetails)
}

// Test for LoadDatabase when the database file has existing data
func TestLoadDatabase_WithExistingData(t *testing.T) {
	// Initialize the database with existing data, including necessary fields in types.Log
	existingData := &storage.Database{
		Telegram: domain.TelegramConfig{
			Token:  "test-token",
			ChatID: 12345,
		},
		Operators: storage.OperatorsData{
			LastProcessedEpoch: 100,
			OperatorDetails: map[string]storage.OperatorDetails{
				"1": {
					Performance: map[string]domain.Report{
						"100": {
							Threshold: "0.85",
						},
					},
					ExitRequests: map[string]domain.ExitRequest{
						"validator1": {
							Event: domain.VeboValidatorExitRequest{
								ValidatorIndex: big.NewInt(1),
								Raw: types.Log{
									Topics: []common.Hash{common.HexToHash("0x0")},
								},
							},
							Status: domain.StatusActiveOngoing,
						},
					},
				},
			},
		},
	}

	tmpFile := CreateTempDatabaseFile(t, existingData)
	defer os.Remove(tmpFile.Name())

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	db, err := storageAdapter.LoadDatabase()
	assert.NoError(t, err)

	// Validate loaded data matches the existing data
	assert.Equal(t, "test-token", db.Telegram.Token)
	assert.Equal(t, uint64(100), db.Operators.LastProcessedEpoch)
	assert.Contains(t, db.Operators.OperatorDetails, "1")
	assert.Contains(t, db.Operators.OperatorDetails["1"].Performance, "100")
	assert.Contains(t, db.Operators.OperatorDetails["1"].ExitRequests, "validator1")
}

// Test for LoadDatabase to check initialization of missing fields in an existing file
func TestLoadDatabase_MissingFields(t *testing.T) {
	// Create a database file with missing fields (no operators)
	missingFieldsData := &storage.Database{
		Telegram: domain.TelegramConfig{
			Token:  "test-token",
			ChatID: 12345,
		},
		Operators: storage.OperatorsData{}, // Empty OperatorsData
	}

	tmpFile := CreateTempDatabaseFile(t, missingFieldsData)
	defer os.Remove(tmpFile.Name())

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	db, err := storageAdapter.LoadDatabase()
	assert.NoError(t, err)

	// Ensure missing fields are initialized correctly
	assert.Equal(t, "test-token", db.Telegram.Token)
	assert.Equal(t, uint64(0), db.Operators.LastProcessedEpoch)
	assert.NotNil(t, db.Operators.OperatorDetails)
	assert.Empty(t, db.Operators.OperatorDetails)
}

func TestSaveDatabase(t *testing.T) {
	tmpFile := CreateTempDatabaseFile(t, nil)
	defer os.Remove(tmpFile.Name())

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	db := storage.Database{
		Telegram: domain.TelegramConfig{
			Token:  "new-token",
			ChatID: 98765,
		},
		Operators: storage.OperatorsData{
			LastProcessedEpoch: 200,
			OperatorDetails: map[string]storage.OperatorDetails{
				"2": {
					Performance: map[string]domain.Report{
						"200": {
							Threshold: "0.90",
						},
					},
					ExitRequests: map[string]domain.ExitRequest{
						"validator2": {
							Event: domain.VeboValidatorExitRequest{
								ValidatorIndex: big.NewInt(2),
								Raw: types.Log{
									Topics:      []common.Hash{common.HexToHash("0x0")},
									Data:        []uint8{}, // Updated to expect an empty slice instead of nil
									BlockNumber: 0,
									TxHash:      common.HexToHash("0x0"),
									TxIndex:     0,
									BlockHash:   common.HexToHash("0x0"),
									Index:       0,
									Removed:     false,
								},
							},
							Status: domain.StatusExitedUnslashed,
						},
					},
				},
			},
		},
	}

	// Save the database
	err := storageAdapter.SaveDatabase(db)
	assert.NoError(t, err)

	// Load the saved file to verify data was stored correctly
	loadedDB, err := storageAdapter.LoadDatabase()
	assert.NoError(t, err)
	assert.Equal(t, db, loadedDB)
}
