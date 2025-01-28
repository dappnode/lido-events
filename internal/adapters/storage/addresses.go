package storage

import (
	"fmt"
	"lido-events/internal/application/domain"

	"github.com/ethereum/go-ethereum/common"
)

// GetAddressLastProcessedBlock retrieves the last processed block for a specific Ethereum address.
func (fs *Storage) GetAddressLastProcessedBlock(address common.Address) (uint64, error) {
	db, err := fs.LoadDatabase()
	if err != nil {
		return 0, err
	}

	addressData, exists := db.Addresses[address]
	if !exists {
		return 0, fmt.Errorf("no data found for address %s", address)
	}

	return addressData.LastProcessedBlock, nil
}

// SaveAddressLastProcessedBlock saves the last processed block for a specific Ethereum address.
func (fs *Storage) SaveAddressLastProcessedBlock(address common.Address, block uint64) error {
	db, err := fs.LoadDatabase()
	if err != nil {
		return err
	}

	if db.Addresses == nil {
		db.Addresses = make(map[common.Address]domain.AddressEvents)
	}

	addressData := db.Addresses[address]
	addressData.LastProcessedBlock = block
	if addressData.NodeOperatorAdded == nil {
		addressData.NodeOperatorAdded = []domain.CsmoduleNodeOperatorAdded{}
	}
	if addressData.NodeOperatorManagerAddressChanged == nil {
		addressData.NodeOperatorManagerAddressChanged = []domain.CsmoduleNodeOperatorManagerAddressChanged{}
	}
	if addressData.NodeOperatorRewardAddressChanged == nil {
		addressData.NodeOperatorRewardAddressChanged = []domain.CsmoduleNodeOperatorRewardAddressChanged{}
	}
	db.Addresses[address] = addressData

	return fs.SaveDatabase(db)
}

// GetAddressEvents retrieves all Address events for a specific Ethereum address.
func (fs *Storage) GetAddressEvents(address common.Address) (domain.AddressEvents, error) {
	db, err := fs.LoadDatabase()
	if err != nil {
		return domain.AddressEvents{}, err
	}

	addressData, exists := db.Addresses[address]
	if !exists {
		return domain.AddressEvents{}, fmt.Errorf("no data found for address %s", address)
	}

	return domain.AddressEvents{
		LastProcessedBlock:                       addressData.LastProcessedBlock,
		NodeOperatorAdded:                        addressData.NodeOperatorAdded,
		NodeOperatorManagerAddressChanged:        addressData.NodeOperatorManagerAddressChanged,
		NodeOperatorRewardAddressChanged:         addressData.NodeOperatorRewardAddressChanged,
		NodeOperatorRewardAddressChangeProposed:  addressData.NodeOperatorRewardAddressChangeProposed,
		NodeOperatorManagerAddressChangeProposed: addressData.NodeOperatorManagerAddressChangeProposed,
	}, nil
}

// SetNodeOperatorAdded saves the NodeOperatorAdded event for a specific Ethereum address.
func (fs *Storage) SetNodeOperatorAdded(address common.Address, event domain.CsmoduleNodeOperatorAdded) error {
	db, err := fs.LoadDatabase()
	if err != nil {
		return err
	}

	if db.Addresses == nil {
		db.Addresses = make(map[common.Address]domain.AddressEvents)
	}

	// Initialize address if it does not exist
	if _, exists := db.Addresses[address]; !exists {
		db.Addresses[address] = domain.AddressEvents{
			LastProcessedBlock:                       0,
			NodeOperatorAdded:                        []domain.CsmoduleNodeOperatorAdded{},
			NodeOperatorManagerAddressChanged:        []domain.CsmoduleNodeOperatorManagerAddressChanged{},
			NodeOperatorRewardAddressChanged:         []domain.CsmoduleNodeOperatorRewardAddressChanged{},
			NodeOperatorRewardAddressChangeProposed:  []domain.CsmoduleNodeOperatorRewardAddressChangeProposed{},
			NodeOperatorManagerAddressChangeProposed: []domain.CsmoduleNodeOperatorManagerAddressChangeProposed{},
		}
	}

	addressData := db.Addresses[address]
	txHash := event.Raw.TxHash.Hex()

	// Check if the transaction hash already exists
	for _, existingEvent := range addressData.NodeOperatorAdded {
		if existingEvent.Raw.TxHash.Hex() == txHash {
			return nil // Event already exists
		}
	}

	// Add the new event
	addressData.NodeOperatorAdded = append(addressData.NodeOperatorAdded, event)
	db.Addresses[address] = addressData

	return fs.SaveDatabase(db)
}

