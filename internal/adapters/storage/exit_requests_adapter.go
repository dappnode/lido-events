package storage

import (
	"encoding/json"
	"fmt"
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

// Get last processed epoch from the JSON file
func (fs *Storage) GetLastProcessedEpoch() (uint64, error) {
	exitRequests, err := fs.GetExitRequests()
	if err != nil {
		return 0, err
	}
	return exitRequests.LastProcessedEpoch, nil
}

// SaveLastProcessedEpoch saves the last processed epoch to the JSON file
func (fs *Storage) SaveLastProcessedEpoch(epoch uint64) error {
	exitRequests, err := fs.GetExitRequests()
	if err != nil {
		return err
	}
	exitRequests.LastProcessedEpoch = epoch
	return fs.SaveExitRequests(exitRequests)
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

// SaveExitRequest saves a single exit request to the JSON file
func (fs *Storage) SaveExitRequest(pubkey string, exitRequest domain.ExitRequest) error {
	// Retrieve the existing exit requests from storage
	exitRequests, err := fs.GetExitRequests()
	if err != nil {
		return err
	}

	// Save the new exit request, indexed by the pubkey
	exitRequests.Requests[pubkey] = exitRequest

	// Persist the updated exit requests back to storage
	return fs.SaveExitRequests(exitRequests)
}

// UpdateExitRequestStatus Update status of a single exit request in the JSON file
func (fs *Storage) UpdateExitRequestStatus(pubkey string, status domain.ValidatorStatus) error {
	// Retrieve the existing exit requests from storage
	exitRequests, err := fs.GetExitRequests()
	if err != nil {
		return err
	}

	// Update the status of the exit request
	exitRequest, ok := exitRequests.Requests[pubkey]
	if !ok {
		return fmt.Errorf("exit request not found for pubkey %s", pubkey)
	}
	exitRequest.Status = status

	// Save the updated exit request back to storage
	return fs.SaveExitRequest(pubkey, exitRequest)
}
