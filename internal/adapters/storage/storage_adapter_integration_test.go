//go:build integration
// +build integration

package storage_test

import (
	"lido-events/internal/adapters/storage"
	"lido-events/internal/application/domain"
	"math/big"
	"os"
	"path/filepath"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
)

// Test for NewStorageAdapter to ensure it creates the directory if it does not exist
func TestNewStorageAdapter_CreatesDirectory(t *testing.T) {
	// Create a temporary directory and remove it to simulate a missing directory
	tempDir := t.TempDir()
	os.Remove(tempDir) // Ensure the directory does not exist

	// Initialize the storage adapter
	adapter, err := storage.NewStorageAdapter(tempDir)
	assert.NoError(t, err)

	// Check that the directory was created
	_, err = os.Stat(tempDir)
	assert.NoError(t, err)
	assert.True(t, os.IsExist(err) || err == nil, "expected directory to be created")

	// Verify the adapter was initialized with the correct DB file path
	expectedDBFile := filepath.Join(tempDir, "db.json")
	assert.Equal(t, expectedDBFile, adapter.DBFile)
}

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
	assert.Equal(t, uint64(0), db.Events.CsModule.LastProcessedBlock)
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
				Reports: domain.Reports{
					"81055-81504": domain.Report{
						Frame:     [2]int{81055, 81504},
						Threshold: 0.85,
						Data: domain.Data{
							Distributed: int64(3465578901933468),
							Stuck:       false,
							Validators: map[string]domain.Validator{
								"1735661": {
									Perf: domain.Performance{
										Assigned: 450,
										Included: 450,
									},
									Slashed: false,
								},
							},
						},
					},
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
				NodeOperatorEvents: domain.NodeOperatorEvents{
					NodeOperatorAdded:                 []domain.CsmoduleNodeOperatorAdded{},
					NodeOperatorManagerAddressChanged: []domain.CsmoduleNodeOperatorManagerAddressChanged{},
					NodeOperatorRewardAddressChanged:  []domain.CsmoduleNodeOperatorRewardAddressChanged{},
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
			CsModule: struct {
				LastProcessedBlock uint64 `json:"lastProcessedBlock"`
			}{
				LastProcessedBlock: 300,
			},
		},
	}
	tmpFile := CreateTempDatabaseFile(t, existingData)
	defer os.Remove(tmpFile.Name())

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	db, err := storageAdapter.LoadDatabase()
	assert.NoError(t, err)

	assert.Equal(t, "test-token", db.Telegram.Token)
	assert.Equal(t, uint64(100), db.Events.DistributionLogUpdated.LastProcessedBlock)
	assert.Equal(t, []string{"hash1", "hash2"}, db.Events.DistributionLogUpdated.PendingHashes)
	assert.Equal(t, uint64(200), db.Events.ValidatorExitRequest.LastProcessedBlock)
	assert.Equal(t, uint64(300), db.Events.CsModule.LastProcessedBlock)
	assert.Contains(t, db.Operators, "1")
	assert.Contains(t, db.Operators["1"].Reports, "81055-81504")
	assert.Equal(t, 0.85, db.Operators["1"].Reports["81055-81504"].Threshold)
	assert.Equal(t, int64(3465578901933468), db.Operators["1"].Reports["81055-81504"].Data.Distributed)
	assert.Equal(t, 450, db.Operators["1"].Reports["81055-81504"].Data.Validators["1735661"].Perf.Assigned)
	assert.False(t, db.Operators["1"].Reports["81055-81504"].Data.Validators["1735661"].Slashed)
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
	assert.Equal(t, uint64(0), db.Events.CsModule.LastProcessedBlock)
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
				Reports: domain.Reports{
					"81055-81504": domain.Report{
						Frame:     [2]int{81055, 81504},
						Threshold: 0.90,
						Data: domain.Data{
							Distributed: 500000000,
							Stuck:       false,
							Validators: map[string]domain.Validator{
								"1797927": {
									Perf: domain.Performance{
										Assigned: 450,
										Included: 0,
									},
									Slashed: false,
								},
							},
						},
					},
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
				NodeOperatorEvents: domain.NodeOperatorEvents{
					NodeOperatorAdded:                 []domain.CsmoduleNodeOperatorAdded{},
					NodeOperatorManagerAddressChanged: []domain.CsmoduleNodeOperatorManagerAddressChanged{},
					NodeOperatorRewardAddressChanged:  []domain.CsmoduleNodeOperatorRewardAddressChanged{},
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
			CsModule: struct {
				LastProcessedBlock uint64 `json:"lastProcessedBlock"`
			}{
				LastProcessedBlock: 500,
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
