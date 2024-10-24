package infrastructure

import (
	"encoding/json"
	"lido-events/internal/domain/entities"
	"log"
	"os"
)

// Config struct representing the full configuration for the application
type Config struct {
	OperatorID entities.OperatorId     `json:"operatorId"`
	Telegram   entities.TelegramConfig `json:"telegram"`
	Network    NetworkConfig           `json:"network"`
}

// NetworkConfig struct representing relevant network-specific configurations
type NetworkConfig struct {
	CSMStakingModuleID int
	EtherscanURL       string
	BeaconchainURL     string
	CSMUIURL           string
	ContractABIs       map[string]string // Map of contract addresses to ABI file paths
}

// LoadAppConfig loads the operatorId and telegram details from a JSON file
func LoadAppConfig(filePath string) (Config, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
		return Config{}, err
	}

	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatalf("Failed to parse config JSON: %v", err)
		return Config{}, err
	}

	return config, nil
}

// LoadNetworkConfig loads the relevant network-specific configuration and contract ABIs
func LoadNetworkConfig() (NetworkConfig, error) {
	network := os.Getenv("NETWORK")

	var config NetworkConfig

	switch network {
	case "holesky":
		config = NetworkConfig{
			CSMStakingModuleID: 4,
			EtherscanURL:       "https://holesky.etherscan.io",
			BeaconchainURL:     "https://holesky.beaconcha.in",
			CSMUIURL:           "https://csm.testnet.fi",
			ContractABIs: map[string]string{
				"0x4562c3e63c2e586cD1651B958C22F88135aCAd4f": "abi/CSAccounting.json",
				"0xc093e53e8F4b55A223c18A2Da6fA00e60DD5EFE1": "abi/CSFeeDistributor.json",
				"0xffDDF7025410412deaa05E3E1cE68FE53208afcb": "abi/VEBO.json",
			},
		}
	case "mainnet":
		config = NetworkConfig{
			CSMStakingModuleID: 3,
			EtherscanURL:       "https://etherscan.io",
			BeaconchainURL:     "https://beaconcha.in",
			CSMUIURL:           "https://csm.lido.fi",
			ContractABIs: map[string]string{
				"0xdA7dE2ECdDfccC6c3AF10108Db212ACBBf9EA83F": "abi/CSAccounting.json",
				"0x4d72BFF1BeaC69925F8Bd12526a39BAAb069e5Da": "abi/CSFeeDistributor.json",
				"0x0De4Ea0184c2ad0BacA7183356Aea5B8d5Bf5c6e": "abi/VEBO.json",
			},
		}
	default:
		log.Fatalf("Unknown network: %s", network)
	}

	return config, nil
}
