// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vebo

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

// ValidatorsExitBusOracleProcessingState is an auto generated low-level Go binding around an user-defined struct.
type ValidatorsExitBusOracleProcessingState struct {
	CurrentFrameRefSlot    *big.Int
	ProcessingDeadlineTime *big.Int
	DataHash               [32]byte
	DataSubmitted          bool
	DataFormat             *big.Int
	RequestsCount          *big.Int
	RequestsSubmitted      *big.Int
}

// ValidatorsExitBusOracleReportData is an auto generated low-level Go binding around an user-defined struct.
type ValidatorsExitBusOracleReportData struct {
	ConsensusVersion *big.Int
	RefSlot          *big.Int
	RequestsCount    *big.Int
	DataFormat       *big.Int
	Data             []byte
}

// VeboMetaData contains all meta data concerning the Vebo contract.
var VeboMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"secondsPerSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"genesisTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"lidoLocator\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AddressCannotBeSame\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AdminCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ArgumentOutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HashCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialRefSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"processingRefSlot\",\"type\":\"uint256\"}],\"name\":\"InitialRefSlotCannotBeLessThanProcessingOne\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidContractVersionIncrement\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRequestsData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRequestsDataLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRequestsDataSortOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoConsensusReportToProcess\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"moduleId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nodeOpId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevRequestedValidatorIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestedValidatorIndex\",\"type\":\"uint256\"}],\"name\":\"NodeOpValidatorIndexMustIncrease\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonZeroContractVersionOnInit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PauseUntilMustBeInFuture\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PausedExpected\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"ProcessingDeadlineMissed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RefSlotAlreadyProcessing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevRefSlot\",\"type\":\"uint256\"}],\"name\":\"RefSlotCannotDecrease\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"processingRefSlot\",\"type\":\"uint256\"}],\"name\":\"RefSlotMustBeGreaterThanProcessingOne\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ResumedExpected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SecondsPerSlotCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderIsNotTheConsensusContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedChainConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedVersion\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"receivedVersion\",\"type\":\"uint256\"}],\"name\":\"UnexpectedConsensusVersion\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"received\",\"type\":\"uint256\"}],\"name\":\"UnexpectedContractVersion\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"consensusHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"receivedHash\",\"type\":\"bytes32\"}],\"name\":\"UnexpectedDataHash\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"consensusRefSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dataRefSlot\",\"type\":\"uint256\"}],\"name\":\"UnexpectedRefSlot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedRequestsDataLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"format\",\"type\":\"uint256\"}],\"name\":\"UnsupportedRequestsDataFormat\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionCannotBeSame\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroPauseDuration\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prevAddr\",\"type\":\"address\"}],\"name\":\"ConsensusHashContractSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"prevVersion\",\"type\":\"uint256\"}],\"name\":\"ConsensusVersionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"ContractVersionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"ProcessingStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"ReportDiscarded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"processingDeadlineTime\",\"type\":\"uint256\"}],\"name\":\"ReportSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Resumed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakingModuleId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"validatorPubkey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ValidatorExitRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestsProcessed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestsCount\",\"type\":\"uint256\"}],\"name\":\"WarnDataIncompleteProcessing\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"}],\"name\":\"WarnProcessingMissed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DATA_FORMAT_LIST\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GENESIS_TIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGE_CONSENSUS_CONTRACT_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGE_CONSENSUS_VERSION_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSE_INFINITELY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RESUME_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SECONDS_PER_SLOT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SUBMIT_DATA_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"}],\"name\":\"discardConsensusReport\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConsensusContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConsensusReport\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"processingDeadlineTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"processingStarted\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConsensusVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastProcessingRefSlot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"moduleId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"nodeOpIds\",\"type\":\"uint256[]\"}],\"name\":\"getLastRequestedValidatorIndices\",\"outputs\":[{\"internalType\":\"int256[]\",\"name\":\"\",\"type\":\"int256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProcessingState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"currentFrameRefSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"processingDeadlineTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"dataSubmitted\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"dataFormat\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestsSubmitted\",\"type\":\"uint256\"}],\"internalType\":\"structValidatorsExitBusOracle.ProcessingState\",\"name\":\"result\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getResumeSinceTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalRequestsProcessed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"consensusContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"consensusVersion\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastProcessingRefSlot\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"pauseFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pauseUntilInclusive\",\"type\":\"uint256\"}],\"name\":\"pauseUntil\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resume\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setConsensusContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"setConsensusVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"reportHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"submitConsensusReport\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"consensusVersion\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dataFormat\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structValidatorsExitBusOracle.ReportData\",\"name\":\"data\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"contractVersion\",\"type\":\"uint256\"}],\"name\":\"submitReportData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VeboABI is the input ABI used to generate the binding from.
// Deprecated: Use VeboMetaData.ABI instead.
var VeboABI = VeboMetaData.ABI

// Vebo is an auto generated Go binding around an Ethereum contract.
type Vebo struct {
	VeboCaller     // Read-only binding to the contract
	VeboTransactor // Write-only binding to the contract
	VeboFilterer   // Log filterer for contract events
}

