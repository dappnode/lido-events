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
		"lido-events/internal/application/domain"
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

// CSFeeOracleReportData is an auto generated low-level Go binding around an user-defined struct.
type CSFeeOracleReportData struct {
	ConsensusVersion *big.Int
	RefSlot          *big.Int
	TreeRoot         [32]byte
	TreeCid          string
	LogCid           string
	Distributed      *big.Int
}

// BindingsMetaData contains all meta data concerning the Bindings contract.
var BindingsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"secondsPerSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"genesisTime\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AddressCannotBeSame\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedToSendEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HashCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialRefSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"processingRefSlot\",\"type\":\"uint256\"}],\"name\":\"InitialRefSlotCannotBeLessThanProcessingOne\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidContractVersion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidContractVersionIncrement\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPerfLeeway\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoConsensusReportToProcess\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonZeroContractVersionOnInit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllowedToRecover\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PauseUntilMustBeInFuture\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PausedExpected\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"ProcessingDeadlineMissed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RefSlotAlreadyProcessing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevRefSlot\",\"type\":\"uint256\"}],\"name\":\"RefSlotCannotDecrease\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"processingRefSlot\",\"type\":\"uint256\"}],\"name\":\"RefSlotMustBeGreaterThanProcessingOne\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ResumedExpected\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SecondsPerSlotCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderIsNotTheConsensusContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedChainConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedVersion\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"receivedVersion\",\"type\":\"uint256\"}],\"name\":\"UnexpectedConsensusVersion\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"received\",\"type\":\"uint256\"}],\"name\":\"UnexpectedContractVersion\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"consensusHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"receivedHash\",\"type\":\"bytes32\"}],\"name\":\"UnexpectedDataHash\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"consensusRefSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dataRefSlot\",\"type\":\"uint256\"}],\"name\":\"UnexpectedRefSlot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionCannotBeSame\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAdminAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroFeeDistributorAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroPauseDuration\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prevAddr\",\"type\":\"address\"}],\"name\":\"ConsensusHashContractSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"prevVersion\",\"type\":\"uint256\"}],\"name\":\"ConsensusVersionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"ContractVersionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC1155Recovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC20Recovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"ERC721Recovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EtherRecovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeDistributorContract\",\"type\":\"address\"}],\"name\":\"FeeDistributorContractSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"valueBP\",\"type\":\"uint256\"}],\"name\":\"PerfLeewaySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"ProcessingStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"ReportDiscarded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"processingDeadlineTime\",\"type\":\"uint256\"}],\"name\":\"ReportSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Resumed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"StETHSharesRecovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"}],\"name\":\"WarnProcessingMissed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CONTRACT_MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GENESIS_TIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGE_CONSENSUS_CONTRACT_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGE_CONSENSUS_VERSION_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSE_INFINITELY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RECOVERER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RESUME_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SECONDS_PER_SLOT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SUBMIT_DATA_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"avgPerfLeewayBP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"}],\"name\":\"discardConsensusReport\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeDistributor\",\"outputs\":[{\"internalType\":\"contractICSFeeDistributor\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConsensusContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConsensusReport\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"processingDeadlineTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"processingStarted\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConsensusVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastProcessingRefSlot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getResumeSinceTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeDistributorContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"consensusContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"consensusVersion\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_avgPerfLeewayBP\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"pauseFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pauseUntilInclusive\",\"type\":\"uint256\"}],\"name\":\"pauseUntil\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"recoverERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"recoverERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"recoverERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recoverEther\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resume\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setConsensusContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"setConsensusVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeDistributorContract\",\"type\":\"address\"}],\"name\":\"setFeeDistributorContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"valueBP\",\"type\":\"uint256\"}],\"name\":\"setPerformanceLeeway\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"reportHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"submitConsensusReport\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"consensusVersion\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"treeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"treeCid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"logCid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"distributed\",\"type\":\"uint256\"}],\"internalType\":\"structCSFeeOracle.ReportData\",\"name\":\"data\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"contractVersion\",\"type\":\"uint256\"}],\"name\":\"submitReportData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// CONTRACTMANAGERROLE is a free data retrieval call binding the contract method 0x9d28b506.
//
// Solidity: function CONTRACT_MANAGER_ROLE() view returns(bytes32)
func (_Bindings *BindingsCaller) CONTRACTMANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "CONTRACT_MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CONTRACTMANAGERROLE is a free data retrieval call binding the contract method 0x9d28b506.
//
// Solidity: function CONTRACT_MANAGER_ROLE() view returns(bytes32)
func (_Bindings *BindingsSession) CONTRACTMANAGERROLE() ([32]byte, error) {
	return _Bindings.Contract.CONTRACTMANAGERROLE(&_Bindings.CallOpts)
}

