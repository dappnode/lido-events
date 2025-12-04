// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ICSFeeDistributorDistributionData is an auto generated low-level Go binding around an user-defined struct.
type ICSFeeDistributorDistributionData struct {
	RefSlot     *big.Int
	TreeRoot    [32]byte
	TreeCid     string
	LogCid      string
	Distributed *big.Int
	Rebate      *big.Int
}

// BindingsMetaData contains all meta data concerning the Bindings contract.
var BindingsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stETH\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"accounting\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedToSendEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeSharesDecrease\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLogCID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidReportData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShares\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTreeCid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTreeRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllowedToRecover\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughShares\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderIsNotAccounting\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderIsNotOracle\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAccountingAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAdminAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroOracleAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroRebateRecipientAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroStEthAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalClaimableShares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"treeRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"treeCid\",\"type\":\"string\"}],\"name\":\"DistributionDataUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"logCid\",\"type\":\"string\"}],\"name\":\"DistributionLogUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC1155Recovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC20Recovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"ERC721Recovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EtherRecovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"ModuleFeeDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"OperatorFeeDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"RebateRecipientSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"RebateTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"StETHSharesRecovered\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ACCOUNTING\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ORACLE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RECOVERER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STETH\",\"outputs\":[{\"internalType\":\"contractIStETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cumulativeFeeShares\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"distributeFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"sharesToDistribute\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"distributedShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"distributed\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributionDataHistoryCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rebateRecipient\",\"type\":\"address\"}],\"name\":\"finalizeUpgradeV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cumulativeFeeShares\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"getFeesToDistribute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"sharesToDistribute\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getHistoricalDistributionData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"treeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"treeCid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"logCid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"distributed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rebate\",\"type\":\"uint256\"}],\"internalType\":\"structICSFeeDistributor.DistributionData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInitializedVersion\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"hashLeaf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rebateRecipient\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"logCid\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingSharesToDistribute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_treeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_treeCid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_logCid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"distributed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rebate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"}],\"name\":\"processOracleReport\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rebateRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"recoverERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"recoverERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"recoverERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recoverEther\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rebateRecipient\",\"type\":\"address\"}],\"name\":\"setRebateRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalClaimableShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treeCid\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treeRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BindingsABI is the input ABI used to generate the binding from.
// Deprecated: Use BindingsMetaData.ABI instead.
var BindingsABI = BindingsMetaData.ABI

// Bindings is an auto generated Go binding around an Ethereum contract.
type Bindings struct {
	BindingsCaller     // Read-only binding to the contract
	BindingsTransactor // Write-only binding to the contract
	BindingsFilterer   // Log filterer for contract events
}

