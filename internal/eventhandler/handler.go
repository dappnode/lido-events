package eventhandler

import (
	"lido-events/internal/domain/entities"
	"lido-events/internal/service"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
)

type EventHandler struct {
	operatorService *service.StorageService
	notifier        *service.NotifierService
}

func NewEventHandler(service *service.StorageService, notifier *service.NotifierService) *EventHandler {
	return &EventHandler{operatorService: service}
}

func (eh *EventHandler) HandleEvent(eventName string, vLog types.Log) {
	switch eventName {
	case "ValidatorExitRequest":
		// create a fake map[string]string)
		exitRequests := entities.ExitRequest{
			"validator1": "exit-request",
		}
		err := eh.operatorService.SetExitRequests(exitRequests)
		if err != nil {
			log.Printf("Failed to set exit request: %v", err)
		}
	case "SubmittedReport":
		// Aquí se podría parsear el log para obtener los detalles del reporte
		report := entities.Report{
			Threshold:  "some-threshold", // Ajustar con los valores del evento
			Validators: map[string]entities.ValidatorPerformance{},
		}
		epoch := "some-epoch-value" // Ajustar con los valores del evento
		// create a fake report map[string]domain.Report
		lidoReport := map[string]entities.Report{
			epoch: report,
		}
		err := eh.operatorService.SaveLidoReport(lidoReport)
		if err != nil {
			log.Printf("Failed to set Lido report: %v", err)
		}
	default:
		message := getEventNotificationMessage(eventName)
		if message != "" {
			err := eh.notifier.Send(message)
			if err != nil {
				log.Printf("Failed to send notification: %v", err)
			}
		}
	}
}

func getEventNotificationMessage(eventName string) string {
	events := map[string]string{
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
