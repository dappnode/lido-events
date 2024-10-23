package storage

import (
	"encoding/json"
	"lido-events/internal/domain"
	"lido-events/internal/infrastructure/config"
	"os"
)

// SaveTelegramConfig updates the Telegram token in the config file
func (fs *Storage) SaveTelegramConfig(telegramConfig domain.TelegramConfig) error {
	cfg, err := loadConfig(fs)
	if err != nil {
		return err
	}
	cfg.Telegram = config.TelegramConfig(telegramConfig)
	return saveConfig(fs, cfg)
}

// GetTelegramConfig loads the Telegram configuration from the config file
func (fs *Storage) GetTelegramConfig() (domain.TelegramConfig, error) {
	cfg, err := loadConfig(fs)
	if err != nil {
		return domain.TelegramConfig{}, err
	}
	return domain.TelegramConfig(cfg.Telegram), nil
}

// SaveOperatorId updates the operator ID in the config file
func (fs *Storage) SaveOperatorId(operatorID string) error {
	cfg, err := loadConfig(fs)
	if err != nil {
		return err
	}
	cfg.OperatorID = operatorID
	return saveConfig(fs, cfg)
}

func (fs *Storage) GetOperatorId() (string, error) {
	cfg, err := loadConfig(fs)
	if err != nil {
		return "", err
	}
	return cfg.OperatorID, nil
}

func loadConfig(fs *Storage) (config.Config, error) {
	file, err := os.ReadFile(fs.ConfigFile)
	if err != nil {
		return config.Config{}, err
	}
	var cfg config.Config
	err = json.Unmarshal(file, &cfg)
	return cfg, err
}

func saveConfig(fs *Storage, cfg config.Config) error {
	file, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(fs.ConfigFile, file, 0644)
	return err
}
