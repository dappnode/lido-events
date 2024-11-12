package storage

import "lido-events/internal/application/domain"

// Telegram Configuration Methods
func (fs *Storage) SaveTelegramConfig(config domain.TelegramConfig) error {
	db, err := fs.LoadDatabase()
	if err != nil {
		return err
	}
	db.Telegram = config
	return fs.SaveDatabase(db)
}

func (fs *Storage) GetTelegramConfig() (domain.TelegramConfig, error) {
	db, err := fs.LoadDatabase()
	if err != nil {
		return domain.TelegramConfig{}, err
	}
	return db.Telegram, nil
}