// VeboCaller is an auto generated read-only Go binding around an Ethereum contract.
type VeboCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VeboTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VeboTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VeboFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VeboFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VeboSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VeboSession struct {
	Contract     *Vebo             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VeboCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VeboCallerSession struct {
	Contract *VeboCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VeboTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VeboTransactorSession struct {
	Contract     *VeboTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VeboRaw is an auto generated low-level Go binding around an Ethereum contract.
type VeboRaw struct {
	Contract *Vebo // Generic contract binding to access the raw methods on
}

// VeboCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VeboCallerRaw struct {
	Contract *VeboCaller // Generic read-only contract binding to access the raw methods on
}

// VeboTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VeboTransactorRaw struct {
	Contract *VeboTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVebo creates a new instance of Vebo, bound to a specific deployed contract.
func NewVebo(address common.Address, backend bind.ContractBackend) (*Vebo, error) {
	contract, err := bindVebo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vebo{VeboCaller: VeboCaller{contract: contract}, VeboTransactor: VeboTransactor{contract: contract}, VeboFilterer: VeboFilterer{contract: contract}}, nil
}

// NewVeboCaller creates a new read-only instance of Vebo, bound to a specific deployed contract.
func NewVeboCaller(address common.Address, caller bind.ContractCaller) (*VeboCaller, error) {
	contract, err := bindVebo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VeboCaller{contract: contract}, nil
}

// NewVeboTransactor creates a new write-only instance of Vebo, bound to a specific deployed contract.
func NewVeboTransactor(address common.Address, transactor bind.ContractTransactor) (*VeboTransactor, error) {
	contract, err := bindVebo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VeboTransactor{contract: contract}, nil
}

// NewVeboFilterer creates a new log filterer instance of Vebo, bound to a specific deployed contract.
func NewVeboFilterer(address common.Address, filterer bind.ContractFilterer) (*VeboFilterer, error) {
	contract, err := bindVebo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VeboFilterer{contract: contract}, nil
}

// bindVebo binds a generic wrapper to an already deployed contract.
func bindVebo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VeboMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vebo *VeboRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vebo.Contract.VeboCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vebo *VeboRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vebo.Contract.VeboTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vebo *VeboRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vebo.Contract.VeboTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vebo *VeboCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vebo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vebo *VeboTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vebo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vebo *VeboTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vebo.Contract.contract.Transact(opts, method, params...)
}

// DATAFORMATLIST is a free data retrieval call binding the contract method 0xe271b774.
//
// Solidity: function DATA_FORMAT_LIST() view returns(uint256)
func (_Vebo *VeboCaller) DATAFORMATLIST(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "DATA_FORMAT_LIST")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DATAFORMATLIST is a free data retrieval call binding the contract method 0xe271b774.
//
// Solidity: function DATA_FORMAT_LIST() view returns(uint256)
func (_Vebo *VeboSession) DATAFORMATLIST() (*big.Int, error) {
	return _Vebo.Contract.DATAFORMATLIST(&_Vebo.CallOpts)
}

// DATAFORMATLIST is a free data retrieval call binding the contract method 0xe271b774.
//
// Solidity: function DATA_FORMAT_LIST() view returns(uint256)
func (_Vebo *VeboCallerSession) DATAFORMATLIST() (*big.Int, error) {
	return _Vebo.Contract.DATAFORMATLIST(&_Vebo.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Vebo *VeboCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Vebo *VeboSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Vebo.Contract.DEFAULTADMINROLE(&_Vebo.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Vebo *VeboCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Vebo.Contract.DEFAULTADMINROLE(&_Vebo.CallOpts)
}

// GENESISTIME is a free data retrieval call binding the contract method 0xf2882461.
//
// Solidity: function GENESIS_TIME() view returns(uint256)
func (_Vebo *VeboCaller) GENESISTIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "GENESIS_TIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GENESISTIME is a free data retrieval call binding the contract method 0xf2882461.
//
// Solidity: function GENESIS_TIME() view returns(uint256)
func (_Vebo *VeboSession) GENESISTIME() (*big.Int, error) {
	return _Vebo.Contract.GENESISTIME(&_Vebo.CallOpts)
}

// GENESISTIME is a free data retrieval call binding the contract method 0xf2882461.
//
// Solidity: function GENESIS_TIME() view returns(uint256)
func (_Vebo *VeboCallerSession) GENESISTIME() (*big.Int, error) {
	return _Vebo.Contract.GENESISTIME(&_Vebo.CallOpts)
}

// MANAGECONSENSUSCONTRACTROLE is a free data retrieval call binding the contract method 0xad5cac4e.
//
// Solidity: function MANAGE_CONSENSUS_CONTRACT_ROLE() view returns(bytes32)
func (_Vebo *VeboCaller) MANAGECONSENSUSCONTRACTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "MANAGE_CONSENSUS_CONTRACT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGECONSENSUSCONTRACTROLE is a free data retrieval call binding the contract method 0xad5cac4e.
//
// Solidity: function MANAGE_CONSENSUS_CONTRACT_ROLE() view returns(bytes32)
func (_Vebo *VeboSession) MANAGECONSENSUSCONTRACTROLE() ([32]byte, error) {
	return _Vebo.Contract.MANAGECONSENSUSCONTRACTROLE(&_Vebo.CallOpts)
}

// MANAGECONSENSUSCONTRACTROLE is a free data retrieval call binding the contract method 0xad5cac4e.
//
// Solidity: function MANAGE_CONSENSUS_CONTRACT_ROLE() view returns(bytes32)
func (_Vebo *VeboCallerSession) MANAGECONSENSUSCONTRACTROLE() ([32]byte, error) {
	return _Vebo.Contract.MANAGECONSENSUSCONTRACTROLE(&_Vebo.CallOpts)
}

// MANAGECONSENSUSVERSIONROLE is a free data retrieval call binding the contract method 0x9cc23c79.
//
// Solidity: function MANAGE_CONSENSUS_VERSION_ROLE() view returns(bytes32)
func (_Vebo *VeboCaller) MANAGECONSENSUSVERSIONROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "MANAGE_CONSENSUS_VERSION_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGECONSENSUSVERSIONROLE is a free data retrieval call binding the contract method 0x9cc23c79.
//
// Solidity: function MANAGE_CONSENSUS_VERSION_ROLE() view returns(bytes32)
func (_Vebo *VeboSession) MANAGECONSENSUSVERSIONROLE() ([32]byte, error) {
	return _Vebo.Contract.MANAGECONSENSUSVERSIONROLE(&_Vebo.CallOpts)
}

// MANAGECONSENSUSVERSIONROLE is a free data retrieval call binding the contract method 0x9cc23c79.
//
// Solidity: function MANAGE_CONSENSUS_VERSION_ROLE() view returns(bytes32)
func (_Vebo *VeboCallerSession) MANAGECONSENSUSVERSIONROLE() ([32]byte, error) {
	return _Vebo.Contract.MANAGECONSENSUSVERSIONROLE(&_Vebo.CallOpts)
}

// PAUSEINFINITELY is a free data retrieval call binding the contract method 0xa302ee38.
//
// Solidity: function PAUSE_INFINITELY() view returns(uint256)
func (_Vebo *VeboCaller) PAUSEINFINITELY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "PAUSE_INFINITELY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PAUSEINFINITELY is a free data retrieval call binding the contract method 0xa302ee38.
//
// Solidity: function PAUSE_INFINITELY() view returns(uint256)
func (_Vebo *VeboSession) PAUSEINFINITELY() (*big.Int, error) {
	return _Vebo.Contract.PAUSEINFINITELY(&_Vebo.CallOpts)
}

// PAUSEINFINITELY is a free data retrieval call binding the contract method 0xa302ee38.
//
// Solidity: function PAUSE_INFINITELY() view returns(uint256)
func (_Vebo *VeboCallerSession) PAUSEINFINITELY() (*big.Int, error) {
	return _Vebo.Contract.PAUSEINFINITELY(&_Vebo.CallOpts)
}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_Vebo *VeboCaller) PAUSEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "PAUSE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_Vebo *VeboSession) PAUSEROLE() ([32]byte, error) {
	return _Vebo.Contract.PAUSEROLE(&_Vebo.CallOpts)
}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_Vebo *VeboCallerSession) PAUSEROLE() ([32]byte, error) {
	return _Vebo.Contract.PAUSEROLE(&_Vebo.CallOpts)
}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_Vebo *VeboCaller) RESUMEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "RESUME_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_Vebo *VeboSession) RESUMEROLE() ([32]byte, error) {
	return _Vebo.Contract.RESUMEROLE(&_Vebo.CallOpts)
}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_Vebo *VeboCallerSession) RESUMEROLE() ([32]byte, error) {
	return _Vebo.Contract.RESUMEROLE(&_Vebo.CallOpts)
}

// SECONDSPERSLOT is a free data retrieval call binding the contract method 0x304b9071.
//
// Solidity: function SECONDS_PER_SLOT() view returns(uint256)
func (_Vebo *VeboCaller) SECONDSPERSLOT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "SECONDS_PER_SLOT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SECONDSPERSLOT is a free data retrieval call binding the contract method 0x304b9071.
//
// Solidity: function SECONDS_PER_SLOT() view returns(uint256)
func (_Vebo *VeboSession) SECONDSPERSLOT() (*big.Int, error) {
	return _Vebo.Contract.SECONDSPERSLOT(&_Vebo.CallOpts)
}

// SECONDSPERSLOT is a free data retrieval call binding the contract method 0x304b9071.
//
// Solidity: function SECONDS_PER_SLOT() view returns(uint256)
func (_Vebo *VeboCallerSession) SECONDSPERSLOT() (*big.Int, error) {
	return _Vebo.Contract.SECONDSPERSLOT(&_Vebo.CallOpts)
}

// SUBMITDATAROLE is a free data retrieval call binding the contract method 0x46e1f576.
//
// Solidity: function SUBMIT_DATA_ROLE() view returns(bytes32)
func (_Vebo *VeboCaller) SUBMITDATAROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "SUBMIT_DATA_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SUBMITDATAROLE is a free data retrieval call binding the contract method 0x46e1f576.
//
// Solidity: function SUBMIT_DATA_ROLE() view returns(bytes32)
func (_Vebo *VeboSession) SUBMITDATAROLE() ([32]byte, error) {
	return _Vebo.Contract.SUBMITDATAROLE(&_Vebo.CallOpts)
}

// SUBMITDATAROLE is a free data retrieval call binding the contract method 0x46e1f576.
//
// Solidity: function SUBMIT_DATA_ROLE() view returns(bytes32)
func (_Vebo *VeboCallerSession) SUBMITDATAROLE() ([32]byte, error) {
	return _Vebo.Contract.SUBMITDATAROLE(&_Vebo.CallOpts)
}

// GetConsensusContract is a free data retrieval call binding the contract method 0x8f55b571.
//
// Solidity: function getConsensusContract() view returns(address)
func (_Vebo *VeboCaller) GetConsensusContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "getConsensusContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetConsensusContract is a free data retrieval call binding the contract method 0x8f55b571.
//
// Solidity: function getConsensusContract() view returns(address)
func (_Vebo *VeboSession) GetConsensusContract() (common.Address, error) {
	return _Vebo.Contract.GetConsensusContract(&_Vebo.CallOpts)
}

// GetConsensusContract is a free data retrieval call binding the contract method 0x8f55b571.
//
// Solidity: function getConsensusContract() view returns(address)
func (_Vebo *VeboCallerSession) GetConsensusContract() (common.Address, error) {
	return _Vebo.Contract.GetConsensusContract(&_Vebo.CallOpts)
}

