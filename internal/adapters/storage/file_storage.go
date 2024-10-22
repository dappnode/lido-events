package storage

import (
	"encoding/json"
	"lido-events/internal/domain"
	"lido-events/internal/infrastructure/config"
	"os"
	"strconv"
)

// FileStorage handles file-based storage for validator performance and exit requests
type FileStorage struct {
	PerformanceFile string
	ExitRequestFile string
	ConfigFile      string
}

// UpdateTelegramToken updates the Telegram token in the config file
func (fs *FileStorage) UpdateTelegramToken(token string) error {
	cfg, err := fs.LoadConfig()
	if err != nil {
		return err
	}
	cfg.Telegram.Token = token
	return fs.saveConfig(cfg)
}

// UpdateOperatorID updates the operator ID in the config file
func (fs *FileStorage) UpdateOperatorID(operatorID string) error {
	cfg, err := fs.LoadConfig()
	if err != nil {
		return err
	}
	cfg.OperatorID = operatorID
	return fs.saveConfig(cfg)
}

// GetLidoReport loads the Lido report data from the performance JSON file for the given epoch range
func (fs *FileStorage) GetLidoReport(start, end int) (map[string]domain.Report, error) {
	performanceData, err := fs.LoadPerformance()
	if err != nil {
		return nil, err
	}
	reportData := make(map[string]domain.Report)
	for epoch, report := range performanceData {
		epochInt, err := strconv.Atoi(epoch)
		if err != nil {
			continue
		}
		if epochInt >= start && epochInt <= end {
			reportData[epoch] = report
		}
	}
	return reportData, nil
}

// GetExitRequests loads exit requests from the JSON file
func (fs *FileStorage) GetExitRequests() (map[string]string, error) {
	file, err := os.ReadFile(fs.ExitRequestFile)
	if err != nil {
		return nil, err
	}
	var data map[string]string
	err = json.Unmarshal(file, &data)
	return data, err
}

// LoadPerformance loads the validator performance data from the JSON file
func (fs *FileStorage) LoadPerformance() (map[string]domain.Report, error) {
	file, err := os.ReadFile(fs.PerformanceFile)
	if err != nil {
		return nil, err
	}
	var data map[string]domain.Report
	err = json.Unmarshal(file, &data)
	return data, err
}

// LoadConfig loads the configuration from the JSON file using the infrastructure config struct
func (fs *FileStorage) LoadConfig() (config.Config, error) {
	file, err := os.ReadFile(fs.ConfigFile)
	if err != nil {
		return config.Config{}, err
	}
	var cfg config.Config
	err = json.Unmarshal(file, &cfg)
	return cfg, err
}

// saveConfig saves the updated configuration back to the config file
func (fs *FileStorage) saveConfig(cfg config.Config) error {
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	return os.WriteFile(fs.ConfigFile, data, 0644)
}
