package subscriber

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type SubscriberAdapter struct {
	client       *ethclient.Client
	contractABIs map[string]interface{} // Contract address mapped to its ABI JSON data
}

type DecodedLog struct {
	Address     common.Address
	Topics      []common.Hash
	Data        []byte
	BlockNumber uint64
	TxHash      common.Hash
	EventData   map[string]interface{} // Decoded event fields
}

func NewSubscriberAdapter(wsURL string, contractABIs map[string]interface{}) (*SubscriberAdapter, error) {
	client, err := ethclient.Dial(wsURL)
	if err != nil {
		return nil, err
	}
	return &SubscriberAdapter{
		client:       client,
		contractABIs: contractABIs,
	}, nil
}

// SubscribeToEvents subscribes to events and calls the provided handler function for each event log.
func (es *SubscriberAdapter) SubscribeToEvents(ctx context.Context, handleLog func(logData interface{}) error) error {
	for contractAddr, abiData := range es.contractABIs {
		contractAddress := common.HexToAddress(contractAddr)

		// Assert that the ABI data is of type abi.ABI
		_, ok := abiData.(abi.ABI)
		if !ok {
			log.Printf("Failed to convert ABI data for contract %s to abi.ABI", contractAddr)
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
		// go es.handleLogs(ctx, logs, sub, handleLog, parsedABI)
		go es.handleLogs(ctx, logs, sub, handleLog)
	}
	return nil // Adjust as needed
}

// TODO: logData interface{} should be replaced with a custom type. How does an ethereum log look like?
func (es *SubscriberAdapter) handleLogs(ctx context.Context, logs <-chan types.Log, sub ethereum.Subscription, handleLog func(logData interface{}) error) {
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
			// This function needs to be implemented by the core aplication, since
			// it will store the log data in the database
			if err := handleLog(vLog); err != nil {
				log.Printf("Error handling log: %v", err)
			}
		}
	}
}
