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
	rpcClient               *ethclient.Client
	CsFeeDistributorAddress common.Address
	BlockChunkSize          uint64 // Configurable block chunk size
}

func NewCsFeeDistributorImplAdapter(
	rpcClient *ethclient.Client,
	csFeeDistributorAddress common.Address,
	blockChunkSize uint64,
) (*CsFeeDistributorImplAdapter, error) {

	return &CsFeeDistributorImplAdapter{
		rpcClient:               rpcClient,
		CsFeeDistributorAddress: csFeeDistributorAddress,
		BlockChunkSize:          blockChunkSize,
	}, nil
}

// ScanDistributionLogUpdatedEvents scans the CsFeeDistributor contract for DistributionLogUpdated events in chunks.
func (cs *CsFeeDistributorImplAdapter) ScanDistributionLogUpdatedEvents(
	ctx context.Context,
	start uint64,
	end *uint64,
	handleDistributionLogUpdated func(*domain.BindingsDistributionLogUpdated) error,
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

			if err := handleDistributionLogUpdated(distributionLogUpdated.Event); err != nil {
				return fmt.Errorf("failed to handle DistributionLogUpdated event: %w", err)
			}
		}

		return nil
	}

	// Iterate through block ranges in chunks
	for current := start; current <= *end; current += cs.BlockChunkSize {
		chunkEnd := current + cs.BlockChunkSize - 1
		if chunkEnd > *end {
			chunkEnd = *end
		}

		// Ensure chunkEnd is not less than current
		if chunkEnd <= current {
			break // Exit the loop to avoid invalid block ranges
		}

		if err := scanChunk(current, chunkEnd); err != nil {
			return fmt.Errorf("error scanning block range %d to %d: %w", current, chunkEnd, err)
		}
	}

	return nil
}
