//go:build integration
// +build integration

package csfeeoracle_test

import (
	"context"
	"math/big"
	"os"
	"testing"
	"time"

	csfeeoracle "lido-events/internal/adapters/csFeeOracle"
	"lido-events/internal/application/domain"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
)

// setupCsFeeOracleAdapter initializes the CsFeeOracleAdapter using the RPC_URL environment variable.
func setupCsFeeOracleAdapter(t *testing.T) (*csfeeoracle.CsFeeOracleAdapter, error) {
	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		t.Fatal("RPC_URL environment variable not set")
	}
	rpcClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}

	// Use the provided proxy address for csFeeOracle
	csFeeOracleAddress := common.HexToAddress("0x4D4074628678Bd302921c20573EEa1ed38DdF7FB")
	// Deployment block used as the filter value for the indexed refSlot parameter.
	deploymentBlock := uint64(20935462)
	blockChunkSize := uint64(1000) // Adjust the chunk size as needed

	adapter, err := csfeeoracle.NewCsFeeOracleAdapter(deploymentBlock, rpcClient, csFeeOracleAddress, blockChunkSize)
	return adapter, err
}

// TestScanProcessingStartedEventsIntegration scans for ProcessingStarted events in the given block range.
func TestScanProcessingStartedEventsIntegration(t *testing.T) {
	// Uncomment the following line to skip integration tests if needed.
	t.Skip()

	adapter, err := setupCsFeeOracleAdapter(t)
	assert.NoError(t, err)

	// Set the block range for scanning
	startBlock := uint64(21659279)
	endBlock := uint64(21859289)

	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	// Array to store found events *domain.BindingsProcessingStarted
	foundEvents := make([]*domain.BindingsProcessingStarted, 0)

	// Execute the scan and handle each found event.
	err = adapter.ScanProcessingStartedEvents(ctx, startBlock, &endBlock, func(event *domain.BindingsProcessingStarted, blockNumber *big.Int) error {
		t.Logf("Found ProcessingStarted event at block %s: refSlot=%s, hash=%s",
			blockNumber.String(),
			event.RefSlot.String(),
			event.Hash,
		)
		foundEvents = append(foundEvents, event)
		return nil
	})
	assert.NoError(t, err)
	// Optionally, assert that at least one event was found.
	assert.Greater(t, len(foundEvents), 0, "expected at least one ProcessingStarted event")
}
