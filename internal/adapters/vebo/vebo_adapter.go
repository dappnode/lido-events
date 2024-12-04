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
	Client         *ethclient.Client
	VeboAddress    common.Address
	StorageAdapter ports.StoragePort
}

func NewVeboAdapter(
	wsURL string,
	veboAddress common.Address,
	storageAdapter ports.StoragePort,
) (*VeboAdapter, error) {
	client, err := ethclient.Dial(wsURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum client at %s: %w", wsURL, err)
	}

	return &VeboAdapter{
		client,
		veboAddress,
		storageAdapter,
	}, nil
}

// ScanVeboValidatorExitRequestEvent scans the Vebo contract for ValidatorExitRequest events.
func (va *VeboAdapter) ScanVeboValidatorExitRequestEvent(ctx context.Context, start uint64, end *uint64, handleValidatorExitRequestEvent func(*domain.VeboValidatorExitRequest) error) error {
	veboContract, err := bindings.NewVebo(va.VeboAddress, va.Client)
	if err != nil {
		return fmt.Errorf("failed to create Vebo contract instance: %w", err)
	}

	// Get operator ids
	operatorIds, err := va.StorageAdapter.GetOperatorIds()
	if err != nil {
		return fmt.Errorf("failed to retrieve operator IDs from storage: %w", err)
	}

	// Filter for ValidatorExitRequest events
	validatorExitRequestEvents, err := veboContract.FilterValidatorExitRequest(&bind.FilterOpts{Context: ctx, Start: start, End: end}, []*big.Int{}, operatorIds, []*big.Int{})
	if err != nil {
		return fmt.Errorf("failed to filter ValidatorExitRequest events: %w", err)
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

// WatchReportSubmittedEvents subscribes to Ethereum events and handles them.
func (va *VeboAdapter) WatchReportSubmittedEvents(ctx context.Context, handleReportSubmittedEvent func(*domain.VeboReportSubmitted) error) error {
	veboContract, err := bindings.NewVebo(va.VeboAddress, va.Client)
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
				return
			// case err := <-subReport.Err():
			// 	// Exit on subscription error
			// 	return
			case <-ctx.Done():
				return
			}
		}
	}()

	return nil
}
