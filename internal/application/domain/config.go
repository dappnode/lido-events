package domain

type TelegramConfig struct {
	Token  string `json:"token"`
	UserID int64  `json:"userId"`
}

type Config struct {
	OperatorID string         `json:"operatorId"`
	Telegram   TelegramConfig `json:"telegram"`
}
