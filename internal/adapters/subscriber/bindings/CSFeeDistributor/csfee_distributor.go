// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package csfeedistributor

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

// CsfeedistributorMetaData contains all meta data concerning the Csfeedistributor contract.
var CsfeedistributorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stETH\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"accounting\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedToSendEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeSharesDecrease\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShares\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTreeCID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTreeRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAccounting\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllowedToRecover\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughShares\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOracle\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAccountingAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAdminAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroOracleAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroStEthAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalClaimableShares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"treeRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"treeCid\",\"type\":\"string\"}],\"name\":\"DistributionDataUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC1155Recovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC20Recovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"ERC721Recovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EtherRecovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"FeeDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"StETHSharesRecovered\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ACCOUNTING\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ORACLE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RECOVERER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STETH\",\"outputs\":[{\"internalType\":\"contractIStETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"distributeFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"sharesToDistribute\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"distributedShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"getFeesToDistribute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"sharesToDistribute\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"hashLeaf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingSharesToDistribute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_treeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_treeCid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"distributed\",\"type\":\"uint256\"}],\"name\":\"processOracleReport\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"recoverERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"recoverERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"recoverERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recoverEther\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalClaimableShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treeCid\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treeRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CsfeedistributorABI is the input ABI used to generate the binding from.
// Deprecated: Use CsfeedistributorMetaData.ABI instead.
var CsfeedistributorABI = CsfeedistributorMetaData.ABI

// Csfeedistributor is an auto generated Go binding around an Ethereum contract.
type Csfeedistributor struct {
	CsfeedistributorCaller     // Read-only binding to the contract
	CsfeedistributorTransactor // Write-only binding to the contract
	CsfeedistributorFilterer   // Log filterer for contract events
}

// CsfeedistributorCaller is an auto generated read-only Go binding around an Ethereum contract.
type CsfeedistributorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CsfeedistributorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CsfeedistributorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CsfeedistributorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CsfeedistributorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CsfeedistributorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CsfeedistributorSession struct {
	Contract     *Csfeedistributor // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CsfeedistributorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CsfeedistributorCallerSession struct {
	Contract *CsfeedistributorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// CsfeedistributorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CsfeedistributorTransactorSession struct {
	Contract     *CsfeedistributorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// CsfeedistributorRaw is an auto generated low-level Go binding around an Ethereum contract.
type CsfeedistributorRaw struct {
	Contract *Csfeedistributor // Generic contract binding to access the raw methods on
}

// CsfeedistributorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CsfeedistributorCallerRaw struct {
	Contract *CsfeedistributorCaller // Generic read-only contract binding to access the raw methods on
}

// CsfeedistributorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CsfeedistributorTransactorRaw struct {
	Contract *CsfeedistributorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCsfeedistributor creates a new instance of Csfeedistributor, bound to a specific deployed contract.
func NewCsfeedistributor(address common.Address, backend bind.ContractBackend) (*Csfeedistributor, error) {
	contract, err := bindCsfeedistributor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Csfeedistributor{CsfeedistributorCaller: CsfeedistributorCaller{contract: contract}, CsfeedistributorTransactor: CsfeedistributorTransactor{contract: contract}, CsfeedistributorFilterer: CsfeedistributorFilterer{contract: contract}}, nil
}

// NewCsfeedistributorCaller creates a new read-only instance of Csfeedistributor, bound to a specific deployed contract.
func NewCsfeedistributorCaller(address common.Address, caller bind.ContractCaller) (*CsfeedistributorCaller, error) {
	contract, err := bindCsfeedistributor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CsfeedistributorCaller{contract: contract}, nil
}

// NewCsfeedistributorTransactor creates a new write-only instance of Csfeedistributor, bound to a specific deployed contract.
func NewCsfeedistributorTransactor(address common.Address, transactor bind.ContractTransactor) (*CsfeedistributorTransactor, error) {
	contract, err := bindCsfeedistributor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CsfeedistributorTransactor{contract: contract}, nil
}

// NewCsfeedistributorFilterer creates a new log filterer instance of Csfeedistributor, bound to a specific deployed contract.
func NewCsfeedistributorFilterer(address common.Address, filterer bind.ContractFilterer) (*CsfeedistributorFilterer, error) {
	contract, err := bindCsfeedistributor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CsfeedistributorFilterer{contract: contract}, nil
}

// bindCsfeedistributor binds a generic wrapper to an already deployed contract.
func bindCsfeedistributor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CsfeedistributorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Csfeedistributor *CsfeedistributorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Csfeedistributor.Contract.CsfeedistributorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Csfeedistributor *CsfeedistributorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.CsfeedistributorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Csfeedistributor *CsfeedistributorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.CsfeedistributorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Csfeedistributor *CsfeedistributorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Csfeedistributor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Csfeedistributor *CsfeedistributorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Csfeedistributor *CsfeedistributorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.contract.Transact(opts, method, params...)
}

