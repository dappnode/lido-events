//go:build integration
// +build integration

package storage_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"lido-events/internal/adapters/storage"
	"lido-events/internal/application/domain"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
)

// Helper function to create a temporary file with specified content.
func createTempFileWithContent(t *testing.T, content []byte) *os.File {
	tmpFile, err := os.CreateTemp("", "exit_requests.json")
	assert.NoError(t, err)

	if content != nil {
		_, err = tmpFile.Write(content)
		assert.NoError(t, err)
	}

	err = tmpFile.Close()
	assert.NoError(t, err)

	return tmpFile
}

func TestLoadOrInitializeExitRequests_LoadsExistingFile(t *testing.T) {
	// Create initial data with required fields in types.Log
	initialData := domain.ExitRequests{
		LastProcessedEpoch: 42,
		Requests: map[string]domain.ExitRequest{
			"validatorPubKey1": {
				Event: domain.VeboValidatorExitRequest{
					Raw: types.Log{
						Address: common.HexToAddress("0x0000000000000000000000000000000000000000"),
						Topics:  []common.Hash{common.HexToHash("0x0")},
						Data:    []byte{},
					},
				},
				Status: domain.StatusActiveOngoing,
			},
		},
	}
	data, _ := json.Marshal(initialData)
	tmpFile := createTempFileWithContent(t, data)
	defer os.Remove(tmpFile.Name()) // Clean up

	storage := &storage.Storage{ExitRequestFile: tmpFile.Name()}
	exitRequests, err := storage.LoadOrInitializeExitRequests()
	assert.NoError(t, err)
	assert.Equal(t, initialData, exitRequests)
}

func TestSaveExitRequests_CreatesAndSavesData(t *testing.T) {
	tmpFile := createTempFileWithContent(t, nil)
	defer os.Remove(tmpFile.Name()) // Clean up

	storage := &storage.Storage{ExitRequestFile: tmpFile.Name()}

	exitRequests := domain.ExitRequests{
		LastProcessedEpoch: 99,
		Requests: map[string]domain.ExitRequest{
			"validatorPubKey1": {
				Event: domain.VeboValidatorExitRequest{
					Raw: types.Log{
						Address: common.HexToAddress("0x0000000000000000000000000000000000000000"),
						Topics:  []common.Hash{common.HexToHash("0x0")},
						Data:    []byte{},
					},
				},
				Status: domain.StatusExitedUnslashed,
			},
		},
	}

	err := storage.SaveExitRequests(exitRequests)
	assert.NoError(t, err)

	// Verify the file contents
	data, err := os.ReadFile(tmpFile.Name())
	assert.NoError(t, err)

	var fileData domain.ExitRequests
	err = json.Unmarshal(data, &fileData)
	assert.NoError(t, err)
	assert.Equal(t, exitRequests, fileData)
}

func TestUpdateExitRequestStatus(t *testing.T) {
	// Initial data with a valid Log structure
	initialData := domain.ExitRequests{
		LastProcessedEpoch: 0,
		Requests: map[string]domain.ExitRequest{
			"validatorPubKey1": {
				Event: domain.VeboValidatorExitRequest{
					Raw: types.Log{
						Address: common.HexToAddress("0x0000000000000000000000000000000000000000"),
						Topics:  []common.Hash{common.HexToHash("0x0")},
						Data:    []byte{},
					},
				},
				Status: domain.StatusPendingInitialized,
			},
		},
	}
	data, _ := json.Marshal(initialData)
	tmpFile := createTempFileWithContent(t, data)
	defer os.Remove(tmpFile.Name()) // Clean up

	storage := &storage.Storage{ExitRequestFile: tmpFile.Name()}
	err := storage.UpdateExitRequestStatus("validatorPubKey1", domain.StatusActiveOngoing)
	assert.NoError(t, err)

	// Verify the updated status
	data, err = os.ReadFile(tmpFile.Name())
	assert.NoError(t, err)

	var fileData domain.ExitRequests
	err = json.Unmarshal(data, &fileData)
	assert.NoError(t, err)
	assert.Equal(t, domain.StatusActiveOngoing, fileData.Requests["validatorPubKey1"].Status)
}

func TestSaveExitRequest(t *testing.T) {
	tmpFile := createTempFileWithContent(t, []byte(`{"LastProcessedEpoch":0,"Requests":{}}`))
	defer os.Remove(tmpFile.Name()) // Clean up

	storage := &storage.Storage{ExitRequestFile: tmpFile.Name()}
	exitRequest := domain.ExitRequest{
		Event: domain.VeboValidatorExitRequest{
			Raw: types.Log{
				Address:     common.HexToAddress("0x0000000000000000000000000000000000000000"),
				Topics:      []common.Hash{common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000")},
				Data:        []byte{},
				BlockNumber: 0,
				TxHash:      common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
				BlockHash:   common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
				Index:       0,
				Removed:     false,
			},
		},
		Status: domain.StatusExitedUnslashed,
	}

	err := storage.SaveExitRequest("validatorPubKey2", exitRequest)
	assert.NoError(t, err)

	// Verify that the new exit request was added
	data, err := os.ReadFile(tmpFile.Name())
	assert.NoError(t, err)

	var fileData domain.ExitRequests
	err = json.Unmarshal(data, &fileData)
	assert.NoError(t, err)

	// Debug output to compare the structures
	fmt.Printf("Expected: %+v\n", exitRequest)
	fmt.Printf("Actual: %+v\n", fileData.Requests["validatorPubKey2"])

	// Compare the expected and actual data
	assert.Equal(t, exitRequest, fileData.Requests["validatorPubKey2"])
}



