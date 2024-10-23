package domain

type Operator struct {
	ID           string
	Reports      map[string]Report
	ExitRequests map[string]string
	Telegram     TelegramConfig
}

type TelegramConfig struct {
	Token  string `json:"token"`
	ChatID int64  `json:"chatId"`
}

type Report struct {
	Threshold  string
	Validators map[string]ValidatorPerformance // indexed by validator pubkey
}

type ValidatorPerformance struct {
	Included string
	Assigned string
}
