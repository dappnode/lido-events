package config

import (
	"lido-events/internal/logger"
	"math/big"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
)

type Config struct {
	SignerUrl          string
	IpfsUrl            string
	WsURL              string
	RpcUrl             string
	CSMStakingModuleID *big.Int
	EtherscanURL       string
	BeaconchainURL     string
	CSMUIURL           string
	ApiPort            uint64

	// Individual contract addresses
	CSAccountingAddress         common.Address
	CSFeeDistributorAddress     common.Address
	CSFeeDistributorImplAddress common.Address
	VEBOAddress                 common.Address
	CSModuleAddress             common.Address

	// Block number of the deployment of the VEBO contract and the CSFeeDistributor contract
	VeboBlockDeployment             uint64
	CsFeeDistributorBlockDeployment uint64
}

func LoadNetworkConfig() (Config, error) {
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
	network := os.Getenv("NETWORK")
	// Default to holesky
	if network == "" {
		network = "holesky"
	}

	ipfsUrl := os.Getenv("IPFS_URL")
	// Default to local http://ipfs.dappnode:5001
	if ipfsUrl == "" {
		ipfsUrl = "http://ipfs.dappnode:5001"
	}

	// Retrieve other necessary environment variables
	wsURL := os.Getenv("WS_URL")
	rpcURL := os.Getenv("RPC_URL")
	beaconchainURL := os.Getenv("BEACONCHAIN_URL")
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "INFO" // Default log level
	}

	var config Config

	switch network {
	case "holesky":
		// Configure default values for the holesky network
		if wsURL == "" {
			wsURL = "ws://execution.holesky.dncore.dappnode:8546"
		}
		if rpcURL == "" {
			rpcURL = "http://execution.holesky.dncore.dappnode:8545"
		}
		if beaconchainURL == "" {
			beaconchainURL = "http://beacon-chain.holesky.dncore.dappnode:3500"
		}
		config = Config{
			SignerUrl:                       "http://signer.holesky.dncore.dappnode",
			IpfsUrl:                         ipfsUrl,
			WsURL:                           wsURL,
			RpcUrl:                          rpcURL,
			CSMStakingModuleID:              big.NewInt(4),
			EtherscanURL:                    "https://holesky.etherscan.io",
			BeaconchainURL:                  beaconchainURL,
			CSMUIURL:                        "https://csm.testnet.fi",
			ApiPort:                         apiPort,
			CSAccountingAddress:             common.HexToAddress("0x4562c3e63c2e586cD1651B958C22F88135aCAd4f"),
			CSFeeDistributorAddress:         common.HexToAddress("0xD7ba648C8F72669C6aE649648B516ec03D07c8ED"),
			CSFeeDistributorImplAddress:     common.HexToAddress("0xe1863C61d2AF2899f06223152ebaaf993C29aEa7"),
			VEBOAddress:                     common.HexToAddress("0xffDDF7025410412deaa05E3E1cE68FE53208afcb"),
			VeboBlockDeployment:             uint64(30701),
			CsFeeDistributorBlockDeployment: uint64(1774650),
			CSModuleAddress:                 common.HexToAddress("0x4562c3e63c2e586cD1651B958C22F88135aCAd4f"),
		}
	case "mainnet":
		// Configure default values for the mainnet
		if wsURL == "" {
			wsURL = "ws://execution.mainnet.dncore.dappnode:8546"
		}
		if rpcURL == "" {
			rpcURL = "http://execution.mainnet.dncore.dappnode:8545"
		}
		if beaconchainURL == "" {
			beaconchainURL = "http://beacon-chain.mainnet.dncore.dappnode:3500"
		}
		config = Config{
			SignerUrl:                       "http://signer.mainnet.dncore.dappnode",
			IpfsUrl:                         ipfsUrl,
			WsURL:                           wsURL,
			RpcUrl:                          rpcURL,
			CSMStakingModuleID:              big.NewInt(3),
			EtherscanURL:                    "https://etherscan.io",
			BeaconchainURL:                  "https://beaconcha.in",
			CSMUIURL:                        "https://csm.lido.fi",
			ApiPort:                         apiPort,
			CSAccountingAddress:             common.HexToAddress("0xdA7dE2ECdDfccC6c3AF10108Db212ACBBf9EA83F"),
			CSFeeDistributorAddress:         common.HexToAddress("0xD99CC66fEC647E68294C6477B40fC7E0F6F618D0"),
			CSFeeDistributorImplAddress:     common.HexToAddress("0x17Fc610ecbbAc3f99751b3B2aAc1bA2b22E444f0"),
			VEBOAddress:                     common.HexToAddress("0x0De4Ea0184c2ad0BacA7183356Aea5B8d5Bf5c6e"),
			VeboBlockDeployment:             uint64(17172556),
			CsFeeDistributorBlockDeployment: uint64(20935463),
			CSModuleAddress:                 common.HexToAddress("0xdA7dE2ECdDfccC6c3AF10108Db212ACBBf9EA83F"),
		}
	default:
		logger.Fatal("Unknown network: %s", network)
	}

	return config, nil
}
