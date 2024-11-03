package storage

import (
	"encoding/json"
	"lido-events/internal/application/domain"
	"os"
)

// GetExitRequests loads exit requests from the JSON file
func (fs *Storage) GetExitRequests() (domain.ExitRequests, error) {
	file, err := os.ReadFile(fs.ExitRequestFile)
	if err != nil {
		return domain.ExitRequests{}, err
	}
	var data domain.ExitRequests
	err = json.Unmarshal(file, &data)
	return data, err
}

// SaveExitRequests saves exit requests to the JSON file
func (fs *Storage) SaveExitRequests(exitRequests domain.ExitRequests) error {
	file, err := json.MarshalIndent(exitRequests, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(fs.ExitRequestFile, file, 0644)
	return err
}
