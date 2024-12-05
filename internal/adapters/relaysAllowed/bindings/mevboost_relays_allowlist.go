// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

		"lido-events/internal/application/domain"

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

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
// type Struct0 struct {
// 	Uri         string
// 	Operator    string
// 	IsMandatory bool
// 	Description string
// }

// BindingsMetaData contains all meta data concerning the Bindings contract.
var BindingsMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"RelayAdded\",\"inputs\":[{\"name\":\"uri_hash\",\"type\":\"string\",\"indexed\":true},{\"name\":\"relay\",\"type\":\"tuple\",\"components\":[{\"name\":\"uri\",\"type\":\"string\"},{\"name\":\"operator\",\"type\":\"string\"},{\"name\":\"is_mandatory\",\"type\":\"bool\"},{\"name\":\"description\",\"type\":\"string\"}],\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RelayRemoved\",\"inputs\":[{\"name\":\"uri_hash\",\"type\":\"string\",\"indexed\":true},{\"name\":\"uri\",\"type\":\"string\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"AllowedListUpdated\",\"inputs\":[{\"name\":\"allowed_list_version\",\"type\":\"uint256\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"OwnerChanged\",\"inputs\":[{\"name\":\"new_owner\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ManagerChanged\",\"inputs\":[{\"name\":\"new_manager\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ERC20Recovered\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_relays_amount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_manager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_relays\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"components\":[{\"name\":\"uri\",\"type\":\"string\"},{\"name\":\"operator\",\"type\":\"string\"},{\"name\":\"is_mandatory\",\"type\":\"bool\"},{\"name\":\"description\",\"type\":\"string\"}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_relay_by_uri\",\"inputs\":[{\"name\":\"relay_uri\",\"type\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"components\":[{\"name\":\"uri\",\"type\":\"string\"},{\"name\":\"operator\",\"type\":\"string\"},{\"name\":\"is_mandatory\",\"type\":\"bool\"},{\"name\":\"description\",\"type\":\"string\"}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_allowed_list_version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_relay\",\"inputs\":[{\"name\":\"uri\",\"type\":\"string\"},{\"name\":\"operator\",\"type\":\"string\"},{\"name\":\"is_mandatory\",\"type\":\"bool\"},{\"name\":\"description\",\"type\":\"string\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_relay\",\"inputs\":[{\"name\":\"uri\",\"type\":\"string\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"change_owner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_manager\",\"inputs\":[{\"name\":\"manager\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"dismiss_manager\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"recover_erc20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"}]",
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

// GetAllowedListVersion is a free data retrieval call binding the contract method 0x76650ad3.
//
// Solidity: function get_allowed_list_version() view returns(uint256)
func (_Bindings *BindingsCaller) GetAllowedListVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "get_allowed_list_version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAllowedListVersion is a free data retrieval call binding the contract method 0x76650ad3.
//
// Solidity: function get_allowed_list_version() view returns(uint256)
func (_Bindings *BindingsSession) GetAllowedListVersion() (*big.Int, error) {
	return _Bindings.Contract.GetAllowedListVersion(&_Bindings.CallOpts)
}

// GetAllowedListVersion is a free data retrieval call binding the contract method 0x76650ad3.
//
// Solidity: function get_allowed_list_version() view returns(uint256)
func (_Bindings *BindingsCallerSession) GetAllowedListVersion() (*big.Int, error) {
	return _Bindings.Contract.GetAllowedListVersion(&_Bindings.CallOpts)
}

// GetManager is a free data retrieval call binding the contract method 0x9e4a0fc4.
//
// Solidity: function get_manager() view returns(address)
func (_Bindings *BindingsCaller) GetManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "get_manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetManager is a free data retrieval call binding the contract method 0x9e4a0fc4.
//
// Solidity: function get_manager() view returns(address)
func (_Bindings *BindingsSession) GetManager() (common.Address, error) {
	return _Bindings.Contract.GetManager(&_Bindings.CallOpts)
}

