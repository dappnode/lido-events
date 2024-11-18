package notifier

import (
	"context"
	"fmt"
	"lido-events/internal/application/ports"
	"log"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
	Bot    *tgbotapi.BotAPI
	UserID int64
	mu     sync.RWMutex
}

func NewNotifierAdapter(ctx context.Context, storageAdapter ports.StoragePort) (*TelegramBot, error) {
	// Load the initial configuration for Telegram
	initialConfig, err := storageAdapter.GetTelegramConfig()
	if err != nil {
		log.Printf("NotifierAdapter: Failed to load Telegram configuration: %v", err)
		return nil, err
	}

	// Initialize the bot with the initial token
	bot, err := tgbotapi.NewBotAPI(initialConfig.Token)
	if err != nil {
		return nil, err
	}

	adapter := &TelegramBot{
		Bot:    bot,
		UserID: initialConfig.UserID,
	}

	// Listen for configuration updates in a separate goroutine
	telegramConfigUpdates := storageAdapter.RegisterTelegramConfigListener()
	go func() {
		for newConfig := range telegramConfigUpdates {
			adapter.mu.Lock()
			adapter.UserID = newConfig.UserID
			// Update the bot API instance with the new token
			updatedBot, err := tgbotapi.NewBotAPI(newConfig.Token)
			if err == nil {
				adapter.Bot = updatedBot
			} else {
				log.Printf("NotifierAdapter: Failed to update bot API with new token: %v", err)
			}
			adapter.mu.Unlock()
		}
	}()

	return adapter, nil
}

// SendNotification sends a message via Telegram
func (tb *TelegramBot) SendNotification(message string) error {
	tb.mu.RLock()
	defer tb.mu.RUnlock()

	msg := tgbotapi.NewMessage(tb.UserID, message)
	_, err := tb.Bot.Send(msg)
	if err != nil {
		return fmt.Errorf("failed to send Telegram message: %w", err)
	}
	return nil
}
