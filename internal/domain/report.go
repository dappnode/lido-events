package domain

type Report struct {
	Epoch      string // TODO: epoch should be the index of the report
	Threshold  string
	Validators map[string]ValidatorPerformance
}

type ValidatorPerformance struct {
	Included int
	Assigned int
}
