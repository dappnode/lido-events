package vebo

import (
	"context"
	"fmt"
	"lido-events/internal/adapters/vebo/bindings"
	"lido-events/internal/application/domain"
	"lido-events/internal/application/ports"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type VeboAdapter struct {
	WsClient       *ethclient.Client
	RpcClient      *ethclient.Client
	VeboAddress    common.Address
	StorageAdapter ports.StoragePort
	blockChunkSize uint64
}

func NewVeboAdapter(
	wsClient *ethclient.Client,
	rpcClient *ethclient.Client,
	veboAddress common.Address,
	storageAdapter ports.StoragePort,
	blockChunkSize uint64,
) (*VeboAdapter, error) {
	return &VeboAdapter{
		WsClient:       wsClient,
		RpcClient:      rpcClient,
		VeboAddress:    veboAddress,
		StorageAdapter: storageAdapter,
		blockChunkSize: blockChunkSize,
	}, nil
}

// ScanVeboValidatorExitRequestEvent scans the Vebo contract for ValidatorExitRequest events in chunks.
func (va *VeboAdapter) ScanVeboValidatorExitRequestEvent(
	ctx context.Context,
	start uint64,
	end *uint64,
	handleValidatorExitRequestEvent func(*domain.VeboValidatorExitRequest) error,
) error {
	if end == nil {
		return fmt.Errorf("end block cannot be nil")
	}

	veboContract, err := bindings.NewVebo(va.VeboAddress, va.RpcClient)
	if err != nil {
		return fmt.Errorf("failed to create Vebo contract instance: %w", err)
	}

	// Get operator IDs
	operatorIds, err := va.StorageAdapter.GetOperatorIds()
	if err != nil {
		return fmt.Errorf("failed to retrieve operator IDs from storage: %w", err)
	}

	// Helper function to scan a chunk
	scanChunk := func(chunkStart, chunkEnd uint64) error {
		validatorExitRequestEvents, err := veboContract.FilterValidatorExitRequest(
			&bind.FilterOpts{Context: ctx, Start: chunkStart, End: &chunkEnd},
			[]*big.Int{}, // No filter for `stakingModuleId`
			operatorIds,  // Operator IDs filter
			[]*big.Int{}, // No filter for `validatorIndex`
		)
		if err != nil {
			return fmt.Errorf("failed to filter ValidatorExitRequest events for block range %d to %d: %w", chunkStart, chunkEnd, err)
		}

		for validatorExitRequestEvents.Next() {
			if err := validatorExitRequestEvents.Error(); err != nil {
				return fmt.Errorf("error reading ValidatorExitRequest event: %w", err)
			}

			if err := handleValidatorExitRequestEvent(validatorExitRequestEvents.Event); err != nil {
				return fmt.Errorf("failed to handle ValidatorExitRequest event: %w", err)
			}
		}

		return nil
	}

	// Iterate through block ranges in chunks
	for currentStart := start; currentStart <= *end; currentStart += va.blockChunkSize {
		// Determine the end of the current chunk
		currentEnd := currentStart + va.blockChunkSize - 1
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

// WatchReportSubmittedEvents subscribes to Ethereum events and handles them.
func (va *VeboAdapter) WatchReportSubmittedEvents(ctx context.Context, handleReportSubmittedEvent func(*domain.VeboReportSubmitted) error) error {
	veboContract, err := bindings.NewVebo(va.VeboAddress, va.WsClient)
	if err != nil {
		return fmt.Errorf("failed to create Vebo contract instance: %w", err)
	}

	// Watch for ReportSubmitted events
	reportSubmittedChan := make(chan *domain.VeboReportSubmitted)
	subReport, err := veboContract.WatchReportSubmitted(&bind.WatchOpts{Context: ctx}, reportSubmittedChan, []*big.Int{})
	if err != nil {
		return fmt.Errorf("failed to subscribe to ReportSubmitted events: %w", err)
	}

	go func() {
		defer subReport.Unsubscribe()
		for {
			select {
			case event := <-reportSubmittedChan:
				handleReportSubmittedEvent(event)
			case <-ctx.Done():
				return
			}
		}
	}()

	return nil
}
