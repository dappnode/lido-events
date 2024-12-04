package notifier

import (
	"context"
	"fmt"
	"lido-events/internal/application/ports"
	"lido-events/internal/logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
	Bot           *tgbotapi.BotAPI
	UserID        int64
	servicePrefix string
}

func NewNotifierAdapter(ctx context.Context, storageAdapter ports.StoragePort) (*TelegramBot, error) {
	const servicePrefix = "TelegramNotifier"

	adapter := &TelegramBot{
		servicePrefix: servicePrefix,
	}

	// Attempt to load the initial configuration for Telegram
	initialConfig, err := storageAdapter.GetTelegramConfig()
	if err != nil {
		logger.WarnWithPrefix(servicePrefix, "Failed to load initial Telegram configuration: %v", err)
	} else if initialConfig.Token != "" && initialConfig.UserID != 0 {
		// Initialize the bot if the configuration is valid
		bot, err := tgbotapi.NewBotAPI(initialConfig.Token)
		if err != nil {
			logger.WarnWithPrefix(servicePrefix, "Failed to initialize Telegram bot: %v", err)
		} else {
			adapter.Bot = bot
			adapter.UserID = initialConfig.UserID
			logger.InfoWithPrefix(servicePrefix, "Telegram bot initialized successfully.")
		}
	} else {
		logger.WarnWithPrefix(servicePrefix, "Initial Telegram configuration is incomplete. Notifications will be disabled.")
	}

	// Listen for configuration updates in a separate goroutine
	telegramConfigUpdates := storageAdapter.RegisterTelegramConfigListener()
	go func() {
		for newConfig := range telegramConfigUpdates {
			adapter.UserID = newConfig.UserID
			// Update the bot API instance with the new token
			if newConfig.Token != "" {
				updatedBot, err := tgbotapi.NewBotAPI(newConfig.Token)
				if err == nil {
					adapter.Bot = updatedBot
					logger.InfoWithPrefix(servicePrefix, "Telegram bot configuration updated successfully.")

					// Send a test notification to confirm the new configuration works
					testMessage := "🔑 Updated telegram configuration successfully"
					testErr := adapter.SendNotification(testMessage)
					if testErr != nil {
						logger.ErrorWithPrefix(servicePrefix, "Failed to send test notification after configuration update: %v", testErr)
					} else {
						logger.InfoWithPrefix(servicePrefix, "Test notification sent successfully.")
					}
				} else {
					logger.ErrorWithPrefix(servicePrefix, "Failed to update Telegram bot: %v", err)
					adapter.Bot = nil // Disable notifications on failure
				}
			} else {
				logger.WarnWithPrefix(servicePrefix, "Received incomplete Telegram configuration. Notifications will be disabled.")
				adapter.Bot = nil // Disable notifications if the new config is invalid
			}
		}
	}()

	return adapter, nil
}

// SendNotification sends a message via Telegram
func (tb *TelegramBot) SendNotification(message string) error {
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
