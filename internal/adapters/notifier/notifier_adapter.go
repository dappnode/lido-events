package notifier

import (
	"context"
	"fmt"
	"lido-events/internal/application/ports"
	"lido-events/internal/logger"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
	Bot           *tgbotapi.BotAPI
	UserID        int64
	servicePrefix string
	mu            sync.RWMutex
}

func NewNotifierAdapter(ctx context.Context, storageAdapter ports.StoragePort) (*TelegramBot, error) {
	const servicePrefix = "TelegramNotifier"

	// Load the initial configuration for Telegram
	initialConfig, err := storageAdapter.GetTelegramConfig()
	if err != nil {
		logger.WarnWithPrefix(servicePrefix, "Failed to load initial Telegram configuration: %v", err)
		// Return a TelegramBot instance without initializing the bot
		return &TelegramBot{servicePrefix: servicePrefix}, nil
	}

	if initialConfig.Token == "" || initialConfig.UserID == 0 {
		logger.WarnWithPrefix(servicePrefix, "Telegram configuration is incomplete. Notifications will be disabled.")
		// Return a TelegramBot instance without initializing the bot
		return &TelegramBot{servicePrefix: servicePrefix}, nil
	}

	// Initialize the bot with the initial token
	bot, err := tgbotapi.NewBotAPI(initialConfig.Token)
	if err != nil {
		logger.WarnWithPrefix(servicePrefix, "Failed to initialize Telegram bot: %v", err)
		// Return a TelegramBot instance without initializing the bot
		return &TelegramBot{servicePrefix: servicePrefix}, nil
	}

	adapter := &TelegramBot{
		Bot:           bot,
		UserID:        initialConfig.UserID,
		servicePrefix: servicePrefix,
	}

	// Listen for configuration updates in a separate goroutine
	telegramConfigUpdates := storageAdapter.RegisterTelegramConfigListener()
	go func() {
		for newConfig := range telegramConfigUpdates {
			adapter.mu.Lock()
			adapter.UserID = newConfig.UserID
			// Update the bot API instance with the new token
			if newConfig.Token != "" {
				updatedBot, err := tgbotapi.NewBotAPI(newConfig.Token)
				if err == nil {
					adapter.Bot = updatedBot
					logger.InfoWithPrefix(servicePrefix, "Telegram bot configuration updated successfully.")
				} else {
					logger.ErrorWithPrefix(servicePrefix, "Failed to update Telegram bot: %v", err)
				}
			} else {
				logger.WarnWithPrefix(servicePrefix, "Received incomplete Telegram configuration. Notifications will be disabled.")
				adapter.Bot = nil // Disable notifications if the new config is invalid
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

	// Check if the bot is initialized
	if tb.Bot == nil {
		logger.WarnWithPrefix(tb.servicePrefix, "Telegram bot is not initialized. Skipping notification.")
		return nil
	}

	msg := tgbotapi.NewMessage(tb.UserID, message)
	_, err := tb.Bot.Send(msg)
	if err != nil {
		return fmt.Errorf("failed to send Telegram message: %w", err)
	}
	logger.DebugWithPrefix(tb.servicePrefix, "Notification sent successfully.")
	return nil
}
