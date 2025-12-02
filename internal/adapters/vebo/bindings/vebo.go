// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"lido-events/internal/application/domain"
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

// ValidatorsExitBusExitRequestsData is an auto generated low-level Go binding around an user-defined struct.
type ValidatorsExitBusExitRequestsData struct {
	Data       []byte
	DataFormat *big.Int
}

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

// BindingsMetaData contains all meta data concerning the Bindings contract.
var BindingsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"secondsPerSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"genesisTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"lidoLocator\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AddressCannotBeSame\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AdminCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"exitDataIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestsCount\",\"type\":\"uint256\"}],\"name\":\"ExitDataIndexOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExitHashAlreadySubmitted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExitHashNotSubmitted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingLimit\",\"type\":\"uint256\"}],\"name\":\"ExitRequestsLimitExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HashCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialRefSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"processingRefSlot\",\"type\":\"uint256\"}],\"name\":\"InitialRefSlotCannotBeLessThanProcessingOne\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidContractVersionIncrement\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidExitDataIndexSortOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidModuleId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRequestsDataLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRequestsDataSortOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LimitExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoConsensusReportToProcess\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonZeroContractVersionOnInit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PauseUntilMustBeInFuture\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PausedExpected\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"ProcessingDeadlineMissed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RefSlotAlreadyProcessing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevRefSlot\",\"type\":\"uint256\"}],\"name\":\"RefSlotCannotDecrease\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"processingRefSlot\",\"type\":\"uint256\"}],\"name\":\"RefSlotMustBeGreaterThanProcessingOne\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RequestsAlreadyDelivered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RequestsNotDelivered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ResumedExpected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SecondsPerSlotCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderIsNotTheConsensusContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooLargeExitsPerFrame\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooLargeFrameDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooLargeMaxExitRequestsLimit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRequestsPerReport\",\"type\":\"uint256\"}],\"name\":\"TooManyExitRequestsInReport\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedChainConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedVersion\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"receivedVersion\",\"type\":\"uint256\"}],\"name\":\"UnexpectedConsensusVersion\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"received\",\"type\":\"uint256\"}],\"name\":\"UnexpectedContractVersion\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"consensusHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"receivedHash\",\"type\":\"bytes32\"}],\"name\":\"UnexpectedDataHash\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"consensusRefSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dataRefSlot\",\"type\":\"uint256\"}],\"name\":\"UnexpectedRefSlot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedRequestsDataLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"format\",\"type\":\"uint256\"}],\"name\":\"UnsupportedRequestsDataFormat\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionCannotBeSame\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"ZeroArgument\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroFrameDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroPauseDuration\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prevAddr\",\"type\":\"address\"}],\"name\":\"ConsensusHashContractSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"prevVersion\",\"type\":\"uint256\"}],\"name\":\"ConsensusVersionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"ContractVersionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"exitRequestsHash\",\"type\":\"bytes32\"}],\"name\":\"ExitDataProcessing\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxExitRequestsLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exitsPerFrame\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"frameDurationInSec\",\"type\":\"uint256\"}],\"name\":\"ExitRequestsLimitSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"ProcessingStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"ReportDiscarded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"processingDeadlineTime\",\"type\":\"uint256\"}],\"name\":\"ReportSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"exitRequestsHash\",\"type\":\"bytes32\"}],\"name\":\"RequestsHashSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Resumed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxValidatorsPerReport\",\"type\":\"uint256\"}],\"name\":\"SetMaxValidatorsPerReport\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakingModuleId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"validatorPubkey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ValidatorExitRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestsProcessed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestsCount\",\"type\":\"uint256\"}],\"name\":\"WarnDataIncompleteProcessing\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"}],\"name\":\"WarnProcessingMissed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DATA_FORMAT_LIST\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXIT_REQUEST_LIMIT_MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXIT_TYPE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GENESIS_TIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGE_CONSENSUS_CONTRACT_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGE_CONSENSUS_VERSION_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSE_INFINITELY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RESUME_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SECONDS_PER_SLOT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SUBMIT_DATA_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SUBMIT_REPORT_HASH_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"}],\"name\":\"discardConsensusReport\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxValidatorsPerReport\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxExitRequestsLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exitsPerFrame\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"frameDurationInSec\",\"type\":\"uint256\"}],\"name\":\"finalizeUpgrade_v2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConsensusContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConsensusReport\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"processingDeadlineTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"processingStarted\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConsensusVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"exitRequestsHash\",\"type\":\"bytes32\"}],\"name\":\"getDeliveryTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"deliveryDateTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExitRequestLimitFullInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxExitRequestsLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exitsPerFrame\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"frameDurationInSec\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevExitRequestsLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentExitRequestsLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastProcessingRefSlot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxValidatorsPerReport\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProcessingState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"currentFrameRefSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"processingDeadlineTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"dataSubmitted\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"dataFormat\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestsSubmitted\",\"type\":\"uint256\"}],\"internalType\":\"structValidatorsExitBusOracle.ProcessingState\",\"name\":\"result\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getResumeSinceTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalRequestsProcessed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"consensusContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"consensusVersion\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastProcessingRefSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValidatorsPerRequest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxExitRequestsLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exitsPerFrame\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"frameDurationInSec\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"pauseFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pauseUntilInclusive\",\"type\":\"uint256\"}],\"name\":\"pauseUntil\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resume\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setConsensusContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"setConsensusVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxExitRequestsLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exitsPerFrame\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"frameDurationInSec\",\"type\":\"uint256\"}],\"name\":\"setExitRequestLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxRequests\",\"type\":\"uint256\"}],\"name\":\"setMaxValidatorsPerReport\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"reportHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"submitConsensusReport\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"dataFormat\",\"type\":\"uint256\"}],\"internalType\":\"structValidatorsExitBus.ExitRequestsData\",\"name\":\"request\",\"type\":\"tuple\"}],\"name\":\"submitExitRequestsData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"exitRequestsHash\",\"type\":\"bytes32\"}],\"name\":\"submitExitRequestsHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"consensusVersion\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dataFormat\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structValidatorsExitBusOracle.ReportData\",\"name\":\"data\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"contractVersion\",\"type\":\"uint256\"}],\"name\":\"submitReportData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"dataFormat\",\"type\":\"uint256\"}],\"internalType\":\"structValidatorsExitBus.ExitRequestsData\",\"name\":\"exitsData\",\"type\":\"tuple\"},{\"internalType\":\"uint256[]\",\"name\":\"exitDataIndexes\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"refundRecipient\",\"type\":\"address\"}],\"name\":\"triggerExits\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"exitRequests\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"dataFormat\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"unpackExitRequest\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"nodeOpId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"moduleId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valIndex\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
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

