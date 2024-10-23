package entities

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
