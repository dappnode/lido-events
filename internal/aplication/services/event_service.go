package services

import (
	"lido-events/internal/domain"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
)

type EventService struct {
	storageService  *StorageService
	notifierService *NotifierService
}

func NewEventService(storageService *StorageService, notifierService *NotifierService) *EventService {
	return &EventService{storageService, notifierService}
}

func (eh *EventService) ProcessEvent(eventName domain.EventName, vLog types.Log) error {
	switch eventName {
	case "ValidatorExitRequest":
		// create a fake map[string]string)
		exitRequests := domain.ExitRequest{
			"validator1": "exit-request",
		}
		err := eh.storageService.SetExitRequests(exitRequests)
		if err != nil {
			log.Printf("Failed to set exit request: %v", err)
			return err
		}
		return nil
	case "SubmittedReport":
		// Aquí se podría parsear el log para obtener los detalles del reporte
		report := domain.Report{
			Threshold:  "some-threshold", // Ajustar con los valores del evento
			Validators: map[string]domain.ValidatorPerformance{},
		}
		epoch := "some-epoch-value" // Ajustar con los valores del evento
		// create a fake report map[string]domain.Report
		lidoReport := map[string]domain.Report{
			epoch: report,
		}
		err := eh.storageService.SaveLidoReport(lidoReport)
		if err != nil {
			log.Printf("Failed to set Lido report: %v", err)
			return err
		}
		return nil
	default:
		message := getEventNotificationMessage(eventName)
		if message != "" {
			err := eh.notifierService.Send(message)
			if err != nil {
				log.Printf("Failed to send notification: %v", err)
				return err
			}
		}
		return nil
	}

}

func getEventNotificationMessage(eventName domain.EventName) string {
	events := map[domain.EventName]string{
		"DepositedSigningKeysCountChanged":         "- 🤩 Node Operator's keys received deposits",
		"ELRewardsStealingPenaltyReported":         "- 🚨 Penalty for stealing EL rewards reported",
		"ELRewardsStealingPenaltySettled":          "- 🚨 EL rewards stealing penalty confirmed and applied",
		"ELRewardsStealingPenaltyCancelled":        "- 😮‍💨 Cancelled penalty for stealing EL rewards",
		"InitialSlashingSubmitted":                 "- 🚨 Initial slashing submitted for one of the validators",
		"KeyRemovalChargeApplied":                  "- 🔑 Applied charge for key removal",
		"NodeOperatorManagerAddressChangeProposed": "- ℹ️ New manager address proposed",
		"NodeOperatorManagerAddressChanged":        "- ✅ Manager address changed",
		"NodeOperatorRewardAddressChangeProposed":  "- ℹ️ New rewards address proposed",
		"NodeOperatorRewardAddressChanged":         "- ✅ Rewards address changed",
		"StuckSigningKeysCountChanged":             "- 🚨 Reported stuck keys that were not exited in time",
		"VettedSigningKeysCountDecreased":          "- 🚨 Uploaded invalid keys",
		"WithdrawalSubmitted":                      "- 👀 Key withdrawal information submitted",
		"TotalSigningKeysCountChanged":             "- 👀 New keys uploaded or removed",
		"ValidatorExitRequest":                     "- 🚨 One of the validators requested to exit",
		"PublicRelease":                            "- 🎉 Public release of CSM!",
		"DistributionDataUpdated":                  "- 📈 New rewards distributed",
	}

	return events[eventName]
}
