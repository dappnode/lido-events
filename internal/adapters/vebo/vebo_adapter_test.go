// vebo/vebo_adapter_test.go
package vebo

import (
	"context"
	"math/big"
	"testing"
	"time"

	"lido-events/internal/application/domain"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
)

// Updated setupVeboAdapter function with real ethclient and hardcoded values
func setupVeboAdapter() (*VeboAdapter, error) {
	// Real WebSocket connection to Ethereum node
	client, err := ethclient.Dial("wss://holesky.infura.io/ws/v3/e6c920580178424bbdf6dde266bfb5bd")
	if err != nil {
		return nil, err
	}

	veboAddress := common.HexToAddress("0xffDDF7025410412deaa05E3E1cE68FE53208afcb")
	stakingModuleId := []*big.Int{big.NewInt(1)}
	nodeOperatorId := []*big.Int{big.NewInt(19)}
	validatorIndex := []*big.Int{big.NewInt(1686582), big.NewInt(1686586)}
	refSlot := []*big.Int{big.NewInt(1)}

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
func TestScanVeboValidatorExitRequestEvent(t *testing.T) {
	start := uint64(2597000)
	end := uint64(2597999)
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
	assert.Contains(t, foundEvents, "1686582")
	assert.Equal(t, "b6736f5a5b7df8c6a1e03d6c3225807a5fc6de1b244b40ce18e52d9faee9033f2aed0370bc3e0736231eb002b0c88fb9", foundEvents["1686582"].ValidatorPubkey)
	assert.Equal(t, uint64(2597903), foundEvents["1686582"].BlockNumber)

	assert.Contains(t, foundEvents, "1686586")
	assert.Equal(t, "b1266f71ef236b60a66b033ab3f5273850261a87b82cad2b1427f1023c7ff029b76688ee6ca0eb4beb3ce3875cbef3c5", foundEvents["1686586"].ValidatorPubkey)
	assert.Equal(t, uint64(2597903), foundEvents["1686586"].BlockNumber)
}
