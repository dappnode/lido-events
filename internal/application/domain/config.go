package domain

import "math/big"

type TelegramConfig struct {
	Token  string `json:"token"`
	ChatID int64  `json:"chatId"`
}

type OperatorId *big.Int

type Config struct {
	OperatorID OperatorId     `json:"operatorId"`
	Telegram   TelegramConfig `json:"telegram"`
}
