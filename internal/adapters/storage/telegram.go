package storage

import (
	"lido-events/internal/application/domain"
)

// Telegram Config Methods

func (fs *Storage) SaveTelegramConfig(config domain.TelegramConfig) error {
	db, err := fs.LoadDatabase()
	if err != nil {
		return err
	}
	db.Telegram = config

	if err := fs.SaveDatabase(db); err != nil {
		return err
	}

	// No notification to listeners needed anymore since we are doing synchronous updates.
	return nil
}

func (fs *Storage) GetTelegramConfig() (domain.TelegramConfig, error) {
	db, err := fs.LoadDatabase()
	if err != nil {
		return domain.TelegramConfig{}, err
	}
	return db.Telegram, nil
}
