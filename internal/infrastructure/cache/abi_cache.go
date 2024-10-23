package cache

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type ABICache struct {
	cache map[string]interface{} // Cache of contract addresses to parsed ABIs
}

// NewABICache initializes the ABI cache with the given paths
func NewABICache(abiPaths map[string]string) (*ABICache, error) {
	cache := make(map[string]interface{})

	for address, path := range abiPaths {
		abi, err := loadABIFromFile(path)
		if err != nil {
			return nil, fmt.Errorf("failed to load ABI for contract %s from %s: %w", address, path, err)
		}
		cache[address] = abi
	}

	return &ABICache{cache: cache}, nil
}

func (a *ABICache) GetAllABIs() map[string]interface{} {
	return a.cache
}

// loadABIFromFile reads and unmarshals the ABI from a file
func loadABIFromFile(filePath string) (interface{}, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file) // Replaced ioutil.ReadAll with io.ReadAll
	if err != nil {
		return nil, err
	}

	var abi interface{}
	if err := json.Unmarshal(byteValue, &abi); err != nil {
		return nil, err
	}

	return abi, nil
}
