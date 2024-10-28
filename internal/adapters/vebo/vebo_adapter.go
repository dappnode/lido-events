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

func (va *VeboAdapter) WatchVeboEvents(ctx context.Context, handlers ports.VeboHandlers) error {
	veboContract, err := bindings.NewVebo(va.VeboAddress, va.client)
	if err != nil {
		return err
	}

	// Watch for ValidatorExitRequest
	validatorExitRequestChan := make(chan *domain.VeboValidatorExitRequest)
	subExit, err := veboContract.WatchValidatorExitRequest(&bind.WatchOpts{Context: ctx}, validatorExitRequestChan, va.StakingModuleId, va.NodeOperatorId, va.ValidatorIndex)
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
		defer subExit.Unsubscribe()
		defer subReport.Unsubscribe()
		for {
			select {
			case event := <-validatorExitRequestChan:
				if err := handlers.HandleValidatorExitRequestEvent(event); err != nil {
					log.Printf("Error handling ValidatorExitRequest: %v", err)
				}
			case event := <-reportSubmittedChan:
				if err := handlers.HandleReportSubmittedEvent(event); err != nil {
					log.Printf("Error handling ReportSubmitted: %v", err)
				}
			case err := <-subExit.Err():
				log.Printf("Subscription error (ValidatorExitRequest): %v", err)
				return
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
