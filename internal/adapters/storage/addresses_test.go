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

func TestGetAddressLastProcessedBlock(t *testing.T) {

	initialData := &storage.Database{
		Addresses: map[common.Address]domain.AddressEvents{
			common.HexToAddress("0x1111111111111111111111111111111111111111"): {
				LastProcessedBlock: 12345,
			},
		},
	}
	tmpFile := CreateTempDatabaseFile(t, initialData)
	defer os.Remove(tmpFile.Name())

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	address := common.HexToAddress("0x1111111111111111111111111111111111111111")

	block, err := storageAdapter.GetAddressLastProcessedBlock(address)
	assert.NoError(t, err)
	assert.Equal(t, initialData.Addresses[address].LastProcessedBlock, block)
}

func TestSaveAddressLastProcessedBlock(t *testing.T) {
	tmpFile := CreateTempDatabaseFile(t, nil)
	defer os.Remove(tmpFile.Name())

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	address := common.HexToAddress("0x1111111111111111111111111111111111111111")

	err := storageAdapter.SaveAddressLastProcessedBlock(address, 12345)
	assert.NoError(t, err)

	db, err := storageAdapter.LoadDatabase()
	assert.NoError(t, err)
	assert.Equal(t, uint64(12345), db.Addresses[address].LastProcessedBlock)
}

func TestSetNodeOperatorAdded(t *testing.T) {
	tmpFile := CreateTempDatabaseFile(t, nil)
	defer os.Remove(tmpFile.Name())

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	address := common.HexToAddress("0x1111111111111111111111111111111111111111")
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

	err := storageAdapter.SetNodeOperatorAdded(address, event)
	assert.NoError(t, err)

	db, err := storageAdapter.LoadDatabase()
	assert.NoError(t, err)
	assert.Contains(t, db.Addresses[address].NodeOperatorAdded, event)
}

func TestSetNodeOperatorAdded_DuplicateEvent(t *testing.T) {
	tmpFile := CreateTempDatabaseFile(t, nil)
	defer os.Remove(tmpFile.Name())

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	address := common.HexToAddress("0x1111111111111111111111111111111111111111")
	event := domain.CsmoduleNodeOperatorAdded{
		NodeOperatorId: big.NewInt(1),
		ManagerAddress: common.HexToAddress("0x1111111111111111111111111111111111111111"),
		RewardAddress:  common.HexToAddress("0x2222222222222222222222222222222222222222"),
		Raw: types.Log{
			TxHash: common.HexToHash("0x1234567890abcdef"),
			Topics: []common.Hash{},
		},
	}

	err := storageAdapter.SetNodeOperatorAdded(address, event)
	assert.NoError(t, err)

	err = storageAdapter.SetNodeOperatorAdded(address, event)
	assert.NoError(t, err)

	db, err := storageAdapter.LoadDatabase()
	assert.NoError(t, err)

	assert.Len(t, db.Addresses[address].NodeOperatorAdded, 1)
}

func TestSetNodeOperatorAdded_UninitializedStructures(t *testing.T) {
	initialData := &storage.Database{}
	tmpFile := CreateTempDatabaseFile(t, initialData)
	defer os.Remove(tmpFile.Name())

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	address := common.HexToAddress("0x1111111111111111111111111111111111111111")
	event := domain.CsmoduleNodeOperatorAdded{
		NodeOperatorId: big.NewInt(1),
		ManagerAddress: common.HexToAddress("0x1111111111111111111111111111111111111111"),
		RewardAddress:  common.HexToAddress("0x2222222222222222222222222222222222222222"),
		Raw: types.Log{
			TxHash: common.HexToHash("0xabcdef1234567890"),
			Topics: []common.Hash{},
		},
	}

	err := storageAdapter.SetNodeOperatorAdded(address, event)
	assert.NoError(t, err)

	db, err := storageAdapter.LoadDatabase()
	assert.NoError(t, err)

	assert.Contains(t, db.Addresses, address)
	assert.NotNil(t, db.Addresses[address].NodeOperatorAdded)
	assert.Len(t, db.Addresses[address].NodeOperatorAdded, 1)
}

func TestSetNodeOperatorManagerAddressChanged(t *testing.T) {
	tmpFile := CreateTempDatabaseFile(t, nil)
	defer os.Remove(tmpFile.Name())

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	address := common.HexToAddress("0x1111111111111111111111111111111111111111")
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

	err := storageAdapter.SetNodeOperatorManagerAddressChanged(address, event)
	assert.NoError(t, err)

	db, err := storageAdapter.LoadDatabase()
	assert.NoError(t, err)
	assert.Contains(t, db.Addresses[address].NodeOperatorManagerAddressChanged, event)
}

func TestSetNodeOperatorRewardAddressChanged(t *testing.T) {
	tmpFile := CreateTempDatabaseFile(t, nil)
	defer os.Remove(tmpFile.Name())

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	address := common.HexToAddress("0x1111111111111111111111111111111111111111")
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

	err := storageAdapter.SetNodeOperatorRewardAddressChanged(address, event)
	assert.NoError(t, err)

	db, err := storageAdapter.LoadDatabase()
	assert.NoError(t, err)
	assert.Contains(t, db.Addresses[address].NodeOperatorRewardAddressChanged, event)
}

func TestGetAddressEvents(t *testing.T) {
	initialData := &storage.Database{
		Addresses: map[common.Address]domain.AddressEvents{
			common.HexToAddress("0x1111111111111111111111111111111111111111"): {
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
	}
	tmpFile := CreateTempDatabaseFile(t, initialData)
	defer os.Remove(tmpFile.Name())

	storageAdapter := &storage.Storage{DBFile: tmpFile.Name()}
	address := common.HexToAddress("0x1111111111111111111111111111111111111111")

	events, err := storageAdapter.GetAddressEvents(address)
	assert.NoError(t, err)

	expectedEvents := initialData.Addresses[address]
	assert.Equal(t, expectedEvents.NodeOperatorAdded, events.NodeOperatorAdded)
	assert.Equal(t, expectedEvents.NodeOperatorManagerAddressChanged, events.NodeOperatorManagerAddressChanged)
	assert.Equal(t, expectedEvents.NodeOperatorRewardAddressChanged, events.NodeOperatorRewardAddressChanged)
}
