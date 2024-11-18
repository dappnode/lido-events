package config

import (
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
)

// TODO: consider reading ENV OPERATOR_IDS as a string comma separated list and converting it to a slice of big.Int
// this enva could be either used to overwrite the default operator ids or to add new ones

type Config struct {
	SignerUrl          string
	IpfsUrl            string
	WsURL              string
	RpcUrl             string
	CSMStakingModuleID *big.Int
	EtherscanURL       string
	BeaconchainURL     string
	CSMUIURL           string

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
	network := os.Getenv("NETWORK")
	// default to holesky
	if network == "" {
		network = "holesky"
	}
	ipfsUrl := os.Getenv("IPFS_URL")
	// default to local http://ipfs.dappnode:5001
	if ipfsUrl == "" {
		ipfsUrl = "http://ipfs.dappnode:5001"
	}

	// Retrieve the WS_URL and RPC_URL from the environment variable
	wsURL := os.Getenv("WS_URL")
	rpcURL := os.Getenv("RPC_URL")

	// Retrieve the BEACONCHAIN_URL from the environment variable
	beaconchainURL := os.Getenv("BEACONCHAIN_URL")

	var config Config

	switch network {
	case "holesky":
		if wsURL == "" {
			wsURL = "ws://execution.holesky.dncore.dappnode:8546" // Default holesky WS URL
		}
		if rpcURL == "" {
			rpcURL = "http://execution.holesky.dncore.dappnode:8545" // Default holesky RPC URL
		}
		if beaconchainURL == "" {
			beaconchainURL = "http://beacon-chain.holesky.dncore.dappnode:3500" // Default holesky beaconchain URL
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
			CSAccountingAddress:             common.HexToAddress("0x4562c3e63c2e586cD1651B958C22F88135aCAd4f"),
			CSFeeDistributorAddress:         common.HexToAddress("0xD7ba648C8F72669C6aE649648B516ec03D07c8ED"),
			CSFeeDistributorImplAddress:     common.HexToAddress("0xe1863C61d2AF2899f06223152ebaaf993C29aEa7"), // https://holesky.etherscan.io/address/0xe1863c61d2af2899f06223152ebaaf993c29aea7#code
			VEBOAddress:                     common.HexToAddress("0xffDDF7025410412deaa05E3E1cE68FE53208afcb"),
			VeboBlockDeployment:             uint64(30701),
			CsFeeDistributorBlockDeployment: uint64(1774650),
			CSModuleAddress:                 common.HexToAddress("0x4562c3e63c2e586cD1651B958C22F88135aCAd4f"),
		}
	case "mainnet":
		if wsURL == "" {
			wsURL = "ws://execution.mainnet.dncore.dappnode:8546" // Default mainnet WS URL
		}
		if rpcURL == "" {
			rpcURL = "http://execution.mainnet.dncore.dappnode:8545" // Default mainnet RPC URL
		}
		if beaconchainURL == "" {
			beaconchainURL = "http://beacon-chain.mainnet.dncore.dappnode:3500" // Default mainnet beaconchain URL
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
			CSAccountingAddress:             common.HexToAddress("0xdA7dE2ECdDfccC6c3AF10108Db212ACBBf9EA83F"),
			CSFeeDistributorAddress:         common.HexToAddress("0xD99CC66fEC647E68294C6477B40fC7E0F6F618D0"),
			CSFeeDistributorImplAddress:     common.HexToAddress("0x17Fc610ecbbAc3f99751b3B2aAc1bA2b22E444f0 "), // https://etherscan.io/address/0x17Fc610ecbbAc3f99751b3B2aAc1bA2b22E444f0#code
			VEBOAddress:                     common.HexToAddress("0x0De4Ea0184c2ad0BacA7183356Aea5B8d5Bf5c6e"),
			VeboBlockDeployment:             uint64(17172556),
			CsFeeDistributorBlockDeployment: uint64(20935463),
			CSModuleAddress:                 common.HexToAddress("0xdA7dE2ECdDfccC6c3AF10108Db212ACBBf9EA83F"),
		}
	default:
		log.Fatalf("Unknown network: %s", network)
	}

	return config, nil
}
