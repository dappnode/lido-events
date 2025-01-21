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
	RpcClient                      *ethclient.Client
	MevBoostRelaysAllowListAddress common.Address
	DappmanagerUrl                 string
	MevBoostDnpName                string
}

func NewRelaysAllowedAdapter(
	rpcClient *ethclient.Client,
	mevBoostRelaysAllowListAddress common.Address,
	dappmanagerUrl string,
	mevBoostDnpName string,
) (*RelaysAllowedAdapter, error) {

	return &RelaysAllowedAdapter{
		rpcClient,
		mevBoostRelaysAllowListAddress,
		dappmanagerUrl,
		mevBoostDnpName,
	}, nil
}

// getRelaysAllowList fetches relays allow list from the SC
func (mev *RelaysAllowedAdapter) GetRelaysAllowList(ctx context.Context) ([]domain.RelayAllowed, error) {
	relaysAllowListContract, err := bindings.NewBindings(mev.MevBoostRelaysAllowListAddress, mev.RpcClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create MevBoostRelaysAllowList contract instance: %w", err)
	}

	relaysAllowList, err := relaysAllowListContract.GetRelays(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("failed to get relays allow list: %w", err)
	}

	return relaysAllowList, nil
}
