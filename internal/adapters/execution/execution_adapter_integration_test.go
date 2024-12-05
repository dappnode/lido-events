//go:build integration
// +build integration

package execution_test

import (
	"os"
	"testing"

	"lido-events/internal/adapters/execution"

	"github.com/stretchr/testify/assert"
)

// setupExecutionAdapter initializes the ExecutionAdapter
func setupExecutionAdapter(t *testing.T) (*execution.ExecutionAdapter, error) {
	rpcUrl := os.Getenv("RPC_URL")
	if rpcUrl == "" {
		t.Fatal("RPC_URL environment variable not set")
	}
	// Initialize the adapter
	adapter := execution.NewExecutionAdapter(rpcUrl)
	return adapter, nil
}

// TestGetMostRecentBlockNumberIntegration tests retrieving the most recent block number
func TestGetMostRecentBlockNumberIntegration(t *testing.T) {
	adapter, err := setupExecutionAdapter(t)
	assert.NoError(t, err)

	// Call the GetMostRecentBlockNumber method
	blockNumber, err := adapter.GetMostRecentBlockNumber()
	assert.NoError(t, err)

	// Ensure blockNumber is greater than 0, indicating a valid response
	assert.Greater(t, blockNumber, uint64(0), "Expected a non-zero block number")

	// Log the block number for debugging
	t.Logf("Most recent block number: %d", blockNumber)
}

// TestGetBlockTimestampByNumberIntegration tests fetching the timestamp of a specific block
func TestGetBlockTimestampByNumberIntegration(t *testing.T) {
	adapter, err := setupExecutionAdapter(t)
	assert.NoError(t, err)

	// Specify the block number to test
	blockNumber := uint64(2876079)

	// Call the GetBlockTimestampByNumber method
	timestamp, err := adapter.GetBlockTimestampByNumber(blockNumber)
	assert.NoError(t, err)

	// Ensure timestamp is greater than 0, indicating a valid response
	assert.Greater(t, timestamp, uint64(0), "Expected a non-zero timestamp")

	// ensure timestamp is 1733395440
	assert.Equal(t, uint64(1733395440), timestamp, "Expected timestamp to be 1733395440")

	// Log the timestamp for debugging
	t.Logf("Timestamp for block %d: %d (Unix)", blockNumber, timestamp)
}

// TestIsSyncingIntegration tests the IsSyncing method of the ExecutionAdapter
func TestIsSyncingIntegration(t *testing.T) {
	// Set up the execution adapter
	adapter, err := setupExecutionAdapter(t)
	assert.NoError(t, err)

	// Call the IsSyncing method
	isSyncing, err := adapter.IsSyncing()
	assert.NoError(t, err)

	// Assert the result is a valid boolean
	assert.IsType(t, isSyncing, false, "Expected the result to be a boolean")

	// Log the result for debugging
	if isSyncing {
		t.Log("The Ethereum node is syncing.")
	} else {
		t.Log("The Ethereum node is not syncing.")
	}
}
