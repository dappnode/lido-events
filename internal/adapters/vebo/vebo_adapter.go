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

// ScanVeboValidatorExitRequestEvent scans the Vebo contract for ValidatorExitRequest events. error logs must be printed since its not executed from main
func (va *VeboAdapter) ScanVeboValidatorExitRequestEvent(ctx context.Context, start uint64, end *uint64, handleValidatorExitRequestEvent func(*domain.VeboValidatorExitRequest) error) error {
	log.Printf("Scanning ValidatorExitRequest events from block %d to %d", start, end)

	veboContract, err := bindings.NewVebo(va.VeboAddress, va.Client)
	if err != nil {
		log.Printf("Error creating Vebo contract: %v", err)
		return err
	}

	// Get operator ids
	operatorIds, err := va.StorageAdapter.GetOperatorIds()
	if err != nil {
		return err
	}

	validatorExitRequestEvents, err := veboContract.FilterValidatorExitRequest(&bind.FilterOpts{Context: ctx, Start: start, End: end}, []*big.Int{}, operatorIds, []*big.Int{})
	if err != nil {
		log.Printf("Error filtering ValidatorExitRequest events: %v", err)
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

// WatchReportSubmittedEvents subscribes to Ethereum events and handles them. not required to print error logs since will be printed when initializing in main
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
