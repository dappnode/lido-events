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
		t.Skip("RPC_URL is not set")
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
