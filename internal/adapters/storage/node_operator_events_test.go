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

// GetNodeOperatorEvents tests retrieving all NodeOperator events.
func TestGetNodeOperatorEvents(t *testing.T) {
	initialData := &storage.Database{
		Operators: map[string]storage.OperatorData{
			"1": {
				NodeOperatorEvents: domain.NodeOperatorEvents{
					NodeOperatorAdded: []domain.CsmoduleNodeOperatorAdded{
						{
							NodeOperatorId: big.NewInt(1),
							ManagerAddress: common.HexToAddress("0x0000000000000000000000000000000000000000"),
							RewardAddress:  common.HexToAddress("0x0000000000000000000000000000000000000000"),
							Raw: types.Log{
								Address: common.HexToAddress("0x0000000000000000000000000000000000000000"),
								Topics:  []common.Hash{},
								Data:    []byte{},
							},
						},
					},
					NodeOperatorManagerAddressChanged: []domain.CsmoduleNodeOperatorManagerAddressChanged{
						{
							NodeOperatorId: big.NewInt(1),
							OldAddress:     common.HexToAddress("0x0000000000000000000000000000000000000000"),
							NewAddress:     common.HexToAddress("0x0000000000000000000000000000000000000000"),
							Raw: types.Log{
								Address: common.HexToAddress("0x0000000000000000000000000000000000000000"),
								Topics:  []common.Hash{},
								Data:    []byte{},
							},
						},
					},
					NodeOperatorRewardAddressChanged: []domain.CsmoduleNodeOperatorRewardAddressChanged{
						{
							NodeOperatorId: big.NewInt(1),
							OldAddress:     common.HexToAddress("0x0000000000000000000000000000000000000000"),
							NewAddress:     common.HexToAddress("0x0000000000000000000000000000000000000000"),
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

	events, err := storageAdapter.GetNodeOperatorEvents(operatorID)
	assert.NoError(t, err)

	expectedEvents := initialData.Operators[operatorID].NodeOperatorEvents
	assert.Equal(t, expectedEvents, events)
}