// SetNodeOperatorManagerAddressChanged saves the NodeOperatorManagerAddressChanged event for a specific Ethereum address.
func (fs *Storage) SetNodeOperatorManagerAddressChanged(address common.Address, event domain.CsmoduleNodeOperatorManagerAddressChanged) error {
	db, err := fs.LoadDatabase()
	if err != nil {
		return err
	}

	if db.Addresses == nil {
		db.Addresses = make(map[common.Address]domain.AddressEvents)
	}

	// Initialize address if it does not exist
	if _, exists := db.Addresses[address]; !exists {
		db.Addresses[address] = domain.AddressEvents{
			LastProcessedBlock:                       0,
			NodeOperatorAdded:                        []domain.CsmoduleNodeOperatorAdded{},
			NodeOperatorManagerAddressChanged:        []domain.CsmoduleNodeOperatorManagerAddressChanged{},
			NodeOperatorRewardAddressChanged:         []domain.CsmoduleNodeOperatorRewardAddressChanged{},
			NodeOperatorRewardAddressChangeProposed:  []domain.CsmoduleNodeOperatorRewardAddressChangeProposed{},
			NodeOperatorManagerAddressChangeProposed: []domain.CsmoduleNodeOperatorManagerAddressChangeProposed{},
		}
	}

	addressData := db.Addresses[address]
	txHash := event.Raw.TxHash.Hex()

	// Check if the transaction hash already exists
	for _, existingEvent := range addressData.NodeOperatorManagerAddressChanged {
		if existingEvent.Raw.TxHash.Hex() == txHash {
			return nil // Event already exists
		}
	}

	// Add the new event
	addressData.NodeOperatorManagerAddressChanged = append(addressData.NodeOperatorManagerAddressChanged, event)
	db.Addresses[address] = addressData

	return fs.SaveDatabase(db)
}

// SetNodeOperatorRewardAddressChanged saves the NodeOperatorRewardAddressChanged event for a specific Ethereum address.
func (fs *Storage) SetNodeOperatorRewardAddressChanged(address common.Address, event domain.CsmoduleNodeOperatorRewardAddressChanged) error {
	db, err := fs.LoadDatabase()
	if err != nil {
		return err
	}

	if db.Addresses == nil {
		db.Addresses = make(map[common.Address]domain.AddressEvents)
	}

	// Initialize address if it does not exist
	if _, exists := db.Addresses[address]; !exists {
		db.Addresses[address] = domain.AddressEvents{
			LastProcessedBlock:                       0,
			NodeOperatorAdded:                        []domain.CsmoduleNodeOperatorAdded{},
			NodeOperatorManagerAddressChanged:        []domain.CsmoduleNodeOperatorManagerAddressChanged{},
			NodeOperatorRewardAddressChanged:         []domain.CsmoduleNodeOperatorRewardAddressChanged{},
			NodeOperatorRewardAddressChangeProposed:  []domain.CsmoduleNodeOperatorRewardAddressChangeProposed{},
			NodeOperatorManagerAddressChangeProposed: []domain.CsmoduleNodeOperatorManagerAddressChangeProposed{},
		}
	}

	addressData := db.Addresses[address]
	txHash := event.Raw.TxHash.Hex()

	// Check if the transaction hash already exists
	for _, existingEvent := range addressData.NodeOperatorRewardAddressChanged {
		if existingEvent.Raw.TxHash.Hex() == txHash {
			return nil // Event already exists
		}
	}

	// Add the new event
	addressData.NodeOperatorRewardAddressChanged = append(addressData.NodeOperatorRewardAddressChanged, event)
	db.Addresses[address] = addressData

	return fs.SaveDatabase(db)
}

