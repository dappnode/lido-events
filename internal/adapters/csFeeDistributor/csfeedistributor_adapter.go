package csfeedistributor

import (
	"context"
	"fmt"
	"lido-events/internal/adapters/csFeeDistributor/bindings"
	"lido-events/internal/application/domain"

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
		return nil, fmt.Errorf("failed to connect to Ethereum client at %s: %w", wsURL, err)
	}

	return &CsFeeDistributorAdapter{
		client:                  client,
		CsFeeDistributorAddress: csFeeDistributorAddress,
	}, nil
}

// WatchCsFeeDistributorEvents watches for CsFeeDistributor events and calls the handleDistributionDataUpdated function when an event is received.
func (csfa *CsFeeDistributorAdapter) WatchCsFeeDistributorEvents(ctx context.Context, handleDistributionDataUpdated func(reportSubmitted *domain.CsfeedistributorDistributionDataUpdated) error) error {
	csFeeDistributorContract, err := bindings.NewCsfeedistributor(csfa.CsFeeDistributorAddress, csfa.client)
	if err != nil {
		return fmt.Errorf("failed to create CsFeeDistributor contract instance: %w", err)
	}

	distributionDataUpdatedChan := make(chan *domain.CsfeedistributorDistributionDataUpdated)
	sub, err := csFeeDistributorContract.WatchDistributionDataUpdated(&bind.WatchOpts{Context: ctx}, distributionDataUpdatedChan)
	if err != nil {
		return fmt.Errorf("failed to watch DistributionDataUpdated event: %w", err)
	}

	go func() {
		defer sub.Unsubscribe()
		for {
			select {
			case event := <-distributionDataUpdatedChan:
				handleDistributionDataUpdated(event)
				return
			// case err := <-sub.Err():
			// 	// Subscription error should be handled by returning it to the service layer.
			// 	return
			case <-ctx.Done():
				return
			}
		}
	}()

	return nil
}
