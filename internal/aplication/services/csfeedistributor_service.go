package services

import (
	"context"
	"fmt"
	"lido-events/internal/aplication/domain"
	"lido-events/internal/aplication/ports"
	"log"
)

type CsFeeDistributorService struct {
	notifierPort         ports.NotifierPort
	csFeeDistributorPort ports.CsFeeDistributorPort
}

func NewCsFeeDistributorService(storagePort ports.StoragePort, notifierPort ports.NotifierPort, subscriberPort ports.CsFeeDistributorPort) *CsFeeDistributorService {
	return &CsFeeDistributorService{
		notifierPort:         notifierPort,
		csFeeDistributorPort: subscriberPort,
	}
}

// WatchCsFeeDistributorEvents subscribes to Ethereum events and handles them.
// It passes to the csFeeDistrubtor port a function that processes the log data.
func (cfds *CsFeeDistributorService) WatchCsFeeDistributorEvents(ctx context.Context) error {
	return cfds.csFeeDistributorPort.WatchCsFeeDistributorEvents(ctx, cfds)
}

func (vs *CsFeeDistributorService) HandleDistributionDataUpdated(rewardsDistributed *domain.CsfeedistributorDistributionDataUpdated) error {
	// send the notification message
	message := fmt.Sprintf("- 📈 New rewards distributed: %s", rewardsDistributed.TotalClaimableShares)
	if err := vs.notifierPort.SendNotification(message); err != nil {
		log.Printf("Failed to send notification: %v", err)
		return err
	}

	return nil
}
