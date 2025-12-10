package execution

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

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

// GetBlockReceipts retrieves all transaction receipts for a given block
// identifier using the eth_getBlockReceipts RPC method. The block identifier
// can be a hex-encoded block number (e.g. "0x1"), a block hash, or tags such
// as "latest".
func (e *Execution) GetBlockReceipts(blockID string) ([]map[string]interface{}, error) {
	// Create the request payload for eth_getBlockReceipts
	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_getBlockReceipts",
		"params":  []interface{}{blockID},
		"id":      1,
	}

	// Marshal the payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request payload for eth_getBlockReceipts: %w", err)
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
		Result []map[string]interface{} `json:"result"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response from execution client: %w", err)
	}

	// If result is nil, treat it as no receipts
	if result.Result == nil {
		return nil, nil
	}

	return result.Result, nil
}

// GetBlockHasReceipts checks if any transaction receipts are available for a
// given block identifier using eth_getBlockReceipts. It returns true if the
// RPC call succeeds and the returned list of receipts is non-empty.
func (e *Execution) GetBlockHasReceipts(blockID string) (bool, error) {
	receipts, err := e.GetBlockReceipts(blockID)
	if err != nil {
		return false, err
	}

	if len(receipts) == 0 {
		return false, nil
	}

	return true, nil
}
