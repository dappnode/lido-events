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

// TestSaveReport_NewOperator tests saving report data for a new operator ID.
func TestSaveReport_NewOperator(t *testing.T) {
	tmpFile := CreateTempDatabaseFile(t, nil) // Helper function to create a temporary database file
	defer os.Remove(tmpFile.Name())           // Clean up

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	operatorID := big.NewInt(1)
	report := domain.Report{
		Frame:     [2]int{81055, 81504},
		Threshold: 0.85,
		Data: domain.Data{
			Validators: map[string]domain.Validator{
				"validator1": {Perf: domain.Performance{Assigned: 123, Included: 120}},
			},
		},
	}

	// Save the report data for the new operator
	err := storageAdapter.SaveReport(operatorID, report)
	assert.NoError(t, err)

	// Load database and verify the data was saved
	db, err := storageAdapter.LoadDatabase()
	assert.NoError(t, err)
	frameKey := "81055-81504"
	savedReport, exists := db.Operators[operatorID.String()].Reports[frameKey]
	assert.True(t, exists)
	assert.Equal(t, report, savedReport)
}

// TestSaveReport_ExistingOperator tests updating report data for an existing operator ID.
func TestSaveReport_ExistingOperator(t *testing.T) {
	// Initialize database with existing report data for an operator
	initialData := &storage.Database{
		Operators: map[string]storage.OperatorData{
			"1": {
				Reports: domain.Reports{
					"81055-81504": {Threshold: 0.80},
				},
			},
		},
	}

	tmpFile := CreateTempDatabaseFile(t, initialData)
	defer os.Remove(tmpFile.Name()) // Clean up

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	operatorID := big.NewInt(1)
	updatedReport := domain.Report{
		Frame:     [2]int{81055, 81504},
		Threshold: 0.85,
		Data: domain.Data{
			Validators: map[string]domain.Validator{
				"validator1": {Perf: domain.Performance{Assigned: 130, Included: 125}},
			},
		},
	}

	// Update the report data for the existing operator
	err := storageAdapter.SaveReport(operatorID, updatedReport)
	assert.NoError(t, err)

	// Verify that the report data was updated
	db, err := storageAdapter.LoadDatabase()
	assert.NoError(t, err)
	frameKey := "81055-81504"
	assert.Equal(t, updatedReport, db.Operators[operatorID.String()].Reports[frameKey])
}

// TestGetReports tests retrieving all report data for a specified operator ID.
func TestGetReports(t *testing.T) {
	// Initialize database with multiple reports for an operator
	initialData := &storage.Database{
		Operators: map[string]storage.OperatorData{
			"2": {
				Reports: domain.Reports{
					"81055-81504": {Threshold: 0.80},
					"81505-82000": {Threshold: 0.85},
				},
			},
		},
	}

	tmpFile := CreateTempDatabaseFile(t, initialData)
	defer os.Remove(tmpFile.Name()) // Clean up

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	operatorID := big.NewInt(2)

	// Retrieve all report data for the operator
	reports, err := storageAdapter.GetReports(operatorID)
	assert.NoError(t, err)

	// Verify that all reports are returned
	expectedReports := domain.Reports{
		"81055-81504": {Threshold: 0.80},
		"81505-82000": {Threshold: 0.85},
	}
	assert.Equal(t, expectedReports, reports)
}

// TestGetReports_NoData tests retrieving report data when no data exists for the specified operator.
func TestGetReports_NoData(t *testing.T) {
	tmpFile := CreateTempDatabaseFile(t, nil) // Empty initial data
	defer os.Remove(tmpFile.Name())           // Clean up

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	operatorID := big.NewInt(3)

	// Attempt to retrieve report data for a non-existent operator
	reports, err := storageAdapter.GetReports(operatorID)
	assert.Error(t, err)
	assert.Nil(t, reports)
	assert.EqualError(t, err, "operator ID 3 not found")
}

// TestGetReports_EmptyData tests retrieving data for an operator with no reports.
func TestGetReports_EmptyData(t *testing.T) {
	// Initialize database with operator data but no reports
	initialData := &storage.Database{
		Operators: map[string]storage.OperatorData{
			"4": {
				Reports: make(domain.Reports),
			},
		},
	}

	tmpFile := CreateTempDatabaseFile(t, initialData)
	defer os.Remove(tmpFile.Name()) // Clean up

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	operatorID := big.NewInt(4)

	// Retrieve report data for the operator (expecting no reports)
	reports, err := storageAdapter.GetReports(operatorID)
	assert.NoError(t, err)
	assert.Empty(t, reports) // Expecting an empty map
}
