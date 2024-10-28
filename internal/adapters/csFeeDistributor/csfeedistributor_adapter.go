package csfeedistributor

import (
	"context"
	"lido-events/internal/adapters/csFeeDistributor/bindings"
	"lido-events/internal/application/domain"
	"lido-events/internal/application/ports"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type CsFeeDistributorAdapter struct {
	client                  *ethclient.Client
	CsFeeDistributorAddress common.Address
}

func NewCsFeeDistributorAdapter(
	wsURL string,
	csFeeDistributorAddress common.Address,
) (*CsFeeDistributorAdapter, error) {
	client, err := ethclient.Dial(wsURL)
	if err != nil {
		return nil, err
	}

	return &CsFeeDistributorAdapter{
		client:                  client,
		CsFeeDistributorAddress: csFeeDistributorAddress,
	}, nil
}

func (csfa *CsFeeDistributorAdapter) WatchCsFeeDistributorEvents(ctx context.Context, handlers ports.CsFeeDistributorHandlers) error {
	csFeeDistributorContract, err := bindings.NewCsfeedistributor(csfa.CsFeeDistributorAddress, csfa.client)
	if err != nil {
		return err
	}

	distributionDataUpdatedChan := make(chan *domain.CsfeedistributorDistributionDataUpdated)
	sub, err := csFeeDistributorContract.WatchDistributionDataUpdated(&bind.WatchOpts{Context: ctx}, distributionDataUpdatedChan)
	if err != nil {
		return err
	}

	go func() {
		defer sub.Unsubscribe()
		for {
			select {
			case event := <-distributionDataUpdatedChan:
				if err := handlers.HandleDistributionDataUpdated(event); err != nil {
					log.Printf("Error handling log: %v", err)
				}
			case err := <-sub.Err():
				log.Printf("Subscription error: %v", err)
				return
			case <-ctx.Done():
				return
			}
		}
	}()

	return nil
}
