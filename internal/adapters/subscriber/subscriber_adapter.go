package subscriber

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type SubscriberAdapter struct {
	client                  *ethclient.Client
	CsFeeDistributorAddress common.Address
	CsModuleAddress         common.Address
	VeboAddress             common.Address
	StakingModuleId         []*big.Int
	NodeOperatorId          []*big.Int
	ValidatorIndex          []*big.Int // TODO: where to get it?
	RefSlot                 []*big.Int // TODO: where to get it?
}

func NewSubscriberAdapter(
	wsURL string,
	csFeeDistributorAddress, csModuleAddress, veboAddress common.Address,
) (*SubscriberAdapter, error) {
	client, err := ethclient.Dial(wsURL)
	if err != nil {
		return nil, err
	}

	return &SubscriberAdapter{
		client:                  client,
		CsFeeDistributorAddress: csFeeDistributorAddress,
		CsModuleAddress:         csModuleAddress,
		VeboAddress:             veboAddress,
	}, nil
}
