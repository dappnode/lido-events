package csfeeoracle

import (
	"context"
	"fmt"
	"lido-events/internal/adapters/csFeeOracle/bindings"
	"lido-events/internal/application/domain"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// CsFeeOracleAdapter holds the RPC client, contract address, and block chunk size.
type CsFeeOracleAdapter struct {
	rpcClient          *ethclient.Client
	CsFeeOracleAddress common.Address
	blockChunkSize     uint64 // Configurable block chunk size
}

// NewCsFeeOracleAdapter creates a new instance of CsFeeOracleAdapter.
func NewCsFeeOracleAdapter(
	rpcClient *ethclient.Client,
	csFeeOracleAddress common.Address,
	blockChunkSize uint64,
) (*CsFeeOracleAdapter, error) {
	return &CsFeeOracleAdapter{
		rpcClient:          rpcClient,
		CsFeeOracleAddress: csFeeOracleAddress,
		blockChunkSize:     blockChunkSize,
	}, nil
}

// ScanProcessingStartedEvents scans the csFeeOracle contract for ProcessingStarted events in chunks.
// The handleProcessingStarted callback is called for each event found.
func (oracle *CsFeeOracleAdapter) ScanProcessingStartedEvents(
	ctx context.Context,
	start uint64,
	end *uint64,
	handleProcessingStarted func(*domain.BindingsProcessingStarted, *big.Int) error,
) error {
	if end == nil {
		return fmt.Errorf("end block cannot be nil")
	}

	// Create a new instance of the contract using the generated bindings.
	csFeeOracleContract, err := bindings.NewBindings(oracle.CsFeeOracleAddress, oracle.rpcClient)
	if err != nil {
		return fmt.Errorf("failed to create csFeeOracle contract instance: %w", err)
	}

	// scanChunk is a helper that scans a range of blocks.
	scanChunk := func(chunkStart, chunkEnd uint64) error {
		// Filter for ProcessingStarted events within the given block range.
		iter, err := csFeeOracleContract.FilterProcessingStarted(&bind.FilterOpts{
			Context: ctx,
			Start:   chunkStart,
			End:     &chunkEnd,
		}, nil)
		if err != nil {
			return fmt.Errorf("failed to filter ProcessingStarted events for block range %d to %d: %w", chunkStart, chunkEnd, err)
		}

		// Iterate through the event log results.
		for iter.Next() {
			if err := iter.Error(); err != nil {
				return fmt.Errorf("error reading ProcessingStarted event: %w", err)
			}
			event := iter.Event
			// The handler is passed the event and the block number (here using the chunk end as a reference).
			if err := handleProcessingStarted(event, big.NewInt(int64(chunkEnd))); err != nil {
				return fmt.Errorf("failed to handle ProcessingStarted event: %w", err)
			}
		}

		return nil
	}

	// Loop through the block range in chunks.
	for currentStart := start; currentStart <= *end; currentStart += oracle.blockChunkSize {
		currentEnd := currentStart + oracle.blockChunkSize - 1
		if currentEnd > *end {
			currentEnd = *end
		}

		if err := scanChunk(currentStart, currentEnd); err != nil {
			return fmt.Errorf("error scanning block range %d to %d: %w", currentStart, currentEnd, err)
		}
	}

	return nil
}
