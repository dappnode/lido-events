package storage

import (
	"encoding/json"
	"os"
)

// GetExitRequests loads exit requests from the JSON file
func (fs *Storage) GetExitRequests() (map[string]string, error) {
	file, err := os.ReadFile(fs.ExitRequestFile)
	if err != nil {
		return nil, err
	}
	var data map[string]string
	err = json.Unmarshal(file, &data)
	return data, err
}

// SaveExitRequests saves exit requests to the JSON file
func (fs *Storage) SaveExitRequests(exitRequests map[string]string) error {
	file, err := json.MarshalIndent(exitRequests, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(fs.ExitRequestFile, file, 0644)
	return err
}
