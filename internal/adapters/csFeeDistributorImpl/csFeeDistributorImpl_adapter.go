package csfeedistributorimpl

import (
	"context"
	"lido-events/internal/adapters/csFeeDistributorImpl/bindings"
	"lido-events/internal/application/domain"
	"log"

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
		return nil, err
	}

	return &CsFeeDistributorImplAdapter{
		client:                  client,
		CsFeeDistributorAddress: csFeeDistributorAddress,
	}, nil
}

// ScanDistributionLogUpdatedEvents scans the Vebo contract for DistributionLogUpdated events.
func (cs *CsFeeDistributorImplAdapter) ScanDistributionLogUpdatedEvents(ctx context.Context, start uint64, end *uint64, handleDistributionLogUpdated func(*domain.BindingsDistributionLogUpdated) error) error {
	// print something
	log.Printf("Scanning DistributionLogUpdated events from block %d to %d", start, *end)
	// print the address
	log.Printf("CsFeeDistributorAddress: %s", cs.CsFeeDistributorAddress.String())
	csFeeDistributorContract, err := bindings.NewBindings(cs.CsFeeDistributorAddress, cs.client)
	if err != nil {
		return err
	}

	// Retrieve DistributionLogUpdated events from the specified block range
	distributionLogUpdated, err := csFeeDistributorContract.FilterDistributionLogUpdated(&bind.FilterOpts{Context: ctx, Start: start, End: end})
	if err != nil {
		return err
	}

	for distributionLogUpdated.Next() {
		if distributionLogUpdated.Error() != nil {
			log.Printf("Error fetching DistributionLogUpdated event: %v", distributionLogUpdated.Error())
			continue
		}

		if err := handleDistributionLogUpdated(distributionLogUpdated.Event); err != nil {
			log.Printf("Error handling DistributionLogUpdated event: %v", err)
		}
	}

	return nil
}
