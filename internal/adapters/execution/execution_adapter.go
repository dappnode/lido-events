package execution

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

// ExecutionAdapter interacts with the Ethereum execution client to fetch block information.
type ExecutionAdapter struct {
	rpcURL string
}

// NewExecutionAdapter creates a new instance of ExecutionAdapter with the provided execution client URL.
func NewExecutionAdapter(rpcURL string) *ExecutionAdapter {
	return &ExecutionAdapter{
		rpcURL,
	}
}

// GetMostRecentBlockNumber retrieves the most recent block number from the Ethereum execution client.
// It calls the eth_blockNumber method on the client.
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
		log.Printf("failed to marshal request payload: %v", err)
		return 0, err
	}

	// Send the request to the execution client
	resp, err := http.Post(e.rpcURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Printf("failed to send request to execution client: %v", err)
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errorMessage := fmt.Sprintf("unexpected status code fetching block number: %d", resp.StatusCode)
		log.Println(errorMessage)
		return 0, errors.New(errorMessage)
	}

	// Parse the response
	var result struct {
		Result string `json:"result"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Printf("failed to decode response: %v", err)
		return 0, err
	}

	// Convert the hexadecimal block number to uint64
	var blockNumber uint64
	_, err = fmt.Sscanf(result.Result, "0x%x", &blockNumber)
	if err != nil {
		log.Printf("failed to parse block number: %v", err)
		return 0, err
	}

	return blockNumber, nil
}
