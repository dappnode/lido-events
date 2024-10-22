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

// EthereumAdapter handles Ethereum client and contract ABIs
type EthereumAdapter struct {
	Client    *ethclient.Client
	Contracts map[string]common.Address // Contract addresses
	ABICache  *cache.ABICache           // ABI cache to fetch ABIs when needed
}

// EventDescriptionMap holds the description for each event
var EventDescriptionMap = map[string]string{
	"DepositedSigningKeysCountChanged":         "🤩 Node Operator's keys received deposits",
	"ELRewardsStealingPenaltyReported":         "🚨 Penalty for stealing EL rewards reported",
	"ELRewardsStealingPenaltySettled":          "🚨 EL rewards stealing penalty confirmed and applied",
	"ELRewardsStealingPenaltyCancelled":        "😮‍💨 Cancelled penalty for stealing EL rewards",
	"InitialSlashingSubmitted":                 "🚨 Initial slashing submitted for one of the validators",
	"KeyRemovalChargeApplied":                  "🔑 Applied charge for key removal",
	"NodeOperatorManagerAddressChangeProposed": "ℹ️ New manager address proposed",
	"NodeOperatorManagerAddressChanged":        "✅ Manager address changed",
	"NodeOperatorRewardAddressChangeProposed":  "ℹ️ New rewards address proposed",
	"NodeOperatorRewardAddressChanged":         "✅ Rewards address changed",
	"StuckSigningKeysCountChanged":             "🚨 Reported stuck keys that were not exited in time",
	"VettedSigningKeysCountDecreased":          "🚨 Uploaded invalid keys",
	"WithdrawalSubmitted":                      "👀 Key withdrawal information submitted",
	"TotalSigningKeysCountChanged":             "👀 New keys uploaded or removed",
	"ValidatorExitRequest":                     "🚨 One of the validators requested to exit",
	"PublicRelease":                            "🎉 Public release of CSM!",
	"DistributionDataUpdated":                  "📈 New rewards distributed",
}

// NewEthereumAdapter initializes the Ethereum adapter with cached ABIs and contract addresses
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

// SubscribeToEvents subscribes to contract events and logs/handles them
func (e *EthereumAdapter) SubscribeToEvents() {
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

		go func() {
			for {
				select {
				case err := <-sub.Err():
					log.Println("Subscription error:", err)
				case vLog := <-logs:
					// Parse the log
					eventName := e.getEventName(parsedABI, vLog.Topics[0])
					if description, found := EventDescriptionMap[eventName]; found {
						log.Printf("Event: %s - %s", eventName, description)
						// Handle event: you can trigger Telegram notification here or store data
					} else {
						log.Printf("Unknown event received: %v", vLog.Topics[0])
					}
				}
			}
		}()
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
