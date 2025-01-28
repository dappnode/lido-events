//go:build integration
// +build integration

package csfeedistributorimpl_test

import (
	"context"
	"math/big"
	"os"
	"testing"
	"time"

	csfeedistributorimpl "lido-events/internal/adapters/csFeeDistributorImpl"
	"lido-events/internal/application/domain"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
)

// setupCsFeeDistributorImplAdapter initializes VeboAdapter with a mocked StoragePort
func setupCsFeeDistributorImplAdapter(t *testing.T) (*csfeedistributorimpl.CsFeeDistributorImplAdapter, error) {
	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		t.Fatal("RPC_URL environment variable not set")
	}
	rpcClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}

	csfeedistributorimplAddress := common.HexToAddress("0xD7ba648C8F72669C6aE649648B516ec03D07c8ED")

	blockChunkSize := uint64(10000)

	// Initialize the adapter with the mock storage
	adapter, err := csfeedistributorimpl.NewCsFeeDistributorImplAdapter(rpcClient, csfeedistributorimplAddress, blockChunkSize)
	return adapter, err
}

// ScanDistributionLogUpdatedEventsIntegration tests scanning for DistributionLogUpdated events
func TestScanDistributionLogUpdatedEventsIntegration(t *testing.T) {
	t.Skip()
	adapter, err := setupCsFeeDistributorImplAdapter(t)
	assert.NoError(t, err)

	// Set the start and end blocks for the scan
	start := uint64(30701)
	end := uint64(2733678)
	operatorId := big.NewInt(1)

	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	// Array to store found logCids
	foundLogCids := make([]string, 0)

	// Execute the scan and handle each found event
	err = adapter.ScanDistributionLogUpdatedEvents(ctx, start, &end, operatorId, func(event *domain.BindingsDistributionLogUpdated, operatorId *big.Int) error {
		foundLogCids = append(foundLogCids, event.LogCid)

		// Log event details for debugging
		t.Logf("Found DistributionLogUpdated event with log cid:: %s", event.LogCid)

		return nil
	})

	assert.NoError(t, err)
	assert.Contains(t, foundLogCids, "QmZvaRtSDQnmtqHAg8Y2sAQKQ6cfDPgL3YtcBFQoVqera5")
}
