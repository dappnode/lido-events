package config

import (
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
)

type Config struct {
	SignerUrl          string
	IpfsUrl            string
	WsURL              string
	CSMStakingModuleID *big.Int
	EtherscanURL       string
	BeaconchainURL     string
	CSMUIURL           string

	// Individual contract addresses
	CSAccountingAddress     common.Address
	CSFeeDistributorAddress common.Address
	VEBOAddress             common.Address
	CSModuleAddress         common.Address

	// Block number of the deployment of the Vebo contract
	VeboBlockDeployment uint64
}

func LoadNetworkConfig() (Config, error) {
	network := os.Getenv("NETWORK")
	wsURL := os.Getenv("WS_URL")
	ipfsUrl := os.Getenv("IPFS_URL")

	var config Config

	switch network {
	case "holesky":
		config = Config{
			SignerUrl:               "http://signer.holesky.dncore.dappnode",
			IpfsUrl:                 ipfsUrl,
			WsURL:                   wsURL,
			CSMStakingModuleID:      big.NewInt(4),
			EtherscanURL:            "https://holesky.etherscan.io",
			BeaconchainURL:          "https://holesky.beaconcha.in",
			CSMUIURL:                "https://csm.testnet.fi",
			CSAccountingAddress:     common.HexToAddress("0x4562c3e63c2e586cD1651B958C22F88135aCAd4f"),
			CSFeeDistributorAddress: common.HexToAddress("0xc093e53e8F4b55A223c18A2Da6fA00e60DD5EFE1"),
			VEBOAddress:             common.HexToAddress("0xffDDF7025410412deaa05E3E1cE68FE53208afcb"),
			VeboBlockDeployment:     uint64(30701),
			CSModuleAddress:         common.HexToAddress("common.HexToAddress(0x4562c3e63c2e586cD1651B958C22F88135aCAd4f"),
		}
	case "mainnet":
		config = Config{
			SignerUrl:               "http://signer.mainnet.dncore.dappnode",
			IpfsUrl:                 ipfsUrl,
			WsURL:                   wsURL,
			CSMStakingModuleID:      big.NewInt(3),
			EtherscanURL:            "https://etherscan.io",
			BeaconchainURL:          "https://beaconcha.in",
			CSMUIURL:                "https://csm.lido.fi",
			CSAccountingAddress:     common.HexToAddress("0xdA7dE2ECdDfccC6c3AF10108Db212ACBBf9EA83F"),
			CSFeeDistributorAddress: common.HexToAddress("0x4d72BFF1BeaC69925F8Bd12526a39BAAb069e5Da"),
			VEBOAddress:             common.HexToAddress("0x0De4Ea0184c2ad0BacA7183356Aea5B8d5Bf5c6e"),
			VeboBlockDeployment:     uint64(17172556),
			CSModuleAddress:         common.HexToAddress("0xdA7dE2ECdDfccC6c3AF10108Db212ACBBf9EA83F"),
		}
	default:
		log.Fatalf("Unknown network: %s", network)
	}

	return config, nil
}
