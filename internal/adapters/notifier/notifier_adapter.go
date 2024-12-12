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
	storagePort   ports.StoragePort
}

func NewNotifierAdapter(ctx context.Context, storageAdapter ports.StoragePort) (*TelegramBot, error) {
	const servicePrefix = "TelegramNotifier"

	adapter := &TelegramBot{
		servicePrefix: servicePrefix,
		storagePort:   storageAdapter,
	}

	// Initialize if config exists.
	config, err := storageAdapter.GetTelegramConfig()
	if err != nil {
		logger.WarnWithPrefix(servicePrefix, "Failed to load initial Telegram configuration: %v", err)
	} else if config.Token != "" && config.UserID != 0 {
		bot, err := tgbotapi.NewBotAPI(config.Token)
		if err != nil {
			logger.WarnWithPrefix(servicePrefix, "Failed to initialize Telegram bot: %v", err)
		} else {
			adapter.Bot = bot
			adapter.UserID = config.UserID
			logger.InfoWithPrefix(servicePrefix, "Telegram bot initialized successfully.")
		}
	} else {
		logger.WarnWithPrefix(servicePrefix, "Initial Telegram configuration is incomplete. Notifications disabled.")
	}

	return adapter, nil
}

func (tb *TelegramBot) UpdateBotConfig() error {
	config, err := tb.storagePort.GetTelegramConfig()
	if err != nil {
		return fmt.Errorf("failed to get Telegram config: %v", err)
	}

	if config.Token == "" || config.UserID == 0 {
		tb.Bot = nil
		tb.UserID = 0
		logger.WarnWithPrefix(tb.servicePrefix, "Incomplete Telegram configuration. Notifications disabled.")
		return nil
	}

	bot, err := tgbotapi.NewBotAPI(config.Token)
	if err != nil {
		tb.Bot = nil
		return fmt.Errorf("failed to update Telegram bot: %v", err)
	}

	tb.Bot = bot
	tb.UserID = config.UserID

	// Send a test notification after the bot has been updated
	if err := tb.SendNotification("🔑 Updated telegram configuration successfully"); err != nil {
		return fmt.Errorf("failed to send test notification: %w", err)
	}

	logger.InfoWithPrefix(tb.servicePrefix, "Telegram bot configuration updated successfully.")

	return nil
}

func (tb *TelegramBot) SendNotification(message string) error {
	if tb.Bot == nil {
		logger.WarnWithPrefix(tb.servicePrefix, "Telegram bot is not initialized. Skipping notification.")
		return nil
	}

	logger.DebugWithPrefix(tb.servicePrefix, "Sending notification to user ID: %d", tb.UserID)
	msg := tgbotapi.NewMessage(tb.UserID, message)
	_, err := tb.Bot.Send(msg)
	if err != nil {
		return fmt.Errorf("failed to send Telegram message: %w", err)
	}
	logger.DebugWithPrefix(tb.servicePrefix, "Notification sent successfully.")
	return nil
}
