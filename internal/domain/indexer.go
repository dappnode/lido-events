package domain

type Indexer interface {
	UpdateTelegramToken(token string) error
	UpdateOperatorID(operatorID string) error
	GetLidoReport(start, end int) (Report, error)
	GetExitRequests() (map[string]string, error)
}

// Notifier interface to send notifications
type Notifier interface {
	Notify(message string) error
}

type Report struct {
	Epoch      string // TODO: epoch should be the index of the report
	Threshold  string
	Validators map[string]ValidatorPerformance
}

type ValidatorPerformance struct {
	Included int
	Assigned int
}

// Config struct representing configuration settings for the application
type Config struct {
	TelegramToken string // e.g., 123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11
	OperatorID    string // e.g., 2536
}
