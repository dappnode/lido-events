//go:build integration
// +build integration

package csfeedistributor_test

import (
	"context"
	"os"
	"testing"
	"time"

	csfeedistributor "lido-events/internal/adapters/csFeeDistributor"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
)

// setupCsFeeDistributorAdapter initializes CsFeeDistributorAdapter using RPC_URL and NETWORK env vars.
func setupCsFeeDistributorAdapter(t *testing.T) (*csfeedistributor.CsFeeDistributorAdapter, error) {
	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		t.Fatal("RPC_URL environment variable not set")
	}

	network := os.Getenv("NETWORK")
	if network == "" {
		t.Fatal("NETWORK environment variable not set")
	}

	rpcClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}

	var csFeeDistributorAddress common.Address
	switch network {
	case "hoodi":
		csFeeDistributorAddress = common.HexToAddress("0xaCd9820b0A2229a82dc1A0770307ce5522FF3582")
	case "mainnet":
		csFeeDistributorAddress = common.HexToAddress("0xD99CC66fEC647E68294C6477B40fC7E0F6F618D0")
	default:
		t.Fatalf("unsupported NETWORK value: %s", network)
	}

	adapter, err := csfeedistributor.NewCsFeeDistributorAdapter(rpcClient, csFeeDistributorAddress)
	return adapter, err
}

// TestGetAllLogCidsIntegration verifies that GetAllLogCids runs without errors
// and prints the returned CIDs.
func TestGetAllLogCidsIntegration(t *testing.T) {
	adapter, err := setupCsFeeDistributorAdapter(t)
	assert.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	t.Log("starting GetAllLogCids against csFeeDistributor contract")
	logCids, err := adapter.GetAllLogCids(ctx)
	assert.NoError(t, err)
	t.Logf("GetAllLogCids returned %d CIDs", len(logCids))
	for _, cid := range logCids {
		t.Logf("Found historical log CID: %s", cid)
	}
}
