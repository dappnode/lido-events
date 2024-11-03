package domain

type ExitRequests struct {
	LastProcessedEpochFinalized int                    `json:"lastProcessedEpochFinalized"`
	Validators                  map[string]ExitRequest `json:"validators"` // indexed by validator pubkey
}

type ExitRequest struct {
	Event  VeboValidatorExitRequest `json:"event"`
	Status ValidatorStatus          `json:"status"`
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
