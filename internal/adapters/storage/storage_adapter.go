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

	storage := &Storage{
		DBFile: dbFile,
	}

	// Run migrations
	if err := RunMigrations(storage); err != nil {
		return nil, fmt.Errorf("database migration failed: %w", err)
	}

	return storage, nil
}

// Database structure for the new storage format
type Database struct {
	Version                            int                                     `json:"version"`
	Telegram                           domain.TelegramConfig                   `json:"telegram"`
	Operators                          map[string]OperatorData                 `json:"operators"`                          // indexed by operator ID
	Addresses                          map[common.Address]domain.AddressEvents `json:"addresses"`                          // indexed by Ethereum address
	ElRewardsStealingPenaltiesReported ElRewardsStealingPenaltiesReported      `json:"elRewardsStealingPenaltiesReported"` // not indexed by anything
	ProcessingStarted                  ProcessingStarted                       `json:"processingStarted"`                  // not indexed by anything
}

type OperatorData struct {
	DistributionLogsUpdated DistributionLogsUpdated `json:"distributionLogsUpdated"`
	ExitsRequests           ExitsRequests           `json:"exitsRequests"`
	WithdrawalsSubmitted    WithdrawalsSubmitted    `json:"withdrawalsSubmitted"`
}

type DistributionLogsUpdated struct {
	LastProcessedBlock uint64   `json:"lastProcessedBlock"`
	PendingHashes      []string `json:"pendingHashes"`
	Reports            domain.Reports
}

type ExitsRequests struct {
	LastProcessedBlock uint64              `json:"lastProcessedBlock"`
	Exits              domain.ExitRequests `json:"exits"`
}

type WithdrawalsSubmitted struct {
	LastProcessedBlock uint64                               `json:"lastProcessedBlock"`
	Withdrawals        []domain.CsmoduleWithdrawalSubmitted `json:"withdrawals"`
}

type ElRewardsStealingPenaltiesReported struct {
	LastProcessedBlock uint64                                            `json:"lastProcessedBlock"`
	Penalties          []domain.CsmoduleELRewardsStealingPenaltyReported `json:"penalties"`
}

type ProcessingStarted struct {
	LastProcessedBlock uint64                             `json:"lastProcessedBlock"`
	Events             []domain.BindingsProcessingStarted `json:"events"`
}

func (fs *Storage) LoadDatabase() (Database, error) {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	// Initialize an empty Database with default values
	db := Database{
		Version:   1,
		Telegram:  domain.TelegramConfig{},
		Operators: make(map[string]OperatorData),
		Addresses: make(map[common.Address]domain.AddressEvents),
		ElRewardsStealingPenaltiesReported: ElRewardsStealingPenaltiesReported{
			Penalties:          []domain.CsmoduleELRewardsStealingPenaltyReported{},
			LastProcessedBlock: 0,
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
