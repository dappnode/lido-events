package subscriber

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	bindings "lido-events/internal/adapters/subscriber/bindings"
)

type SubscriberAdapter struct {
	client           *ethclient.Client
	contractBindings *bindings.ContractBindings
}

func NewSubscriberAdapter(
	wsURL string,
	csAccountingAddress, csFeeDistributorAddress, csModuleAddress, veboAddress string,
) (*SubscriberAdapter, error) {

	client, err := ethclient.Dial(wsURL)
	if err != nil {
		return nil, err
	}

	// Initialize the contract bindings with the provided addresses
	contractBindings, err := bindings.InitializeContracts(
		client,
		csAccountingAddress,
		csFeeDistributorAddress,
		csModuleAddress,
		veboAddress,
	)
	if err != nil {
		return nil, err
	}

	return &SubscriberAdapter{
		client:           client,
		contractBindings: contractBindings,
	}, nil
}

// SubscribeToEvents subscribes to events and calls the provided handler function for each event log.
func (es *SubscriberAdapter) SubscribeToEvents(ctx context.Context, handleLog func(logData interface{}) error) error {
	// Prepare contract addresses from ContractBindings
	contractAddresses := []common.Address{
		es.contractBindings.CSAccountingAddress,
		es.contractBindings.CSFeeDistributorAddress,
		es.contractBindings.CSModuleAddress,
		es.contractBindings.VEBOAddress,
	}

	for _, contractAddress := range contractAddresses {
		query := ethereum.FilterQuery{
			Addresses: []common.Address{contractAddress},
		}

		logs := make(chan types.Log)
		sub, err := es.client.SubscribeFilterLogs(ctx, query, logs)
		if err != nil {
			log.Printf("Failed to subscribe to contract %s: %v", contractAddress.Hex(), err)
			continue
		}

		// Handle logs for each contract
		go es.handleLogs(ctx, logs, sub, handleLog)
	}
	return nil
}

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
			var eventData interface{}
			var parseErr error

			// Determine which contract to use for parsing based on the log's address
			switch vLog.Address {
			case es.contractBindings.CSAccountingAddress:
				// Parse using CSAccounting's specific event parser
				eventData, parseErr = es.contractBindings.CSAccounting.parse(vLog)
			case es.contractBindings.CSFeeDistributorAddress:
				// Parse using CSFeeDistributor's specific event parser
				eventData, parseErr = es.contractBindings.CSFeeDistributor.ParseEventFromLog(vLog)
			case es.contractBindings.CSModuleAddress:
				// Parse using CSModule's specific event parser
				eventData, parseErr = es.contractBindings.CSModule.ParseEventFromLog(vLog)
			case es.contractBindings.VEBOAddress:
				// Parse using VEBO's specific event parser
				eventData, parseErr = es.contractBindings.VEBO.ParseValidatorExitRequest(vLog) // -> determine the event name to parse first the logs
			default:
				log.Printf("Unknown contract address: %s", vLog.Address.Hex())
				continue
			}

			if parseErr != nil {
				log.Printf("Failed to parse log data for contract %s: %v", vLog.Address.Hex(), parseErr)
				continue
			}

			// Call the provided handleLog function with the decoded event data
			if err := handleLog(eventData); err != nil {
				log.Printf("Error handling log: %v", err)
			}
		}
	}
}
