package vebo

import (
	"context"
	"lido-events/internal/adapters/vebo/bindings"
	"lido-events/internal/application/domain"
	"lido-events/internal/application/ports"
	"log"
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
		return nil, err
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
		return err
	}

	// Get operator ids
	operatorIds, err := va.StorageAdapter.GetOperatorIds()
	if err != nil {
		return err
	}

	validatorExitRequestEvents, err := veboContract.FilterValidatorExitRequest(&bind.FilterOpts{Context: ctx, Start: start, End: end}, []*big.Int{}, operatorIds, []*big.Int{})
	if err != nil {
		return err
	}

	for validatorExitRequestEvents.Next() {
		if validatorExitRequestEvents.Error() != nil {
			log.Printf("Error fetching ValidatorExitRequest event: %v", validatorExitRequestEvents.Error())
			continue
		}

		if err := handleValidatorExitRequestEvent(validatorExitRequestEvents.Event); err != nil {
			log.Printf("Error handling ValidatorExitRequest: %v", err)
		}
	}

	return nil
}

// ScanReportSubmittedEvents scans the Vebo contract for ReportSubmitted events.
func (va *VeboAdapter) ScanReportSubmittedEvents(ctx context.Context, start uint64, end *uint64, handleReportSubmittedEvent func(*domain.VeboReportSubmitted) error) error {
	veboContract, err := bindings.NewVebo(va.VeboAddress, va.Client)
	if err != nil {
		return err
	}

	// Retrieve ReportSubmitted events from the specified block range
	reportSubmittedEvents, err := veboContract.FilterReportSubmitted(&bind.FilterOpts{Context: ctx, Start: start, End: end}, []*big.Int{})
	if err != nil {
		return err
	}

	for reportSubmittedEvents.Next() {
		if reportSubmittedEvents.Error() != nil {
			log.Printf("Error fetching ReportSubmitted event: %v", reportSubmittedEvents.Error())
			continue
		}

		if err := handleReportSubmittedEvent(reportSubmittedEvents.Event); err != nil {
			log.Printf("Error handling ReportSubmitted event: %v", err)
		}
	}

	return nil
}
