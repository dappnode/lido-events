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
	assert.NotNil(t, db.Operators)
	assert.Empty(t, db.Operators)
	assert.Equal(t, uint64(0), db.Events.DistributionLogUpdated.LastProcessedBlock)
	assert.Empty(t, db.Events.DistributionLogUpdated.PendingHashes)
	assert.Equal(t, uint64(0), db.Events.ValidatorExitRequest.LastProcessedBlock)
}

// Test for LoadDatabase when the database file has existing data
func TestLoadDatabase_WithExistingData(t *testing.T) {
	// Initialize the database with existing data
	existingData := &storage.Database{
		Telegram: domain.TelegramConfig{
			Token:  "test-token",
			UserID: 12345,
		},
		Operators: map[string]storage.OperatorData{
			"1": {
				Performance: domain.Reports{
					"100": domain.Report{Threshold: "0.85"},
				},
				ExitRequests: domain.ExitRequests{
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
		Events: storage.Events{
			DistributionLogUpdated: struct {
				LastProcessedBlock uint64   `json:"lastProcessedBlock"`
				PendingHashes      []string `json:"pendingHashes"`
			}{
				LastProcessedBlock: 100,
				PendingHashes:      []string{"hash1", "hash2"},
			},
			ValidatorExitRequest: struct {
				LastProcessedBlock uint64 `json:"lastProcessedBlock"`
			}{
				LastProcessedBlock: 200,
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
	assert.Equal(t, uint64(100), db.Events.DistributionLogUpdated.LastProcessedBlock)
	assert.Equal(t, []string{"hash1", "hash2"}, db.Events.DistributionLogUpdated.PendingHashes)
	assert.Equal(t, uint64(200), db.Events.ValidatorExitRequest.LastProcessedBlock)
	assert.Contains(t, db.Operators, "1")
	assert.Contains(t, db.Operators["1"].Performance, "100")
	assert.Contains(t, db.Operators["1"].ExitRequests, "validator1")
}

// Test for LoadDatabase to check initialization of missing fields in an existing file
func TestLoadDatabase_MissingFields(t *testing.T) {
	// Create a database file with missing fields (no operators and missing events fields)
	missingFieldsData := &storage.Database{
		Telegram: domain.TelegramConfig{
			Token:  "test-token",
			UserID: 12345,
		},
		Operators: make(map[string]storage.OperatorData), // Empty Operators
		Events: storage.Events{
			DistributionLogUpdated: struct {
				LastProcessedBlock uint64   `json:"lastProcessedBlock"`
				PendingHashes      []string `json:"pendingHashes"`
			}{},
			ValidatorExitRequest: struct {
				LastProcessedBlock uint64 `json:"lastProcessedBlock"`
			}{},
		},
	}

	tmpFile := CreateTempDatabaseFile(t, missingFieldsData)
	defer os.Remove(tmpFile.Name())

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	db, err := storageAdapter.LoadDatabase()
	assert.NoError(t, err)

	// Ensure missing fields are initialized correctly
	assert.Equal(t, "test-token", db.Telegram.Token)
	assert.Equal(t, uint64(0), db.Events.DistributionLogUpdated.LastProcessedBlock)
	assert.Empty(t, db.Events.DistributionLogUpdated.PendingHashes)
	assert.Equal(t, uint64(0), db.Events.ValidatorExitRequest.LastProcessedBlock)
	assert.NotNil(t, db.Operators)
	assert.Empty(t, db.Operators)
}

func TestSaveDatabase(t *testing.T) {
	tmpFile := CreateTempDatabaseFile(t, nil)
	defer os.Remove(tmpFile.Name())

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	db := storage.Database{
		Telegram: domain.TelegramConfig{
			Token:  "new-token",
			UserID: 98765,
		},
		Operators: map[string]storage.OperatorData{
			"2": {
				Performance: domain.Reports{
					"200": domain.Report{Threshold: "0.90"},
				},
				ExitRequests: domain.ExitRequests{
					"validator2": {
						Event: domain.VeboValidatorExitRequest{
							ValidatorIndex: big.NewInt(2),
							Raw: types.Log{
								Topics:      []common.Hash{common.HexToHash("0x0")},
								Data:        []uint8{},
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
		Events: storage.Events{
			DistributionLogUpdated: struct {
				LastProcessedBlock uint64   `json:"lastProcessedBlock"`
				PendingHashes      []string `json:"pendingHashes"`
			}{
				LastProcessedBlock: 300,
				PendingHashes:      []string{"hash3", "hash4"},
			},
			ValidatorExitRequest: struct {
				LastProcessedBlock uint64 `json:"lastProcessedBlock"`
			}{
				LastProcessedBlock: 400,
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
