package storage

import (
	"fmt"
	"lido-events/internal/application/domain"

	"github.com/ethereum/go-ethereum/common"
)

// Migration defines the function signature for migrations
type Migration func(*Database) error

// MigrationRegistry stores migrations mapped by version number
var MigrationRegistry = map[int]Migration{
	1: migrateToVersion1,
}

// RunMigrations runs all required migrations sequentially
func RunMigrations(storage *Storage) error {
	// Load the current database
	db, err := storage.LoadDatabase()
	if err != nil {
		return fmt.Errorf("failed to load database for migration: %v", err)
	}

	// Default version if not set
	if db.Version == 0 {
		db.Version = 1 // Assume the initial version is 1
	}

	// Apply all necessary migrations
	for version, migration := range MigrationRegistry {
		if db.Version < version {
			fmt.Printf("Applying migration: version %d\n", version)
			err := migration(&db)
			if err != nil {
				return fmt.Errorf("migration to version %d failed: %v", version, err)
			}
			db.Version = version
		}
	}

	// Save the migrated database
	err = storage.SaveDatabase(db)
	if err != nil {
		return fmt.Errorf("failed to save migrated database: %v", err)
	}

	fmt.Println("Database migrations completed successfully.")
	return nil
}

// migrateToVersion1 removes all data except Telegram config and Operator IDs
func migrateToVersion1(db *Database) error {
	fmt.Println("Performing migration to version 1...")

	// Extract only Telegram config and Operator IDs
	newOperators := make(map[string]OperatorData)
	for id := range db.Operators {
		newOperators[id] = OperatorData{}
	}

	// Reset the database with only these fields
	*db = Database{
		Version:                            1,
		Telegram:                           db.Telegram,
		Operators:                          newOperators,
		Addresses:                          make(map[common.Address]domain.AddressEvents),
		ElRewardsStealingPenaltiesReported: ElRewardsStealingPenaltiesReported{},
	}

	return nil
}
