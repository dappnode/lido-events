package storage

import (
	"encoding/json"
	"lido-events/internal/aplication/domain"
	"math/big"
	"os"
)

// SaveTelegramConfig updates the Telegram token in the config file
func (fs *Storage) SaveTelegramConfig(telegramConfig domain.TelegramConfig) error {
	cfg, err := loadConfig(fs)
	if err != nil {
		return err
	}
	cfg.Telegram = telegramConfig
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
func (fs *Storage) SaveOperatorId(operatorID domain.OperatorId) error {
	cfg, err := loadConfig(fs)
	if err != nil {
		return err
	}
	cfg.OperatorID = operatorID
	return saveConfig(fs, cfg)
}

func (fs *Storage) GetOperatorId() (domain.OperatorId, error) {
	cfg, err := loadConfig(fs)
	if err != nil {
		return big.NewInt(0), err
	}
	return domain.OperatorId(cfg.OperatorID), nil
}

func loadConfig(fs *Storage) (domain.Config, error) {
	file, err := os.ReadFile(fs.ConfigFile)
	if err != nil {
		return domain.Config{}, err
	}
	var cfg domain.Config
	err = json.Unmarshal(file, &cfg)
	return cfg, err
}

func saveConfig(fs *Storage, cfg domain.Config) error {
	file, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(fs.ConfigFile, file, 0644)
	return err
}
