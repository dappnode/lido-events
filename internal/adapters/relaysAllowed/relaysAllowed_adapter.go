package relaysallowed

import (
	"context"
	"fmt"

	"lido-events/internal/adapters/relaysAllowed/bindings"
	"lido-events/internal/application/domain"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type RelaysAllowedAdapter struct {
	Client                         *ethclient.Client
	MevBoostRelaysAllowListAddress common.Address
	DappmanagerUrl                 string
	MevBoostDnpName                string
}

func NewRelaysAllowedAdapter(
	wsURL string,
	mevBoostRelaysAllowListAddress common.Address,
	dappmanagerUrl string,
	mevBoostDnpName string,
) (*RelaysAllowedAdapter, error) {
	client, err := ethclient.Dial(wsURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum client at %s: %w", wsURL, err)
	}

	return &RelaysAllowedAdapter{
		client,
		mevBoostRelaysAllowListAddress,
		dappmanagerUrl,
		mevBoostDnpName,
	}, nil
}

// getRelaysAllowList fetches relays allow list from the SC
func (mev *RelaysAllowedAdapter) GetRelaysAllowList(ctx context.Context) ([]domain.RelayAllowed, error) {
	relaysAllowListContract, err := bindings.NewBindings(mev.MevBoostRelaysAllowListAddress, mev.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to create MevBoostRelaysAllowList contract instance: %w", err)
	}

	relaysAllowList, err := relaysAllowListContract.GetRelays(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("failed to get relays allow list: %w", err)
	}

	return relaysAllowList, nil
}
