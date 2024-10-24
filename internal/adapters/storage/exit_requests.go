package storage

import (
	"encoding/json"
	"lido-events/internal/domain"
	"os"
)

// GetExitRequests loads exit requests from the JSON file
func (fs *Storage) GetExitRequests() (domain.ExitRequest, error) {
	file, err := os.ReadFile(fs.ExitRequestFile)
	if err != nil {
		return nil, err
	}
	var data map[string]string
	err = json.Unmarshal(file, &data)
	return data, err
}

// SaveExitRequests saves exit requests to the JSON file
func (fs *Storage) SaveExitRequests(exitRequests domain.ExitRequest) error {
	file, err := json.MarshalIndent(exitRequests, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(fs.ExitRequestFile, file, 0644)
	return err
}
