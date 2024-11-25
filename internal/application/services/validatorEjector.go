package services

import (
	"context"
	"fmt"
	"lido-events/internal/application/domain"
	"lido-events/internal/application/ports"
	"lido-events/internal/logger"
	"sync"
	"time"
)

type ValidatorEjector struct {
	storagePort       ports.StoragePort
	notifierPort      ports.NotifierPort
	exitValidatorPort ports.ExitValidator
	beaconchainPort   ports.Beaconchain
	servicePrefix     string
}

func NewValidatorEjectorService(storagePort ports.StoragePort, notifierPort ports.NotifierPort, exitValidatorPort ports.ExitValidator, beaconchainPort ports.Beaconchain) *ValidatorEjector {
	return &ValidatorEjector{
		storagePort,
		notifierPort,
		exitValidatorPort,
		beaconchainPort,
		"ValidatorEjector",
	}
}

// ValidatorEjectorCron starts a periodic service to exit validators that requested to exit
func (ve *ValidatorEjector) ValidatorEjectorCron(ctx context.Context, interval time.Duration, wg *sync.WaitGroup, firstExecutionComplete chan struct{}) {
	defer wg.Done() // Decrement the counter when the goroutine finishes
	wg.Add(1)       // Increment the WaitGroup counter

	// Wait for the signal from cron event scanner
	<-firstExecutionComplete
	logger.DebugWithPrefix(ve.servicePrefix, "Signal received, starting periodic ejector for ValidatorExitRequest events")

	// Execute immediately on startup
	if err := ve.EjectValidator(); err != nil {
		logger.ErrorWithPrefix(ve.servicePrefix, "Error ejecting validators: %v", err)
	}

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Call the scan method periodically
			if err := ve.EjectValidator(); err != nil {
				logger.ErrorWithPrefix(ve.servicePrefix, "Error ejecting validators: %v", err)
			}
		case <-ctx.Done():
			// Stop the periodic scan if the context is canceled
			logger.InfoWithPrefix(ve.servicePrefix, "Stopping periodic ejector for ValidatorExitRequest events")
			return
		}
	}
}