// GetManager is a free data retrieval call binding the contract method 0x9e4a0fc4.
//
// Solidity: function get_manager() view returns(address)
func (_Bindings *BindingsCallerSession) GetManager() (common.Address, error) {
	return _Bindings.Contract.GetManager(&_Bindings.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x0ac298dc.
//
// Solidity: function get_owner() view returns(address)
func (_Bindings *BindingsCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "get_owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x0ac298dc.
//
// Solidity: function get_owner() view returns(address)
func (_Bindings *BindingsSession) GetOwner() (common.Address, error) {
	return _Bindings.Contract.GetOwner(&_Bindings.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x0ac298dc.
//
// Solidity: function get_owner() view returns(address)
func (_Bindings *BindingsCallerSession) GetOwner() (common.Address, error) {
	return _Bindings.Contract.GetOwner(&_Bindings.CallOpts)
}

// GetRelayByUri is a free data retrieval call binding the contract method 0xf5f33c7b.
//
// Solidity: function get_relay_by_uri(string relay_uri) view returns((string,string,bool,string))
func (_Bindings *BindingsCaller) GetRelayByUri(opts *bind.CallOpts, relay_uri string) (domain.RelayAllowed, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "get_relay_by_uri", relay_uri)

	if err != nil {
		return *new(domain.RelayAllowed), err
	}

	out0 := *abi.ConvertType(out[0], new(domain.RelayAllowed)).(*domain.RelayAllowed)

	return out0, err

}

// GetRelayByUri is a free data retrieval call binding the contract method 0xf5f33c7b.
//
// Solidity: function get_relay_by_uri(string relay_uri) view returns((string,string,bool,string))
func (_Bindings *BindingsSession) GetRelayByUri(relay_uri string) (domain.RelayAllowed, error) {
	return _Bindings.Contract.GetRelayByUri(&_Bindings.CallOpts, relay_uri)
}

// GetRelayByUri is a free data retrieval call binding the contract method 0xf5f33c7b.
//
// Solidity: function get_relay_by_uri(string relay_uri) view returns((string,string,bool,string))
func (_Bindings *BindingsCallerSession) GetRelayByUri(relay_uri string) (domain.RelayAllowed, error) {
	return _Bindings.Contract.GetRelayByUri(&_Bindings.CallOpts, relay_uri)
}

// GetRelays is a free data retrieval call binding the contract method 0x04e469ea.
//
// Solidity: function get_relays() view returns((string,string,bool,string)[])
func (_Bindings *BindingsCaller) GetRelays(opts *bind.CallOpts) ([]domain.RelayAllowed, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "get_relays")

	if err != nil {
		return *new([]domain.RelayAllowed), err
	}

	out0 := *abi.ConvertType(out[0], new([]domain.RelayAllowed)).(*[]domain.RelayAllowed)

	return out0, err

}

// GetRelays is a free data retrieval call binding the contract method 0x04e469ea.
//
// Solidity: function get_relays() view returns((string,string,bool,string)[])
func (_Bindings *BindingsSession) GetRelays() ([]domain.RelayAllowed, error) {
	return _Bindings.Contract.GetRelays(&_Bindings.CallOpts)
}

// GetRelays is a free data retrieval call binding the contract method 0x04e469ea.
//
// Solidity: function get_relays() view returns((string,string,bool,string)[])
func (_Bindings *BindingsCallerSession) GetRelays() ([]domain.RelayAllowed, error) {
	return _Bindings.Contract.GetRelays(&_Bindings.CallOpts)
}

// GetRelaysAmount is a free data retrieval call binding the contract method 0x312c3165.
//
// Solidity: function get_relays_amount() view returns(uint256)
func (_Bindings *BindingsCaller) GetRelaysAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "get_relays_amount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRelaysAmount is a free data retrieval call binding the contract method 0x312c3165.
//
// Solidity: function get_relays_amount() view returns(uint256)
func (_Bindings *BindingsSession) GetRelaysAmount() (*big.Int, error) {
	return _Bindings.Contract.GetRelaysAmount(&_Bindings.CallOpts)
}

// GetRelaysAmount is a free data retrieval call binding the contract method 0x312c3165.
//
// Solidity: function get_relays_amount() view returns(uint256)
func (_Bindings *BindingsCallerSession) GetRelaysAmount() (*big.Int, error) {
	return _Bindings.Contract.GetRelaysAmount(&_Bindings.CallOpts)
}

// AddRelay is a paid mutator transaction binding the contract method 0x2e21ecef.
//
// Solidity: function add_relay(string uri, string operator, bool is_mandatory, string description) returns()
func (_Bindings *BindingsTransactor) AddRelay(opts *bind.TransactOpts, uri string, operator string, is_mandatory bool, description string) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "add_relay", uri, operator, is_mandatory, description)
}

// AddRelay is a paid mutator transaction binding the contract method 0x2e21ecef.
//
// Solidity: function add_relay(string uri, string operator, bool is_mandatory, string description) returns()
func (_Bindings *BindingsSession) AddRelay(uri string, operator string, is_mandatory bool, description string) (*types.Transaction, error) {
	return _Bindings.Contract.AddRelay(&_Bindings.TransactOpts, uri, operator, is_mandatory, description)
}

// AddRelay is a paid mutator transaction binding the contract method 0x2e21ecef.
//
// Solidity: function add_relay(string uri, string operator, bool is_mandatory, string description) returns()
func (_Bindings *BindingsTransactorSession) AddRelay(uri string, operator string, is_mandatory bool, description string) (*types.Transaction, error) {
	return _Bindings.Contract.AddRelay(&_Bindings.TransactOpts, uri, operator, is_mandatory, description)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0x253c8bd4.
//
// Solidity: function change_owner(address owner) returns()
func (_Bindings *BindingsTransactor) ChangeOwner(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "change_owner", owner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0x253c8bd4.
//
// Solidity: function change_owner(address owner) returns()
func (_Bindings *BindingsSession) ChangeOwner(owner common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.ChangeOwner(&_Bindings.TransactOpts, owner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0x253c8bd4.
//
// Solidity: function change_owner(address owner) returns()
func (_Bindings *BindingsTransactorSession) ChangeOwner(owner common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.ChangeOwner(&_Bindings.TransactOpts, owner)
}

// DismissManager is a paid mutator transaction binding the contract method 0x417a02b4.
//
// Solidity: function dismiss_manager() returns()
func (_Bindings *BindingsTransactor) DismissManager(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "dismiss_manager")
}

// DismissManager is a paid mutator transaction binding the contract method 0x417a02b4.
//
// Solidity: function dismiss_manager() returns()
func (_Bindings *BindingsSession) DismissManager() (*types.Transaction, error) {
	return _Bindings.Contract.DismissManager(&_Bindings.TransactOpts)
}

// DismissManager is a paid mutator transaction binding the contract method 0x417a02b4.
//
// Solidity: function dismiss_manager() returns()
func (_Bindings *BindingsTransactorSession) DismissManager() (*types.Transaction, error) {
	return _Bindings.Contract.DismissManager(&_Bindings.TransactOpts)
}

// RecoverErc20 is a paid mutator transaction binding the contract method 0xedd885b4.
//
// Solidity: function recover_erc20(address token, uint256 amount, address recipient) returns()
func (_Bindings *BindingsTransactor) RecoverErc20(opts *bind.TransactOpts, token common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "recover_erc20", token, amount, recipient)
}

// RecoverErc20 is a paid mutator transaction binding the contract method 0xedd885b4.
//
// Solidity: function recover_erc20(address token, uint256 amount, address recipient) returns()
func (_Bindings *BindingsSession) RecoverErc20(token common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.RecoverErc20(&_Bindings.TransactOpts, token, amount, recipient)
}

// RecoverErc20 is a paid mutator transaction binding the contract method 0xedd885b4.
//
// Solidity: function recover_erc20(address token, uint256 amount, address recipient) returns()
func (_Bindings *BindingsTransactorSession) RecoverErc20(token common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.RecoverErc20(&_Bindings.TransactOpts, token, amount, recipient)
}

// RemoveRelay is a paid mutator transaction binding the contract method 0xf5a70a80.
//
// Solidity: function remove_relay(string uri) returns()
func (_Bindings *BindingsTransactor) RemoveRelay(opts *bind.TransactOpts, uri string) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "remove_relay", uri)
}

// RemoveRelay is a paid mutator transaction binding the contract method 0xf5a70a80.
//
// Solidity: function remove_relay(string uri) returns()
func (_Bindings *BindingsSession) RemoveRelay(uri string) (*types.Transaction, error) {
	return _Bindings.Contract.RemoveRelay(&_Bindings.TransactOpts, uri)
}

// RemoveRelay is a paid mutator transaction binding the contract method 0xf5a70a80.
//
// Solidity: function remove_relay(string uri) returns()
func (_Bindings *BindingsTransactorSession) RemoveRelay(uri string) (*types.Transaction, error) {
	return _Bindings.Contract.RemoveRelay(&_Bindings.TransactOpts, uri)
}

// SetManager is a paid mutator transaction binding the contract method 0x9aece83e.
//
// Solidity: function set_manager(address manager) returns()
func (_Bindings *BindingsTransactor) SetManager(opts *bind.TransactOpts, manager common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "set_manager", manager)
}

