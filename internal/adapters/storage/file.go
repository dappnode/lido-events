package storage

import (
	"encoding/json"
	"os"

	"lido-events/internal/domain"
)

type FileStorage struct {
	PerformanceFile string
	ExitRequestFile string
	ConfigFile      string
}

func (fs *FileStorage) LoadPerformance() (map[string]domain.Report, error) {
	file, err := os.ReadFile(fs.PerformanceFile)
	if err != nil {
		return nil, err
	}
	var data map[string]domain.Report
	err = json.Unmarshal(file, &data)
	return data, err
}

func (fs *FileStorage) LoadExitRequests() (map[string]string, error) {
	file, err := os.ReadFile(fs.ExitRequestFile)
	if err != nil {
		return nil, err
	}
	var data map[string]string
	err = json.Unmarshal(file, &data)
	return data, err
}

func (fs *FileStorage) LoadConfig() (domain.Config, error) {
	file, err := os.ReadFile(fs.ConfigFile)
	if err != nil {
		return domain.Config{}, err
	}
	var data domain.Config
	err = json.Unmarshal(file, &data)
	return data, err
}
