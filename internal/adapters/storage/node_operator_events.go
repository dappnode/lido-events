package storage

import (
	"fmt"
	"lido-events/internal/application/domain"
)

// GetNodeOperatorEvents retrieves all NodeOperator events for a specific operator ID.
func (fs *Storage) GetNodeOperatorEvents(operatorID string) (domain.NodeOperatorEvents, error) {
	db, err := fs.LoadDatabase()
	if err != nil {
		return domain.NodeOperatorEvents{}, err
	}

	opData, exists := db.Operators[operatorID]
	if !exists {
		return domain.NodeOperatorEvents{}, fmt.Errorf("no data found for operator ID %s", operatorID)
	}

	return opData.NodeOperatorEvents, nil
}

// SetNodeOperatorAdded saves the NodeOperatorAdded event for a specific operator ID.
func (fs *Storage) SetNodeOperatorAdded(operatorID string, event domain.CsmoduleNodeOperatorAdded) error {
	db, err := fs.LoadDatabase()
	if err != nil {
		return err
	}

	if db.Operators == nil {
		db.Operators = make(map[string]OperatorData)
	}

	opData := db.Operators[operatorID]

	// Use transaction hash from the event's Raw log for comparison
	txHash := event.Raw.TxHash.Hex()

	// Check if the transaction hash already exists in the event array
	for _, existingEvent := range opData.NodeOperatorEvents.NodeOperatorAdded {
		if existingEvent.Raw.TxHash.Hex() == txHash {
			// Event already exists; no need to add it
			return nil
		}
	}

	// Add the new event to the array
	opData.NodeOperatorEvents.NodeOperatorAdded = append(opData.NodeOperatorEvents.NodeOperatorAdded, event)
	db.Operators[operatorID] = opData

	return fs.SaveDatabase(db)
}

// SetNodeOperatorManagerAddressChanged saves the NodeOperatorManagerAddressChanged event for a specific operator ID.
func (fs *Storage) SetNodeOperatorManagerAddressChanged(operatorID string, event domain.CsmoduleNodeOperatorManagerAddressChanged) error {
	db, err := fs.LoadDatabase()
	if err != nil {
		return err
	}

	if db.Operators == nil {
		db.Operators = make(map[string]OperatorData)
	}

	opData := db.Operators[operatorID]

	// Use transaction hash from the event's Raw log for comparison
	txHash := event.Raw.TxHash.Hex()

	// Check if the transaction hash already exists in the event array
	for _, existingEvent := range opData.NodeOperatorEvents.NodeOperatorManagerAddressChanged {
		if existingEvent.Raw.TxHash.Hex() == txHash {
			// Event already exists; no need to add it
			return nil
		}
	}

	// Add the new event to the array
	opData.NodeOperatorEvents.NodeOperatorManagerAddressChanged = append(opData.NodeOperatorEvents.NodeOperatorManagerAddressChanged, event)
	db.Operators[operatorID] = opData

	return fs.SaveDatabase(db)
}

// SetNodeOperatorRewardAddressChanged saves the NodeOperatorRewardAddressChanged event for a specific operator ID.
func (fs *Storage) SetNodeOperatorRewardAddressChanged(operatorID string, event domain.CsmoduleNodeOperatorRewardAddressChanged) error {
	db, err := fs.LoadDatabase()
	if err != nil {
		return err
	}

	if db.Operators == nil {
		db.Operators = make(map[string]OperatorData)
	}

	opData := db.Operators[operatorID]

	// Use transaction hash from the event's Raw log for comparison
	txHash := event.Raw.TxHash.Hex()

	// Check if the transaction hash already exists in the event array
	for _, existingEvent := range opData.NodeOperatorEvents.NodeOperatorRewardAddressChanged {
		if existingEvent.Raw.TxHash.Hex() == txHash {
			// Event already exists; no need to add it
			return nil
		}
	}

	// Add the new event to the array
	opData.NodeOperatorEvents.NodeOperatorRewardAddressChanged = append(opData.NodeOperatorEvents.NodeOperatorRewardAddressChanged, event)
	db.Operators[operatorID] = opData

	return fs.SaveDatabase(db)
}