// CONTRACTMANAGERROLE is a free data retrieval call binding the contract method 0x9d28b506.
//
// Solidity: function CONTRACT_MANAGER_ROLE() view returns(bytes32)
func (_Bindings *BindingsCallerSession) CONTRACTMANAGERROLE() ([32]byte, error) {
	return _Bindings.Contract.CONTRACTMANAGERROLE(&_Bindings.CallOpts)
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

// GENESISTIME is a free data retrieval call binding the contract method 0xf2882461.
//
// Solidity: function GENESIS_TIME() view returns(uint256)
func (_Bindings *BindingsCaller) GENESISTIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "GENESIS_TIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GENESISTIME is a free data retrieval call binding the contract method 0xf2882461.
//
// Solidity: function GENESIS_TIME() view returns(uint256)
func (_Bindings *BindingsSession) GENESISTIME() (*big.Int, error) {
	return _Bindings.Contract.GENESISTIME(&_Bindings.CallOpts)
}

// GENESISTIME is a free data retrieval call binding the contract method 0xf2882461.
//
// Solidity: function GENESIS_TIME() view returns(uint256)
func (_Bindings *BindingsCallerSession) GENESISTIME() (*big.Int, error) {
	return _Bindings.Contract.GENESISTIME(&_Bindings.CallOpts)
}

// MANAGECONSENSUSCONTRACTROLE is a free data retrieval call binding the contract method 0xad5cac4e.
//
// Solidity: function MANAGE_CONSENSUS_CONTRACT_ROLE() view returns(bytes32)
func (_Bindings *BindingsCaller) MANAGECONSENSUSCONTRACTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "MANAGE_CONSENSUS_CONTRACT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGECONSENSUSCONTRACTROLE is a free data retrieval call binding the contract method 0xad5cac4e.
//
// Solidity: function MANAGE_CONSENSUS_CONTRACT_ROLE() view returns(bytes32)
func (_Bindings *BindingsSession) MANAGECONSENSUSCONTRACTROLE() ([32]byte, error) {
	return _Bindings.Contract.MANAGECONSENSUSCONTRACTROLE(&_Bindings.CallOpts)
}

// MANAGECONSENSUSCONTRACTROLE is a free data retrieval call binding the contract method 0xad5cac4e.
//
// Solidity: function MANAGE_CONSENSUS_CONTRACT_ROLE() view returns(bytes32)
func (_Bindings *BindingsCallerSession) MANAGECONSENSUSCONTRACTROLE() ([32]byte, error) {
	return _Bindings.Contract.MANAGECONSENSUSCONTRACTROLE(&_Bindings.CallOpts)
}

// MANAGECONSENSUSVERSIONROLE is a free data retrieval call binding the contract method 0x9cc23c79.
//
// Solidity: function MANAGE_CONSENSUS_VERSION_ROLE() view returns(bytes32)
func (_Bindings *BindingsCaller) MANAGECONSENSUSVERSIONROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "MANAGE_CONSENSUS_VERSION_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGECONSENSUSVERSIONROLE is a free data retrieval call binding the contract method 0x9cc23c79.
//
// Solidity: function MANAGE_CONSENSUS_VERSION_ROLE() view returns(bytes32)
func (_Bindings *BindingsSession) MANAGECONSENSUSVERSIONROLE() ([32]byte, error) {
	return _Bindings.Contract.MANAGECONSENSUSVERSIONROLE(&_Bindings.CallOpts)
}

// MANAGECONSENSUSVERSIONROLE is a free data retrieval call binding the contract method 0x9cc23c79.
//
// Solidity: function MANAGE_CONSENSUS_VERSION_ROLE() view returns(bytes32)
func (_Bindings *BindingsCallerSession) MANAGECONSENSUSVERSIONROLE() ([32]byte, error) {
	return _Bindings.Contract.MANAGECONSENSUSVERSIONROLE(&_Bindings.CallOpts)
}

// PAUSEINFINITELY is a free data retrieval call binding the contract method 0xa302ee38.
//
// Solidity: function PAUSE_INFINITELY() view returns(uint256)
func (_Bindings *BindingsCaller) PAUSEINFINITELY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "PAUSE_INFINITELY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PAUSEINFINITELY is a free data retrieval call binding the contract method 0xa302ee38.
//
// Solidity: function PAUSE_INFINITELY() view returns(uint256)
func (_Bindings *BindingsSession) PAUSEINFINITELY() (*big.Int, error) {
	return _Bindings.Contract.PAUSEINFINITELY(&_Bindings.CallOpts)
}

// PAUSEINFINITELY is a free data retrieval call binding the contract method 0xa302ee38.
//
// Solidity: function PAUSE_INFINITELY() view returns(uint256)
func (_Bindings *BindingsCallerSession) PAUSEINFINITELY() (*big.Int, error) {
	return _Bindings.Contract.PAUSEINFINITELY(&_Bindings.CallOpts)
}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_Bindings *BindingsCaller) PAUSEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "PAUSE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_Bindings *BindingsSession) PAUSEROLE() ([32]byte, error) {
	return _Bindings.Contract.PAUSEROLE(&_Bindings.CallOpts)
}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_Bindings *BindingsCallerSession) PAUSEROLE() ([32]byte, error) {
	return _Bindings.Contract.PAUSEROLE(&_Bindings.CallOpts)
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

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_Bindings *BindingsCaller) RESUMEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "RESUME_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_Bindings *BindingsSession) RESUMEROLE() ([32]byte, error) {
	return _Bindings.Contract.RESUMEROLE(&_Bindings.CallOpts)
}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_Bindings *BindingsCallerSession) RESUMEROLE() ([32]byte, error) {
	return _Bindings.Contract.RESUMEROLE(&_Bindings.CallOpts)
}

// SECONDSPERSLOT is a free data retrieval call binding the contract method 0x304b9071.
//
// Solidity: function SECONDS_PER_SLOT() view returns(uint256)
func (_Bindings *BindingsCaller) SECONDSPERSLOT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "SECONDS_PER_SLOT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SECONDSPERSLOT is a free data retrieval call binding the contract method 0x304b9071.
//
// Solidity: function SECONDS_PER_SLOT() view returns(uint256)
func (_Bindings *BindingsSession) SECONDSPERSLOT() (*big.Int, error) {
	return _Bindings.Contract.SECONDSPERSLOT(&_Bindings.CallOpts)
}

// SECONDSPERSLOT is a free data retrieval call binding the contract method 0x304b9071.
//
// Solidity: function SECONDS_PER_SLOT() view returns(uint256)
func (_Bindings *BindingsCallerSession) SECONDSPERSLOT() (*big.Int, error) {
	return _Bindings.Contract.SECONDSPERSLOT(&_Bindings.CallOpts)
}

// SUBMITDATAROLE is a free data retrieval call binding the contract method 0x46e1f576.
//
// Solidity: function SUBMIT_DATA_ROLE() view returns(bytes32)
func (_Bindings *BindingsCaller) SUBMITDATAROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "SUBMIT_DATA_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SUBMITDATAROLE is a free data retrieval call binding the contract method 0x46e1f576.
//
// Solidity: function SUBMIT_DATA_ROLE() view returns(bytes32)
func (_Bindings *BindingsSession) SUBMITDATAROLE() ([32]byte, error) {
	return _Bindings.Contract.SUBMITDATAROLE(&_Bindings.CallOpts)
}

// SUBMITDATAROLE is a free data retrieval call binding the contract method 0x46e1f576.
//
// Solidity: function SUBMIT_DATA_ROLE() view returns(bytes32)
func (_Bindings *BindingsCallerSession) SUBMITDATAROLE() ([32]byte, error) {
	return _Bindings.Contract.SUBMITDATAROLE(&_Bindings.CallOpts)
}

// AvgPerfLeewayBP is a free data retrieval call binding the contract method 0xf19b4006.
//
// Solidity: function avgPerfLeewayBP() view returns(uint256)
func (_Bindings *BindingsCaller) AvgPerfLeewayBP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "avgPerfLeewayBP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvgPerfLeewayBP is a free data retrieval call binding the contract method 0xf19b4006.
//
// Solidity: function avgPerfLeewayBP() view returns(uint256)
func (_Bindings *BindingsSession) AvgPerfLeewayBP() (*big.Int, error) {
	return _Bindings.Contract.AvgPerfLeewayBP(&_Bindings.CallOpts)
}

// AvgPerfLeewayBP is a free data retrieval call binding the contract method 0xf19b4006.
//
// Solidity: function avgPerfLeewayBP() view returns(uint256)
func (_Bindings *BindingsCallerSession) AvgPerfLeewayBP() (*big.Int, error) {
	return _Bindings.Contract.AvgPerfLeewayBP(&_Bindings.CallOpts)
}

// FeeDistributor is a free data retrieval call binding the contract method 0x0d43e8ad.
//
// Solidity: function feeDistributor() view returns(address)
func (_Bindings *BindingsCaller) FeeDistributor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "feeDistributor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeDistributor is a free data retrieval call binding the contract method 0x0d43e8ad.
//
// Solidity: function feeDistributor() view returns(address)
func (_Bindings *BindingsSession) FeeDistributor() (common.Address, error) {
	return _Bindings.Contract.FeeDistributor(&_Bindings.CallOpts)
}

