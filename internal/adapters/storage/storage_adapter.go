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

	"github.com/ethereum/go-ethereum/common"
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
	Telegram  domain.TelegramConfig                   `json:"telegram"`
	Operators map[string]OperatorData                 `json:"operators"` // indexed by operator ID
	Addresses map[common.Address]domain.AddressEvents `json:"addresses"` // indexed by Ethereum address
	Events    Events                                  `json:"events"`
}

type OperatorData struct {
	Reports      domain.Reports      `json:"reports"`
	ExitRequests domain.ExitRequests `json:"exitRequests"`
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

func (fs *Storage) LoadDatabase() (Database, error) {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	// Initialize an empty Database with default values
	db := Database{
		Telegram:  domain.TelegramConfig{},
		Operators: make(map[string]OperatorData),
		Addresses: make(map[common.Address]domain.AddressEvents),
		Events: Events{
			DistributionLogUpdated: struct {
				LastProcessedBlock uint64   `json:"lastProcessedBlock"`
				PendingHashes      []string `json:"pendingHashes"`
			}{
				LastProcessedBlock: 0,
				PendingHashes:      []string{},
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
	if db.Addresses == nil {
		db.Addresses = make(map[common.Address]domain.AddressEvents)
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
		return fmt.Errorf("error marshalling database: %v", err)
	}
	err = os.WriteFile(fs.DBFile, file, 0644)
	if err != nil {
		return fmt.Errorf("error writing database file: %v", err)
	}

	return nil
}
