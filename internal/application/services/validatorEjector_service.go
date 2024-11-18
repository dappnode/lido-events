package services

import (
	"context"
	"fmt"
	"lido-events/internal/application/domain"
	"lido-events/internal/application/ports"
	"lido-events/internal/logger"
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
func (ve *ValidatorEjector) ValidatorEjectorCron(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Call the scan method periodically
			if err := ve.ejectValidator(); err != nil {
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
func (ve *ValidatorEjector) ejectValidator() error {
	logger.DebugWithPrefix(ve.servicePrefix, "Ejecting validators if any")

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
			// if the validator is not active_ongoing, skip
			if exitRequest.Status != domain.StatusActiveOngoing {
				logger.InfoWithPrefix(ve.servicePrefix, "Validator %s is %s and no exit is required, skipping", exitRequest.Event.ValidatorIndex, exitRequest.Status)
				continue
			}

			logger.WarnWithPrefix(ve.servicePrefix, "Validator %s is %s, executing exit", exitRequest.Event.ValidatorIndex, exitRequest.Status)

			// send notification and skip on error
			message := fmt.Sprintf("- 🚨 One of the validators requested to exit: %s", exitRequest.Event.ValidatorIndex)
			ve.notifierPort.SendNotification(message)

			// exit the validator
			if err := ve.exitValidatorPort.ExitValidator(string(exitRequest.Event.ValidatorPubkey), exitRequest.Event.ValidatorIndex.String()); err != nil {
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
			// a maximum of 40 times with a 3 second sleep between each call
			for i := 0; i < 40; i++ {
				logger.DebugWithPrefix(ve.servicePrefix, "Waiting for validator %s to exit", exitRequest.Event.ValidatorIndex)

				validatorStatus, err := ve.beaconchainPort.GetValidatorStatus(string(exitRequest.Event.ValidatorPubkey))
				if err != nil {
					logger.ErrorWithPrefix(ve.servicePrefix, "Error getting validator status", err)
					continue
				}

				if validatorStatus == domain.StatusActiveExiting {
					logger.InfoWithPrefix(ve.servicePrefix, "Validator %s has been exited", exitRequest.Event.ValidatorIndex)

					// update the status on the db using UpdateExitRequest and skip on error
					if err := ve.storagePort.UpdateExitRequestStatus(operatorID.String(), string(exitRequest.Event.ValidatorPubkey), domain.StatusActiveExiting); err != nil {
						logger.ErrorWithPrefix(ve.servicePrefix, "Error updating exit request status", err)
					}

					// send notification and skip on error
					message = fmt.Sprintf("- 🚪 Validator %s has been exited automatically, no manual action required", exitRequest.Event.ValidatorIndex)
					if err := ve.notifierPort.SendNotification(message); err != nil {
						logger.ErrorWithPrefix(ve.servicePrefix, "Error sending exit notification", err)
					}

					break
				}

				time.Sleep(3 * time.Second)
			}

		}
	}

	logger.DebugWithPrefix(ve.servicePrefix, "Ejecting validators completed")
	return nil
}
