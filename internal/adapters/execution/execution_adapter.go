package execution

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// ExecutionAdapter interacts with the Ethereum execution client to fetch block information.
type ExecutionAdapter struct {
	rpcURL string
}

// NewExecutionAdapter creates a new instance of ExecutionAdapter with the provided execution client URL.
func NewExecutionAdapter(rpcURL string) *ExecutionAdapter {
	return &ExecutionAdapter{
		rpcURL: rpcURL,
	}
}

// GetMostRecentBlockNumber retrieves the most recent block number from the Ethereum execution client.
func (e *ExecutionAdapter) GetMostRecentBlockNumber() (uint64, error) {
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

	// Convert the hexadecimal block number to uint64
	var blockNumber uint64
	_, err = fmt.Sscanf(result.Result, "0x%x", &blockNumber)
	if err != nil {
		return 0, fmt.Errorf("failed to parse block number from result %s: %w", result.Result, err)
	}

	return blockNumber, nil
}

// GetBlockTimestampByNumber retrieves the timestamp of the block with the specified number from the Ethereum execution client.
func (e *ExecutionAdapter) GetBlockTimestampByNumber(blockNumber uint64) (uint64, error) {
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
