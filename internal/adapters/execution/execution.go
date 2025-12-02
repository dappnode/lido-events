package execution

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Execution interacts with the Ethereum execution client to fetch block information.
type Execution struct {
	rpcURL string
}

// NewExecution creates a new instance of ExecutionAdapter with the provided execution client URL.
func NewExecution(rpcURL string) *Execution {
	return &Execution{
		rpcURL: rpcURL,
	}
}

// GetMostRecentBlockNumber retrieves the most recent block number from the Ethereum execution client.
func (e *Execution) GetMostRecentBlockNumber() (uint64, error) {
	// Create the request payload for eth_blockNumber
	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_blockNumber",
		"params":  []interface{}{},
		"id":      1,
	}

	// Marshal the payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal request payload for eth_blockNumber: %w", err)
	}

	// Send the request to the execution client
	resp, err := http.Post(e.rpcURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return 0, fmt.Errorf("failed to send request to execution client at %s: %w", e.rpcURL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("unexpected status code %d received from execution client", resp.StatusCode)
	}

	// Parse the response
	var result struct {
		Result string `json:"result"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, fmt.Errorf("failed to decode response from execution client: %w", err)
	}

	// Convert the hexadecimal block number to uint64 using hexutil.DecodeUint64
	blockNumber, err := hexutil.DecodeUint64(result.Result)
	if err != nil {
		return 0, fmt.Errorf("failed to decode block number from result %s: %w", result.Result, err)
	}

	return blockNumber, nil
}

// GetBlockTimestampByNumber retrieves the timestamp of the block with the specified number from the Ethereum execution client.
func (e *Execution) GetBlockTimestampByNumber(blockNumber uint64) (uint64, error) {
	// Convert block number to hexadecimal
	blockNumberHex := fmt.Sprintf("0x%x", blockNumber)

	// Create the request payload for eth_getBlockByNumber
	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_getBlockByNumber",
		"params":  []interface{}{blockNumberHex, false}, // `false` indicates we don't need full transaction objects
		"id":      1,
	}

	// Marshal the payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal request payload for eth_getBlockByNumber: %w", err)
	}

	// Send the request to the execution client
	resp, err := http.Post(e.rpcURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return 0, fmt.Errorf("failed to send request to execution client at %s: %w", e.rpcURL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("unexpected status code %d received from execution client", resp.StatusCode)
	}

	// Parse the response
	var result struct {
		Result struct {
			Timestamp string `json:"timestamp"`
		} `json:"result"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, fmt.Errorf("failed to decode response from execution client: %w", err)
	}

	// Convert the hexadecimal timestamp to uint64
	var timestamp uint64
	_, err = fmt.Sscanf(result.Result.Timestamp, "0x%x", &timestamp)
	if err != nil {
		return 0, fmt.Errorf("failed to parse timestamp from result %s: %w", result.Result.Timestamp, err)
	}

	return timestamp, nil
}

// IsSyncing checks if the Ethereum execution client is currently syncing.
func (e *Execution) IsSyncing() (bool, error) {
	// Create the request payload for eth_syncing
	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_syncing",
		"params":  []interface{}{},
		"id":      1,
	}

	// Marshal the payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return false, fmt.Errorf("failed to marshal request payload for eth_syncing: %w", err)
	}

	// Send the request to the execution client
	resp, err := http.Post(e.rpcURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return false, fmt.Errorf("failed to send request to execution client at %s: %w", e.rpcURL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("unexpected status code %d received from execution client", resp.StatusCode)
	}

	// Parse the response
	var result struct {
		Result interface{} `json:"result"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, fmt.Errorf("failed to decode response from execution client: %w", err)
	}

	// If result is false, the node is not syncing
	if result.Result == false {
		return false, nil
	}

	// If result is a map or object, the node is syncing
	return true, nil
}

// GetTransactionReceipt retrieves the transaction receipt for a given transaction hash.
func (e *Execution) GetTransactionReceipt(txHash common.Hash) (map[string]interface{}, error) {
	// Create the request payload for eth_getTransactionReceipt
	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_getTransactionReceipt",
		"params":  []interface{}{txHash},
		"id":      1,
	}

	// Marshal the payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request payload for eth_getTransactionReceipt: %w", err)
	}

	// Send the request to the execution client
	resp, err := http.Post(e.rpcURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, fmt.Errorf("failed to send request to execution client at %s: %w", e.rpcURL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d received from execution client", resp.StatusCode)
	}

	// Parse the response
	var result struct {
		Result map[string]interface{} `json:"result"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response from execution client: %w", err)
	}

	// Check if the result is null
	if result.Result == nil {
		return nil, nil // Returning nil to indicate no receipt is available
	}

	return result.Result, nil
}

// GetTransactionReceiptExists checks if the transaction receipt exists for a given transaction hash.
// - Reth running as fullnode returns "result": null if the transaction receipt does not exist in the database.
// TODO: test erigon response running it with config not to store receipts
func (e *Execution) GetTransactionReceiptExists(txHash common.Hash) (bool, error) {
	receipt, err := e.GetTransactionReceipt(txHash)
	if err != nil {
		return false, err
	}

	// Check if the receipt is nil
	if receipt == nil {
		return false, nil
	}

	// Check if the receipt is an empty map
	if len(receipt) == 0 {
		return false, nil
	}

	// Check specific fields in the receipt to ensure it is valid
	if _, ok := receipt["transactionHash"]; !ok {
		return false, nil
	}
	if _, ok := receipt["blockNumber"]; !ok {
		return false, nil
	}

	return true, nil
}
