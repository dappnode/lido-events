package storage

import (
	"lido-events/internal/application/domain"
)

// TODO: determine if token should be stored hashed

// Telegram Configuration Methods

// SaveTelegramConfig saves the Telegram configuration to storage and notifies listeners of the update.
func (fs *Storage) SaveTelegramConfig(config domain.TelegramConfig) error {
	db, err := fs.LoadDatabase()
	if err != nil {
		return err
	}
	db.Telegram = config

	if err := fs.SaveDatabase(db); err != nil {
		return err
	}

	fs.notifyTelegramConfigListeners() // Notify listeners of the change
	return nil
}

// GetTelegramConfig retrieves the Telegram configuration from storage.
func (fs *Storage) GetTelegramConfig() (domain.TelegramConfig, error) {
	db, err := fs.LoadDatabase()
	if err != nil {
		return domain.TelegramConfig{}, err
	}
	return db.Telegram, nil
}

// RegisterTelegramConfigListener registers a channel to receive updates when the Telegram config changes.
func (fs *Storage) RegisterTelegramConfigListener() chan domain.TelegramConfig {
	updateChan := make(chan domain.TelegramConfig, 1)
	fs.telegramConfigListeners = append(fs.telegramConfigListeners, updateChan)
	return updateChan
}

// notifyTelegramConfigListeners sends updates to all registered listeners of Telegram config changes.
func (fs *Storage) notifyTelegramConfigListeners() {
	config, err := fs.GetTelegramConfig()
	if err != nil {
		return
	}

	for _, listener := range fs.telegramConfigListeners {
		select {
		case listener <- config:
		default:
			// Ignore if channel is full to prevent blocking
		}
	}
}