// SetNodeOperatorRewardAddressChangeProposed saves the NodeOperatorRewardAddressChangeProposed event for a specific Ethereum address.
func (fs *Storage) SetNodeOperatorRewardAddressChangeProposed(address common.Address, event domain.CsmoduleNodeOperatorRewardAddressChangeProposed) error {
	db, err := fs.LoadDatabase()
	if err != nil {
		return err
	}

	if db.Addresses == nil {
		db.Addresses = make(map[common.Address]domain.AddressEvents)
	}

	// Initialize address if it does not exist
	if _, exists := db.Addresses[address]; !exists {
		db.Addresses[address] = domain.AddressEvents{
			LastProcessedBlock:                       0,
			NodeOperatorAdded:                        []domain.CsmoduleNodeOperatorAdded{},
			NodeOperatorManagerAddressChanged:        []domain.CsmoduleNodeOperatorManagerAddressChanged{},
			NodeOperatorRewardAddressChanged:         []domain.CsmoduleNodeOperatorRewardAddressChanged{},
			NodeOperatorRewardAddressChangeProposed:  []domain.CsmoduleNodeOperatorRewardAddressChangeProposed{},
			NodeOperatorManagerAddressChangeProposed: []domain.CsmoduleNodeOperatorManagerAddressChangeProposed{},
		}
	}

	addressData := db.Addresses[address]
	txHash := event.Raw.TxHash.Hex()

	// Check if the transaction hash already exists
	for _, existingEvent := range addressData.NodeOperatorRewardAddressChangeProposed {
		if existingEvent.Raw.TxHash.Hex() == txHash {
			return nil // Event already exists
		}
	}

	// Add the new event
	addressData.NodeOperatorRewardAddressChangeProposed = append(addressData.NodeOperatorRewardAddressChangeProposed, event)
	db.Addresses[address] = addressData

	return fs.SaveDatabase(db)
}

// SetNodeOperatorManagerAddressChangeProposed saves the NodeOperatorManagerAddressChangeProposed event for a specific Ethereum address.
func (fs *Storage) SetNodeOperatorManagerAddressChangeProposed(address common.Address, event domain.CsmoduleNodeOperatorManagerAddressChangeProposed) error {
	db, err := fs.LoadDatabase()
	if err != nil {
		return err
	}

	if db.Addresses == nil {
		db.Addresses = make(map[common.Address]domain.AddressEvents)
	}

	// Initialize address if it does not exist
	if _, exists := db.Addresses[address]; !exists {
		db.Addresses[address] = domain.AddressEvents{
			LastProcessedBlock:                       0,
			NodeOperatorAdded:                        []domain.CsmoduleNodeOperatorAdded{},
			NodeOperatorManagerAddressChanged:        []domain.CsmoduleNodeOperatorManagerAddressChanged{},
			NodeOperatorRewardAddressChanged:         []domain.CsmoduleNodeOperatorRewardAddressChanged{},
			NodeOperatorRewardAddressChangeProposed:  []domain.CsmoduleNodeOperatorRewardAddressChangeProposed{},
			NodeOperatorManagerAddressChangeProposed: []domain.CsmoduleNodeOperatorManagerAddressChangeProposed{},
		}
	}

	addressData := db.Addresses[address]
	txHash := event.Raw.TxHash.Hex()

	// Check if the transaction hash already exists
	for _, existingEvent := range addressData.NodeOperatorManagerAddressChangeProposed {
		if existingEvent.Raw.TxHash.Hex() == txHash {
			return nil // Event already exists
		}
	}

	// Add the new event
	addressData.NodeOperatorManagerAddressChangeProposed = append(addressData.NodeOperatorManagerAddressChangeProposed, event)
	db.Addresses[address] = addressData

	return fs.SaveDatabase(db)
}