// BindingsCaller is an auto generated read-only Go binding around an Ethereum contract.
type BindingsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BindingsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BindingsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BindingsSession struct {
	Contract     *Bindings         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BindingsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BindingsCallerSession struct {
	Contract *BindingsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BindingsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BindingsTransactorSession struct {
	Contract     *BindingsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BindingsRaw is an auto generated low-level Go binding around an Ethereum contract.
type BindingsRaw struct {
	Contract *Bindings // Generic contract binding to access the raw methods on
}

// BindingsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BindingsCallerRaw struct {
	Contract *BindingsCaller // Generic read-only contract binding to access the raw methods on
}

// BindingsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BindingsTransactorRaw struct {
	Contract *BindingsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBindings creates a new instance of Bindings, bound to a specific deployed contract.
func NewBindings(address common.Address, backend bind.ContractBackend) (*Bindings, error) {
	contract, err := bindBindings(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bindings{BindingsCaller: BindingsCaller{contract: contract}, BindingsTransactor: BindingsTransactor{contract: contract}, BindingsFilterer: BindingsFilterer{contract: contract}}, nil
}

// NewBindingsCaller creates a new read-only instance of Bindings, bound to a specific deployed contract.
func NewBindingsCaller(address common.Address, caller bind.ContractCaller) (*BindingsCaller, error) {
	contract, err := bindBindings(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BindingsCaller{contract: contract}, nil
}

// NewBindingsTransactor creates a new write-only instance of Bindings, bound to a specific deployed contract.
func NewBindingsTransactor(address common.Address, transactor bind.ContractTransactor) (*BindingsTransactor, error) {
	contract, err := bindBindings(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BindingsTransactor{contract: contract}, nil
}

// NewBindingsFilterer creates a new log filterer instance of Bindings, bound to a specific deployed contract.
func NewBindingsFilterer(address common.Address, filterer bind.ContractFilterer) (*BindingsFilterer, error) {
	contract, err := bindBindings(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BindingsFilterer{contract: contract}, nil
}

// bindBindings binds a generic wrapper to an already deployed contract.
func bindBindings(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BindingsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bindings *BindingsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bindings.Contract.BindingsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bindings *BindingsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.Contract.BindingsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bindings *BindingsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bindings.Contract.BindingsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bindings *BindingsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bindings.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bindings *BindingsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bindings *BindingsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bindings.Contract.contract.Transact(opts, method, params...)
}

// ACCOUNTING is a free data retrieval call binding the contract method 0x6dc3f2bd.
//
// Solidity: function ACCOUNTING() view returns(address)
func (_Bindings *BindingsCaller) ACCOUNTING(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "ACCOUNTING")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ACCOUNTING is a free data retrieval call binding the contract method 0x6dc3f2bd.
//
// Solidity: function ACCOUNTING() view returns(address)
func (_Bindings *BindingsSession) ACCOUNTING() (common.Address, error) {
	return _Bindings.Contract.ACCOUNTING(&_Bindings.CallOpts)
}

// ACCOUNTING is a free data retrieval call binding the contract method 0x6dc3f2bd.
//
// Solidity: function ACCOUNTING() view returns(address)
func (_Bindings *BindingsCallerSession) ACCOUNTING() (common.Address, error) {
	return _Bindings.Contract.ACCOUNTING(&_Bindings.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bindings *BindingsCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bindings *BindingsSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Bindings.Contract.DEFAULTADMINROLE(&_Bindings.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bindings *BindingsCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Bindings.Contract.DEFAULTADMINROLE(&_Bindings.CallOpts)
}

// ORACLE is a free data retrieval call binding the contract method 0x38013f02.
//
// Solidity: function ORACLE() view returns(address)
func (_Bindings *BindingsCaller) ORACLE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "ORACLE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ORACLE is a free data retrieval call binding the contract method 0x38013f02.
//
// Solidity: function ORACLE() view returns(address)
func (_Bindings *BindingsSession) ORACLE() (common.Address, error) {
	return _Bindings.Contract.ORACLE(&_Bindings.CallOpts)
}

// ORACLE is a free data retrieval call binding the contract method 0x38013f02.
//
// Solidity: function ORACLE() view returns(address)
func (_Bindings *BindingsCallerSession) ORACLE() (common.Address, error) {
	return _Bindings.Contract.ORACLE(&_Bindings.CallOpts)
}

// RECOVERERROLE is a free data retrieval call binding the contract method 0xacf1c948.
//
// Solidity: function RECOVERER_ROLE() view returns(bytes32)
func (_Bindings *BindingsCaller) RECOVERERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "RECOVERER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RECOVERERROLE is a free data retrieval call binding the contract method 0xacf1c948.
//
// Solidity: function RECOVERER_ROLE() view returns(bytes32)
func (_Bindings *BindingsSession) RECOVERERROLE() ([32]byte, error) {
	return _Bindings.Contract.RECOVERERROLE(&_Bindings.CallOpts)
}

// RECOVERERROLE is a free data retrieval call binding the contract method 0xacf1c948.
//
// Solidity: function RECOVERER_ROLE() view returns(bytes32)
func (_Bindings *BindingsCallerSession) RECOVERERROLE() ([32]byte, error) {
	return _Bindings.Contract.RECOVERERROLE(&_Bindings.CallOpts)
}

// STETH is a free data retrieval call binding the contract method 0xe00bfe50.
//
// Solidity: function STETH() view returns(address)
func (_Bindings *BindingsCaller) STETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "STETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STETH is a free data retrieval call binding the contract method 0xe00bfe50.
//
// Solidity: function STETH() view returns(address)
func (_Bindings *BindingsSession) STETH() (common.Address, error) {
	return _Bindings.Contract.STETH(&_Bindings.CallOpts)
}

// STETH is a free data retrieval call binding the contract method 0xe00bfe50.
//
// Solidity: function STETH() view returns(address)
func (_Bindings *BindingsCallerSession) STETH() (common.Address, error) {
	return _Bindings.Contract.STETH(&_Bindings.CallOpts)
}

// DistributedShares is a free data retrieval call binding the contract method 0xea6301ab.
//
// Solidity: function distributedShares(uint256 nodeOperatorId) view returns(uint256 distributed)
func (_Bindings *BindingsCaller) DistributedShares(opts *bind.CallOpts, nodeOperatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "distributedShares", nodeOperatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DistributedShares is a free data retrieval call binding the contract method 0xea6301ab.
//
// Solidity: function distributedShares(uint256 nodeOperatorId) view returns(uint256 distributed)
func (_Bindings *BindingsSession) DistributedShares(nodeOperatorId *big.Int) (*big.Int, error) {
	return _Bindings.Contract.DistributedShares(&_Bindings.CallOpts, nodeOperatorId)
}

// DistributedShares is a free data retrieval call binding the contract method 0xea6301ab.
//
// Solidity: function distributedShares(uint256 nodeOperatorId) view returns(uint256 distributed)
func (_Bindings *BindingsCallerSession) DistributedShares(nodeOperatorId *big.Int) (*big.Int, error) {
	return _Bindings.Contract.DistributedShares(&_Bindings.CallOpts, nodeOperatorId)
}

// DistributionDataHistoryCount is a free data retrieval call binding the contract method 0x3d18b6f3.
//
// Solidity: function distributionDataHistoryCount() view returns(uint256)
func (_Bindings *BindingsCaller) DistributionDataHistoryCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "distributionDataHistoryCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DistributionDataHistoryCount is a free data retrieval call binding the contract method 0x3d18b6f3.
//
// Solidity: function distributionDataHistoryCount() view returns(uint256)
func (_Bindings *BindingsSession) DistributionDataHistoryCount() (*big.Int, error) {
	return _Bindings.Contract.DistributionDataHistoryCount(&_Bindings.CallOpts)
}

// DistributionDataHistoryCount is a free data retrieval call binding the contract method 0x3d18b6f3.
//
// Solidity: function distributionDataHistoryCount() view returns(uint256)
func (_Bindings *BindingsCallerSession) DistributionDataHistoryCount() (*big.Int, error) {
	return _Bindings.Contract.DistributionDataHistoryCount(&_Bindings.CallOpts)
}

// GetFeesToDistribute is a free data retrieval call binding the contract method 0x5e8e8f6f.
//
// Solidity: function getFeesToDistribute(uint256 nodeOperatorId, uint256 cumulativeFeeShares, bytes32[] proof) view returns(uint256 sharesToDistribute)
func (_Bindings *BindingsCaller) GetFeesToDistribute(opts *bind.CallOpts, nodeOperatorId *big.Int, cumulativeFeeShares *big.Int, proof [][32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getFeesToDistribute", nodeOperatorId, cumulativeFeeShares, proof)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFeesToDistribute is a free data retrieval call binding the contract method 0x5e8e8f6f.
//
// Solidity: function getFeesToDistribute(uint256 nodeOperatorId, uint256 cumulativeFeeShares, bytes32[] proof) view returns(uint256 sharesToDistribute)
func (_Bindings *BindingsSession) GetFeesToDistribute(nodeOperatorId *big.Int, cumulativeFeeShares *big.Int, proof [][32]byte) (*big.Int, error) {
	return _Bindings.Contract.GetFeesToDistribute(&_Bindings.CallOpts, nodeOperatorId, cumulativeFeeShares, proof)
}

// GetFeesToDistribute is a free data retrieval call binding the contract method 0x5e8e8f6f.
//
// Solidity: function getFeesToDistribute(uint256 nodeOperatorId, uint256 cumulativeFeeShares, bytes32[] proof) view returns(uint256 sharesToDistribute)
func (_Bindings *BindingsCallerSession) GetFeesToDistribute(nodeOperatorId *big.Int, cumulativeFeeShares *big.Int, proof [][32]byte) (*big.Int, error) {
	return _Bindings.Contract.GetFeesToDistribute(&_Bindings.CallOpts, nodeOperatorId, cumulativeFeeShares, proof)
}

// GetHistoricalDistributionData is a free data retrieval call binding the contract method 0x4e5b3a62.
//
// Solidity: function getHistoricalDistributionData(uint256 index) view returns((uint256,bytes32,string,string,uint256,uint256))
func (_Bindings *BindingsCaller) GetHistoricalDistributionData(opts *bind.CallOpts, index *big.Int) (ICSFeeDistributorDistributionData, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getHistoricalDistributionData", index)

	if err != nil {
		return *new(ICSFeeDistributorDistributionData), err
	}

	out0 := *abi.ConvertType(out[0], new(ICSFeeDistributorDistributionData)).(*ICSFeeDistributorDistributionData)

	return out0, err

}

// GetHistoricalDistributionData is a free data retrieval call binding the contract method 0x4e5b3a62.
//
// Solidity: function getHistoricalDistributionData(uint256 index) view returns((uint256,bytes32,string,string,uint256,uint256))
func (_Bindings *BindingsSession) GetHistoricalDistributionData(index *big.Int) (ICSFeeDistributorDistributionData, error) {
	return _Bindings.Contract.GetHistoricalDistributionData(&_Bindings.CallOpts, index)
}

// GetHistoricalDistributionData is a free data retrieval call binding the contract method 0x4e5b3a62.
//
// Solidity: function getHistoricalDistributionData(uint256 index) view returns((uint256,bytes32,string,string,uint256,uint256))
func (_Bindings *BindingsCallerSession) GetHistoricalDistributionData(index *big.Int) (ICSFeeDistributorDistributionData, error) {
	return _Bindings.Contract.GetHistoricalDistributionData(&_Bindings.CallOpts, index)
}

// GetInitializedVersion is a free data retrieval call binding the contract method 0xb3c65015.
//
// Solidity: function getInitializedVersion() view returns(uint64)
func (_Bindings *BindingsCaller) GetInitializedVersion(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getInitializedVersion")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetInitializedVersion is a free data retrieval call binding the contract method 0xb3c65015.
//
// Solidity: function getInitializedVersion() view returns(uint64)
func (_Bindings *BindingsSession) GetInitializedVersion() (uint64, error) {
	return _Bindings.Contract.GetInitializedVersion(&_Bindings.CallOpts)
}

// GetInitializedVersion is a free data retrieval call binding the contract method 0xb3c65015.
//
// Solidity: function getInitializedVersion() view returns(uint64)
func (_Bindings *BindingsCallerSession) GetInitializedVersion() (uint64, error) {
	return _Bindings.Contract.GetInitializedVersion(&_Bindings.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bindings *BindingsCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bindings *BindingsSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Bindings.Contract.GetRoleAdmin(&_Bindings.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bindings *BindingsCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Bindings.Contract.GetRoleAdmin(&_Bindings.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Bindings *BindingsCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Bindings *BindingsSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Bindings.Contract.GetRoleMember(&_Bindings.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Bindings *BindingsCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Bindings.Contract.GetRoleMember(&_Bindings.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Bindings *BindingsCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Bindings *BindingsSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Bindings.Contract.GetRoleMemberCount(&_Bindings.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Bindings *BindingsCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Bindings.Contract.GetRoleMemberCount(&_Bindings.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bindings *BindingsCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bindings *BindingsSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Bindings.Contract.HasRole(&_Bindings.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bindings *BindingsCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Bindings.Contract.HasRole(&_Bindings.CallOpts, role, account)
}

// HashLeaf is a free data retrieval call binding the contract method 0x7e9f27ad.
//
// Solidity: function hashLeaf(uint256 nodeOperatorId, uint256 shares) pure returns(bytes32)
func (_Bindings *BindingsCaller) HashLeaf(opts *bind.CallOpts, nodeOperatorId *big.Int, shares *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "hashLeaf", nodeOperatorId, shares)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashLeaf is a free data retrieval call binding the contract method 0x7e9f27ad.
//
// Solidity: function hashLeaf(uint256 nodeOperatorId, uint256 shares) pure returns(bytes32)
func (_Bindings *BindingsSession) HashLeaf(nodeOperatorId *big.Int, shares *big.Int) ([32]byte, error) {
	return _Bindings.Contract.HashLeaf(&_Bindings.CallOpts, nodeOperatorId, shares)
}

// HashLeaf is a free data retrieval call binding the contract method 0x7e9f27ad.
//
// Solidity: function hashLeaf(uint256 nodeOperatorId, uint256 shares) pure returns(bytes32)
func (_Bindings *BindingsCallerSession) HashLeaf(nodeOperatorId *big.Int, shares *big.Int) ([32]byte, error) {
	return _Bindings.Contract.HashLeaf(&_Bindings.CallOpts, nodeOperatorId, shares)
}

// LogCid is a free data retrieval call binding the contract method 0x6f962e5c.
//
// Solidity: function logCid() view returns(string)
func (_Bindings *BindingsCaller) LogCid(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "logCid")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LogCid is a free data retrieval call binding the contract method 0x6f962e5c.
//
// Solidity: function logCid() view returns(string)
func (_Bindings *BindingsSession) LogCid() (string, error) {
	return _Bindings.Contract.LogCid(&_Bindings.CallOpts)
}

// LogCid is a free data retrieval call binding the contract method 0x6f962e5c.
//
// Solidity: function logCid() view returns(string)
func (_Bindings *BindingsCallerSession) LogCid() (string, error) {
	return _Bindings.Contract.LogCid(&_Bindings.CallOpts)
}

// PendingSharesToDistribute is a free data retrieval call binding the contract method 0xd257cf2a.
//
// Solidity: function pendingSharesToDistribute() view returns(uint256)
func (_Bindings *BindingsCaller) PendingSharesToDistribute(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "pendingSharesToDistribute")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingSharesToDistribute is a free data retrieval call binding the contract method 0xd257cf2a.
//
// Solidity: function pendingSharesToDistribute() view returns(uint256)
func (_Bindings *BindingsSession) PendingSharesToDistribute() (*big.Int, error) {
	return _Bindings.Contract.PendingSharesToDistribute(&_Bindings.CallOpts)
}

// PendingSharesToDistribute is a free data retrieval call binding the contract method 0xd257cf2a.
//
// Solidity: function pendingSharesToDistribute() view returns(uint256)
func (_Bindings *BindingsCallerSession) PendingSharesToDistribute() (*big.Int, error) {
	return _Bindings.Contract.PendingSharesToDistribute(&_Bindings.CallOpts)
}

// RebateRecipient is a free data retrieval call binding the contract method 0x3333e109.
//
// Solidity: function rebateRecipient() view returns(address)
func (_Bindings *BindingsCaller) RebateRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "rebateRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RebateRecipient is a free data retrieval call binding the contract method 0x3333e109.
//
// Solidity: function rebateRecipient() view returns(address)
func (_Bindings *BindingsSession) RebateRecipient() (common.Address, error) {
	return _Bindings.Contract.RebateRecipient(&_Bindings.CallOpts)
}

// RebateRecipient is a free data retrieval call binding the contract method 0x3333e109.
//
// Solidity: function rebateRecipient() view returns(address)
func (_Bindings *BindingsCallerSession) RebateRecipient() (common.Address, error) {
	return _Bindings.Contract.RebateRecipient(&_Bindings.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bindings *BindingsCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bindings *BindingsSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Bindings.Contract.SupportsInterface(&_Bindings.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bindings *BindingsCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Bindings.Contract.SupportsInterface(&_Bindings.CallOpts, interfaceId)
}

// TotalClaimableShares is a free data retrieval call binding the contract method 0x47d17d9d.
//
// Solidity: function totalClaimableShares() view returns(uint256)
func (_Bindings *BindingsCaller) TotalClaimableShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "totalClaimableShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalClaimableShares is a free data retrieval call binding the contract method 0x47d17d9d.
//
// Solidity: function totalClaimableShares() view returns(uint256)
func (_Bindings *BindingsSession) TotalClaimableShares() (*big.Int, error) {
	return _Bindings.Contract.TotalClaimableShares(&_Bindings.CallOpts)
}

// TotalClaimableShares is a free data retrieval call binding the contract method 0x47d17d9d.
//
// Solidity: function totalClaimableShares() view returns(uint256)
func (_Bindings *BindingsCallerSession) TotalClaimableShares() (*big.Int, error) {
	return _Bindings.Contract.TotalClaimableShares(&_Bindings.CallOpts)
}

// TreeCid is a free data retrieval call binding the contract method 0xfe3c9b9b.
//
// Solidity: function treeCid() view returns(string)
func (_Bindings *BindingsCaller) TreeCid(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "treeCid")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TreeCid is a free data retrieval call binding the contract method 0xfe3c9b9b.
//
// Solidity: function treeCid() view returns(string)
func (_Bindings *BindingsSession) TreeCid() (string, error) {
	return _Bindings.Contract.TreeCid(&_Bindings.CallOpts)
}

// TreeCid is a free data retrieval call binding the contract method 0xfe3c9b9b.
//
// Solidity: function treeCid() view returns(string)
func (_Bindings *BindingsCallerSession) TreeCid() (string, error) {
	return _Bindings.Contract.TreeCid(&_Bindings.CallOpts)
}

// TreeRoot is a free data retrieval call binding the contract method 0x14dc6c14.
//
// Solidity: function treeRoot() view returns(bytes32)
func (_Bindings *BindingsCaller) TreeRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "treeRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TreeRoot is a free data retrieval call binding the contract method 0x14dc6c14.
//
// Solidity: function treeRoot() view returns(bytes32)
func (_Bindings *BindingsSession) TreeRoot() ([32]byte, error) {
	return _Bindings.Contract.TreeRoot(&_Bindings.CallOpts)
}

// TreeRoot is a free data retrieval call binding the contract method 0x14dc6c14.
//
// Solidity: function treeRoot() view returns(bytes32)
func (_Bindings *BindingsCallerSession) TreeRoot() ([32]byte, error) {
	return _Bindings.Contract.TreeRoot(&_Bindings.CallOpts)
}

// DistributeFees is a paid mutator transaction binding the contract method 0x21893f7b.
//
// Solidity: function distributeFees(uint256 nodeOperatorId, uint256 cumulativeFeeShares, bytes32[] proof) returns(uint256 sharesToDistribute)
func (_Bindings *BindingsTransactor) DistributeFees(opts *bind.TransactOpts, nodeOperatorId *big.Int, cumulativeFeeShares *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "distributeFees", nodeOperatorId, cumulativeFeeShares, proof)
}

// DistributeFees is a paid mutator transaction binding the contract method 0x21893f7b.
//
// Solidity: function distributeFees(uint256 nodeOperatorId, uint256 cumulativeFeeShares, bytes32[] proof) returns(uint256 sharesToDistribute)
func (_Bindings *BindingsSession) DistributeFees(nodeOperatorId *big.Int, cumulativeFeeShares *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _Bindings.Contract.DistributeFees(&_Bindings.TransactOpts, nodeOperatorId, cumulativeFeeShares, proof)
}

// DistributeFees is a paid mutator transaction binding the contract method 0x21893f7b.
//
// Solidity: function distributeFees(uint256 nodeOperatorId, uint256 cumulativeFeeShares, bytes32[] proof) returns(uint256 sharesToDistribute)
func (_Bindings *BindingsTransactorSession) DistributeFees(nodeOperatorId *big.Int, cumulativeFeeShares *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _Bindings.Contract.DistributeFees(&_Bindings.TransactOpts, nodeOperatorId, cumulativeFeeShares, proof)
}

// FinalizeUpgradeV2 is a paid mutator transaction binding the contract method 0xd5ba2dcf.
//
// Solidity: function finalizeUpgradeV2(address _rebateRecipient) returns()
func (_Bindings *BindingsTransactor) FinalizeUpgradeV2(opts *bind.TransactOpts, _rebateRecipient common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "finalizeUpgradeV2", _rebateRecipient)
}

// FinalizeUpgradeV2 is a paid mutator transaction binding the contract method 0xd5ba2dcf.
//
// Solidity: function finalizeUpgradeV2(address _rebateRecipient) returns()
func (_Bindings *BindingsSession) FinalizeUpgradeV2(_rebateRecipient common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.FinalizeUpgradeV2(&_Bindings.TransactOpts, _rebateRecipient)
}

// FinalizeUpgradeV2 is a paid mutator transaction binding the contract method 0xd5ba2dcf.
//
// Solidity: function finalizeUpgradeV2(address _rebateRecipient) returns()
func (_Bindings *BindingsTransactorSession) FinalizeUpgradeV2(_rebateRecipient common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.FinalizeUpgradeV2(&_Bindings.TransactOpts, _rebateRecipient)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bindings *BindingsTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bindings *BindingsSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.GrantRole(&_Bindings.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bindings *BindingsTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.GrantRole(&_Bindings.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address admin, address _rebateRecipient) returns()
func (_Bindings *BindingsTransactor) Initialize(opts *bind.TransactOpts, admin common.Address, _rebateRecipient common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "initialize", admin, _rebateRecipient)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address admin, address _rebateRecipient) returns()
func (_Bindings *BindingsSession) Initialize(admin common.Address, _rebateRecipient common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.Initialize(&_Bindings.TransactOpts, admin, _rebateRecipient)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address admin, address _rebateRecipient) returns()
func (_Bindings *BindingsTransactorSession) Initialize(admin common.Address, _rebateRecipient common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.Initialize(&_Bindings.TransactOpts, admin, _rebateRecipient)
}

// ProcessOracleReport is a paid mutator transaction binding the contract method 0x2ffa14e1.
//
// Solidity: function processOracleReport(bytes32 _treeRoot, string _treeCid, string _logCid, uint256 distributed, uint256 rebate, uint256 refSlot) returns()
func (_Bindings *BindingsTransactor) ProcessOracleReport(opts *bind.TransactOpts, _treeRoot [32]byte, _treeCid string, _logCid string, distributed *big.Int, rebate *big.Int, refSlot *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "processOracleReport", _treeRoot, _treeCid, _logCid, distributed, rebate, refSlot)
}

// ProcessOracleReport is a paid mutator transaction binding the contract method 0x2ffa14e1.
//
// Solidity: function processOracleReport(bytes32 _treeRoot, string _treeCid, string _logCid, uint256 distributed, uint256 rebate, uint256 refSlot) returns()
func (_Bindings *BindingsSession) ProcessOracleReport(_treeRoot [32]byte, _treeCid string, _logCid string, distributed *big.Int, rebate *big.Int, refSlot *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.ProcessOracleReport(&_Bindings.TransactOpts, _treeRoot, _treeCid, _logCid, distributed, rebate, refSlot)
}

// ProcessOracleReport is a paid mutator transaction binding the contract method 0x2ffa14e1.
//
// Solidity: function processOracleReport(bytes32 _treeRoot, string _treeCid, string _logCid, uint256 distributed, uint256 rebate, uint256 refSlot) returns()
func (_Bindings *BindingsTransactorSession) ProcessOracleReport(_treeRoot [32]byte, _treeCid string, _logCid string, distributed *big.Int, rebate *big.Int, refSlot *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.ProcessOracleReport(&_Bindings.TransactOpts, _treeRoot, _treeCid, _logCid, distributed, rebate, refSlot)
}

// RecoverERC1155 is a paid mutator transaction binding the contract method 0x5c654ad9.
//
// Solidity: function recoverERC1155(address token, uint256 tokenId) returns()
func (_Bindings *BindingsTransactor) RecoverERC1155(opts *bind.TransactOpts, token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "recoverERC1155", token, tokenId)
}

// RecoverERC1155 is a paid mutator transaction binding the contract method 0x5c654ad9.
//
// Solidity: function recoverERC1155(address token, uint256 tokenId) returns()
func (_Bindings *BindingsSession) RecoverERC1155(token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.RecoverERC1155(&_Bindings.TransactOpts, token, tokenId)
}

// RecoverERC1155 is a paid mutator transaction binding the contract method 0x5c654ad9.
//
// Solidity: function recoverERC1155(address token, uint256 tokenId) returns()
func (_Bindings *BindingsTransactorSession) RecoverERC1155(token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.RecoverERC1155(&_Bindings.TransactOpts, token, tokenId)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address token, uint256 amount) returns()
func (_Bindings *BindingsTransactor) RecoverERC20(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "recoverERC20", token, amount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address token, uint256 amount) returns()
func (_Bindings *BindingsSession) RecoverERC20(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.RecoverERC20(&_Bindings.TransactOpts, token, amount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address token, uint256 amount) returns()
func (_Bindings *BindingsTransactorSession) RecoverERC20(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.RecoverERC20(&_Bindings.TransactOpts, token, amount)
}

// RecoverERC721 is a paid mutator transaction binding the contract method 0x819d4cc6.
//
// Solidity: function recoverERC721(address token, uint256 tokenId) returns()
func (_Bindings *BindingsTransactor) RecoverERC721(opts *bind.TransactOpts, token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "recoverERC721", token, tokenId)
}

// RecoverERC721 is a paid mutator transaction binding the contract method 0x819d4cc6.
//
// Solidity: function recoverERC721(address token, uint256 tokenId) returns()
func (_Bindings *BindingsSession) RecoverERC721(token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.RecoverERC721(&_Bindings.TransactOpts, token, tokenId)
}

// RecoverERC721 is a paid mutator transaction binding the contract method 0x819d4cc6.
//
// Solidity: function recoverERC721(address token, uint256 tokenId) returns()
func (_Bindings *BindingsTransactorSession) RecoverERC721(token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.RecoverERC721(&_Bindings.TransactOpts, token, tokenId)
}

// RecoverEther is a paid mutator transaction binding the contract method 0x52d8bfc2.
//
// Solidity: function recoverEther() returns()
func (_Bindings *BindingsTransactor) RecoverEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "recoverEther")
}

// RecoverEther is a paid mutator transaction binding the contract method 0x52d8bfc2.
//
// Solidity: function recoverEther() returns()
func (_Bindings *BindingsSession) RecoverEther() (*types.Transaction, error) {
	return _Bindings.Contract.RecoverEther(&_Bindings.TransactOpts)
}

// RecoverEther is a paid mutator transaction binding the contract method 0x52d8bfc2.
//
// Solidity: function recoverEther() returns()
func (_Bindings *BindingsTransactorSession) RecoverEther() (*types.Transaction, error) {
	return _Bindings.Contract.RecoverEther(&_Bindings.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Bindings *BindingsTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Bindings *BindingsSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.RenounceRole(&_Bindings.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Bindings *BindingsTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.RenounceRole(&_Bindings.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bindings *BindingsTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bindings *BindingsSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.RevokeRole(&_Bindings.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bindings *BindingsTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.RevokeRole(&_Bindings.TransactOpts, role, account)
}

// SetRebateRecipient is a paid mutator transaction binding the contract method 0xe877f068.
//
// Solidity: function setRebateRecipient(address _rebateRecipient) returns()
func (_Bindings *BindingsTransactor) SetRebateRecipient(opts *bind.TransactOpts, _rebateRecipient common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setRebateRecipient", _rebateRecipient)
}

// SetRebateRecipient is a paid mutator transaction binding the contract method 0xe877f068.
//
// Solidity: function setRebateRecipient(address _rebateRecipient) returns()
func (_Bindings *BindingsSession) SetRebateRecipient(_rebateRecipient common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.SetRebateRecipient(&_Bindings.TransactOpts, _rebateRecipient)
}

// SetRebateRecipient is a paid mutator transaction binding the contract method 0xe877f068.
//
// Solidity: function setRebateRecipient(address _rebateRecipient) returns()
func (_Bindings *BindingsTransactorSession) SetRebateRecipient(_rebateRecipient common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.SetRebateRecipient(&_Bindings.TransactOpts, _rebateRecipient)
}

// BindingsDistributionDataUpdatedIterator is returned from FilterDistributionDataUpdated and is used to iterate over the raw logs and unpacked data for DistributionDataUpdated events raised by the Bindings contract.
type BindingsDistributionDataUpdatedIterator struct {
	Event *BindingsDistributionDataUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsDistributionDataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsDistributionDataUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsDistributionDataUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsDistributionDataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsDistributionDataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsDistributionDataUpdated represents a DistributionDataUpdated event raised by the Bindings contract.
type BindingsDistributionDataUpdated struct {
	TotalClaimableShares *big.Int
	TreeRoot             [32]byte
	TreeCid              string
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterDistributionDataUpdated is a free log retrieval operation binding the contract event 0x26dec7cc117e9b3907dc1f90d2dc5f6e04dbb9f285f5898be2c82ec524dcd424.
//
// Solidity: event DistributionDataUpdated(uint256 totalClaimableShares, bytes32 treeRoot, string treeCid)
func (_Bindings *BindingsFilterer) FilterDistributionDataUpdated(opts *bind.FilterOpts) (*BindingsDistributionDataUpdatedIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "DistributionDataUpdated")
	if err != nil {
		return nil, err
	}
	return &BindingsDistributionDataUpdatedIterator{contract: _Bindings.contract, event: "DistributionDataUpdated", logs: logs, sub: sub}, nil
}

// WatchDistributionDataUpdated is a free log subscription operation binding the contract event 0x26dec7cc117e9b3907dc1f90d2dc5f6e04dbb9f285f5898be2c82ec524dcd424.
//
// Solidity: event DistributionDataUpdated(uint256 totalClaimableShares, bytes32 treeRoot, string treeCid)
func (_Bindings *BindingsFilterer) WatchDistributionDataUpdated(opts *bind.WatchOpts, sink chan<- *BindingsDistributionDataUpdated) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "DistributionDataUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsDistributionDataUpdated)
				if err := _Bindings.contract.UnpackLog(event, "DistributionDataUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDistributionDataUpdated is a log parse operation binding the contract event 0x26dec7cc117e9b3907dc1f90d2dc5f6e04dbb9f285f5898be2c82ec524dcd424.
//
// Solidity: event DistributionDataUpdated(uint256 totalClaimableShares, bytes32 treeRoot, string treeCid)
func (_Bindings *BindingsFilterer) ParseDistributionDataUpdated(log types.Log) (*BindingsDistributionDataUpdated, error) {
	event := new(BindingsDistributionDataUpdated)
	if err := _Bindings.contract.UnpackLog(event, "DistributionDataUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsDistributionLogUpdatedIterator is returned from FilterDistributionLogUpdated and is used to iterate over the raw logs and unpacked data for DistributionLogUpdated events raised by the Bindings contract.
type BindingsDistributionLogUpdatedIterator struct {
	Event *BindingsDistributionLogUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsDistributionLogUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsDistributionLogUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsDistributionLogUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsDistributionLogUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsDistributionLogUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsDistributionLogUpdated represents a DistributionLogUpdated event raised by the Bindings contract.
type BindingsDistributionLogUpdated struct {
	LogCid string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDistributionLogUpdated is a free log retrieval operation binding the contract event 0x1f1a488b71a099a0d9cb71f60e14cf90bd1b5b188ca593111a40f533a3130b3b.
//
// Solidity: event DistributionLogUpdated(string logCid)
func (_Bindings *BindingsFilterer) FilterDistributionLogUpdated(opts *bind.FilterOpts) (*BindingsDistributionLogUpdatedIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "DistributionLogUpdated")
	if err != nil {
		return nil, err
	}
	return &BindingsDistributionLogUpdatedIterator{contract: _Bindings.contract, event: "DistributionLogUpdated", logs: logs, sub: sub}, nil
}

// WatchDistributionLogUpdated is a free log subscription operation binding the contract event 0x1f1a488b71a099a0d9cb71f60e14cf90bd1b5b188ca593111a40f533a3130b3b.
//
// Solidity: event DistributionLogUpdated(string logCid)
func (_Bindings *BindingsFilterer) WatchDistributionLogUpdated(opts *bind.WatchOpts, sink chan<- *BindingsDistributionLogUpdated) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "DistributionLogUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsDistributionLogUpdated)
				if err := _Bindings.contract.UnpackLog(event, "DistributionLogUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDistributionLogUpdated is a log parse operation binding the contract event 0x1f1a488b71a099a0d9cb71f60e14cf90bd1b5b188ca593111a40f533a3130b3b.
//
// Solidity: event DistributionLogUpdated(string logCid)
func (_Bindings *BindingsFilterer) ParseDistributionLogUpdated(log types.Log) (*BindingsDistributionLogUpdated, error) {
	event := new(BindingsDistributionLogUpdated)
	if err := _Bindings.contract.UnpackLog(event, "DistributionLogUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsERC1155RecoveredIterator is returned from FilterERC1155Recovered and is used to iterate over the raw logs and unpacked data for ERC1155Recovered events raised by the Bindings contract.
type BindingsERC1155RecoveredIterator struct {
	Event *BindingsERC1155Recovered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsERC1155RecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsERC1155Recovered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsERC1155Recovered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsERC1155RecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsERC1155RecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsERC1155Recovered represents a ERC1155Recovered event raised by the Bindings contract.
type BindingsERC1155Recovered struct {
	Token     common.Address
	TokenId   *big.Int
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERC1155Recovered is a free log retrieval operation binding the contract event 0x5cf02e753b3eb0f4bee4460a72817d8e5e3c75cd4d65c1d0b06dca88b8032936.
//
// Solidity: event ERC1155Recovered(address indexed token, uint256 tokenId, address indexed recipient, uint256 amount)
func (_Bindings *BindingsFilterer) FilterERC1155Recovered(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*BindingsERC1155RecoveredIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "ERC1155Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &BindingsERC1155RecoveredIterator{contract: _Bindings.contract, event: "ERC1155Recovered", logs: logs, sub: sub}, nil
}

// WatchERC1155Recovered is a free log subscription operation binding the contract event 0x5cf02e753b3eb0f4bee4460a72817d8e5e3c75cd4d65c1d0b06dca88b8032936.
//
// Solidity: event ERC1155Recovered(address indexed token, uint256 tokenId, address indexed recipient, uint256 amount)
func (_Bindings *BindingsFilterer) WatchERC1155Recovered(opts *bind.WatchOpts, sink chan<- *BindingsERC1155Recovered, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "ERC1155Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsERC1155Recovered)
				if err := _Bindings.contract.UnpackLog(event, "ERC1155Recovered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseERC1155Recovered is a log parse operation binding the contract event 0x5cf02e753b3eb0f4bee4460a72817d8e5e3c75cd4d65c1d0b06dca88b8032936.
//
// Solidity: event ERC1155Recovered(address indexed token, uint256 tokenId, address indexed recipient, uint256 amount)
func (_Bindings *BindingsFilterer) ParseERC1155Recovered(log types.Log) (*BindingsERC1155Recovered, error) {
	event := new(BindingsERC1155Recovered)
	if err := _Bindings.contract.UnpackLog(event, "ERC1155Recovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsERC20RecoveredIterator is returned from FilterERC20Recovered and is used to iterate over the raw logs and unpacked data for ERC20Recovered events raised by the Bindings contract.
type BindingsERC20RecoveredIterator struct {
	Event *BindingsERC20Recovered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsERC20RecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsERC20Recovered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsERC20Recovered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsERC20RecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsERC20RecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsERC20Recovered represents a ERC20Recovered event raised by the Bindings contract.
type BindingsERC20Recovered struct {
	Token     common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERC20Recovered is a free log retrieval operation binding the contract event 0xaca8fb252cde442184e5f10e0f2e6e4029e8cd7717cae63559079610702436aa.
//
// Solidity: event ERC20Recovered(address indexed token, address indexed recipient, uint256 amount)
func (_Bindings *BindingsFilterer) FilterERC20Recovered(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*BindingsERC20RecoveredIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "ERC20Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &BindingsERC20RecoveredIterator{contract: _Bindings.contract, event: "ERC20Recovered", logs: logs, sub: sub}, nil
}

// WatchERC20Recovered is a free log subscription operation binding the contract event 0xaca8fb252cde442184e5f10e0f2e6e4029e8cd7717cae63559079610702436aa.
//
// Solidity: event ERC20Recovered(address indexed token, address indexed recipient, uint256 amount)
func (_Bindings *BindingsFilterer) WatchERC20Recovered(opts *bind.WatchOpts, sink chan<- *BindingsERC20Recovered, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "ERC20Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsERC20Recovered)
				if err := _Bindings.contract.UnpackLog(event, "ERC20Recovered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseERC20Recovered is a log parse operation binding the contract event 0xaca8fb252cde442184e5f10e0f2e6e4029e8cd7717cae63559079610702436aa.
//
// Solidity: event ERC20Recovered(address indexed token, address indexed recipient, uint256 amount)
func (_Bindings *BindingsFilterer) ParseERC20Recovered(log types.Log) (*BindingsERC20Recovered, error) {
	event := new(BindingsERC20Recovered)
	if err := _Bindings.contract.UnpackLog(event, "ERC20Recovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsERC721RecoveredIterator is returned from FilterERC721Recovered and is used to iterate over the raw logs and unpacked data for ERC721Recovered events raised by the Bindings contract.
type BindingsERC721RecoveredIterator struct {
	Event *BindingsERC721Recovered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsERC721RecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsERC721Recovered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsERC721Recovered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsERC721RecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsERC721RecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsERC721Recovered represents a ERC721Recovered event raised by the Bindings contract.
type BindingsERC721Recovered struct {
	Token     common.Address
	TokenId   *big.Int
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERC721Recovered is a free log retrieval operation binding the contract event 0x8166bf75d2ff2fa3c8f3c44410540bf42e9a5359b48409e8d660291dc9f788c8.
//
// Solidity: event ERC721Recovered(address indexed token, uint256 tokenId, address indexed recipient)
func (_Bindings *BindingsFilterer) FilterERC721Recovered(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*BindingsERC721RecoveredIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "ERC721Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &BindingsERC721RecoveredIterator{contract: _Bindings.contract, event: "ERC721Recovered", logs: logs, sub: sub}, nil
}

// WatchERC721Recovered is a free log subscription operation binding the contract event 0x8166bf75d2ff2fa3c8f3c44410540bf42e9a5359b48409e8d660291dc9f788c8.
//
// Solidity: event ERC721Recovered(address indexed token, uint256 tokenId, address indexed recipient)
func (_Bindings *BindingsFilterer) WatchERC721Recovered(opts *bind.WatchOpts, sink chan<- *BindingsERC721Recovered, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "ERC721Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsERC721Recovered)
				if err := _Bindings.contract.UnpackLog(event, "ERC721Recovered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseERC721Recovered is a log parse operation binding the contract event 0x8166bf75d2ff2fa3c8f3c44410540bf42e9a5359b48409e8d660291dc9f788c8.
//
// Solidity: event ERC721Recovered(address indexed token, uint256 tokenId, address indexed recipient)
func (_Bindings *BindingsFilterer) ParseERC721Recovered(log types.Log) (*BindingsERC721Recovered, error) {
	event := new(BindingsERC721Recovered)
	if err := _Bindings.contract.UnpackLog(event, "ERC721Recovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsEtherRecoveredIterator is returned from FilterEtherRecovered and is used to iterate over the raw logs and unpacked data for EtherRecovered events raised by the Bindings contract.
type BindingsEtherRecoveredIterator struct {
	Event *BindingsEtherRecovered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsEtherRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsEtherRecovered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsEtherRecovered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsEtherRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsEtherRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsEtherRecovered represents a EtherRecovered event raised by the Bindings contract.
type BindingsEtherRecovered struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEtherRecovered is a free log retrieval operation binding the contract event 0x8e274e42262a7f013b700b35c2b4629ccce1702f8fe83f8dfb7eacbb26a4382c.
//
// Solidity: event EtherRecovered(address indexed recipient, uint256 amount)
func (_Bindings *BindingsFilterer) FilterEtherRecovered(opts *bind.FilterOpts, recipient []common.Address) (*BindingsEtherRecoveredIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "EtherRecovered", recipientRule)
	if err != nil {
		return nil, err
	}
	return &BindingsEtherRecoveredIterator{contract: _Bindings.contract, event: "EtherRecovered", logs: logs, sub: sub}, nil
}

// WatchEtherRecovered is a free log subscription operation binding the contract event 0x8e274e42262a7f013b700b35c2b4629ccce1702f8fe83f8dfb7eacbb26a4382c.
//
// Solidity: event EtherRecovered(address indexed recipient, uint256 amount)
func (_Bindings *BindingsFilterer) WatchEtherRecovered(opts *bind.WatchOpts, sink chan<- *BindingsEtherRecovered, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "EtherRecovered", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsEtherRecovered)
				if err := _Bindings.contract.UnpackLog(event, "EtherRecovered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEtherRecovered is a log parse operation binding the contract event 0x8e274e42262a7f013b700b35c2b4629ccce1702f8fe83f8dfb7eacbb26a4382c.
//
// Solidity: event EtherRecovered(address indexed recipient, uint256 amount)
func (_Bindings *BindingsFilterer) ParseEtherRecovered(log types.Log) (*BindingsEtherRecovered, error) {
	event := new(BindingsEtherRecovered)
	if err := _Bindings.contract.UnpackLog(event, "EtherRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Bindings contract.
type BindingsInitializedIterator struct {
	Event *BindingsInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsInitialized represents a Initialized event raised by the Bindings contract.
type BindingsInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Bindings *BindingsFilterer) FilterInitialized(opts *bind.FilterOpts) (*BindingsInitializedIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BindingsInitializedIterator{contract: _Bindings.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Bindings *BindingsFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BindingsInitialized) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsInitialized)
				if err := _Bindings.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Bindings *BindingsFilterer) ParseInitialized(log types.Log) (*BindingsInitialized, error) {
	event := new(BindingsInitialized)
	if err := _Bindings.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsModuleFeeDistributedIterator is returned from FilterModuleFeeDistributed and is used to iterate over the raw logs and unpacked data for ModuleFeeDistributed events raised by the Bindings contract.
type BindingsModuleFeeDistributedIterator struct {
	Event *BindingsModuleFeeDistributed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsModuleFeeDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsModuleFeeDistributed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsModuleFeeDistributed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsModuleFeeDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsModuleFeeDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsModuleFeeDistributed represents a ModuleFeeDistributed event raised by the Bindings contract.
type BindingsModuleFeeDistributed struct {
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterModuleFeeDistributed is a free log retrieval operation binding the contract event 0x010f65f5f56ba52d759f7b1dc49a3d277570cc2aa631e9c865b073a0ffc2af41.
//
// Solidity: event ModuleFeeDistributed(uint256 shares)
func (_Bindings *BindingsFilterer) FilterModuleFeeDistributed(opts *bind.FilterOpts) (*BindingsModuleFeeDistributedIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "ModuleFeeDistributed")
	if err != nil {
		return nil, err
	}
	return &BindingsModuleFeeDistributedIterator{contract: _Bindings.contract, event: "ModuleFeeDistributed", logs: logs, sub: sub}, nil
}

// WatchModuleFeeDistributed is a free log subscription operation binding the contract event 0x010f65f5f56ba52d759f7b1dc49a3d277570cc2aa631e9c865b073a0ffc2af41.
//
// Solidity: event ModuleFeeDistributed(uint256 shares)
func (_Bindings *BindingsFilterer) WatchModuleFeeDistributed(opts *bind.WatchOpts, sink chan<- *BindingsModuleFeeDistributed) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "ModuleFeeDistributed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsModuleFeeDistributed)
				if err := _Bindings.contract.UnpackLog(event, "ModuleFeeDistributed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseModuleFeeDistributed is a log parse operation binding the contract event 0x010f65f5f56ba52d759f7b1dc49a3d277570cc2aa631e9c865b073a0ffc2af41.
//
// Solidity: event ModuleFeeDistributed(uint256 shares)
func (_Bindings *BindingsFilterer) ParseModuleFeeDistributed(log types.Log) (*BindingsModuleFeeDistributed, error) {
	event := new(BindingsModuleFeeDistributed)
	if err := _Bindings.contract.UnpackLog(event, "ModuleFeeDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsOperatorFeeDistributedIterator is returned from FilterOperatorFeeDistributed and is used to iterate over the raw logs and unpacked data for OperatorFeeDistributed events raised by the Bindings contract.
type BindingsOperatorFeeDistributedIterator struct {
	Event *BindingsOperatorFeeDistributed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsOperatorFeeDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsOperatorFeeDistributed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsOperatorFeeDistributed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsOperatorFeeDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsOperatorFeeDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsOperatorFeeDistributed represents a OperatorFeeDistributed event raised by the Bindings contract.
type BindingsOperatorFeeDistributed struct {
	NodeOperatorId *big.Int
	Shares         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOperatorFeeDistributed is a free log retrieval operation binding the contract event 0x4b7ab1c192267e83350d06490a852b8dbbb25bfa00fd065b1862cf7accd2ab90.
//
// Solidity: event OperatorFeeDistributed(uint256 indexed nodeOperatorId, uint256 shares)
func (_Bindings *BindingsFilterer) FilterOperatorFeeDistributed(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*BindingsOperatorFeeDistributedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "OperatorFeeDistributed", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &BindingsOperatorFeeDistributedIterator{contract: _Bindings.contract, event: "OperatorFeeDistributed", logs: logs, sub: sub}, nil
}

// WatchOperatorFeeDistributed is a free log subscription operation binding the contract event 0x4b7ab1c192267e83350d06490a852b8dbbb25bfa00fd065b1862cf7accd2ab90.
//
// Solidity: event OperatorFeeDistributed(uint256 indexed nodeOperatorId, uint256 shares)
func (_Bindings *BindingsFilterer) WatchOperatorFeeDistributed(opts *bind.WatchOpts, sink chan<- *BindingsOperatorFeeDistributed, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "OperatorFeeDistributed", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsOperatorFeeDistributed)
				if err := _Bindings.contract.UnpackLog(event, "OperatorFeeDistributed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorFeeDistributed is a log parse operation binding the contract event 0x4b7ab1c192267e83350d06490a852b8dbbb25bfa00fd065b1862cf7accd2ab90.
//
// Solidity: event OperatorFeeDistributed(uint256 indexed nodeOperatorId, uint256 shares)
func (_Bindings *BindingsFilterer) ParseOperatorFeeDistributed(log types.Log) (*BindingsOperatorFeeDistributed, error) {
	event := new(BindingsOperatorFeeDistributed)
	if err := _Bindings.contract.UnpackLog(event, "OperatorFeeDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsRebateRecipientSetIterator is returned from FilterRebateRecipientSet and is used to iterate over the raw logs and unpacked data for RebateRecipientSet events raised by the Bindings contract.
type BindingsRebateRecipientSetIterator struct {
	Event *BindingsRebateRecipientSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsRebateRecipientSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsRebateRecipientSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsRebateRecipientSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsRebateRecipientSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsRebateRecipientSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsRebateRecipientSet represents a RebateRecipientSet event raised by the Bindings contract.
type BindingsRebateRecipientSet struct {
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRebateRecipientSet is a free log retrieval operation binding the contract event 0x9f8636d85f90aba9c7d2c9e076c6102a5459d2e063afb71d81328bbb3608a234.
//
// Solidity: event RebateRecipientSet(address recipient)
func (_Bindings *BindingsFilterer) FilterRebateRecipientSet(opts *bind.FilterOpts) (*BindingsRebateRecipientSetIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "RebateRecipientSet")
	if err != nil {
		return nil, err
	}
	return &BindingsRebateRecipientSetIterator{contract: _Bindings.contract, event: "RebateRecipientSet", logs: logs, sub: sub}, nil
}

// WatchRebateRecipientSet is a free log subscription operation binding the contract event 0x9f8636d85f90aba9c7d2c9e076c6102a5459d2e063afb71d81328bbb3608a234.
//
// Solidity: event RebateRecipientSet(address recipient)
func (_Bindings *BindingsFilterer) WatchRebateRecipientSet(opts *bind.WatchOpts, sink chan<- *BindingsRebateRecipientSet) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "RebateRecipientSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsRebateRecipientSet)
				if err := _Bindings.contract.UnpackLog(event, "RebateRecipientSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRebateRecipientSet is a log parse operation binding the contract event 0x9f8636d85f90aba9c7d2c9e076c6102a5459d2e063afb71d81328bbb3608a234.
//
// Solidity: event RebateRecipientSet(address recipient)
func (_Bindings *BindingsFilterer) ParseRebateRecipientSet(log types.Log) (*BindingsRebateRecipientSet, error) {
	event := new(BindingsRebateRecipientSet)
	if err := _Bindings.contract.UnpackLog(event, "RebateRecipientSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsRebateTransferredIterator is returned from FilterRebateTransferred and is used to iterate over the raw logs and unpacked data for RebateTransferred events raised by the Bindings contract.
type BindingsRebateTransferredIterator struct {
	Event *BindingsRebateTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsRebateTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsRebateTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsRebateTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsRebateTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsRebateTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsRebateTransferred represents a RebateTransferred event raised by the Bindings contract.
type BindingsRebateTransferred struct {
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRebateTransferred is a free log retrieval operation binding the contract event 0x7462935fb42d34d84233f737293310ca24e851021f9cb7f2549470cdf6de56bf.
//
// Solidity: event RebateTransferred(uint256 shares)
func (_Bindings *BindingsFilterer) FilterRebateTransferred(opts *bind.FilterOpts) (*BindingsRebateTransferredIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "RebateTransferred")
	if err != nil {
		return nil, err
	}
	return &BindingsRebateTransferredIterator{contract: _Bindings.contract, event: "RebateTransferred", logs: logs, sub: sub}, nil
}

// WatchRebateTransferred is a free log subscription operation binding the contract event 0x7462935fb42d34d84233f737293310ca24e851021f9cb7f2549470cdf6de56bf.
//
// Solidity: event RebateTransferred(uint256 shares)
func (_Bindings *BindingsFilterer) WatchRebateTransferred(opts *bind.WatchOpts, sink chan<- *BindingsRebateTransferred) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "RebateTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsRebateTransferred)
				if err := _Bindings.contract.UnpackLog(event, "RebateTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRebateTransferred is a log parse operation binding the contract event 0x7462935fb42d34d84233f737293310ca24e851021f9cb7f2549470cdf6de56bf.
//
// Solidity: event RebateTransferred(uint256 shares)
func (_Bindings *BindingsFilterer) ParseRebateTransferred(log types.Log) (*BindingsRebateTransferred, error) {
	event := new(BindingsRebateTransferred)
	if err := _Bindings.contract.UnpackLog(event, "RebateTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Bindings contract.
type BindingsRoleAdminChangedIterator struct {
	Event *BindingsRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsRoleAdminChanged represents a RoleAdminChanged event raised by the Bindings contract.
type BindingsRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Bindings *BindingsFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BindingsRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BindingsRoleAdminChangedIterator{contract: _Bindings.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Bindings *BindingsFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BindingsRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsRoleAdminChanged)
				if err := _Bindings.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Bindings *BindingsFilterer) ParseRoleAdminChanged(log types.Log) (*BindingsRoleAdminChanged, error) {
	event := new(BindingsRoleAdminChanged)
	if err := _Bindings.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Bindings contract.
type BindingsRoleGrantedIterator struct {
	Event *BindingsRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsRoleGranted represents a RoleGranted event raised by the Bindings contract.
type BindingsRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bindings *BindingsFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BindingsRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BindingsRoleGrantedIterator{contract: _Bindings.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bindings *BindingsFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BindingsRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsRoleGranted)
				if err := _Bindings.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bindings *BindingsFilterer) ParseRoleGranted(log types.Log) (*BindingsRoleGranted, error) {
	event := new(BindingsRoleGranted)
	if err := _Bindings.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Bindings contract.
type BindingsRoleRevokedIterator struct {
	Event *BindingsRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsRoleRevoked represents a RoleRevoked event raised by the Bindings contract.
type BindingsRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bindings *BindingsFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BindingsRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BindingsRoleRevokedIterator{contract: _Bindings.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bindings *BindingsFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BindingsRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsRoleRevoked)
				if err := _Bindings.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bindings *BindingsFilterer) ParseRoleRevoked(log types.Log) (*BindingsRoleRevoked, error) {
	event := new(BindingsRoleRevoked)
	if err := _Bindings.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsStETHSharesRecoveredIterator is returned from FilterStETHSharesRecovered and is used to iterate over the raw logs and unpacked data for StETHSharesRecovered events raised by the Bindings contract.
type BindingsStETHSharesRecoveredIterator struct {
	Event *BindingsStETHSharesRecovered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsStETHSharesRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsStETHSharesRecovered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsStETHSharesRecovered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsStETHSharesRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsStETHSharesRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsStETHSharesRecovered represents a StETHSharesRecovered event raised by the Bindings contract.
type BindingsStETHSharesRecovered struct {
	Recipient common.Address
	Shares    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStETHSharesRecovered is a free log retrieval operation binding the contract event 0x426e7e0100db57255d4af4a46cd49552ef74f5f002bbdc8d4ebb6371c0070a02.
//
// Solidity: event StETHSharesRecovered(address indexed recipient, uint256 shares)
func (_Bindings *BindingsFilterer) FilterStETHSharesRecovered(opts *bind.FilterOpts, recipient []common.Address) (*BindingsStETHSharesRecoveredIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "StETHSharesRecovered", recipientRule)
	if err != nil {
		return nil, err
	}
	return &BindingsStETHSharesRecoveredIterator{contract: _Bindings.contract, event: "StETHSharesRecovered", logs: logs, sub: sub}, nil
}

// WatchStETHSharesRecovered is a free log subscription operation binding the contract event 0x426e7e0100db57255d4af4a46cd49552ef74f5f002bbdc8d4ebb6371c0070a02.
//
// Solidity: event StETHSharesRecovered(address indexed recipient, uint256 shares)
func (_Bindings *BindingsFilterer) WatchStETHSharesRecovered(opts *bind.WatchOpts, sink chan<- *BindingsStETHSharesRecovered, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "StETHSharesRecovered", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsStETHSharesRecovered)
				if err := _Bindings.contract.UnpackLog(event, "StETHSharesRecovered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStETHSharesRecovered is a log parse operation binding the contract event 0x426e7e0100db57255d4af4a46cd49552ef74f5f002bbdc8d4ebb6371c0070a02.
//
// Solidity: event StETHSharesRecovered(address indexed recipient, uint256 shares)
func (_Bindings *BindingsFilterer) ParseStETHSharesRecovered(log types.Log) (*BindingsStETHSharesRecovered, error) {
	event := new(BindingsStETHSharesRecovered)
	if err := _Bindings.contract.UnpackLog(event, "StETHSharesRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
