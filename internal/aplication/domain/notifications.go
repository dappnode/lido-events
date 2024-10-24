package domain

// NotificationMessageMap maps event names to their associated notification messages.
var NotificationMessageMap = map[EventName]string{
	DepositedSigningKeysCountChanged:         "- 🤩 Node Operator's keys received deposits",
	ELRewardsStealingPenaltyReported:         "- 🚨 Penalty for stealing EL rewards reported",
	ELRewardsStealingPenaltySettled:          "- 🚨 EL rewards stealing penalty confirmed and applied",
	ELRewardsStealingPenaltyCancelled:        "- 😮‍💨 Cancelled penalty for stealing EL rewards",
	InitialSlashingSubmitted:                 "- 🚨 Initial slashing submitted for one of the validators",
	KeyRemovalChargeApplied:                  "- 🔑 Applied charge for key removal",
	NodeOperatorManagerAddressChangeProposed: "- ℹ️ New manager address proposed",
	NodeOperatorManagerAddressChanged:        "- ✅ Manager address changed",
	NodeOperatorRewardAddressChangeProposed:  "- ℹ️ New rewards address proposed",
	NodeOperatorRewardAddressChanged:         "- ✅ Rewards address changed",
	StuckSigningKeysCountChanged:             "- 🚨 Reported stuck keys that were not exited in time",
	VettedSigningKeysCountDecreased:          "- 🚨 Uploaded invalid keys",
	WithdrawalSubmitted:                      "- 👀 Key withdrawal information submitted",
	TotalSigningKeysCountChanged:             "- 👀 New keys uploaded or removed",
	ValidatorExitRequest:                     "- 🚨 One of the validators requested to exit",
	PublicRelease:                            "- 🎉 Public release of CSM!",
	DistributionDataUpdated:                  "- 📈 New rewards distributed",
	ReportSubmitted:                          "- 📈 New submitted report",
	// Additional custom notifications can be added here as needed.
	"TelegramTokenUpdated":        "- ✅ Telegram token changed",
	"TelegramChatIdUpdated":       "- ✅ Telegram token changed",
	"OperatorIdUpdated":           "- ✅ Operator ID changed",
	"ValidatorSuccessfullyExited": "- ✅ Validator exited successfully",
	"ValidatorCouldNotExit":       "- 🚨 Could not exit validator, a manual exit is required",
}
