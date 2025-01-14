//go:build integration
// +build integration

package storage_test

import (
	"lido-events/internal/adapters/storage"
	"lido-events/internal/application/domain"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
)

// TestSetNodeOperatorAdded tests saving a NodeOperatorAdded event.
func TestSetNodeOperatorAdded(t *testing.T) {
	tmpFile := CreateTempDatabaseFile(t, nil)
	defer os.Remove(tmpFile.Name())

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	operatorID := "1"
	event := domain.CsmoduleNodeOperatorAdded{
		NodeOperatorId: big.NewInt(1),
		ManagerAddress: common.HexToAddress("0x1111111111111111111111111111111111111111"),
		RewardAddress:  common.HexToAddress("0x2222222222222222222222222222222222222222"),
		Raw: types.Log{
			Address: common.HexToAddress("0x0000000000000000000000000000000000000000"),
			Topics:  []common.Hash{},
			Data:    []byte{},
		},
	}

	err := storageAdapter.SetNodeOperatorAdded(operatorID, event)
	assert.NoError(t, err)

	db, err := storageAdapter.LoadDatabase()
	assert.NoError(t, err)
	assert.Contains(t, db.Operators[operatorID].NodeOperatorEvents.NodeOperatorAdded, event)
}

// TestGetNodeOperatorAdded tests retrieving a NodeOperatorAdded event.
func TestGetNodeOperatorAdded(t *testing.T) {
	initialData := &storage.Database{
		Operators: map[string]storage.OperatorData{
			"1": {
				NodeOperatorEvents: storage.NodeOperatorEvents{
					NodeOperatorAdded: []domain.CsmoduleNodeOperatorAdded{
						{
							NodeOperatorId: big.NewInt(1),
							ManagerAddress: common.HexToAddress("0x1111111111111111111111111111111111111111"),
							RewardAddress:  common.HexToAddress("0x2222222222222222222222222222222222222222"),
							Raw: types.Log{
								Address: common.HexToAddress("0x0000000000000000000000000000000000000000"),
								Topics:  []common.Hash{},
								Data:    []byte{},
							},
						},
					},
				},
			},
		},
	}

	tmpFile := CreateTempDatabaseFile(t, initialData)
	defer os.Remove(tmpFile.Name())

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	operatorID := "1"

	events, err := storageAdapter.GetNodeOperatorAdded(operatorID)
	assert.NoError(t, err)

	expectedEvents := initialData.Operators[operatorID].NodeOperatorEvents.NodeOperatorAdded
	assert.Equal(t, expectedEvents, events)
}

// TestSetNodeOperatorManagerAddressChanged tests saving a NodeOperatorManagerAddressChanged event.
func TestSetNodeOperatorManagerAddressChanged(t *testing.T) {
	tmpFile := CreateTempDatabaseFile(t, nil)
	defer os.Remove(tmpFile.Name())

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	operatorID := "1"
	event := domain.CsmoduleNodeOperatorManagerAddressChanged{
		NodeOperatorId: big.NewInt(1),
		OldAddress:     common.HexToAddress("0x3333333333333333333333333333333333333333"),
		NewAddress:     common.HexToAddress("0x4444444444444444444444444444444444444444"),
		Raw: types.Log{
			Address: common.HexToAddress("0x0000000000000000000000000000000000000000"),
			Topics:  []common.Hash{},
			Data:    []byte{},
		},
	}

	err := storageAdapter.SetNodeOperatorManagerAddressChanged(operatorID, event)
	assert.NoError(t, err)

	db, err := storageAdapter.LoadDatabase()
	assert.NoError(t, err)
	assert.Contains(t, db.Operators[operatorID].NodeOperatorEvents.NodeOperatorManagerAddressChanged, event)
}

// TestGetNodeOperatorManagerAddressChanged tests retrieving a NodeOperatorManagerAddressChanged event.
func TestGetNodeOperatorManagerAddressChanged(t *testing.T) {
	initialData := &storage.Database{
		Operators: map[string]storage.OperatorData{
			"1": {
				NodeOperatorEvents: storage.NodeOperatorEvents{
					NodeOperatorManagerAddressChanged: []domain.CsmoduleNodeOperatorManagerAddressChanged{
						{
							NodeOperatorId: big.NewInt(1),
							OldAddress:     common.HexToAddress("0x3333333333333333333333333333333333333333"),
							NewAddress:     common.HexToAddress("0x4444444444444444444444444444444444444444"),
							Raw: types.Log{
								Address: common.HexToAddress("0x0000000000000000000000000000000000000000"),
								Topics:  []common.Hash{},
								Data:    []byte{},
							},
						},
					},
				},
			},
		},
	}

	tmpFile := CreateTempDatabaseFile(t, initialData)
	defer os.Remove(tmpFile.Name())

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	operatorID := "1"

	events, err := storageAdapter.GetNodeOperatorManagerAddressChanged(operatorID)
	assert.NoError(t, err)

	expectedEvents := initialData.Operators[operatorID].NodeOperatorEvents.NodeOperatorManagerAddressChanged
	assert.Equal(t, expectedEvents, events)
}

// TestSetNodeOperatorRewardAddressChanged tests saving a NodeOperatorRewardAddressChanged event.
func TestSetNodeOperatorRewardAddressChanged(t *testing.T) {
	tmpFile := CreateTempDatabaseFile(t, nil)
	defer os.Remove(tmpFile.Name())

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	operatorID := "1"
	event := domain.CsmoduleNodeOperatorRewardAddressChanged{
		NodeOperatorId: big.NewInt(1),
		OldAddress:     common.HexToAddress("0x5555555555555555555555555555555555555555"),
		NewAddress:     common.HexToAddress("0x6666666666666666666666666666666666666666"),
		Raw: types.Log{
			Address: common.HexToAddress("0x0000000000000000000000000000000000000000"),
			Topics:  []common.Hash{},
			Data:    []byte{},
		},
	}

	err := storageAdapter.SetNodeOperatorRewardAddressChanged(operatorID, event)
	assert.NoError(t, err)

	db, err := storageAdapter.LoadDatabase()
	assert.NoError(t, err)
	assert.Contains(t, db.Operators[operatorID].NodeOperatorEvents.NodeOperatorRewardAddressChanged, event)
}

// TestGetNodeOperatorRewardAddressChanged tests retrieving a NodeOperatorRewardAddressChanged event.
func TestGetNodeOperatorRewardAddressChanged(t *testing.T) {
	initialData := &storage.Database{
		Operators: map[string]storage.OperatorData{
			"1": {
				NodeOperatorEvents: storage.NodeOperatorEvents{
					NodeOperatorRewardAddressChanged: []domain.CsmoduleNodeOperatorRewardAddressChanged{
						{
							NodeOperatorId: big.NewInt(1),
							OldAddress:     common.HexToAddress("0x5555555555555555555555555555555555555555"),
							NewAddress:     common.HexToAddress("0x6666666666666666666666666666666666666666"),
							Raw: types.Log{
								Address: common.HexToAddress("0x0000000000000000000000000000000000000000"),
								Topics:  []common.Hash{},
								Data:    []byte{},
							},
						},
					},
				},
			},
		},
	}

	tmpFile := CreateTempDatabaseFile(t, initialData)
	defer os.Remove(tmpFile.Name())

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	operatorID := "1"

	events, err := storageAdapter.GetNodeOperatorRewardAddressChanged(operatorID)
	assert.NoError(t, err)

	expectedEvents := initialData.Operators[operatorID].NodeOperatorEvents.NodeOperatorRewardAddressChanged
	assert.Equal(t, expectedEvents, events)
}