// FeeDistributor is a free data retrieval call binding the contract method 0x0d43e8ad.
//
// Solidity: function feeDistributor() view returns(address)
func (_Bindings *BindingsCallerSession) FeeDistributor() (common.Address, error) {
	return _Bindings.Contract.FeeDistributor(&_Bindings.CallOpts)
}

// GetConsensusContract is a free data retrieval call binding the contract method 0x8f55b571.
//
// Solidity: function getConsensusContract() view returns(address)
func (_Bindings *BindingsCaller) GetConsensusContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getConsensusContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetConsensusContract is a free data retrieval call binding the contract method 0x8f55b571.
//
// Solidity: function getConsensusContract() view returns(address)
func (_Bindings *BindingsSession) GetConsensusContract() (common.Address, error) {
	return _Bindings.Contract.GetConsensusContract(&_Bindings.CallOpts)
}

// GetConsensusContract is a free data retrieval call binding the contract method 0x8f55b571.
//
// Solidity: function getConsensusContract() view returns(address)
func (_Bindings *BindingsCallerSession) GetConsensusContract() (common.Address, error) {
	return _Bindings.Contract.GetConsensusContract(&_Bindings.CallOpts)
}

// GetConsensusReport is a free data retrieval call binding the contract method 0x60d64d38.
//
// Solidity: function getConsensusReport() view returns(bytes32 hash, uint256 refSlot, uint256 processingDeadlineTime, bool processingStarted)
func (_Bindings *BindingsCaller) GetConsensusReport(opts *bind.CallOpts) (struct {
	Hash                   [32]byte
	RefSlot                *big.Int
	ProcessingDeadlineTime *big.Int
	ProcessingStarted      bool
}, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getConsensusReport")

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
func (_Bindings *BindingsSession) GetConsensusReport() (struct {
	Hash                   [32]byte
	RefSlot                *big.Int
	ProcessingDeadlineTime *big.Int
	ProcessingStarted      bool
}, error) {
	return _Bindings.Contract.GetConsensusReport(&_Bindings.CallOpts)
}

// GetConsensusReport is a free data retrieval call binding the contract method 0x60d64d38.
//
// Solidity: function getConsensusReport() view returns(bytes32 hash, uint256 refSlot, uint256 processingDeadlineTime, bool processingStarted)
func (_Bindings *BindingsCallerSession) GetConsensusReport() (struct {
	Hash                   [32]byte
	RefSlot                *big.Int
	ProcessingDeadlineTime *big.Int
	ProcessingStarted      bool
}, error) {
	return _Bindings.Contract.GetConsensusReport(&_Bindings.CallOpts)
}

// GetConsensusVersion is a free data retrieval call binding the contract method 0x5be20425.
//
// Solidity: function getConsensusVersion() view returns(uint256)
func (_Bindings *BindingsCaller) GetConsensusVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getConsensusVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetConsensusVersion is a free data retrieval call binding the contract method 0x5be20425.
//
// Solidity: function getConsensusVersion() view returns(uint256)
func (_Bindings *BindingsSession) GetConsensusVersion() (*big.Int, error) {
	return _Bindings.Contract.GetConsensusVersion(&_Bindings.CallOpts)
}

// GetConsensusVersion is a free data retrieval call binding the contract method 0x5be20425.
//
// Solidity: function getConsensusVersion() view returns(uint256)
func (_Bindings *BindingsCallerSession) GetConsensusVersion() (*big.Int, error) {
	return _Bindings.Contract.GetConsensusVersion(&_Bindings.CallOpts)
}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_Bindings *BindingsCaller) GetContractVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getContractVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_Bindings *BindingsSession) GetContractVersion() (*big.Int, error) {
	return _Bindings.Contract.GetContractVersion(&_Bindings.CallOpts)
}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_Bindings *BindingsCallerSession) GetContractVersion() (*big.Int, error) {
	return _Bindings.Contract.GetContractVersion(&_Bindings.CallOpts)
}

// GetLastProcessingRefSlot is a free data retrieval call binding the contract method 0x3584d59c.
//
// Solidity: function getLastProcessingRefSlot() view returns(uint256)
func (_Bindings *BindingsCaller) GetLastProcessingRefSlot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getLastProcessingRefSlot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastProcessingRefSlot is a free data retrieval call binding the contract method 0x3584d59c.
//
// Solidity: function getLastProcessingRefSlot() view returns(uint256)
func (_Bindings *BindingsSession) GetLastProcessingRefSlot() (*big.Int, error) {
	return _Bindings.Contract.GetLastProcessingRefSlot(&_Bindings.CallOpts)
}

// GetLastProcessingRefSlot is a free data retrieval call binding the contract method 0x3584d59c.
//
// Solidity: function getLastProcessingRefSlot() view returns(uint256)
func (_Bindings *BindingsCallerSession) GetLastProcessingRefSlot() (*big.Int, error) {
	return _Bindings.Contract.GetLastProcessingRefSlot(&_Bindings.CallOpts)
}

// GetResumeSinceTimestamp is a free data retrieval call binding the contract method 0x589ff76c.
//
// Solidity: function getResumeSinceTimestamp() view returns(uint256)
func (_Bindings *BindingsCaller) GetResumeSinceTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getResumeSinceTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetResumeSinceTimestamp is a free data retrieval call binding the contract method 0x589ff76c.
//
// Solidity: function getResumeSinceTimestamp() view returns(uint256)
func (_Bindings *BindingsSession) GetResumeSinceTimestamp() (*big.Int, error) {
	return _Bindings.Contract.GetResumeSinceTimestamp(&_Bindings.CallOpts)
}

