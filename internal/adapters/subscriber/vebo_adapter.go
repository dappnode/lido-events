package subscriber

import (
	"context"
	"lido-events/internal/adapters/subscriber/bindings/vebo"
	"lido-events/internal/aplication/ports"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func (sa *SubscriberAdapter) WatchVeboEvents(ctx context.Context, handlers ports.VeboHandlers) error {
	veboContract, err := vebo.NewVebo(sa.VeboAddress, sa.client)
	if err != nil {
		return err
	}

	// Watch for ValidatorExitRequest
	validatorExitRequestChan := make(chan *vebo.VeboValidatorExitRequest)
	subExit, err := veboContract.WatchValidatorExitRequest(&bind.WatchOpts{Context: ctx}, validatorExitRequestChan, sa.StakingModuleId, sa.NodeOperatorId, sa.ValidatorIndex)
	if err != nil {
		return err
	}

	// Watch for ReportSubmitted
	reportSubmittedChan := make(chan *vebo.VeboReportSubmitted)
	subReport, err := veboContract.WatchReportSubmitted(&bind.WatchOpts{Context: ctx}, reportSubmittedChan, sa.RefSlot)
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