// GetConsensusReport is a free data retrieval call binding the contract method 0x60d64d38.
//
// Solidity: function getConsensusReport() view returns(bytes32 hash, uint256 refSlot, uint256 processingDeadlineTime, bool processingStarted)
func (_Vebo *VeboCaller) GetConsensusReport(opts *bind.CallOpts) (struct {
	Hash                   [32]byte
	RefSlot                *big.Int
	ProcessingDeadlineTime *big.Int
	ProcessingStarted      bool
}, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "getConsensusReport")

	outstruct := new(struct {
		Hash                   [32]byte
		RefSlot                *big.Int
		ProcessingDeadlineTime *big.Int
		ProcessingStarted      bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Hash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.RefSlot = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ProcessingDeadlineTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ProcessingStarted = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// GetConsensusReport is a free data retrieval call binding the contract method 0x60d64d38.
//
// Solidity: function getConsensusReport() view returns(bytes32 hash, uint256 refSlot, uint256 processingDeadlineTime, bool processingStarted)
func (_Vebo *VeboSession) GetConsensusReport() (struct {
	Hash                   [32]byte
	RefSlot                *big.Int
	ProcessingDeadlineTime *big.Int
	ProcessingStarted      bool
}, error) {
	return _Vebo.Contract.GetConsensusReport(&_Vebo.CallOpts)
}

// GetConsensusReport is a free data retrieval call binding the contract method 0x60d64d38.
//
// Solidity: function getConsensusReport() view returns(bytes32 hash, uint256 refSlot, uint256 processingDeadlineTime, bool processingStarted)
func (_Vebo *VeboCallerSession) GetConsensusReport() (struct {
	Hash                   [32]byte
	RefSlot                *big.Int
	ProcessingDeadlineTime *big.Int
	ProcessingStarted      bool
}, error) {
	return _Vebo.Contract.GetConsensusReport(&_Vebo.CallOpts)
}

// GetConsensusVersion is a free data retrieval call binding the contract method 0x5be20425.
//
// Solidity: function getConsensusVersion() view returns(uint256)
func (_Vebo *VeboCaller) GetConsensusVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "getConsensusVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetConsensusVersion is a free data retrieval call binding the contract method 0x5be20425.
//
// Solidity: function getConsensusVersion() view returns(uint256)
func (_Vebo *VeboSession) GetConsensusVersion() (*big.Int, error) {
	return _Vebo.Contract.GetConsensusVersion(&_Vebo.CallOpts)
}

// GetConsensusVersion is a free data retrieval call binding the contract method 0x5be20425.
//
// Solidity: function getConsensusVersion() view returns(uint256)
func (_Vebo *VeboCallerSession) GetConsensusVersion() (*big.Int, error) {
	return _Vebo.Contract.GetConsensusVersion(&_Vebo.CallOpts)
}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_Vebo *VeboCaller) GetContractVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "getContractVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_Vebo *VeboSession) GetContractVersion() (*big.Int, error) {
	return _Vebo.Contract.GetContractVersion(&_Vebo.CallOpts)
}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_Vebo *VeboCallerSession) GetContractVersion() (*big.Int, error) {
	return _Vebo.Contract.GetContractVersion(&_Vebo.CallOpts)
}

// GetLastProcessingRefSlot is a free data retrieval call binding the contract method 0x3584d59c.
//
// Solidity: function getLastProcessingRefSlot() view returns(uint256)
func (_Vebo *VeboCaller) GetLastProcessingRefSlot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "getLastProcessingRefSlot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastProcessingRefSlot is a free data retrieval call binding the contract method 0x3584d59c.
//
// Solidity: function getLastProcessingRefSlot() view returns(uint256)
func (_Vebo *VeboSession) GetLastProcessingRefSlot() (*big.Int, error) {
	return _Vebo.Contract.GetLastProcessingRefSlot(&_Vebo.CallOpts)
}

// GetLastProcessingRefSlot is a free data retrieval call binding the contract method 0x3584d59c.
//
// Solidity: function getLastProcessingRefSlot() view returns(uint256)
func (_Vebo *VeboCallerSession) GetLastProcessingRefSlot() (*big.Int, error) {
	return _Vebo.Contract.GetLastProcessingRefSlot(&_Vebo.CallOpts)
}

// GetLastRequestedValidatorIndices is a free data retrieval call binding the contract method 0xef9bf37e.
//
// Solidity: function getLastRequestedValidatorIndices(uint256 moduleId, uint256[] nodeOpIds) view returns(int256[])
func (_Vebo *VeboCaller) GetLastRequestedValidatorIndices(opts *bind.CallOpts, moduleId *big.Int, nodeOpIds []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "getLastRequestedValidatorIndices", moduleId, nodeOpIds)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetLastRequestedValidatorIndices is a free data retrieval call binding the contract method 0xef9bf37e.
//
// Solidity: function getLastRequestedValidatorIndices(uint256 moduleId, uint256[] nodeOpIds) view returns(int256[])
func (_Vebo *VeboSession) GetLastRequestedValidatorIndices(moduleId *big.Int, nodeOpIds []*big.Int) ([]*big.Int, error) {
	return _Vebo.Contract.GetLastRequestedValidatorIndices(&_Vebo.CallOpts, moduleId, nodeOpIds)
}

// GetLastRequestedValidatorIndices is a free data retrieval call binding the contract method 0xef9bf37e.
//
// Solidity: function getLastRequestedValidatorIndices(uint256 moduleId, uint256[] nodeOpIds) view returns(int256[])
func (_Vebo *VeboCallerSession) GetLastRequestedValidatorIndices(moduleId *big.Int, nodeOpIds []*big.Int) ([]*big.Int, error) {
	return _Vebo.Contract.GetLastRequestedValidatorIndices(&_Vebo.CallOpts, moduleId, nodeOpIds)
}

// GetProcessingState is a free data retrieval call binding the contract method 0x8f7797c2.
//
// Solidity: function getProcessingState() view returns((uint256,uint256,bytes32,bool,uint256,uint256,uint256) result)
func (_Vebo *VeboCaller) GetProcessingState(opts *bind.CallOpts) (ValidatorsExitBusOracleProcessingState, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "getProcessingState")

	if err != nil {
		return *new(ValidatorsExitBusOracleProcessingState), err
	}

	out0 := *abi.ConvertType(out[0], new(ValidatorsExitBusOracleProcessingState)).(*ValidatorsExitBusOracleProcessingState)

	return out0, err

}

// GetProcessingState is a free data retrieval call binding the contract method 0x8f7797c2.
//
// Solidity: function getProcessingState() view returns((uint256,uint256,bytes32,bool,uint256,uint256,uint256) result)
func (_Vebo *VeboSession) GetProcessingState() (ValidatorsExitBusOracleProcessingState, error) {
	return _Vebo.Contract.GetProcessingState(&_Vebo.CallOpts)
}

// GetProcessingState is a free data retrieval call binding the contract method 0x8f7797c2.
//
// Solidity: function getProcessingState() view returns((uint256,uint256,bytes32,bool,uint256,uint256,uint256) result)
func (_Vebo *VeboCallerSession) GetProcessingState() (ValidatorsExitBusOracleProcessingState, error) {
	return _Vebo.Contract.GetProcessingState(&_Vebo.CallOpts)
}

// GetResumeSinceTimestamp is a free data retrieval call binding the contract method 0x589ff76c.
//
// Solidity: function getResumeSinceTimestamp() view returns(uint256)
func (_Vebo *VeboCaller) GetResumeSinceTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "getResumeSinceTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetResumeSinceTimestamp is a free data retrieval call binding the contract method 0x589ff76c.
//
// Solidity: function getResumeSinceTimestamp() view returns(uint256)
func (_Vebo *VeboSession) GetResumeSinceTimestamp() (*big.Int, error) {
	return _Vebo.Contract.GetResumeSinceTimestamp(&_Vebo.CallOpts)
}

