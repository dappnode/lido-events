package storage

import (
	"encoding/json"
	"lido-events/internal/application/domain"
	"math/big"
	"os"
	"sync"
)

// Storage handles file-based storage for configuration and operator data
type Storage struct {
	DBFile                  string
	mu                      sync.RWMutex // RWMutex for concurrent access
	operatorIdListeners     []chan []*big.Int
	telegramConfigListeners []chan domain.TelegramConfig
}

func NewStorageAdapter() *Storage {
	return &Storage{
		DBFile: "db.json",
	}
}

// Database structure for the new storage format
type Database struct {
	Telegram  domain.TelegramConfig   `json:"telegram"`
	Operators map[string]OperatorData `json:"operators"` // indexed by operator ID
	Events    Events                  `json:"events"`
}

type Events struct {
	DistributionLogUpdated struct {
		LastProcessedBlock uint64   `json:"lastProcessedBlock"`
		PendingHashes      []string `json:"pendingHashes"`
	} `json:"distributionLogUpdated"`
	ValidatorExitRequest struct {
		LastProcessedBlock uint64 `json:"lastProcessedBlock"`
	} `json:"validatorExitRequest"`
}

type OperatorData struct {
	Performance  domain.Reports      `json:"performance"`
	ExitRequests domain.ExitRequests `json:"exitRequests"`
}

func (fs *Storage) LoadDatabase() (Database, error) {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	// Initialize an empty Database with default values
	db := Database{
		Telegram:  domain.TelegramConfig{},
		Operators: make(map[string]OperatorData), // Initialize Operators as an empty map
		Events: Events{
			DistributionLogUpdated: struct {
				LastProcessedBlock uint64   `json:"lastProcessedBlock"`
				PendingHashes      []string `json:"pendingHashes"`
			}{
				LastProcessedBlock: 0,
				PendingHashes:      []string{}, // Initialize as empty slice
			},
			ValidatorExitRequest: struct {
				LastProcessedBlock uint64 `json:"lastProcessedBlock"`
			}{
				LastProcessedBlock: 0,
			},
		},
	}

	file, err := os.ReadFile(fs.DBFile)
	if err != nil {
		if os.IsNotExist(err) {
			return db, nil
		}
		return Database{}, err
	}

	// Check if file is empty before attempting to unmarshal
	if len(file) == 0 {
		return db, nil
	}

	// Unmarshal the file content into db
	err = json.Unmarshal(file, &db)
	if err != nil {
		return Database{}, err
	}

	// Ensure nested fields are properly initialized
	if db.Operators == nil {
		db.Operators = make(map[string]OperatorData)
	}
	if db.Events.DistributionLogUpdated.PendingHashes == nil {
		db.Events.DistributionLogUpdated.PendingHashes = []string{}
	}

	return db, nil
}

func (fs *Storage) SaveDatabase(db Database) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	file, err := json.MarshalIndent(db, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fs.DBFile, file, 0644)
}