// GetResumeSinceTimestamp is a free data retrieval call binding the contract method 0x589ff76c.
//
// Solidity: function getResumeSinceTimestamp() view returns(uint256)
func (_Bindings *BindingsCallerSession) GetResumeSinceTimestamp() (*big.Int, error) {
	return _Bindings.Contract.GetResumeSinceTimestamp(&_Bindings.CallOpts)
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

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Bindings *BindingsCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Bindings *BindingsSession) IsPaused() (bool, error) {
	return _Bindings.Contract.IsPaused(&_Bindings.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Bindings *BindingsCallerSession) IsPaused() (bool, error) {
	return _Bindings.Contract.IsPaused(&_Bindings.CallOpts)
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

// DiscardConsensusReport is a paid mutator transaction binding the contract method 0xd4381217.
//
// Solidity: function discardConsensusReport(uint256 refSlot) returns()
func (_Bindings *BindingsTransactor) DiscardConsensusReport(opts *bind.TransactOpts, refSlot *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "discardConsensusReport", refSlot)
}

// DiscardConsensusReport is a paid mutator transaction binding the contract method 0xd4381217.
//
// Solidity: function discardConsensusReport(uint256 refSlot) returns()
func (_Bindings *BindingsSession) DiscardConsensusReport(refSlot *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.DiscardConsensusReport(&_Bindings.TransactOpts, refSlot)
}

// DiscardConsensusReport is a paid mutator transaction binding the contract method 0xd4381217.
//
// Solidity: function discardConsensusReport(uint256 refSlot) returns()
func (_Bindings *BindingsTransactorSession) DiscardConsensusReport(refSlot *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.DiscardConsensusReport(&_Bindings.TransactOpts, refSlot)
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

// Initialize is a paid mutator transaction binding the contract method 0xa6b63eb8.
//
// Solidity: function initialize(address admin, address feeDistributorContract, address consensusContract, uint256 consensusVersion, uint256 _avgPerfLeewayBP) returns()
func (_Bindings *BindingsTransactor) Initialize(opts *bind.TransactOpts, admin common.Address, feeDistributorContract common.Address, consensusContract common.Address, consensusVersion *big.Int, _avgPerfLeewayBP *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "initialize", admin, feeDistributorContract, consensusContract, consensusVersion, _avgPerfLeewayBP)
}

// Initialize is a paid mutator transaction binding the contract method 0xa6b63eb8.
//
// Solidity: function initialize(address admin, address feeDistributorContract, address consensusContract, uint256 consensusVersion, uint256 _avgPerfLeewayBP) returns()
func (_Bindings *BindingsSession) Initialize(admin common.Address, feeDistributorContract common.Address, consensusContract common.Address, consensusVersion *big.Int, _avgPerfLeewayBP *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Initialize(&_Bindings.TransactOpts, admin, feeDistributorContract, consensusContract, consensusVersion, _avgPerfLeewayBP)
}

// Initialize is a paid mutator transaction binding the contract method 0xa6b63eb8.
//
// Solidity: function initialize(address admin, address feeDistributorContract, address consensusContract, uint256 consensusVersion, uint256 _avgPerfLeewayBP) returns()
func (_Bindings *BindingsTransactorSession) Initialize(admin common.Address, feeDistributorContract common.Address, consensusContract common.Address, consensusVersion *big.Int, _avgPerfLeewayBP *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Initialize(&_Bindings.TransactOpts, admin, feeDistributorContract, consensusContract, consensusVersion, _avgPerfLeewayBP)
}

// PauseFor is a paid mutator transaction binding the contract method 0xf3f449c7.
//
// Solidity: function pauseFor(uint256 duration) returns()
func (_Bindings *BindingsTransactor) PauseFor(opts *bind.TransactOpts, duration *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "pauseFor", duration)
}

// PauseFor is a paid mutator transaction binding the contract method 0xf3f449c7.
//
// Solidity: function pauseFor(uint256 duration) returns()
func (_Bindings *BindingsSession) PauseFor(duration *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.PauseFor(&_Bindings.TransactOpts, duration)
}

// PauseFor is a paid mutator transaction binding the contract method 0xf3f449c7.
//
// Solidity: function pauseFor(uint256 duration) returns()
func (_Bindings *BindingsTransactorSession) PauseFor(duration *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.PauseFor(&_Bindings.TransactOpts, duration)
}

// PauseUntil is a paid mutator transaction binding the contract method 0xabe9cfc8.
//
// Solidity: function pauseUntil(uint256 pauseUntilInclusive) returns()
func (_Bindings *BindingsTransactor) PauseUntil(opts *bind.TransactOpts, pauseUntilInclusive *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "pauseUntil", pauseUntilInclusive)
}

// PauseUntil is a paid mutator transaction binding the contract method 0xabe9cfc8.
//
// Solidity: function pauseUntil(uint256 pauseUntilInclusive) returns()
func (_Bindings *BindingsSession) PauseUntil(pauseUntilInclusive *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.PauseUntil(&_Bindings.TransactOpts, pauseUntilInclusive)
}