// ACCOUNTING is a free data retrieval call binding the contract method 0x6dc3f2bd.
//
// Solidity: function ACCOUNTING() view returns(address)
func (_Csfeedistributor *CsfeedistributorCaller) ACCOUNTING(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Csfeedistributor.contract.Call(opts, &out, "ACCOUNTING")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ACCOUNTING is a free data retrieval call binding the contract method 0x6dc3f2bd.
//
// Solidity: function ACCOUNTING() view returns(address)
func (_Csfeedistributor *CsfeedistributorSession) ACCOUNTING() (common.Address, error) {
	return _Csfeedistributor.Contract.ACCOUNTING(&_Csfeedistributor.CallOpts)
}

// ACCOUNTING is a free data retrieval call binding the contract method 0x6dc3f2bd.
//
// Solidity: function ACCOUNTING() view returns(address)
func (_Csfeedistributor *CsfeedistributorCallerSession) ACCOUNTING() (common.Address, error) {
	return _Csfeedistributor.Contract.ACCOUNTING(&_Csfeedistributor.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Csfeedistributor *CsfeedistributorCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Csfeedistributor.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Csfeedistributor *CsfeedistributorSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Csfeedistributor.Contract.DEFAULTADMINROLE(&_Csfeedistributor.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Csfeedistributor *CsfeedistributorCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Csfeedistributor.Contract.DEFAULTADMINROLE(&_Csfeedistributor.CallOpts)
}

// ORACLE is a free data retrieval call binding the contract method 0x38013f02.
//
// Solidity: function ORACLE() view returns(address)
func (_Csfeedistributor *CsfeedistributorCaller) ORACLE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Csfeedistributor.contract.Call(opts, &out, "ORACLE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ORACLE is a free data retrieval call binding the contract method 0x38013f02.
//
// Solidity: function ORACLE() view returns(address)
func (_Csfeedistributor *CsfeedistributorSession) ORACLE() (common.Address, error) {
	return _Csfeedistributor.Contract.ORACLE(&_Csfeedistributor.CallOpts)
}

// ORACLE is a free data retrieval call binding the contract method 0x38013f02.
//
// Solidity: function ORACLE() view returns(address)
func (_Csfeedistributor *CsfeedistributorCallerSession) ORACLE() (common.Address, error) {
	return _Csfeedistributor.Contract.ORACLE(&_Csfeedistributor.CallOpts)
}

// RECOVERERROLE is a free data retrieval call binding the contract method 0xacf1c948.
//
// Solidity: function RECOVERER_ROLE() view returns(bytes32)
func (_Csfeedistributor *CsfeedistributorCaller) RECOVERERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Csfeedistributor.contract.Call(opts, &out, "RECOVERER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RECOVERERROLE is a free data retrieval call binding the contract method 0xacf1c948.
//
// Solidity: function RECOVERER_ROLE() view returns(bytes32)
func (_Csfeedistributor *CsfeedistributorSession) RECOVERERROLE() ([32]byte, error) {
	return _Csfeedistributor.Contract.RECOVERERROLE(&_Csfeedistributor.CallOpts)
}

// RECOVERERROLE is a free data retrieval call binding the contract method 0xacf1c948.
//
// Solidity: function RECOVERER_ROLE() view returns(bytes32)
func (_Csfeedistributor *CsfeedistributorCallerSession) RECOVERERROLE() ([32]byte, error) {
	return _Csfeedistributor.Contract.RECOVERERROLE(&_Csfeedistributor.CallOpts)
}

// STETH is a free data retrieval call binding the contract method 0xe00bfe50.
//
// Solidity: function STETH() view returns(address)
func (_Csfeedistributor *CsfeedistributorCaller) STETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Csfeedistributor.contract.Call(opts, &out, "STETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STETH is a free data retrieval call binding the contract method 0xe00bfe50.
//
// Solidity: function STETH() view returns(address)
func (_Csfeedistributor *CsfeedistributorSession) STETH() (common.Address, error) {
	return _Csfeedistributor.Contract.STETH(&_Csfeedistributor.CallOpts)
}

// STETH is a free data retrieval call binding the contract method 0xe00bfe50.
//
// Solidity: function STETH() view returns(address)
func (_Csfeedistributor *CsfeedistributorCallerSession) STETH() (common.Address, error) {
	return _Csfeedistributor.Contract.STETH(&_Csfeedistributor.CallOpts)
}

// DistributedShares is a free data retrieval call binding the contract method 0xea6301ab.
//
// Solidity: function distributedShares(uint256 ) view returns(uint256)
func (_Csfeedistributor *CsfeedistributorCaller) DistributedShares(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Csfeedistributor.contract.Call(opts, &out, "distributedShares", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DistributedShares is a free data retrieval call binding the contract method 0xea6301ab.
//
// Solidity: function distributedShares(uint256 ) view returns(uint256)
func (_Csfeedistributor *CsfeedistributorSession) DistributedShares(arg0 *big.Int) (*big.Int, error) {
	return _Csfeedistributor.Contract.DistributedShares(&_Csfeedistributor.CallOpts, arg0)
}

// DistributedShares is a free data retrieval call binding the contract method 0xea6301ab.
//
// Solidity: function distributedShares(uint256 ) view returns(uint256)
func (_Csfeedistributor *CsfeedistributorCallerSession) DistributedShares(arg0 *big.Int) (*big.Int, error) {
	return _Csfeedistributor.Contract.DistributedShares(&_Csfeedistributor.CallOpts, arg0)
}

// GetFeesToDistribute is a free data retrieval call binding the contract method 0x5e8e8f6f.
//
// Solidity: function getFeesToDistribute(uint256 nodeOperatorId, uint256 shares, bytes32[] proof) view returns(uint256 sharesToDistribute)
func (_Csfeedistributor *CsfeedistributorCaller) GetFeesToDistribute(opts *bind.CallOpts, nodeOperatorId *big.Int, shares *big.Int, proof [][32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Csfeedistributor.contract.Call(opts, &out, "getFeesToDistribute", nodeOperatorId, shares, proof)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFeesToDistribute is a free data retrieval call binding the contract method 0x5e8e8f6f.
//
// Solidity: function getFeesToDistribute(uint256 nodeOperatorId, uint256 shares, bytes32[] proof) view returns(uint256 sharesToDistribute)
func (_Csfeedistributor *CsfeedistributorSession) GetFeesToDistribute(nodeOperatorId *big.Int, shares *big.Int, proof [][32]byte) (*big.Int, error) {
	return _Csfeedistributor.Contract.GetFeesToDistribute(&_Csfeedistributor.CallOpts, nodeOperatorId, shares, proof)
}

// GetFeesToDistribute is a free data retrieval call binding the contract method 0x5e8e8f6f.
//
// Solidity: function getFeesToDistribute(uint256 nodeOperatorId, uint256 shares, bytes32[] proof) view returns(uint256 sharesToDistribute)
func (_Csfeedistributor *CsfeedistributorCallerSession) GetFeesToDistribute(nodeOperatorId *big.Int, shares *big.Int, proof [][32]byte) (*big.Int, error) {
	return _Csfeedistributor.Contract.GetFeesToDistribute(&_Csfeedistributor.CallOpts, nodeOperatorId, shares, proof)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Csfeedistributor *CsfeedistributorCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Csfeedistributor.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Csfeedistributor *CsfeedistributorSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Csfeedistributor.Contract.GetRoleAdmin(&_Csfeedistributor.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Csfeedistributor *CsfeedistributorCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Csfeedistributor.Contract.GetRoleAdmin(&_Csfeedistributor.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Csfeedistributor *CsfeedistributorCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Csfeedistributor.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Csfeedistributor *CsfeedistributorSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Csfeedistributor.Contract.GetRoleMember(&_Csfeedistributor.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Csfeedistributor *CsfeedistributorCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Csfeedistributor.Contract.GetRoleMember(&_Csfeedistributor.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Csfeedistributor *CsfeedistributorCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Csfeedistributor.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Csfeedistributor *CsfeedistributorSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Csfeedistributor.Contract.GetRoleMemberCount(&_Csfeedistributor.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Csfeedistributor *CsfeedistributorCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Csfeedistributor.Contract.GetRoleMemberCount(&_Csfeedistributor.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Csfeedistributor *CsfeedistributorCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Csfeedistributor.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Csfeedistributor *CsfeedistributorSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Csfeedistributor.Contract.HasRole(&_Csfeedistributor.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Csfeedistributor *CsfeedistributorCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Csfeedistributor.Contract.HasRole(&_Csfeedistributor.CallOpts, role, account)
}

// HashLeaf is a free data retrieval call binding the contract method 0x7e9f27ad.
//
// Solidity: function hashLeaf(uint256 nodeOperatorId, uint256 shares) pure returns(bytes32)
func (_Csfeedistributor *CsfeedistributorCaller) HashLeaf(opts *bind.CallOpts, nodeOperatorId *big.Int, shares *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Csfeedistributor.contract.Call(opts, &out, "hashLeaf", nodeOperatorId, shares)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashLeaf is a free data retrieval call binding the contract method 0x7e9f27ad.
//
// Solidity: function hashLeaf(uint256 nodeOperatorId, uint256 shares) pure returns(bytes32)
func (_Csfeedistributor *CsfeedistributorSession) HashLeaf(nodeOperatorId *big.Int, shares *big.Int) ([32]byte, error) {
	return _Csfeedistributor.Contract.HashLeaf(&_Csfeedistributor.CallOpts, nodeOperatorId, shares)
}

// HashLeaf is a free data retrieval call binding the contract method 0x7e9f27ad.
//
// Solidity: function hashLeaf(uint256 nodeOperatorId, uint256 shares) pure returns(bytes32)
func (_Csfeedistributor *CsfeedistributorCallerSession) HashLeaf(nodeOperatorId *big.Int, shares *big.Int) ([32]byte, error) {
	return _Csfeedistributor.Contract.HashLeaf(&_Csfeedistributor.CallOpts, nodeOperatorId, shares)
}

// PendingSharesToDistribute is a free data retrieval call binding the contract method 0xd257cf2a.
//
// Solidity: function pendingSharesToDistribute() view returns(uint256)
func (_Csfeedistributor *CsfeedistributorCaller) PendingSharesToDistribute(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Csfeedistributor.contract.Call(opts, &out, "pendingSharesToDistribute")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingSharesToDistribute is a free data retrieval call binding the contract method 0xd257cf2a.
//
// Solidity: function pendingSharesToDistribute() view returns(uint256)
func (_Csfeedistributor *CsfeedistributorSession) PendingSharesToDistribute() (*big.Int, error) {
	return _Csfeedistributor.Contract.PendingSharesToDistribute(&_Csfeedistributor.CallOpts)
}

// PendingSharesToDistribute is a free data retrieval call binding the contract method 0xd257cf2a.
//
// Solidity: function pendingSharesToDistribute() view returns(uint256)
func (_Csfeedistributor *CsfeedistributorCallerSession) PendingSharesToDistribute() (*big.Int, error) {
	return _Csfeedistributor.Contract.PendingSharesToDistribute(&_Csfeedistributor.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Csfeedistributor *CsfeedistributorCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Csfeedistributor.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Csfeedistributor *CsfeedistributorSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Csfeedistributor.Contract.SupportsInterface(&_Csfeedistributor.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Csfeedistributor *CsfeedistributorCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Csfeedistributor.Contract.SupportsInterface(&_Csfeedistributor.CallOpts, interfaceId)
}

// TotalClaimableShares is a free data retrieval call binding the contract method 0x47d17d9d.
//
// Solidity: function totalClaimableShares() view returns(uint256)
func (_Csfeedistributor *CsfeedistributorCaller) TotalClaimableShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Csfeedistributor.contract.Call(opts, &out, "totalClaimableShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalClaimableShares is a free data retrieval call binding the contract method 0x47d17d9d.
//
// Solidity: function totalClaimableShares() view returns(uint256)
func (_Csfeedistributor *CsfeedistributorSession) TotalClaimableShares() (*big.Int, error) {
	return _Csfeedistributor.Contract.TotalClaimableShares(&_Csfeedistributor.CallOpts)
}

// TotalClaimableShares is a free data retrieval call binding the contract method 0x47d17d9d.
//
// Solidity: function totalClaimableShares() view returns(uint256)
func (_Csfeedistributor *CsfeedistributorCallerSession) TotalClaimableShares() (*big.Int, error) {
	return _Csfeedistributor.Contract.TotalClaimableShares(&_Csfeedistributor.CallOpts)
}

// TreeCid is a free data retrieval call binding the contract method 0xfe3c9b9b.
//
// Solidity: function treeCid() view returns(string)
func (_Csfeedistributor *CsfeedistributorCaller) TreeCid(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Csfeedistributor.contract.Call(opts, &out, "treeCid")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TreeCid is a free data retrieval call binding the contract method 0xfe3c9b9b.
//
// Solidity: function treeCid() view returns(string)
func (_Csfeedistributor *CsfeedistributorSession) TreeCid() (string, error) {
	return _Csfeedistributor.Contract.TreeCid(&_Csfeedistributor.CallOpts)
}

// TreeCid is a free data retrieval call binding the contract method 0xfe3c9b9b.
//
// Solidity: function treeCid() view returns(string)
func (_Csfeedistributor *CsfeedistributorCallerSession) TreeCid() (string, error) {
	return _Csfeedistributor.Contract.TreeCid(&_Csfeedistributor.CallOpts)
}

// TreeRoot is a free data retrieval call binding the contract method 0x14dc6c14.
//
// Solidity: function treeRoot() view returns(bytes32)
func (_Csfeedistributor *CsfeedistributorCaller) TreeRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Csfeedistributor.contract.Call(opts, &out, "treeRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TreeRoot is a free data retrieval call binding the contract method 0x14dc6c14.
//
// Solidity: function treeRoot() view returns(bytes32)
func (_Csfeedistributor *CsfeedistributorSession) TreeRoot() ([32]byte, error) {
	return _Csfeedistributor.Contract.TreeRoot(&_Csfeedistributor.CallOpts)
}

// TreeRoot is a free data retrieval call binding the contract method 0x14dc6c14.
//
// Solidity: function treeRoot() view returns(bytes32)
func (_Csfeedistributor *CsfeedistributorCallerSession) TreeRoot() ([32]byte, error) {
	return _Csfeedistributor.Contract.TreeRoot(&_Csfeedistributor.CallOpts)
}

// DistributeFees is a paid mutator transaction binding the contract method 0x21893f7b.
//
// Solidity: function distributeFees(uint256 nodeOperatorId, uint256 shares, bytes32[] proof) returns(uint256 sharesToDistribute)
func (_Csfeedistributor *CsfeedistributorTransactor) DistributeFees(opts *bind.TransactOpts, nodeOperatorId *big.Int, shares *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _Csfeedistributor.contract.Transact(opts, "distributeFees", nodeOperatorId, shares, proof)
}

// DistributeFees is a paid mutator transaction binding the contract method 0x21893f7b.
//
// Solidity: function distributeFees(uint256 nodeOperatorId, uint256 shares, bytes32[] proof) returns(uint256 sharesToDistribute)
func (_Csfeedistributor *CsfeedistributorSession) DistributeFees(nodeOperatorId *big.Int, shares *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.DistributeFees(&_Csfeedistributor.TransactOpts, nodeOperatorId, shares, proof)
}

// DistributeFees is a paid mutator transaction binding the contract method 0x21893f7b.
//
// Solidity: function distributeFees(uint256 nodeOperatorId, uint256 shares, bytes32[] proof) returns(uint256 sharesToDistribute)
func (_Csfeedistributor *CsfeedistributorTransactorSession) DistributeFees(nodeOperatorId *big.Int, shares *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.DistributeFees(&_Csfeedistributor.TransactOpts, nodeOperatorId, shares, proof)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Csfeedistributor *CsfeedistributorTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Csfeedistributor.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Csfeedistributor *CsfeedistributorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.GrantRole(&_Csfeedistributor.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Csfeedistributor *CsfeedistributorTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.GrantRole(&_Csfeedistributor.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address admin) returns()
func (_Csfeedistributor *CsfeedistributorTransactor) Initialize(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _Csfeedistributor.contract.Transact(opts, "initialize", admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address admin) returns()
func (_Csfeedistributor *CsfeedistributorSession) Initialize(admin common.Address) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.Initialize(&_Csfeedistributor.TransactOpts, admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address admin) returns()
func (_Csfeedistributor *CsfeedistributorTransactorSession) Initialize(admin common.Address) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.Initialize(&_Csfeedistributor.TransactOpts, admin)
}

// ProcessOracleReport is a paid mutator transaction binding the contract method 0xb66cf058.
//
// Solidity: function processOracleReport(bytes32 _treeRoot, string _treeCid, uint256 distributed) returns()
func (_Csfeedistributor *CsfeedistributorTransactor) ProcessOracleReport(opts *bind.TransactOpts, _treeRoot [32]byte, _treeCid string, distributed *big.Int) (*types.Transaction, error) {
	return _Csfeedistributor.contract.Transact(opts, "processOracleReport", _treeRoot, _treeCid, distributed)
}

// ProcessOracleReport is a paid mutator transaction binding the contract method 0xb66cf058.
//
// Solidity: function processOracleReport(bytes32 _treeRoot, string _treeCid, uint256 distributed) returns()
func (_Csfeedistributor *CsfeedistributorSession) ProcessOracleReport(_treeRoot [32]byte, _treeCid string, distributed *big.Int) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.ProcessOracleReport(&_Csfeedistributor.TransactOpts, _treeRoot, _treeCid, distributed)
}

// ProcessOracleReport is a paid mutator transaction binding the contract method 0xb66cf058.
//
// Solidity: function processOracleReport(bytes32 _treeRoot, string _treeCid, uint256 distributed) returns()
func (_Csfeedistributor *CsfeedistributorTransactorSession) ProcessOracleReport(_treeRoot [32]byte, _treeCid string, distributed *big.Int) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.ProcessOracleReport(&_Csfeedistributor.TransactOpts, _treeRoot, _treeCid, distributed)
}

// RecoverERC1155 is a paid mutator transaction binding the contract method 0x5c654ad9.
//
// Solidity: function recoverERC1155(address token, uint256 tokenId) returns()
func (_Csfeedistributor *CsfeedistributorTransactor) RecoverERC1155(opts *bind.TransactOpts, token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Csfeedistributor.contract.Transact(opts, "recoverERC1155", token, tokenId)
}

// RecoverERC1155 is a paid mutator transaction binding the contract method 0x5c654ad9.
//
// Solidity: function recoverERC1155(address token, uint256 tokenId) returns()
func (_Csfeedistributor *CsfeedistributorSession) RecoverERC1155(token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.RecoverERC1155(&_Csfeedistributor.TransactOpts, token, tokenId)
}

// RecoverERC1155 is a paid mutator transaction binding the contract method 0x5c654ad9.
//
// Solidity: function recoverERC1155(address token, uint256 tokenId) returns()
func (_Csfeedistributor *CsfeedistributorTransactorSession) RecoverERC1155(token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.RecoverERC1155(&_Csfeedistributor.TransactOpts, token, tokenId)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address token, uint256 amount) returns()
func (_Csfeedistributor *CsfeedistributorTransactor) RecoverERC20(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Csfeedistributor.contract.Transact(opts, "recoverERC20", token, amount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address token, uint256 amount) returns()
func (_Csfeedistributor *CsfeedistributorSession) RecoverERC20(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.RecoverERC20(&_Csfeedistributor.TransactOpts, token, amount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address token, uint256 amount) returns()
func (_Csfeedistributor *CsfeedistributorTransactorSession) RecoverERC20(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.RecoverERC20(&_Csfeedistributor.TransactOpts, token, amount)
}

// RecoverERC721 is a paid mutator transaction binding the contract method 0x819d4cc6.
//
// Solidity: function recoverERC721(address token, uint256 tokenId) returns()
func (_Csfeedistributor *CsfeedistributorTransactor) RecoverERC721(opts *bind.TransactOpts, token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Csfeedistributor.contract.Transact(opts, "recoverERC721", token, tokenId)
}

// RecoverERC721 is a paid mutator transaction binding the contract method 0x819d4cc6.
//
// Solidity: function recoverERC721(address token, uint256 tokenId) returns()
func (_Csfeedistributor *CsfeedistributorSession) RecoverERC721(token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.RecoverERC721(&_Csfeedistributor.TransactOpts, token, tokenId)
}

// RecoverERC721 is a paid mutator transaction binding the contract method 0x819d4cc6.
//
// Solidity: function recoverERC721(address token, uint256 tokenId) returns()
func (_Csfeedistributor *CsfeedistributorTransactorSession) RecoverERC721(token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.RecoverERC721(&_Csfeedistributor.TransactOpts, token, tokenId)
}

// RecoverEther is a paid mutator transaction binding the contract method 0x52d8bfc2.
//
// Solidity: function recoverEther() returns()
func (_Csfeedistributor *CsfeedistributorTransactor) RecoverEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Csfeedistributor.contract.Transact(opts, "recoverEther")
}

// RecoverEther is a paid mutator transaction binding the contract method 0x52d8bfc2.
//
// Solidity: function recoverEther() returns()
func (_Csfeedistributor *CsfeedistributorSession) RecoverEther() (*types.Transaction, error) {
	return _Csfeedistributor.Contract.RecoverEther(&_Csfeedistributor.TransactOpts)
}

// RecoverEther is a paid mutator transaction binding the contract method 0x52d8bfc2.
//
// Solidity: function recoverEther() returns()
func (_Csfeedistributor *CsfeedistributorTransactorSession) RecoverEther() (*types.Transaction, error) {
	return _Csfeedistributor.Contract.RecoverEther(&_Csfeedistributor.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Csfeedistributor *CsfeedistributorTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Csfeedistributor.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Csfeedistributor *CsfeedistributorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.RenounceRole(&_Csfeedistributor.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Csfeedistributor *CsfeedistributorTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.RenounceRole(&_Csfeedistributor.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Csfeedistributor *CsfeedistributorTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Csfeedistributor.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Csfeedistributor *CsfeedistributorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.RevokeRole(&_Csfeedistributor.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Csfeedistributor *CsfeedistributorTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.RevokeRole(&_Csfeedistributor.TransactOpts, role, account)
}

// CsfeedistributorDistributionDataUpdatedIterator is returned from FilterDistributionDataUpdated and is used to iterate over the raw logs and unpacked data for DistributionDataUpdated events raised by the Csfeedistributor contract.
type CsfeedistributorDistributionDataUpdatedIterator struct {
	Event *CsfeedistributorDistributionDataUpdated // Event containing the contract specifics and raw log

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
func (it *CsfeedistributorDistributionDataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsfeedistributorDistributionDataUpdated)
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
		it.Event = new(CsfeedistributorDistributionDataUpdated)
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
func (it *CsfeedistributorDistributionDataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsfeedistributorDistributionDataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsfeedistributorDistributionDataUpdated represents a DistributionDataUpdated event raised by the Csfeedistributor contract.
type CsfeedistributorDistributionDataUpdated struct {
	TotalClaimableShares *big.Int
	TreeRoot             [32]byte
	TreeCid              string
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterDistributionDataUpdated is a free log retrieval operation binding the contract event 0x26dec7cc117e9b3907dc1f90d2dc5f6e04dbb9f285f5898be2c82ec524dcd424.
//
// Solidity: event DistributionDataUpdated(uint256 totalClaimableShares, bytes32 treeRoot, string treeCid)
func (_Csfeedistributor *CsfeedistributorFilterer) FilterDistributionDataUpdated(opts *bind.FilterOpts) (*CsfeedistributorDistributionDataUpdatedIterator, error) {

	logs, sub, err := _Csfeedistributor.contract.FilterLogs(opts, "DistributionDataUpdated")
	if err != nil {
		return nil, err
	}
	return &CsfeedistributorDistributionDataUpdatedIterator{contract: _Csfeedistributor.contract, event: "DistributionDataUpdated", logs: logs, sub: sub}, nil
}

// WatchDistributionDataUpdated is a free log subscription operation binding the contract event 0x26dec7cc117e9b3907dc1f90d2dc5f6e04dbb9f285f5898be2c82ec524dcd424.
//
// Solidity: event DistributionDataUpdated(uint256 totalClaimableShares, bytes32 treeRoot, string treeCid)
func (_Csfeedistributor *CsfeedistributorFilterer) WatchDistributionDataUpdated(opts *bind.WatchOpts, sink chan<- *CsfeedistributorDistributionDataUpdated) (event.Subscription, error) {

	logs, sub, err := _Csfeedistributor.contract.WatchLogs(opts, "DistributionDataUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsfeedistributorDistributionDataUpdated)
				if err := _Csfeedistributor.contract.UnpackLog(event, "DistributionDataUpdated", log); err != nil {
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
func (_Csfeedistributor *CsfeedistributorFilterer) ParseDistributionDataUpdated(log types.Log) (*CsfeedistributorDistributionDataUpdated, error) {
	event := new(CsfeedistributorDistributionDataUpdated)
	if err := _Csfeedistributor.contract.UnpackLog(event, "DistributionDataUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsfeedistributorERC1155RecoveredIterator is returned from FilterERC1155Recovered and is used to iterate over the raw logs and unpacked data for ERC1155Recovered events raised by the Csfeedistributor contract.
type CsfeedistributorERC1155RecoveredIterator struct {
	Event *CsfeedistributorERC1155Recovered // Event containing the contract specifics and raw log

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
func (it *CsfeedistributorERC1155RecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsfeedistributorERC1155Recovered)
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
		it.Event = new(CsfeedistributorERC1155Recovered)
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
func (it *CsfeedistributorERC1155RecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsfeedistributorERC1155RecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsfeedistributorERC1155Recovered represents a ERC1155Recovered event raised by the Csfeedistributor contract.
type CsfeedistributorERC1155Recovered struct {
	Token     common.Address
	TokenId   *big.Int
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERC1155Recovered is a free log retrieval operation binding the contract event 0x5cf02e753b3eb0f4bee4460a72817d8e5e3c75cd4d65c1d0b06dca88b8032936.
//
// Solidity: event ERC1155Recovered(address indexed token, uint256 tokenId, address indexed recipient, uint256 amount)
func (_Csfeedistributor *CsfeedistributorFilterer) FilterERC1155Recovered(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*CsfeedistributorERC1155RecoveredIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csfeedistributor.contract.FilterLogs(opts, "ERC1155Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &CsfeedistributorERC1155RecoveredIterator{contract: _Csfeedistributor.contract, event: "ERC1155Recovered", logs: logs, sub: sub}, nil
}

// WatchERC1155Recovered is a free log subscription operation binding the contract event 0x5cf02e753b3eb0f4bee4460a72817d8e5e3c75cd4d65c1d0b06dca88b8032936.
//
// Solidity: event ERC1155Recovered(address indexed token, uint256 tokenId, address indexed recipient, uint256 amount)
func (_Csfeedistributor *CsfeedistributorFilterer) WatchERC1155Recovered(opts *bind.WatchOpts, sink chan<- *CsfeedistributorERC1155Recovered, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csfeedistributor.contract.WatchLogs(opts, "ERC1155Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsfeedistributorERC1155Recovered)
				if err := _Csfeedistributor.contract.UnpackLog(event, "ERC1155Recovered", log); err != nil {
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
func (_Csfeedistributor *CsfeedistributorFilterer) ParseERC1155Recovered(log types.Log) (*CsfeedistributorERC1155Recovered, error) {
	event := new(CsfeedistributorERC1155Recovered)
	if err := _Csfeedistributor.contract.UnpackLog(event, "ERC1155Recovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsfeedistributorERC20RecoveredIterator is returned from FilterERC20Recovered and is used to iterate over the raw logs and unpacked data for ERC20Recovered events raised by the Csfeedistributor contract.
type CsfeedistributorERC20RecoveredIterator struct {
	Event *CsfeedistributorERC20Recovered // Event containing the contract specifics and raw log

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
func (it *CsfeedistributorERC20RecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsfeedistributorERC20Recovered)
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
		it.Event = new(CsfeedistributorERC20Recovered)
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
func (it *CsfeedistributorERC20RecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsfeedistributorERC20RecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsfeedistributorERC20Recovered represents a ERC20Recovered event raised by the Csfeedistributor contract.
type CsfeedistributorERC20Recovered struct {
	Token     common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERC20Recovered is a free log retrieval operation binding the contract event 0xaca8fb252cde442184e5f10e0f2e6e4029e8cd7717cae63559079610702436aa.
//
// Solidity: event ERC20Recovered(address indexed token, address indexed recipient, uint256 amount)
func (_Csfeedistributor *CsfeedistributorFilterer) FilterERC20Recovered(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*CsfeedistributorERC20RecoveredIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csfeedistributor.contract.FilterLogs(opts, "ERC20Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &CsfeedistributorERC20RecoveredIterator{contract: _Csfeedistributor.contract, event: "ERC20Recovered", logs: logs, sub: sub}, nil
}

// WatchERC20Recovered is a free log subscription operation binding the contract event 0xaca8fb252cde442184e5f10e0f2e6e4029e8cd7717cae63559079610702436aa.
//
// Solidity: event ERC20Recovered(address indexed token, address indexed recipient, uint256 amount)
func (_Csfeedistributor *CsfeedistributorFilterer) WatchERC20Recovered(opts *bind.WatchOpts, sink chan<- *CsfeedistributorERC20Recovered, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csfeedistributor.contract.WatchLogs(opts, "ERC20Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsfeedistributorERC20Recovered)
				if err := _Csfeedistributor.contract.UnpackLog(event, "ERC20Recovered", log); err != nil {
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
func (_Csfeedistributor *CsfeedistributorFilterer) ParseERC20Recovered(log types.Log) (*CsfeedistributorERC20Recovered, error) {
	event := new(CsfeedistributorERC20Recovered)
	if err := _Csfeedistributor.contract.UnpackLog(event, "ERC20Recovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsfeedistributorERC721RecoveredIterator is returned from FilterERC721Recovered and is used to iterate over the raw logs and unpacked data for ERC721Recovered events raised by the Csfeedistributor contract.
type CsfeedistributorERC721RecoveredIterator struct {
	Event *CsfeedistributorERC721Recovered // Event containing the contract specifics and raw log

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
func (it *CsfeedistributorERC721RecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsfeedistributorERC721Recovered)
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
		it.Event = new(CsfeedistributorERC721Recovered)
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
func (it *CsfeedistributorERC721RecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsfeedistributorERC721RecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsfeedistributorERC721Recovered represents a ERC721Recovered event raised by the Csfeedistributor contract.
type CsfeedistributorERC721Recovered struct {
	Token     common.Address
	TokenId   *big.Int
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERC721Recovered is a free log retrieval operation binding the contract event 0x8166bf75d2ff2fa3c8f3c44410540bf42e9a5359b48409e8d660291dc9f788c8.
//
// Solidity: event ERC721Recovered(address indexed token, uint256 tokenId, address indexed recipient)
func (_Csfeedistributor *CsfeedistributorFilterer) FilterERC721Recovered(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*CsfeedistributorERC721RecoveredIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csfeedistributor.contract.FilterLogs(opts, "ERC721Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &CsfeedistributorERC721RecoveredIterator{contract: _Csfeedistributor.contract, event: "ERC721Recovered", logs: logs, sub: sub}, nil
}

// WatchERC721Recovered is a free log subscription operation binding the contract event 0x8166bf75d2ff2fa3c8f3c44410540bf42e9a5359b48409e8d660291dc9f788c8.
//
// Solidity: event ERC721Recovered(address indexed token, uint256 tokenId, address indexed recipient)
func (_Csfeedistributor *CsfeedistributorFilterer) WatchERC721Recovered(opts *bind.WatchOpts, sink chan<- *CsfeedistributorERC721Recovered, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csfeedistributor.contract.WatchLogs(opts, "ERC721Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsfeedistributorERC721Recovered)
				if err := _Csfeedistributor.contract.UnpackLog(event, "ERC721Recovered", log); err != nil {
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
func (_Csfeedistributor *CsfeedistributorFilterer) ParseERC721Recovered(log types.Log) (*CsfeedistributorERC721Recovered, error) {
	event := new(CsfeedistributorERC721Recovered)
	if err := _Csfeedistributor.contract.UnpackLog(event, "ERC721Recovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsfeedistributorEtherRecoveredIterator is returned from FilterEtherRecovered and is used to iterate over the raw logs and unpacked data for EtherRecovered events raised by the Csfeedistributor contract.
type CsfeedistributorEtherRecoveredIterator struct {
	Event *CsfeedistributorEtherRecovered // Event containing the contract specifics and raw log

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
func (it *CsfeedistributorEtherRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsfeedistributorEtherRecovered)
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
		it.Event = new(CsfeedistributorEtherRecovered)
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
func (it *CsfeedistributorEtherRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsfeedistributorEtherRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsfeedistributorEtherRecovered represents a EtherRecovered event raised by the Csfeedistributor contract.
type CsfeedistributorEtherRecovered struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEtherRecovered is a free log retrieval operation binding the contract event 0x8e274e42262a7f013b700b35c2b4629ccce1702f8fe83f8dfb7eacbb26a4382c.
//
// Solidity: event EtherRecovered(address indexed recipient, uint256 amount)
func (_Csfeedistributor *CsfeedistributorFilterer) FilterEtherRecovered(opts *bind.FilterOpts, recipient []common.Address) (*CsfeedistributorEtherRecoveredIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csfeedistributor.contract.FilterLogs(opts, "EtherRecovered", recipientRule)
	if err != nil {
		return nil, err
	}
	return &CsfeedistributorEtherRecoveredIterator{contract: _Csfeedistributor.contract, event: "EtherRecovered", logs: logs, sub: sub}, nil
}

// WatchEtherRecovered is a free log subscription operation binding the contract event 0x8e274e42262a7f013b700b35c2b4629ccce1702f8fe83f8dfb7eacbb26a4382c.
//
// Solidity: event EtherRecovered(address indexed recipient, uint256 amount)
func (_Csfeedistributor *CsfeedistributorFilterer) WatchEtherRecovered(opts *bind.WatchOpts, sink chan<- *CsfeedistributorEtherRecovered, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csfeedistributor.contract.WatchLogs(opts, "EtherRecovered", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsfeedistributorEtherRecovered)
				if err := _Csfeedistributor.contract.UnpackLog(event, "EtherRecovered", log); err != nil {
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
func (_Csfeedistributor *CsfeedistributorFilterer) ParseEtherRecovered(log types.Log) (*CsfeedistributorEtherRecovered, error) {
	event := new(CsfeedistributorEtherRecovered)
	if err := _Csfeedistributor.contract.UnpackLog(event, "EtherRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsfeedistributorFeeDistributedIterator is returned from FilterFeeDistributed and is used to iterate over the raw logs and unpacked data for FeeDistributed events raised by the Csfeedistributor contract.
type CsfeedistributorFeeDistributedIterator struct {
	Event *CsfeedistributorFeeDistributed // Event containing the contract specifics and raw log

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
func (it *CsfeedistributorFeeDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsfeedistributorFeeDistributed)
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
		it.Event = new(CsfeedistributorFeeDistributed)
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
func (it *CsfeedistributorFeeDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsfeedistributorFeeDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsfeedistributorFeeDistributed represents a FeeDistributed event raised by the Csfeedistributor contract.
type CsfeedistributorFeeDistributed struct {
	NodeOperatorId *big.Int
	Shares         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterFeeDistributed is a free log retrieval operation binding the contract event 0x61930a6c1553eab59d5766da6e1bab8eba982aec848ae7683452f4a6423b6e4a.
//
// Solidity: event FeeDistributed(uint256 indexed nodeOperatorId, uint256 shares)
func (_Csfeedistributor *CsfeedistributorFilterer) FilterFeeDistributed(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsfeedistributorFeeDistributedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csfeedistributor.contract.FilterLogs(opts, "FeeDistributed", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsfeedistributorFeeDistributedIterator{contract: _Csfeedistributor.contract, event: "FeeDistributed", logs: logs, sub: sub}, nil
}

// WatchFeeDistributed is a free log subscription operation binding the contract event 0x61930a6c1553eab59d5766da6e1bab8eba982aec848ae7683452f4a6423b6e4a.
//
// Solidity: event FeeDistributed(uint256 indexed nodeOperatorId, uint256 shares)
func (_Csfeedistributor *CsfeedistributorFilterer) WatchFeeDistributed(opts *bind.WatchOpts, sink chan<- *CsfeedistributorFeeDistributed, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csfeedistributor.contract.WatchLogs(opts, "FeeDistributed", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsfeedistributorFeeDistributed)
				if err := _Csfeedistributor.contract.UnpackLog(event, "FeeDistributed", log); err != nil {
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

// ParseFeeDistributed is a log parse operation binding the contract event 0x61930a6c1553eab59d5766da6e1bab8eba982aec848ae7683452f4a6423b6e4a.
//
// Solidity: event FeeDistributed(uint256 indexed nodeOperatorId, uint256 shares)
func (_Csfeedistributor *CsfeedistributorFilterer) ParseFeeDistributed(log types.Log) (*CsfeedistributorFeeDistributed, error) {
	event := new(CsfeedistributorFeeDistributed)
	if err := _Csfeedistributor.contract.UnpackLog(event, "FeeDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsfeedistributorInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Csfeedistributor contract.
type CsfeedistributorInitializedIterator struct {
	Event *CsfeedistributorInitialized // Event containing the contract specifics and raw log

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
func (it *CsfeedistributorInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsfeedistributorInitialized)
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
		it.Event = new(CsfeedistributorInitialized)
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
func (it *CsfeedistributorInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsfeedistributorInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsfeedistributorInitialized represents a Initialized event raised by the Csfeedistributor contract.
type CsfeedistributorInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Csfeedistributor *CsfeedistributorFilterer) FilterInitialized(opts *bind.FilterOpts) (*CsfeedistributorInitializedIterator, error) {

	logs, sub, err := _Csfeedistributor.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CsfeedistributorInitializedIterator{contract: _Csfeedistributor.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Csfeedistributor *CsfeedistributorFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CsfeedistributorInitialized) (event.Subscription, error) {

	logs, sub, err := _Csfeedistributor.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsfeedistributorInitialized)
				if err := _Csfeedistributor.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Csfeedistributor *CsfeedistributorFilterer) ParseInitialized(log types.Log) (*CsfeedistributorInitialized, error) {
	event := new(CsfeedistributorInitialized)
	if err := _Csfeedistributor.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsfeedistributorRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Csfeedistributor contract.
type CsfeedistributorRoleAdminChangedIterator struct {
	Event *CsfeedistributorRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *CsfeedistributorRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsfeedistributorRoleAdminChanged)
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
		it.Event = new(CsfeedistributorRoleAdminChanged)
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
func (it *CsfeedistributorRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsfeedistributorRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsfeedistributorRoleAdminChanged represents a RoleAdminChanged event raised by the Csfeedistributor contract.
type CsfeedistributorRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Csfeedistributor *CsfeedistributorFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*CsfeedistributorRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Csfeedistributor.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &CsfeedistributorRoleAdminChangedIterator{contract: _Csfeedistributor.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Csfeedistributor *CsfeedistributorFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *CsfeedistributorRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Csfeedistributor.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsfeedistributorRoleAdminChanged)
				if err := _Csfeedistributor.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Csfeedistributor *CsfeedistributorFilterer) ParseRoleAdminChanged(log types.Log) (*CsfeedistributorRoleAdminChanged, error) {
	event := new(CsfeedistributorRoleAdminChanged)
	if err := _Csfeedistributor.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsfeedistributorRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Csfeedistributor contract.
type CsfeedistributorRoleGrantedIterator struct {
	Event *CsfeedistributorRoleGranted // Event containing the contract specifics and raw log

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
func (it *CsfeedistributorRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsfeedistributorRoleGranted)
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
		it.Event = new(CsfeedistributorRoleGranted)
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
func (it *CsfeedistributorRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsfeedistributorRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsfeedistributorRoleGranted represents a RoleGranted event raised by the Csfeedistributor contract.
type CsfeedistributorRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Csfeedistributor *CsfeedistributorFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CsfeedistributorRoleGrantedIterator, error) {

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

	logs, sub, err := _Csfeedistributor.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CsfeedistributorRoleGrantedIterator{contract: _Csfeedistributor.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Csfeedistributor *CsfeedistributorFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *CsfeedistributorRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Csfeedistributor.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsfeedistributorRoleGranted)
				if err := _Csfeedistributor.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Csfeedistributor *CsfeedistributorFilterer) ParseRoleGranted(log types.Log) (*CsfeedistributorRoleGranted, error) {
	event := new(CsfeedistributorRoleGranted)
	if err := _Csfeedistributor.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsfeedistributorRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Csfeedistributor contract.
type CsfeedistributorRoleRevokedIterator struct {
	Event *CsfeedistributorRoleRevoked // Event containing the contract specifics and raw log

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
func (it *CsfeedistributorRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsfeedistributorRoleRevoked)
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
		it.Event = new(CsfeedistributorRoleRevoked)
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
func (it *CsfeedistributorRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsfeedistributorRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsfeedistributorRoleRevoked represents a RoleRevoked event raised by the Csfeedistributor contract.
type CsfeedistributorRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Csfeedistributor *CsfeedistributorFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CsfeedistributorRoleRevokedIterator, error) {

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

	logs, sub, err := _Csfeedistributor.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CsfeedistributorRoleRevokedIterator{contract: _Csfeedistributor.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Csfeedistributor *CsfeedistributorFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *CsfeedistributorRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Csfeedistributor.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsfeedistributorRoleRevoked)
				if err := _Csfeedistributor.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Csfeedistributor *CsfeedistributorFilterer) ParseRoleRevoked(log types.Log) (*CsfeedistributorRoleRevoked, error) {
	event := new(CsfeedistributorRoleRevoked)
	if err := _Csfeedistributor.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsfeedistributorStETHSharesRecoveredIterator is returned from FilterStETHSharesRecovered and is used to iterate over the raw logs and unpacked data for StETHSharesRecovered events raised by the Csfeedistributor contract.
type CsfeedistributorStETHSharesRecoveredIterator struct {
	Event *CsfeedistributorStETHSharesRecovered // Event containing the contract specifics and raw log

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
func (it *CsfeedistributorStETHSharesRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsfeedistributorStETHSharesRecovered)
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
		it.Event = new(CsfeedistributorStETHSharesRecovered)
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
func (it *CsfeedistributorStETHSharesRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsfeedistributorStETHSharesRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsfeedistributorStETHSharesRecovered represents a StETHSharesRecovered event raised by the Csfeedistributor contract.
type CsfeedistributorStETHSharesRecovered struct {
	Recipient common.Address
	Shares    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStETHSharesRecovered is a free log retrieval operation binding the contract event 0x426e7e0100db57255d4af4a46cd49552ef74f5f002bbdc8d4ebb6371c0070a02.
//
// Solidity: event StETHSharesRecovered(address indexed recipient, uint256 shares)
func (_Csfeedistributor *CsfeedistributorFilterer) FilterStETHSharesRecovered(opts *bind.FilterOpts, recipient []common.Address) (*CsfeedistributorStETHSharesRecoveredIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csfeedistributor.contract.FilterLogs(opts, "StETHSharesRecovered", recipientRule)
	if err != nil {
		return nil, err
	}
	return &CsfeedistributorStETHSharesRecoveredIterator{contract: _Csfeedistributor.contract, event: "StETHSharesRecovered", logs: logs, sub: sub}, nil
}

// WatchStETHSharesRecovered is a free log subscription operation binding the contract event 0x426e7e0100db57255d4af4a46cd49552ef74f5f002bbdc8d4ebb6371c0070a02.
//
// Solidity: event StETHSharesRecovered(address indexed recipient, uint256 shares)
func (_Csfeedistributor *CsfeedistributorFilterer) WatchStETHSharesRecovered(opts *bind.WatchOpts, sink chan<- *CsfeedistributorStETHSharesRecovered, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csfeedistributor.contract.WatchLogs(opts, "StETHSharesRecovered", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsfeedistributorStETHSharesRecovered)
				if err := _Csfeedistributor.contract.UnpackLog(event, "StETHSharesRecovered", log); err != nil {
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
func (_Csfeedistributor *CsfeedistributorFilterer) ParseStETHSharesRecovered(log types.Log) (*CsfeedistributorStETHSharesRecovered, error) {
	event := new(CsfeedistributorStETHSharesRecovered)
	if err := _Csfeedistributor.contract.UnpackLog(event, "StETHSharesRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
