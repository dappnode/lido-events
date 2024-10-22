package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// TelegramBot implements the Notifier interface
type TelegramBot struct {
	Bot    *tgbotapi.BotAPI
	ChatID int64
}

// NewTelegramNotifier creates a new TelegramBot instance
func NewTelegramNotifier(token string, chatID int64) (*TelegramBot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	return &TelegramBot{Bot: bot, ChatID: chatID}, nil
}

// Notify sends a notification message via Telegram
func (tb *TelegramBot) Notify(message string) error {
	msg := tgbotapi.NewMessage(tb.ChatID, message)
	_, err := tb.Bot.Send(msg)
	if err != nil {
		log.Printf("Error sending message via Telegram: %s", err)
		return err
	}
	return nil
}
