package config

import (
	"lido-events/internal/logger"
	"math/big"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

type Config struct {
	Network string

	DBDirectory        string
	PerformanceDBPath  string
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
	ApiPort            uint64
	LidoDnpName        string

	CORS []string

	// Individual contract addresses
	CSFeeDistributorProxyAddress  common.Address
	VEBOAddress                   common.Address
	MEVBoostRelaysAllowListAddres common.Address

	// Block number of the deployment of the VEBO contract
	VeboBlockDeployment uint64

	// tx receipts
	CSModuleTxReceipt common.Hash

	// Lido specifics
	LidoKeysApiUrl string
	ProxyApiPort   uint64

	// Blockchain
	MinGenesisTime          uint64
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
	performanceDbPath := filepath.Join(dbDirectory, "performance.db")

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
			PerformanceDBPath:             performanceDbPath,
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
			ApiPort:                       apiPort,
			LidoDnpName:                   "llido-csm-hoodi.dnp.dappnode.eth",
			CORS:                          parseCORS(corsEnv, []string{"http://ui.lido-csm-hoodi.dappnode", "http://my.dappnode"}),
			CSFeeDistributorProxyAddress:  common.HexToAddress("0xaCd9820b0A2229a82dc1A0770307ce5522FF3582"),
			VEBOAddress:                   common.HexToAddress("0x8664d394C2B3278F26A1B44B967aEf99707eeAB2"),
			MEVBoostRelaysAllowListAddres: common.HexToAddress("0x279d3A456212a1294DaEd0faEE98675a52E8A4Bf"),
			VeboBlockDeployment:           uint64(750),
			CSModuleTxReceipt:             common.HexToHash("0xebc45a0fa30a3f9badbcc4448ea22cef1a5d18b97825802a70df31cecb59127d"),
			LidoKeysApiUrl:                "https://keys-api-hoodi.testnet.fi",
			ProxyApiPort:                  proxyApiPort,
			MinGenesisTime:                uint64(1742277084), // timestamp when the CSFeeDistributorAddress SC was deployed:  https://hoodi.etherscan.io/tx/0xb5bf7cc66bc3b04eee6fb8ec650dd699d23b4ff53028fa5ec333d8e8fbe5201e
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
			PerformanceDBPath:             performanceDbPath,
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
			LidoDnpName:                   "llido-csm-mainnet.dnp.dappnode.eth",
			ApiPort:                       apiPort,
			CORS:                          parseCORS(corsEnv, []string{"http://ui.lido-csm-mainnet.dappnode", "http://my.dappnode"}),
			CSFeeDistributorProxyAddress:  common.HexToAddress("0xD99CC66fEC647E68294C6477B40fC7E0F6F618D0"),
			VEBOAddress:                   common.HexToAddress("0x0De4Ea0184c2ad0BacA7183356Aea5B8d5Bf5c6e"),
			MEVBoostRelaysAllowListAddres: common.HexToAddress("0xF95f069F9AD107938F6ba802a3da87892298610E"),
			VeboBlockDeployment:           uint64(17172556),
			CSModuleTxReceipt:             common.HexToHash("0xf5330dbcf09885ed145c4435e356b5d8a10054751bb8009d3a2605d476ac173f"),
			LidoKeysApiUrl:                "https://keys-api.lido.fi",
			ProxyApiPort:                  proxyApiPort,
			MinGenesisTime:                uint64(1606824023),
			BlockChunkSize:                blockChunkSize,
			BlockScannerMinDistance:       blockScannerMinDistance,
		}
	default:
		logger.Fatal("Unknown network: %s", network)
	}

	return config, nil
}
