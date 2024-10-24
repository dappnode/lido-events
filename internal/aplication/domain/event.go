package domain

// EventName is a custom type to represent event names.
type EventName string

// Define constants for each event name.
const (
	DepositedSigningKeysCountChanged         EventName = "DepositedSigningKeysCountChanged"
	ELRewardsStealingPenaltyReported         EventName = "ELRewardsStealingPenaltyReported"
	ELRewardsStealingPenaltySettled          EventName = "ELRewardsStealingPenaltySettled"
	ELRewardsStealingPenaltyCancelled        EventName = "ELRewardsStealingPenaltyCancelled"
	InitialSlashingSubmitted                 EventName = "InitialSlashingSubmitted"
	KeyRemovalChargeApplied                  EventName = "KeyRemovalChargeApplied"
	NodeOperatorManagerAddressChangeProposed EventName = "NodeOperatorManagerAddressChangeProposed"
	NodeOperatorManagerAddressChanged        EventName = "NodeOperatorManagerAddressChanged"
	NodeOperatorRewardAddressChangeProposed  EventName = "NodeOperatorRewardAddressChangeProposed"
	NodeOperatorRewardAddressChanged         EventName = "NodeOperatorRewardAddressChanged"
	StuckSigningKeysCountChanged             EventName = "StuckSigningKeysCountChanged"
	VettedSigningKeysCountDecreased          EventName = "VettedSigningKeysCountDecreased"
	WithdrawalSubmitted                      EventName = "WithdrawalSubmitted"
	TotalSigningKeysCountChanged             EventName = "TotalSigningKeysCountChanged"
	ValidatorExitRequest                     EventName = "ValidatorExitRequest" // requires store in db and exit
	PublicRelease                            EventName = "PublicRelease"
	DistributionDataUpdated                  EventName = "DistributionDataUpdated"
	ReportSubmitted                          EventName = "ReportSubmitted" // requires store in db
)
