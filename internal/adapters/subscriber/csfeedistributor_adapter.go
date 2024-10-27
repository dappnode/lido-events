package subscriber

import (
	"context"
	"lido-events/internal/adapters/subscriber/bindings/csfeedistributor"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func (sa *SubscriberAdapter) WatchCsFeeDistributorEvents(ctx context.Context, handleLog func(logData interface{}) error) error {
	csFeeDistributorContract, err := csfeedistributor.NewCsfeedistributor(sa.CsFeeDistributorAddress, sa.client)
	if err != nil {
		return err
	}

	distributionDataUpdatedChan := make(chan *csfeedistributor.CsfeedistributorDistributionDataUpdated)
	sub, err := csFeeDistributorContract.WatchDistributionDataUpdated(&bind.WatchOpts{Context: ctx}, distributionDataUpdatedChan)
	if err != nil {
		return err
	}

	go func() {
		defer sub.Unsubscribe()
		for {
			select {
			case event := <-distributionDataUpdatedChan:
				if err := handleLog(event); err != nil {
					log.Printf("Error handling log: %v", err)
				}
			case err := <-sub.Err():
				log.Printf("Subscription error: %v", err)
				return
			case <-ctx.Done():
				return
			}
		}
	}()

	return nil
}
