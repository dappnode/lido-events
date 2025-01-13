package storage

import (
	"fmt"
	"lido-events/internal/application/domain"
)

// GetNodeOperatorAdded retrieves the NodeOperatorAdded event for a specific operator ID.
func (fs *Storage) GetNodeOperatorAdded(operatorID string) (domain.CsmoduleNodeOperatorAdded, error) {
	db, err := fs.LoadDatabase()
	if err != nil {
		return domain.CsmoduleNodeOperatorAdded{}, err
	}

	opData, exists := db.Operators[operatorID]
	if !exists {
		return domain.CsmoduleNodeOperatorAdded{}, fmt.Errorf("no data found for operator ID %s", operatorID)
	}

	return opData.NodeOperatorEvents.NodeOperatorAdded, nil
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
	opData.NodeOperatorEvents.NodeOperatorAdded = event
	db.Operators[operatorID] = opData

	return fs.SaveDatabase(db)
}

// GetNodeOperatorManagerAddressChanged retrieves the NodeOperatorManagerAddressChanged event for a specific operator ID.
func (fs *Storage) GetNodeOperatorManagerAddressChanged(operatorID string) (domain.CsmoduleNodeOperatorManagerAddressChanged, error) {
	db, err := fs.LoadDatabase()
	if err != nil {
		return domain.CsmoduleNodeOperatorManagerAddressChanged{}, err
	}

	opData, exists := db.Operators[operatorID]
	if !exists {
		return domain.CsmoduleNodeOperatorManagerAddressChanged{}, fmt.Errorf("no data found for operator ID %s", operatorID)
	}

	return opData.NodeOperatorEvents.NodeOperatorManagerAddressChanged, nil
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
	opData.NodeOperatorEvents.NodeOperatorManagerAddressChanged = event
	db.Operators[operatorID] = opData

	return fs.SaveDatabase(db)
}

// GetNodeOperatorRewardAddressChanged retrieves the NodeOperatorRewardAddressChanged event for a specific operator ID.
func (fs *Storage) GetNodeOperatorRewardAddressChanged(operatorID string) (domain.CsmoduleNodeOperatorRewardAddressChanged, error) {
	db, err := fs.LoadDatabase()
	if err != nil {
		return domain.CsmoduleNodeOperatorRewardAddressChanged{}, err
	}

	opData, exists := db.Operators[operatorID]
	if !exists {
		return domain.CsmoduleNodeOperatorRewardAddressChanged{}, fmt.Errorf("no data found for operator ID %s", operatorID)
	}

	return opData.NodeOperatorEvents.NodeOperatorRewardAddressChanged, nil
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
	opData.NodeOperatorEvents.NodeOperatorRewardAddressChanged = event
	db.Operators[operatorID] = opData

	return fs.SaveDatabase(db)
}
