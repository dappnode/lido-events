package subscriber

import (
	"context"
	"encoding/json"
	"lido-events/internal/eventhandler"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthereumSubscriber struct {
	client        *ethclient.Client
	contractABIs  map[string]interface{} // Contract address mapped to its ABI JSON data
	eventHandlers *eventhandler.EventHandler
}

func NewEthereumSubscriber(wsURL string, contractABIs map[string]interface{}, handler *eventhandler.EventHandler) (*EthereumSubscriber, error) {
	client, err := ethclient.Dial(wsURL)
	if err != nil {
		return nil, err
	}
	return &EthereumSubscriber{
		client:        client,
		contractABIs:  contractABIs,
		eventHandlers: handler,
	}, nil
}

func (es *EthereumSubscriber) SubscribeToEvents() {
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
		sub, err := es.client.SubscribeFilterLogs(context.Background(), query, logs)
		if err != nil {
			log.Printf("Failed to subscribe to contract %s: %v", contractAddr, err)
			continue
		}

		go es.handleLogs(parsedABI, logs, sub)
	}
}

func (es *EthereumSubscriber) handleLogs(parsedABI abi.ABI, logs <-chan types.Log, sub ethereum.Subscription) {
	for {
		select {
		case err := <-sub.Err():
			log.Printf("Error in subscription: %v", err)
			return
		case vLog := <-logs:
			eventName := ""
			for name, event := range parsedABI.Events {
				if vLog.Topics[0] == event.ID {
					eventName = name
					break
				}
			}

			if eventName != "" {
				es.eventHandlers.HandleEvent(eventName, vLog)
			}
		}
	}
}
