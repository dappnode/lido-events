package ethereum

import (
	"bytes"
	"context"
	"io"
	"log"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type MyEthereumAdapter struct {
	Client       *ethclient.Client
	ContractABIs map[string]abi.ABI
	Contracts    map[string]common.Address
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

// NewEthereumAdapter initializes the Ethereum adapter for multiple contracts
func NewEthereumAdapter(rpcURL string, abiFilePaths []string, contractAddresses []string) (*MyEthereumAdapter, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}

	contractABIs := make(map[string]abi.ABI)
	contracts := make(map[string]common.Address)

	// Load each ABI
	for i, abiFilePath := range abiFilePaths {
		abiFile, err := os.ReadFile(abiFilePath) // Replaced ioutil.ReadFile with os.ReadFile
		if err != nil {
			return nil, err
		}

		parsedABI, err := abi.JSON(io.NopCloser(bytes.NewReader(abiFile))) // Replaced ioutil.NopCloser
		if err != nil {
			return nil, err
		}

		contractABIs[abiFilePath] = parsedABI
		contracts[abiFilePath] = common.HexToAddress(contractAddresses[i])
	}

	return &MyEthereumAdapter{
		Client:       client,
		ContractABIs: contractABIs,
		Contracts:    contracts,
	}, nil
}

// SubscribeToEvents subscribes to contract events and logs/handles them
func (e *MyEthereumAdapter) SubscribeToEvents() {
	for contractName, contractABI := range e.ContractABIs {
		query := ethereum.FilterQuery{
			Addresses: []common.Address{e.Contracts[contractName]},
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
					eventName := e.getEventName(contractABI, vLog.Topics[0])
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
func (e *MyEthereumAdapter) getEventName(contractABI abi.ABI, topic common.Hash) string {
	for _, event := range contractABI.Events {
		if event.ID == topic {
			return event.Name
		}
	}
	return "Unknown"
}
