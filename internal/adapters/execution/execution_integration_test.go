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
func setupExecutionAdapter(t *testing.T) (*execution.Execution, error) {
	rpcUrl := os.Getenv("RPC_URL")
	if rpcUrl == "" {
		t.Fatal("RPC_URL environment variable not set")
	}
	// Initialize the adapter
	adapter := execution.NewExecution(rpcUrl)
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

// TestGetBlockReceiptsIntegration tests retrieving block receipts for a given block.
// This is a dummy test that queries block 1 and logs the response.
func TestGetBlockReceiptsIntegration(t *testing.T) {
	adapter, err := setupExecutionAdapter(t)
	assert.NoError(t, err)

	// Block identifier for block 1693623 as a hex quantity.
	blockID := "0x19f0d7"

	receipts, err := adapter.GetBlockReceipts(blockID)
	assert.NoError(t, err)

	t.Logf("Block receipts for block %s: %+v", blockID, receipts)
}

// TestGetBlockHasReceiptsIntegration tests checking whether a block has any
// receipts using eth_getBlockReceipts. This is a dummy test for block 1.
func TestGetBlockHasReceiptsIntegration(t *testing.T) {
	adapter, err := setupExecutionAdapter(t)
	assert.NoError(t, err)

	blockID := "0x19f0d7"

	hasReceipts, err := adapter.GetBlockHasReceipts(blockID)
	assert.NoError(t, err)

	t.Logf("Block %s has receipts: %v", blockID, hasReceipts)
}
