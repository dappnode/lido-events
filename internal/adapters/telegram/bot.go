package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type TelegramBot struct {
    Bot *tgbotapi.BotAPI
}

func (tb *TelegramBot) SendMessage(chatID int64, text string) error {
    msg := tgbotapi.NewMessage(chatID, text)
    _, err := tb.Bot.Send(msg)
    if err != nil {
        log.Printf("Error sending message: %s", err)
        return err
    }
    return nil
}