// PauseUntil is a paid mutator transaction binding the contract method 0xabe9cfc8.
//
// Solidity: function pauseUntil(uint256 pauseUntilInclusive) returns()
func (_Bindings *BindingsTransactorSession) PauseUntil(pauseUntilInclusive *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.PauseUntil(&_Bindings.TransactOpts, pauseUntilInclusive)
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

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Bindings *BindingsTransactor) Resume(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "resume")
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Bindings *BindingsSession) Resume() (*types.Transaction, error) {
	return _Bindings.Contract.Resume(&_Bindings.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Bindings *BindingsTransactorSession) Resume() (*types.Transaction, error) {
	return _Bindings.Contract.Resume(&_Bindings.TransactOpts)
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

// SetConsensusContract is a paid mutator transaction binding the contract method 0xc469c307.
//
// Solidity: function setConsensusContract(address addr) returns()
func (_Bindings *BindingsTransactor) SetConsensusContract(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setConsensusContract", addr)
}

// SetConsensusContract is a paid mutator transaction binding the contract method 0xc469c307.
//
// Solidity: function setConsensusContract(address addr) returns()
func (_Bindings *BindingsSession) SetConsensusContract(addr common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.SetConsensusContract(&_Bindings.TransactOpts, addr)
}

// SetConsensusContract is a paid mutator transaction binding the contract method 0xc469c307.
//
// Solidity: function setConsensusContract(address addr) returns()
func (_Bindings *BindingsTransactorSession) SetConsensusContract(addr common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.SetConsensusContract(&_Bindings.TransactOpts, addr)
}

// SetConsensusVersion is a paid mutator transaction binding the contract method 0x8d591474.
//
// Solidity: function setConsensusVersion(uint256 version) returns()
func (_Bindings *BindingsTransactor) SetConsensusVersion(opts *bind.TransactOpts, version *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setConsensusVersion", version)
}

// SetConsensusVersion is a paid mutator transaction binding the contract method 0x8d591474.
//
// Solidity: function setConsensusVersion(uint256 version) returns()
func (_Bindings *BindingsSession) SetConsensusVersion(version *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetConsensusVersion(&_Bindings.TransactOpts, version)
}

// SetConsensusVersion is a paid mutator transaction binding the contract method 0x8d591474.
//
// Solidity: function setConsensusVersion(uint256 version) returns()
func (_Bindings *BindingsTransactorSession) SetConsensusVersion(version *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetConsensusVersion(&_Bindings.TransactOpts, version)
}

// SetFeeDistributorContract is a paid mutator transaction binding the contract method 0xec18df3b.
//
// Solidity: function setFeeDistributorContract(address feeDistributorContract) returns()
func (_Bindings *BindingsTransactor) SetFeeDistributorContract(opts *bind.TransactOpts, feeDistributorContract common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setFeeDistributorContract", feeDistributorContract)
}

// SetFeeDistributorContract is a paid mutator transaction binding the contract method 0xec18df3b.
//
// Solidity: function setFeeDistributorContract(address feeDistributorContract) returns()
func (_Bindings *BindingsSession) SetFeeDistributorContract(feeDistributorContract common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.SetFeeDistributorContract(&_Bindings.TransactOpts, feeDistributorContract)
}

// SetFeeDistributorContract is a paid mutator transaction binding the contract method 0xec18df3b.
//
// Solidity: function setFeeDistributorContract(address feeDistributorContract) returns()
func (_Bindings *BindingsTransactorSession) SetFeeDistributorContract(feeDistributorContract common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.SetFeeDistributorContract(&_Bindings.TransactOpts, feeDistributorContract)
}

// SetPerformanceLeeway is a paid mutator transaction binding the contract method 0xf6b8549c.
//
// Solidity: function setPerformanceLeeway(uint256 valueBP) returns()
func (_Bindings *BindingsTransactor) SetPerformanceLeeway(opts *bind.TransactOpts, valueBP *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setPerformanceLeeway", valueBP)
}

// SetPerformanceLeeway is a paid mutator transaction binding the contract method 0xf6b8549c.
//
// Solidity: function setPerformanceLeeway(uint256 valueBP) returns()
func (_Bindings *BindingsSession) SetPerformanceLeeway(valueBP *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetPerformanceLeeway(&_Bindings.TransactOpts, valueBP)
}

// SetPerformanceLeeway is a paid mutator transaction binding the contract method 0xf6b8549c.
//
// Solidity: function setPerformanceLeeway(uint256 valueBP) returns()
func (_Bindings *BindingsTransactorSession) SetPerformanceLeeway(valueBP *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetPerformanceLeeway(&_Bindings.TransactOpts, valueBP)
}

// SubmitConsensusReport is a paid mutator transaction binding the contract method 0x063f36ad.
//
// Solidity: function submitConsensusReport(bytes32 reportHash, uint256 refSlot, uint256 deadline) returns()
func (_Bindings *BindingsTransactor) SubmitConsensusReport(opts *bind.TransactOpts, reportHash [32]byte, refSlot *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "submitConsensusReport", reportHash, refSlot, deadline)
}

// SubmitConsensusReport is a paid mutator transaction binding the contract method 0x063f36ad.
//
// Solidity: function submitConsensusReport(bytes32 reportHash, uint256 refSlot, uint256 deadline) returns()
func (_Bindings *BindingsSession) SubmitConsensusReport(reportHash [32]byte, refSlot *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SubmitConsensusReport(&_Bindings.TransactOpts, reportHash, refSlot, deadline)
}

// SubmitConsensusReport is a paid mutator transaction binding the contract method 0x063f36ad.
//
// Solidity: function submitConsensusReport(bytes32 reportHash, uint256 refSlot, uint256 deadline) returns()
func (_Bindings *BindingsTransactorSession) SubmitConsensusReport(reportHash [32]byte, refSlot *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SubmitConsensusReport(&_Bindings.TransactOpts, reportHash, refSlot, deadline)
}

// SubmitReportData is a paid mutator transaction binding the contract method 0xade4e312.
//
// Solidity: function submitReportData((uint256,uint256,bytes32,string,string,uint256) data, uint256 contractVersion) returns()
func (_Bindings *BindingsTransactor) SubmitReportData(opts *bind.TransactOpts, data CSFeeOracleReportData, contractVersion *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "submitReportData", data, contractVersion)
}

// SubmitReportData is a paid mutator transaction binding the contract method 0xade4e312.
//
// Solidity: function submitReportData((uint256,uint256,bytes32,string,string,uint256) data, uint256 contractVersion) returns()
func (_Bindings *BindingsSession) SubmitReportData(data CSFeeOracleReportData, contractVersion *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SubmitReportData(&_Bindings.TransactOpts, data, contractVersion)
}

// SubmitReportData is a paid mutator transaction binding the contract method 0xade4e312.
//
// Solidity: function submitReportData((uint256,uint256,bytes32,string,string,uint256) data, uint256 contractVersion) returns()
func (_Bindings *BindingsTransactorSession) SubmitReportData(data CSFeeOracleReportData, contractVersion *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SubmitReportData(&_Bindings.TransactOpts, data, contractVersion)
}

