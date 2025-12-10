//go:build integration
// +build integration

package csparameters_test

import (
	"context"
	"os"
	"testing"
	"time"

	"lido-events/internal/adapters/csparameters"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
)

// setupCsParametersAdapter initializes CsParameters adapter using RPC_URL and CSPARAMETERS_ADDRESS env vars.
func setupCsParametersAdapter(t *testing.T) (*csparameters.CsParameters, error) {
	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		t.Fatal("RPC_URL environment variable not set")
	}

	csParamsAddressHex := os.Getenv("CSPARAMETERS_ADDRESS")
	if csParamsAddressHex == "" {
		t.Fatal("CSPARAMETERS_ADDRESS environment variable not set")
	}

	rpcClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}

	csParamsAddress := common.HexToAddress(csParamsAddressHex)
	adapter, err := csparameters.NewCsParameters(rpcClient, csParamsAddress)
	return adapter, err
}

// TestGetDefaultAllowedExitDelayIntegration verifies that GetDefaultAllowedExitDelay runs without errors
// and logs the returned value.
func TestGetDefaultAllowedExitDelayIntegration(t *testing.T) {
	adapter, err := setupCsParametersAdapter(t)
	assert.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	t.Log("starting GetDefaultAllowedExitDelay against csParameters contract")
	delay, err := adapter.GetDefaultAllowedExitDelay(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, delay)
	assert.True(t, delay.Sign() >= 0)

	t.Logf("GetDefaultAllowedExitDelay returned: %s", delay.String())
}