package ethereum

import (
	"context"
	"log"

	"lido-events/internal/infrastructure/cache"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthereumAdapter struct {
	Client    *ethclient.Client
	Contracts map[string]common.Address
	ABICache  *cache.ABICache
}

// NewEthereumAdapter initializes the Ethereum adapter
func NewEthereumAdapter(rpcURL string, contractAddresses map[string]string, abiCache *cache.ABICache) (*EthereumAdapter, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}

	contracts := make(map[string]common.Address)
	for name, address := range contractAddresses {
		contracts[name] = common.HexToAddress(address)
	}

	return &EthereumAdapter{
		Client:    client,
		Contracts: contracts,
		ABICache:  abiCache,
	}, nil
}

// SubscribeToEvents subscribes to contract events and passes them to the provided event handler
func (e *EthereumAdapter) SubscribeToEvents(eventHandler func(contractName string, eventName string, vLog types.Log)) {
	for contractName, contractAddress := range e.Contracts {
		// Retrieve the ABI from the cache
		cachedABI, exists := e.ABICache.GetABI(contractAddress.Hex())
		if !exists {
			log.Fatalf("No ABI found for contract address: %s", contractAddress.Hex())
		}

		// Cast the ABI to the correct type
		parsedABI, ok := cachedABI.(abi.ABI)
		if !ok {
			log.Fatalf("Invalid ABI format for contract address: %s", contractAddress.Hex())
		}

		query := ethereum.FilterQuery{
			Addresses: []common.Address{contractAddress},
		}

		logs := make(chan types.Log)
		sub, err := e.Client.SubscribeFilterLogs(context.Background(), query, logs)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Subscribed to events from contract: %s", contractName)

		// Capture variables contractName and parsedABI in the closure
		go func(contractName string, parsedABI abi.ABI) {
			for {
				select {
				case err := <-sub.Err():
					log.Println("Subscription error:", err)
				case vLog := <-logs:
					// Parse the log and delegate handling to the service layer
					eventName := e.getEventName(parsedABI, vLog.Topics[0])
					eventHandler(contractName, eventName, vLog) // Pass the event to the handler
				}
			}
		}(contractName, parsedABI) // Pass contractName and parsedABI to the goroutine
	}
}

// Helper to get event name from the ABI
func (e *EthereumAdapter) getEventName(contractABI abi.ABI, topic common.Hash) string {
	for _, event := range contractABI.Events {
		if event.ID == topic {
			return event.Name
		}
	}
	return "Unknown"
}
