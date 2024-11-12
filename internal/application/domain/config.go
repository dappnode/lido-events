package domain

type TelegramConfig struct {
	Token  string `json:"token"`
	ChatID int64  `json:"chatId"`
}

type Config struct {
	OperatorID string         `json:"operatorId"`
	Telegram   TelegramConfig `json:"telegram"`
}
