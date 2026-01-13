//go:build integration
// +build integration

package relays_test

import (
	"context"
	"os"
	"testing"
	"time"

	"lido-events/internal/adapters/relays"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestRelaysAdapterIntegration(t *testing.T) {
	t.Skip("integration test temporarily disabled")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	rpcURL := os.Getenv("EXECUTION_RPC_URL")
	if rpcURL == "" {
		rpcURL = os.Getenv("RPC_URL")
	}
	if rpcURL == "" {
		t.Fatal("EXECUTION_RPC_URL (or RPC_URL) environment variable not set")
	}

	allowListAddrHex := os.Getenv("MEVBOOST_RELAYS_ALLOWLIST_ADDRESS")
	if allowListAddrHex == "" {
		t.Fatal("MEVBOOST_RELAYS_ALLOWLIST_ADDRESS environment variable not set")
	}
	allowListAddr := common.HexToAddress(allowListAddrHex)
	if allowListAddr == (common.Address{}) {
		t.Fatalf("MEVBOOST_RELAYS_ALLOWLIST_ADDRESS resolves to zero address: %q", allowListAddrHex)
	}

	dappmanagerURL := os.Getenv("DAPPMANAGER_URL")
	if dappmanagerURL == "" {
		t.Fatal("DAPPMANAGER_URL environment variable not set")
	}

	mevBoostDnpName := os.Getenv("MEVBOOST_DNP_NAME")
	if mevBoostDnpName == "" {
		t.Fatal("MEVBOOST_DNP_NAME environment variable not set")
	}

	rpcClient, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		t.Fatalf("failed to dial execution RPC %s: %v", rpcURL, err)
	}
	defer rpcClient.Close()

	adapter, err := relays.NewARelays(rpcClient, allowListAddr, dappmanagerURL, mevBoostDnpName)
	if err != nil {
		t.Fatalf("failed to create relays adapter: %v", err)
	}

	allowList, err := adapter.GetRelaysAllowList(ctx)
	if err != nil {
		t.Fatalf("GetRelaysAllowList failed: %v", err)
	}
	if len(allowList) == 0 {
		t.Fatalf("GetRelaysAllowList returned empty list")
	}
	t.Logf("allow list relays count: %d", len(allowList))
	for i, r := range allowList {
		if i >= 5 {
			break
		}
		t.Logf("allow list relay[%d]: %+v", i, r)
	}

	usedRelays, err := adapter.GetRelaysUsed(ctx)
	if err != nil {
		t.Fatalf("GetRelaysUsed failed: %v", err)
	}
	t.Logf("used relays count: %d", len(usedRelays))
	for i, r := range usedRelays {
		if i >= 10 {
			break
		}
		t.Logf("used relay[%d]: %s", i, r)
	}
}
