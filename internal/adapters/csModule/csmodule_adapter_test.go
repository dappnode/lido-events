//go:build integration
// +build integration

package csmodule_test

import (
	"context"
	"math/big"
	"os"
	"testing"
	"time"

	csmodule "lido-events/internal/adapters/csModule"
	"lido-events/internal/application/domain"
	"lido-events/internal/mocks"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func setupCsModuleAdapter(t *testing.T) (*csmodule.CsModuleAdapter, *mocks.MockStoragePort, error) {
	wsURL := os.Getenv("WS_URL")
	if wsURL == "" {
		t.Fatal("WS_URL environment variable not set")
	}

	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		t.Fatal("RPC_URL environment variable not set")
	}

	// Create the mock StoragePort
	mockStorage := new(mocks.MockStoragePort)

	// Define initial operator IDs as required by the test
	mockStorage.On("GetOperatorIds").Return([]*big.Int{
		big.NewInt(306),
		big.NewInt(283),
		big.NewInt(262),
	}, nil)

	// Define the behavior for RegisterOperatorIdListener
	operatorIdChan := make(chan []*big.Int, 1)
	mockStorage.On("RegisterOperatorIdListener").Return(operatorIdChan)

	csModuleAddress := common.HexToAddress("0xdA7dE2ECdDfccC6c3AF10108Db212ACBBf9EA83F")
	blockChunkSize := uint64(10000)

	// Initialize the adapter with the mock storage
	adapter, err := csmodule.NewCsModuleAdapter(wsURL, rpcURL, csModuleAddress, mockStorage, blockChunkSize)
	return adapter, mockStorage, err
}

// TestScanNodeOperatorEventsIntegration tests scanning for NodeOperator events with mocked data
func TestScanNodeOperatorEventsIntegration(t *testing.T) {
	t.Skip()
	adapter, mockStorage, err := setupCsModuleAdapter(t)
	assert.NoError(t, err)

	// Set the start and end blocks for the scan
	start := uint64(20935463)
	end := uint64(21657327)

	ctx, cancel := context.WithTimeout(context.Background(), 260*time.Second)
	defer cancel()

	// Map to store and verify events found during the scan
	foundEvents := struct {
		NodeOperatorAdded                 map[string]struct{}
		NodeOperatorManagerAddressChanged map[string]struct{}
		NodeOperatorRewardAddressChanged  map[string]struct{}
	}{
		NodeOperatorAdded:                 make(map[string]struct{}),
		NodeOperatorManagerAddressChanged: make(map[string]struct{}),
		NodeOperatorRewardAddressChanged:  make(map[string]struct{}),
	}

	// Execute the scan and handle each found event
	t.Log("Scanning for NodeOperator events...")
	// print start and end
	t.Logf("Start block: %d, End block: %d", start, end)
	err = adapter.ScanNodeOperatorEvents(ctx, common.HexToAddress("0x18a2869554df268828bc366fb5333df89c9c2b7b"), start, &end,
		func(event *domain.CsmoduleNodeOperatorAdded, address common.Address) error {
			foundEvents.NodeOperatorAdded[event.NodeOperatorId.String()] = struct{}{}
			t.Logf("NodeOperatorAdded: NodeOperatorId=%s, ManagerAddress=%s, RewardAddress=%s, BlockNumber=%d",
				event.NodeOperatorId.String(),
				event.ManagerAddress.Hex(),
				event.RewardAddress.Hex(),
				event.Raw.BlockNumber,
			)
			return nil
		},
		func(event *domain.CsmoduleNodeOperatorManagerAddressChanged, address common.Address) error {
			foundEvents.NodeOperatorManagerAddressChanged[event.NodeOperatorId.String()] = struct{}{}
			t.Logf("NodeOperatorManagerAddressChanged: NodeOperatorId=%s, OldAddress=%s, NewAddress=%s, BlockNumber=%d",
				event.NodeOperatorId.String(),
				event.OldAddress.Hex(),
				event.NewAddress.Hex(),
				event.Raw.BlockNumber,
			)
			return nil
		},
		func(event *domain.CsmoduleNodeOperatorRewardAddressChanged, address common.Address) error {
			foundEvents.NodeOperatorRewardAddressChanged[event.NodeOperatorId.String()] = struct{}{}
			t.Logf("NodeOperatorRewardAddressChanged: NodeOperatorId=%s, OldAddress=%s, NewAddress=%s, BlockNumber=%d",
				event.NodeOperatorId.String(),
				event.OldAddress.Hex(),
				event.NewAddress.Hex(),
				event.Raw.BlockNumber,
			)
			return nil
		},
	)
	t.Log("Scan complete")

	// print the errr
	t.Log(err)

	// Assertions for the expected events
	assert.NoError(t, err)

	// Ensure all expected mock calls were made
	mockStorage.AssertCalled(t, "GetOperatorIds")
}
