//go:build integration
// +build integration

package vebo

import (
	"context"
	"math/big"
	"os"
	"testing"
	"time"

	"lido-events/internal/application/domain"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
)

// Updated setupVeboAdapter function with real ethclient and hardcoded values
func setupVeboAdapter() (*VeboAdapter, error) {
	// Read wsURL from  env
	wsURL := os.Getenv("WS_URL")
	// Real WebSocket connection to Ethereum node
	client, err := ethclient.Dial(wsURL)
	if err != nil {
		return nil, err
	}

	veboAddress := common.HexToAddress("0x0De4Ea0184c2ad0BacA7183356Aea5B8d5Bf5c6e")
	stakingModuleId := []*big.Int{}
	nodeOperatorId := []*big.Int{big.NewInt(20)}
	validatorIndex := []*big.Int{}
	refSlot := []*big.Int{}

	return &VeboAdapter{
		client:          client,
		VeboAddress:     veboAddress,
		StakingModuleId: stakingModuleId,
		NodeOperatorId:  nodeOperatorId,
		ValidatorIndex:  validatorIndex,
		RefSlot:         refSlot,
	}, nil
}

// Test ScanVeboValidatorExitRequestEvent with real query
func TestScanVeboValidatorExitRequestEventIntegration(t *testing.T) {
	start := uint64(21071257)
	end := uint64(21071259)
	adapter, err := setupVeboAdapter()
	assert.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Use a map to track found events by ValidatorIndex
	foundEvents := map[string]struct {
		ValidatorPubkey string
		BlockNumber     uint64
	}{}

	// Scan real events from the blockchain
	err = adapter.ScanVeboValidatorExitRequestEvent(ctx, start, &end, func(event *domain.VeboValidatorExitRequest) error {
		// Store the event details in foundEvents for verification
		foundEvents[event.ValidatorIndex.String()] = struct {
			ValidatorPubkey string
			BlockNumber     uint64
		}{
			ValidatorPubkey: common.Bytes2Hex(event.ValidatorPubkey),
			BlockNumber:     event.Raw.BlockNumber,
		}

		// Print event details for debugging
		t.Logf("ValidatorIndex: %s, ValidatorPubkey: %s, BlockNumber: %d",
			event.ValidatorIndex.String(),
			common.Bytes2Hex(event.ValidatorPubkey),
			event.Raw.BlockNumber)

		return nil
	})

	// Check that the expected events are present
	assert.NoError(t, err)
	assert.Contains(t, foundEvents, "370637")
	assert.Equal(t, "b07c5cc5abd773d24c460140f872d24fb6584027626a4cba7b43e4f79f22c7b4846ea55d11ba96a322ba663c55cf3523", foundEvents["370637"].ValidatorPubkey)
	assert.Equal(t, uint64(21071258), foundEvents["370637"].BlockNumber)

	assert.Contains(t, foundEvents, "370638")
	assert.Equal(t, "92404310ad54447f6ab3277351a0da37f9b75696409056e855cfbe9838c649def1ca24cba994f6394f436c940d4cc3bc", foundEvents["370638"].ValidatorPubkey)
	assert.Equal(t, uint64(21071258), foundEvents["370638"].BlockNumber)
}
