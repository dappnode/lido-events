package domain

type ExitRequests map[string]ExitRequest // indexed by validator index

type ExitRequest struct {
	Event              VeboValidatorExitRequest `json:"event"`
	Status             ValidatorStatus          `json:"status"`
	ValidatorPubkeyHex string                   `json:"validator_pubkey_hex"`
	// ValidatorPubkeyHex is added because VeboValidatorExitRequest stores the validator pubkey as a byte array,
	// and we need to have it as a hex string for the ejector to use later on.
}

type ValidatorStatus string

const (
	StatusPendingInitialized ValidatorStatus = "pending_initialized"
	StatusPendingQueued      ValidatorStatus = "pending_queued"
	StatusActiveOngoing      ValidatorStatus = "active_ongoing"
	StatusActiveExiting      ValidatorStatus = "active_exiting"
	StatusActiveSlashed      ValidatorStatus = "active_slashed"
	StatusExitedUnslashed    ValidatorStatus = "exited_unslashed"
	StatusExitedSlashed      ValidatorStatus = "exited_slashed"
	StatusWithdrawalPossible ValidatorStatus = "withdrawal_possible"
	StatusWithdrawalDone     ValidatorStatus = "withdrawal_done"
)
