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

// Database structure for unified storage
type Database struct {
	Telegram  domain.TelegramConfig `json:"telegram"`
	Operators OperatorsData         `json:"operators"`
}

// OperatorsData represents the top-level structure for all operators,
// with a global last processed epoch, pending hashes, and individual operator entries.
type OperatorsData struct {
	LastProcessedEpoch uint64                     `json:"lastProcessedEpoch"`
	PendingHashes      []string                   `json:"pendingHashes"`
	OperatorDetails    map[string]OperatorDetails `json:"operatorDetails"` // indexed by operator ID
}

// OperatorDetails holds the data specific to each operator, including
// performance reports and exit requests.
type OperatorDetails struct {
	Performance  map[string]domain.Report      `json:"performance"`  // indexed by epoch
	ExitRequests map[string]domain.ExitRequest `json:"exitRequests"` // indexed by validator pubkey
}

func (fs *Storage) LoadDatabase() (Database, error) {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	// Initialize an empty Database with default values
	db := Database{
		Telegram: domain.TelegramConfig{},
		Operators: OperatorsData{
			LastProcessedEpoch: 0,                                // Initialize LastProcessedEpoch as 0
			PendingHashes:      []string{},                       // Initialize PendingHashes as an empty slice
			OperatorDetails:    make(map[string]OperatorDetails), // Initialize OperatorDetails as an empty map
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
	if db.Operators.OperatorDetails == nil {
		db.Operators.OperatorDetails = make(map[string]OperatorDetails)
	}
	if db.Operators.PendingHashes == nil {
		db.Operators.PendingHashes = []string{} // Initialize if nil after unmarshalling
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
