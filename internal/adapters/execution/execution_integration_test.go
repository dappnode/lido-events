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
	t.Skip("integration test temporarily disabled")

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
	t.Skip("integration test temporarily disabled")

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

// TestGetBlockReceiptsIntegration tests retrieving block receipts for blocks
// that are known to have or not have receipts.
func TestGetBlockReceiptsIntegration(t *testing.T) {
	t.Skip("integration test temporarily disabled")

	adapter, err := setupExecutionAdapter(t)
	assert.NoError(t, err)

	// Block 0x1E240 should have receipts.
	blockWithReceipts := "0x1e240"
	receipts, err := adapter.GetBlockReceipts(blockWithReceipts)
	assert.NoError(t, err)
	assert.NotNil(t, receipts)
	assert.Greater(t, len(receipts), 0, "expected block %s to have receipts", blockWithReceipts)
	t.Logf("Block receipts for block %s: %+v", blockWithReceipts, receipts)

	// Block 0x3039 should not have receipts.
	blockWithoutReceipts := "0x3039"
	receipts, err = adapter.GetBlockReceipts(blockWithoutReceipts)
	assert.NoError(t, err)
	if receipts != nil {
		assert.Equal(t, 0, len(receipts), "expected block %s to have no receipts", blockWithoutReceipts)
	}
	t.Logf("Block receipts for block %s: %+v", blockWithoutReceipts, receipts)
}

// TestGetBlockHasReceiptsIntegration tests checking whether blocks have any
// receipts using eth_getBlockReceipts.
func TestGetBlockHasReceiptsIntegration(t *testing.T) {
	t.Skip("integration test temporarily disabled")

	adapter, err := setupExecutionAdapter(t)
	assert.NoError(t, err)

	// Block 0x1E240 should have receipts.
	blockWithReceipts := "0x1e240"
	hasReceipts, err := adapter.GetBlockHasReceipts(blockWithReceipts)
	assert.NoError(t, err)
	assert.True(t, hasReceipts, "expected block %s to have receipts", blockWithReceipts)
	t.Logf("Block %s has receipts: %v", blockWithReceipts, hasReceipts)

	// Block 0x3039 should not have receipts.
	blockWithoutReceipts := "0x3039"
	hasReceipts, err = adapter.GetBlockHasReceipts(blockWithoutReceipts)
	assert.NoError(t, err)
	assert.False(t, hasReceipts, "expected block %s to have no receipts", blockWithoutReceipts)
	t.Logf("Block %s has receipts: %v", blockWithoutReceipts, hasReceipts)
}