// BindingsConsensusHashContractSetIterator is returned from FilterConsensusHashContractSet and is used to iterate over the raw logs and unpacked data for ConsensusHashContractSet events raised by the Bindings contract.
type BindingsConsensusHashContractSetIterator struct {
	Event *BindingsConsensusHashContractSet // Event containing the contract specifics and raw log

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
func (it *BindingsConsensusHashContractSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsConsensusHashContractSet)
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
		it.Event = new(BindingsConsensusHashContractSet)
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
func (it *BindingsConsensusHashContractSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsConsensusHashContractSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsConsensusHashContractSet represents a ConsensusHashContractSet event raised by the Bindings contract.
type BindingsConsensusHashContractSet struct {
	Addr     common.Address
	PrevAddr common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterConsensusHashContractSet is a free log retrieval operation binding the contract event 0x25421480fb7f52d18947876279a213696b58d7e0e5416ce5e2c9f9942661c34c.
//
// Solidity: event ConsensusHashContractSet(address indexed addr, address indexed prevAddr)
func (_Bindings *BindingsFilterer) FilterConsensusHashContractSet(opts *bind.FilterOpts, addr []common.Address, prevAddr []common.Address) (*BindingsConsensusHashContractSetIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var prevAddrRule []interface{}
	for _, prevAddrItem := range prevAddr {
		prevAddrRule = append(prevAddrRule, prevAddrItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "ConsensusHashContractSet", addrRule, prevAddrRule)
	if err != nil {
		return nil, err
	}
	return &BindingsConsensusHashContractSetIterator{contract: _Bindings.contract, event: "ConsensusHashContractSet", logs: logs, sub: sub}, nil
}

// WatchConsensusHashContractSet is a free log subscription operation binding the contract event 0x25421480fb7f52d18947876279a213696b58d7e0e5416ce5e2c9f9942661c34c.
//
// Solidity: event ConsensusHashContractSet(address indexed addr, address indexed prevAddr)
func (_Bindings *BindingsFilterer) WatchConsensusHashContractSet(opts *bind.WatchOpts, sink chan<- *BindingsConsensusHashContractSet, addr []common.Address, prevAddr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var prevAddrRule []interface{}
	for _, prevAddrItem := range prevAddr {
		prevAddrRule = append(prevAddrRule, prevAddrItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "ConsensusHashContractSet", addrRule, prevAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsConsensusHashContractSet)
				if err := _Bindings.contract.UnpackLog(event, "ConsensusHashContractSet", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParseConsensusHashContractSet(log types.Log) (*BindingsConsensusHashContractSet, error) {
	event := new(BindingsConsensusHashContractSet)
	if err := _Bindings.contract.UnpackLog(event, "ConsensusHashContractSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsConsensusVersionSetIterator is returned from FilterConsensusVersionSet and is used to iterate over the raw logs and unpacked data for ConsensusVersionSet events raised by the Bindings contract.
type BindingsConsensusVersionSetIterator struct {
	Event *BindingsConsensusVersionSet // Event containing the contract specifics and raw log

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
func (it *BindingsConsensusVersionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsConsensusVersionSet)
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
		it.Event = new(BindingsConsensusVersionSet)
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
func (it *BindingsConsensusVersionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsConsensusVersionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsConsensusVersionSet represents a ConsensusVersionSet event raised by the Bindings contract.
type BindingsConsensusVersionSet struct {
	Version     *big.Int
	PrevVersion *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConsensusVersionSet is a free log retrieval operation binding the contract event 0xfa5304972d4ec3e3207f0bbf91155a49d0dfa62488f9529403a2a49e4b29a895.
//
// Solidity: event ConsensusVersionSet(uint256 indexed version, uint256 indexed prevVersion)
func (_Bindings *BindingsFilterer) FilterConsensusVersionSet(opts *bind.FilterOpts, version []*big.Int, prevVersion []*big.Int) (*BindingsConsensusVersionSetIterator, error) {

	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}
	var prevVersionRule []interface{}
	for _, prevVersionItem := range prevVersion {
		prevVersionRule = append(prevVersionRule, prevVersionItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "ConsensusVersionSet", versionRule, prevVersionRule)
	if err != nil {
		return nil, err
	}
	return &BindingsConsensusVersionSetIterator{contract: _Bindings.contract, event: "ConsensusVersionSet", logs: logs, sub: sub}, nil
}

// WatchConsensusVersionSet is a free log subscription operation binding the contract event 0xfa5304972d4ec3e3207f0bbf91155a49d0dfa62488f9529403a2a49e4b29a895.
//
// Solidity: event ConsensusVersionSet(uint256 indexed version, uint256 indexed prevVersion)
func (_Bindings *BindingsFilterer) WatchConsensusVersionSet(opts *bind.WatchOpts, sink chan<- *BindingsConsensusVersionSet, version []*big.Int, prevVersion []*big.Int) (event.Subscription, error) {

	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}
	var prevVersionRule []interface{}
	for _, prevVersionItem := range prevVersion {
		prevVersionRule = append(prevVersionRule, prevVersionItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "ConsensusVersionSet", versionRule, prevVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsConsensusVersionSet)
				if err := _Bindings.contract.UnpackLog(event, "ConsensusVersionSet", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParseConsensusVersionSet(log types.Log) (*BindingsConsensusVersionSet, error) {
	event := new(BindingsConsensusVersionSet)
	if err := _Bindings.contract.UnpackLog(event, "ConsensusVersionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsContractVersionSetIterator is returned from FilterContractVersionSet and is used to iterate over the raw logs and unpacked data for ContractVersionSet events raised by the Bindings contract.
type BindingsContractVersionSetIterator struct {
	Event *BindingsContractVersionSet // Event containing the contract specifics and raw log

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
func (it *BindingsContractVersionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsContractVersionSet)
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
		it.Event = new(BindingsContractVersionSet)
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
func (it *BindingsContractVersionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsContractVersionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsContractVersionSet represents a ContractVersionSet event raised by the Bindings contract.
type BindingsContractVersionSet struct {
	Version *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterContractVersionSet is a free log retrieval operation binding the contract event 0xfddcded6b4f4730c226821172046b48372d3cd963c159701ae1b7c3bcac541bb.
//
// Solidity: event ContractVersionSet(uint256 version)
func (_Bindings *BindingsFilterer) FilterContractVersionSet(opts *bind.FilterOpts) (*BindingsContractVersionSetIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "ContractVersionSet")
	if err != nil {
		return nil, err
	}
	return &BindingsContractVersionSetIterator{contract: _Bindings.contract, event: "ContractVersionSet", logs: logs, sub: sub}, nil
}

// WatchContractVersionSet is a free log subscription operation binding the contract event 0xfddcded6b4f4730c226821172046b48372d3cd963c159701ae1b7c3bcac541bb.
//
// Solidity: event ContractVersionSet(uint256 version)
func (_Bindings *BindingsFilterer) WatchContractVersionSet(opts *bind.WatchOpts, sink chan<- *BindingsContractVersionSet) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "ContractVersionSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsContractVersionSet)
				if err := _Bindings.contract.UnpackLog(event, "ContractVersionSet", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParseContractVersionSet(log types.Log) (*BindingsContractVersionSet, error) {
	event := new(BindingsContractVersionSet)
	if err := _Bindings.contract.UnpackLog(event, "ContractVersionSet", log); err != nil {
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

// BindingsFeeDistributorContractSetIterator is returned from FilterFeeDistributorContractSet and is used to iterate over the raw logs and unpacked data for FeeDistributorContractSet events raised by the Bindings contract.
type BindingsFeeDistributorContractSetIterator struct {
	Event *BindingsFeeDistributorContractSet // Event containing the contract specifics and raw log

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
func (it *BindingsFeeDistributorContractSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsFeeDistributorContractSet)
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
		it.Event = new(BindingsFeeDistributorContractSet)
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
func (it *BindingsFeeDistributorContractSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsFeeDistributorContractSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsFeeDistributorContractSet represents a FeeDistributorContractSet event raised by the Bindings contract.
type BindingsFeeDistributorContractSet struct {
	FeeDistributorContract common.Address
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterFeeDistributorContractSet is a free log retrieval operation binding the contract event 0xbd6deacc5dfdad351f7c5eb0a4c23a43ed5fafc59e427d7fa17878564c4d36cc.
//
// Solidity: event FeeDistributorContractSet(address feeDistributorContract)
func (_Bindings *BindingsFilterer) FilterFeeDistributorContractSet(opts *bind.FilterOpts) (*BindingsFeeDistributorContractSetIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "FeeDistributorContractSet")
	if err != nil {
		return nil, err
	}
	return &BindingsFeeDistributorContractSetIterator{contract: _Bindings.contract, event: "FeeDistributorContractSet", logs: logs, sub: sub}, nil
}

// WatchFeeDistributorContractSet is a free log subscription operation binding the contract event 0xbd6deacc5dfdad351f7c5eb0a4c23a43ed5fafc59e427d7fa17878564c4d36cc.
//
// Solidity: event FeeDistributorContractSet(address feeDistributorContract)
func (_Bindings *BindingsFilterer) WatchFeeDistributorContractSet(opts *bind.WatchOpts, sink chan<- *BindingsFeeDistributorContractSet) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "FeeDistributorContractSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsFeeDistributorContractSet)
				if err := _Bindings.contract.UnpackLog(event, "FeeDistributorContractSet", log); err != nil {
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

// ParseFeeDistributorContractSet is a log parse operation binding the contract event 0xbd6deacc5dfdad351f7c5eb0a4c23a43ed5fafc59e427d7fa17878564c4d36cc.
//
// Solidity: event FeeDistributorContractSet(address feeDistributorContract)
func (_Bindings *BindingsFilterer) ParseFeeDistributorContractSet(log types.Log) (*BindingsFeeDistributorContractSet, error) {
	event := new(BindingsFeeDistributorContractSet)
	if err := _Bindings.contract.UnpackLog(event, "FeeDistributorContractSet", log); err != nil {
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

// BindingsPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Bindings contract.
type BindingsPausedIterator struct {
	Event *BindingsPaused // Event containing the contract specifics and raw log

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
func (it *BindingsPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsPaused)
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
		it.Event = new(BindingsPaused)
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
func (it *BindingsPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsPaused represents a Paused event raised by the Bindings contract.
type BindingsPaused struct {
	Duration *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e.
//
// Solidity: event Paused(uint256 duration)
func (_Bindings *BindingsFilterer) FilterPaused(opts *bind.FilterOpts) (*BindingsPausedIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BindingsPausedIterator{contract: _Bindings.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e.
//
// Solidity: event Paused(uint256 duration)
func (_Bindings *BindingsFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BindingsPaused) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsPaused)
				if err := _Bindings.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParsePaused(log types.Log) (*BindingsPaused, error) {
	event := new(BindingsPaused)
	if err := _Bindings.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsPerfLeewaySetIterator is returned from FilterPerfLeewaySet and is used to iterate over the raw logs and unpacked data for PerfLeewaySet events raised by the Bindings contract.
type BindingsPerfLeewaySetIterator struct {
	Event *BindingsPerfLeewaySet // Event containing the contract specifics and raw log

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
func (it *BindingsPerfLeewaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsPerfLeewaySet)
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
		it.Event = new(BindingsPerfLeewaySet)
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
func (it *BindingsPerfLeewaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsPerfLeewaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsPerfLeewaySet represents a PerfLeewaySet event raised by the Bindings contract.
type BindingsPerfLeewaySet struct {
	ValueBP *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPerfLeewaySet is a free log retrieval operation binding the contract event 0xbbc9c6755631437a2dbadd5c3d3a1771ac765b634e2371b5493a167751675f65.
//
// Solidity: event PerfLeewaySet(uint256 valueBP)
func (_Bindings *BindingsFilterer) FilterPerfLeewaySet(opts *bind.FilterOpts) (*BindingsPerfLeewaySetIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "PerfLeewaySet")
	if err != nil {
		return nil, err
	}
	return &BindingsPerfLeewaySetIterator{contract: _Bindings.contract, event: "PerfLeewaySet", logs: logs, sub: sub}, nil
}

// WatchPerfLeewaySet is a free log subscription operation binding the contract event 0xbbc9c6755631437a2dbadd5c3d3a1771ac765b634e2371b5493a167751675f65.
//
// Solidity: event PerfLeewaySet(uint256 valueBP)
func (_Bindings *BindingsFilterer) WatchPerfLeewaySet(opts *bind.WatchOpts, sink chan<- *BindingsPerfLeewaySet) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "PerfLeewaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsPerfLeewaySet)
				if err := _Bindings.contract.UnpackLog(event, "PerfLeewaySet", log); err != nil {
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

// ParsePerfLeewaySet is a log parse operation binding the contract event 0xbbc9c6755631437a2dbadd5c3d3a1771ac765b634e2371b5493a167751675f65.
//
// Solidity: event PerfLeewaySet(uint256 valueBP)
func (_Bindings *BindingsFilterer) ParsePerfLeewaySet(log types.Log) (*BindingsPerfLeewaySet, error) {
	event := new(BindingsPerfLeewaySet)
	if err := _Bindings.contract.UnpackLog(event, "PerfLeewaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsProcessingStartedIterator is returned from FilterProcessingStarted and is used to iterate over the raw logs and unpacked data for ProcessingStarted events raised by the Bindings contract.
type BindingsProcessingStartedIterator struct {
	Event *domain.BindingsProcessingStarted // Event containing the contract specifics and raw log

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
func (it *BindingsProcessingStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(domain.BindingsProcessingStarted)
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
		it.Event = new(domain.BindingsProcessingStarted)
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
func (it *BindingsProcessingStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsProcessingStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsProcessingStarted represents a ProcessingStarted event raised by the Bindings contract.
// type BindingsProcessingStarted struct {
// 	RefSlot *big.Int
// 	Hash    [32]byte
// 	Raw     types.Log // Blockchain specific contextual infos
// }

// FilterProcessingStarted is a free log retrieval operation binding the contract event 0xf73febded7d4502284718948a3e1d75406151c6326bde069424a584a4f6af87a.
//
// Solidity: event ProcessingStarted(uint256 indexed refSlot, bytes32 hash)
func (_Bindings *BindingsFilterer) FilterProcessingStarted(opts *bind.FilterOpts, refSlot []*big.Int) (*BindingsProcessingStartedIterator, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "ProcessingStarted", refSlotRule)
	if err != nil {
		return nil, err
	}
	return &BindingsProcessingStartedIterator{contract: _Bindings.contract, event: "ProcessingStarted", logs: logs, sub: sub}, nil
}

// WatchProcessingStarted is a free log subscription operation binding the contract event 0xf73febded7d4502284718948a3e1d75406151c6326bde069424a584a4f6af87a.
//
// Solidity: event ProcessingStarted(uint256 indexed refSlot, bytes32 hash)
func (_Bindings *BindingsFilterer) WatchProcessingStarted(opts *bind.WatchOpts, sink chan<- *domain.BindingsProcessingStarted, refSlot []*big.Int) (event.Subscription, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "ProcessingStarted", refSlotRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(domain.BindingsProcessingStarted)
				if err := _Bindings.contract.UnpackLog(event, "ProcessingStarted", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParseProcessingStarted(log types.Log) (*domain.BindingsProcessingStarted, error) {
	event := new(domain.BindingsProcessingStarted)
	if err := _Bindings.contract.UnpackLog(event, "ProcessingStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsReportDiscardedIterator is returned from FilterReportDiscarded and is used to iterate over the raw logs and unpacked data for ReportDiscarded events raised by the Bindings contract.
type BindingsReportDiscardedIterator struct {
	Event *BindingsReportDiscarded // Event containing the contract specifics and raw log

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
func (it *BindingsReportDiscardedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsReportDiscarded)
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
		it.Event = new(BindingsReportDiscarded)
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
func (it *BindingsReportDiscardedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsReportDiscardedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsReportDiscarded represents a ReportDiscarded event raised by the Bindings contract.
type BindingsReportDiscarded struct {
	RefSlot *big.Int
	Hash    [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReportDiscarded is a free log retrieval operation binding the contract event 0xe21266bc27ee721ac10034efaf7fd724656ef471c75b8402cd8f07850af6b676.
//
// Solidity: event ReportDiscarded(uint256 indexed refSlot, bytes32 hash)
func (_Bindings *BindingsFilterer) FilterReportDiscarded(opts *bind.FilterOpts, refSlot []*big.Int) (*BindingsReportDiscardedIterator, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "ReportDiscarded", refSlotRule)
	if err != nil {
		return nil, err
	}
	return &BindingsReportDiscardedIterator{contract: _Bindings.contract, event: "ReportDiscarded", logs: logs, sub: sub}, nil
}

// WatchReportDiscarded is a free log subscription operation binding the contract event 0xe21266bc27ee721ac10034efaf7fd724656ef471c75b8402cd8f07850af6b676.
//
// Solidity: event ReportDiscarded(uint256 indexed refSlot, bytes32 hash)
func (_Bindings *BindingsFilterer) WatchReportDiscarded(opts *bind.WatchOpts, sink chan<- *BindingsReportDiscarded, refSlot []*big.Int) (event.Subscription, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "ReportDiscarded", refSlotRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsReportDiscarded)
				if err := _Bindings.contract.UnpackLog(event, "ReportDiscarded", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParseReportDiscarded(log types.Log) (*BindingsReportDiscarded, error) {
	event := new(BindingsReportDiscarded)
	if err := _Bindings.contract.UnpackLog(event, "ReportDiscarded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsReportSubmittedIterator is returned from FilterReportSubmitted and is used to iterate over the raw logs and unpacked data for ReportSubmitted events raised by the Bindings contract.
type BindingsReportSubmittedIterator struct {
	Event *BindingsReportSubmitted // Event containing the contract specifics and raw log

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
func (it *BindingsReportSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsReportSubmitted)
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
		it.Event = new(BindingsReportSubmitted)
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
func (it *BindingsReportSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsReportSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsReportSubmitted represents a ReportSubmitted event raised by the Bindings contract.
type BindingsReportSubmitted struct {
	RefSlot                *big.Int
	Hash                   [32]byte
	ProcessingDeadlineTime *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterReportSubmitted is a free log retrieval operation binding the contract event 0xaed7d1a7a1831158dcda1e4214f5862f450bd3eb5721a5f322bf8c9fe1790b0a.
//
// Solidity: event ReportSubmitted(uint256 indexed refSlot, bytes32 hash, uint256 processingDeadlineTime)
func (_Bindings *BindingsFilterer) FilterReportSubmitted(opts *bind.FilterOpts, refSlot []*big.Int) (*BindingsReportSubmittedIterator, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "ReportSubmitted", refSlotRule)
	if err != nil {
		return nil, err
	}
	return &BindingsReportSubmittedIterator{contract: _Bindings.contract, event: "ReportSubmitted", logs: logs, sub: sub}, nil
}

// WatchReportSubmitted is a free log subscription operation binding the contract event 0xaed7d1a7a1831158dcda1e4214f5862f450bd3eb5721a5f322bf8c9fe1790b0a.
//
// Solidity: event ReportSubmitted(uint256 indexed refSlot, bytes32 hash, uint256 processingDeadlineTime)
func (_Bindings *BindingsFilterer) WatchReportSubmitted(opts *bind.WatchOpts, sink chan<- *BindingsReportSubmitted, refSlot []*big.Int) (event.Subscription, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "ReportSubmitted", refSlotRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsReportSubmitted)
				if err := _Bindings.contract.UnpackLog(event, "ReportSubmitted", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParseReportSubmitted(log types.Log) (*BindingsReportSubmitted, error) {
	event := new(BindingsReportSubmitted)
	if err := _Bindings.contract.UnpackLog(event, "ReportSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsResumedIterator is returned from FilterResumed and is used to iterate over the raw logs and unpacked data for Resumed events raised by the Bindings contract.
type BindingsResumedIterator struct {
	Event *BindingsResumed // Event containing the contract specifics and raw log

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
func (it *BindingsResumedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsResumed)
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
		it.Event = new(BindingsResumed)
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
func (it *BindingsResumedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsResumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsResumed represents a Resumed event raised by the Bindings contract.
type BindingsResumed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterResumed is a free log retrieval operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_Bindings *BindingsFilterer) FilterResumed(opts *bind.FilterOpts) (*BindingsResumedIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "Resumed")
	if err != nil {
		return nil, err
	}
	return &BindingsResumedIterator{contract: _Bindings.contract, event: "Resumed", logs: logs, sub: sub}, nil
}

// WatchResumed is a free log subscription operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_Bindings *BindingsFilterer) WatchResumed(opts *bind.WatchOpts, sink chan<- *BindingsResumed) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "Resumed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsResumed)
				if err := _Bindings.contract.UnpackLog(event, "Resumed", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParseResumed(log types.Log) (*BindingsResumed, error) {
	event := new(BindingsResumed)
	if err := _Bindings.contract.UnpackLog(event, "Resumed", log); err != nil {
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

// BindingsWarnProcessingMissedIterator is returned from FilterWarnProcessingMissed and is used to iterate over the raw logs and unpacked data for WarnProcessingMissed events raised by the Bindings contract.
type BindingsWarnProcessingMissedIterator struct {
	Event *BindingsWarnProcessingMissed // Event containing the contract specifics and raw log

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
func (it *BindingsWarnProcessingMissedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsWarnProcessingMissed)
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
		it.Event = new(BindingsWarnProcessingMissed)
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
func (it *BindingsWarnProcessingMissedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsWarnProcessingMissedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsWarnProcessingMissed represents a WarnProcessingMissed event raised by the Bindings contract.
type BindingsWarnProcessingMissed struct {
	RefSlot *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWarnProcessingMissed is a free log retrieval operation binding the contract event 0x800b849c8bf80718cf786c99d1091c079fe2c5e420a3ba7ba9b0ef8179ef2c38.
//
// Solidity: event WarnProcessingMissed(uint256 indexed refSlot)
func (_Bindings *BindingsFilterer) FilterWarnProcessingMissed(opts *bind.FilterOpts, refSlot []*big.Int) (*BindingsWarnProcessingMissedIterator, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "WarnProcessingMissed", refSlotRule)
	if err != nil {
		return nil, err
	}
	return &BindingsWarnProcessingMissedIterator{contract: _Bindings.contract, event: "WarnProcessingMissed", logs: logs, sub: sub}, nil
}

// WatchWarnProcessingMissed is a free log subscription operation binding the contract event 0x800b849c8bf80718cf786c99d1091c079fe2c5e420a3ba7ba9b0ef8179ef2c38.
//
// Solidity: event WarnProcessingMissed(uint256 indexed refSlot)
func (_Bindings *BindingsFilterer) WatchWarnProcessingMissed(opts *bind.WatchOpts, sink chan<- *BindingsWarnProcessingMissed, refSlot []*big.Int) (event.Subscription, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "WarnProcessingMissed", refSlotRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsWarnProcessingMissed)
				if err := _Bindings.contract.UnpackLog(event, "WarnProcessingMissed", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParseWarnProcessingMissed(log types.Log) (*BindingsWarnProcessingMissed, error) {
	event := new(BindingsWarnProcessingMissed)
	if err := _Bindings.contract.UnpackLog(event, "WarnProcessingMissed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
