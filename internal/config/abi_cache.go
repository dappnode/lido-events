package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

// This component manages the caching of Lido Contract ABIs. It loads the ABIs from the
// given path and stores them in a map for easy access. It is a Support Tool for other adapters
// like EthereumSubscriber
type ABICache struct {
	cache map[string]interface{} // Cache of contract addresses to parsed ABIs
}

// NewABICache initializes the ABI cache with the given paths
func NewABICache(abiPaths map[string]string) (*ABICache, error) {
	cache := make(map[string]interface{})

	for address, path := range abiPaths {
		parsedABI, err := loadABIFromFile(path)
		if err != nil {
			return nil, fmt.Errorf("failed to load ABI for contract %s from %s: %w", address, path, err)
		}
		cache[address] = parsedABI
	}

	return &ABICache{cache: cache}, nil
}

func (a *ABICache) GetAllABIs() map[string]interface{} {
	return a.cache
}

// loadABIFromFile reads and unmarshals the ABI from a file
func loadABIFromFile(path string) (abi.ABI, error) {
	abiJSON, err := os.ReadFile(path)
	if err != nil {
		return abi.ABI{}, fmt.Errorf("failed to read ABI file: %w", err)
	}

	var parsedABI abi.ABI
	if err := json.Unmarshal(abiJSON, &parsedABI); err != nil {
		return abi.ABI{}, fmt.Errorf("failed to parse ABI JSON: %w", err)
	}

	return parsedABI, nil
}
