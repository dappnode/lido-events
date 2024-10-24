package domain

type TelegramConfig struct {
	Token  string `json:"token"`
	ChatID int64  `json:"chatId"`
}

type OperatorId string

type Config struct {
	OperatorID OperatorId     `json:"operatorId"`
	Telegram   TelegramConfig `json:"telegram"`
}
