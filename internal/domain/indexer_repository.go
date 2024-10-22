package domain

type IndexerRepository interface {
	UpdateTelegramToken(token string) error
	UpdateOperatorID(operatorID string) error
	GetLidoReport(start, end int) (map[string]Report, error) // indexed by epoch
	GetExitRequests() (map[string]string, error)
}

type Report struct {
	Threshold  string
	Validators map[string]ValidatorPerformance // indexed by validator pubkey
}

type ValidatorPerformance struct {
	Included int
	Assigned int
}
