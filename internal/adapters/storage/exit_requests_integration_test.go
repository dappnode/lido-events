//go:build integration
// +build integration

package storage_test

import (
	"fmt"
	"lido-events/internal/adapters/storage"
	"lido-events/internal/application/domain"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
)

// TestSaveExitRequests_NewOperator tests saving multiple exit requests for a new operator.
func TestSaveExitRequests_NewOperator(t *testing.T) {
	tmpFile := CreateTempDatabaseFile(t, nil)
	defer os.Remove(tmpFile.Name())

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	operatorID := big.NewInt(1)
	exitRequests := map[string]domain.ExitRequest{
		"validator1": {
			Status: domain.StatusActiveOngoing,
			Event: domain.VeboValidatorExitRequest{
				Raw: types.Log{
					Address: common.HexToAddress("0x0000000000000000000000000000000000000000"),
					Topics:  []common.Hash{}, // Explicitly initialize topics
					Data:    []byte{},
				},
			},
		},
		"validator2": {
			Status: domain.StatusExitedUnslashed,
			Event: domain.VeboValidatorExitRequest{
				Raw: types.Log{
					Address: common.HexToAddress("0x0000000000000000000000000000000000000000"),
					Topics:  []common.Hash{}, // Explicitly initialize topics
					Data:    []byte{},
				},
			},
		},
	}

	// Save multiple exit requests
	err := storageAdapter.SaveExitRequests(operatorID, exitRequests)
	assert.NoError(t, err)

	// Verify that exit requests were saved
	db, err := storageAdapter.LoadDatabase()
	assert.NoError(t, err)
	assert.Equal(t, exitRequests, db.Operators[operatorID.String()].ExitRequests)
}

// TestSaveExitRequest tests saving an individual exit request for a specific operator ID and pubkey.
func TestSaveExitRequest(t *testing.T) {
	tmpFile := CreateTempDatabaseFile(t, nil)
	defer os.Remove(tmpFile.Name())

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	operatorID := big.NewInt(2)
	pubkey := "validator1"
	exitRequest := domain.ExitRequest{
		Status: domain.StatusActiveOngoing,
		Event: domain.VeboValidatorExitRequest{
			Raw: types.Log{
				Address: common.HexToAddress("0x0000000000000000000000000000000000000000"),
				Topics:  []common.Hash{}, // Explicitly initialize topics
				Data:    []byte{},
			},
		},
	}

	// Save individual exit request
	err := storageAdapter.SaveExitRequest(operatorID, pubkey, exitRequest)
	assert.NoError(t, err)

	// Verify that the exit request was saved
	db, err := storageAdapter.LoadDatabase()
	assert.NoError(t, err)
	assert.Equal(t, exitRequest, db.Operators[operatorID.String()].ExitRequests[pubkey])
}

// TestGetExitRequests tests retrieving all exit requests for a specific operator ID.
func TestGetExitRequests(t *testing.T) {
	// Initialize the database with exit requests for an operator
	initialData := &storage.Database{
		Operators: map[string]storage.OperatorData{
			"1": {
				ExitRequests: map[string]domain.ExitRequest{
					"validator1": {
						Status: domain.StatusActiveOngoing,
						Event: domain.VeboValidatorExitRequest{
							Raw: types.Log{
								Address: common.HexToAddress("0x0000000000000000000000000000000000000000"),
								Topics:  []common.Hash{}, // Explicitly initialize topics
								Data:    []byte{},
							},
						},
					},
					"validator2": {
						Status: domain.StatusExitedUnslashed,
						Event: domain.VeboValidatorExitRequest{
							Raw: types.Log{
								Address: common.HexToAddress("0x0000000000000000000000000000000000000000"),
								Topics:  []common.Hash{}, // Explicitly initialize topics
								Data:    []byte{},
							},
						},
					},
				},
			},
		},
	}

	tmpFile := CreateTempDatabaseFile(t, initialData)
	defer os.Remove(tmpFile.Name())

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	operatorID := "1"

	// Retrieve exit requests
	exitRequests, err := storageAdapter.GetExitRequests(operatorID)
	assert.NoError(t, err)

	// Verify the retrieved data matches the initial data
	expectedExitRequests := initialData.Operators[operatorID].ExitRequests
	assert.Equal(t, expectedExitRequests, exitRequests)
}

// TestUpdateExitRequestStatus tests updating the status of an existing exit request.
func TestUpdateExitRequestStatus(t *testing.T) {
	// Initialize the database with an existing exit request
	initialData := &storage.Database{
		Operators: map[string]storage.OperatorData{
			"1": {
				ExitRequests: map[string]domain.ExitRequest{
					"validator1": {
						Status: domain.StatusActiveOngoing,
						Event: domain.VeboValidatorExitRequest{
							Raw: types.Log{
								Address: common.HexToAddress("0x0000000000000000000000000000000000000000"),
								Topics:  []common.Hash{}, // Explicitly initialize topics
								Data:    []byte{},
							},
						},
					},
				},
			},
		},
	}

	tmpFile := CreateTempDatabaseFile(t, initialData)
	defer os.Remove(tmpFile.Name())

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	operatorID := "1"
	pubkey := "validator1"
	newStatus := domain.StatusExitedUnslashed

	// Update the exit request status
	err := storageAdapter.UpdateExitRequestStatus(operatorID, pubkey, newStatus)
	assert.NoError(t, err)

	// Verify that the status was updated
	db, err := storageAdapter.LoadDatabase()
	assert.NoError(t, err)
	assert.Equal(t, newStatus, db.Operators[operatorID].ExitRequests[pubkey].Status)
}

// TestUpdateExitRequestStatus_NonExistent tests updating the status of a non-existent exit request.
func TestUpdateExitRequestStatus_NonExistent(t *testing.T) {
	tmpFile := CreateTempDatabaseFile(t, nil)
	defer os.Remove(tmpFile.Name())

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	operatorID := "3" // Operator ID that doesn't exist
	pubkey := "nonexistent"

	// Attempt to update the status of a non-existent exit request
	err := storageAdapter.UpdateExitRequestStatus(operatorID, pubkey, domain.StatusExitedUnslashed)
	assert.Error(t, err)

	// Adjust the error message expectation based on the missing operator ID
	assert.EqualError(t, err, fmt.Sprintf("operator ID %s not found", operatorID))
}
