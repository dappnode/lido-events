package relays

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"lido-events/internal/adapters/relays/bindings"
	"lido-events/internal/application/domain"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Adapter combines access to the relays allow list smart contract and
// the currently used relays configured in the MEV Boost DNP.
type Adapter struct {
	rpcClient                      *ethclient.Client
	mevBoostRelaysAllowListAddress common.Address
	dappmanagerURL                 string
	mevBoostDnpName                string
}

func NewAdapter(
	rpcClient *ethclient.Client,
	mevBoostRelaysAllowListAddress common.Address,
	dappmanagerURL string,
	mevBoostDnpName string,
) (*Adapter, error) {
	return &Adapter{
		rpcClient:                      rpcClient,
		mevBoostRelaysAllowListAddress: mevBoostRelaysAllowListAddress,
		dappmanagerURL:                 dappmanagerURL,
		mevBoostDnpName:                mevBoostDnpName,
	}, nil
}

// GetRelaysAllowList fetches the relays allow list from the smart contract.
func (a *Adapter) GetRelaysAllowList(ctx context.Context) ([]domain.RelayAllowed, error) {
	contract, err := bindings.NewBindings(a.mevBoostRelaysAllowListAddress, a.rpcClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create MevBoostRelaysAllowList contract instance: %w", err)
	}

	relaysAllowList, err := contract.GetRelays(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("failed to get relays allow list: %w", err)
	}

	return relaysAllowList, nil
}

// GetRelaysUsed fetches the RELAYS env from the MEV Boost DNP via Dappmanager.
func (a *Adapter) GetRelaysUsed(ctx context.Context) ([]string, error) {
	url := fmt.Sprintf("%s/env/%s?envName=RELAYS", a.dappmanagerURL, a.mevBoostDnpName)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create GET request for URL %s: %w", url, err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request to Dappmanager at %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response status: %d", resp.StatusCode)
	}

	var responseString string
	if err := json.NewDecoder(resp.Body).Decode(&responseString); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	if responseString == "" {
		return []string{}, nil
	}

	return splitAndTrim(responseString, ","), nil
}

// splitAndTrim splits a separated string and trims whitespace around each part.
func splitAndTrim(input, sep string) []string {
	parts := strings.Split(input, sep)
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}
