package storage

import (
	"encoding/json"
	"fmt"
	"lido-events/internal/application/domain"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// Storage handles file-based storage for configuration and operator data
type Storage struct {
	DBFile              string
	mu                  sync.RWMutex // RWMutex for concurrent access
	operatorIdListeners []chan []*big.Int
}

// NewStorageAdapter creates a new instance of Storage and ensures the DB directory exists
func NewStorageAdapter(dbDirectory string) (*Storage, error) {
	// Trim any trailing slash from the directory path
	dbDirectory = strings.TrimRight(dbDirectory, "/")

	// Check if the directory exists
	if _, err := os.Stat(dbDirectory); os.IsNotExist(err) {
		// Attempt to create the directory
		if err := os.MkdirAll(dbDirectory, os.ModePerm); err != nil {
			return nil, fmt.Errorf("failed to create directory %s: %w", dbDirectory, err)
		}
	}

	// Construct the DB file path
	dbFile := filepath.Join(dbDirectory, "db.json")

	return &Storage{
		DBFile: dbFile,
	}, nil
}

// Database structure for the new storage format
type Database struct {
	Telegram  domain.TelegramConfig   `json:"telegram"`
	Operators map[string]OperatorData `json:"operators"` // indexed by operator ID
	Events    Events                  `json:"events"`
}

type OperatorData struct {
	Reports            domain.Reports            `json:"reports"`
	ExitRequests       domain.ExitRequests       `json:"exitRequests"`
	NodeOperatorEvents domain.NodeOperatorEvents `json:"nodeOperatorEvents"`
}

type Events struct {
	DistributionLogUpdated struct {
		LastProcessedBlock uint64   `json:"lastProcessedBlock"`
		PendingHashes      []string `json:"pendingHashes"`
	} `json:"distributionLogUpdated"`
	ValidatorExitRequest struct {
		LastProcessedBlock uint64 `json:"lastProcessedBlock"`
	} `json:"validatorExitRequest"`
	CsModule struct {
		LastProcessedBlock uint64 `json:"lastProcessedBlock"`
	} `json:"csModule"`
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
			CsModule: struct {
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
		return Database{}, fmt.Errorf("error reading database file: %v", err)
	}

	// Check if file is empty before attempting to unmarshal
	if len(file) == 0 {
		return db, nil
	}

	// Unmarshal the file content into db
	err = json.Unmarshal(file, &db)
	if err != nil {
		return Database{}, fmt.Errorf("error unmarshalling database file: %v", err)
	}

	// Ensure nested fields are properly initialized
	if db.Operators == nil {
		db.Operators = make(map[string]OperatorData)
	}
	if db.Events.DistributionLogUpdated.PendingHashes == nil {
		db.Events.DistributionLogUpdated.PendingHashes = []string{}
	}

	// Ensure each OperatorData's nested fields are initialized
	for key, operatorData := range db.Operators {
		if operatorData.Reports == nil {
			operatorData.Reports = make(domain.Reports)
		}
		if operatorData.ExitRequests == nil {
			operatorData.ExitRequests = make(domain.ExitRequests)
		}
		if operatorData.NodeOperatorEvents.NodeOperatorAdded == nil {
			operatorData.NodeOperatorEvents.NodeOperatorAdded = []domain.CsmoduleNodeOperatorAdded{}
		}
		if operatorData.NodeOperatorEvents.NodeOperatorManagerAddressChanged == nil {
			operatorData.NodeOperatorEvents.NodeOperatorManagerAddressChanged = []domain.CsmoduleNodeOperatorManagerAddressChanged{}
		}
		if operatorData.NodeOperatorEvents.NodeOperatorRewardAddressChanged == nil {
			operatorData.NodeOperatorEvents.NodeOperatorRewardAddressChanged = []domain.CsmoduleNodeOperatorRewardAddressChanged{}
		}
		db.Operators[key] = operatorData // Update map with initialized OperatorData
	}

	return db, nil
}

func (fs *Storage) SaveDatabase(db Database) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	file, err := json.MarshalIndent(db, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling database: %v", err)
	}
	err = os.WriteFile(fs.DBFile, file, 0644)
	if err != nil {
		return fmt.Errorf("error writing database file: %v", err)
	}

	return nil
}
