//go:build integration
// +build integration

package storage_test

import (
	"lido-events/internal/adapters/storage"
	"lido-events/internal/application/domain"
	"math/big"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSaveOperatorId_NewOperator tests adding a new operator ID that doesn't already exist.
func TestSaveOperatorId_NewOperator(t *testing.T) {
	tmpFile := CreateTempDatabaseFile(t, nil) // Helper function to create a temporary database file
	defer os.Remove(tmpFile.Name())           // Clean up

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	operatorID := big.NewInt(1)

	// Save a new operator ID
	err := storageAdapter.SaveOperatorId(operatorID.String())
	assert.NoError(t, err)

	// Verify the operator ID was saved
	db, err := storageAdapter.LoadDatabase()
	assert.NoError(t, err)
	assert.Contains(t, db.Operators, operatorID.String())
}

// TestSaveOperatorId_ExistingOperator tests attempting to save an operator ID that already exists.
func TestSaveOperatorId_ExistingOperator(t *testing.T) {
	// Initialize the database with an existing operator ID
	initialData := &storage.Database{
		Operators: map[string]storage.OperatorData{
			"1": { // operatorID as string
				Reports:      make(domain.Reports),
				ExitRequests: make(domain.ExitRequests),
				NodeOperatorEvents: domain.NodeOperatorEvents{
					NodeOperatorAdded:                 []domain.CsmoduleNodeOperatorAdded{},
					NodeOperatorManagerAddressChanged: []domain.CsmoduleNodeOperatorManagerAddressChanged{},
					NodeOperatorRewardAddressChanged:  []domain.CsmoduleNodeOperatorRewardAddressChanged{},
				},
			},
		},
	}

	tmpFile := CreateTempDatabaseFile(t, initialData)
	defer os.Remove(tmpFile.Name()) // Clean up

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	operatorID := big.NewInt(1)

	// Attempt to save the same operator ID again
	err := storageAdapter.SaveOperatorId(operatorID.String())
	assert.NoError(t, err) // No error expected, as it should recognize the ID already exists

	// Verify that no duplicate entries were created
	db, err := storageAdapter.LoadDatabase()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(db.Operators))
}

// TestGetOperatorIds tests retrieving a list of all operator IDs.
func TestGetOperatorIds(t *testing.T) {
	// Initialize database with multiple operator IDs
	initialData := &storage.Database{
		Operators: map[string]storage.OperatorData{
			"1": {},
			"2": {},
			"3": {},
		},
	}

	tmpFile := CreateTempDatabaseFile(t, initialData)
	defer os.Remove(tmpFile.Name()) // Clean up

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}

	// Retrieve operator IDs
	operatorIDs, err := storageAdapter.GetOperatorIds()
	assert.NoError(t, err)

	// Verify the retrieved IDs match the expected IDs
	expectedIDs := []*big.Int{big.NewInt(1), big.NewInt(2), big.NewInt(3)}
	assert.Equal(t, len(expectedIDs), len(operatorIDs))
	for _, expectedID := range expectedIDs {
		assert.Contains(t, operatorIDs, expectedID)
	}
}

// TestGetOperatorIds_EmptyDatabase tests retrieving operator IDs when no operators exist.
func TestGetOperatorIds_EmptyDatabase(t *testing.T) {
	tmpFile := CreateTempDatabaseFile(t, nil) // Empty initial data
	defer os.Remove(tmpFile.Name())           // Clean up

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}

	// Attempt to retrieve operator IDs from an empty database
	operatorIDs, err := storageAdapter.GetOperatorIds()
	assert.NoError(t, err)
	assert.Empty(t, operatorIDs) // Expecting an empty slice as there are no operators
}

// TestGetOperatorIds_InvalidOperatorId tests retrieving operator IDs with an invalid operator ID string.
func TestGetOperatorIds_InvalidOperatorId(t *testing.T) {
	// Initialize database with an invalid operator ID string (non-numeric)
	initialData := &storage.Database{
		Operators: map[string]storage.OperatorData{
			"invalid": {}, // Non-numeric ID
		},
	}

	tmpFile := CreateTempDatabaseFile(t, initialData)
	defer os.Remove(tmpFile.Name()) // Clean up

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}

	// Attempt to retrieve operator IDs, expecting an error
	operatorIDs, err := storageAdapter.GetOperatorIds()
	assert.Error(t, err)
	assert.Nil(t, operatorIDs)
	assert.EqualError(t, err, "failed to convert operator ID invalid to big.Int")
}

// TestDeleteOperator_ExistingOperator tests deleting an existing operator ID.
func TestDeleteOperator_ExistingOperator(t *testing.T) {
	// Initialize the database with an operator ID
	initialData := &storage.Database{
		Operators: map[string]storage.OperatorData{
			"1": { // operatorID as string
				Reports:      make(domain.Reports),
				ExitRequests: make(domain.ExitRequests),
			},
		},
	}

	tmpFile := CreateTempDatabaseFile(t, initialData)
	defer os.Remove(tmpFile.Name()) // Clean up

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	operatorID := "1"

	// Delete the operator ID
	err := storageAdapter.DeleteOperator(operatorID)
	assert.NoError(t, err)

	// Verify the operator ID was deleted
	db, err := storageAdapter.LoadDatabase()
	assert.NoError(t, err)
	assert.NotContains(t, db.Operators, operatorID)
}

// TestDeleteOperator_NonExistentOperator tests attempting to delete a non-existent operator ID.
func TestDeleteOperator_NonExistentOperator(t *testing.T) {
	tmpFile := CreateTempDatabaseFile(t, nil) // Empty initial data
	defer os.Remove(tmpFile.Name())           // Clean up

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	operatorID := "non-existent"

	// Attempt to delete a non-existent operator ID
	err := storageAdapter.DeleteOperator(operatorID)
	assert.Error(t, err)
	assert.EqualError(t, err, "operator ID non-existent not found")
}

// TestDeleteOperator_EmptyDatabase tests deleting an operator ID from an empty database.
func TestDeleteOperator_EmptyDatabase(t *testing.T) {
	tmpFile := CreateTempDatabaseFile(t, nil) // Empty initial data
	defer os.Remove(tmpFile.Name())           // Clean up

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	operatorID := "1"

	// Attempt to delete an operator ID from an empty database
	err := storageAdapter.DeleteOperator(operatorID)
	assert.Error(t, err)
	assert.EqualError(t, err, "operator ID 1 not found")
}
