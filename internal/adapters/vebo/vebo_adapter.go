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

	// print operator ids
	for _, operatorId := range operatorIds {
		log.Printf("Operator ID: %v", operatorId)
	}

	validatorExitRequestEvents, err := veboContract.FilterValidatorExitRequest(&bind.FilterOpts{Context: ctx, Start: start, End: end}, []*big.Int{}, operatorIds, []*big.Int{})
	if err != nil {
		return err
	}

	for validatorExitRequestEvents.Next() {
		if err := handleValidatorExitRequestEvent(&domain.VeboValidatorExitRequest{
			ValidatorIndex:  validatorExitRequestEvents.Event.ValidatorIndex,
			ValidatorPubkey: validatorExitRequestEvents.Event.ValidatorPubkey,
			Raw:             validatorExitRequestEvents.Event.Raw,
		}); err != nil {
			log.Printf("Error handling ValidatorExitRequest: %v", err)
		}
	}

	return nil
}

// WatchReportSubmittedEvents subscribes to Ethereum events and handles them.
func (va *VeboAdapter) WatchReportSubmittedEvents(ctx context.Context, handleReportSubmittedEvent func(*domain.VeboReportSubmitted) error) error {
	veboContract, err := bindings.NewVebo(va.VeboAddress, va.Client)
	if err != nil {
		return err
	}

	// Watch for ReportSubmitted
	reportSubmittedChan := make(chan *domain.VeboReportSubmitted)
	subReport, err := veboContract.WatchReportSubmitted(&bind.WatchOpts{Context: ctx}, reportSubmittedChan, []*big.Int{})
	if err != nil {
		return err
	}

	go func() {
		defer subReport.Unsubscribe()
		for {
			select {
			case event := <-reportSubmittedChan:
				if err := handleReportSubmittedEvent(event); err != nil {
					log.Printf("Error handling ReportSubmitted: %v", err)
				}
			case err := <-subReport.Err():
				log.Printf("Subscription error (ReportSubmitted): %v", err)
				return
			case <-ctx.Done():
				return
			}
		}
	}()

	return nil
}
