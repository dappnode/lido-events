package config

import (
	"lido-events/internal/logger"
	"math/big"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

type Config struct {
	Network string

	DBDirectory        string
	MevBoostDnpName    string
	DappmanagerUrl     string
	SignerUrl          string
	IpfsUrl            string
	RpcUrl             string
	CSMStakingModuleID *big.Int
	EtherscanURL       string
	BeaconchainURL     string
	BeaconchaUrl       string
	CSMUIURL           string
	StakersUiUrl       string
	BrainUrl           string
	ApiPort            uint64
	LidoDnpName        string

	CORS []string

	// Individual contract addresses
	CSFeeDistributorProxyAddress  common.Address
	VEBOAddress                   common.Address
	MEVBoostRelaysAllowListAddres common.Address
	CSParametersRegistryAddress   common.Address

	// Lido specifics
	LidoKeysApiUrl          string
	ProxyApiPort            uint64
	DefaultAllowedExitDelay uint64
	ExitDelayMultiplier     uint64

	// Blockchain
	SecondsPerSlot          uint64
	BlockChunkSize          uint64
	BlockScannerMinDistance uint64
}

// Helper function to parse and validate CORS from environment variable
func parseCORS(envCORS string, defaultCORS []string) []string {
	if envCORS == "" {
		// Return default CORS if environment variable is not set
		return defaultCORS
	}

	// Split the CORS string by commas and trim spaces
	corsList := strings.Split(envCORS, ",")
	for i, origin := range corsList {
		corsList[i] = strings.TrimSpace(origin)
		// Validate that each origin is not empty
		if corsList[i] == "" {
			logger.Fatal("Invalid CORS entry in environment variable")
		}
	}

	return corsList
}

func LoadNetworkConfig() (Config, error) {
	dappmanagerUrl := os.Getenv("DAPPMANAGER_URL")
	if dappmanagerUrl == "" {
		dappmanagerUrl = "http://my.dappnode"
	}

	dbDirectory := os.Getenv("DB_PATH")
	if dbDirectory == "" {
		cwd, err := os.Getwd()
		if err != nil {
			logger.Fatal("Failed to get working directory: %v", err)
		}
		dbDirectory = cwd
	}

	apiPortStr := os.Getenv("API_PORT")
	apiPort := uint64(8080)
	if apiPortStr != "" {
		// Try to parse the port as uint64
		if port, err := strconv.ParseUint(apiPortStr, 10, 64); err == nil {
			apiPort = port
		} else {
			logger.Fatal("Invalid API_PORT value: %s", apiPortStr)
		}
	}

	proxyApiPortStr := os.Getenv("PROXY_API_PORT")
	proxyApiPort := uint64(8081)
	if proxyApiPortStr != "" {
		// Try to parse the port as uint64
		if port, err := strconv.ParseUint(proxyApiPortStr, 10, 64); err == nil {
			proxyApiPort = port
		} else {
			logger.Fatal("Invalid PROXY_API_PORT value: %s", proxyApiPortStr)
		}
	}

	network := os.Getenv("NETWORK")
	// Default to hoodi
	if network == "" {
		network = "hoodi"
	}

	ipfsUrl := os.Getenv("IPFS_URL")
	// Default to local http://ipfs.dappnode:5001
	if ipfsUrl == "" {
		ipfsUrl = "http://ipfs.dappnode:5001"
	}

	// Retrieve other necessary environment variables
	rpcURL := os.Getenv("RPC_URL")
	beaconchainURL := os.Getenv("BEACONCHAIN_URL")
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "INFO" // Default log level
	}
	blockChunkSize := uint64(100000)
	blockChunkSizeStr := os.Getenv("BLOCK_CHUNK_SIZE")
	if blockChunkSizeStr != "" {
		// Try to parse the block chunk size as uint64
		if size, err := strconv.ParseUint(blockChunkSizeStr, 10, 64); err == nil {
			blockChunkSize = size
		} else {
			logger.Fatal("Invalid BLOCK_CHUNK_SIZE value: %s", blockChunkSizeStr)
		}
	}

	blockScannerMinDistance := uint64(320) // 320 blocks = 10 epochs
	blockScannerMinDistanceStr := os.Getenv("BLOCK_SCANNER_MIN_DISTANCE")
	if blockScannerMinDistanceStr != "" {
		// Try to parse the block scanner min distance as uint64
		if distance, err := strconv.ParseUint(blockScannerMinDistanceStr, 10, 64); err == nil {
			blockScannerMinDistance = distance
		} else {
			logger.Fatal("Invalid BLOCK_SCANNER_MIN_DISTANCE value: %s", blockScannerMinDistanceStr)
		}
	}

	defaultAllowedExitDelay := uint64(345600) // 4 days in seconds
	defaultAllowedExitDelayStr := os.Getenv("DEFAULT_ALLOWED_EXIT_DELAY")
	if defaultAllowedExitDelayStr != "" {
		// Try to parse the default allowed exit delay as uint64
		if delay, err := strconv.ParseUint(defaultAllowedExitDelayStr, 10, 64); err == nil {
			defaultAllowedExitDelay = delay
		} else {
			logger.Fatal("Invalid DEFAULT_ALLOWED_EXIT_DELAY value: %s", defaultAllowedExitDelayStr)
		}
	}

	// must be integer and greater than or equal to 1
	exitDelayMultiplier := uint64(2)
	exitDelayMultiplierStr := os.Getenv("EXIT_DELAY_MULTIPLIER")
	if exitDelayMultiplierStr != "" {
		// Try to parse the exit delay multiplier as uint64
		if multiplier, err := strconv.ParseUint(exitDelayMultiplierStr, 10, 64); err == nil {
			if multiplier < 1 {
				logger.Fatal("EXIT_DELAY_MULTIPLIER must be greater than or equal to 1")
			}
			exitDelayMultiplier = multiplier
		} else {
			logger.Fatal("Invalid EXIT_DELAY_MULTIPLIER value: %s", exitDelayMultiplierStr)
		}
	}

	corsEnv := os.Getenv("CORS")
	var config Config

	switch network {
	case "hoodi":
		// Configure default values for the hoodi network
		if rpcURL == "" {
			rpcURL = "http://execution.hoodi.dncore.dappnode:8545"
		}
		if beaconchainURL == "" {
			beaconchainURL = "http://beacon-chain.hoodi.dncore.dappnode:3500"
		}
		config = Config{
			Network:                       network,
			DBDirectory:                   dbDirectory,
			MevBoostDnpName:               "mev-boost-hoodi.dnp.dappnode.eth",
			DappmanagerUrl:                dappmanagerUrl,
			SignerUrl:                     "http://signer.hoodi.dncore.dappnode:9000",
			IpfsUrl:                       ipfsUrl,
			RpcUrl:                        rpcURL,
			CSMStakingModuleID:            big.NewInt(4),
			EtherscanURL:                  "https://hoodi.etherscan.io",
			BeaconchainURL:                beaconchainURL,
			BeaconchaUrl:                  "https://hoodi.beaconcha.in",
			CSMUIURL:                      "https://csm.testnet.fi",
			StakersUiUrl:                  "http://my.dappnode/stakers/hoodi",
			BrainUrl:                      "http://brain.web3signer-hoodi.dappnode",
			ApiPort:                       apiPort,
			LidoDnpName:                   "llido-csm-hoodi.dnp.dappnode.eth",
			CORS:                          parseCORS(corsEnv, []string{"http://ui.lido-csm-hoodi.dappnode", "http://my.dappnode"}),
			CSFeeDistributorProxyAddress:  common.HexToAddress("0xaCd9820b0A2229a82dc1A0770307ce5522FF3582"),
			VEBOAddress:                   common.HexToAddress("0x8664d394C2B3278F26A1B44B967aEf99707eeAB2"),
			MEVBoostRelaysAllowListAddres: common.HexToAddress("0x279d3A456212a1294DaEd0faEE98675a52E8A4Bf"),
			CSParametersRegistryAddress:   common.HexToAddress("0xA4aD5236963f9Fe4229864712269D8d79B65C5Ad"),
			LidoKeysApiUrl:                "https://keys-api-hoodi.testnet.fi",
			ProxyApiPort:                  proxyApiPort,
			DefaultAllowedExitDelay:       defaultAllowedExitDelay,
			ExitDelayMultiplier:           exitDelayMultiplier,
			SecondsPerSlot:                12,
			BlockChunkSize:                blockChunkSize,
			BlockScannerMinDistance:       blockScannerMinDistance,
		}
	case "mainnet":
		// Configure default values for the mainnet
		if rpcURL == "" {
			rpcURL = "http://execution.mainnet.dncore.dappnode:8545"
		}
		if beaconchainURL == "" {
			beaconchainURL = "http://beacon-chain.mainnet.dncore.dappnode:3500"
		}
		config = Config{
			Network:                       network,
			DBDirectory:                   dbDirectory,
			MevBoostDnpName:               "mev-boost.dnp.dappnode.eth",
			DappmanagerUrl:                dappmanagerUrl,
			SignerUrl:                     "http://signer.mainnet.dncore.dappnode:9000",
			IpfsUrl:                       ipfsUrl,
			RpcUrl:                        rpcURL,
			CSMStakingModuleID:            big.NewInt(3),
			EtherscanURL:                  "https://etherscan.io",
			BeaconchainURL:                beaconchainURL,
			BeaconchaUrl:                  "https://beaconcha.in",
			CSMUIURL:                      "https://csm.lido.fi",
			StakersUiUrl:                  "http://my.dappnode/stakers/ethereum",
			BrainUrl:                      "http://brain.web3signer.dappnode",
			LidoDnpName:                   "llido-csm-mainnet.dnp.dappnode.eth",
			ApiPort:                       apiPort,
			CORS:                          parseCORS(corsEnv, []string{"http://ui.lido-csm-mainnet.dappnode", "http://my.dappnode"}),
			CSFeeDistributorProxyAddress:  common.HexToAddress("0xD99CC66fEC647E68294C6477B40fC7E0F6F618D0"),
			VEBOAddress:                   common.HexToAddress("0x0De4Ea0184c2ad0BacA7183356Aea5B8d5Bf5c6e"),
			MEVBoostRelaysAllowListAddres: common.HexToAddress("0xF95f069F9AD107938F6ba802a3da87892298610E"),
			CSParametersRegistryAddress:   common.HexToAddress("0x9D28ad303C90DF524BA960d7a2DAC56DcC31e428"),
			LidoKeysApiUrl:                "https://keys-api.lido.fi",
			ProxyApiPort:                  proxyApiPort,
			DefaultAllowedExitDelay:       defaultAllowedExitDelay,
			ExitDelayMultiplier:           exitDelayMultiplier,
			SecondsPerSlot:                12,
			BlockChunkSize:                blockChunkSize,
			BlockScannerMinDistance:       blockScannerMinDistance,
		}
	default:
		logger.Fatal("Unknown network: %s", network)
	}

	logConfig(config)

	return config, nil
}

// logConfig iterates over all fields of Config to ensure nothing is missed
// when logging. This avoids having to maintain a manual list when new fields
// are added.
func logConfig(cfg Config) {
	v := reflect.ValueOf(cfg)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()
		logger.DebugWithPrefix("CONFIG", "Config.%s: %v", field.Name, value)
	}
}
