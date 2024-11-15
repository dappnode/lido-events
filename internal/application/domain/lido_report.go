package domain

import "github.com/ethereum/go-ethereum/core/types"

type Reports map[string]Report // indexed by epoch

type Report struct {
	Threshold  string
	Validators Validators
}

type Validators map[string]ValidatorPerformance // indexed by validator pubkey

type ValidatorPerformance struct {
	Included string
	Assigned string
}

type BindingsDistributionLogUpdated struct {
	LogCid string
	Raw    types.Log // Blockchain specific contextual infos
}
