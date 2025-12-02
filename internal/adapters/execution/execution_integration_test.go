//go:build integration
// +build integration

package execution_test

import (
	"os"
	"testing"

	"lido-events/internal/adapters/execution"

	"github.com/ethereum/go-ethereum/common"
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

// TestGetTransactionReceiptIntegration tests retrieving the transaction receipt
func TestGetTransactionReceiptIntegration(t *testing.T) {
	adapter, err := setupExecutionAdapter(t)
	assert.NoError(t, err)

	// Specify a transaction hash to test
	txHash := common.HexToHash("0x1475719ecbb73b28bc531bb54b37695df1bf6b71c6d2bf1d28b4efa404867e26")

	// Call the GetTransactionReceipt method
	receipt, err := adapter.GetTransactionReceipt(txHash)
	assert.NoError(t, err)

	// Ensure receipt is not nil
	assert.NotNil(t, receipt, "Expected a non-nil transaction receipt")

	// Log the receipt for debugging
	t.Logf("Transaction receipt: %+v", receipt)
}

// TestGetTransactionReceiptExistsIntegration tests checking if a transaction receipt exists
func TestGetTransactionReceiptExistsIntegration(t *testing.T) {
	adapter, err := setupExecutionAdapter(t)
	assert.NoError(t, err)

	// Specify a transaction hash to test
	txHash := common.HexToHash("0x1475719ecbb73b28bc531bb54b37695df1bf6b71c6d2bf1d28b4efa404867e26")

	// Call the GetTransactionReceiptExists method
	exists, err := adapter.GetTransactionReceiptExists(txHash)
	assert.NoError(t, err)

	// Ensure exists is true
	assert.True(t, exists, "Expected the transaction receipt to exist")

	// Log the result for debugging
	t.Logf("Transaction receipt exists for hash %s: %v", txHash.Hex(), exists)
}
