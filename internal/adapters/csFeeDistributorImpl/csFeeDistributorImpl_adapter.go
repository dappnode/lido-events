package csfeedistributorimpl

import (
	"context"
	"fmt"
	"lido-events/internal/adapters/csFeeDistributorImpl/bindings"
	"lido-events/internal/application/domain"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type CsFeeDistributorImplAdapter struct {
	rpcClient               *ethclient.Client
	CsFeeDistributorAddress common.Address
	blockChunkSize          uint64 // Configurable block chunk size
}

func NewCsFeeDistributorImplAdapter(
	rpcClient *ethclient.Client,
	csFeeDistributorAddress common.Address,
	blockChunkSize uint64,
) (*CsFeeDistributorImplAdapter, error) {
	return &CsFeeDistributorImplAdapter{
		rpcClient:               rpcClient,
		CsFeeDistributorAddress: csFeeDistributorAddress,
		blockChunkSize:          blockChunkSize,
	}, nil
}

// ScanDistributionLogUpdatedEvents scans the CsFeeDistributor contract for DistributionLogUpdated events in chunks.
func (cs *CsFeeDistributorImplAdapter) ScanDistributionLogUpdatedEvents(
	ctx context.Context,
	start uint64,
	end *uint64,
	operatorId *big.Int,
	handleDistributionLogUpdated func(*domain.BindingsDistributionLogUpdated, *big.Int) error,
) error {
	if end == nil {
		return fmt.Errorf("end block cannot be nil")
	}

	csFeeDistributorContract, err := bindings.NewBindings(cs.CsFeeDistributorAddress, cs.rpcClient)
	if err != nil {
		return fmt.Errorf("failed to create CsFeeDistributor contract instance: %w", err)
	}

	// Helper function to scan a chunk
	scanChunk := func(chunkStart, chunkEnd uint64) error {
		distributionLogUpdated, err := csFeeDistributorContract.FilterDistributionLogUpdated(
			&bind.FilterOpts{Context: ctx, Start: chunkStart, End: &chunkEnd},
		)
		if err != nil {
			return fmt.Errorf("failed to filter DistributionLogUpdated events for block range %d to %d: %w", chunkStart, chunkEnd, err)
		}

		for distributionLogUpdated.Next() {
			if err := distributionLogUpdated.Error(); err != nil {
				return fmt.Errorf("error reading DistributionLogUpdated event: %w", err)
			}

			if err := handleDistributionLogUpdated(distributionLogUpdated.Event, operatorId); err != nil {
				return fmt.Errorf("failed to handle DistributionLogUpdated event: %w", err)
			}
		}

		return nil
	}

	// Iterate through block ranges in chunks
	for currentStart := start; currentStart <= *end; currentStart += cs.blockChunkSize {
		// Determine the end of the current chunk
		currentEnd := currentStart + cs.blockChunkSize - 1
		if currentEnd > *end {
			currentEnd = *end
		}

		// Scan the current chunk
		if err := scanChunk(currentStart, currentEnd); err != nil {
			return fmt.Errorf("error scanning block range %d to %d: %w", currentStart, currentEnd, err)
		}
	}

	return nil
}