// ejectValidator orchestrates the voluntary exit process for a validator
func (ve *ValidatorEjector) EjectValidator() error {
	logger.DebugWithPrefix(ve.servicePrefix, "Validator Ejector cron started")

	operatorIDs, err := ve.storagePort.GetOperatorIds()
	if err != nil {
		return err
	}

	for _, operatorID := range operatorIDs {

		//  get exit requests
		exitRequests, err := ve.storagePort.GetExitRequests(operatorID.String())
		if err != nil {
			continue
		}

		for _, exitRequest := range exitRequests {

			// First thing we do is to check the onchain status of the validator. This way we make sure we dont try to exit a validator that is already exiting
			onchainStatus, err := ve.beaconchainPort.GetValidatorStatus(exitRequest.Event.ValidatorIndex.String())
			if err != nil {
				logger.ErrorWithPrefix(ve.servicePrefix, "Error getting validator status from beaconchain, skipping.", err)
				continue
			}

			// if the validator status is not active ongoing or active slashed we skip the exit request because it is already exiting.
			if onchainStatus != domain.StatusActiveOngoing && onchainStatus != domain.StatusActiveSlashed {
				if onchainStatus != domain.StatusPendingInitialized && onchainStatus != domain.StatusPendingQueued {
					logger.InfoWithPrefix(ve.servicePrefix, "Validator %s is %s so no exit request is required, deleting the exit request from db", exitRequest.Event.ValidatorIndex, exitRequest.Status)
					//Since the validator is already exiting, we remove the exit request from the db
					if err := ve.storagePort.DeleteExitRequest(operatorID.String(), exitRequest.Event.ValidatorIndex.String()); err != nil {
						// An error here is no big deal, we will retry to delete this in the next iteration of the cron
						logger.ErrorWithPrefix(ve.servicePrefix, "Error deleting exit request from db", err)
					}
				} else {
					logger.DebugWithPrefix(ve.servicePrefix, "Validator %s is exited to request but it is in a pending status, %s waiting for it to be active", exitRequest.Event.ValidatorIndex, exitRequest.Status)
				}
				continue
			}

			// send notification and skip on error
			message := fmt.Sprintf("- 🚨 Your validator %s is requested to exit.", exitRequest.Event.ValidatorIndex)
			if err := ve.notifierPort.SendNotification(message); err != nil {
				logger.ErrorWithPrefix(ve.servicePrefix, "Error sending exit notification", err)
			}

			// exit the validator
			logger.InfoWithPrefix(ve.servicePrefix, "Exiting validator %s with status %s", exitRequest.Event.ValidatorIndex, exitRequest.Status)
			if err := ve.exitValidatorPort.ExitValidator(exitRequest.ValidatorPubkeyHex, exitRequest.Event.ValidatorIndex.String()); err != nil {
				logger.WarnWithPrefix(ve.servicePrefix, "Failed to exit validator %s, a manual exit is required: %v", exitRequest.Event.ValidatorIndex, err)
				// send notification with manual exit link and skip on errror
				// TODO: wait for PR in docs to add the proper link
				message = fmt.Sprintf("- 🚪 Validator %s failed to exit, a manual exit is required. Click here to learn how to do the exit manually %s", exitRequest.Event.ValidatorIndex, "https://docs.dappnode.io/docs/user/staking/gnosis-chain/solo#1-exit-the-validator-from-the-dappnode-ui")
				if err := ve.notifierPort.SendNotification(message); err != nil {
					logger.ErrorWithPrefix(ve.servicePrefix, "Error sending manual exit notification", err)
				}
				continue
			}

			// wait for the transaction to be included
			// call ve.beaconchainPort.GetValidatorStatus(string(validator.Event.ValidatorPubkey)) in a loop until the status is domain.StatusActiveExiting
			// a maximum of 64 times with a 30 second sleep between each call (check for 32 minutes, two times x minute)
			// TODO: If this ends before the status is ActiveExiting, the user will never get the notification that the validator has been exited successfully.
			// IMPORTANT: This should never take longer to finish than the next cron iteration (defined in main)
			for i := 0; i < 64; i++ {
				logger.DebugWithPrefix(ve.servicePrefix, "Waiting for validator %s to exit", exitRequest.Event.ValidatorIndex)

				validatorStatus, err := ve.beaconchainPort.GetValidatorStatus(exitRequest.ValidatorPubkeyHex)
				if err != nil {
					logger.ErrorWithPrefix(ve.servicePrefix, "Error getting validator status", err)
					continue
				}

				if validatorStatus == domain.StatusActiveExiting || validatorStatus == domain.StatusExitedUnslashed || validatorStatus == domain.StatusExitedSlashed {
					logger.InfoWithPrefix(ve.servicePrefix, "Validator %s has entered the exit queue", exitRequest.Event.ValidatorIndex)

					// send notification and skip on error
					message = fmt.Sprintf("- 🚪 Validator %s has entered the exit queue automatically, no manual action required", exitRequest.Event.ValidatorIndex)
					if err := ve.notifierPort.SendNotification(message); err != nil {
						logger.ErrorWithPrefix(ve.servicePrefix, "Error sending exit notification", err)
					}

					// remove the exit request from the db
					logger.DebugWithPrefix(ve.servicePrefix, "Deleting exit request for validator %s from db", exitRequest.Event.ValidatorIndex)
					if err := ve.storagePort.DeleteExitRequest(operatorID.String(), exitRequest.Event.ValidatorIndex.String()); err != nil {
						logger.ErrorWithPrefix(ve.servicePrefix, "Error deleting exit request from db", err)
					}
					break
				}

				time.Sleep(30 * time.Second)
			}

		}
	}

	logger.DebugWithPrefix(ve.servicePrefix, "Validator Ejector cron finished")
	return nil
}
