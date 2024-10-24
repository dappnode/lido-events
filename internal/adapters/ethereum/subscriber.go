package ethereum

import (
	"context"
	"encoding/json"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthereumSubscriber struct {
	client       *ethclient.Client
	contractABIs map[string]interface{} // Contract address mapped to its ABI JSON data
}

func NewEthereumSubscriber(wsURL string, contractABIs map[string]interface{}) (*EthereumSubscriber, error) {
	client, err := ethclient.Dial(wsURL)
	if err != nil {
		return nil, err
	}
	return &EthereumSubscriber{
		client:       client,
		contractABIs: contractABIs,
	}, nil
}

// SubscribeToEvents subscribes to events and calls the provided handler function for each event log.
func (es *EthereumSubscriber) SubscribeToEvents(ctx context.Context, handleLog func(logData interface{}) error) error {
	for contractAddr, abiData := range es.contractABIs {
		contractAddress := common.HexToAddress(contractAddr)

		// Assert that the ABI data is a JSON string and parse it
		abiJSONBytes, ok := abiData.([]byte)
		if !ok {
			log.Printf("Failed to convert ABI data for contract %s to []byte", contractAddr)
			continue
		}

		// Parse the ABI JSON
		var parsedABI abi.ABI
		if err := json.Unmarshal(abiJSONBytes, &parsedABI); err != nil {
			log.Printf("Failed to parse ABI JSON for contract %s: %v", contractAddr, err)
			continue
		}

		query := ethereum.FilterQuery{
			Addresses: []common.Address{contractAddress},
		}

		logs := make(chan types.Log)
		sub, err := es.client.SubscribeFilterLogs(ctx, query, logs)
		if err != nil {
			log.Printf("Failed to subscribe to contract %s: %v", contractAddr, err)
			continue
		}

		// Handle logs
		go es.handleLogs(ctx, logs, sub, handleLog)
	}
	return nil // TODO: should we return something?
}

// TODO: logData interface{} should be replaced with a custom type. How does an ethereum log look like?
func (es *EthereumSubscriber) handleLogs(ctx context.Context, logs <-chan types.Log, sub ethereum.Subscription, handleLog func(logData interface{}) error) {
	for {
		select {
		case <-ctx.Done():
			log.Printf("Subscription context cancelled: %v", ctx.Err())
			return
		case err := <-sub.Err():
			log.Printf("Error in subscription: %v", err)
			return
		case vLog := <-logs:
			// Call the provided handleLog function with the log data
			// This function needs to be implemented by the core application, since
			// it will store the log data in the database
			if err := handleLog(vLog); err != nil {
				log.Printf("Error handling log: %v", err)
			}
		}
	}
}