// SetManager is a paid mutator transaction binding the contract method 0x9aece83e.
//
// Solidity: function set_manager(address manager) returns()
func (_Bindings *BindingsSession) SetManager(manager common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.SetManager(&_Bindings.TransactOpts, manager)
}

// SetManager is a paid mutator transaction binding the contract method 0x9aece83e.
//
// Solidity: function set_manager(address manager) returns()
func (_Bindings *BindingsTransactorSession) SetManager(manager common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.SetManager(&_Bindings.TransactOpts, manager)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Bindings *BindingsTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Bindings.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Bindings *BindingsSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Bindings.Contract.Fallback(&_Bindings.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Bindings *BindingsTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Bindings.Contract.Fallback(&_Bindings.TransactOpts, calldata)
}

// BindingsAllowedListUpdatedIterator is returned from FilterAllowedListUpdated and is used to iterate over the raw logs and unpacked data for AllowedListUpdated events raised by the Bindings contract.
type BindingsAllowedListUpdatedIterator struct {
	Event *BindingsAllowedListUpdated // Event containing the contract specifics and raw log

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
func (it *BindingsAllowedListUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsAllowedListUpdated)
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
		it.Event = new(BindingsAllowedListUpdated)
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
func (it *BindingsAllowedListUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsAllowedListUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsAllowedListUpdated represents a AllowedListUpdated event raised by the Bindings contract.
type BindingsAllowedListUpdated struct {
	AllowedListVersion *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterAllowedListUpdated is a free log retrieval operation binding the contract event 0x49f5627aa055ec3fcd474f99c8b7799b798c04af7b9f215305512c867e5a1839.
//
// Solidity: event AllowedListUpdated(uint256 indexed allowed_list_version)
func (_Bindings *BindingsFilterer) FilterAllowedListUpdated(opts *bind.FilterOpts, allowed_list_version []*big.Int) (*BindingsAllowedListUpdatedIterator, error) {

	var allowed_list_versionRule []interface{}
	for _, allowed_list_versionItem := range allowed_list_version {
		allowed_list_versionRule = append(allowed_list_versionRule, allowed_list_versionItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "AllowedListUpdated", allowed_list_versionRule)
	if err != nil {
		return nil, err
	}
	return &BindingsAllowedListUpdatedIterator{contract: _Bindings.contract, event: "AllowedListUpdated", logs: logs, sub: sub}, nil
}

// WatchAllowedListUpdated is a free log subscription operation binding the contract event 0x49f5627aa055ec3fcd474f99c8b7799b798c04af7b9f215305512c867e5a1839.
//
// Solidity: event AllowedListUpdated(uint256 indexed allowed_list_version)
func (_Bindings *BindingsFilterer) WatchAllowedListUpdated(opts *bind.WatchOpts, sink chan<- *BindingsAllowedListUpdated, allowed_list_version []*big.Int) (event.Subscription, error) {

	var allowed_list_versionRule []interface{}
	for _, allowed_list_versionItem := range allowed_list_version {
		allowed_list_versionRule = append(allowed_list_versionRule, allowed_list_versionItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "AllowedListUpdated", allowed_list_versionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsAllowedListUpdated)
				if err := _Bindings.contract.UnpackLog(event, "AllowedListUpdated", log); err != nil {
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

// ParseAllowedListUpdated is a log parse operation binding the contract event 0x49f5627aa055ec3fcd474f99c8b7799b798c04af7b9f215305512c867e5a1839.
//
// Solidity: event AllowedListUpdated(uint256 indexed allowed_list_version)
func (_Bindings *BindingsFilterer) ParseAllowedListUpdated(log types.Log) (*BindingsAllowedListUpdated, error) {
	event := new(BindingsAllowedListUpdated)
	if err := _Bindings.contract.UnpackLog(event, "AllowedListUpdated", log); err != nil {
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
	Amount    *big.Int
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERC20Recovered is a free log retrieval operation binding the contract event 0x8619312ed4eff1cf9f0116e6db2f49d9570a86f0350d1c5ad1bd0f7b0cf9e132.
//
// Solidity: event ERC20Recovered(address indexed token, uint256 amount, address indexed recipient)
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

// WatchERC20Recovered is a free log subscription operation binding the contract event 0x8619312ed4eff1cf9f0116e6db2f49d9570a86f0350d1c5ad1bd0f7b0cf9e132.
//
// Solidity: event ERC20Recovered(address indexed token, uint256 amount, address indexed recipient)
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

// ParseERC20Recovered is a log parse operation binding the contract event 0x8619312ed4eff1cf9f0116e6db2f49d9570a86f0350d1c5ad1bd0f7b0cf9e132.
//
// Solidity: event ERC20Recovered(address indexed token, uint256 amount, address indexed recipient)
func (_Bindings *BindingsFilterer) ParseERC20Recovered(log types.Log) (*BindingsERC20Recovered, error) {
	event := new(BindingsERC20Recovered)
	if err := _Bindings.contract.UnpackLog(event, "ERC20Recovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsManagerChangedIterator is returned from FilterManagerChanged and is used to iterate over the raw logs and unpacked data for ManagerChanged events raised by the Bindings contract.
type BindingsManagerChangedIterator struct {
	Event *BindingsManagerChanged // Event containing the contract specifics and raw log

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
func (it *BindingsManagerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsManagerChanged)
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
		it.Event = new(BindingsManagerChanged)
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
func (it *BindingsManagerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsManagerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsManagerChanged represents a ManagerChanged event raised by the Bindings contract.
type BindingsManagerChanged struct {
	NewManager common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterManagerChanged is a free log retrieval operation binding the contract event 0x198db6e425fb8aafd1823c6ca50be2d51e5764571a5ae0f0f21c6812e45def0b.
//
// Solidity: event ManagerChanged(address indexed new_manager)
func (_Bindings *BindingsFilterer) FilterManagerChanged(opts *bind.FilterOpts, new_manager []common.Address) (*BindingsManagerChangedIterator, error) {

	var new_managerRule []interface{}
	for _, new_managerItem := range new_manager {
		new_managerRule = append(new_managerRule, new_managerItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "ManagerChanged", new_managerRule)
	if err != nil {
		return nil, err
	}
	return &BindingsManagerChangedIterator{contract: _Bindings.contract, event: "ManagerChanged", logs: logs, sub: sub}, nil
}

// WatchManagerChanged is a free log subscription operation binding the contract event 0x198db6e425fb8aafd1823c6ca50be2d51e5764571a5ae0f0f21c6812e45def0b.
//
// Solidity: event ManagerChanged(address indexed new_manager)
func (_Bindings *BindingsFilterer) WatchManagerChanged(opts *bind.WatchOpts, sink chan<- *BindingsManagerChanged, new_manager []common.Address) (event.Subscription, error) {

	var new_managerRule []interface{}
	for _, new_managerItem := range new_manager {
		new_managerRule = append(new_managerRule, new_managerItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "ManagerChanged", new_managerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsManagerChanged)
				if err := _Bindings.contract.UnpackLog(event, "ManagerChanged", log); err != nil {
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

// ParseManagerChanged is a log parse operation binding the contract event 0x198db6e425fb8aafd1823c6ca50be2d51e5764571a5ae0f0f21c6812e45def0b.
//
// Solidity: event ManagerChanged(address indexed new_manager)
func (_Bindings *BindingsFilterer) ParseManagerChanged(log types.Log) (*BindingsManagerChanged, error) {
	event := new(BindingsManagerChanged)
	if err := _Bindings.contract.UnpackLog(event, "ManagerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the Bindings contract.
type BindingsOwnerChangedIterator struct {
	Event *BindingsOwnerChanged // Event containing the contract specifics and raw log

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
func (it *BindingsOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsOwnerChanged)
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
		it.Event = new(BindingsOwnerChanged)
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
func (it *BindingsOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsOwnerChanged represents a OwnerChanged event raised by the Bindings contract.
type BindingsOwnerChanged struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xa2ea9883a321a3e97b8266c2b078bfeec6d50c711ed71f874a90d500ae2eaf36.
//
// Solidity: event OwnerChanged(address indexed new_owner)
func (_Bindings *BindingsFilterer) FilterOwnerChanged(opts *bind.FilterOpts, new_owner []common.Address) (*BindingsOwnerChangedIterator, error) {

	var new_ownerRule []interface{}
	for _, new_ownerItem := range new_owner {
		new_ownerRule = append(new_ownerRule, new_ownerItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "OwnerChanged", new_ownerRule)
	if err != nil {
		return nil, err
	}
	return &BindingsOwnerChangedIterator{contract: _Bindings.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xa2ea9883a321a3e97b8266c2b078bfeec6d50c711ed71f874a90d500ae2eaf36.
//
// Solidity: event OwnerChanged(address indexed new_owner)
func (_Bindings *BindingsFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *BindingsOwnerChanged, new_owner []common.Address) (event.Subscription, error) {

	var new_ownerRule []interface{}
	for _, new_ownerItem := range new_owner {
		new_ownerRule = append(new_ownerRule, new_ownerItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "OwnerChanged", new_ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsOwnerChanged)
				if err := _Bindings.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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

// ParseOwnerChanged is a log parse operation binding the contract event 0xa2ea9883a321a3e97b8266c2b078bfeec6d50c711ed71f874a90d500ae2eaf36.
//
// Solidity: event OwnerChanged(address indexed new_owner)
func (_Bindings *BindingsFilterer) ParseOwnerChanged(log types.Log) (*BindingsOwnerChanged, error) {
	event := new(BindingsOwnerChanged)
	if err := _Bindings.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsRelayAddedIterator is returned from FilterRelayAdded and is used to iterate over the raw logs and unpacked data for RelayAdded events raised by the Bindings contract.
type BindingsRelayAddedIterator struct {
	Event *BindingsRelayAdded // Event containing the contract specifics and raw log

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
func (it *BindingsRelayAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsRelayAdded)
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
		it.Event = new(BindingsRelayAdded)
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
func (it *BindingsRelayAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsRelayAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsRelayAdded represents a RelayAdded event raised by the Bindings contract.
type BindingsRelayAdded struct {
	UriHash common.Hash
	Relay   domain.RelayAllowed
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRelayAdded is a free log retrieval operation binding the contract event 0xeee5faa84d45af657ab405cdbf2c6a8a6d466e83fa694a358fee5ff84431d0bf.
//
// Solidity: event RelayAdded(string indexed uri_hash, (string,string,bool,string) relay)
func (_Bindings *BindingsFilterer) FilterRelayAdded(opts *bind.FilterOpts, uri_hash []string) (*BindingsRelayAddedIterator, error) {

	var uri_hashRule []interface{}
	for _, uri_hashItem := range uri_hash {
		uri_hashRule = append(uri_hashRule, uri_hashItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "RelayAdded", uri_hashRule)
	if err != nil {
		return nil, err
	}
	return &BindingsRelayAddedIterator{contract: _Bindings.contract, event: "RelayAdded", logs: logs, sub: sub}, nil
}

// WatchRelayAdded is a free log subscription operation binding the contract event 0xeee5faa84d45af657ab405cdbf2c6a8a6d466e83fa694a358fee5ff84431d0bf.
//
// Solidity: event RelayAdded(string indexed uri_hash, (string,string,bool,string) relay)
func (_Bindings *BindingsFilterer) WatchRelayAdded(opts *bind.WatchOpts, sink chan<- *BindingsRelayAdded, uri_hash []string) (event.Subscription, error) {

	var uri_hashRule []interface{}
	for _, uri_hashItem := range uri_hash {
		uri_hashRule = append(uri_hashRule, uri_hashItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "RelayAdded", uri_hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsRelayAdded)
				if err := _Bindings.contract.UnpackLog(event, "RelayAdded", log); err != nil {
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

// ParseRelayAdded is a log parse operation binding the contract event 0xeee5faa84d45af657ab405cdbf2c6a8a6d466e83fa694a358fee5ff84431d0bf.
//
// Solidity: event RelayAdded(string indexed uri_hash, (string,string,bool,string) relay)
func (_Bindings *BindingsFilterer) ParseRelayAdded(log types.Log) (*BindingsRelayAdded, error) {
	event := new(BindingsRelayAdded)
	if err := _Bindings.contract.UnpackLog(event, "RelayAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsRelayRemovedIterator is returned from FilterRelayRemoved and is used to iterate over the raw logs and unpacked data for RelayRemoved events raised by the Bindings contract.
type BindingsRelayRemovedIterator struct {
	Event *BindingsRelayRemoved // Event containing the contract specifics and raw log

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
func (it *BindingsRelayRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsRelayRemoved)
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
		it.Event = new(BindingsRelayRemoved)
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
func (it *BindingsRelayRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsRelayRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsRelayRemoved represents a RelayRemoved event raised by the Bindings contract.
type BindingsRelayRemoved struct {
	UriHash common.Hash
	Uri     string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRelayRemoved is a free log retrieval operation binding the contract event 0xef756634af7ee7210f786ec0f91930afa63fda84d9ff6493ae681c332055dadb.
//
// Solidity: event RelayRemoved(string indexed uri_hash, string uri)
func (_Bindings *BindingsFilterer) FilterRelayRemoved(opts *bind.FilterOpts, uri_hash []string) (*BindingsRelayRemovedIterator, error) {

	var uri_hashRule []interface{}
	for _, uri_hashItem := range uri_hash {
		uri_hashRule = append(uri_hashRule, uri_hashItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "RelayRemoved", uri_hashRule)
	if err != nil {
		return nil, err
	}
	return &BindingsRelayRemovedIterator{contract: _Bindings.contract, event: "RelayRemoved", logs: logs, sub: sub}, nil
}

// WatchRelayRemoved is a free log subscription operation binding the contract event 0xef756634af7ee7210f786ec0f91930afa63fda84d9ff6493ae681c332055dadb.
//
// Solidity: event RelayRemoved(string indexed uri_hash, string uri)
func (_Bindings *BindingsFilterer) WatchRelayRemoved(opts *bind.WatchOpts, sink chan<- *BindingsRelayRemoved, uri_hash []string) (event.Subscription, error) {

	var uri_hashRule []interface{}
	for _, uri_hashItem := range uri_hash {
		uri_hashRule = append(uri_hashRule, uri_hashItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "RelayRemoved", uri_hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsRelayRemoved)
				if err := _Bindings.contract.UnpackLog(event, "RelayRemoved", log); err != nil {
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

// ParseRelayRemoved is a log parse operation binding the contract event 0xef756634af7ee7210f786ec0f91930afa63fda84d9ff6493ae681c332055dadb.
//
// Solidity: event RelayRemoved(string indexed uri_hash, string uri)
func (_Bindings *BindingsFilterer) ParseRelayRemoved(log types.Log) (*BindingsRelayRemoved, error) {
	event := new(BindingsRelayRemoved)
	if err := _Bindings.contract.UnpackLog(event, "RelayRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
