package domain

type Indexer interface {
    UpdateTelegramToken(token string) error // 123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11
    UpdateOperatorID(operatorID string) error // 2536
    GetLidoReport(start, end int) (Report, error) 
    GetExitRequests() (map[string]string, error)
}

type Report struct {
    Epoch       string // TODO: epoch should be the index of the report
    Threshold   string
    Validators  map[string]ValidatorPerformance 
}

type ValidatorPerformance struct {
    Included int
    Assigned int
}