// DATAFORMATLIST is a free data retrieval call binding the contract method 0xe271b774.
//
// Solidity: function DATA_FORMAT_LIST() view returns(uint256)
func (_Bindings *BindingsCaller) DATAFORMATLIST(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "DATA_FORMAT_LIST")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DATAFORMATLIST is a free data retrieval call binding the contract method 0xe271b774.
//
// Solidity: function DATA_FORMAT_LIST() view returns(uint256)
func (_Bindings *BindingsSession) DATAFORMATLIST() (*big.Int, error) {
	return _Bindings.Contract.DATAFORMATLIST(&_Bindings.CallOpts)
}

// DATAFORMATLIST is a free data retrieval call binding the contract method 0xe271b774.
//
// Solidity: function DATA_FORMAT_LIST() view returns(uint256)
func (_Bindings *BindingsCallerSession) DATAFORMATLIST() (*big.Int, error) {
	return _Bindings.Contract.DATAFORMATLIST(&_Bindings.CallOpts)
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

// EXITREQUESTLIMITMANAGERROLE is a free data retrieval call binding the contract method 0xab53ac48.
//
// Solidity: function EXIT_REQUEST_LIMIT_MANAGER_ROLE() view returns(bytes32)
func (_Bindings *BindingsCaller) EXITREQUESTLIMITMANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "EXIT_REQUEST_LIMIT_MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EXITREQUESTLIMITMANAGERROLE is a free data retrieval call binding the contract method 0xab53ac48.
//
// Solidity: function EXIT_REQUEST_LIMIT_MANAGER_ROLE() view returns(bytes32)
func (_Bindings *BindingsSession) EXITREQUESTLIMITMANAGERROLE() ([32]byte, error) {
	return _Bindings.Contract.EXITREQUESTLIMITMANAGERROLE(&_Bindings.CallOpts)
}

// EXITREQUESTLIMITMANAGERROLE is a free data retrieval call binding the contract method 0xab53ac48.
//
// Solidity: function EXIT_REQUEST_LIMIT_MANAGER_ROLE() view returns(bytes32)
func (_Bindings *BindingsCallerSession) EXITREQUESTLIMITMANAGERROLE() ([32]byte, error) {
	return _Bindings.Contract.EXITREQUESTLIMITMANAGERROLE(&_Bindings.CallOpts)
}

// EXITTYPE is a free data retrieval call binding the contract method 0x06e41389.
//
// Solidity: function EXIT_TYPE() view returns(uint256)
func (_Bindings *BindingsCaller) EXITTYPE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "EXIT_TYPE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EXITTYPE is a free data retrieval call binding the contract method 0x06e41389.
//
// Solidity: function EXIT_TYPE() view returns(uint256)
func (_Bindings *BindingsSession) EXITTYPE() (*big.Int, error) {
	return _Bindings.Contract.EXITTYPE(&_Bindings.CallOpts)
}

// EXITTYPE is a free data retrieval call binding the contract method 0x06e41389.
//
// Solidity: function EXIT_TYPE() view returns(uint256)
func (_Bindings *BindingsCallerSession) EXITTYPE() (*big.Int, error) {
	return _Bindings.Contract.EXITTYPE(&_Bindings.CallOpts)
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

// SUBMITREPORTHASHROLE is a free data retrieval call binding the contract method 0xd072f014.
//
// Solidity: function SUBMIT_REPORT_HASH_ROLE() view returns(bytes32)
func (_Bindings *BindingsCaller) SUBMITREPORTHASHROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "SUBMIT_REPORT_HASH_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SUBMITREPORTHASHROLE is a free data retrieval call binding the contract method 0xd072f014.
//
// Solidity: function SUBMIT_REPORT_HASH_ROLE() view returns(bytes32)
func (_Bindings *BindingsSession) SUBMITREPORTHASHROLE() ([32]byte, error) {
	return _Bindings.Contract.SUBMITREPORTHASHROLE(&_Bindings.CallOpts)
}

// SUBMITREPORTHASHROLE is a free data retrieval call binding the contract method 0xd072f014.
//
// Solidity: function SUBMIT_REPORT_HASH_ROLE() view returns(bytes32)
func (_Bindings *BindingsCallerSession) SUBMITREPORTHASHROLE() ([32]byte, error) {
	return _Bindings.Contract.SUBMITREPORTHASHROLE(&_Bindings.CallOpts)
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

// GetDeliveryTimestamp is a free data retrieval call binding the contract method 0xa52289bf.
//
// Solidity: function getDeliveryTimestamp(bytes32 exitRequestsHash) view returns(uint256 deliveryDateTimestamp)
func (_Bindings *BindingsCaller) GetDeliveryTimestamp(opts *bind.CallOpts, exitRequestsHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getDeliveryTimestamp", exitRequestsHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDeliveryTimestamp is a free data retrieval call binding the contract method 0xa52289bf.
//
// Solidity: function getDeliveryTimestamp(bytes32 exitRequestsHash) view returns(uint256 deliveryDateTimestamp)
func (_Bindings *BindingsSession) GetDeliveryTimestamp(exitRequestsHash [32]byte) (*big.Int, error) {
	return _Bindings.Contract.GetDeliveryTimestamp(&_Bindings.CallOpts, exitRequestsHash)
}

// GetDeliveryTimestamp is a free data retrieval call binding the contract method 0xa52289bf.
//
// Solidity: function getDeliveryTimestamp(bytes32 exitRequestsHash) view returns(uint256 deliveryDateTimestamp)
func (_Bindings *BindingsCallerSession) GetDeliveryTimestamp(exitRequestsHash [32]byte) (*big.Int, error) {
	return _Bindings.Contract.GetDeliveryTimestamp(&_Bindings.CallOpts, exitRequestsHash)
}

// GetExitRequestLimitFullInfo is a free data retrieval call binding the contract method 0xb6b764b2.
//
// Solidity: function getExitRequestLimitFullInfo() view returns(uint256 maxExitRequestsLimit, uint256 exitsPerFrame, uint256 frameDurationInSec, uint256 prevExitRequestsLimit, uint256 currentExitRequestsLimit)
func (_Bindings *BindingsCaller) GetExitRequestLimitFullInfo(opts *bind.CallOpts) (struct {
	MaxExitRequestsLimit     *big.Int
	ExitsPerFrame            *big.Int
	FrameDurationInSec       *big.Int
	PrevExitRequestsLimit    *big.Int
	CurrentExitRequestsLimit *big.Int
}, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getExitRequestLimitFullInfo")

	outstruct := new(struct {
		MaxExitRequestsLimit     *big.Int
		ExitsPerFrame            *big.Int
		FrameDurationInSec       *big.Int
		PrevExitRequestsLimit    *big.Int
		CurrentExitRequestsLimit *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MaxExitRequestsLimit = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ExitsPerFrame = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.FrameDurationInSec = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.PrevExitRequestsLimit = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.CurrentExitRequestsLimit = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetExitRequestLimitFullInfo is a free data retrieval call binding the contract method 0xb6b764b2.
//
// Solidity: function getExitRequestLimitFullInfo() view returns(uint256 maxExitRequestsLimit, uint256 exitsPerFrame, uint256 frameDurationInSec, uint256 prevExitRequestsLimit, uint256 currentExitRequestsLimit)
func (_Bindings *BindingsSession) GetExitRequestLimitFullInfo() (struct {
	MaxExitRequestsLimit     *big.Int
	ExitsPerFrame            *big.Int
	FrameDurationInSec       *big.Int
	PrevExitRequestsLimit    *big.Int
	CurrentExitRequestsLimit *big.Int
}, error) {
	return _Bindings.Contract.GetExitRequestLimitFullInfo(&_Bindings.CallOpts)
}

// GetExitRequestLimitFullInfo is a free data retrieval call binding the contract method 0xb6b764b2.
//
// Solidity: function getExitRequestLimitFullInfo() view returns(uint256 maxExitRequestsLimit, uint256 exitsPerFrame, uint256 frameDurationInSec, uint256 prevExitRequestsLimit, uint256 currentExitRequestsLimit)
func (_Bindings *BindingsCallerSession) GetExitRequestLimitFullInfo() (struct {
	MaxExitRequestsLimit     *big.Int
	ExitsPerFrame            *big.Int
	FrameDurationInSec       *big.Int
	PrevExitRequestsLimit    *big.Int
	CurrentExitRequestsLimit *big.Int
}, error) {
	return _Bindings.Contract.GetExitRequestLimitFullInfo(&_Bindings.CallOpts)
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

// GetMaxValidatorsPerReport is a free data retrieval call binding the contract method 0xc1f665bc.
//
// Solidity: function getMaxValidatorsPerReport() view returns(uint256)
func (_Bindings *BindingsCaller) GetMaxValidatorsPerReport(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getMaxValidatorsPerReport")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxValidatorsPerReport is a free data retrieval call binding the contract method 0xc1f665bc.
//
// Solidity: function getMaxValidatorsPerReport() view returns(uint256)
func (_Bindings *BindingsSession) GetMaxValidatorsPerReport() (*big.Int, error) {
	return _Bindings.Contract.GetMaxValidatorsPerReport(&_Bindings.CallOpts)
}

// GetMaxValidatorsPerReport is a free data retrieval call binding the contract method 0xc1f665bc.
//
// Solidity: function getMaxValidatorsPerReport() view returns(uint256)
func (_Bindings *BindingsCallerSession) GetMaxValidatorsPerReport() (*big.Int, error) {
	return _Bindings.Contract.GetMaxValidatorsPerReport(&_Bindings.CallOpts)
}

// GetProcessingState is a free data retrieval call binding the contract method 0x8f7797c2.
//
// Solidity: function getProcessingState() view returns((uint256,uint256,bytes32,bool,uint256,uint256,uint256) result)
func (_Bindings *BindingsCaller) GetProcessingState(opts *bind.CallOpts) (ValidatorsExitBusOracleProcessingState, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getProcessingState")

	if err != nil {
		return *new(ValidatorsExitBusOracleProcessingState), err
	}

	out0 := *abi.ConvertType(out[0], new(ValidatorsExitBusOracleProcessingState)).(*ValidatorsExitBusOracleProcessingState)

	return out0, err

}

// GetProcessingState is a free data retrieval call binding the contract method 0x8f7797c2.
//
// Solidity: function getProcessingState() view returns((uint256,uint256,bytes32,bool,uint256,uint256,uint256) result)
func (_Bindings *BindingsSession) GetProcessingState() (ValidatorsExitBusOracleProcessingState, error) {
	return _Bindings.Contract.GetProcessingState(&_Bindings.CallOpts)
}

// GetProcessingState is a free data retrieval call binding the contract method 0x8f7797c2.
//
// Solidity: function getProcessingState() view returns((uint256,uint256,bytes32,bool,uint256,uint256,uint256) result)
func (_Bindings *BindingsCallerSession) GetProcessingState() (ValidatorsExitBusOracleProcessingState, error) {
	return _Bindings.Contract.GetProcessingState(&_Bindings.CallOpts)
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

// GetTotalRequestsProcessed is a free data retrieval call binding the contract method 0xe2793e72.
//
// Solidity: function getTotalRequestsProcessed() view returns(uint256)
func (_Bindings *BindingsCaller) GetTotalRequestsProcessed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getTotalRequestsProcessed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalRequestsProcessed is a free data retrieval call binding the contract method 0xe2793e72.
//
// Solidity: function getTotalRequestsProcessed() view returns(uint256)
func (_Bindings *BindingsSession) GetTotalRequestsProcessed() (*big.Int, error) {
	return _Bindings.Contract.GetTotalRequestsProcessed(&_Bindings.CallOpts)
}

// GetTotalRequestsProcessed is a free data retrieval call binding the contract method 0xe2793e72.
//
// Solidity: function getTotalRequestsProcessed() view returns(uint256)
func (_Bindings *BindingsCallerSession) GetTotalRequestsProcessed() (*big.Int, error) {
	return _Bindings.Contract.GetTotalRequestsProcessed(&_Bindings.CallOpts)
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

// UnpackExitRequest is a free data retrieval call binding the contract method 0x7dad759d.
//
// Solidity: function unpackExitRequest(bytes exitRequests, uint256 dataFormat, uint256 index) pure returns(bytes pubkey, uint256 nodeOpId, uint256 moduleId, uint256 valIndex)
func (_Bindings *BindingsCaller) UnpackExitRequest(opts *bind.CallOpts, exitRequests []byte, dataFormat *big.Int, index *big.Int) (struct {
	Pubkey   []byte
	NodeOpId *big.Int
	ModuleId *big.Int
	ValIndex *big.Int
}, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "unpackExitRequest", exitRequests, dataFormat, index)

	outstruct := new(struct {
		Pubkey   []byte
		NodeOpId *big.Int
		ModuleId *big.Int
		ValIndex *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Pubkey = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.NodeOpId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ModuleId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ValIndex = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UnpackExitRequest is a free data retrieval call binding the contract method 0x7dad759d.
//
// Solidity: function unpackExitRequest(bytes exitRequests, uint256 dataFormat, uint256 index) pure returns(bytes pubkey, uint256 nodeOpId, uint256 moduleId, uint256 valIndex)
func (_Bindings *BindingsSession) UnpackExitRequest(exitRequests []byte, dataFormat *big.Int, index *big.Int) (struct {
	Pubkey   []byte
	NodeOpId *big.Int
	ModuleId *big.Int
	ValIndex *big.Int
}, error) {
	return _Bindings.Contract.UnpackExitRequest(&_Bindings.CallOpts, exitRequests, dataFormat, index)
}

// UnpackExitRequest is a free data retrieval call binding the contract method 0x7dad759d.
//
// Solidity: function unpackExitRequest(bytes exitRequests, uint256 dataFormat, uint256 index) pure returns(bytes pubkey, uint256 nodeOpId, uint256 moduleId, uint256 valIndex)
func (_Bindings *BindingsCallerSession) UnpackExitRequest(exitRequests []byte, dataFormat *big.Int, index *big.Int) (struct {
	Pubkey   []byte
	NodeOpId *big.Int
	ModuleId *big.Int
	ValIndex *big.Int
}, error) {
	return _Bindings.Contract.UnpackExitRequest(&_Bindings.CallOpts, exitRequests, dataFormat, index)
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

// FinalizeUpgradeV2 is a paid mutator transaction binding the contract method 0xff406d19.
//
// Solidity: function finalizeUpgrade_v2(uint256 maxValidatorsPerReport, uint256 maxExitRequestsLimit, uint256 exitsPerFrame, uint256 frameDurationInSec) returns()
func (_Bindings *BindingsTransactor) FinalizeUpgradeV2(opts *bind.TransactOpts, maxValidatorsPerReport *big.Int, maxExitRequestsLimit *big.Int, exitsPerFrame *big.Int, frameDurationInSec *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "finalizeUpgrade_v2", maxValidatorsPerReport, maxExitRequestsLimit, exitsPerFrame, frameDurationInSec)
}

// FinalizeUpgradeV2 is a paid mutator transaction binding the contract method 0xff406d19.
//
// Solidity: function finalizeUpgrade_v2(uint256 maxValidatorsPerReport, uint256 maxExitRequestsLimit, uint256 exitsPerFrame, uint256 frameDurationInSec) returns()
func (_Bindings *BindingsSession) FinalizeUpgradeV2(maxValidatorsPerReport *big.Int, maxExitRequestsLimit *big.Int, exitsPerFrame *big.Int, frameDurationInSec *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.FinalizeUpgradeV2(&_Bindings.TransactOpts, maxValidatorsPerReport, maxExitRequestsLimit, exitsPerFrame, frameDurationInSec)
}

// FinalizeUpgradeV2 is a paid mutator transaction binding the contract method 0xff406d19.
//
// Solidity: function finalizeUpgrade_v2(uint256 maxValidatorsPerReport, uint256 maxExitRequestsLimit, uint256 exitsPerFrame, uint256 frameDurationInSec) returns()
func (_Bindings *BindingsTransactorSession) FinalizeUpgradeV2(maxValidatorsPerReport *big.Int, maxExitRequestsLimit *big.Int, exitsPerFrame *big.Int, frameDurationInSec *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.FinalizeUpgradeV2(&_Bindings.TransactOpts, maxValidatorsPerReport, maxExitRequestsLimit, exitsPerFrame, frameDurationInSec)
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

// Initialize is a paid mutator transaction binding the contract method 0x8ba796af.
//
// Solidity: function initialize(address admin, address consensusContract, uint256 consensusVersion, uint256 lastProcessingRefSlot, uint256 maxValidatorsPerRequest, uint256 maxExitRequestsLimit, uint256 exitsPerFrame, uint256 frameDurationInSec) returns()
func (_Bindings *BindingsTransactor) Initialize(opts *bind.TransactOpts, admin common.Address, consensusContract common.Address, consensusVersion *big.Int, lastProcessingRefSlot *big.Int, maxValidatorsPerRequest *big.Int, maxExitRequestsLimit *big.Int, exitsPerFrame *big.Int, frameDurationInSec *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "initialize", admin, consensusContract, consensusVersion, lastProcessingRefSlot, maxValidatorsPerRequest, maxExitRequestsLimit, exitsPerFrame, frameDurationInSec)
}

// Initialize is a paid mutator transaction binding the contract method 0x8ba796af.
//
// Solidity: function initialize(address admin, address consensusContract, uint256 consensusVersion, uint256 lastProcessingRefSlot, uint256 maxValidatorsPerRequest, uint256 maxExitRequestsLimit, uint256 exitsPerFrame, uint256 frameDurationInSec) returns()
func (_Bindings *BindingsSession) Initialize(admin common.Address, consensusContract common.Address, consensusVersion *big.Int, lastProcessingRefSlot *big.Int, maxValidatorsPerRequest *big.Int, maxExitRequestsLimit *big.Int, exitsPerFrame *big.Int, frameDurationInSec *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Initialize(&_Bindings.TransactOpts, admin, consensusContract, consensusVersion, lastProcessingRefSlot, maxValidatorsPerRequest, maxExitRequestsLimit, exitsPerFrame, frameDurationInSec)
}

// Initialize is a paid mutator transaction binding the contract method 0x8ba796af.
//
// Solidity: function initialize(address admin, address consensusContract, uint256 consensusVersion, uint256 lastProcessingRefSlot, uint256 maxValidatorsPerRequest, uint256 maxExitRequestsLimit, uint256 exitsPerFrame, uint256 frameDurationInSec) returns()
func (_Bindings *BindingsTransactorSession) Initialize(admin common.Address, consensusContract common.Address, consensusVersion *big.Int, lastProcessingRefSlot *big.Int, maxValidatorsPerRequest *big.Int, maxExitRequestsLimit *big.Int, exitsPerFrame *big.Int, frameDurationInSec *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Initialize(&_Bindings.TransactOpts, admin, consensusContract, consensusVersion, lastProcessingRefSlot, maxValidatorsPerRequest, maxExitRequestsLimit, exitsPerFrame, frameDurationInSec)
}

// PauseFor is a paid mutator transaction binding the contract method 0xf3f449c7.
//
// Solidity: function pauseFor(uint256 _duration) returns()
func (_Bindings *BindingsTransactor) PauseFor(opts *bind.TransactOpts, _duration *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "pauseFor", _duration)
}

// PauseFor is a paid mutator transaction binding the contract method 0xf3f449c7.
//
// Solidity: function pauseFor(uint256 _duration) returns()
func (_Bindings *BindingsSession) PauseFor(_duration *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.PauseFor(&_Bindings.TransactOpts, _duration)
}

// PauseFor is a paid mutator transaction binding the contract method 0xf3f449c7.
//
// Solidity: function pauseFor(uint256 _duration) returns()
func (_Bindings *BindingsTransactorSession) PauseFor(_duration *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.PauseFor(&_Bindings.TransactOpts, _duration)
}

// PauseUntil is a paid mutator transaction binding the contract method 0xabe9cfc8.
//
// Solidity: function pauseUntil(uint256 _pauseUntilInclusive) returns()
func (_Bindings *BindingsTransactor) PauseUntil(opts *bind.TransactOpts, _pauseUntilInclusive *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "pauseUntil", _pauseUntilInclusive)
}

// PauseUntil is a paid mutator transaction binding the contract method 0xabe9cfc8.
//
// Solidity: function pauseUntil(uint256 _pauseUntilInclusive) returns()
func (_Bindings *BindingsSession) PauseUntil(_pauseUntilInclusive *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.PauseUntil(&_Bindings.TransactOpts, _pauseUntilInclusive)
}

// PauseUntil is a paid mutator transaction binding the contract method 0xabe9cfc8.
//
// Solidity: function pauseUntil(uint256 _pauseUntilInclusive) returns()
func (_Bindings *BindingsTransactorSession) PauseUntil(_pauseUntilInclusive *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.PauseUntil(&_Bindings.TransactOpts, _pauseUntilInclusive)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Bindings *BindingsTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Bindings *BindingsSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.RenounceRole(&_Bindings.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Bindings *BindingsTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.RenounceRole(&_Bindings.TransactOpts, role, account)
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

// SetExitRequestLimit is a paid mutator transaction binding the contract method 0x56254a97.
//
// Solidity: function setExitRequestLimit(uint256 maxExitRequestsLimit, uint256 exitsPerFrame, uint256 frameDurationInSec) returns()
func (_Bindings *BindingsTransactor) SetExitRequestLimit(opts *bind.TransactOpts, maxExitRequestsLimit *big.Int, exitsPerFrame *big.Int, frameDurationInSec *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setExitRequestLimit", maxExitRequestsLimit, exitsPerFrame, frameDurationInSec)
}

// SetExitRequestLimit is a paid mutator transaction binding the contract method 0x56254a97.
//
// Solidity: function setExitRequestLimit(uint256 maxExitRequestsLimit, uint256 exitsPerFrame, uint256 frameDurationInSec) returns()
func (_Bindings *BindingsSession) SetExitRequestLimit(maxExitRequestsLimit *big.Int, exitsPerFrame *big.Int, frameDurationInSec *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetExitRequestLimit(&_Bindings.TransactOpts, maxExitRequestsLimit, exitsPerFrame, frameDurationInSec)
}

// SetExitRequestLimit is a paid mutator transaction binding the contract method 0x56254a97.
//
// Solidity: function setExitRequestLimit(uint256 maxExitRequestsLimit, uint256 exitsPerFrame, uint256 frameDurationInSec) returns()
func (_Bindings *BindingsTransactorSession) SetExitRequestLimit(maxExitRequestsLimit *big.Int, exitsPerFrame *big.Int, frameDurationInSec *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetExitRequestLimit(&_Bindings.TransactOpts, maxExitRequestsLimit, exitsPerFrame, frameDurationInSec)
}

// SetMaxValidatorsPerReport is a paid mutator transaction binding the contract method 0x6f2c322d.
//
// Solidity: function setMaxValidatorsPerReport(uint256 maxRequests) returns()
func (_Bindings *BindingsTransactor) SetMaxValidatorsPerReport(opts *bind.TransactOpts, maxRequests *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setMaxValidatorsPerReport", maxRequests)
}

// SetMaxValidatorsPerReport is a paid mutator transaction binding the contract method 0x6f2c322d.
//
// Solidity: function setMaxValidatorsPerReport(uint256 maxRequests) returns()
func (_Bindings *BindingsSession) SetMaxValidatorsPerReport(maxRequests *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetMaxValidatorsPerReport(&_Bindings.TransactOpts, maxRequests)
}

// SetMaxValidatorsPerReport is a paid mutator transaction binding the contract method 0x6f2c322d.
//
// Solidity: function setMaxValidatorsPerReport(uint256 maxRequests) returns()
func (_Bindings *BindingsTransactorSession) SetMaxValidatorsPerReport(maxRequests *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetMaxValidatorsPerReport(&_Bindings.TransactOpts, maxRequests)
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

// SubmitExitRequestsData is a paid mutator transaction binding the contract method 0xb8fe0ad0.
//
// Solidity: function submitExitRequestsData((bytes,uint256) request) returns()
func (_Bindings *BindingsTransactor) SubmitExitRequestsData(opts *bind.TransactOpts, request ValidatorsExitBusExitRequestsData) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "submitExitRequestsData", request)
}

// SubmitExitRequestsData is a paid mutator transaction binding the contract method 0xb8fe0ad0.
//
// Solidity: function submitExitRequestsData((bytes,uint256) request) returns()
func (_Bindings *BindingsSession) SubmitExitRequestsData(request ValidatorsExitBusExitRequestsData) (*types.Transaction, error) {
	return _Bindings.Contract.SubmitExitRequestsData(&_Bindings.TransactOpts, request)
}

// SubmitExitRequestsData is a paid mutator transaction binding the contract method 0xb8fe0ad0.
//
// Solidity: function submitExitRequestsData((bytes,uint256) request) returns()
func (_Bindings *BindingsTransactorSession) SubmitExitRequestsData(request ValidatorsExitBusExitRequestsData) (*types.Transaction, error) {
	return _Bindings.Contract.SubmitExitRequestsData(&_Bindings.TransactOpts, request)
}

// SubmitExitRequestsHash is a paid mutator transaction binding the contract method 0xb1b19f57.
//
// Solidity: function submitExitRequestsHash(bytes32 exitRequestsHash) returns()
func (_Bindings *BindingsTransactor) SubmitExitRequestsHash(opts *bind.TransactOpts, exitRequestsHash [32]byte) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "submitExitRequestsHash", exitRequestsHash)
}

// SubmitExitRequestsHash is a paid mutator transaction binding the contract method 0xb1b19f57.
//
// Solidity: function submitExitRequestsHash(bytes32 exitRequestsHash) returns()
func (_Bindings *BindingsSession) SubmitExitRequestsHash(exitRequestsHash [32]byte) (*types.Transaction, error) {
	return _Bindings.Contract.SubmitExitRequestsHash(&_Bindings.TransactOpts, exitRequestsHash)
}

// SubmitExitRequestsHash is a paid mutator transaction binding the contract method 0xb1b19f57.
//
// Solidity: function submitExitRequestsHash(bytes32 exitRequestsHash) returns()
func (_Bindings *BindingsTransactorSession) SubmitExitRequestsHash(exitRequestsHash [32]byte) (*types.Transaction, error) {
	return _Bindings.Contract.SubmitExitRequestsHash(&_Bindings.TransactOpts, exitRequestsHash)
}

// SubmitReportData is a paid mutator transaction binding the contract method 0x294492c8.
//
// Solidity: function submitReportData((uint256,uint256,uint256,uint256,bytes) data, uint256 contractVersion) returns()
func (_Bindings *BindingsTransactor) SubmitReportData(opts *bind.TransactOpts, data ValidatorsExitBusOracleReportData, contractVersion *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "submitReportData", data, contractVersion)
}

// SubmitReportData is a paid mutator transaction binding the contract method 0x294492c8.
//
// Solidity: function submitReportData((uint256,uint256,uint256,uint256,bytes) data, uint256 contractVersion) returns()
func (_Bindings *BindingsSession) SubmitReportData(data ValidatorsExitBusOracleReportData, contractVersion *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SubmitReportData(&_Bindings.TransactOpts, data, contractVersion)
}

// SubmitReportData is a paid mutator transaction binding the contract method 0x294492c8.
//
// Solidity: function submitReportData((uint256,uint256,uint256,uint256,bytes) data, uint256 contractVersion) returns()
func (_Bindings *BindingsTransactorSession) SubmitReportData(data ValidatorsExitBusOracleReportData, contractVersion *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SubmitReportData(&_Bindings.TransactOpts, data, contractVersion)
}

// TriggerExits is a paid mutator transaction binding the contract method 0xa2ab7065.
//
// Solidity: function triggerExits((bytes,uint256) exitsData, uint256[] exitDataIndexes, address refundRecipient) payable returns()
func (_Bindings *BindingsTransactor) TriggerExits(opts *bind.TransactOpts, exitsData ValidatorsExitBusExitRequestsData, exitDataIndexes []*big.Int, refundRecipient common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "triggerExits", exitsData, exitDataIndexes, refundRecipient)
}

// TriggerExits is a paid mutator transaction binding the contract method 0xa2ab7065.
//
// Solidity: function triggerExits((bytes,uint256) exitsData, uint256[] exitDataIndexes, address refundRecipient) payable returns()
func (_Bindings *BindingsSession) TriggerExits(exitsData ValidatorsExitBusExitRequestsData, exitDataIndexes []*big.Int, refundRecipient common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.TriggerExits(&_Bindings.TransactOpts, exitsData, exitDataIndexes, refundRecipient)
}

// TriggerExits is a paid mutator transaction binding the contract method 0xa2ab7065.
//
// Solidity: function triggerExits((bytes,uint256) exitsData, uint256[] exitDataIndexes, address refundRecipient) payable returns()
func (_Bindings *BindingsTransactorSession) TriggerExits(exitsData ValidatorsExitBusExitRequestsData, exitDataIndexes []*big.Int, refundRecipient common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.TriggerExits(&_Bindings.TransactOpts, exitsData, exitDataIndexes, refundRecipient)
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

// BindingsExitDataProcessingIterator is returned from FilterExitDataProcessing and is used to iterate over the raw logs and unpacked data for ExitDataProcessing events raised by the Bindings contract.
type BindingsExitDataProcessingIterator struct {
	Event *BindingsExitDataProcessing // Event containing the contract specifics and raw log

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
func (it *BindingsExitDataProcessingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsExitDataProcessing)
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
		it.Event = new(BindingsExitDataProcessing)
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
func (it *BindingsExitDataProcessingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsExitDataProcessingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsExitDataProcessing represents a ExitDataProcessing event raised by the Bindings contract.
type BindingsExitDataProcessing struct {
	ExitRequestsHash [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterExitDataProcessing is a free log retrieval operation binding the contract event 0x01b8de053572c3c2104259b555c485ccac8017196b3471e8483b7e96f071608a.
//
// Solidity: event ExitDataProcessing(bytes32 exitRequestsHash)
func (_Bindings *BindingsFilterer) FilterExitDataProcessing(opts *bind.FilterOpts) (*BindingsExitDataProcessingIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "ExitDataProcessing")
	if err != nil {
		return nil, err
	}
	return &BindingsExitDataProcessingIterator{contract: _Bindings.contract, event: "ExitDataProcessing", logs: logs, sub: sub}, nil
}

// WatchExitDataProcessing is a free log subscription operation binding the contract event 0x01b8de053572c3c2104259b555c485ccac8017196b3471e8483b7e96f071608a.
//
// Solidity: event ExitDataProcessing(bytes32 exitRequestsHash)
func (_Bindings *BindingsFilterer) WatchExitDataProcessing(opts *bind.WatchOpts, sink chan<- *BindingsExitDataProcessing) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "ExitDataProcessing")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsExitDataProcessing)
				if err := _Bindings.contract.UnpackLog(event, "ExitDataProcessing", log); err != nil {
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

// ParseExitDataProcessing is a log parse operation binding the contract event 0x01b8de053572c3c2104259b555c485ccac8017196b3471e8483b7e96f071608a.
//
// Solidity: event ExitDataProcessing(bytes32 exitRequestsHash)
func (_Bindings *BindingsFilterer) ParseExitDataProcessing(log types.Log) (*BindingsExitDataProcessing, error) {
	event := new(BindingsExitDataProcessing)
	if err := _Bindings.contract.UnpackLog(event, "ExitDataProcessing", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsExitRequestsLimitSetIterator is returned from FilterExitRequestsLimitSet and is used to iterate over the raw logs and unpacked data for ExitRequestsLimitSet events raised by the Bindings contract.
type BindingsExitRequestsLimitSetIterator struct {
	Event *BindingsExitRequestsLimitSet // Event containing the contract specifics and raw log

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
func (it *BindingsExitRequestsLimitSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsExitRequestsLimitSet)
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
		it.Event = new(BindingsExitRequestsLimitSet)
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
func (it *BindingsExitRequestsLimitSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsExitRequestsLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsExitRequestsLimitSet represents a ExitRequestsLimitSet event raised by the Bindings contract.
type BindingsExitRequestsLimitSet struct {
	MaxExitRequestsLimit *big.Int
	ExitsPerFrame        *big.Int
	FrameDurationInSec   *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterExitRequestsLimitSet is a free log retrieval operation binding the contract event 0x3119d910326e0f179e121df55f23f45b8a5022ff10c73c02aabf2b48ae36070a.
//
// Solidity: event ExitRequestsLimitSet(uint256 maxExitRequestsLimit, uint256 exitsPerFrame, uint256 frameDurationInSec)
func (_Bindings *BindingsFilterer) FilterExitRequestsLimitSet(opts *bind.FilterOpts) (*BindingsExitRequestsLimitSetIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "ExitRequestsLimitSet")
	if err != nil {
		return nil, err
	}
	return &BindingsExitRequestsLimitSetIterator{contract: _Bindings.contract, event: "ExitRequestsLimitSet", logs: logs, sub: sub}, nil
}

// WatchExitRequestsLimitSet is a free log subscription operation binding the contract event 0x3119d910326e0f179e121df55f23f45b8a5022ff10c73c02aabf2b48ae36070a.
//
// Solidity: event ExitRequestsLimitSet(uint256 maxExitRequestsLimit, uint256 exitsPerFrame, uint256 frameDurationInSec)
func (_Bindings *BindingsFilterer) WatchExitRequestsLimitSet(opts *bind.WatchOpts, sink chan<- *BindingsExitRequestsLimitSet) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "ExitRequestsLimitSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsExitRequestsLimitSet)
				if err := _Bindings.contract.UnpackLog(event, "ExitRequestsLimitSet", log); err != nil {
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

// ParseExitRequestsLimitSet is a log parse operation binding the contract event 0x3119d910326e0f179e121df55f23f45b8a5022ff10c73c02aabf2b48ae36070a.
//
// Solidity: event ExitRequestsLimitSet(uint256 maxExitRequestsLimit, uint256 exitsPerFrame, uint256 frameDurationInSec)
func (_Bindings *BindingsFilterer) ParseExitRequestsLimitSet(log types.Log) (*BindingsExitRequestsLimitSet, error) {
	event := new(BindingsExitRequestsLimitSet)
	if err := _Bindings.contract.UnpackLog(event, "ExitRequestsLimitSet", log); err != nil {
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

// BindingsProcessingStartedIterator is returned from FilterProcessingStarted and is used to iterate over the raw logs and unpacked data for ProcessingStarted events raised by the Bindings contract.
type BindingsProcessingStartedIterator struct {
	Event *BindingsProcessingStarted // Event containing the contract specifics and raw log

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
			it.Event = new(BindingsProcessingStarted)
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
		it.Event = new(BindingsProcessingStarted)
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
type BindingsProcessingStarted struct {
	RefSlot *big.Int
	Hash    [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

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
func (_Bindings *BindingsFilterer) WatchProcessingStarted(opts *bind.WatchOpts, sink chan<- *BindingsProcessingStarted, refSlot []*big.Int) (event.Subscription, error) {

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
				event := new(BindingsProcessingStarted)
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
func (_Bindings *BindingsFilterer) ParseProcessingStarted(log types.Log) (*BindingsProcessingStarted, error) {
	event := new(BindingsProcessingStarted)
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

// BindingsRequestsHashSubmittedIterator is returned from FilterRequestsHashSubmitted and is used to iterate over the raw logs and unpacked data for RequestsHashSubmitted events raised by the Bindings contract.
type BindingsRequestsHashSubmittedIterator struct {
	Event *BindingsRequestsHashSubmitted // Event containing the contract specifics and raw log

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
func (it *BindingsRequestsHashSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsRequestsHashSubmitted)
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
		it.Event = new(BindingsRequestsHashSubmitted)
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
func (it *BindingsRequestsHashSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsRequestsHashSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsRequestsHashSubmitted represents a RequestsHashSubmitted event raised by the Bindings contract.
type BindingsRequestsHashSubmitted struct {
	ExitRequestsHash [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRequestsHashSubmitted is a free log retrieval operation binding the contract event 0x76d8359ea28964b79f7fa8bb502ec325fd0d1e956c42a0436940e35d0e99f2de.
//
// Solidity: event RequestsHashSubmitted(bytes32 exitRequestsHash)
func (_Bindings *BindingsFilterer) FilterRequestsHashSubmitted(opts *bind.FilterOpts) (*BindingsRequestsHashSubmittedIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "RequestsHashSubmitted")
	if err != nil {
		return nil, err
	}
	return &BindingsRequestsHashSubmittedIterator{contract: _Bindings.contract, event: "RequestsHashSubmitted", logs: logs, sub: sub}, nil
}

// WatchRequestsHashSubmitted is a free log subscription operation binding the contract event 0x76d8359ea28964b79f7fa8bb502ec325fd0d1e956c42a0436940e35d0e99f2de.
//
// Solidity: event RequestsHashSubmitted(bytes32 exitRequestsHash)
func (_Bindings *BindingsFilterer) WatchRequestsHashSubmitted(opts *bind.WatchOpts, sink chan<- *BindingsRequestsHashSubmitted) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "RequestsHashSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsRequestsHashSubmitted)
				if err := _Bindings.contract.UnpackLog(event, "RequestsHashSubmitted", log); err != nil {
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

// ParseRequestsHashSubmitted is a log parse operation binding the contract event 0x76d8359ea28964b79f7fa8bb502ec325fd0d1e956c42a0436940e35d0e99f2de.
//
// Solidity: event RequestsHashSubmitted(bytes32 exitRequestsHash)
func (_Bindings *BindingsFilterer) ParseRequestsHashSubmitted(log types.Log) (*BindingsRequestsHashSubmitted, error) {
	event := new(BindingsRequestsHashSubmitted)
	if err := _Bindings.contract.UnpackLog(event, "RequestsHashSubmitted", log); err != nil {
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

// BindingsSetMaxValidatorsPerReportIterator is returned from FilterSetMaxValidatorsPerReport and is used to iterate over the raw logs and unpacked data for SetMaxValidatorsPerReport events raised by the Bindings contract.
type BindingsSetMaxValidatorsPerReportIterator struct {
	Event *BindingsSetMaxValidatorsPerReport // Event containing the contract specifics and raw log

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
func (it *BindingsSetMaxValidatorsPerReportIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsSetMaxValidatorsPerReport)
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
		it.Event = new(BindingsSetMaxValidatorsPerReport)
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
func (it *BindingsSetMaxValidatorsPerReportIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsSetMaxValidatorsPerReportIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsSetMaxValidatorsPerReport represents a SetMaxValidatorsPerReport event raised by the Bindings contract.
type BindingsSetMaxValidatorsPerReport struct {
	MaxValidatorsPerReport *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetMaxValidatorsPerReport is a free log retrieval operation binding the contract event 0x9b17a153b6e933d8497c6b713fbd70c893891d75639ede17ce6e4cea08e7cfc3.
//
// Solidity: event SetMaxValidatorsPerReport(uint256 maxValidatorsPerReport)
func (_Bindings *BindingsFilterer) FilterSetMaxValidatorsPerReport(opts *bind.FilterOpts) (*BindingsSetMaxValidatorsPerReportIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "SetMaxValidatorsPerReport")
	if err != nil {
		return nil, err
	}
	return &BindingsSetMaxValidatorsPerReportIterator{contract: _Bindings.contract, event: "SetMaxValidatorsPerReport", logs: logs, sub: sub}, nil
}

// WatchSetMaxValidatorsPerReport is a free log subscription operation binding the contract event 0x9b17a153b6e933d8497c6b713fbd70c893891d75639ede17ce6e4cea08e7cfc3.
//
// Solidity: event SetMaxValidatorsPerReport(uint256 maxValidatorsPerReport)
func (_Bindings *BindingsFilterer) WatchSetMaxValidatorsPerReport(opts *bind.WatchOpts, sink chan<- *BindingsSetMaxValidatorsPerReport) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "SetMaxValidatorsPerReport")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsSetMaxValidatorsPerReport)
				if err := _Bindings.contract.UnpackLog(event, "SetMaxValidatorsPerReport", log); err != nil {
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

// ParseSetMaxValidatorsPerReport is a log parse operation binding the contract event 0x9b17a153b6e933d8497c6b713fbd70c893891d75639ede17ce6e4cea08e7cfc3.
//
// Solidity: event SetMaxValidatorsPerReport(uint256 maxValidatorsPerReport)
func (_Bindings *BindingsFilterer) ParseSetMaxValidatorsPerReport(log types.Log) (*BindingsSetMaxValidatorsPerReport, error) {
	event := new(BindingsSetMaxValidatorsPerReport)
	if err := _Bindings.contract.UnpackLog(event, "SetMaxValidatorsPerReport", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsValidatorExitRequestIterator is returned from FilterValidatorExitRequest and is used to iterate over the raw logs and unpacked data for ValidatorExitRequest events raised by the Bindings contract.
type BindingsValidatorExitRequestIterator struct {
	Event *domain.VeboValidatorExitRequest // Event containing the contract specifics and raw log

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
func (it *BindingsValidatorExitRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(domain.VeboValidatorExitRequest)
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
		it.Event = new(domain.VeboValidatorExitRequest)
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
func (it *BindingsValidatorExitRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsValidatorExitRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsValidatorExitRequest represents a ValidatorExitRequest event raised by the Bindings contract.
// type BindingsValidatorExitRequest struct {
// 	StakingModuleId *big.Int
// 	NodeOperatorId  *big.Int
// 	ValidatorIndex  *big.Int
// 	ValidatorPubkey []byte
// 	Timestamp       *big.Int
// 	Raw             types.Log // Blockchain specific contextual infos
// }

// FilterValidatorExitRequest is a free log retrieval operation binding the contract event 0x96395f55c4997466e5035d777f0e1ba82b8cae217aaad05cf07839eb7c75bcf2.
//
// Solidity: event ValidatorExitRequest(uint256 indexed stakingModuleId, uint256 indexed nodeOperatorId, uint256 indexed validatorIndex, bytes validatorPubkey, uint256 timestamp)
func (_Bindings *BindingsFilterer) FilterValidatorExitRequest(opts *bind.FilterOpts, stakingModuleId []*big.Int, nodeOperatorId []*big.Int, validatorIndex []*big.Int) (*BindingsValidatorExitRequestIterator, error) {

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

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "ValidatorExitRequest", stakingModuleIdRule, nodeOperatorIdRule, validatorIndexRule)
	if err != nil {
		return nil, err
	}
	return &BindingsValidatorExitRequestIterator{contract: _Bindings.contract, event: "ValidatorExitRequest", logs: logs, sub: sub}, nil
}

// WatchValidatorExitRequest is a free log subscription operation binding the contract event 0x96395f55c4997466e5035d777f0e1ba82b8cae217aaad05cf07839eb7c75bcf2.
//
// Solidity: event ValidatorExitRequest(uint256 indexed stakingModuleId, uint256 indexed nodeOperatorId, uint256 indexed validatorIndex, bytes validatorPubkey, uint256 timestamp)
func (_Bindings *BindingsFilterer) WatchValidatorExitRequest(opts *bind.WatchOpts, sink chan<- *domain.VeboValidatorExitRequest, stakingModuleId []*big.Int, nodeOperatorId []*big.Int, validatorIndex []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "ValidatorExitRequest", stakingModuleIdRule, nodeOperatorIdRule, validatorIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(domain.VeboValidatorExitRequest)
				if err := _Bindings.contract.UnpackLog(event, "ValidatorExitRequest", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParseValidatorExitRequest(log types.Log) (*domain.VeboValidatorExitRequest, error) {
	event := new(domain.VeboValidatorExitRequest)
	if err := _Bindings.contract.UnpackLog(event, "ValidatorExitRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsWarnDataIncompleteProcessingIterator is returned from FilterWarnDataIncompleteProcessing and is used to iterate over the raw logs and unpacked data for WarnDataIncompleteProcessing events raised by the Bindings contract.
type BindingsWarnDataIncompleteProcessingIterator struct {
	Event *BindingsWarnDataIncompleteProcessing // Event containing the contract specifics and raw log

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
func (it *BindingsWarnDataIncompleteProcessingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsWarnDataIncompleteProcessing)
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
		it.Event = new(BindingsWarnDataIncompleteProcessing)
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
func (it *BindingsWarnDataIncompleteProcessingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsWarnDataIncompleteProcessingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsWarnDataIncompleteProcessing represents a WarnDataIncompleteProcessing event raised by the Bindings contract.
type BindingsWarnDataIncompleteProcessing struct {
	RefSlot           *big.Int
	RequestsProcessed *big.Int
	RequestsCount     *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterWarnDataIncompleteProcessing is a free log retrieval operation binding the contract event 0xefc67aab43195093a8d8ed25d52281d96de480748ece2787888c586e8e1e79b4.
//
// Solidity: event WarnDataIncompleteProcessing(uint256 indexed refSlot, uint256 requestsProcessed, uint256 requestsCount)
func (_Bindings *BindingsFilterer) FilterWarnDataIncompleteProcessing(opts *bind.FilterOpts, refSlot []*big.Int) (*BindingsWarnDataIncompleteProcessingIterator, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "WarnDataIncompleteProcessing", refSlotRule)
	if err != nil {
		return nil, err
	}
	return &BindingsWarnDataIncompleteProcessingIterator{contract: _Bindings.contract, event: "WarnDataIncompleteProcessing", logs: logs, sub: sub}, nil
}

// WatchWarnDataIncompleteProcessing is a free log subscription operation binding the contract event 0xefc67aab43195093a8d8ed25d52281d96de480748ece2787888c586e8e1e79b4.
//
// Solidity: event WarnDataIncompleteProcessing(uint256 indexed refSlot, uint256 requestsProcessed, uint256 requestsCount)
func (_Bindings *BindingsFilterer) WatchWarnDataIncompleteProcessing(opts *bind.WatchOpts, sink chan<- *BindingsWarnDataIncompleteProcessing, refSlot []*big.Int) (event.Subscription, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "WarnDataIncompleteProcessing", refSlotRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsWarnDataIncompleteProcessing)
				if err := _Bindings.contract.UnpackLog(event, "WarnDataIncompleteProcessing", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParseWarnDataIncompleteProcessing(log types.Log) (*BindingsWarnDataIncompleteProcessing, error) {
	event := new(BindingsWarnDataIncompleteProcessing)
	if err := _Bindings.contract.UnpackLog(event, "WarnDataIncompleteProcessing", log); err != nil {
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
