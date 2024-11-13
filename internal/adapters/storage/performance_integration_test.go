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

// TestSaveOperatorPerformance_NewOperator tests saving performance data for a new operator ID.
func TestSaveOperatorPerformance_NewOperator(t *testing.T) {
	tmpFile := CreateTempDatabaseFile(t, nil) // Helper function to create a temporary database file
	defer os.Remove(tmpFile.Name())           // Clean up

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	operatorID := big.NewInt(1)
	epoch := "100"
	report := domain.Report{
		Threshold: "0.85",
		Validators: map[string]domain.ValidatorPerformance{
			"validator1": {Included: "120", Assigned: "123"},
		},
	}

	// Save the performance data for the new operator
	err := storageAdapter.SaveOperatorPerformance(operatorID, epoch, report)
	assert.NoError(t, err)

	// Load database and verify the data was saved
	db, err := storageAdapter.LoadDatabase()
	assert.NoError(t, err)
	savedReport, exists := db.Operators.OperatorDetails[operatorID.String()].Performance[epoch]
	assert.True(t, exists)
	assert.Equal(t, report, savedReport)
}

// TestSaveOperatorPerformance_ExistingOperator tests updating performance data for an existing operator ID.
func TestSaveOperatorPerformance_ExistingOperator(t *testing.T) {
	// Initialize database with existing performance data for an operator
	initialData := &storage.Database{
		Operators: storage.OperatorsData{
			OperatorDetails: map[string]storage.OperatorDetails{
				"1": {
					Performance: map[string]domain.Report{
						"100": {
							Threshold: "0.80",
						},
					},
				},
			},
		},
	}

	tmpFile := CreateTempDatabaseFile(t, initialData)
	defer os.Remove(tmpFile.Name()) // Clean up

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	operatorID := big.NewInt(1)
	epoch := "100"
	updatedReport := domain.Report{
		Threshold: "0.85",
		Validators: map[string]domain.ValidatorPerformance{
			"validator1": {Included: "125", Assigned: "130"},
		},
	}

	// Update the performance data for the existing operator and epoch
	err := storageAdapter.SaveOperatorPerformance(operatorID, epoch, updatedReport)
	assert.NoError(t, err)

	// Verify that the performance data was updated
	db, err := storageAdapter.LoadDatabase()
	assert.NoError(t, err)
	assert.Equal(t, updatedReport, db.Operators.OperatorDetails[operatorID.String()].Performance[epoch])
}

// TestGetOperatorPerformance_Range tests retrieving performance data within a specified epoch range.
func TestGetOperatorPerformance_Range(t *testing.T) {
	// Initialize database with multiple epochs for an operator
	initialData := &storage.Database{
		Operators: storage.OperatorsData{
			OperatorDetails: map[string]storage.OperatorDetails{
				"2": {
					Performance: map[string]domain.Report{
						"100": {Threshold: "0.80"},
						"101": {Threshold: "0.85"},
						"102": {Threshold: "0.90"},
					},
				},
			},
		},
	}

	tmpFile := CreateTempDatabaseFile(t, initialData)
	defer os.Remove(tmpFile.Name()) // Clean up

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	operatorID := big.NewInt(2)
	startEpoch := "100"
	endEpoch := "101"

	// Retrieve performance data within the specified range
	reports, err := storageAdapter.GetOperatorPerformance(operatorID, startEpoch, endEpoch)
	assert.NoError(t, err)

	// Verify that only the reports within the specified range are returned
	expectedReports := map[string]domain.Report{
		"100": {Threshold: "0.80"},
		"101": {Threshold: "0.85"},
	}
	assert.Equal(t, expectedReports, reports)
}

// TestAddPendingHash tests adding a pending hash to the global pendingHashes array.
func TestAddPendingHash(t *testing.T) {
	tmpFile := CreateTempDatabaseFile(t, nil) // Empty initial data
	defer os.Remove(tmpFile.Name())           // Clean up

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	hash := "newPendingHash"

	// Add the pending hash
	err := storageAdapter.AddPendingHash(hash)
	assert.NoError(t, err)

	// Load database and verify the pending hash was added
	db, err := storageAdapter.LoadDatabase()
	assert.NoError(t, err)
	assert.Contains(t, db.Operators.PendingHashes, hash)
}

// TestDeletePendingHash tests removing a pending hash from the global pendingHashes array.
func TestDeletePendingHash(t *testing.T) {
	// Initialize database with a pending hash
	initialData := &storage.Database{
		Operators: storage.OperatorsData{
			PendingHashes: []string{"hashToDelete"},
		},
	}

	tmpFile := CreateTempDatabaseFile(t, initialData)
	defer os.Remove(tmpFile.Name()) // Clean up

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	hash := "hashToDelete"

	// Delete the pending hash
	err := storageAdapter.DeletePendingHash(hash)
	assert.NoError(t, err)

	// Load database and verify the pending hash was removed
	db, err := storageAdapter.LoadDatabase()
	assert.NoError(t, err)
	assert.NotContains(t, db.Operators.PendingHashes, hash)
}

// TestGetOperatorPerformance_NoData tests retrieving performance data when no data exists for the specified operator.
func TestGetOperatorPerformance_NoData(t *testing.T) {
	tmpFile := CreateTempDatabaseFile(t, nil) // Empty initial data
	defer os.Remove(tmpFile.Name())           // Clean up

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	operatorID := big.NewInt(3)
	startEpoch := "100"
	endEpoch := "200"

	// Attempt to retrieve performance data for a non-existent operator
	reports, err := storageAdapter.GetOperatorPerformance(operatorID, startEpoch, endEpoch)
	assert.Error(t, err)
	assert.Nil(t, reports)
	assert.EqualError(t, err, "operator ID 3 not found")
}

// TestGetOperatorPerformance_EmptyRange tests retrieving data with no epochs within the range.
func TestGetOperatorPerformance_EmptyRange(t *testing.T) {
	// Initialize database with operator data but no epochs within the specified range
	initialData := &storage.Database{
		Operators: storage.OperatorsData{
			OperatorDetails: map[string]storage.OperatorDetails{
				"4": {
					Performance: map[string]domain.Report{
						"150": {Threshold: "0.75"},
					},
				},
			},
		},
	}

	tmpFile := CreateTempDatabaseFile(t, initialData)
	defer os.Remove(tmpFile.Name()) // Clean up

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	operatorID := big.NewInt(4)
	startEpoch := "100"
	endEpoch := "120"

	// Retrieve performance data within the specified range (which has no data)
	reports, err := storageAdapter.GetOperatorPerformance(operatorID, startEpoch, endEpoch)
	assert.NoError(t, err)
	assert.Empty(t, reports) // Expecting an empty map
}
