package services

import (
	"context"
	"fmt"
	"lido-events/internal/application/domain"
	"lido-events/internal/application/ports"
	"log"
	"time"
)

type ValidatorEjector struct {
	storagePort       ports.StoragePort
	notifierPort      ports.NotifierPort
	exitValidatorPort ports.ExitValidator
	beaconchainPort   ports.Beaconchain
}

func NewValidatorEjectorService(storagePort ports.StoragePort, notifierPort ports.NotifierPort, exitValidatorPort ports.ExitValidator, beaconchainPort ports.Beaconchain) *ValidatorEjector {
	return &ValidatorEjector{
		storagePort,
		notifierPort,
		exitValidatorPort,
		beaconchainPort,
	}
}

// ValidatorEjectorCron starts a periodic service to exit validators that requested to exit
func (ve *ValidatorEjector) ValidatorEjectorCron(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Call the scan method periodically
			ve.ejectValidator()
		case <-ctx.Done():
			// Stop the periodic scan if the context is canceled
			log.Println("Stopping periodic ejector for ValidatorExitRequest events")
			return
		}
	}
}

// ejectValidator orchestrates the voluntary exit process for a validator
func (ve *ValidatorEjector) ejectValidator() error {
	// get Operator IDs
	operatorIDs, err := ve.storagePort.GetOperatorIds()
	if err != nil {
		return err
	}

	for _, operatorID := range operatorIDs {

		//  get exit requests
		exitRequests, err := ve.storagePort.GetExitRequests(operatorID.String())
		if err != nil {
			return err
		}

		for _, exitRequest := range exitRequests {
			// if the validator is not active_ongoing, skip
			if exitRequest.Status != domain.StatusActiveOngoing {
				log.Printf("Validator %s is %s, skipping", exitRequest.Event.ValidatorIndex, exitRequest.Status)
				continue
			}
			// send notification
			log.Printf("Validator %s is %s, starting exit", exitRequest.Event.ValidatorIndex, exitRequest.Status)
			message := fmt.Sprintf("- 🚨 One of the validators requested to exit: %s", exitRequest.Event.ValidatorIndex)
			if err := ve.notifierPort.SendNotification(message); err != nil {
				log.Printf("Failed to send notification: %v", err)
				// continue on error
			}

			// exit the validator
			if err := ve.exitValidatorPort.ExitValidator(string(exitRequest.Event.ValidatorPubkey), exitRequest.Event.ValidatorIndex.String()); err != nil {
				log.Printf("Failed to exit validator %s: %v", exitRequest.Event.ValidatorIndex, err)
				// send manual exit link
				// TODO: wait for PR in docs to add the proper link
				message = fmt.Sprintf("- 🚪 Validator %s failed to exit, a manual exit is required. Click here to learn how to do the exit manually %s", exitRequest.Event.ValidatorIndex, "https://docs.dappnode.io/docs/user/staking/gnosis-chain/solo#1-exit-the-validator-from-the-dappnode-ui")
				if err := ve.notifierPort.SendNotification(message); err != nil {
					log.Printf("Failed to send notification: %v", err)
				}
				continue
			}

			// wait for the transaction to be included
			// call ve.beaconchainPort.GetValidatorStatus(string(validator.Event.ValidatorPubkey)) in a loop until the status is domain.StatusActiveExiting
			// a maximum of 40 times with a 3 second sleep between each call
			for i := 0; i < 40; i++ {
				validatorStatus, err := ve.beaconchainPort.GetValidatorStatus(string(exitRequest.Event.ValidatorPubkey))
				if err != nil {
					log.Printf("Failed to get validator status: %v", err)
					// continue on error
				}

				if validatorStatus == domain.StatusActiveExiting {
					log.Printf("Validator %s has been exited", exitRequest.Event.ValidatorIndex)

					// update the status on the db using UpdateExitRequest
					if err := ve.storagePort.UpdateExitRequestStatus(operatorID.String(), string(exitRequest.Event.ValidatorPubkey), domain.StatusActiveExiting); err != nil {
						log.Printf("Failed to update exit request: %v", err)
						// continue on error
					}

					// send notification
					message = fmt.Sprintf("- 🚪 Validator %s has been exited automatically, no manual action required", exitRequest.Event.ValidatorIndex)
					if err := ve.notifierPort.SendNotification(message); err != nil {
						log.Printf("Failed to send notification: %v", err)
					}
					break
				}

				time.Sleep(3 * time.Second)
			}

		}
	}

	return nil
}
