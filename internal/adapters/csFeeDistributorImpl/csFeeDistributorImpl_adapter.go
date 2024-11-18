package csfeedistributorimpl

import (
	"context"
	"fmt"
	"lido-events/internal/adapters/csFeeDistributorImpl/bindings"
	"lido-events/internal/application/domain"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type CsFeeDistributorImplAdapter struct {
	client                  *ethclient.Client
	CsFeeDistributorAddress common.Address
}

func NewCsFeeDistributorImplAdapter(
	wsURL string,
	csFeeDistributorAddress common.Address,
) (*CsFeeDistributorImplAdapter, error) {
	client, err := ethclient.Dial(wsURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum client at %s: %w", wsURL, err)
	}

	return &CsFeeDistributorImplAdapter{
		client:                  client,
		CsFeeDistributorAddress: csFeeDistributorAddress,
	}, nil
}

// ScanDistributionLogUpdatedEvents scans the CsFeeDistributor contract for DistributionLogUpdated events.
func (cs *CsFeeDistributorImplAdapter) ScanDistributionLogUpdatedEvents(ctx context.Context, start uint64, end *uint64, handleDistributionLogUpdated func(*domain.BindingsDistributionLogUpdated) error) error {
	csFeeDistributorContract, err := bindings.NewBindings(cs.CsFeeDistributorAddress, cs.client)
	if err != nil {
		return fmt.Errorf("failed to create CsFeeDistributor contract instance: %w", err)
	}

	// Retrieve DistributionLogUpdated events from the specified block range
	distributionLogUpdated, err := csFeeDistributorContract.FilterDistributionLogUpdated(&bind.FilterOpts{Context: ctx, Start: start, End: end})
	if err != nil {
		return fmt.Errorf("failed to filter DistributionLogUpdated events for block range %d to %v: %w", start, end, err)
	}

	for distributionLogUpdated.Next() {
		if err := distributionLogUpdated.Error(); err != nil {
			// Skip this event if there is an error retrieving it
			continue
		}

		if err := handleDistributionLogUpdated(distributionLogUpdated.Event); err != nil {
			// Continue to the next event if handling fails
			continue
		}
	}

	return nil
}
