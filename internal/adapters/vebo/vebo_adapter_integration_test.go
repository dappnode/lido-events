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
		t.Skip("WS_URL is not set")
	}

	// Create the mock StoragePort
	mockStorage := new(mocks.MockStoragePort)

	// Define initial operator IDs as required by the test
	mockStorage.On("GetOperatorIds").Return([]*big.Int{big.NewInt(20)}, nil)

	veboAddress := common.HexToAddress("0x0De4Ea0184c2ad0BacA7183356Aea5B8d5Bf5c6e")

	// Initialize the adapter with the mock storage
	adapter, err := vebo.NewVeboAdapter(wsURL, veboAddress, mockStorage)
	return adapter, mockStorage, err
}

// TestScanVeboValidatorExitRequestEventIntegration tests scanning for ValidatorExitRequest events with mocked data
func TestScanVeboValidatorExitRequestEventIntegration(t *testing.T) {
	adapter, mockStorage, err := setupVeboAdapter(t)
	assert.NoError(t, err)

	// Set the start and end blocks for the scan
	start := uint64(21071257)
	end := uint64(21071259)

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
	assert.Contains(t, foundEvents, "370637")
	assert.Equal(t, "b07c5cc5abd773d24c460140f872d24fb6584027626a4cba7b43e4f79f22c7b4846ea55d11ba96a322ba663c55cf3523", foundEvents["370637"].ValidatorPubkey)
	assert.Equal(t, uint64(21071258), foundEvents["370637"].BlockNumber)

	assert.Contains(t, foundEvents, "370638")
	assert.Equal(t, "92404310ad54447f6ab3277351a0da37f9b75696409056e855cfbe9838c649def1ca24cba994f6394f436c940d4cc3bc", foundEvents["370638"].ValidatorPubkey)
	assert.Equal(t, uint64(21071258), foundEvents["370638"].BlockNumber)

	// Ensure all expected mock calls were made
	mockStorage.AssertCalled(t, "GetOperatorIds") // Ensure GetOperatorIds was actually called
}
