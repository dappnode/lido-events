package domain

// EventName is a custom type to represent event names.
type EventName string

// Define constants for each event name.
const (
	DepositedSigningKeysCountChanged         EventName = "DepositedSigningKeysCountChanged"         // CSModule
	ELRewardsStealingPenaltyReported         EventName = "ELRewardsStealingPenaltyReported"         // CSModule
	ELRewardsStealingPenaltySettled          EventName = "ELRewardsStealingPenaltySettled"          // CSModule
	ELRewardsStealingPenaltyCancelled        EventName = "ELRewardsStealingPenaltyCancelled"        // CSModule
	InitialSlashingSubmitted                 EventName = "InitialSlashingSubmitted"                 // CSModule
	KeyRemovalChargeApplied                  EventName = "KeyRemovalChargeApplied"                  // CSModule
	NodeOperatorManagerAddressChangeProposed EventName = "NodeOperatorManagerAddressChangeProposed" // CSModule
	NodeOperatorManagerAddressChanged        EventName = "NodeOperatorManagerAddressChanged"        // CSModule
	NodeOperatorRewardAddressChangeProposed  EventName = "NodeOperatorRewardAddressChangeProposed"  // CSModule
	NodeOperatorRewardAddressChanged         EventName = "NodeOperatorRewardAddressChanged"         // CSModule
	StuckSigningKeysCountChanged             EventName = "StuckSigningKeysCountChanged"             // CSModule
	VettedSigningKeysCountDecreased          EventName = "VettedSigningKeysCountDecreased"          // CSModule
	WithdrawalSubmitted                      EventName = "WithdrawalSubmitted"                      // CSModule
	TotalSigningKeysCountChanged             EventName = "TotalSigningKeysCountChanged"             // CSModule
	ValidatorExitRequest                     EventName = "ValidatorExitRequest"                     // VEBO
	PublicRelease                            EventName = "PublicRelease"                            // CSModule
	DistributionDataUpdated                  EventName = "DistributionDataUpdated"                  // CSFeeDistributor
	ReportSubmitted                          EventName = "ReportSubmitted"                          // VEBO
)
