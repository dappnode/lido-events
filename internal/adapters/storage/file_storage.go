package storage

import (
	"encoding/json"
	"lido-events/internal/domain"
	"os"
	"strconv"
)

type FileStorage struct {
	PerformanceFile string
	ExitRequestFile string
	ConfigFile      string
}

// Implement UpdateTelegramToken (required by domain.Storage)
func (fs *FileStorage) UpdateTelegramToken(token string) error {
	config, err := fs.LoadConfig()
	if err != nil {
		return err
	}
	config.TelegramToken = token
	return fs.saveConfig(config)
}

// Implement UpdateOperatorID (required by domain.Storage)
func (fs *FileStorage) UpdateOperatorID(operatorID string) error {
	config, err := fs.LoadConfig()
	if err != nil {
		return err
	}
	config.OperatorID = operatorID
	return fs.saveConfig(config)
}

// Implement GetLidoReport (required by domain.Storage)
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

// Implement GetExitRequests (required by domain.Storage)
func (fs *FileStorage) GetExitRequests() (map[string]string, error) {
	file, err := os.ReadFile(fs.ExitRequestFile)
	if err != nil {
		return nil, err
	}
	var data map[string]string
	err = json.Unmarshal(file, &data)
	return data, err
}

// LoadConfig loads the configuration from the JSON file
func (fs *FileStorage) LoadConfig() (domain.Config, error) {
	file, err := os.ReadFile(fs.ConfigFile)
	if err != nil {
		return domain.Config{}, err
	}
	var config domain.Config
	err = json.Unmarshal(file, &config)
	return config, err
}

// saveConfig saves the updated configuration back to the config file
func (fs *FileStorage) saveConfig(config domain.Config) error {
	data, err := json.Marshal(config)
	if err != nil {
		return err
	}
	return os.WriteFile(fs.ConfigFile, data, 0644)
}

// LoadPerformance loads the performance data from the JSON file
func (fs *FileStorage) LoadPerformance() (map[string]domain.Report, error) {
	file, err := os.ReadFile(fs.PerformanceFile)
	if err != nil {
		return nil, err
	}
	var data map[string]domain.Report
	err = json.Unmarshal(file, &data)
	return data, err
}
