package vebo

import (
	"context"
	"lido-events/internal/adapters/vebo/bindings"
	"lido-events/internal/application/domain"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type VeboAdapter struct {
	client          *ethclient.Client
	VeboAddress     common.Address
	StakingModuleId []*big.Int
	NodeOperatorId  []*big.Int
	ValidatorIndex  []*big.Int // TODO: where to get it?
	RefSlot         []*big.Int // TODO: where to get it?
}

func NewVeboAdapter(
	wsURL string,
	veboAddress common.Address,
	stakingModuleId []*big.Int,
	nodeOperatorId []*big.Int,
	validatorIndex []*big.Int,
	refSlot []*big.Int,
) (*VeboAdapter, error) {
	client, err := ethclient.Dial(wsURL)
	if err != nil {
		return nil, err
	}

	return &VeboAdapter{
		client:          client,
		VeboAddress:     veboAddress,
		StakingModuleId: stakingModuleId,
		NodeOperatorId:  nodeOperatorId,
		ValidatorIndex:  validatorIndex,
		RefSlot:         refSlot,
	}, nil
}

// ScanVeboValidatorExitRequestEvent scans the Vebo contract for ValidatorExitRequest events.
func (va *VeboAdapter) ScanVeboValidatorExitRequestEvent(ctx context.Context, handleValidatorExitRequestEvent func(*domain.VeboValidatorExitRequest) error) error {
	veboContract, err := bindings.NewVebo(va.VeboAddress, va.client)
	if err != nil {
		return err
	}

	validatorExitRequestEvents, err := veboContract.FilterValidatorExitRequest(&bind.FilterOpts{Context: ctx}, va.StakingModuleId, va.NodeOperatorId, va.ValidatorIndex)
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
	veboContract, err := bindings.NewVebo(va.VeboAddress, va.client)
	if err != nil {
		return err
	}

	// Watch for ReportSubmitted
	reportSubmittedChan := make(chan *domain.VeboReportSubmitted)
	subReport, err := veboContract.WatchReportSubmitted(&bind.WatchOpts{Context: ctx}, reportSubmittedChan, va.RefSlot)
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
