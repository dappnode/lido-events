package notifier

import (
	"context"
	"lido-events/internal/application/domain"
	"lido-events/internal/config"
	"log"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// TelegramBot implements the Notifier interface
type TelegramBot struct {
	Bot    *tgbotapi.BotAPI
	ChatID int64
	mu     sync.Mutex // Mutex for safe concurrent updates
}

// NewNotifierAdapter creates a new TelegramBot instance and starts listening for config updates.
func NewNotifierAdapter(ctx context.Context, configManager *config.ConfigManager) (*TelegramBot, error) {
	// Get the initial Telegram configuration
	initialConfig := configManager.GetTelegramConfig()

	// Create the Telegram bot with the initial configuration
	bot, err := tgbotapi.NewBotAPI(initialConfig.Token)
	if err != nil {
		return nil, err
	}

	// Initialize the TelegramBot instance
	telegramBot := &TelegramBot{
		Bot:    bot,
		ChatID: initialConfig.ChatID,
	}

	// Start listening for configuration updates
	updateChan := configManager.GetTelegramConfigUpdateChannel()
	telegramBot.ListenForTelegramConfigUpdates(ctx, updateChan)

	return telegramBot, nil
}

func (tb *TelegramBot) ListenForTelegramConfigUpdates(ctx context.Context, updateChan <-chan domain.TelegramConfig) {
	go func() {
		for {
			select {
			case newConfig := <-updateChan:
				tb.mu.Lock()
				tb.Bot, _ = tgbotapi.NewBotAPI(newConfig.Token) // Reinitialize with new token
				tb.ChatID = newConfig.ChatID
				tb.mu.Unlock()
				log.Println("Updated TelegramBot config")
			case <-ctx.Done():
				return
			}
		}
	}()
}

// Notify sends a notification message via Telegram
func (tb *TelegramBot) SendNotification(message string) error {
	msg := tgbotapi.NewMessage(tb.ChatID, message)
	_, err := tb.Bot.Send(msg)
	if err != nil {
		log.Printf("Error sending message via Telegram: %s", err)
		return err
	}
	return nil
}
