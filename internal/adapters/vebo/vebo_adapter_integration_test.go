//go:build integration
// +build integration

package vebo_test

import (
	"context"
	"math/big"
	"os"
	"testing"
	"time"

	"lido-events/internal/adapters/vebo"
	"lido-events/internal/application/domain"
	"lido-events/internal/mocks"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

// setupVeboAdapter initializes VeboAdapter with a mocked StoragePort
func setupVeboAdapter(t *testing.T) (*vebo.VeboAdapter, *mocks.MockStoragePort, error) {
	wsURL := os.Getenv("WS_URL")
	if wsURL == "" {
		t.Fatal("WS_URL environment variable not set")
	}

	// Create the mock StoragePort
	mockStorage := new(mocks.MockStoragePort)

	// Define initial operator IDs as required by the test
	mockStorage.On("GetOperatorIds").Return([]*big.Int{big.NewInt(2535)}, nil)

	veboAddress := common.HexToAddress("0xffDDF7025410412deaa05E3E1cE68FE53208afcb")

	blockChunkSize := uint64(10000)

	// Initialize the adapter with the mock storage
	adapter, err := vebo.NewVeboAdapter(wsURL, veboAddress, mockStorage, blockChunkSize)
	return adapter, mockStorage, err
}

// TestScanVeboValidatorExitRequestEventIntegration tests scanning for ValidatorExitRequest events with mocked data
func TestScanVeboValidatorExitRequestEventIntegration(t *testing.T) {
	adapter, mockStorage, err := setupVeboAdapter(t)
	assert.NoError(t, err)

	// Set the start and end blocks for the scan
	start := uint64(2689810)
	end := uint64(2812210)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Map to store and verify events found during the scan
	foundEvents := map[string]struct {
		ValidatorPubkey string
		BlockNumber     uint64
	}{}

	// Execute the scan and handle each found event
	err = adapter.ScanVeboValidatorExitRequestEvent(ctx, start, &end, func(event *domain.VeboValidatorExitRequest) error {
		foundEvents[event.ValidatorIndex.String()] = struct {
			ValidatorPubkey string
			BlockNumber     uint64
		}{
			ValidatorPubkey: common.Bytes2Hex(event.ValidatorPubkey),
			BlockNumber:     event.Raw.BlockNumber,
		}

		// Log event details for debugging
		t.Logf("ValidatorIndex: %s, ValidatorPubkey: %s, BlockNumber: %d",
			event.ValidatorIndex.String(),
			common.Bytes2Hex(event.ValidatorPubkey),
			event.Raw.BlockNumber)

		return nil
	})

	// Assertions for the expected events
	assert.NoError(t, err)
	assert.Contains(t, foundEvents, "1802081")
	assert.Equal(t, "972255d9a5085d082d485f1e17999b38967e022057aba66a477cd93bce5cfa980bc42df82b208987ed46b9cdbc7b5fcb", foundEvents["1802081"].ValidatorPubkey)
	assert.Equal(t, uint64(2790523), foundEvents["1802081"].BlockNumber)

	// Ensure all expected mock calls were made
	mockStorage.AssertCalled(t, "GetOperatorIds") // Ensure GetOperatorIds was actually called
}