// GetResumeSinceTimestamp is a free data retrieval call binding the contract method 0x589ff76c.
//
// Solidity: function getResumeSinceTimestamp() view returns(uint256)
func (_Vebo *VeboCallerSession) GetResumeSinceTimestamp() (*big.Int, error) {
	return _Vebo.Contract.GetResumeSinceTimestamp(&_Vebo.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Vebo *VeboCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Vebo *VeboSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Vebo.Contract.GetRoleAdmin(&_Vebo.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Vebo *VeboCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Vebo.Contract.GetRoleAdmin(&_Vebo.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Vebo *VeboCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Vebo *VeboSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Vebo.Contract.GetRoleMember(&_Vebo.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Vebo *VeboCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Vebo.Contract.GetRoleMember(&_Vebo.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Vebo *VeboCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Vebo *VeboSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Vebo.Contract.GetRoleMemberCount(&_Vebo.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Vebo *VeboCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Vebo.Contract.GetRoleMemberCount(&_Vebo.CallOpts, role)
}

// GetTotalRequestsProcessed is a free data retrieval call binding the contract method 0xe2793e72.
//
// Solidity: function getTotalRequestsProcessed() view returns(uint256)
func (_Vebo *VeboCaller) GetTotalRequestsProcessed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "getTotalRequestsProcessed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalRequestsProcessed is a free data retrieval call binding the contract method 0xe2793e72.
//
// Solidity: function getTotalRequestsProcessed() view returns(uint256)
func (_Vebo *VeboSession) GetTotalRequestsProcessed() (*big.Int, error) {
	return _Vebo.Contract.GetTotalRequestsProcessed(&_Vebo.CallOpts)
}

// GetTotalRequestsProcessed is a free data retrieval call binding the contract method 0xe2793e72.
//
// Solidity: function getTotalRequestsProcessed() view returns(uint256)
func (_Vebo *VeboCallerSession) GetTotalRequestsProcessed() (*big.Int, error) {
	return _Vebo.Contract.GetTotalRequestsProcessed(&_Vebo.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Vebo *VeboCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Vebo *VeboSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Vebo.Contract.HasRole(&_Vebo.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Vebo *VeboCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Vebo.Contract.HasRole(&_Vebo.CallOpts, role, account)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Vebo *VeboCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Vebo *VeboSession) IsPaused() (bool, error) {
	return _Vebo.Contract.IsPaused(&_Vebo.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Vebo *VeboCallerSession) IsPaused() (bool, error) {
	return _Vebo.Contract.IsPaused(&_Vebo.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Vebo *VeboCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Vebo *VeboSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Vebo.Contract.SupportsInterface(&_Vebo.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Vebo *VeboCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Vebo.Contract.SupportsInterface(&_Vebo.CallOpts, interfaceId)
}

// DiscardConsensusReport is a paid mutator transaction binding the contract method 0xd4381217.
//
// Solidity: function discardConsensusReport(uint256 refSlot) returns()
func (_Vebo *VeboTransactor) DiscardConsensusReport(opts *bind.TransactOpts, refSlot *big.Int) (*types.Transaction, error) {
	return _Vebo.contract.Transact(opts, "discardConsensusReport", refSlot)
}

// DiscardConsensusReport is a paid mutator transaction binding the contract method 0xd4381217.
//
// Solidity: function discardConsensusReport(uint256 refSlot) returns()
func (_Vebo *VeboSession) DiscardConsensusReport(refSlot *big.Int) (*types.Transaction, error) {
	return _Vebo.Contract.DiscardConsensusReport(&_Vebo.TransactOpts, refSlot)
}

// DiscardConsensusReport is a paid mutator transaction binding the contract method 0xd4381217.
//
// Solidity: function discardConsensusReport(uint256 refSlot) returns()
func (_Vebo *VeboTransactorSession) DiscardConsensusReport(refSlot *big.Int) (*types.Transaction, error) {
	return _Vebo.Contract.DiscardConsensusReport(&_Vebo.TransactOpts, refSlot)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Vebo *VeboTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vebo.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Vebo *VeboSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vebo.Contract.GrantRole(&_Vebo.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Vebo *VeboTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vebo.Contract.GrantRole(&_Vebo.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xeb990c59.
//
// Solidity: function initialize(address admin, address consensusContract, uint256 consensusVersion, uint256 lastProcessingRefSlot) returns()
func (_Vebo *VeboTransactor) Initialize(opts *bind.TransactOpts, admin common.Address, consensusContract common.Address, consensusVersion *big.Int, lastProcessingRefSlot *big.Int) (*types.Transaction, error) {
	return _Vebo.contract.Transact(opts, "initialize", admin, consensusContract, consensusVersion, lastProcessingRefSlot)
}

// Initialize is a paid mutator transaction binding the contract method 0xeb990c59.
//
// Solidity: function initialize(address admin, address consensusContract, uint256 consensusVersion, uint256 lastProcessingRefSlot) returns()
func (_Vebo *VeboSession) Initialize(admin common.Address, consensusContract common.Address, consensusVersion *big.Int, lastProcessingRefSlot *big.Int) (*types.Transaction, error) {
	return _Vebo.Contract.Initialize(&_Vebo.TransactOpts, admin, consensusContract, consensusVersion, lastProcessingRefSlot)
}

// Initialize is a paid mutator transaction binding the contract method 0xeb990c59.
//
// Solidity: function initialize(address admin, address consensusContract, uint256 consensusVersion, uint256 lastProcessingRefSlot) returns()
func (_Vebo *VeboTransactorSession) Initialize(admin common.Address, consensusContract common.Address, consensusVersion *big.Int, lastProcessingRefSlot *big.Int) (*types.Transaction, error) {
	return _Vebo.Contract.Initialize(&_Vebo.TransactOpts, admin, consensusContract, consensusVersion, lastProcessingRefSlot)
}

// PauseFor is a paid mutator transaction binding the contract method 0xf3f449c7.
//
// Solidity: function pauseFor(uint256 _duration) returns()
func (_Vebo *VeboTransactor) PauseFor(opts *bind.TransactOpts, _duration *big.Int) (*types.Transaction, error) {
	return _Vebo.contract.Transact(opts, "pauseFor", _duration)
}

// PauseFor is a paid mutator transaction binding the contract method 0xf3f449c7.
//
// Solidity: function pauseFor(uint256 _duration) returns()
func (_Vebo *VeboSession) PauseFor(_duration *big.Int) (*types.Transaction, error) {
	return _Vebo.Contract.PauseFor(&_Vebo.TransactOpts, _duration)
}

// PauseFor is a paid mutator transaction binding the contract method 0xf3f449c7.
//
// Solidity: function pauseFor(uint256 _duration) returns()
func (_Vebo *VeboTransactorSession) PauseFor(_duration *big.Int) (*types.Transaction, error) {
	return _Vebo.Contract.PauseFor(&_Vebo.TransactOpts, _duration)
}

// PauseUntil is a paid mutator transaction binding the contract method 0xabe9cfc8.
//
// Solidity: function pauseUntil(uint256 _pauseUntilInclusive) returns()
func (_Vebo *VeboTransactor) PauseUntil(opts *bind.TransactOpts, _pauseUntilInclusive *big.Int) (*types.Transaction, error) {
	return _Vebo.contract.Transact(opts, "pauseUntil", _pauseUntilInclusive)
}

// PauseUntil is a paid mutator transaction binding the contract method 0xabe9cfc8.
//
// Solidity: function pauseUntil(uint256 _pauseUntilInclusive) returns()
func (_Vebo *VeboSession) PauseUntil(_pauseUntilInclusive *big.Int) (*types.Transaction, error) {
	return _Vebo.Contract.PauseUntil(&_Vebo.TransactOpts, _pauseUntilInclusive)
}

// PauseUntil is a paid mutator transaction binding the contract method 0xabe9cfc8.
//
// Solidity: function pauseUntil(uint256 _pauseUntilInclusive) returns()
func (_Vebo *VeboTransactorSession) PauseUntil(_pauseUntilInclusive *big.Int) (*types.Transaction, error) {
	return _Vebo.Contract.PauseUntil(&_Vebo.TransactOpts, _pauseUntilInclusive)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Vebo *VeboTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vebo.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Vebo *VeboSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vebo.Contract.RenounceRole(&_Vebo.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Vebo *VeboTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vebo.Contract.RenounceRole(&_Vebo.TransactOpts, role, account)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Vebo *VeboTransactor) Resume(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vebo.contract.Transact(opts, "resume")
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Vebo *VeboSession) Resume() (*types.Transaction, error) {
	return _Vebo.Contract.Resume(&_Vebo.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Vebo *VeboTransactorSession) Resume() (*types.Transaction, error) {
	return _Vebo.Contract.Resume(&_Vebo.TransactOpts)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Vebo *VeboTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vebo.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Vebo *VeboSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vebo.Contract.RevokeRole(&_Vebo.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Vebo *VeboTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vebo.Contract.RevokeRole(&_Vebo.TransactOpts, role, account)
}

// SetConsensusContract is a paid mutator transaction binding the contract method 0xc469c307.
//
// Solidity: function setConsensusContract(address addr) returns()
func (_Vebo *VeboTransactor) SetConsensusContract(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Vebo.contract.Transact(opts, "setConsensusContract", addr)
}

// SetConsensusContract is a paid mutator transaction binding the contract method 0xc469c307.
//
// Solidity: function setConsensusContract(address addr) returns()
func (_Vebo *VeboSession) SetConsensusContract(addr common.Address) (*types.Transaction, error) {
	return _Vebo.Contract.SetConsensusContract(&_Vebo.TransactOpts, addr)
}

// SetConsensusContract is a paid mutator transaction binding the contract method 0xc469c307.
//
// Solidity: function setConsensusContract(address addr) returns()
func (_Vebo *VeboTransactorSession) SetConsensusContract(addr common.Address) (*types.Transaction, error) {
	return _Vebo.Contract.SetConsensusContract(&_Vebo.TransactOpts, addr)
}

// SetConsensusVersion is a paid mutator transaction binding the contract method 0x8d591474.
//
// Solidity: function setConsensusVersion(uint256 version) returns()
func (_Vebo *VeboTransactor) SetConsensusVersion(opts *bind.TransactOpts, version *big.Int) (*types.Transaction, error) {
	return _Vebo.contract.Transact(opts, "setConsensusVersion", version)
}

// SetConsensusVersion is a paid mutator transaction binding the contract method 0x8d591474.
//
// Solidity: function setConsensusVersion(uint256 version) returns()
func (_Vebo *VeboSession) SetConsensusVersion(version *big.Int) (*types.Transaction, error) {
	return _Vebo.Contract.SetConsensusVersion(&_Vebo.TransactOpts, version)
}

// SetConsensusVersion is a paid mutator transaction binding the contract method 0x8d591474.
//
// Solidity: function setConsensusVersion(uint256 version) returns()
func (_Vebo *VeboTransactorSession) SetConsensusVersion(version *big.Int) (*types.Transaction, error) {
	return _Vebo.Contract.SetConsensusVersion(&_Vebo.TransactOpts, version)
}

// SubmitConsensusReport is a paid mutator transaction binding the contract method 0x063f36ad.
//
// Solidity: function submitConsensusReport(bytes32 reportHash, uint256 refSlot, uint256 deadline) returns()
func (_Vebo *VeboTransactor) SubmitConsensusReport(opts *bind.TransactOpts, reportHash [32]byte, refSlot *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Vebo.contract.Transact(opts, "submitConsensusReport", reportHash, refSlot, deadline)
}

// SubmitConsensusReport is a paid mutator transaction binding the contract method 0x063f36ad.
//
// Solidity: function submitConsensusReport(bytes32 reportHash, uint256 refSlot, uint256 deadline) returns()
func (_Vebo *VeboSession) SubmitConsensusReport(reportHash [32]byte, refSlot *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Vebo.Contract.SubmitConsensusReport(&_Vebo.TransactOpts, reportHash, refSlot, deadline)
}

// SubmitConsensusReport is a paid mutator transaction binding the contract method 0x063f36ad.
//
// Solidity: function submitConsensusReport(bytes32 reportHash, uint256 refSlot, uint256 deadline) returns()
func (_Vebo *VeboTransactorSession) SubmitConsensusReport(reportHash [32]byte, refSlot *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Vebo.Contract.SubmitConsensusReport(&_Vebo.TransactOpts, reportHash, refSlot, deadline)
}

// SubmitReportData is a paid mutator transaction binding the contract method 0x294492c8.
//
// Solidity: function submitReportData((uint256,uint256,uint256,uint256,bytes) data, uint256 contractVersion) returns()
func (_Vebo *VeboTransactor) SubmitReportData(opts *bind.TransactOpts, data ValidatorsExitBusOracleReportData, contractVersion *big.Int) (*types.Transaction, error) {
	return _Vebo.contract.Transact(opts, "submitReportData", data, contractVersion)
}

// SubmitReportData is a paid mutator transaction binding the contract method 0x294492c8.
//
// Solidity: function submitReportData((uint256,uint256,uint256,uint256,bytes) data, uint256 contractVersion) returns()
func (_Vebo *VeboSession) SubmitReportData(data ValidatorsExitBusOracleReportData, contractVersion *big.Int) (*types.Transaction, error) {
	return _Vebo.Contract.SubmitReportData(&_Vebo.TransactOpts, data, contractVersion)
}

// SubmitReportData is a paid mutator transaction binding the contract method 0x294492c8.
//
// Solidity: function submitReportData((uint256,uint256,uint256,uint256,bytes) data, uint256 contractVersion) returns()
func (_Vebo *VeboTransactorSession) SubmitReportData(data ValidatorsExitBusOracleReportData, contractVersion *big.Int) (*types.Transaction, error) {
	return _Vebo.Contract.SubmitReportData(&_Vebo.TransactOpts, data, contractVersion)
}

// VeboConsensusHashContractSetIterator is returned from FilterConsensusHashContractSet and is used to iterate over the raw logs and unpacked data for ConsensusHashContractSet events raised by the Vebo contract.
type VeboConsensusHashContractSetIterator struct {
	Event *VeboConsensusHashContractSet // Event containing the contract specifics and raw log

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
func (it *VeboConsensusHashContractSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeboConsensusHashContractSet)
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
		it.Event = new(VeboConsensusHashContractSet)
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
func (it *VeboConsensusHashContractSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeboConsensusHashContractSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeboConsensusHashContractSet represents a ConsensusHashContractSet event raised by the Vebo contract.
type VeboConsensusHashContractSet struct {
	Addr     common.Address
	PrevAddr common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterConsensusHashContractSet is a free log retrieval operation binding the contract event 0x25421480fb7f52d18947876279a213696b58d7e0e5416ce5e2c9f9942661c34c.
//
// Solidity: event ConsensusHashContractSet(address indexed addr, address indexed prevAddr)
func (_Vebo *VeboFilterer) FilterConsensusHashContractSet(opts *bind.FilterOpts, addr []common.Address, prevAddr []common.Address) (*VeboConsensusHashContractSetIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var prevAddrRule []interface{}
	for _, prevAddrItem := range prevAddr {
		prevAddrRule = append(prevAddrRule, prevAddrItem)
	}

	logs, sub, err := _Vebo.contract.FilterLogs(opts, "ConsensusHashContractSet", addrRule, prevAddrRule)
	if err != nil {
		return nil, err
	}
	return &VeboConsensusHashContractSetIterator{contract: _Vebo.contract, event: "ConsensusHashContractSet", logs: logs, sub: sub}, nil
}

// WatchConsensusHashContractSet is a free log subscription operation binding the contract event 0x25421480fb7f52d18947876279a213696b58d7e0e5416ce5e2c9f9942661c34c.
//
// Solidity: event ConsensusHashContractSet(address indexed addr, address indexed prevAddr)
func (_Vebo *VeboFilterer) WatchConsensusHashContractSet(opts *bind.WatchOpts, sink chan<- *VeboConsensusHashContractSet, addr []common.Address, prevAddr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var prevAddrRule []interface{}
	for _, prevAddrItem := range prevAddr {
		prevAddrRule = append(prevAddrRule, prevAddrItem)
	}

	logs, sub, err := _Vebo.contract.WatchLogs(opts, "ConsensusHashContractSet", addrRule, prevAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeboConsensusHashContractSet)
				if err := _Vebo.contract.UnpackLog(event, "ConsensusHashContractSet", log); err != nil {
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

// ParseConsensusHashContractSet is a log parse operation binding the contract event 0x25421480fb7f52d18947876279a213696b58d7e0e5416ce5e2c9f9942661c34c.
//
// Solidity: event ConsensusHashContractSet(address indexed addr, address indexed prevAddr)
func (_Vebo *VeboFilterer) ParseConsensusHashContractSet(log types.Log) (*VeboConsensusHashContractSet, error) {
	event := new(VeboConsensusHashContractSet)
	if err := _Vebo.contract.UnpackLog(event, "ConsensusHashContractSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeboConsensusVersionSetIterator is returned from FilterConsensusVersionSet and is used to iterate over the raw logs and unpacked data for ConsensusVersionSet events raised by the Vebo contract.
type VeboConsensusVersionSetIterator struct {
	Event *VeboConsensusVersionSet // Event containing the contract specifics and raw log

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
func (it *VeboConsensusVersionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeboConsensusVersionSet)
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
		it.Event = new(VeboConsensusVersionSet)
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
func (it *VeboConsensusVersionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeboConsensusVersionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeboConsensusVersionSet represents a ConsensusVersionSet event raised by the Vebo contract.
type VeboConsensusVersionSet struct {
	Version     *big.Int
	PrevVersion *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConsensusVersionSet is a free log retrieval operation binding the contract event 0xfa5304972d4ec3e3207f0bbf91155a49d0dfa62488f9529403a2a49e4b29a895.
//
// Solidity: event ConsensusVersionSet(uint256 indexed version, uint256 indexed prevVersion)
func (_Vebo *VeboFilterer) FilterConsensusVersionSet(opts *bind.FilterOpts, version []*big.Int, prevVersion []*big.Int) (*VeboConsensusVersionSetIterator, error) {

	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}
	var prevVersionRule []interface{}
	for _, prevVersionItem := range prevVersion {
		prevVersionRule = append(prevVersionRule, prevVersionItem)
	}

	logs, sub, err := _Vebo.contract.FilterLogs(opts, "ConsensusVersionSet", versionRule, prevVersionRule)
	if err != nil {
		return nil, err
	}
	return &VeboConsensusVersionSetIterator{contract: _Vebo.contract, event: "ConsensusVersionSet", logs: logs, sub: sub}, nil
}

// WatchConsensusVersionSet is a free log subscription operation binding the contract event 0xfa5304972d4ec3e3207f0bbf91155a49d0dfa62488f9529403a2a49e4b29a895.
//
// Solidity: event ConsensusVersionSet(uint256 indexed version, uint256 indexed prevVersion)
func (_Vebo *VeboFilterer) WatchConsensusVersionSet(opts *bind.WatchOpts, sink chan<- *VeboConsensusVersionSet, version []*big.Int, prevVersion []*big.Int) (event.Subscription, error) {

	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}
	var prevVersionRule []interface{}
	for _, prevVersionItem := range prevVersion {
		prevVersionRule = append(prevVersionRule, prevVersionItem)
	}

	logs, sub, err := _Vebo.contract.WatchLogs(opts, "ConsensusVersionSet", versionRule, prevVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeboConsensusVersionSet)
				if err := _Vebo.contract.UnpackLog(event, "ConsensusVersionSet", log); err != nil {
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

// ParseConsensusVersionSet is a log parse operation binding the contract event 0xfa5304972d4ec3e3207f0bbf91155a49d0dfa62488f9529403a2a49e4b29a895.
//
// Solidity: event ConsensusVersionSet(uint256 indexed version, uint256 indexed prevVersion)
func (_Vebo *VeboFilterer) ParseConsensusVersionSet(log types.Log) (*VeboConsensusVersionSet, error) {
	event := new(VeboConsensusVersionSet)
	if err := _Vebo.contract.UnpackLog(event, "ConsensusVersionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeboContractVersionSetIterator is returned from FilterContractVersionSet and is used to iterate over the raw logs and unpacked data for ContractVersionSet events raised by the Vebo contract.
type VeboContractVersionSetIterator struct {
	Event *VeboContractVersionSet // Event containing the contract specifics and raw log

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
func (it *VeboContractVersionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeboContractVersionSet)
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
		it.Event = new(VeboContractVersionSet)
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
func (it *VeboContractVersionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeboContractVersionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeboContractVersionSet represents a ContractVersionSet event raised by the Vebo contract.
type VeboContractVersionSet struct {
	Version *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterContractVersionSet is a free log retrieval operation binding the contract event 0xfddcded6b4f4730c226821172046b48372d3cd963c159701ae1b7c3bcac541bb.
//
// Solidity: event ContractVersionSet(uint256 version)
func (_Vebo *VeboFilterer) FilterContractVersionSet(opts *bind.FilterOpts) (*VeboContractVersionSetIterator, error) {

	logs, sub, err := _Vebo.contract.FilterLogs(opts, "ContractVersionSet")
	if err != nil {
		return nil, err
	}
	return &VeboContractVersionSetIterator{contract: _Vebo.contract, event: "ContractVersionSet", logs: logs, sub: sub}, nil
}

// WatchContractVersionSet is a free log subscription operation binding the contract event 0xfddcded6b4f4730c226821172046b48372d3cd963c159701ae1b7c3bcac541bb.
//
// Solidity: event ContractVersionSet(uint256 version)
func (_Vebo *VeboFilterer) WatchContractVersionSet(opts *bind.WatchOpts, sink chan<- *VeboContractVersionSet) (event.Subscription, error) {

	logs, sub, err := _Vebo.contract.WatchLogs(opts, "ContractVersionSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeboContractVersionSet)
				if err := _Vebo.contract.UnpackLog(event, "ContractVersionSet", log); err != nil {
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

// ParseContractVersionSet is a log parse operation binding the contract event 0xfddcded6b4f4730c226821172046b48372d3cd963c159701ae1b7c3bcac541bb.
//
// Solidity: event ContractVersionSet(uint256 version)
func (_Vebo *VeboFilterer) ParseContractVersionSet(log types.Log) (*VeboContractVersionSet, error) {
	event := new(VeboContractVersionSet)
	if err := _Vebo.contract.UnpackLog(event, "ContractVersionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeboPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Vebo contract.
type VeboPausedIterator struct {
	Event *VeboPaused // Event containing the contract specifics and raw log

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
func (it *VeboPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeboPaused)
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
		it.Event = new(VeboPaused)
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
func (it *VeboPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeboPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeboPaused represents a Paused event raised by the Vebo contract.
type VeboPaused struct {
	Duration *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e.
//
// Solidity: event Paused(uint256 duration)
func (_Vebo *VeboFilterer) FilterPaused(opts *bind.FilterOpts) (*VeboPausedIterator, error) {

	logs, sub, err := _Vebo.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &VeboPausedIterator{contract: _Vebo.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e.
//
// Solidity: event Paused(uint256 duration)
func (_Vebo *VeboFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *VeboPaused) (event.Subscription, error) {

	logs, sub, err := _Vebo.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeboPaused)
				if err := _Vebo.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e.
//
// Solidity: event Paused(uint256 duration)
func (_Vebo *VeboFilterer) ParsePaused(log types.Log) (*VeboPaused, error) {
	event := new(VeboPaused)
	if err := _Vebo.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeboProcessingStartedIterator is returned from FilterProcessingStarted and is used to iterate over the raw logs and unpacked data for ProcessingStarted events raised by the Vebo contract.
type VeboProcessingStartedIterator struct {
	Event *VeboProcessingStarted // Event containing the contract specifics and raw log

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
func (it *VeboProcessingStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeboProcessingStarted)
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
		it.Event = new(VeboProcessingStarted)
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
func (it *VeboProcessingStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeboProcessingStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeboProcessingStarted represents a ProcessingStarted event raised by the Vebo contract.
type VeboProcessingStarted struct {
	RefSlot *big.Int
	Hash    [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterProcessingStarted is a free log retrieval operation binding the contract event 0xf73febded7d4502284718948a3e1d75406151c6326bde069424a584a4f6af87a.
//
// Solidity: event ProcessingStarted(uint256 indexed refSlot, bytes32 hash)
func (_Vebo *VeboFilterer) FilterProcessingStarted(opts *bind.FilterOpts, refSlot []*big.Int) (*VeboProcessingStartedIterator, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _Vebo.contract.FilterLogs(opts, "ProcessingStarted", refSlotRule)
	if err != nil {
		return nil, err
	}
	return &VeboProcessingStartedIterator{contract: _Vebo.contract, event: "ProcessingStarted", logs: logs, sub: sub}, nil
}

// WatchProcessingStarted is a free log subscription operation binding the contract event 0xf73febded7d4502284718948a3e1d75406151c6326bde069424a584a4f6af87a.
//
// Solidity: event ProcessingStarted(uint256 indexed refSlot, bytes32 hash)
func (_Vebo *VeboFilterer) WatchProcessingStarted(opts *bind.WatchOpts, sink chan<- *VeboProcessingStarted, refSlot []*big.Int) (event.Subscription, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _Vebo.contract.WatchLogs(opts, "ProcessingStarted", refSlotRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeboProcessingStarted)
				if err := _Vebo.contract.UnpackLog(event, "ProcessingStarted", log); err != nil {
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

// ParseProcessingStarted is a log parse operation binding the contract event 0xf73febded7d4502284718948a3e1d75406151c6326bde069424a584a4f6af87a.
//
// Solidity: event ProcessingStarted(uint256 indexed refSlot, bytes32 hash)
func (_Vebo *VeboFilterer) ParseProcessingStarted(log types.Log) (*VeboProcessingStarted, error) {
	event := new(VeboProcessingStarted)
	if err := _Vebo.contract.UnpackLog(event, "ProcessingStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeboReportDiscardedIterator is returned from FilterReportDiscarded and is used to iterate over the raw logs and unpacked data for ReportDiscarded events raised by the Vebo contract.
type VeboReportDiscardedIterator struct {
	Event *VeboReportDiscarded // Event containing the contract specifics and raw log

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
func (it *VeboReportDiscardedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeboReportDiscarded)
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
		it.Event = new(VeboReportDiscarded)
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
func (it *VeboReportDiscardedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeboReportDiscardedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeboReportDiscarded represents a ReportDiscarded event raised by the Vebo contract.
type VeboReportDiscarded struct {
	RefSlot *big.Int
	Hash    [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReportDiscarded is a free log retrieval operation binding the contract event 0xe21266bc27ee721ac10034efaf7fd724656ef471c75b8402cd8f07850af6b676.
//
// Solidity: event ReportDiscarded(uint256 indexed refSlot, bytes32 hash)
func (_Vebo *VeboFilterer) FilterReportDiscarded(opts *bind.FilterOpts, refSlot []*big.Int) (*VeboReportDiscardedIterator, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _Vebo.contract.FilterLogs(opts, "ReportDiscarded", refSlotRule)
	if err != nil {
		return nil, err
	}
	return &VeboReportDiscardedIterator{contract: _Vebo.contract, event: "ReportDiscarded", logs: logs, sub: sub}, nil
}

// WatchReportDiscarded is a free log subscription operation binding the contract event 0xe21266bc27ee721ac10034efaf7fd724656ef471c75b8402cd8f07850af6b676.
//
// Solidity: event ReportDiscarded(uint256 indexed refSlot, bytes32 hash)
func (_Vebo *VeboFilterer) WatchReportDiscarded(opts *bind.WatchOpts, sink chan<- *VeboReportDiscarded, refSlot []*big.Int) (event.Subscription, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _Vebo.contract.WatchLogs(opts, "ReportDiscarded", refSlotRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeboReportDiscarded)
				if err := _Vebo.contract.UnpackLog(event, "ReportDiscarded", log); err != nil {
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

// ParseReportDiscarded is a log parse operation binding the contract event 0xe21266bc27ee721ac10034efaf7fd724656ef471c75b8402cd8f07850af6b676.
//
// Solidity: event ReportDiscarded(uint256 indexed refSlot, bytes32 hash)
func (_Vebo *VeboFilterer) ParseReportDiscarded(log types.Log) (*VeboReportDiscarded, error) {
	event := new(VeboReportDiscarded)
	if err := _Vebo.contract.UnpackLog(event, "ReportDiscarded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeboReportSubmittedIterator is returned from FilterReportSubmitted and is used to iterate over the raw logs and unpacked data for ReportSubmitted events raised by the Vebo contract.
type VeboReportSubmittedIterator struct {
	Event *VeboReportSubmitted // Event containing the contract specifics and raw log

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
func (it *VeboReportSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeboReportSubmitted)
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
		it.Event = new(VeboReportSubmitted)
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
func (it *VeboReportSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeboReportSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeboReportSubmitted represents a ReportSubmitted event raised by the Vebo contract.
type VeboReportSubmitted struct {
	RefSlot                *big.Int
	Hash                   [32]byte
	ProcessingDeadlineTime *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterReportSubmitted is a free log retrieval operation binding the contract event 0xaed7d1a7a1831158dcda1e4214f5862f450bd3eb5721a5f322bf8c9fe1790b0a.
//
// Solidity: event ReportSubmitted(uint256 indexed refSlot, bytes32 hash, uint256 processingDeadlineTime)
func (_Vebo *VeboFilterer) FilterReportSubmitted(opts *bind.FilterOpts, refSlot []*big.Int) (*VeboReportSubmittedIterator, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _Vebo.contract.FilterLogs(opts, "ReportSubmitted", refSlotRule)
	if err != nil {
		return nil, err
	}
	return &VeboReportSubmittedIterator{contract: _Vebo.contract, event: "ReportSubmitted", logs: logs, sub: sub}, nil
}

// WatchReportSubmitted is a free log subscription operation binding the contract event 0xaed7d1a7a1831158dcda1e4214f5862f450bd3eb5721a5f322bf8c9fe1790b0a.
//
// Solidity: event ReportSubmitted(uint256 indexed refSlot, bytes32 hash, uint256 processingDeadlineTime)
func (_Vebo *VeboFilterer) WatchReportSubmitted(opts *bind.WatchOpts, sink chan<- *VeboReportSubmitted, refSlot []*big.Int) (event.Subscription, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _Vebo.contract.WatchLogs(opts, "ReportSubmitted", refSlotRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeboReportSubmitted)
				if err := _Vebo.contract.UnpackLog(event, "ReportSubmitted", log); err != nil {
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

// ParseReportSubmitted is a log parse operation binding the contract event 0xaed7d1a7a1831158dcda1e4214f5862f450bd3eb5721a5f322bf8c9fe1790b0a.
//
// Solidity: event ReportSubmitted(uint256 indexed refSlot, bytes32 hash, uint256 processingDeadlineTime)
func (_Vebo *VeboFilterer) ParseReportSubmitted(log types.Log) (*VeboReportSubmitted, error) {
	event := new(VeboReportSubmitted)
	if err := _Vebo.contract.UnpackLog(event, "ReportSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeboResumedIterator is returned from FilterResumed and is used to iterate over the raw logs and unpacked data for Resumed events raised by the Vebo contract.
type VeboResumedIterator struct {
	Event *VeboResumed // Event containing the contract specifics and raw log

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
func (it *VeboResumedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeboResumed)
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
		it.Event = new(VeboResumed)
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
func (it *VeboResumedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeboResumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeboResumed represents a Resumed event raised by the Vebo contract.
type VeboResumed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterResumed is a free log retrieval operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_Vebo *VeboFilterer) FilterResumed(opts *bind.FilterOpts) (*VeboResumedIterator, error) {

	logs, sub, err := _Vebo.contract.FilterLogs(opts, "Resumed")
	if err != nil {
		return nil, err
	}
	return &VeboResumedIterator{contract: _Vebo.contract, event: "Resumed", logs: logs, sub: sub}, nil
}

// WatchResumed is a free log subscription operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_Vebo *VeboFilterer) WatchResumed(opts *bind.WatchOpts, sink chan<- *VeboResumed) (event.Subscription, error) {

	logs, sub, err := _Vebo.contract.WatchLogs(opts, "Resumed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeboResumed)
				if err := _Vebo.contract.UnpackLog(event, "Resumed", log); err != nil {
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

// ParseResumed is a log parse operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_Vebo *VeboFilterer) ParseResumed(log types.Log) (*VeboResumed, error) {
	event := new(VeboResumed)
	if err := _Vebo.contract.UnpackLog(event, "Resumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeboRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Vebo contract.
type VeboRoleAdminChangedIterator struct {
	Event *VeboRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *VeboRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeboRoleAdminChanged)
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
		it.Event = new(VeboRoleAdminChanged)
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
func (it *VeboRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeboRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeboRoleAdminChanged represents a RoleAdminChanged event raised by the Vebo contract.
type VeboRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Vebo *VeboFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*VeboRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Vebo.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &VeboRoleAdminChangedIterator{contract: _Vebo.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Vebo *VeboFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *VeboRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Vebo.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeboRoleAdminChanged)
				if err := _Vebo.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Vebo *VeboFilterer) ParseRoleAdminChanged(log types.Log) (*VeboRoleAdminChanged, error) {
	event := new(VeboRoleAdminChanged)
	if err := _Vebo.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeboRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Vebo contract.
type VeboRoleGrantedIterator struct {
	Event *VeboRoleGranted // Event containing the contract specifics and raw log

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
func (it *VeboRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeboRoleGranted)
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
		it.Event = new(VeboRoleGranted)
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
func (it *VeboRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeboRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeboRoleGranted represents a RoleGranted event raised by the Vebo contract.
type VeboRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Vebo *VeboFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*VeboRoleGrantedIterator, error) {

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

	logs, sub, err := _Vebo.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VeboRoleGrantedIterator{contract: _Vebo.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Vebo *VeboFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *VeboRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Vebo.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeboRoleGranted)
				if err := _Vebo.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Vebo *VeboFilterer) ParseRoleGranted(log types.Log) (*VeboRoleGranted, error) {
	event := new(VeboRoleGranted)
	if err := _Vebo.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeboRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Vebo contract.
type VeboRoleRevokedIterator struct {
	Event *VeboRoleRevoked // Event containing the contract specifics and raw log

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
func (it *VeboRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeboRoleRevoked)
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
		it.Event = new(VeboRoleRevoked)
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
func (it *VeboRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeboRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeboRoleRevoked represents a RoleRevoked event raised by the Vebo contract.
type VeboRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Vebo *VeboFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*VeboRoleRevokedIterator, error) {

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

	logs, sub, err := _Vebo.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VeboRoleRevokedIterator{contract: _Vebo.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Vebo *VeboFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *VeboRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Vebo.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeboRoleRevoked)
				if err := _Vebo.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Vebo *VeboFilterer) ParseRoleRevoked(log types.Log) (*VeboRoleRevoked, error) {
	event := new(VeboRoleRevoked)
	if err := _Vebo.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeboValidatorExitRequestIterator is returned from FilterValidatorExitRequest and is used to iterate over the raw logs and unpacked data for ValidatorExitRequest events raised by the Vebo contract.
type VeboValidatorExitRequestIterator struct {
	Event *VeboValidatorExitRequest // Event containing the contract specifics and raw log

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
func (it *VeboValidatorExitRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeboValidatorExitRequest)
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
		it.Event = new(VeboValidatorExitRequest)
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
func (it *VeboValidatorExitRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeboValidatorExitRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeboValidatorExitRequest represents a ValidatorExitRequest event raised by the Vebo contract.
type VeboValidatorExitRequest struct {
	StakingModuleId *big.Int
	NodeOperatorId  *big.Int
	ValidatorIndex  *big.Int
	ValidatorPubkey []byte
	Timestamp       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterValidatorExitRequest is a free log retrieval operation binding the contract event 0x96395f55c4997466e5035d777f0e1ba82b8cae217aaad05cf07839eb7c75bcf2.
//
// Solidity: event ValidatorExitRequest(uint256 indexed stakingModuleId, uint256 indexed nodeOperatorId, uint256 indexed validatorIndex, bytes validatorPubkey, uint256 timestamp)
func (_Vebo *VeboFilterer) FilterValidatorExitRequest(opts *bind.FilterOpts, stakingModuleId []*big.Int, nodeOperatorId []*big.Int, validatorIndex []*big.Int) (*VeboValidatorExitRequestIterator, error) {

	var stakingModuleIdRule []interface{}
	for _, stakingModuleIdItem := range stakingModuleId {
		stakingModuleIdRule = append(stakingModuleIdRule, stakingModuleIdItem)
	}
	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var validatorIndexRule []interface{}
	for _, validatorIndexItem := range validatorIndex {
		validatorIndexRule = append(validatorIndexRule, validatorIndexItem)
	}

	logs, sub, err := _Vebo.contract.FilterLogs(opts, "ValidatorExitRequest", stakingModuleIdRule, nodeOperatorIdRule, validatorIndexRule)
	if err != nil {
		return nil, err
	}
	return &VeboValidatorExitRequestIterator{contract: _Vebo.contract, event: "ValidatorExitRequest", logs: logs, sub: sub}, nil
}

// WatchValidatorExitRequest is a free log subscription operation binding the contract event 0x96395f55c4997466e5035d777f0e1ba82b8cae217aaad05cf07839eb7c75bcf2.
//
// Solidity: event ValidatorExitRequest(uint256 indexed stakingModuleId, uint256 indexed nodeOperatorId, uint256 indexed validatorIndex, bytes validatorPubkey, uint256 timestamp)
func (_Vebo *VeboFilterer) WatchValidatorExitRequest(opts *bind.WatchOpts, sink chan<- *VeboValidatorExitRequest, stakingModuleId []*big.Int, nodeOperatorId []*big.Int, validatorIndex []*big.Int) (event.Subscription, error) {

	var stakingModuleIdRule []interface{}
	for _, stakingModuleIdItem := range stakingModuleId {
		stakingModuleIdRule = append(stakingModuleIdRule, stakingModuleIdItem)
	}
	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var validatorIndexRule []interface{}
	for _, validatorIndexItem := range validatorIndex {
		validatorIndexRule = append(validatorIndexRule, validatorIndexItem)
	}

	logs, sub, err := _Vebo.contract.WatchLogs(opts, "ValidatorExitRequest", stakingModuleIdRule, nodeOperatorIdRule, validatorIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeboValidatorExitRequest)
				if err := _Vebo.contract.UnpackLog(event, "ValidatorExitRequest", log); err != nil {
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

// ParseValidatorExitRequest is a log parse operation binding the contract event 0x96395f55c4997466e5035d777f0e1ba82b8cae217aaad05cf07839eb7c75bcf2.
//
// Solidity: event ValidatorExitRequest(uint256 indexed stakingModuleId, uint256 indexed nodeOperatorId, uint256 indexed validatorIndex, bytes validatorPubkey, uint256 timestamp)
func (_Vebo *VeboFilterer) ParseValidatorExitRequest(log types.Log) (*VeboValidatorExitRequest, error) {
	event := new(VeboValidatorExitRequest)
	if err := _Vebo.contract.UnpackLog(event, "ValidatorExitRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeboWarnDataIncompleteProcessingIterator is returned from FilterWarnDataIncompleteProcessing and is used to iterate over the raw logs and unpacked data for WarnDataIncompleteProcessing events raised by the Vebo contract.
type VeboWarnDataIncompleteProcessingIterator struct {
	Event *VeboWarnDataIncompleteProcessing // Event containing the contract specifics and raw log

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
func (it *VeboWarnDataIncompleteProcessingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeboWarnDataIncompleteProcessing)
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
		it.Event = new(VeboWarnDataIncompleteProcessing)
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
func (it *VeboWarnDataIncompleteProcessingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeboWarnDataIncompleteProcessingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeboWarnDataIncompleteProcessing represents a WarnDataIncompleteProcessing event raised by the Vebo contract.
type VeboWarnDataIncompleteProcessing struct {
	RefSlot           *big.Int
	RequestsProcessed *big.Int
	RequestsCount     *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterWarnDataIncompleteProcessing is a free log retrieval operation binding the contract event 0xefc67aab43195093a8d8ed25d52281d96de480748ece2787888c586e8e1e79b4.
//
// Solidity: event WarnDataIncompleteProcessing(uint256 indexed refSlot, uint256 requestsProcessed, uint256 requestsCount)
func (_Vebo *VeboFilterer) FilterWarnDataIncompleteProcessing(opts *bind.FilterOpts, refSlot []*big.Int) (*VeboWarnDataIncompleteProcessingIterator, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _Vebo.contract.FilterLogs(opts, "WarnDataIncompleteProcessing", refSlotRule)
	if err != nil {
		return nil, err
	}
	return &VeboWarnDataIncompleteProcessingIterator{contract: _Vebo.contract, event: "WarnDataIncompleteProcessing", logs: logs, sub: sub}, nil
}

// WatchWarnDataIncompleteProcessing is a free log subscription operation binding the contract event 0xefc67aab43195093a8d8ed25d52281d96de480748ece2787888c586e8e1e79b4.
//
// Solidity: event WarnDataIncompleteProcessing(uint256 indexed refSlot, uint256 requestsProcessed, uint256 requestsCount)
func (_Vebo *VeboFilterer) WatchWarnDataIncompleteProcessing(opts *bind.WatchOpts, sink chan<- *VeboWarnDataIncompleteProcessing, refSlot []*big.Int) (event.Subscription, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _Vebo.contract.WatchLogs(opts, "WarnDataIncompleteProcessing", refSlotRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeboWarnDataIncompleteProcessing)
				if err := _Vebo.contract.UnpackLog(event, "WarnDataIncompleteProcessing", log); err != nil {
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

// ParseWarnDataIncompleteProcessing is a log parse operation binding the contract event 0xefc67aab43195093a8d8ed25d52281d96de480748ece2787888c586e8e1e79b4.
//
// Solidity: event WarnDataIncompleteProcessing(uint256 indexed refSlot, uint256 requestsProcessed, uint256 requestsCount)
func (_Vebo *VeboFilterer) ParseWarnDataIncompleteProcessing(log types.Log) (*VeboWarnDataIncompleteProcessing, error) {
	event := new(VeboWarnDataIncompleteProcessing)
	if err := _Vebo.contract.UnpackLog(event, "WarnDataIncompleteProcessing", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeboWarnProcessingMissedIterator is returned from FilterWarnProcessingMissed and is used to iterate over the raw logs and unpacked data for WarnProcessingMissed events raised by the Vebo contract.
type VeboWarnProcessingMissedIterator struct {
	Event *VeboWarnProcessingMissed // Event containing the contract specifics and raw log

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
func (it *VeboWarnProcessingMissedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeboWarnProcessingMissed)
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
		it.Event = new(VeboWarnProcessingMissed)
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
func (it *VeboWarnProcessingMissedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeboWarnProcessingMissedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeboWarnProcessingMissed represents a WarnProcessingMissed event raised by the Vebo contract.
type VeboWarnProcessingMissed struct {
	RefSlot *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWarnProcessingMissed is a free log retrieval operation binding the contract event 0x800b849c8bf80718cf786c99d1091c079fe2c5e420a3ba7ba9b0ef8179ef2c38.
//
// Solidity: event WarnProcessingMissed(uint256 indexed refSlot)
func (_Vebo *VeboFilterer) FilterWarnProcessingMissed(opts *bind.FilterOpts, refSlot []*big.Int) (*VeboWarnProcessingMissedIterator, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _Vebo.contract.FilterLogs(opts, "WarnProcessingMissed", refSlotRule)
	if err != nil {
		return nil, err
	}
	return &VeboWarnProcessingMissedIterator{contract: _Vebo.contract, event: "WarnProcessingMissed", logs: logs, sub: sub}, nil
}

// WatchWarnProcessingMissed is a free log subscription operation binding the contract event 0x800b849c8bf80718cf786c99d1091c079fe2c5e420a3ba7ba9b0ef8179ef2c38.
//
// Solidity: event WarnProcessingMissed(uint256 indexed refSlot)
func (_Vebo *VeboFilterer) WatchWarnProcessingMissed(opts *bind.WatchOpts, sink chan<- *VeboWarnProcessingMissed, refSlot []*big.Int) (event.Subscription, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _Vebo.contract.WatchLogs(opts, "WarnProcessingMissed", refSlotRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeboWarnProcessingMissed)
				if err := _Vebo.contract.UnpackLog(event, "WarnProcessingMissed", log); err != nil {
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

// ParseWarnProcessingMissed is a log parse operation binding the contract event 0x800b849c8bf80718cf786c99d1091c079fe2c5e420a3ba7ba9b0ef8179ef2c38.
//
// Solidity: event WarnProcessingMissed(uint256 indexed refSlot)
func (_Vebo *VeboFilterer) ParseWarnProcessingMissed(log types.Log) (*VeboWarnProcessingMissed, error) {
	event := new(VeboWarnProcessingMissed)
	if err := _Vebo.contract.UnpackLog(event, "WarnProcessingMissed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
