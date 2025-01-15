package storage

import (
	"fmt"
	"lido-events/internal/application/domain"
)

// GetAddressEvents retrieves all Address events for a specific Ethereum address.
func (fs *Storage) GetAddressEvents(address string) (domain.AddressEvents, error) {
	db, err := fs.LoadDatabase()
	if err != nil {
		return domain.AddressEvents{}, err
	}

	addressData, exists := db.Addresses[address]
	if !exists {
		return domain.AddressEvents{}, fmt.Errorf("no data found for address %s", address)
	}

	return domain.AddressEvents{
		NodeOperatorAdded:                 addressData.NodeOperatorAdded,
		NodeOperatorManagerAddressChanged: addressData.NodeOperatorManagerAddressChanged,
		NodeOperatorRewardAddressChanged:  addressData.NodeOperatorRewardAddressChanged,
	}, nil
}

// SetNodeOperatorAdded saves the NodeOperatorAdded event for a specific Ethereum address.
func (fs *Storage) SetNodeOperatorAdded(address string, event domain.CsmoduleNodeOperatorAdded) error {
	db, err := fs.LoadDatabase()
	if err != nil {
		return err
	}

	if db.Addresses == nil {
		db.Addresses = make(map[string]domain.AddressEvents)
	}

	addressData := db.Addresses[address]

	// Use transaction hash from the event's Raw log for comparison
	txHash := event.Raw.TxHash.Hex()

	// Check if the transaction hash already exists in the event array
	for _, existingEvent := range addressData.NodeOperatorAdded {
		if existingEvent.Raw.TxHash.Hex() == txHash {
			// Event already exists; no need to add it
			return nil
		}
	}

	// Add the new event to the array
	addressData.NodeOperatorAdded = append(addressData.NodeOperatorAdded, event)
	db.Addresses[address] = addressData

	return fs.SaveDatabase(db)
}

// SetNodeOperatorManagerAddressChanged saves the NodeOperatorManagerAddressChanged event for a specific Ethereum address.
func (fs *Storage) SetNodeOperatorManagerAddressChanged(address string, event domain.CsmoduleNodeOperatorManagerAddressChanged) error {
	db, err := fs.LoadDatabase()
	if err != nil {
		return err
	}

	if db.Addresses == nil {
		db.Addresses = make(map[string]domain.AddressEvents)
	}

	addressData := db.Addresses[address]

	// Use transaction hash from the event's Raw log for comparison
	txHash := event.Raw.TxHash.Hex()

	// Check if the transaction hash already exists in the event array
	for _, existingEvent := range addressData.NodeOperatorManagerAddressChanged {
		if existingEvent.Raw.TxHash.Hex() == txHash {
			// Event already exists; no need to add it
			return nil
		}
	}

	// Add the new event to the array
	addressData.NodeOperatorManagerAddressChanged = append(addressData.NodeOperatorManagerAddressChanged, event)
	db.Addresses[address] = addressData

	return fs.SaveDatabase(db)
}

// SetNodeOperatorRewardAddressChanged saves the NodeOperatorRewardAddressChanged event for a specific Ethereum address.
func (fs *Storage) SetNodeOperatorRewardAddressChanged(address string, event domain.CsmoduleNodeOperatorRewardAddressChanged) error {
	db, err := fs.LoadDatabase()
	if err != nil {
		return err
	}

	if db.Addresses == nil {
		db.Addresses = make(map[string]domain.AddressEvents)
	}

	addressData := db.Addresses[address]

	// Use transaction hash from the event's Raw log for comparison
	txHash := event.Raw.TxHash.Hex()

	// Check if the transaction hash already exists in the event array
	for _, existingEvent := range addressData.NodeOperatorRewardAddressChanged {
		if existingEvent.Raw.TxHash.Hex() == txHash {
			// Event already exists; no need to add it
			return nil
		}
	}

	// Add the new event to the array
	addressData.NodeOperatorRewardAddressChanged = append(addressData.NodeOperatorRewardAddressChanged, event)
	db.Addresses[address] = addressData

	return fs.SaveDatabase(db)
}
