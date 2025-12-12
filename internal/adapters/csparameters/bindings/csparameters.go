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

// ICSParametersRegistryInitializationData is an auto generated low-level Go binding around an user-defined struct.
type ICSParametersRegistryInitializationData struct {
	KeyRemovalCharge                *big.Int
	ElRewardsStealingAdditionalFine *big.Int
	KeysLimit                       *big.Int
	RewardShare                     *big.Int
	PerformanceLeeway               *big.Int
	StrikesLifetime                 *big.Int
	StrikesThreshold                *big.Int
	DefaultQueuePriority            *big.Int
	DefaultQueueMaxDeposits         *big.Int
	BadPerformancePenalty           *big.Int
	AttestationsWeight              *big.Int
	BlocksWeight                    *big.Int
	SyncWeight                      *big.Int
	DefaultAllowedExitDelay         *big.Int
	DefaultExitDelayPenalty         *big.Int
	DefaultMaxWithdrawalRequestFee  *big.Int
}

// ICSParametersRegistryKeyNumberValueInterval is an auto generated low-level Go binding around an user-defined struct.
type ICSParametersRegistryKeyNumberValueInterval struct {
	MinKeyNumber *big.Int
	Value        *big.Int
}

// BindingsMetaData contains all meta data concerning the Bindings contract.
var BindingsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"queueLowestPriority\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAllowedExitDelay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyNumberValueIntervals\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPerformanceCoefficients\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPerformanceLeewayData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRewardShareData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStrikesParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QueueCannotBeUsed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAdminAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroMaxDeposits\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroQueueLowestPriority\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"AllowedExitDelaySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"AllowedExitDelayUnset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"penalty\",\"type\":\"uint256\"}],\"name\":\"BadPerformancePenaltySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"BadPerformancePenaltyUnset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"DefaultAllowedExitDelaySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"DefaultBadPerformancePenaltySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"DefaultElRewardsStealingAdditionalFineSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"penalty\",\"type\":\"uint256\"}],\"name\":\"DefaultExitDelayPenaltySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"DefaultKeyRemovalChargeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"DefaultKeysLimitSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"DefaultMaxWithdrawalRequestFeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"attestationsWeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blocksWeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"syncWeight\",\"type\":\"uint256\"}],\"name\":\"DefaultPerformanceCoefficientsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"DefaultPerformanceLeewaySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"priority\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxDeposits\",\"type\":\"uint256\"}],\"name\":\"DefaultQueueConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"DefaultRewardShareSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lifetime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"DefaultStrikesParamsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fine\",\"type\":\"uint256\"}],\"name\":\"ElRewardsStealingAdditionalFineSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"ElRewardsStealingAdditionalFineUnset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"penalty\",\"type\":\"uint256\"}],\"name\":\"ExitDelayPenaltySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"ExitDelayPenaltyUnset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"keyRemovalCharge\",\"type\":\"uint256\"}],\"name\":\"KeyRemovalChargeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"KeyRemovalChargeUnset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"KeysLimitSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"KeysLimitUnset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"MaxWithdrawalRequestFeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"MaxWithdrawalRequestFeeUnset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"attestationsWeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blocksWeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"syncWeight\",\"type\":\"uint256\"}],\"name\":\"PerformanceCoefficientsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"PerformanceCoefficientsUnset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minKeyNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structICSParametersRegistry.KeyNumberValueInterval[]\",\"name\":\"data\",\"type\":\"tuple[]\"}],\"name\":\"PerformanceLeewayDataSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"PerformanceLeewayDataUnset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"priority\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxDeposits\",\"type\":\"uint256\"}],\"name\":\"QueueConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"QueueConfigUnset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minKeyNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structICSParametersRegistry.KeyNumberValueInterval[]\",\"name\":\"data\",\"type\":\"tuple[]\"}],\"name\":\"RewardShareDataSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"RewardShareDataUnset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lifetime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"StrikesParamsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"StrikesParamsUnset\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"QUEUE_LEGACY_PRIORITY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"QUEUE_LOWEST_PRIORITY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAllowedExitDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultBadPerformancePenalty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultElRewardsStealingAdditionalFine\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultExitDelayPenalty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultKeyRemovalCharge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultKeysLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultMaxWithdrawalRequestFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultPerformanceCoefficients\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"attestationsWeight\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blocksWeight\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"syncWeight\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultPerformanceLeeway\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultQueueConfig\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"priority\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxDeposits\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultRewardShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultStrikesParams\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"lifetime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"threshold\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"getAllowedExitDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"getBadPerformancePenalty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"penalty\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"getElRewardsStealingAdditionalFine\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fine\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"getExitDelayPenalty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"penalty\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInitializedVersion\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"getKeyRemovalCharge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"keyRemovalCharge\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"getKeysLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"getMaxWithdrawalRequestFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"getPerformanceCoefficients\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"attestationsWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blocksWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"syncWeight\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"getPerformanceLeewayData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minKeyNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structICSParametersRegistry.KeyNumberValueInterval[]\",\"name\":\"data\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"getQueueConfig\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"queuePriority\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxDeposits\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"getRewardShareData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minKeyNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structICSParametersRegistry.KeyNumberValueInterval[]\",\"name\":\"data\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"getStrikesParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lifetime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"keyRemovalCharge\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"elRewardsStealingAdditionalFine\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keysLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"performanceLeeway\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strikesLifetime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strikesThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"defaultQueuePriority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"defaultQueueMaxDeposits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"badPerformancePenalty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"attestationsWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blocksWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"syncWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"defaultAllowedExitDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"defaultExitDelayPenalty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"defaultMaxWithdrawalRequestFee\",\"type\":\"uint256\"}],\"internalType\":\"structICSParametersRegistry.InitializationData\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"setAllowedExitDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"penalty\",\"type\":\"uint256\"}],\"name\":\"setBadPerformancePenalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"setDefaultAllowedExitDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"penalty\",\"type\":\"uint256\"}],\"name\":\"setDefaultBadPerformancePenalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fine\",\"type\":\"uint256\"}],\"name\":\"setDefaultElRewardsStealingAdditionalFine\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"penalty\",\"type\":\"uint256\"}],\"name\":\"setDefaultExitDelayPenalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"keyRemovalCharge\",\"type\":\"uint256\"}],\"name\":\"setDefaultKeyRemovalCharge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"setDefaultKeysLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setDefaultMaxWithdrawalRequestFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"attestationsWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blocksWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"syncWeight\",\"type\":\"uint256\"}],\"name\":\"setDefaultPerformanceCoefficients\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"leeway\",\"type\":\"uint256\"}],\"name\":\"setDefaultPerformanceLeeway\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"priority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDeposits\",\"type\":\"uint256\"}],\"name\":\"setDefaultQueueConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"share\",\"type\":\"uint256\"}],\"name\":\"setDefaultRewardShare\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lifetime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"setDefaultStrikesParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fine\",\"type\":\"uint256\"}],\"name\":\"setElRewardsStealingAdditionalFine\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"penalty\",\"type\":\"uint256\"}],\"name\":\"setExitDelayPenalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keyRemovalCharge\",\"type\":\"uint256\"}],\"name\":\"setKeyRemovalCharge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"setKeysLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setMaxWithdrawalRequestFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"attestationsWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blocksWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"syncWeight\",\"type\":\"uint256\"}],\"name\":\"setPerformanceCoefficients\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minKeyNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structICSParametersRegistry.KeyNumberValueInterval[]\",\"name\":\"data\",\"type\":\"tuple[]\"}],\"name\":\"setPerformanceLeewayData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDeposits\",\"type\":\"uint256\"}],\"name\":\"setQueueConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minKeyNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structICSParametersRegistry.KeyNumberValueInterval[]\",\"name\":\"data\",\"type\":\"tuple[]\"}],\"name\":\"setRewardShareData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lifetime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"setStrikesParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"unsetAllowedExitDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"unsetBadPerformancePenalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"unsetElRewardsStealingAdditionalFine\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"unsetExitDelayPenalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"unsetKeyRemovalCharge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"unsetKeysLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"unsetMaxWithdrawalRequestFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"unsetPerformanceCoefficients\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"unsetPerformanceLeewayData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"unsetQueueConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"unsetRewardShareData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"unsetStrikesParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// QUEUELEGACYPRIORITY is a free data retrieval call binding the contract method 0xa6b89b81.
//
// Solidity: function QUEUE_LEGACY_PRIORITY() view returns(uint256)
func (_Bindings *BindingsCaller) QUEUELEGACYPRIORITY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "QUEUE_LEGACY_PRIORITY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QUEUELEGACYPRIORITY is a free data retrieval call binding the contract method 0xa6b89b81.
//
// Solidity: function QUEUE_LEGACY_PRIORITY() view returns(uint256)
func (_Bindings *BindingsSession) QUEUELEGACYPRIORITY() (*big.Int, error) {
	return _Bindings.Contract.QUEUELEGACYPRIORITY(&_Bindings.CallOpts)
}

// QUEUELEGACYPRIORITY is a free data retrieval call binding the contract method 0xa6b89b81.
//
// Solidity: function QUEUE_LEGACY_PRIORITY() view returns(uint256)
func (_Bindings *BindingsCallerSession) QUEUELEGACYPRIORITY() (*big.Int, error) {
	return _Bindings.Contract.QUEUELEGACYPRIORITY(&_Bindings.CallOpts)
}

// QUEUELOWESTPRIORITY is a free data retrieval call binding the contract method 0xd614ae0c.
//
// Solidity: function QUEUE_LOWEST_PRIORITY() view returns(uint256)
func (_Bindings *BindingsCaller) QUEUELOWESTPRIORITY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "QUEUE_LOWEST_PRIORITY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QUEUELOWESTPRIORITY is a free data retrieval call binding the contract method 0xd614ae0c.
//
// Solidity: function QUEUE_LOWEST_PRIORITY() view returns(uint256)
func (_Bindings *BindingsSession) QUEUELOWESTPRIORITY() (*big.Int, error) {
	return _Bindings.Contract.QUEUELOWESTPRIORITY(&_Bindings.CallOpts)
}

// QUEUELOWESTPRIORITY is a free data retrieval call binding the contract method 0xd614ae0c.
//
// Solidity: function QUEUE_LOWEST_PRIORITY() view returns(uint256)
func (_Bindings *BindingsCallerSession) QUEUELOWESTPRIORITY() (*big.Int, error) {
	return _Bindings.Contract.QUEUELOWESTPRIORITY(&_Bindings.CallOpts)
}

// DefaultAllowedExitDelay is a free data retrieval call binding the contract method 0x49a0fae7.
//
// Solidity: function defaultAllowedExitDelay() view returns(uint256)
func (_Bindings *BindingsCaller) DefaultAllowedExitDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "defaultAllowedExitDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultAllowedExitDelay is a free data retrieval call binding the contract method 0x49a0fae7.
//
// Solidity: function defaultAllowedExitDelay() view returns(uint256)
func (_Bindings *BindingsSession) DefaultAllowedExitDelay() (*big.Int, error) {
	return _Bindings.Contract.DefaultAllowedExitDelay(&_Bindings.CallOpts)
}

// DefaultAllowedExitDelay is a free data retrieval call binding the contract method 0x49a0fae7.
//
// Solidity: function defaultAllowedExitDelay() view returns(uint256)
func (_Bindings *BindingsCallerSession) DefaultAllowedExitDelay() (*big.Int, error) {
	return _Bindings.Contract.DefaultAllowedExitDelay(&_Bindings.CallOpts)
}

// DefaultBadPerformancePenalty is a free data retrieval call binding the contract method 0x3914285e.
//
// Solidity: function defaultBadPerformancePenalty() view returns(uint256)
func (_Bindings *BindingsCaller) DefaultBadPerformancePenalty(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "defaultBadPerformancePenalty")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultBadPerformancePenalty is a free data retrieval call binding the contract method 0x3914285e.
//
// Solidity: function defaultBadPerformancePenalty() view returns(uint256)
func (_Bindings *BindingsSession) DefaultBadPerformancePenalty() (*big.Int, error) {
	return _Bindings.Contract.DefaultBadPerformancePenalty(&_Bindings.CallOpts)
}

// DefaultBadPerformancePenalty is a free data retrieval call binding the contract method 0x3914285e.
//
// Solidity: function defaultBadPerformancePenalty() view returns(uint256)
func (_Bindings *BindingsCallerSession) DefaultBadPerformancePenalty() (*big.Int, error) {
	return _Bindings.Contract.DefaultBadPerformancePenalty(&_Bindings.CallOpts)
}

// DefaultElRewardsStealingAdditionalFine is a free data retrieval call binding the contract method 0xf8b84e6d.
//
// Solidity: function defaultElRewardsStealingAdditionalFine() view returns(uint256)
func (_Bindings *BindingsCaller) DefaultElRewardsStealingAdditionalFine(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "defaultElRewardsStealingAdditionalFine")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultElRewardsStealingAdditionalFine is a free data retrieval call binding the contract method 0xf8b84e6d.
//
// Solidity: function defaultElRewardsStealingAdditionalFine() view returns(uint256)
func (_Bindings *BindingsSession) DefaultElRewardsStealingAdditionalFine() (*big.Int, error) {
	return _Bindings.Contract.DefaultElRewardsStealingAdditionalFine(&_Bindings.CallOpts)
}

// DefaultElRewardsStealingAdditionalFine is a free data retrieval call binding the contract method 0xf8b84e6d.
//
// Solidity: function defaultElRewardsStealingAdditionalFine() view returns(uint256)
func (_Bindings *BindingsCallerSession) DefaultElRewardsStealingAdditionalFine() (*big.Int, error) {
	return _Bindings.Contract.DefaultElRewardsStealingAdditionalFine(&_Bindings.CallOpts)
}

// DefaultExitDelayPenalty is a free data retrieval call binding the contract method 0x541ec542.
//
// Solidity: function defaultExitDelayPenalty() view returns(uint256)
func (_Bindings *BindingsCaller) DefaultExitDelayPenalty(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "defaultExitDelayPenalty")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultExitDelayPenalty is a free data retrieval call binding the contract method 0x541ec542.
//
// Solidity: function defaultExitDelayPenalty() view returns(uint256)
func (_Bindings *BindingsSession) DefaultExitDelayPenalty() (*big.Int, error) {
	return _Bindings.Contract.DefaultExitDelayPenalty(&_Bindings.CallOpts)
}

// DefaultExitDelayPenalty is a free data retrieval call binding the contract method 0x541ec542.
//
// Solidity: function defaultExitDelayPenalty() view returns(uint256)
func (_Bindings *BindingsCallerSession) DefaultExitDelayPenalty() (*big.Int, error) {
	return _Bindings.Contract.DefaultExitDelayPenalty(&_Bindings.CallOpts)
}

// DefaultKeyRemovalCharge is a free data retrieval call binding the contract method 0xe5fa9277.
//
// Solidity: function defaultKeyRemovalCharge() view returns(uint256)
func (_Bindings *BindingsCaller) DefaultKeyRemovalCharge(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "defaultKeyRemovalCharge")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultKeyRemovalCharge is a free data retrieval call binding the contract method 0xe5fa9277.
//
// Solidity: function defaultKeyRemovalCharge() view returns(uint256)
func (_Bindings *BindingsSession) DefaultKeyRemovalCharge() (*big.Int, error) {
	return _Bindings.Contract.DefaultKeyRemovalCharge(&_Bindings.CallOpts)
}

// DefaultKeyRemovalCharge is a free data retrieval call binding the contract method 0xe5fa9277.
//
// Solidity: function defaultKeyRemovalCharge() view returns(uint256)
func (_Bindings *BindingsCallerSession) DefaultKeyRemovalCharge() (*big.Int, error) {
	return _Bindings.Contract.DefaultKeyRemovalCharge(&_Bindings.CallOpts)
}

// DefaultKeysLimit is a free data retrieval call binding the contract method 0x1e97abd1.
//
// Solidity: function defaultKeysLimit() view returns(uint256)
func (_Bindings *BindingsCaller) DefaultKeysLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "defaultKeysLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultKeysLimit is a free data retrieval call binding the contract method 0x1e97abd1.
//
// Solidity: function defaultKeysLimit() view returns(uint256)
func (_Bindings *BindingsSession) DefaultKeysLimit() (*big.Int, error) {
	return _Bindings.Contract.DefaultKeysLimit(&_Bindings.CallOpts)
}

// DefaultKeysLimit is a free data retrieval call binding the contract method 0x1e97abd1.
//
// Solidity: function defaultKeysLimit() view returns(uint256)
func (_Bindings *BindingsCallerSession) DefaultKeysLimit() (*big.Int, error) {
	return _Bindings.Contract.DefaultKeysLimit(&_Bindings.CallOpts)
}

// DefaultMaxWithdrawalRequestFee is a free data retrieval call binding the contract method 0xb591d09a.
//
// Solidity: function defaultMaxWithdrawalRequestFee() view returns(uint256)
func (_Bindings *BindingsCaller) DefaultMaxWithdrawalRequestFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "defaultMaxWithdrawalRequestFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultMaxWithdrawalRequestFee is a free data retrieval call binding the contract method 0xb591d09a.
//
// Solidity: function defaultMaxWithdrawalRequestFee() view returns(uint256)
func (_Bindings *BindingsSession) DefaultMaxWithdrawalRequestFee() (*big.Int, error) {
	return _Bindings.Contract.DefaultMaxWithdrawalRequestFee(&_Bindings.CallOpts)
}

// DefaultMaxWithdrawalRequestFee is a free data retrieval call binding the contract method 0xb591d09a.
//
// Solidity: function defaultMaxWithdrawalRequestFee() view returns(uint256)
func (_Bindings *BindingsCallerSession) DefaultMaxWithdrawalRequestFee() (*big.Int, error) {
	return _Bindings.Contract.DefaultMaxWithdrawalRequestFee(&_Bindings.CallOpts)
}

// DefaultPerformanceCoefficients is a free data retrieval call binding the contract method 0x3adb8f9c.
//
// Solidity: function defaultPerformanceCoefficients() view returns(uint32 attestationsWeight, uint32 blocksWeight, uint32 syncWeight)
func (_Bindings *BindingsCaller) DefaultPerformanceCoefficients(opts *bind.CallOpts) (struct {
	AttestationsWeight uint32
	BlocksWeight       uint32
	SyncWeight         uint32
}, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "defaultPerformanceCoefficients")

	outstruct := new(struct {
		AttestationsWeight uint32
		BlocksWeight       uint32
		SyncWeight         uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AttestationsWeight = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BlocksWeight = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.SyncWeight = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// DefaultPerformanceCoefficients is a free data retrieval call binding the contract method 0x3adb8f9c.
//
// Solidity: function defaultPerformanceCoefficients() view returns(uint32 attestationsWeight, uint32 blocksWeight, uint32 syncWeight)
func (_Bindings *BindingsSession) DefaultPerformanceCoefficients() (struct {
	AttestationsWeight uint32
	BlocksWeight       uint32
	SyncWeight         uint32
}, error) {
	return _Bindings.Contract.DefaultPerformanceCoefficients(&_Bindings.CallOpts)
}

// DefaultPerformanceCoefficients is a free data retrieval call binding the contract method 0x3adb8f9c.
//
// Solidity: function defaultPerformanceCoefficients() view returns(uint32 attestationsWeight, uint32 blocksWeight, uint32 syncWeight)
func (_Bindings *BindingsCallerSession) DefaultPerformanceCoefficients() (struct {
	AttestationsWeight uint32
	BlocksWeight       uint32
	SyncWeight         uint32
}, error) {
	return _Bindings.Contract.DefaultPerformanceCoefficients(&_Bindings.CallOpts)
}

// DefaultPerformanceLeeway is a free data retrieval call binding the contract method 0x7f3e3fa7.
//
// Solidity: function defaultPerformanceLeeway() view returns(uint256)
func (_Bindings *BindingsCaller) DefaultPerformanceLeeway(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "defaultPerformanceLeeway")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultPerformanceLeeway is a free data retrieval call binding the contract method 0x7f3e3fa7.
//
// Solidity: function defaultPerformanceLeeway() view returns(uint256)
func (_Bindings *BindingsSession) DefaultPerformanceLeeway() (*big.Int, error) {
	return _Bindings.Contract.DefaultPerformanceLeeway(&_Bindings.CallOpts)
}

// DefaultPerformanceLeeway is a free data retrieval call binding the contract method 0x7f3e3fa7.
//
// Solidity: function defaultPerformanceLeeway() view returns(uint256)
func (_Bindings *BindingsCallerSession) DefaultPerformanceLeeway() (*big.Int, error) {
	return _Bindings.Contract.DefaultPerformanceLeeway(&_Bindings.CallOpts)
}

// DefaultQueueConfig is a free data retrieval call binding the contract method 0xf18c961a.
//
// Solidity: function defaultQueueConfig() view returns(uint32 priority, uint32 maxDeposits)
func (_Bindings *BindingsCaller) DefaultQueueConfig(opts *bind.CallOpts) (struct {
	Priority    uint32
	MaxDeposits uint32
}, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "defaultQueueConfig")

	outstruct := new(struct {
		Priority    uint32
		MaxDeposits uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Priority = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.MaxDeposits = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// DefaultQueueConfig is a free data retrieval call binding the contract method 0xf18c961a.
//
// Solidity: function defaultQueueConfig() view returns(uint32 priority, uint32 maxDeposits)
func (_Bindings *BindingsSession) DefaultQueueConfig() (struct {
	Priority    uint32
	MaxDeposits uint32
}, error) {
	return _Bindings.Contract.DefaultQueueConfig(&_Bindings.CallOpts)
}

// DefaultQueueConfig is a free data retrieval call binding the contract method 0xf18c961a.
//
// Solidity: function defaultQueueConfig() view returns(uint32 priority, uint32 maxDeposits)
func (_Bindings *BindingsCallerSession) DefaultQueueConfig() (struct {
	Priority    uint32
	MaxDeposits uint32
}, error) {
	return _Bindings.Contract.DefaultQueueConfig(&_Bindings.CallOpts)
}

// DefaultRewardShare is a free data retrieval call binding the contract method 0xd96dd07e.
//
// Solidity: function defaultRewardShare() view returns(uint256)
func (_Bindings *BindingsCaller) DefaultRewardShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "defaultRewardShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultRewardShare is a free data retrieval call binding the contract method 0xd96dd07e.
//
// Solidity: function defaultRewardShare() view returns(uint256)
func (_Bindings *BindingsSession) DefaultRewardShare() (*big.Int, error) {
	return _Bindings.Contract.DefaultRewardShare(&_Bindings.CallOpts)
}

// DefaultRewardShare is a free data retrieval call binding the contract method 0xd96dd07e.
//
// Solidity: function defaultRewardShare() view returns(uint256)
func (_Bindings *BindingsCallerSession) DefaultRewardShare() (*big.Int, error) {
	return _Bindings.Contract.DefaultRewardShare(&_Bindings.CallOpts)
}

// DefaultStrikesParams is a free data retrieval call binding the contract method 0x9c9531d0.
//
// Solidity: function defaultStrikesParams() view returns(uint32 lifetime, uint32 threshold)
func (_Bindings *BindingsCaller) DefaultStrikesParams(opts *bind.CallOpts) (struct {
	Lifetime  uint32
	Threshold uint32
}, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "defaultStrikesParams")

	outstruct := new(struct {
		Lifetime  uint32
		Threshold uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Lifetime = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.Threshold = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// DefaultStrikesParams is a free data retrieval call binding the contract method 0x9c9531d0.
//
// Solidity: function defaultStrikesParams() view returns(uint32 lifetime, uint32 threshold)
func (_Bindings *BindingsSession) DefaultStrikesParams() (struct {
	Lifetime  uint32
	Threshold uint32
}, error) {
	return _Bindings.Contract.DefaultStrikesParams(&_Bindings.CallOpts)
}

// DefaultStrikesParams is a free data retrieval call binding the contract method 0x9c9531d0.
//
// Solidity: function defaultStrikesParams() view returns(uint32 lifetime, uint32 threshold)
func (_Bindings *BindingsCallerSession) DefaultStrikesParams() (struct {
	Lifetime  uint32
	Threshold uint32
}, error) {
	return _Bindings.Contract.DefaultStrikesParams(&_Bindings.CallOpts)
}

// GetAllowedExitDelay is a free data retrieval call binding the contract method 0x9a7b0508.
//
// Solidity: function getAllowedExitDelay(uint256 curveId) view returns(uint256 delay)
func (_Bindings *BindingsCaller) GetAllowedExitDelay(opts *bind.CallOpts, curveId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getAllowedExitDelay", curveId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAllowedExitDelay is a free data retrieval call binding the contract method 0x9a7b0508.
//
// Solidity: function getAllowedExitDelay(uint256 curveId) view returns(uint256 delay)
func (_Bindings *BindingsSession) GetAllowedExitDelay(curveId *big.Int) (*big.Int, error) {
	return _Bindings.Contract.GetAllowedExitDelay(&_Bindings.CallOpts, curveId)
}

// GetAllowedExitDelay is a free data retrieval call binding the contract method 0x9a7b0508.
//
// Solidity: function getAllowedExitDelay(uint256 curveId) view returns(uint256 delay)
func (_Bindings *BindingsCallerSession) GetAllowedExitDelay(curveId *big.Int) (*big.Int, error) {
	return _Bindings.Contract.GetAllowedExitDelay(&_Bindings.CallOpts, curveId)
}

// GetBadPerformancePenalty is a free data retrieval call binding the contract method 0x533c60d9.
//
// Solidity: function getBadPerformancePenalty(uint256 curveId) view returns(uint256 penalty)
func (_Bindings *BindingsCaller) GetBadPerformancePenalty(opts *bind.CallOpts, curveId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getBadPerformancePenalty", curveId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBadPerformancePenalty is a free data retrieval call binding the contract method 0x533c60d9.
//
// Solidity: function getBadPerformancePenalty(uint256 curveId) view returns(uint256 penalty)
func (_Bindings *BindingsSession) GetBadPerformancePenalty(curveId *big.Int) (*big.Int, error) {
	return _Bindings.Contract.GetBadPerformancePenalty(&_Bindings.CallOpts, curveId)
}

// GetBadPerformancePenalty is a free data retrieval call binding the contract method 0x533c60d9.
//
// Solidity: function getBadPerformancePenalty(uint256 curveId) view returns(uint256 penalty)
func (_Bindings *BindingsCallerSession) GetBadPerformancePenalty(curveId *big.Int) (*big.Int, error) {
	return _Bindings.Contract.GetBadPerformancePenalty(&_Bindings.CallOpts, curveId)
}

// GetElRewardsStealingAdditionalFine is a free data retrieval call binding the contract method 0x07a994c7.
//
// Solidity: function getElRewardsStealingAdditionalFine(uint256 curveId) view returns(uint256 fine)
func (_Bindings *BindingsCaller) GetElRewardsStealingAdditionalFine(opts *bind.CallOpts, curveId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getElRewardsStealingAdditionalFine", curveId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetElRewardsStealingAdditionalFine is a free data retrieval call binding the contract method 0x07a994c7.
//
// Solidity: function getElRewardsStealingAdditionalFine(uint256 curveId) view returns(uint256 fine)
func (_Bindings *BindingsSession) GetElRewardsStealingAdditionalFine(curveId *big.Int) (*big.Int, error) {
	return _Bindings.Contract.GetElRewardsStealingAdditionalFine(&_Bindings.CallOpts, curveId)
}

// GetElRewardsStealingAdditionalFine is a free data retrieval call binding the contract method 0x07a994c7.
//
// Solidity: function getElRewardsStealingAdditionalFine(uint256 curveId) view returns(uint256 fine)
func (_Bindings *BindingsCallerSession) GetElRewardsStealingAdditionalFine(curveId *big.Int) (*big.Int, error) {
	return _Bindings.Contract.GetElRewardsStealingAdditionalFine(&_Bindings.CallOpts, curveId)
}

// GetExitDelayPenalty is a free data retrieval call binding the contract method 0xf91a79b5.
//
// Solidity: function getExitDelayPenalty(uint256 curveId) view returns(uint256 penalty)
func (_Bindings *BindingsCaller) GetExitDelayPenalty(opts *bind.CallOpts, curveId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getExitDelayPenalty", curveId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExitDelayPenalty is a free data retrieval call binding the contract method 0xf91a79b5.
//
// Solidity: function getExitDelayPenalty(uint256 curveId) view returns(uint256 penalty)
func (_Bindings *BindingsSession) GetExitDelayPenalty(curveId *big.Int) (*big.Int, error) {
	return _Bindings.Contract.GetExitDelayPenalty(&_Bindings.CallOpts, curveId)
}

// GetExitDelayPenalty is a free data retrieval call binding the contract method 0xf91a79b5.
//
// Solidity: function getExitDelayPenalty(uint256 curveId) view returns(uint256 penalty)
func (_Bindings *BindingsCallerSession) GetExitDelayPenalty(curveId *big.Int) (*big.Int, error) {
	return _Bindings.Contract.GetExitDelayPenalty(&_Bindings.CallOpts, curveId)
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

// GetKeyRemovalCharge is a free data retrieval call binding the contract method 0xf42d7db5.
//
// Solidity: function getKeyRemovalCharge(uint256 curveId) view returns(uint256 keyRemovalCharge)
func (_Bindings *BindingsCaller) GetKeyRemovalCharge(opts *bind.CallOpts, curveId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getKeyRemovalCharge", curveId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetKeyRemovalCharge is a free data retrieval call binding the contract method 0xf42d7db5.
//
// Solidity: function getKeyRemovalCharge(uint256 curveId) view returns(uint256 keyRemovalCharge)
func (_Bindings *BindingsSession) GetKeyRemovalCharge(curveId *big.Int) (*big.Int, error) {
	return _Bindings.Contract.GetKeyRemovalCharge(&_Bindings.CallOpts, curveId)
}

// GetKeyRemovalCharge is a free data retrieval call binding the contract method 0xf42d7db5.
//
// Solidity: function getKeyRemovalCharge(uint256 curveId) view returns(uint256 keyRemovalCharge)
func (_Bindings *BindingsCallerSession) GetKeyRemovalCharge(curveId *big.Int) (*big.Int, error) {
	return _Bindings.Contract.GetKeyRemovalCharge(&_Bindings.CallOpts, curveId)
}

// GetKeysLimit is a free data retrieval call binding the contract method 0x29bbbd60.
//
// Solidity: function getKeysLimit(uint256 curveId) view returns(uint256 limit)
func (_Bindings *BindingsCaller) GetKeysLimit(opts *bind.CallOpts, curveId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getKeysLimit", curveId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetKeysLimit is a free data retrieval call binding the contract method 0x29bbbd60.
//
// Solidity: function getKeysLimit(uint256 curveId) view returns(uint256 limit)
func (_Bindings *BindingsSession) GetKeysLimit(curveId *big.Int) (*big.Int, error) {
	return _Bindings.Contract.GetKeysLimit(&_Bindings.CallOpts, curveId)
}

// GetKeysLimit is a free data retrieval call binding the contract method 0x29bbbd60.
//
// Solidity: function getKeysLimit(uint256 curveId) view returns(uint256 limit)
func (_Bindings *BindingsCallerSession) GetKeysLimit(curveId *big.Int) (*big.Int, error) {
	return _Bindings.Contract.GetKeysLimit(&_Bindings.CallOpts, curveId)
}

// GetMaxWithdrawalRequestFee is a free data retrieval call binding the contract method 0xacd98130.
//
// Solidity: function getMaxWithdrawalRequestFee(uint256 curveId) view returns(uint256 fee)
func (_Bindings *BindingsCaller) GetMaxWithdrawalRequestFee(opts *bind.CallOpts, curveId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getMaxWithdrawalRequestFee", curveId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxWithdrawalRequestFee is a free data retrieval call binding the contract method 0xacd98130.
//
// Solidity: function getMaxWithdrawalRequestFee(uint256 curveId) view returns(uint256 fee)
func (_Bindings *BindingsSession) GetMaxWithdrawalRequestFee(curveId *big.Int) (*big.Int, error) {
	return _Bindings.Contract.GetMaxWithdrawalRequestFee(&_Bindings.CallOpts, curveId)
}

// GetMaxWithdrawalRequestFee is a free data retrieval call binding the contract method 0xacd98130.
//
// Solidity: function getMaxWithdrawalRequestFee(uint256 curveId) view returns(uint256 fee)
func (_Bindings *BindingsCallerSession) GetMaxWithdrawalRequestFee(curveId *big.Int) (*big.Int, error) {
	return _Bindings.Contract.GetMaxWithdrawalRequestFee(&_Bindings.CallOpts, curveId)
}

// GetPerformanceCoefficients is a free data retrieval call binding the contract method 0x600ce782.
//
// Solidity: function getPerformanceCoefficients(uint256 curveId) view returns(uint256 attestationsWeight, uint256 blocksWeight, uint256 syncWeight)
func (_Bindings *BindingsCaller) GetPerformanceCoefficients(opts *bind.CallOpts, curveId *big.Int) (struct {
	AttestationsWeight *big.Int
	BlocksWeight       *big.Int
	SyncWeight         *big.Int
}, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getPerformanceCoefficients", curveId)

	outstruct := new(struct {
		AttestationsWeight *big.Int
		BlocksWeight       *big.Int
		SyncWeight         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AttestationsWeight = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BlocksWeight = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SyncWeight = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPerformanceCoefficients is a free data retrieval call binding the contract method 0x600ce782.
//
// Solidity: function getPerformanceCoefficients(uint256 curveId) view returns(uint256 attestationsWeight, uint256 blocksWeight, uint256 syncWeight)
func (_Bindings *BindingsSession) GetPerformanceCoefficients(curveId *big.Int) (struct {
	AttestationsWeight *big.Int
	BlocksWeight       *big.Int
	SyncWeight         *big.Int
}, error) {
	return _Bindings.Contract.GetPerformanceCoefficients(&_Bindings.CallOpts, curveId)
}

// GetPerformanceCoefficients is a free data retrieval call binding the contract method 0x600ce782.
//
// Solidity: function getPerformanceCoefficients(uint256 curveId) view returns(uint256 attestationsWeight, uint256 blocksWeight, uint256 syncWeight)
func (_Bindings *BindingsCallerSession) GetPerformanceCoefficients(curveId *big.Int) (struct {
	AttestationsWeight *big.Int
	BlocksWeight       *big.Int
	SyncWeight         *big.Int
}, error) {
	return _Bindings.Contract.GetPerformanceCoefficients(&_Bindings.CallOpts, curveId)
}

// GetPerformanceLeewayData is a free data retrieval call binding the contract method 0x95453a2d.
//
// Solidity: function getPerformanceLeewayData(uint256 curveId) view returns((uint256,uint256)[] data)
func (_Bindings *BindingsCaller) GetPerformanceLeewayData(opts *bind.CallOpts, curveId *big.Int) ([]ICSParametersRegistryKeyNumberValueInterval, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getPerformanceLeewayData", curveId)

	if err != nil {
		return *new([]ICSParametersRegistryKeyNumberValueInterval), err
	}

	out0 := *abi.ConvertType(out[0], new([]ICSParametersRegistryKeyNumberValueInterval)).(*[]ICSParametersRegistryKeyNumberValueInterval)

	return out0, err

}

// GetPerformanceLeewayData is a free data retrieval call binding the contract method 0x95453a2d.
//
// Solidity: function getPerformanceLeewayData(uint256 curveId) view returns((uint256,uint256)[] data)
func (_Bindings *BindingsSession) GetPerformanceLeewayData(curveId *big.Int) ([]ICSParametersRegistryKeyNumberValueInterval, error) {
	return _Bindings.Contract.GetPerformanceLeewayData(&_Bindings.CallOpts, curveId)
}

// GetPerformanceLeewayData is a free data retrieval call binding the contract method 0x95453a2d.
//
// Solidity: function getPerformanceLeewayData(uint256 curveId) view returns((uint256,uint256)[] data)
func (_Bindings *BindingsCallerSession) GetPerformanceLeewayData(curveId *big.Int) ([]ICSParametersRegistryKeyNumberValueInterval, error) {
	return _Bindings.Contract.GetPerformanceLeewayData(&_Bindings.CallOpts, curveId)
}

// GetQueueConfig is a free data retrieval call binding the contract method 0xdecfec56.
//
// Solidity: function getQueueConfig(uint256 curveId) view returns(uint32 queuePriority, uint32 maxDeposits)
func (_Bindings *BindingsCaller) GetQueueConfig(opts *bind.CallOpts, curveId *big.Int) (struct {
	QueuePriority uint32
	MaxDeposits   uint32
}, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getQueueConfig", curveId)

	outstruct := new(struct {
		QueuePriority uint32
		MaxDeposits   uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.QueuePriority = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.MaxDeposits = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetQueueConfig is a free data retrieval call binding the contract method 0xdecfec56.
//
// Solidity: function getQueueConfig(uint256 curveId) view returns(uint32 queuePriority, uint32 maxDeposits)
func (_Bindings *BindingsSession) GetQueueConfig(curveId *big.Int) (struct {
	QueuePriority uint32
	MaxDeposits   uint32
}, error) {
	return _Bindings.Contract.GetQueueConfig(&_Bindings.CallOpts, curveId)
}

// GetQueueConfig is a free data retrieval call binding the contract method 0xdecfec56.
//
// Solidity: function getQueueConfig(uint256 curveId) view returns(uint32 queuePriority, uint32 maxDeposits)
func (_Bindings *BindingsCallerSession) GetQueueConfig(curveId *big.Int) (struct {
	QueuePriority uint32
	MaxDeposits   uint32
}, error) {
	return _Bindings.Contract.GetQueueConfig(&_Bindings.CallOpts, curveId)
}

// GetRewardShareData is a free data retrieval call binding the contract method 0x4031002d.
//
// Solidity: function getRewardShareData(uint256 curveId) view returns((uint256,uint256)[] data)
func (_Bindings *BindingsCaller) GetRewardShareData(opts *bind.CallOpts, curveId *big.Int) ([]ICSParametersRegistryKeyNumberValueInterval, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getRewardShareData", curveId)

	if err != nil {
		return *new([]ICSParametersRegistryKeyNumberValueInterval), err
	}

	out0 := *abi.ConvertType(out[0], new([]ICSParametersRegistryKeyNumberValueInterval)).(*[]ICSParametersRegistryKeyNumberValueInterval)

	return out0, err

}

// GetRewardShareData is a free data retrieval call binding the contract method 0x4031002d.
//
// Solidity: function getRewardShareData(uint256 curveId) view returns((uint256,uint256)[] data)
func (_Bindings *BindingsSession) GetRewardShareData(curveId *big.Int) ([]ICSParametersRegistryKeyNumberValueInterval, error) {
	return _Bindings.Contract.GetRewardShareData(&_Bindings.CallOpts, curveId)
}

// GetRewardShareData is a free data retrieval call binding the contract method 0x4031002d.
//
// Solidity: function getRewardShareData(uint256 curveId) view returns((uint256,uint256)[] data)
func (_Bindings *BindingsCallerSession) GetRewardShareData(curveId *big.Int) ([]ICSParametersRegistryKeyNumberValueInterval, error) {
	return _Bindings.Contract.GetRewardShareData(&_Bindings.CallOpts, curveId)
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

// GetStrikesParams is a free data retrieval call binding the contract method 0x5684e702.
//
// Solidity: function getStrikesParams(uint256 curveId) view returns(uint256 lifetime, uint256 threshold)
func (_Bindings *BindingsCaller) GetStrikesParams(opts *bind.CallOpts, curveId *big.Int) (struct {
	Lifetime  *big.Int
	Threshold *big.Int
}, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getStrikesParams", curveId)

	outstruct := new(struct {
		Lifetime  *big.Int
		Threshold *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Lifetime = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Threshold = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetStrikesParams is a free data retrieval call binding the contract method 0x5684e702.
//
// Solidity: function getStrikesParams(uint256 curveId) view returns(uint256 lifetime, uint256 threshold)
func (_Bindings *BindingsSession) GetStrikesParams(curveId *big.Int) (struct {
	Lifetime  *big.Int
	Threshold *big.Int
}, error) {
	return _Bindings.Contract.GetStrikesParams(&_Bindings.CallOpts, curveId)
}

// GetStrikesParams is a free data retrieval call binding the contract method 0x5684e702.
//
// Solidity: function getStrikesParams(uint256 curveId) view returns(uint256 lifetime, uint256 threshold)
func (_Bindings *BindingsCallerSession) GetStrikesParams(curveId *big.Int) (struct {
	Lifetime  *big.Int
	Threshold *big.Int
}, error) {
	return _Bindings.Contract.GetStrikesParams(&_Bindings.CallOpts, curveId)
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

// Initialize is a paid mutator transaction binding the contract method 0xdec45c38.
//
// Solidity: function initialize(address admin, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) data) returns()
func (_Bindings *BindingsTransactor) Initialize(opts *bind.TransactOpts, admin common.Address, data ICSParametersRegistryInitializationData) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "initialize", admin, data)
}

// Initialize is a paid mutator transaction binding the contract method 0xdec45c38.
//
// Solidity: function initialize(address admin, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) data) returns()
func (_Bindings *BindingsSession) Initialize(admin common.Address, data ICSParametersRegistryInitializationData) (*types.Transaction, error) {
	return _Bindings.Contract.Initialize(&_Bindings.TransactOpts, admin, data)
}

// Initialize is a paid mutator transaction binding the contract method 0xdec45c38.
//
// Solidity: function initialize(address admin, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) data) returns()
func (_Bindings *BindingsTransactorSession) Initialize(admin common.Address, data ICSParametersRegistryInitializationData) (*types.Transaction, error) {
	return _Bindings.Contract.Initialize(&_Bindings.TransactOpts, admin, data)
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

// SetAllowedExitDelay is a paid mutator transaction binding the contract method 0xa35bb07e.
//
// Solidity: function setAllowedExitDelay(uint256 curveId, uint256 delay) returns()
func (_Bindings *BindingsTransactor) SetAllowedExitDelay(opts *bind.TransactOpts, curveId *big.Int, delay *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setAllowedExitDelay", curveId, delay)
}

// SetAllowedExitDelay is a paid mutator transaction binding the contract method 0xa35bb07e.
//
// Solidity: function setAllowedExitDelay(uint256 curveId, uint256 delay) returns()
func (_Bindings *BindingsSession) SetAllowedExitDelay(curveId *big.Int, delay *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetAllowedExitDelay(&_Bindings.TransactOpts, curveId, delay)
}

// SetAllowedExitDelay is a paid mutator transaction binding the contract method 0xa35bb07e.
//
// Solidity: function setAllowedExitDelay(uint256 curveId, uint256 delay) returns()
func (_Bindings *BindingsTransactorSession) SetAllowedExitDelay(curveId *big.Int, delay *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetAllowedExitDelay(&_Bindings.TransactOpts, curveId, delay)
}

// SetBadPerformancePenalty is a paid mutator transaction binding the contract method 0x808cdedd.
//
// Solidity: function setBadPerformancePenalty(uint256 curveId, uint256 penalty) returns()
func (_Bindings *BindingsTransactor) SetBadPerformancePenalty(opts *bind.TransactOpts, curveId *big.Int, penalty *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setBadPerformancePenalty", curveId, penalty)
}

// SetBadPerformancePenalty is a paid mutator transaction binding the contract method 0x808cdedd.
//
// Solidity: function setBadPerformancePenalty(uint256 curveId, uint256 penalty) returns()
func (_Bindings *BindingsSession) SetBadPerformancePenalty(curveId *big.Int, penalty *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetBadPerformancePenalty(&_Bindings.TransactOpts, curveId, penalty)
}

// SetBadPerformancePenalty is a paid mutator transaction binding the contract method 0x808cdedd.
//
// Solidity: function setBadPerformancePenalty(uint256 curveId, uint256 penalty) returns()
func (_Bindings *BindingsTransactorSession) SetBadPerformancePenalty(curveId *big.Int, penalty *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetBadPerformancePenalty(&_Bindings.TransactOpts, curveId, penalty)
}

// SetDefaultAllowedExitDelay is a paid mutator transaction binding the contract method 0x1a31fb30.
//
// Solidity: function setDefaultAllowedExitDelay(uint256 delay) returns()
func (_Bindings *BindingsTransactor) SetDefaultAllowedExitDelay(opts *bind.TransactOpts, delay *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setDefaultAllowedExitDelay", delay)
}

// SetDefaultAllowedExitDelay is a paid mutator transaction binding the contract method 0x1a31fb30.
//
// Solidity: function setDefaultAllowedExitDelay(uint256 delay) returns()
func (_Bindings *BindingsSession) SetDefaultAllowedExitDelay(delay *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetDefaultAllowedExitDelay(&_Bindings.TransactOpts, delay)
}

// SetDefaultAllowedExitDelay is a paid mutator transaction binding the contract method 0x1a31fb30.
//
// Solidity: function setDefaultAllowedExitDelay(uint256 delay) returns()
func (_Bindings *BindingsTransactorSession) SetDefaultAllowedExitDelay(delay *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetDefaultAllowedExitDelay(&_Bindings.TransactOpts, delay)
}

// SetDefaultBadPerformancePenalty is a paid mutator transaction binding the contract method 0x1472e1cd.
//
// Solidity: function setDefaultBadPerformancePenalty(uint256 penalty) returns()
func (_Bindings *BindingsTransactor) SetDefaultBadPerformancePenalty(opts *bind.TransactOpts, penalty *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setDefaultBadPerformancePenalty", penalty)
}

// SetDefaultBadPerformancePenalty is a paid mutator transaction binding the contract method 0x1472e1cd.
//
// Solidity: function setDefaultBadPerformancePenalty(uint256 penalty) returns()
func (_Bindings *BindingsSession) SetDefaultBadPerformancePenalty(penalty *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetDefaultBadPerformancePenalty(&_Bindings.TransactOpts, penalty)
}

// SetDefaultBadPerformancePenalty is a paid mutator transaction binding the contract method 0x1472e1cd.
//
// Solidity: function setDefaultBadPerformancePenalty(uint256 penalty) returns()
func (_Bindings *BindingsTransactorSession) SetDefaultBadPerformancePenalty(penalty *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetDefaultBadPerformancePenalty(&_Bindings.TransactOpts, penalty)
}

// SetDefaultElRewardsStealingAdditionalFine is a paid mutator transaction binding the contract method 0xd8ef52f4.
//
// Solidity: function setDefaultElRewardsStealingAdditionalFine(uint256 fine) returns()
func (_Bindings *BindingsTransactor) SetDefaultElRewardsStealingAdditionalFine(opts *bind.TransactOpts, fine *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setDefaultElRewardsStealingAdditionalFine", fine)
}

// SetDefaultElRewardsStealingAdditionalFine is a paid mutator transaction binding the contract method 0xd8ef52f4.
//
// Solidity: function setDefaultElRewardsStealingAdditionalFine(uint256 fine) returns()
func (_Bindings *BindingsSession) SetDefaultElRewardsStealingAdditionalFine(fine *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetDefaultElRewardsStealingAdditionalFine(&_Bindings.TransactOpts, fine)
}

// SetDefaultElRewardsStealingAdditionalFine is a paid mutator transaction binding the contract method 0xd8ef52f4.
//
// Solidity: function setDefaultElRewardsStealingAdditionalFine(uint256 fine) returns()
func (_Bindings *BindingsTransactorSession) SetDefaultElRewardsStealingAdditionalFine(fine *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetDefaultElRewardsStealingAdditionalFine(&_Bindings.TransactOpts, fine)
}

// SetDefaultExitDelayPenalty is a paid mutator transaction binding the contract method 0x3deace15.
//
// Solidity: function setDefaultExitDelayPenalty(uint256 penalty) returns()
func (_Bindings *BindingsTransactor) SetDefaultExitDelayPenalty(opts *bind.TransactOpts, penalty *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setDefaultExitDelayPenalty", penalty)
}

// SetDefaultExitDelayPenalty is a paid mutator transaction binding the contract method 0x3deace15.
//
// Solidity: function setDefaultExitDelayPenalty(uint256 penalty) returns()
func (_Bindings *BindingsSession) SetDefaultExitDelayPenalty(penalty *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetDefaultExitDelayPenalty(&_Bindings.TransactOpts, penalty)
}

// SetDefaultExitDelayPenalty is a paid mutator transaction binding the contract method 0x3deace15.
//
// Solidity: function setDefaultExitDelayPenalty(uint256 penalty) returns()
func (_Bindings *BindingsTransactorSession) SetDefaultExitDelayPenalty(penalty *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetDefaultExitDelayPenalty(&_Bindings.TransactOpts, penalty)
}

// SetDefaultKeyRemovalCharge is a paid mutator transaction binding the contract method 0x3373a164.
//
// Solidity: function setDefaultKeyRemovalCharge(uint256 keyRemovalCharge) returns()
func (_Bindings *BindingsTransactor) SetDefaultKeyRemovalCharge(opts *bind.TransactOpts, keyRemovalCharge *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setDefaultKeyRemovalCharge", keyRemovalCharge)
}

// SetDefaultKeyRemovalCharge is a paid mutator transaction binding the contract method 0x3373a164.
//
// Solidity: function setDefaultKeyRemovalCharge(uint256 keyRemovalCharge) returns()
func (_Bindings *BindingsSession) SetDefaultKeyRemovalCharge(keyRemovalCharge *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetDefaultKeyRemovalCharge(&_Bindings.TransactOpts, keyRemovalCharge)
}

// SetDefaultKeyRemovalCharge is a paid mutator transaction binding the contract method 0x3373a164.
//
// Solidity: function setDefaultKeyRemovalCharge(uint256 keyRemovalCharge) returns()
func (_Bindings *BindingsTransactorSession) SetDefaultKeyRemovalCharge(keyRemovalCharge *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetDefaultKeyRemovalCharge(&_Bindings.TransactOpts, keyRemovalCharge)
}

// SetDefaultKeysLimit is a paid mutator transaction binding the contract method 0xf535c7ea.
//
// Solidity: function setDefaultKeysLimit(uint256 limit) returns()
func (_Bindings *BindingsTransactor) SetDefaultKeysLimit(opts *bind.TransactOpts, limit *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setDefaultKeysLimit", limit)
}

// SetDefaultKeysLimit is a paid mutator transaction binding the contract method 0xf535c7ea.
//
// Solidity: function setDefaultKeysLimit(uint256 limit) returns()
func (_Bindings *BindingsSession) SetDefaultKeysLimit(limit *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetDefaultKeysLimit(&_Bindings.TransactOpts, limit)
}

// SetDefaultKeysLimit is a paid mutator transaction binding the contract method 0xf535c7ea.
//
// Solidity: function setDefaultKeysLimit(uint256 limit) returns()
func (_Bindings *BindingsTransactorSession) SetDefaultKeysLimit(limit *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetDefaultKeysLimit(&_Bindings.TransactOpts, limit)
}

// SetDefaultMaxWithdrawalRequestFee is a paid mutator transaction binding the contract method 0xb0a29b7c.
//
// Solidity: function setDefaultMaxWithdrawalRequestFee(uint256 fee) returns()
func (_Bindings *BindingsTransactor) SetDefaultMaxWithdrawalRequestFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setDefaultMaxWithdrawalRequestFee", fee)
}

// SetDefaultMaxWithdrawalRequestFee is a paid mutator transaction binding the contract method 0xb0a29b7c.
//
// Solidity: function setDefaultMaxWithdrawalRequestFee(uint256 fee) returns()
func (_Bindings *BindingsSession) SetDefaultMaxWithdrawalRequestFee(fee *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetDefaultMaxWithdrawalRequestFee(&_Bindings.TransactOpts, fee)
}

// SetDefaultMaxWithdrawalRequestFee is a paid mutator transaction binding the contract method 0xb0a29b7c.
//
// Solidity: function setDefaultMaxWithdrawalRequestFee(uint256 fee) returns()
func (_Bindings *BindingsTransactorSession) SetDefaultMaxWithdrawalRequestFee(fee *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetDefaultMaxWithdrawalRequestFee(&_Bindings.TransactOpts, fee)
}

// SetDefaultPerformanceCoefficients is a paid mutator transaction binding the contract method 0x9edf0ecc.
//
// Solidity: function setDefaultPerformanceCoefficients(uint256 attestationsWeight, uint256 blocksWeight, uint256 syncWeight) returns()
func (_Bindings *BindingsTransactor) SetDefaultPerformanceCoefficients(opts *bind.TransactOpts, attestationsWeight *big.Int, blocksWeight *big.Int, syncWeight *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setDefaultPerformanceCoefficients", attestationsWeight, blocksWeight, syncWeight)
}

// SetDefaultPerformanceCoefficients is a paid mutator transaction binding the contract method 0x9edf0ecc.
//
// Solidity: function setDefaultPerformanceCoefficients(uint256 attestationsWeight, uint256 blocksWeight, uint256 syncWeight) returns()
func (_Bindings *BindingsSession) SetDefaultPerformanceCoefficients(attestationsWeight *big.Int, blocksWeight *big.Int, syncWeight *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetDefaultPerformanceCoefficients(&_Bindings.TransactOpts, attestationsWeight, blocksWeight, syncWeight)
}

// SetDefaultPerformanceCoefficients is a paid mutator transaction binding the contract method 0x9edf0ecc.
//
// Solidity: function setDefaultPerformanceCoefficients(uint256 attestationsWeight, uint256 blocksWeight, uint256 syncWeight) returns()
func (_Bindings *BindingsTransactorSession) SetDefaultPerformanceCoefficients(attestationsWeight *big.Int, blocksWeight *big.Int, syncWeight *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetDefaultPerformanceCoefficients(&_Bindings.TransactOpts, attestationsWeight, blocksWeight, syncWeight)
}

// SetDefaultPerformanceLeeway is a paid mutator transaction binding the contract method 0x33cdbd40.
//
// Solidity: function setDefaultPerformanceLeeway(uint256 leeway) returns()
func (_Bindings *BindingsTransactor) SetDefaultPerformanceLeeway(opts *bind.TransactOpts, leeway *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setDefaultPerformanceLeeway", leeway)
}

// SetDefaultPerformanceLeeway is a paid mutator transaction binding the contract method 0x33cdbd40.
//
// Solidity: function setDefaultPerformanceLeeway(uint256 leeway) returns()
func (_Bindings *BindingsSession) SetDefaultPerformanceLeeway(leeway *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetDefaultPerformanceLeeway(&_Bindings.TransactOpts, leeway)
}

// SetDefaultPerformanceLeeway is a paid mutator transaction binding the contract method 0x33cdbd40.
//
// Solidity: function setDefaultPerformanceLeeway(uint256 leeway) returns()
func (_Bindings *BindingsTransactorSession) SetDefaultPerformanceLeeway(leeway *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetDefaultPerformanceLeeway(&_Bindings.TransactOpts, leeway)
}

// SetDefaultQueueConfig is a paid mutator transaction binding the contract method 0x8083dc48.
//
// Solidity: function setDefaultQueueConfig(uint256 priority, uint256 maxDeposits) returns()
func (_Bindings *BindingsTransactor) SetDefaultQueueConfig(opts *bind.TransactOpts, priority *big.Int, maxDeposits *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setDefaultQueueConfig", priority, maxDeposits)
}

// SetDefaultQueueConfig is a paid mutator transaction binding the contract method 0x8083dc48.
//
// Solidity: function setDefaultQueueConfig(uint256 priority, uint256 maxDeposits) returns()
func (_Bindings *BindingsSession) SetDefaultQueueConfig(priority *big.Int, maxDeposits *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetDefaultQueueConfig(&_Bindings.TransactOpts, priority, maxDeposits)
}

// SetDefaultQueueConfig is a paid mutator transaction binding the contract method 0x8083dc48.
//
// Solidity: function setDefaultQueueConfig(uint256 priority, uint256 maxDeposits) returns()
func (_Bindings *BindingsTransactorSession) SetDefaultQueueConfig(priority *big.Int, maxDeposits *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetDefaultQueueConfig(&_Bindings.TransactOpts, priority, maxDeposits)
}

// SetDefaultRewardShare is a paid mutator transaction binding the contract method 0xf2db2c8b.
//
// Solidity: function setDefaultRewardShare(uint256 share) returns()
func (_Bindings *BindingsTransactor) SetDefaultRewardShare(opts *bind.TransactOpts, share *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setDefaultRewardShare", share)
}

// SetDefaultRewardShare is a paid mutator transaction binding the contract method 0xf2db2c8b.
//
// Solidity: function setDefaultRewardShare(uint256 share) returns()
func (_Bindings *BindingsSession) SetDefaultRewardShare(share *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetDefaultRewardShare(&_Bindings.TransactOpts, share)
}

// SetDefaultRewardShare is a paid mutator transaction binding the contract method 0xf2db2c8b.
//
// Solidity: function setDefaultRewardShare(uint256 share) returns()
func (_Bindings *BindingsTransactorSession) SetDefaultRewardShare(share *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetDefaultRewardShare(&_Bindings.TransactOpts, share)
}

// SetDefaultStrikesParams is a paid mutator transaction binding the contract method 0x75d9e810.
//
// Solidity: function setDefaultStrikesParams(uint256 lifetime, uint256 threshold) returns()
func (_Bindings *BindingsTransactor) SetDefaultStrikesParams(opts *bind.TransactOpts, lifetime *big.Int, threshold *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setDefaultStrikesParams", lifetime, threshold)
}

// SetDefaultStrikesParams is a paid mutator transaction binding the contract method 0x75d9e810.
//
// Solidity: function setDefaultStrikesParams(uint256 lifetime, uint256 threshold) returns()
func (_Bindings *BindingsSession) SetDefaultStrikesParams(lifetime *big.Int, threshold *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetDefaultStrikesParams(&_Bindings.TransactOpts, lifetime, threshold)
}

// SetDefaultStrikesParams is a paid mutator transaction binding the contract method 0x75d9e810.
//
// Solidity: function setDefaultStrikesParams(uint256 lifetime, uint256 threshold) returns()
func (_Bindings *BindingsTransactorSession) SetDefaultStrikesParams(lifetime *big.Int, threshold *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetDefaultStrikesParams(&_Bindings.TransactOpts, lifetime, threshold)
}

// SetElRewardsStealingAdditionalFine is a paid mutator transaction binding the contract method 0x0820659c.
//
// Solidity: function setElRewardsStealingAdditionalFine(uint256 curveId, uint256 fine) returns()
func (_Bindings *BindingsTransactor) SetElRewardsStealingAdditionalFine(opts *bind.TransactOpts, curveId *big.Int, fine *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setElRewardsStealingAdditionalFine", curveId, fine)
}

// SetElRewardsStealingAdditionalFine is a paid mutator transaction binding the contract method 0x0820659c.
//
// Solidity: function setElRewardsStealingAdditionalFine(uint256 curveId, uint256 fine) returns()
func (_Bindings *BindingsSession) SetElRewardsStealingAdditionalFine(curveId *big.Int, fine *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetElRewardsStealingAdditionalFine(&_Bindings.TransactOpts, curveId, fine)
}

// SetElRewardsStealingAdditionalFine is a paid mutator transaction binding the contract method 0x0820659c.
//
// Solidity: function setElRewardsStealingAdditionalFine(uint256 curveId, uint256 fine) returns()
func (_Bindings *BindingsTransactorSession) SetElRewardsStealingAdditionalFine(curveId *big.Int, fine *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetElRewardsStealingAdditionalFine(&_Bindings.TransactOpts, curveId, fine)
}

// SetExitDelayPenalty is a paid mutator transaction binding the contract method 0x71a7a39e.
//
// Solidity: function setExitDelayPenalty(uint256 curveId, uint256 penalty) returns()
func (_Bindings *BindingsTransactor) SetExitDelayPenalty(opts *bind.TransactOpts, curveId *big.Int, penalty *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setExitDelayPenalty", curveId, penalty)
}

// SetExitDelayPenalty is a paid mutator transaction binding the contract method 0x71a7a39e.
//
// Solidity: function setExitDelayPenalty(uint256 curveId, uint256 penalty) returns()
func (_Bindings *BindingsSession) SetExitDelayPenalty(curveId *big.Int, penalty *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetExitDelayPenalty(&_Bindings.TransactOpts, curveId, penalty)
}

// SetExitDelayPenalty is a paid mutator transaction binding the contract method 0x71a7a39e.
//
// Solidity: function setExitDelayPenalty(uint256 curveId, uint256 penalty) returns()
func (_Bindings *BindingsTransactorSession) SetExitDelayPenalty(curveId *big.Int, penalty *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetExitDelayPenalty(&_Bindings.TransactOpts, curveId, penalty)
}

// SetKeyRemovalCharge is a paid mutator transaction binding the contract method 0x1332971c.
//
// Solidity: function setKeyRemovalCharge(uint256 curveId, uint256 keyRemovalCharge) returns()
func (_Bindings *BindingsTransactor) SetKeyRemovalCharge(opts *bind.TransactOpts, curveId *big.Int, keyRemovalCharge *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setKeyRemovalCharge", curveId, keyRemovalCharge)
}

// SetKeyRemovalCharge is a paid mutator transaction binding the contract method 0x1332971c.
//
// Solidity: function setKeyRemovalCharge(uint256 curveId, uint256 keyRemovalCharge) returns()
func (_Bindings *BindingsSession) SetKeyRemovalCharge(curveId *big.Int, keyRemovalCharge *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetKeyRemovalCharge(&_Bindings.TransactOpts, curveId, keyRemovalCharge)
}

// SetKeyRemovalCharge is a paid mutator transaction binding the contract method 0x1332971c.
//
// Solidity: function setKeyRemovalCharge(uint256 curveId, uint256 keyRemovalCharge) returns()
func (_Bindings *BindingsTransactorSession) SetKeyRemovalCharge(curveId *big.Int, keyRemovalCharge *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetKeyRemovalCharge(&_Bindings.TransactOpts, curveId, keyRemovalCharge)
}

// SetKeysLimit is a paid mutator transaction binding the contract method 0x1974fa1d.
//
// Solidity: function setKeysLimit(uint256 curveId, uint256 limit) returns()
func (_Bindings *BindingsTransactor) SetKeysLimit(opts *bind.TransactOpts, curveId *big.Int, limit *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setKeysLimit", curveId, limit)
}

// SetKeysLimit is a paid mutator transaction binding the contract method 0x1974fa1d.
//
// Solidity: function setKeysLimit(uint256 curveId, uint256 limit) returns()
func (_Bindings *BindingsSession) SetKeysLimit(curveId *big.Int, limit *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetKeysLimit(&_Bindings.TransactOpts, curveId, limit)
}

// SetKeysLimit is a paid mutator transaction binding the contract method 0x1974fa1d.
//
// Solidity: function setKeysLimit(uint256 curveId, uint256 limit) returns()
func (_Bindings *BindingsTransactorSession) SetKeysLimit(curveId *big.Int, limit *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetKeysLimit(&_Bindings.TransactOpts, curveId, limit)
}

// SetMaxWithdrawalRequestFee is a paid mutator transaction binding the contract method 0x6496a430.
//
// Solidity: function setMaxWithdrawalRequestFee(uint256 curveId, uint256 fee) returns()
func (_Bindings *BindingsTransactor) SetMaxWithdrawalRequestFee(opts *bind.TransactOpts, curveId *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setMaxWithdrawalRequestFee", curveId, fee)
}

// SetMaxWithdrawalRequestFee is a paid mutator transaction binding the contract method 0x6496a430.
//
// Solidity: function setMaxWithdrawalRequestFee(uint256 curveId, uint256 fee) returns()
func (_Bindings *BindingsSession) SetMaxWithdrawalRequestFee(curveId *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetMaxWithdrawalRequestFee(&_Bindings.TransactOpts, curveId, fee)
}

// SetMaxWithdrawalRequestFee is a paid mutator transaction binding the contract method 0x6496a430.
//
// Solidity: function setMaxWithdrawalRequestFee(uint256 curveId, uint256 fee) returns()
func (_Bindings *BindingsTransactorSession) SetMaxWithdrawalRequestFee(curveId *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetMaxWithdrawalRequestFee(&_Bindings.TransactOpts, curveId, fee)
}

// SetPerformanceCoefficients is a paid mutator transaction binding the contract method 0xbf49bbde.
//
// Solidity: function setPerformanceCoefficients(uint256 curveId, uint256 attestationsWeight, uint256 blocksWeight, uint256 syncWeight) returns()
func (_Bindings *BindingsTransactor) SetPerformanceCoefficients(opts *bind.TransactOpts, curveId *big.Int, attestationsWeight *big.Int, blocksWeight *big.Int, syncWeight *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setPerformanceCoefficients", curveId, attestationsWeight, blocksWeight, syncWeight)
}

// SetPerformanceCoefficients is a paid mutator transaction binding the contract method 0xbf49bbde.
//
// Solidity: function setPerformanceCoefficients(uint256 curveId, uint256 attestationsWeight, uint256 blocksWeight, uint256 syncWeight) returns()
func (_Bindings *BindingsSession) SetPerformanceCoefficients(curveId *big.Int, attestationsWeight *big.Int, blocksWeight *big.Int, syncWeight *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetPerformanceCoefficients(&_Bindings.TransactOpts, curveId, attestationsWeight, blocksWeight, syncWeight)
}

// SetPerformanceCoefficients is a paid mutator transaction binding the contract method 0xbf49bbde.
//
// Solidity: function setPerformanceCoefficients(uint256 curveId, uint256 attestationsWeight, uint256 blocksWeight, uint256 syncWeight) returns()
func (_Bindings *BindingsTransactorSession) SetPerformanceCoefficients(curveId *big.Int, attestationsWeight *big.Int, blocksWeight *big.Int, syncWeight *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetPerformanceCoefficients(&_Bindings.TransactOpts, curveId, attestationsWeight, blocksWeight, syncWeight)
}

// SetPerformanceLeewayData is a paid mutator transaction binding the contract method 0x726f0729.
//
// Solidity: function setPerformanceLeewayData(uint256 curveId, (uint256,uint256)[] data) returns()
func (_Bindings *BindingsTransactor) SetPerformanceLeewayData(opts *bind.TransactOpts, curveId *big.Int, data []ICSParametersRegistryKeyNumberValueInterval) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setPerformanceLeewayData", curveId, data)
}

// SetPerformanceLeewayData is a paid mutator transaction binding the contract method 0x726f0729.
//
// Solidity: function setPerformanceLeewayData(uint256 curveId, (uint256,uint256)[] data) returns()
func (_Bindings *BindingsSession) SetPerformanceLeewayData(curveId *big.Int, data []ICSParametersRegistryKeyNumberValueInterval) (*types.Transaction, error) {
	return _Bindings.Contract.SetPerformanceLeewayData(&_Bindings.TransactOpts, curveId, data)
}

// SetPerformanceLeewayData is a paid mutator transaction binding the contract method 0x726f0729.
//
// Solidity: function setPerformanceLeewayData(uint256 curveId, (uint256,uint256)[] data) returns()
func (_Bindings *BindingsTransactorSession) SetPerformanceLeewayData(curveId *big.Int, data []ICSParametersRegistryKeyNumberValueInterval) (*types.Transaction, error) {
	return _Bindings.Contract.SetPerformanceLeewayData(&_Bindings.TransactOpts, curveId, data)
}

// SetQueueConfig is a paid mutator transaction binding the contract method 0xd5d3687c.
//
// Solidity: function setQueueConfig(uint256 curveId, uint256 priority, uint256 maxDeposits) returns()
func (_Bindings *BindingsTransactor) SetQueueConfig(opts *bind.TransactOpts, curveId *big.Int, priority *big.Int, maxDeposits *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setQueueConfig", curveId, priority, maxDeposits)
}

// SetQueueConfig is a paid mutator transaction binding the contract method 0xd5d3687c.
//
// Solidity: function setQueueConfig(uint256 curveId, uint256 priority, uint256 maxDeposits) returns()
func (_Bindings *BindingsSession) SetQueueConfig(curveId *big.Int, priority *big.Int, maxDeposits *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetQueueConfig(&_Bindings.TransactOpts, curveId, priority, maxDeposits)
}

// SetQueueConfig is a paid mutator transaction binding the contract method 0xd5d3687c.
//
// Solidity: function setQueueConfig(uint256 curveId, uint256 priority, uint256 maxDeposits) returns()
func (_Bindings *BindingsTransactorSession) SetQueueConfig(curveId *big.Int, priority *big.Int, maxDeposits *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetQueueConfig(&_Bindings.TransactOpts, curveId, priority, maxDeposits)
}

// SetRewardShareData is a paid mutator transaction binding the contract method 0x9befc979.
//
// Solidity: function setRewardShareData(uint256 curveId, (uint256,uint256)[] data) returns()
func (_Bindings *BindingsTransactor) SetRewardShareData(opts *bind.TransactOpts, curveId *big.Int, data []ICSParametersRegistryKeyNumberValueInterval) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setRewardShareData", curveId, data)
}

// SetRewardShareData is a paid mutator transaction binding the contract method 0x9befc979.
//
// Solidity: function setRewardShareData(uint256 curveId, (uint256,uint256)[] data) returns()
func (_Bindings *BindingsSession) SetRewardShareData(curveId *big.Int, data []ICSParametersRegistryKeyNumberValueInterval) (*types.Transaction, error) {
	return _Bindings.Contract.SetRewardShareData(&_Bindings.TransactOpts, curveId, data)
}

// SetRewardShareData is a paid mutator transaction binding the contract method 0x9befc979.
//
// Solidity: function setRewardShareData(uint256 curveId, (uint256,uint256)[] data) returns()
func (_Bindings *BindingsTransactorSession) SetRewardShareData(curveId *big.Int, data []ICSParametersRegistryKeyNumberValueInterval) (*types.Transaction, error) {
	return _Bindings.Contract.SetRewardShareData(&_Bindings.TransactOpts, curveId, data)
}

// SetStrikesParams is a paid mutator transaction binding the contract method 0x67f73735.
//
// Solidity: function setStrikesParams(uint256 curveId, uint256 lifetime, uint256 threshold) returns()
func (_Bindings *BindingsTransactor) SetStrikesParams(opts *bind.TransactOpts, curveId *big.Int, lifetime *big.Int, threshold *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setStrikesParams", curveId, lifetime, threshold)
}

// SetStrikesParams is a paid mutator transaction binding the contract method 0x67f73735.
//
// Solidity: function setStrikesParams(uint256 curveId, uint256 lifetime, uint256 threshold) returns()
func (_Bindings *BindingsSession) SetStrikesParams(curveId *big.Int, lifetime *big.Int, threshold *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetStrikesParams(&_Bindings.TransactOpts, curveId, lifetime, threshold)
}

// SetStrikesParams is a paid mutator transaction binding the contract method 0x67f73735.
//
// Solidity: function setStrikesParams(uint256 curveId, uint256 lifetime, uint256 threshold) returns()
func (_Bindings *BindingsTransactorSession) SetStrikesParams(curveId *big.Int, lifetime *big.Int, threshold *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetStrikesParams(&_Bindings.TransactOpts, curveId, lifetime, threshold)
}

// UnsetAllowedExitDelay is a paid mutator transaction binding the contract method 0x9c5ebca1.
//
// Solidity: function unsetAllowedExitDelay(uint256 curveId) returns()
func (_Bindings *BindingsTransactor) UnsetAllowedExitDelay(opts *bind.TransactOpts, curveId *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "unsetAllowedExitDelay", curveId)
}

// UnsetAllowedExitDelay is a paid mutator transaction binding the contract method 0x9c5ebca1.
//
// Solidity: function unsetAllowedExitDelay(uint256 curveId) returns()
func (_Bindings *BindingsSession) UnsetAllowedExitDelay(curveId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.UnsetAllowedExitDelay(&_Bindings.TransactOpts, curveId)
}

// UnsetAllowedExitDelay is a paid mutator transaction binding the contract method 0x9c5ebca1.
//
// Solidity: function unsetAllowedExitDelay(uint256 curveId) returns()
func (_Bindings *BindingsTransactorSession) UnsetAllowedExitDelay(curveId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.UnsetAllowedExitDelay(&_Bindings.TransactOpts, curveId)
}

// UnsetBadPerformancePenalty is a paid mutator transaction binding the contract method 0x05284fb8.
//
// Solidity: function unsetBadPerformancePenalty(uint256 curveId) returns()
func (_Bindings *BindingsTransactor) UnsetBadPerformancePenalty(opts *bind.TransactOpts, curveId *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "unsetBadPerformancePenalty", curveId)
}

// UnsetBadPerformancePenalty is a paid mutator transaction binding the contract method 0x05284fb8.
//
// Solidity: function unsetBadPerformancePenalty(uint256 curveId) returns()
func (_Bindings *BindingsSession) UnsetBadPerformancePenalty(curveId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.UnsetBadPerformancePenalty(&_Bindings.TransactOpts, curveId)
}

// UnsetBadPerformancePenalty is a paid mutator transaction binding the contract method 0x05284fb8.
//
// Solidity: function unsetBadPerformancePenalty(uint256 curveId) returns()
func (_Bindings *BindingsTransactorSession) UnsetBadPerformancePenalty(curveId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.UnsetBadPerformancePenalty(&_Bindings.TransactOpts, curveId)
}

// UnsetElRewardsStealingAdditionalFine is a paid mutator transaction binding the contract method 0xce2bccdf.
//
// Solidity: function unsetElRewardsStealingAdditionalFine(uint256 curveId) returns()
func (_Bindings *BindingsTransactor) UnsetElRewardsStealingAdditionalFine(opts *bind.TransactOpts, curveId *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "unsetElRewardsStealingAdditionalFine", curveId)
}

// UnsetElRewardsStealingAdditionalFine is a paid mutator transaction binding the contract method 0xce2bccdf.
//
// Solidity: function unsetElRewardsStealingAdditionalFine(uint256 curveId) returns()
func (_Bindings *BindingsSession) UnsetElRewardsStealingAdditionalFine(curveId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.UnsetElRewardsStealingAdditionalFine(&_Bindings.TransactOpts, curveId)
}

// UnsetElRewardsStealingAdditionalFine is a paid mutator transaction binding the contract method 0xce2bccdf.
//
// Solidity: function unsetElRewardsStealingAdditionalFine(uint256 curveId) returns()
func (_Bindings *BindingsTransactorSession) UnsetElRewardsStealingAdditionalFine(curveId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.UnsetElRewardsStealingAdditionalFine(&_Bindings.TransactOpts, curveId)
}

// UnsetExitDelayPenalty is a paid mutator transaction binding the contract method 0xd665d195.
//
// Solidity: function unsetExitDelayPenalty(uint256 curveId) returns()
func (_Bindings *BindingsTransactor) UnsetExitDelayPenalty(opts *bind.TransactOpts, curveId *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "unsetExitDelayPenalty", curveId)
}

// UnsetExitDelayPenalty is a paid mutator transaction binding the contract method 0xd665d195.
//
// Solidity: function unsetExitDelayPenalty(uint256 curveId) returns()
func (_Bindings *BindingsSession) UnsetExitDelayPenalty(curveId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.UnsetExitDelayPenalty(&_Bindings.TransactOpts, curveId)
}

// UnsetExitDelayPenalty is a paid mutator transaction binding the contract method 0xd665d195.
//
// Solidity: function unsetExitDelayPenalty(uint256 curveId) returns()
func (_Bindings *BindingsTransactorSession) UnsetExitDelayPenalty(curveId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.UnsetExitDelayPenalty(&_Bindings.TransactOpts, curveId)
}

// UnsetKeyRemovalCharge is a paid mutator transaction binding the contract method 0xcf4e0bb2.
//
// Solidity: function unsetKeyRemovalCharge(uint256 curveId) returns()
func (_Bindings *BindingsTransactor) UnsetKeyRemovalCharge(opts *bind.TransactOpts, curveId *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "unsetKeyRemovalCharge", curveId)
}

// UnsetKeyRemovalCharge is a paid mutator transaction binding the contract method 0xcf4e0bb2.
//
// Solidity: function unsetKeyRemovalCharge(uint256 curveId) returns()
func (_Bindings *BindingsSession) UnsetKeyRemovalCharge(curveId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.UnsetKeyRemovalCharge(&_Bindings.TransactOpts, curveId)
}

// UnsetKeyRemovalCharge is a paid mutator transaction binding the contract method 0xcf4e0bb2.
//
// Solidity: function unsetKeyRemovalCharge(uint256 curveId) returns()
func (_Bindings *BindingsTransactorSession) UnsetKeyRemovalCharge(curveId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.UnsetKeyRemovalCharge(&_Bindings.TransactOpts, curveId)
}

// UnsetKeysLimit is a paid mutator transaction binding the contract method 0x384bfdf4.
//
// Solidity: function unsetKeysLimit(uint256 curveId) returns()
func (_Bindings *BindingsTransactor) UnsetKeysLimit(opts *bind.TransactOpts, curveId *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "unsetKeysLimit", curveId)
}

// UnsetKeysLimit is a paid mutator transaction binding the contract method 0x384bfdf4.
//
// Solidity: function unsetKeysLimit(uint256 curveId) returns()
func (_Bindings *BindingsSession) UnsetKeysLimit(curveId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.UnsetKeysLimit(&_Bindings.TransactOpts, curveId)
}

// UnsetKeysLimit is a paid mutator transaction binding the contract method 0x384bfdf4.
//
// Solidity: function unsetKeysLimit(uint256 curveId) returns()
func (_Bindings *BindingsTransactorSession) UnsetKeysLimit(curveId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.UnsetKeysLimit(&_Bindings.TransactOpts, curveId)
}

// UnsetMaxWithdrawalRequestFee is a paid mutator transaction binding the contract method 0x3c151e22.
//
// Solidity: function unsetMaxWithdrawalRequestFee(uint256 curveId) returns()
func (_Bindings *BindingsTransactor) UnsetMaxWithdrawalRequestFee(opts *bind.TransactOpts, curveId *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "unsetMaxWithdrawalRequestFee", curveId)
}

// UnsetMaxWithdrawalRequestFee is a paid mutator transaction binding the contract method 0x3c151e22.
//
// Solidity: function unsetMaxWithdrawalRequestFee(uint256 curveId) returns()
func (_Bindings *BindingsSession) UnsetMaxWithdrawalRequestFee(curveId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.UnsetMaxWithdrawalRequestFee(&_Bindings.TransactOpts, curveId)
}

// UnsetMaxWithdrawalRequestFee is a paid mutator transaction binding the contract method 0x3c151e22.
//
// Solidity: function unsetMaxWithdrawalRequestFee(uint256 curveId) returns()
func (_Bindings *BindingsTransactorSession) UnsetMaxWithdrawalRequestFee(curveId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.UnsetMaxWithdrawalRequestFee(&_Bindings.TransactOpts, curveId)
}

// UnsetPerformanceCoefficients is a paid mutator transaction binding the contract method 0x351da1e5.
//
// Solidity: function unsetPerformanceCoefficients(uint256 curveId) returns()
func (_Bindings *BindingsTransactor) UnsetPerformanceCoefficients(opts *bind.TransactOpts, curveId *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "unsetPerformanceCoefficients", curveId)
}

// UnsetPerformanceCoefficients is a paid mutator transaction binding the contract method 0x351da1e5.
//
// Solidity: function unsetPerformanceCoefficients(uint256 curveId) returns()
func (_Bindings *BindingsSession) UnsetPerformanceCoefficients(curveId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.UnsetPerformanceCoefficients(&_Bindings.TransactOpts, curveId)
}

// UnsetPerformanceCoefficients is a paid mutator transaction binding the contract method 0x351da1e5.
//
// Solidity: function unsetPerformanceCoefficients(uint256 curveId) returns()
func (_Bindings *BindingsTransactorSession) UnsetPerformanceCoefficients(curveId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.UnsetPerformanceCoefficients(&_Bindings.TransactOpts, curveId)
}

// UnsetPerformanceLeewayData is a paid mutator transaction binding the contract method 0x390530f4.
//
// Solidity: function unsetPerformanceLeewayData(uint256 curveId) returns()
func (_Bindings *BindingsTransactor) UnsetPerformanceLeewayData(opts *bind.TransactOpts, curveId *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "unsetPerformanceLeewayData", curveId)
}

// UnsetPerformanceLeewayData is a paid mutator transaction binding the contract method 0x390530f4.
//
// Solidity: function unsetPerformanceLeewayData(uint256 curveId) returns()
func (_Bindings *BindingsSession) UnsetPerformanceLeewayData(curveId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.UnsetPerformanceLeewayData(&_Bindings.TransactOpts, curveId)
}

// UnsetPerformanceLeewayData is a paid mutator transaction binding the contract method 0x390530f4.
//
// Solidity: function unsetPerformanceLeewayData(uint256 curveId) returns()
func (_Bindings *BindingsTransactorSession) UnsetPerformanceLeewayData(curveId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.UnsetPerformanceLeewayData(&_Bindings.TransactOpts, curveId)
}

// UnsetQueueConfig is a paid mutator transaction binding the contract method 0xc58e4bc7.
//
// Solidity: function unsetQueueConfig(uint256 curveId) returns()
func (_Bindings *BindingsTransactor) UnsetQueueConfig(opts *bind.TransactOpts, curveId *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "unsetQueueConfig", curveId)
}

// UnsetQueueConfig is a paid mutator transaction binding the contract method 0xc58e4bc7.
//
// Solidity: function unsetQueueConfig(uint256 curveId) returns()
func (_Bindings *BindingsSession) UnsetQueueConfig(curveId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.UnsetQueueConfig(&_Bindings.TransactOpts, curveId)
}

// UnsetQueueConfig is a paid mutator transaction binding the contract method 0xc58e4bc7.
//
// Solidity: function unsetQueueConfig(uint256 curveId) returns()
func (_Bindings *BindingsTransactorSession) UnsetQueueConfig(curveId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.UnsetQueueConfig(&_Bindings.TransactOpts, curveId)
}

// UnsetRewardShareData is a paid mutator transaction binding the contract method 0x8d9d8091.
//
// Solidity: function unsetRewardShareData(uint256 curveId) returns()
func (_Bindings *BindingsTransactor) UnsetRewardShareData(opts *bind.TransactOpts, curveId *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "unsetRewardShareData", curveId)
}

// UnsetRewardShareData is a paid mutator transaction binding the contract method 0x8d9d8091.
//
// Solidity: function unsetRewardShareData(uint256 curveId) returns()
func (_Bindings *BindingsSession) UnsetRewardShareData(curveId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.UnsetRewardShareData(&_Bindings.TransactOpts, curveId)
}

// UnsetRewardShareData is a paid mutator transaction binding the contract method 0x8d9d8091.
//
// Solidity: function unsetRewardShareData(uint256 curveId) returns()
func (_Bindings *BindingsTransactorSession) UnsetRewardShareData(curveId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.UnsetRewardShareData(&_Bindings.TransactOpts, curveId)
}

// UnsetStrikesParams is a paid mutator transaction binding the contract method 0x0b8f26e3.
//
// Solidity: function unsetStrikesParams(uint256 curveId) returns()
func (_Bindings *BindingsTransactor) UnsetStrikesParams(opts *bind.TransactOpts, curveId *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "unsetStrikesParams", curveId)
}

// UnsetStrikesParams is a paid mutator transaction binding the contract method 0x0b8f26e3.
//
// Solidity: function unsetStrikesParams(uint256 curveId) returns()
func (_Bindings *BindingsSession) UnsetStrikesParams(curveId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.UnsetStrikesParams(&_Bindings.TransactOpts, curveId)
}

// UnsetStrikesParams is a paid mutator transaction binding the contract method 0x0b8f26e3.
//
// Solidity: function unsetStrikesParams(uint256 curveId) returns()
func (_Bindings *BindingsTransactorSession) UnsetStrikesParams(curveId *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.UnsetStrikesParams(&_Bindings.TransactOpts, curveId)
}

// BindingsAllowedExitDelaySetIterator is returned from FilterAllowedExitDelaySet and is used to iterate over the raw logs and unpacked data for AllowedExitDelaySet events raised by the Bindings contract.
type BindingsAllowedExitDelaySetIterator struct {
	Event *BindingsAllowedExitDelaySet // Event containing the contract specifics and raw log

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
func (it *BindingsAllowedExitDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsAllowedExitDelaySet)
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
		it.Event = new(BindingsAllowedExitDelaySet)
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
func (it *BindingsAllowedExitDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsAllowedExitDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsAllowedExitDelaySet represents a AllowedExitDelaySet event raised by the Bindings contract.
type BindingsAllowedExitDelaySet struct {
	CurveId *big.Int
	Delay   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAllowedExitDelaySet is a free log retrieval operation binding the contract event 0x1b550eacd68ffc9d5aa86266bce448ee53e88cb7d8eeca7eb8de08f81e730f4b.
//
// Solidity: event AllowedExitDelaySet(uint256 indexed curveId, uint256 delay)
func (_Bindings *BindingsFilterer) FilterAllowedExitDelaySet(opts *bind.FilterOpts, curveId []*big.Int) (*BindingsAllowedExitDelaySetIterator, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "AllowedExitDelaySet", curveIdRule)
	if err != nil {
		return nil, err
	}
	return &BindingsAllowedExitDelaySetIterator{contract: _Bindings.contract, event: "AllowedExitDelaySet", logs: logs, sub: sub}, nil
}

// WatchAllowedExitDelaySet is a free log subscription operation binding the contract event 0x1b550eacd68ffc9d5aa86266bce448ee53e88cb7d8eeca7eb8de08f81e730f4b.
//
// Solidity: event AllowedExitDelaySet(uint256 indexed curveId, uint256 delay)
func (_Bindings *BindingsFilterer) WatchAllowedExitDelaySet(opts *bind.WatchOpts, sink chan<- *BindingsAllowedExitDelaySet, curveId []*big.Int) (event.Subscription, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "AllowedExitDelaySet", curveIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsAllowedExitDelaySet)
				if err := _Bindings.contract.UnpackLog(event, "AllowedExitDelaySet", log); err != nil {
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

// ParseAllowedExitDelaySet is a log parse operation binding the contract event 0x1b550eacd68ffc9d5aa86266bce448ee53e88cb7d8eeca7eb8de08f81e730f4b.
//
// Solidity: event AllowedExitDelaySet(uint256 indexed curveId, uint256 delay)
func (_Bindings *BindingsFilterer) ParseAllowedExitDelaySet(log types.Log) (*BindingsAllowedExitDelaySet, error) {
	event := new(BindingsAllowedExitDelaySet)
	if err := _Bindings.contract.UnpackLog(event, "AllowedExitDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsAllowedExitDelayUnsetIterator is returned from FilterAllowedExitDelayUnset and is used to iterate over the raw logs and unpacked data for AllowedExitDelayUnset events raised by the Bindings contract.
type BindingsAllowedExitDelayUnsetIterator struct {
	Event *BindingsAllowedExitDelayUnset // Event containing the contract specifics and raw log

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
func (it *BindingsAllowedExitDelayUnsetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsAllowedExitDelayUnset)
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
		it.Event = new(BindingsAllowedExitDelayUnset)
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
func (it *BindingsAllowedExitDelayUnsetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsAllowedExitDelayUnsetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsAllowedExitDelayUnset represents a AllowedExitDelayUnset event raised by the Bindings contract.
type BindingsAllowedExitDelayUnset struct {
	CurveId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAllowedExitDelayUnset is a free log retrieval operation binding the contract event 0xc5aa054c2c0f4bb3c866f44f8e5062c9a09da86e371118ea7eb9b84848f0ffde.
//
// Solidity: event AllowedExitDelayUnset(uint256 indexed curveId)
func (_Bindings *BindingsFilterer) FilterAllowedExitDelayUnset(opts *bind.FilterOpts, curveId []*big.Int) (*BindingsAllowedExitDelayUnsetIterator, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "AllowedExitDelayUnset", curveIdRule)
	if err != nil {
		return nil, err
	}
	return &BindingsAllowedExitDelayUnsetIterator{contract: _Bindings.contract, event: "AllowedExitDelayUnset", logs: logs, sub: sub}, nil
}

// WatchAllowedExitDelayUnset is a free log subscription operation binding the contract event 0xc5aa054c2c0f4bb3c866f44f8e5062c9a09da86e371118ea7eb9b84848f0ffde.
//
// Solidity: event AllowedExitDelayUnset(uint256 indexed curveId)
func (_Bindings *BindingsFilterer) WatchAllowedExitDelayUnset(opts *bind.WatchOpts, sink chan<- *BindingsAllowedExitDelayUnset, curveId []*big.Int) (event.Subscription, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "AllowedExitDelayUnset", curveIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsAllowedExitDelayUnset)
				if err := _Bindings.contract.UnpackLog(event, "AllowedExitDelayUnset", log); err != nil {
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

// ParseAllowedExitDelayUnset is a log parse operation binding the contract event 0xc5aa054c2c0f4bb3c866f44f8e5062c9a09da86e371118ea7eb9b84848f0ffde.
//
// Solidity: event AllowedExitDelayUnset(uint256 indexed curveId)
func (_Bindings *BindingsFilterer) ParseAllowedExitDelayUnset(log types.Log) (*BindingsAllowedExitDelayUnset, error) {
	event := new(BindingsAllowedExitDelayUnset)
	if err := _Bindings.contract.UnpackLog(event, "AllowedExitDelayUnset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsBadPerformancePenaltySetIterator is returned from FilterBadPerformancePenaltySet and is used to iterate over the raw logs and unpacked data for BadPerformancePenaltySet events raised by the Bindings contract.
type BindingsBadPerformancePenaltySetIterator struct {
	Event *BindingsBadPerformancePenaltySet // Event containing the contract specifics and raw log

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
func (it *BindingsBadPerformancePenaltySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsBadPerformancePenaltySet)
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
		it.Event = new(BindingsBadPerformancePenaltySet)
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
func (it *BindingsBadPerformancePenaltySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsBadPerformancePenaltySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsBadPerformancePenaltySet represents a BadPerformancePenaltySet event raised by the Bindings contract.
type BindingsBadPerformancePenaltySet struct {
	CurveId *big.Int
	Penalty *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBadPerformancePenaltySet is a free log retrieval operation binding the contract event 0x20eae15fa694728a1eff8e6907a511dd23bd4bd04f0d73e8f26e61f9cafed751.
//
// Solidity: event BadPerformancePenaltySet(uint256 indexed curveId, uint256 penalty)
func (_Bindings *BindingsFilterer) FilterBadPerformancePenaltySet(opts *bind.FilterOpts, curveId []*big.Int) (*BindingsBadPerformancePenaltySetIterator, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "BadPerformancePenaltySet", curveIdRule)
	if err != nil {
		return nil, err
	}
	return &BindingsBadPerformancePenaltySetIterator{contract: _Bindings.contract, event: "BadPerformancePenaltySet", logs: logs, sub: sub}, nil
}

// WatchBadPerformancePenaltySet is a free log subscription operation binding the contract event 0x20eae15fa694728a1eff8e6907a511dd23bd4bd04f0d73e8f26e61f9cafed751.
//
// Solidity: event BadPerformancePenaltySet(uint256 indexed curveId, uint256 penalty)
func (_Bindings *BindingsFilterer) WatchBadPerformancePenaltySet(opts *bind.WatchOpts, sink chan<- *BindingsBadPerformancePenaltySet, curveId []*big.Int) (event.Subscription, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "BadPerformancePenaltySet", curveIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsBadPerformancePenaltySet)
				if err := _Bindings.contract.UnpackLog(event, "BadPerformancePenaltySet", log); err != nil {
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

// ParseBadPerformancePenaltySet is a log parse operation binding the contract event 0x20eae15fa694728a1eff8e6907a511dd23bd4bd04f0d73e8f26e61f9cafed751.
//
// Solidity: event BadPerformancePenaltySet(uint256 indexed curveId, uint256 penalty)
func (_Bindings *BindingsFilterer) ParseBadPerformancePenaltySet(log types.Log) (*BindingsBadPerformancePenaltySet, error) {
	event := new(BindingsBadPerformancePenaltySet)
	if err := _Bindings.contract.UnpackLog(event, "BadPerformancePenaltySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsBadPerformancePenaltyUnsetIterator is returned from FilterBadPerformancePenaltyUnset and is used to iterate over the raw logs and unpacked data for BadPerformancePenaltyUnset events raised by the Bindings contract.
type BindingsBadPerformancePenaltyUnsetIterator struct {
	Event *BindingsBadPerformancePenaltyUnset // Event containing the contract specifics and raw log

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
func (it *BindingsBadPerformancePenaltyUnsetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsBadPerformancePenaltyUnset)
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
		it.Event = new(BindingsBadPerformancePenaltyUnset)
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
func (it *BindingsBadPerformancePenaltyUnsetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsBadPerformancePenaltyUnsetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsBadPerformancePenaltyUnset represents a BadPerformancePenaltyUnset event raised by the Bindings contract.
type BindingsBadPerformancePenaltyUnset struct {
	CurveId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBadPerformancePenaltyUnset is a free log retrieval operation binding the contract event 0x3cea17c78f077924551bbcc01c140208cf087e3736225cd921e36d140cf29757.
//
// Solidity: event BadPerformancePenaltyUnset(uint256 indexed curveId)
func (_Bindings *BindingsFilterer) FilterBadPerformancePenaltyUnset(opts *bind.FilterOpts, curveId []*big.Int) (*BindingsBadPerformancePenaltyUnsetIterator, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "BadPerformancePenaltyUnset", curveIdRule)
	if err != nil {
		return nil, err
	}
	return &BindingsBadPerformancePenaltyUnsetIterator{contract: _Bindings.contract, event: "BadPerformancePenaltyUnset", logs: logs, sub: sub}, nil
}

// WatchBadPerformancePenaltyUnset is a free log subscription operation binding the contract event 0x3cea17c78f077924551bbcc01c140208cf087e3736225cd921e36d140cf29757.
//
// Solidity: event BadPerformancePenaltyUnset(uint256 indexed curveId)
func (_Bindings *BindingsFilterer) WatchBadPerformancePenaltyUnset(opts *bind.WatchOpts, sink chan<- *BindingsBadPerformancePenaltyUnset, curveId []*big.Int) (event.Subscription, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "BadPerformancePenaltyUnset", curveIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsBadPerformancePenaltyUnset)
				if err := _Bindings.contract.UnpackLog(event, "BadPerformancePenaltyUnset", log); err != nil {
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

// ParseBadPerformancePenaltyUnset is a log parse operation binding the contract event 0x3cea17c78f077924551bbcc01c140208cf087e3736225cd921e36d140cf29757.
//
// Solidity: event BadPerformancePenaltyUnset(uint256 indexed curveId)
func (_Bindings *BindingsFilterer) ParseBadPerformancePenaltyUnset(log types.Log) (*BindingsBadPerformancePenaltyUnset, error) {
	event := new(BindingsBadPerformancePenaltyUnset)
	if err := _Bindings.contract.UnpackLog(event, "BadPerformancePenaltyUnset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsDefaultAllowedExitDelaySetIterator is returned from FilterDefaultAllowedExitDelaySet and is used to iterate over the raw logs and unpacked data for DefaultAllowedExitDelaySet events raised by the Bindings contract.
type BindingsDefaultAllowedExitDelaySetIterator struct {
	Event *BindingsDefaultAllowedExitDelaySet // Event containing the contract specifics and raw log

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
func (it *BindingsDefaultAllowedExitDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsDefaultAllowedExitDelaySet)
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
		it.Event = new(BindingsDefaultAllowedExitDelaySet)
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
func (it *BindingsDefaultAllowedExitDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsDefaultAllowedExitDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsDefaultAllowedExitDelaySet represents a DefaultAllowedExitDelaySet event raised by the Bindings contract.
type BindingsDefaultAllowedExitDelaySet struct {
	Delay *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDefaultAllowedExitDelaySet is a free log retrieval operation binding the contract event 0x63975a918e6e40eb014f8e933d55d7f120a2bd111e173360153b357249727c5f.
//
// Solidity: event DefaultAllowedExitDelaySet(uint256 delay)
func (_Bindings *BindingsFilterer) FilterDefaultAllowedExitDelaySet(opts *bind.FilterOpts) (*BindingsDefaultAllowedExitDelaySetIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "DefaultAllowedExitDelaySet")
	if err != nil {
		return nil, err
	}
	return &BindingsDefaultAllowedExitDelaySetIterator{contract: _Bindings.contract, event: "DefaultAllowedExitDelaySet", logs: logs, sub: sub}, nil
}

// WatchDefaultAllowedExitDelaySet is a free log subscription operation binding the contract event 0x63975a918e6e40eb014f8e933d55d7f120a2bd111e173360153b357249727c5f.
//
// Solidity: event DefaultAllowedExitDelaySet(uint256 delay)
func (_Bindings *BindingsFilterer) WatchDefaultAllowedExitDelaySet(opts *bind.WatchOpts, sink chan<- *BindingsDefaultAllowedExitDelaySet) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "DefaultAllowedExitDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsDefaultAllowedExitDelaySet)
				if err := _Bindings.contract.UnpackLog(event, "DefaultAllowedExitDelaySet", log); err != nil {
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

// ParseDefaultAllowedExitDelaySet is a log parse operation binding the contract event 0x63975a918e6e40eb014f8e933d55d7f120a2bd111e173360153b357249727c5f.
//
// Solidity: event DefaultAllowedExitDelaySet(uint256 delay)
func (_Bindings *BindingsFilterer) ParseDefaultAllowedExitDelaySet(log types.Log) (*BindingsDefaultAllowedExitDelaySet, error) {
	event := new(BindingsDefaultAllowedExitDelaySet)
	if err := _Bindings.contract.UnpackLog(event, "DefaultAllowedExitDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsDefaultBadPerformancePenaltySetIterator is returned from FilterDefaultBadPerformancePenaltySet and is used to iterate over the raw logs and unpacked data for DefaultBadPerformancePenaltySet events raised by the Bindings contract.
type BindingsDefaultBadPerformancePenaltySetIterator struct {
	Event *BindingsDefaultBadPerformancePenaltySet // Event containing the contract specifics and raw log

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
func (it *BindingsDefaultBadPerformancePenaltySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsDefaultBadPerformancePenaltySet)
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
		it.Event = new(BindingsDefaultBadPerformancePenaltySet)
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
func (it *BindingsDefaultBadPerformancePenaltySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsDefaultBadPerformancePenaltySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsDefaultBadPerformancePenaltySet represents a DefaultBadPerformancePenaltySet event raised by the Bindings contract.
type BindingsDefaultBadPerformancePenaltySet struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDefaultBadPerformancePenaltySet is a free log retrieval operation binding the contract event 0xb8322c3f829eef7cf1ee51739884859c5a7a3421a9069a2ea0eb72219823094d.
//
// Solidity: event DefaultBadPerformancePenaltySet(uint256 value)
func (_Bindings *BindingsFilterer) FilterDefaultBadPerformancePenaltySet(opts *bind.FilterOpts) (*BindingsDefaultBadPerformancePenaltySetIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "DefaultBadPerformancePenaltySet")
	if err != nil {
		return nil, err
	}
	return &BindingsDefaultBadPerformancePenaltySetIterator{contract: _Bindings.contract, event: "DefaultBadPerformancePenaltySet", logs: logs, sub: sub}, nil
}

// WatchDefaultBadPerformancePenaltySet is a free log subscription operation binding the contract event 0xb8322c3f829eef7cf1ee51739884859c5a7a3421a9069a2ea0eb72219823094d.
//
// Solidity: event DefaultBadPerformancePenaltySet(uint256 value)
func (_Bindings *BindingsFilterer) WatchDefaultBadPerformancePenaltySet(opts *bind.WatchOpts, sink chan<- *BindingsDefaultBadPerformancePenaltySet) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "DefaultBadPerformancePenaltySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsDefaultBadPerformancePenaltySet)
				if err := _Bindings.contract.UnpackLog(event, "DefaultBadPerformancePenaltySet", log); err != nil {
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

// ParseDefaultBadPerformancePenaltySet is a log parse operation binding the contract event 0xb8322c3f829eef7cf1ee51739884859c5a7a3421a9069a2ea0eb72219823094d.
//
// Solidity: event DefaultBadPerformancePenaltySet(uint256 value)
func (_Bindings *BindingsFilterer) ParseDefaultBadPerformancePenaltySet(log types.Log) (*BindingsDefaultBadPerformancePenaltySet, error) {
	event := new(BindingsDefaultBadPerformancePenaltySet)
	if err := _Bindings.contract.UnpackLog(event, "DefaultBadPerformancePenaltySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsDefaultElRewardsStealingAdditionalFineSetIterator is returned from FilterDefaultElRewardsStealingAdditionalFineSet and is used to iterate over the raw logs and unpacked data for DefaultElRewardsStealingAdditionalFineSet events raised by the Bindings contract.
type BindingsDefaultElRewardsStealingAdditionalFineSetIterator struct {
	Event *BindingsDefaultElRewardsStealingAdditionalFineSet // Event containing the contract specifics and raw log

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
func (it *BindingsDefaultElRewardsStealingAdditionalFineSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsDefaultElRewardsStealingAdditionalFineSet)
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
		it.Event = new(BindingsDefaultElRewardsStealingAdditionalFineSet)
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
func (it *BindingsDefaultElRewardsStealingAdditionalFineSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsDefaultElRewardsStealingAdditionalFineSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsDefaultElRewardsStealingAdditionalFineSet represents a DefaultElRewardsStealingAdditionalFineSet event raised by the Bindings contract.
type BindingsDefaultElRewardsStealingAdditionalFineSet struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDefaultElRewardsStealingAdditionalFineSet is a free log retrieval operation binding the contract event 0x058c14a22aeef4bebb83ecbff4730ad324e9b41ae29b17a59fa814c4e066d01d.
//
// Solidity: event DefaultElRewardsStealingAdditionalFineSet(uint256 value)
func (_Bindings *BindingsFilterer) FilterDefaultElRewardsStealingAdditionalFineSet(opts *bind.FilterOpts) (*BindingsDefaultElRewardsStealingAdditionalFineSetIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "DefaultElRewardsStealingAdditionalFineSet")
	if err != nil {
		return nil, err
	}
	return &BindingsDefaultElRewardsStealingAdditionalFineSetIterator{contract: _Bindings.contract, event: "DefaultElRewardsStealingAdditionalFineSet", logs: logs, sub: sub}, nil
}

// WatchDefaultElRewardsStealingAdditionalFineSet is a free log subscription operation binding the contract event 0x058c14a22aeef4bebb83ecbff4730ad324e9b41ae29b17a59fa814c4e066d01d.
//
// Solidity: event DefaultElRewardsStealingAdditionalFineSet(uint256 value)
func (_Bindings *BindingsFilterer) WatchDefaultElRewardsStealingAdditionalFineSet(opts *bind.WatchOpts, sink chan<- *BindingsDefaultElRewardsStealingAdditionalFineSet) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "DefaultElRewardsStealingAdditionalFineSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsDefaultElRewardsStealingAdditionalFineSet)
				if err := _Bindings.contract.UnpackLog(event, "DefaultElRewardsStealingAdditionalFineSet", log); err != nil {
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

// ParseDefaultElRewardsStealingAdditionalFineSet is a log parse operation binding the contract event 0x058c14a22aeef4bebb83ecbff4730ad324e9b41ae29b17a59fa814c4e066d01d.
//
// Solidity: event DefaultElRewardsStealingAdditionalFineSet(uint256 value)
func (_Bindings *BindingsFilterer) ParseDefaultElRewardsStealingAdditionalFineSet(log types.Log) (*BindingsDefaultElRewardsStealingAdditionalFineSet, error) {
	event := new(BindingsDefaultElRewardsStealingAdditionalFineSet)
	if err := _Bindings.contract.UnpackLog(event, "DefaultElRewardsStealingAdditionalFineSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsDefaultExitDelayPenaltySetIterator is returned from FilterDefaultExitDelayPenaltySet and is used to iterate over the raw logs and unpacked data for DefaultExitDelayPenaltySet events raised by the Bindings contract.
type BindingsDefaultExitDelayPenaltySetIterator struct {
	Event *BindingsDefaultExitDelayPenaltySet // Event containing the contract specifics and raw log

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
func (it *BindingsDefaultExitDelayPenaltySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsDefaultExitDelayPenaltySet)
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
		it.Event = new(BindingsDefaultExitDelayPenaltySet)
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
func (it *BindingsDefaultExitDelayPenaltySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsDefaultExitDelayPenaltySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsDefaultExitDelayPenaltySet represents a DefaultExitDelayPenaltySet event raised by the Bindings contract.
type BindingsDefaultExitDelayPenaltySet struct {
	Penalty *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDefaultExitDelayPenaltySet is a free log retrieval operation binding the contract event 0x2568cbc15cf2d9e34de76c131b8fd30f3e0690a8928c027e3ae8f6adf792967b.
//
// Solidity: event DefaultExitDelayPenaltySet(uint256 penalty)
func (_Bindings *BindingsFilterer) FilterDefaultExitDelayPenaltySet(opts *bind.FilterOpts) (*BindingsDefaultExitDelayPenaltySetIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "DefaultExitDelayPenaltySet")
	if err != nil {
		return nil, err
	}
	return &BindingsDefaultExitDelayPenaltySetIterator{contract: _Bindings.contract, event: "DefaultExitDelayPenaltySet", logs: logs, sub: sub}, nil
}

// WatchDefaultExitDelayPenaltySet is a free log subscription operation binding the contract event 0x2568cbc15cf2d9e34de76c131b8fd30f3e0690a8928c027e3ae8f6adf792967b.
//
// Solidity: event DefaultExitDelayPenaltySet(uint256 penalty)
func (_Bindings *BindingsFilterer) WatchDefaultExitDelayPenaltySet(opts *bind.WatchOpts, sink chan<- *BindingsDefaultExitDelayPenaltySet) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "DefaultExitDelayPenaltySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsDefaultExitDelayPenaltySet)
				if err := _Bindings.contract.UnpackLog(event, "DefaultExitDelayPenaltySet", log); err != nil {
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

// ParseDefaultExitDelayPenaltySet is a log parse operation binding the contract event 0x2568cbc15cf2d9e34de76c131b8fd30f3e0690a8928c027e3ae8f6adf792967b.
//
// Solidity: event DefaultExitDelayPenaltySet(uint256 penalty)
func (_Bindings *BindingsFilterer) ParseDefaultExitDelayPenaltySet(log types.Log) (*BindingsDefaultExitDelayPenaltySet, error) {
	event := new(BindingsDefaultExitDelayPenaltySet)
	if err := _Bindings.contract.UnpackLog(event, "DefaultExitDelayPenaltySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsDefaultKeyRemovalChargeSetIterator is returned from FilterDefaultKeyRemovalChargeSet and is used to iterate over the raw logs and unpacked data for DefaultKeyRemovalChargeSet events raised by the Bindings contract.
type BindingsDefaultKeyRemovalChargeSetIterator struct {
	Event *BindingsDefaultKeyRemovalChargeSet // Event containing the contract specifics and raw log

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
func (it *BindingsDefaultKeyRemovalChargeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsDefaultKeyRemovalChargeSet)
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
		it.Event = new(BindingsDefaultKeyRemovalChargeSet)
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
func (it *BindingsDefaultKeyRemovalChargeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsDefaultKeyRemovalChargeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsDefaultKeyRemovalChargeSet represents a DefaultKeyRemovalChargeSet event raised by the Bindings contract.
type BindingsDefaultKeyRemovalChargeSet struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDefaultKeyRemovalChargeSet is a free log retrieval operation binding the contract event 0x7f44d94c4ba99a62f393d42a985fed36a76e6b1a2cb378ea18180e1cdb27c9ef.
//
// Solidity: event DefaultKeyRemovalChargeSet(uint256 value)
func (_Bindings *BindingsFilterer) FilterDefaultKeyRemovalChargeSet(opts *bind.FilterOpts) (*BindingsDefaultKeyRemovalChargeSetIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "DefaultKeyRemovalChargeSet")
	if err != nil {
		return nil, err
	}
	return &BindingsDefaultKeyRemovalChargeSetIterator{contract: _Bindings.contract, event: "DefaultKeyRemovalChargeSet", logs: logs, sub: sub}, nil
}

// WatchDefaultKeyRemovalChargeSet is a free log subscription operation binding the contract event 0x7f44d94c4ba99a62f393d42a985fed36a76e6b1a2cb378ea18180e1cdb27c9ef.
//
// Solidity: event DefaultKeyRemovalChargeSet(uint256 value)
func (_Bindings *BindingsFilterer) WatchDefaultKeyRemovalChargeSet(opts *bind.WatchOpts, sink chan<- *BindingsDefaultKeyRemovalChargeSet) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "DefaultKeyRemovalChargeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsDefaultKeyRemovalChargeSet)
				if err := _Bindings.contract.UnpackLog(event, "DefaultKeyRemovalChargeSet", log); err != nil {
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

// ParseDefaultKeyRemovalChargeSet is a log parse operation binding the contract event 0x7f44d94c4ba99a62f393d42a985fed36a76e6b1a2cb378ea18180e1cdb27c9ef.
//
// Solidity: event DefaultKeyRemovalChargeSet(uint256 value)
func (_Bindings *BindingsFilterer) ParseDefaultKeyRemovalChargeSet(log types.Log) (*BindingsDefaultKeyRemovalChargeSet, error) {
	event := new(BindingsDefaultKeyRemovalChargeSet)
	if err := _Bindings.contract.UnpackLog(event, "DefaultKeyRemovalChargeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsDefaultKeysLimitSetIterator is returned from FilterDefaultKeysLimitSet and is used to iterate over the raw logs and unpacked data for DefaultKeysLimitSet events raised by the Bindings contract.
type BindingsDefaultKeysLimitSetIterator struct {
	Event *BindingsDefaultKeysLimitSet // Event containing the contract specifics and raw log

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
func (it *BindingsDefaultKeysLimitSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsDefaultKeysLimitSet)
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
		it.Event = new(BindingsDefaultKeysLimitSet)
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
func (it *BindingsDefaultKeysLimitSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsDefaultKeysLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsDefaultKeysLimitSet represents a DefaultKeysLimitSet event raised by the Bindings contract.
type BindingsDefaultKeysLimitSet struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDefaultKeysLimitSet is a free log retrieval operation binding the contract event 0xdb0c52b8ad6c5b56b45a406e6669dec9bf62c73943389e5b4d55bd59754fb38a.
//
// Solidity: event DefaultKeysLimitSet(uint256 value)
func (_Bindings *BindingsFilterer) FilterDefaultKeysLimitSet(opts *bind.FilterOpts) (*BindingsDefaultKeysLimitSetIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "DefaultKeysLimitSet")
	if err != nil {
		return nil, err
	}
	return &BindingsDefaultKeysLimitSetIterator{contract: _Bindings.contract, event: "DefaultKeysLimitSet", logs: logs, sub: sub}, nil
}

// WatchDefaultKeysLimitSet is a free log subscription operation binding the contract event 0xdb0c52b8ad6c5b56b45a406e6669dec9bf62c73943389e5b4d55bd59754fb38a.
//
// Solidity: event DefaultKeysLimitSet(uint256 value)
func (_Bindings *BindingsFilterer) WatchDefaultKeysLimitSet(opts *bind.WatchOpts, sink chan<- *BindingsDefaultKeysLimitSet) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "DefaultKeysLimitSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsDefaultKeysLimitSet)
				if err := _Bindings.contract.UnpackLog(event, "DefaultKeysLimitSet", log); err != nil {
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

// ParseDefaultKeysLimitSet is a log parse operation binding the contract event 0xdb0c52b8ad6c5b56b45a406e6669dec9bf62c73943389e5b4d55bd59754fb38a.
//
// Solidity: event DefaultKeysLimitSet(uint256 value)
func (_Bindings *BindingsFilterer) ParseDefaultKeysLimitSet(log types.Log) (*BindingsDefaultKeysLimitSet, error) {
	event := new(BindingsDefaultKeysLimitSet)
	if err := _Bindings.contract.UnpackLog(event, "DefaultKeysLimitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsDefaultMaxWithdrawalRequestFeeSetIterator is returned from FilterDefaultMaxWithdrawalRequestFeeSet and is used to iterate over the raw logs and unpacked data for DefaultMaxWithdrawalRequestFeeSet events raised by the Bindings contract.
type BindingsDefaultMaxWithdrawalRequestFeeSetIterator struct {
	Event *BindingsDefaultMaxWithdrawalRequestFeeSet // Event containing the contract specifics and raw log

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
func (it *BindingsDefaultMaxWithdrawalRequestFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsDefaultMaxWithdrawalRequestFeeSet)
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
		it.Event = new(BindingsDefaultMaxWithdrawalRequestFeeSet)
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
func (it *BindingsDefaultMaxWithdrawalRequestFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsDefaultMaxWithdrawalRequestFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsDefaultMaxWithdrawalRequestFeeSet represents a DefaultMaxWithdrawalRequestFeeSet event raised by the Bindings contract.
type BindingsDefaultMaxWithdrawalRequestFeeSet struct {
	Fee *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDefaultMaxWithdrawalRequestFeeSet is a free log retrieval operation binding the contract event 0xdd10741741bd981f5126e08cd0b47bdc8a737bd3d031dda270fd22ddccabbc8c.
//
// Solidity: event DefaultMaxWithdrawalRequestFeeSet(uint256 fee)
func (_Bindings *BindingsFilterer) FilterDefaultMaxWithdrawalRequestFeeSet(opts *bind.FilterOpts) (*BindingsDefaultMaxWithdrawalRequestFeeSetIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "DefaultMaxWithdrawalRequestFeeSet")
	if err != nil {
		return nil, err
	}
	return &BindingsDefaultMaxWithdrawalRequestFeeSetIterator{contract: _Bindings.contract, event: "DefaultMaxWithdrawalRequestFeeSet", logs: logs, sub: sub}, nil
}

// WatchDefaultMaxWithdrawalRequestFeeSet is a free log subscription operation binding the contract event 0xdd10741741bd981f5126e08cd0b47bdc8a737bd3d031dda270fd22ddccabbc8c.
//
// Solidity: event DefaultMaxWithdrawalRequestFeeSet(uint256 fee)
func (_Bindings *BindingsFilterer) WatchDefaultMaxWithdrawalRequestFeeSet(opts *bind.WatchOpts, sink chan<- *BindingsDefaultMaxWithdrawalRequestFeeSet) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "DefaultMaxWithdrawalRequestFeeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsDefaultMaxWithdrawalRequestFeeSet)
				if err := _Bindings.contract.UnpackLog(event, "DefaultMaxWithdrawalRequestFeeSet", log); err != nil {
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

// ParseDefaultMaxWithdrawalRequestFeeSet is a log parse operation binding the contract event 0xdd10741741bd981f5126e08cd0b47bdc8a737bd3d031dda270fd22ddccabbc8c.
//
// Solidity: event DefaultMaxWithdrawalRequestFeeSet(uint256 fee)
func (_Bindings *BindingsFilterer) ParseDefaultMaxWithdrawalRequestFeeSet(log types.Log) (*BindingsDefaultMaxWithdrawalRequestFeeSet, error) {
	event := new(BindingsDefaultMaxWithdrawalRequestFeeSet)
	if err := _Bindings.contract.UnpackLog(event, "DefaultMaxWithdrawalRequestFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsDefaultPerformanceCoefficientsSetIterator is returned from FilterDefaultPerformanceCoefficientsSet and is used to iterate over the raw logs and unpacked data for DefaultPerformanceCoefficientsSet events raised by the Bindings contract.
type BindingsDefaultPerformanceCoefficientsSetIterator struct {
	Event *BindingsDefaultPerformanceCoefficientsSet // Event containing the contract specifics and raw log

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
func (it *BindingsDefaultPerformanceCoefficientsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsDefaultPerformanceCoefficientsSet)
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
		it.Event = new(BindingsDefaultPerformanceCoefficientsSet)
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
func (it *BindingsDefaultPerformanceCoefficientsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsDefaultPerformanceCoefficientsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsDefaultPerformanceCoefficientsSet represents a DefaultPerformanceCoefficientsSet event raised by the Bindings contract.
type BindingsDefaultPerformanceCoefficientsSet struct {
	AttestationsWeight *big.Int
	BlocksWeight       *big.Int
	SyncWeight         *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDefaultPerformanceCoefficientsSet is a free log retrieval operation binding the contract event 0x8458290c9d3d495665598cb2ef88cb10338287228ccea69e96166ad30c7d2a95.
//
// Solidity: event DefaultPerformanceCoefficientsSet(uint256 attestationsWeight, uint256 blocksWeight, uint256 syncWeight)
func (_Bindings *BindingsFilterer) FilterDefaultPerformanceCoefficientsSet(opts *bind.FilterOpts) (*BindingsDefaultPerformanceCoefficientsSetIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "DefaultPerformanceCoefficientsSet")
	if err != nil {
		return nil, err
	}
	return &BindingsDefaultPerformanceCoefficientsSetIterator{contract: _Bindings.contract, event: "DefaultPerformanceCoefficientsSet", logs: logs, sub: sub}, nil
}

// WatchDefaultPerformanceCoefficientsSet is a free log subscription operation binding the contract event 0x8458290c9d3d495665598cb2ef88cb10338287228ccea69e96166ad30c7d2a95.
//
// Solidity: event DefaultPerformanceCoefficientsSet(uint256 attestationsWeight, uint256 blocksWeight, uint256 syncWeight)
func (_Bindings *BindingsFilterer) WatchDefaultPerformanceCoefficientsSet(opts *bind.WatchOpts, sink chan<- *BindingsDefaultPerformanceCoefficientsSet) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "DefaultPerformanceCoefficientsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsDefaultPerformanceCoefficientsSet)
				if err := _Bindings.contract.UnpackLog(event, "DefaultPerformanceCoefficientsSet", log); err != nil {
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

// ParseDefaultPerformanceCoefficientsSet is a log parse operation binding the contract event 0x8458290c9d3d495665598cb2ef88cb10338287228ccea69e96166ad30c7d2a95.
//
// Solidity: event DefaultPerformanceCoefficientsSet(uint256 attestationsWeight, uint256 blocksWeight, uint256 syncWeight)
func (_Bindings *BindingsFilterer) ParseDefaultPerformanceCoefficientsSet(log types.Log) (*BindingsDefaultPerformanceCoefficientsSet, error) {
	event := new(BindingsDefaultPerformanceCoefficientsSet)
	if err := _Bindings.contract.UnpackLog(event, "DefaultPerformanceCoefficientsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsDefaultPerformanceLeewaySetIterator is returned from FilterDefaultPerformanceLeewaySet and is used to iterate over the raw logs and unpacked data for DefaultPerformanceLeewaySet events raised by the Bindings contract.
type BindingsDefaultPerformanceLeewaySetIterator struct {
	Event *BindingsDefaultPerformanceLeewaySet // Event containing the contract specifics and raw log

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
func (it *BindingsDefaultPerformanceLeewaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsDefaultPerformanceLeewaySet)
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
		it.Event = new(BindingsDefaultPerformanceLeewaySet)
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
func (it *BindingsDefaultPerformanceLeewaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsDefaultPerformanceLeewaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsDefaultPerformanceLeewaySet represents a DefaultPerformanceLeewaySet event raised by the Bindings contract.
type BindingsDefaultPerformanceLeewaySet struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDefaultPerformanceLeewaySet is a free log retrieval operation binding the contract event 0x9a4b729d4e365eebc740b9d0763a1302c2d6c665ec429c87f9eca0ce2f6ba2de.
//
// Solidity: event DefaultPerformanceLeewaySet(uint256 value)
func (_Bindings *BindingsFilterer) FilterDefaultPerformanceLeewaySet(opts *bind.FilterOpts) (*BindingsDefaultPerformanceLeewaySetIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "DefaultPerformanceLeewaySet")
	if err != nil {
		return nil, err
	}
	return &BindingsDefaultPerformanceLeewaySetIterator{contract: _Bindings.contract, event: "DefaultPerformanceLeewaySet", logs: logs, sub: sub}, nil
}

// WatchDefaultPerformanceLeewaySet is a free log subscription operation binding the contract event 0x9a4b729d4e365eebc740b9d0763a1302c2d6c665ec429c87f9eca0ce2f6ba2de.
//
// Solidity: event DefaultPerformanceLeewaySet(uint256 value)
func (_Bindings *BindingsFilterer) WatchDefaultPerformanceLeewaySet(opts *bind.WatchOpts, sink chan<- *BindingsDefaultPerformanceLeewaySet) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "DefaultPerformanceLeewaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsDefaultPerformanceLeewaySet)
				if err := _Bindings.contract.UnpackLog(event, "DefaultPerformanceLeewaySet", log); err != nil {
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

// ParseDefaultPerformanceLeewaySet is a log parse operation binding the contract event 0x9a4b729d4e365eebc740b9d0763a1302c2d6c665ec429c87f9eca0ce2f6ba2de.
//
// Solidity: event DefaultPerformanceLeewaySet(uint256 value)
func (_Bindings *BindingsFilterer) ParseDefaultPerformanceLeewaySet(log types.Log) (*BindingsDefaultPerformanceLeewaySet, error) {
	event := new(BindingsDefaultPerformanceLeewaySet)
	if err := _Bindings.contract.UnpackLog(event, "DefaultPerformanceLeewaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsDefaultQueueConfigSetIterator is returned from FilterDefaultQueueConfigSet and is used to iterate over the raw logs and unpacked data for DefaultQueueConfigSet events raised by the Bindings contract.
type BindingsDefaultQueueConfigSetIterator struct {
	Event *BindingsDefaultQueueConfigSet // Event containing the contract specifics and raw log

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
func (it *BindingsDefaultQueueConfigSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsDefaultQueueConfigSet)
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
		it.Event = new(BindingsDefaultQueueConfigSet)
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
func (it *BindingsDefaultQueueConfigSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsDefaultQueueConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsDefaultQueueConfigSet represents a DefaultQueueConfigSet event raised by the Bindings contract.
type BindingsDefaultQueueConfigSet struct {
	Priority    *big.Int
	MaxDeposits *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDefaultQueueConfigSet is a free log retrieval operation binding the contract event 0xfa24f285a48ec646597783f0da25d5b690c6c6cce4df8abbf73dc3ac763ecaa3.
//
// Solidity: event DefaultQueueConfigSet(uint256 priority, uint256 maxDeposits)
func (_Bindings *BindingsFilterer) FilterDefaultQueueConfigSet(opts *bind.FilterOpts) (*BindingsDefaultQueueConfigSetIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "DefaultQueueConfigSet")
	if err != nil {
		return nil, err
	}
	return &BindingsDefaultQueueConfigSetIterator{contract: _Bindings.contract, event: "DefaultQueueConfigSet", logs: logs, sub: sub}, nil
}

// WatchDefaultQueueConfigSet is a free log subscription operation binding the contract event 0xfa24f285a48ec646597783f0da25d5b690c6c6cce4df8abbf73dc3ac763ecaa3.
//
// Solidity: event DefaultQueueConfigSet(uint256 priority, uint256 maxDeposits)
func (_Bindings *BindingsFilterer) WatchDefaultQueueConfigSet(opts *bind.WatchOpts, sink chan<- *BindingsDefaultQueueConfigSet) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "DefaultQueueConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsDefaultQueueConfigSet)
				if err := _Bindings.contract.UnpackLog(event, "DefaultQueueConfigSet", log); err != nil {
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

// ParseDefaultQueueConfigSet is a log parse operation binding the contract event 0xfa24f285a48ec646597783f0da25d5b690c6c6cce4df8abbf73dc3ac763ecaa3.
//
// Solidity: event DefaultQueueConfigSet(uint256 priority, uint256 maxDeposits)
func (_Bindings *BindingsFilterer) ParseDefaultQueueConfigSet(log types.Log) (*BindingsDefaultQueueConfigSet, error) {
	event := new(BindingsDefaultQueueConfigSet)
	if err := _Bindings.contract.UnpackLog(event, "DefaultQueueConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsDefaultRewardShareSetIterator is returned from FilterDefaultRewardShareSet and is used to iterate over the raw logs and unpacked data for DefaultRewardShareSet events raised by the Bindings contract.
type BindingsDefaultRewardShareSetIterator struct {
	Event *BindingsDefaultRewardShareSet // Event containing the contract specifics and raw log

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
func (it *BindingsDefaultRewardShareSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsDefaultRewardShareSet)
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
		it.Event = new(BindingsDefaultRewardShareSet)
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
func (it *BindingsDefaultRewardShareSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsDefaultRewardShareSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsDefaultRewardShareSet represents a DefaultRewardShareSet event raised by the Bindings contract.
type BindingsDefaultRewardShareSet struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDefaultRewardShareSet is a free log retrieval operation binding the contract event 0x5fb2003374f60e2d64b6f54356b8ba7cb4bd5ffe092bd71ee3a509fb29340311.
//
// Solidity: event DefaultRewardShareSet(uint256 value)
func (_Bindings *BindingsFilterer) FilterDefaultRewardShareSet(opts *bind.FilterOpts) (*BindingsDefaultRewardShareSetIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "DefaultRewardShareSet")
	if err != nil {
		return nil, err
	}
	return &BindingsDefaultRewardShareSetIterator{contract: _Bindings.contract, event: "DefaultRewardShareSet", logs: logs, sub: sub}, nil
}

// WatchDefaultRewardShareSet is a free log subscription operation binding the contract event 0x5fb2003374f60e2d64b6f54356b8ba7cb4bd5ffe092bd71ee3a509fb29340311.
//
// Solidity: event DefaultRewardShareSet(uint256 value)
func (_Bindings *BindingsFilterer) WatchDefaultRewardShareSet(opts *bind.WatchOpts, sink chan<- *BindingsDefaultRewardShareSet) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "DefaultRewardShareSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsDefaultRewardShareSet)
				if err := _Bindings.contract.UnpackLog(event, "DefaultRewardShareSet", log); err != nil {
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

// ParseDefaultRewardShareSet is a log parse operation binding the contract event 0x5fb2003374f60e2d64b6f54356b8ba7cb4bd5ffe092bd71ee3a509fb29340311.
//
// Solidity: event DefaultRewardShareSet(uint256 value)
func (_Bindings *BindingsFilterer) ParseDefaultRewardShareSet(log types.Log) (*BindingsDefaultRewardShareSet, error) {
	event := new(BindingsDefaultRewardShareSet)
	if err := _Bindings.contract.UnpackLog(event, "DefaultRewardShareSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsDefaultStrikesParamsSetIterator is returned from FilterDefaultStrikesParamsSet and is used to iterate over the raw logs and unpacked data for DefaultStrikesParamsSet events raised by the Bindings contract.
type BindingsDefaultStrikesParamsSetIterator struct {
	Event *BindingsDefaultStrikesParamsSet // Event containing the contract specifics and raw log

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
func (it *BindingsDefaultStrikesParamsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsDefaultStrikesParamsSet)
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
		it.Event = new(BindingsDefaultStrikesParamsSet)
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
func (it *BindingsDefaultStrikesParamsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsDefaultStrikesParamsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsDefaultStrikesParamsSet represents a DefaultStrikesParamsSet event raised by the Bindings contract.
type BindingsDefaultStrikesParamsSet struct {
	Lifetime  *big.Int
	Threshold *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDefaultStrikesParamsSet is a free log retrieval operation binding the contract event 0x7fe1cdad48b436f21fbbcf184487f60fc38ef145c8d91f0499052abded80b149.
//
// Solidity: event DefaultStrikesParamsSet(uint256 lifetime, uint256 threshold)
func (_Bindings *BindingsFilterer) FilterDefaultStrikesParamsSet(opts *bind.FilterOpts) (*BindingsDefaultStrikesParamsSetIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "DefaultStrikesParamsSet")
	if err != nil {
		return nil, err
	}
	return &BindingsDefaultStrikesParamsSetIterator{contract: _Bindings.contract, event: "DefaultStrikesParamsSet", logs: logs, sub: sub}, nil
}

// WatchDefaultStrikesParamsSet is a free log subscription operation binding the contract event 0x7fe1cdad48b436f21fbbcf184487f60fc38ef145c8d91f0499052abded80b149.
//
// Solidity: event DefaultStrikesParamsSet(uint256 lifetime, uint256 threshold)
func (_Bindings *BindingsFilterer) WatchDefaultStrikesParamsSet(opts *bind.WatchOpts, sink chan<- *BindingsDefaultStrikesParamsSet) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "DefaultStrikesParamsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsDefaultStrikesParamsSet)
				if err := _Bindings.contract.UnpackLog(event, "DefaultStrikesParamsSet", log); err != nil {
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

// ParseDefaultStrikesParamsSet is a log parse operation binding the contract event 0x7fe1cdad48b436f21fbbcf184487f60fc38ef145c8d91f0499052abded80b149.
//
// Solidity: event DefaultStrikesParamsSet(uint256 lifetime, uint256 threshold)
func (_Bindings *BindingsFilterer) ParseDefaultStrikesParamsSet(log types.Log) (*BindingsDefaultStrikesParamsSet, error) {
	event := new(BindingsDefaultStrikesParamsSet)
	if err := _Bindings.contract.UnpackLog(event, "DefaultStrikesParamsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsElRewardsStealingAdditionalFineSetIterator is returned from FilterElRewardsStealingAdditionalFineSet and is used to iterate over the raw logs and unpacked data for ElRewardsStealingAdditionalFineSet events raised by the Bindings contract.
type BindingsElRewardsStealingAdditionalFineSetIterator struct {
	Event *BindingsElRewardsStealingAdditionalFineSet // Event containing the contract specifics and raw log

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
func (it *BindingsElRewardsStealingAdditionalFineSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsElRewardsStealingAdditionalFineSet)
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
		it.Event = new(BindingsElRewardsStealingAdditionalFineSet)
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
func (it *BindingsElRewardsStealingAdditionalFineSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsElRewardsStealingAdditionalFineSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsElRewardsStealingAdditionalFineSet represents a ElRewardsStealingAdditionalFineSet event raised by the Bindings contract.
type BindingsElRewardsStealingAdditionalFineSet struct {
	CurveId *big.Int
	Fine    *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterElRewardsStealingAdditionalFineSet is a free log retrieval operation binding the contract event 0xbfdb11d95e87f282b12252576ce5982bf3136988ec3db94a3ab3a76a4455e38c.
//
// Solidity: event ElRewardsStealingAdditionalFineSet(uint256 indexed curveId, uint256 fine)
func (_Bindings *BindingsFilterer) FilterElRewardsStealingAdditionalFineSet(opts *bind.FilterOpts, curveId []*big.Int) (*BindingsElRewardsStealingAdditionalFineSetIterator, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "ElRewardsStealingAdditionalFineSet", curveIdRule)
	if err != nil {
		return nil, err
	}
	return &BindingsElRewardsStealingAdditionalFineSetIterator{contract: _Bindings.contract, event: "ElRewardsStealingAdditionalFineSet", logs: logs, sub: sub}, nil
}

// WatchElRewardsStealingAdditionalFineSet is a free log subscription operation binding the contract event 0xbfdb11d95e87f282b12252576ce5982bf3136988ec3db94a3ab3a76a4455e38c.
//
// Solidity: event ElRewardsStealingAdditionalFineSet(uint256 indexed curveId, uint256 fine)
func (_Bindings *BindingsFilterer) WatchElRewardsStealingAdditionalFineSet(opts *bind.WatchOpts, sink chan<- *BindingsElRewardsStealingAdditionalFineSet, curveId []*big.Int) (event.Subscription, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "ElRewardsStealingAdditionalFineSet", curveIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsElRewardsStealingAdditionalFineSet)
				if err := _Bindings.contract.UnpackLog(event, "ElRewardsStealingAdditionalFineSet", log); err != nil {
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

// ParseElRewardsStealingAdditionalFineSet is a log parse operation binding the contract event 0xbfdb11d95e87f282b12252576ce5982bf3136988ec3db94a3ab3a76a4455e38c.
//
// Solidity: event ElRewardsStealingAdditionalFineSet(uint256 indexed curveId, uint256 fine)
func (_Bindings *BindingsFilterer) ParseElRewardsStealingAdditionalFineSet(log types.Log) (*BindingsElRewardsStealingAdditionalFineSet, error) {
	event := new(BindingsElRewardsStealingAdditionalFineSet)
	if err := _Bindings.contract.UnpackLog(event, "ElRewardsStealingAdditionalFineSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsElRewardsStealingAdditionalFineUnsetIterator is returned from FilterElRewardsStealingAdditionalFineUnset and is used to iterate over the raw logs and unpacked data for ElRewardsStealingAdditionalFineUnset events raised by the Bindings contract.
type BindingsElRewardsStealingAdditionalFineUnsetIterator struct {
	Event *BindingsElRewardsStealingAdditionalFineUnset // Event containing the contract specifics and raw log

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
func (it *BindingsElRewardsStealingAdditionalFineUnsetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsElRewardsStealingAdditionalFineUnset)
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
		it.Event = new(BindingsElRewardsStealingAdditionalFineUnset)
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
func (it *BindingsElRewardsStealingAdditionalFineUnsetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsElRewardsStealingAdditionalFineUnsetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsElRewardsStealingAdditionalFineUnset represents a ElRewardsStealingAdditionalFineUnset event raised by the Bindings contract.
type BindingsElRewardsStealingAdditionalFineUnset struct {
	CurveId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterElRewardsStealingAdditionalFineUnset is a free log retrieval operation binding the contract event 0x937c833a4a10529235d61dcf2004cbd39968786cd4d5623a5ea1754955fa5224.
//
// Solidity: event ElRewardsStealingAdditionalFineUnset(uint256 indexed curveId)
func (_Bindings *BindingsFilterer) FilterElRewardsStealingAdditionalFineUnset(opts *bind.FilterOpts, curveId []*big.Int) (*BindingsElRewardsStealingAdditionalFineUnsetIterator, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "ElRewardsStealingAdditionalFineUnset", curveIdRule)
	if err != nil {
		return nil, err
	}
	return &BindingsElRewardsStealingAdditionalFineUnsetIterator{contract: _Bindings.contract, event: "ElRewardsStealingAdditionalFineUnset", logs: logs, sub: sub}, nil
}

// WatchElRewardsStealingAdditionalFineUnset is a free log subscription operation binding the contract event 0x937c833a4a10529235d61dcf2004cbd39968786cd4d5623a5ea1754955fa5224.
//
// Solidity: event ElRewardsStealingAdditionalFineUnset(uint256 indexed curveId)
func (_Bindings *BindingsFilterer) WatchElRewardsStealingAdditionalFineUnset(opts *bind.WatchOpts, sink chan<- *BindingsElRewardsStealingAdditionalFineUnset, curveId []*big.Int) (event.Subscription, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "ElRewardsStealingAdditionalFineUnset", curveIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsElRewardsStealingAdditionalFineUnset)
				if err := _Bindings.contract.UnpackLog(event, "ElRewardsStealingAdditionalFineUnset", log); err != nil {
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

// ParseElRewardsStealingAdditionalFineUnset is a log parse operation binding the contract event 0x937c833a4a10529235d61dcf2004cbd39968786cd4d5623a5ea1754955fa5224.
//
// Solidity: event ElRewardsStealingAdditionalFineUnset(uint256 indexed curveId)
func (_Bindings *BindingsFilterer) ParseElRewardsStealingAdditionalFineUnset(log types.Log) (*BindingsElRewardsStealingAdditionalFineUnset, error) {
	event := new(BindingsElRewardsStealingAdditionalFineUnset)
	if err := _Bindings.contract.UnpackLog(event, "ElRewardsStealingAdditionalFineUnset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsExitDelayPenaltySetIterator is returned from FilterExitDelayPenaltySet and is used to iterate over the raw logs and unpacked data for ExitDelayPenaltySet events raised by the Bindings contract.
type BindingsExitDelayPenaltySetIterator struct {
	Event *BindingsExitDelayPenaltySet // Event containing the contract specifics and raw log

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
func (it *BindingsExitDelayPenaltySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsExitDelayPenaltySet)
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
		it.Event = new(BindingsExitDelayPenaltySet)
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
func (it *BindingsExitDelayPenaltySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsExitDelayPenaltySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsExitDelayPenaltySet represents a ExitDelayPenaltySet event raised by the Bindings contract.
type BindingsExitDelayPenaltySet struct {
	CurveId *big.Int
	Penalty *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExitDelayPenaltySet is a free log retrieval operation binding the contract event 0xeadb66df4fb1918232eab1ad26c6e48acae4606e4619aabec4658fa9938889ca.
//
// Solidity: event ExitDelayPenaltySet(uint256 indexed curveId, uint256 penalty)
func (_Bindings *BindingsFilterer) FilterExitDelayPenaltySet(opts *bind.FilterOpts, curveId []*big.Int) (*BindingsExitDelayPenaltySetIterator, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "ExitDelayPenaltySet", curveIdRule)
	if err != nil {
		return nil, err
	}
	return &BindingsExitDelayPenaltySetIterator{contract: _Bindings.contract, event: "ExitDelayPenaltySet", logs: logs, sub: sub}, nil
}

// WatchExitDelayPenaltySet is a free log subscription operation binding the contract event 0xeadb66df4fb1918232eab1ad26c6e48acae4606e4619aabec4658fa9938889ca.
//
// Solidity: event ExitDelayPenaltySet(uint256 indexed curveId, uint256 penalty)
func (_Bindings *BindingsFilterer) WatchExitDelayPenaltySet(opts *bind.WatchOpts, sink chan<- *BindingsExitDelayPenaltySet, curveId []*big.Int) (event.Subscription, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "ExitDelayPenaltySet", curveIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsExitDelayPenaltySet)
				if err := _Bindings.contract.UnpackLog(event, "ExitDelayPenaltySet", log); err != nil {
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

// ParseExitDelayPenaltySet is a log parse operation binding the contract event 0xeadb66df4fb1918232eab1ad26c6e48acae4606e4619aabec4658fa9938889ca.
//
// Solidity: event ExitDelayPenaltySet(uint256 indexed curveId, uint256 penalty)
func (_Bindings *BindingsFilterer) ParseExitDelayPenaltySet(log types.Log) (*BindingsExitDelayPenaltySet, error) {
	event := new(BindingsExitDelayPenaltySet)
	if err := _Bindings.contract.UnpackLog(event, "ExitDelayPenaltySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsExitDelayPenaltyUnsetIterator is returned from FilterExitDelayPenaltyUnset and is used to iterate over the raw logs and unpacked data for ExitDelayPenaltyUnset events raised by the Bindings contract.
type BindingsExitDelayPenaltyUnsetIterator struct {
	Event *BindingsExitDelayPenaltyUnset // Event containing the contract specifics and raw log

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
func (it *BindingsExitDelayPenaltyUnsetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsExitDelayPenaltyUnset)
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
		it.Event = new(BindingsExitDelayPenaltyUnset)
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
func (it *BindingsExitDelayPenaltyUnsetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsExitDelayPenaltyUnsetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsExitDelayPenaltyUnset represents a ExitDelayPenaltyUnset event raised by the Bindings contract.
type BindingsExitDelayPenaltyUnset struct {
	CurveId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExitDelayPenaltyUnset is a free log retrieval operation binding the contract event 0x820479d95e6de13167a46c4d48b6ceb806e2d858fd85b5dd8a7f24a99a97231a.
//
// Solidity: event ExitDelayPenaltyUnset(uint256 indexed curveId)
func (_Bindings *BindingsFilterer) FilterExitDelayPenaltyUnset(opts *bind.FilterOpts, curveId []*big.Int) (*BindingsExitDelayPenaltyUnsetIterator, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "ExitDelayPenaltyUnset", curveIdRule)
	if err != nil {
		return nil, err
	}
	return &BindingsExitDelayPenaltyUnsetIterator{contract: _Bindings.contract, event: "ExitDelayPenaltyUnset", logs: logs, sub: sub}, nil
}

// WatchExitDelayPenaltyUnset is a free log subscription operation binding the contract event 0x820479d95e6de13167a46c4d48b6ceb806e2d858fd85b5dd8a7f24a99a97231a.
//
// Solidity: event ExitDelayPenaltyUnset(uint256 indexed curveId)
func (_Bindings *BindingsFilterer) WatchExitDelayPenaltyUnset(opts *bind.WatchOpts, sink chan<- *BindingsExitDelayPenaltyUnset, curveId []*big.Int) (event.Subscription, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "ExitDelayPenaltyUnset", curveIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsExitDelayPenaltyUnset)
				if err := _Bindings.contract.UnpackLog(event, "ExitDelayPenaltyUnset", log); err != nil {
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

// ParseExitDelayPenaltyUnset is a log parse operation binding the contract event 0x820479d95e6de13167a46c4d48b6ceb806e2d858fd85b5dd8a7f24a99a97231a.
//
// Solidity: event ExitDelayPenaltyUnset(uint256 indexed curveId)
func (_Bindings *BindingsFilterer) ParseExitDelayPenaltyUnset(log types.Log) (*BindingsExitDelayPenaltyUnset, error) {
	event := new(BindingsExitDelayPenaltyUnset)
	if err := _Bindings.contract.UnpackLog(event, "ExitDelayPenaltyUnset", log); err != nil {
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

// BindingsKeyRemovalChargeSetIterator is returned from FilterKeyRemovalChargeSet and is used to iterate over the raw logs and unpacked data for KeyRemovalChargeSet events raised by the Bindings contract.
type BindingsKeyRemovalChargeSetIterator struct {
	Event *BindingsKeyRemovalChargeSet // Event containing the contract specifics and raw log

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
func (it *BindingsKeyRemovalChargeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsKeyRemovalChargeSet)
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
		it.Event = new(BindingsKeyRemovalChargeSet)
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
func (it *BindingsKeyRemovalChargeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsKeyRemovalChargeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsKeyRemovalChargeSet represents a KeyRemovalChargeSet event raised by the Bindings contract.
type BindingsKeyRemovalChargeSet struct {
	CurveId          *big.Int
	KeyRemovalCharge *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterKeyRemovalChargeSet is a free log retrieval operation binding the contract event 0xeb4ecc2a5d1e33322fd211762f890015ef34e3011b4c8f5553cb2cd623358ae2.
//
// Solidity: event KeyRemovalChargeSet(uint256 indexed curveId, uint256 keyRemovalCharge)
func (_Bindings *BindingsFilterer) FilterKeyRemovalChargeSet(opts *bind.FilterOpts, curveId []*big.Int) (*BindingsKeyRemovalChargeSetIterator, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "KeyRemovalChargeSet", curveIdRule)
	if err != nil {
		return nil, err
	}
	return &BindingsKeyRemovalChargeSetIterator{contract: _Bindings.contract, event: "KeyRemovalChargeSet", logs: logs, sub: sub}, nil
}

// WatchKeyRemovalChargeSet is a free log subscription operation binding the contract event 0xeb4ecc2a5d1e33322fd211762f890015ef34e3011b4c8f5553cb2cd623358ae2.
//
// Solidity: event KeyRemovalChargeSet(uint256 indexed curveId, uint256 keyRemovalCharge)
func (_Bindings *BindingsFilterer) WatchKeyRemovalChargeSet(opts *bind.WatchOpts, sink chan<- *BindingsKeyRemovalChargeSet, curveId []*big.Int) (event.Subscription, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "KeyRemovalChargeSet", curveIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsKeyRemovalChargeSet)
				if err := _Bindings.contract.UnpackLog(event, "KeyRemovalChargeSet", log); err != nil {
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

// ParseKeyRemovalChargeSet is a log parse operation binding the contract event 0xeb4ecc2a5d1e33322fd211762f890015ef34e3011b4c8f5553cb2cd623358ae2.
//
// Solidity: event KeyRemovalChargeSet(uint256 indexed curveId, uint256 keyRemovalCharge)
func (_Bindings *BindingsFilterer) ParseKeyRemovalChargeSet(log types.Log) (*BindingsKeyRemovalChargeSet, error) {
	event := new(BindingsKeyRemovalChargeSet)
	if err := _Bindings.contract.UnpackLog(event, "KeyRemovalChargeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsKeyRemovalChargeUnsetIterator is returned from FilterKeyRemovalChargeUnset and is used to iterate over the raw logs and unpacked data for KeyRemovalChargeUnset events raised by the Bindings contract.
type BindingsKeyRemovalChargeUnsetIterator struct {
	Event *BindingsKeyRemovalChargeUnset // Event containing the contract specifics and raw log

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
func (it *BindingsKeyRemovalChargeUnsetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsKeyRemovalChargeUnset)
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
		it.Event = new(BindingsKeyRemovalChargeUnset)
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
func (it *BindingsKeyRemovalChargeUnsetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsKeyRemovalChargeUnsetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsKeyRemovalChargeUnset represents a KeyRemovalChargeUnset event raised by the Bindings contract.
type BindingsKeyRemovalChargeUnset struct {
	CurveId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterKeyRemovalChargeUnset is a free log retrieval operation binding the contract event 0xf25dcadbf389c152c9f293f7c48c878f24d1985fb3568315e7abe669b6a863b5.
//
// Solidity: event KeyRemovalChargeUnset(uint256 indexed curveId)
func (_Bindings *BindingsFilterer) FilterKeyRemovalChargeUnset(opts *bind.FilterOpts, curveId []*big.Int) (*BindingsKeyRemovalChargeUnsetIterator, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "KeyRemovalChargeUnset", curveIdRule)
	if err != nil {
		return nil, err
	}
	return &BindingsKeyRemovalChargeUnsetIterator{contract: _Bindings.contract, event: "KeyRemovalChargeUnset", logs: logs, sub: sub}, nil
}

// WatchKeyRemovalChargeUnset is a free log subscription operation binding the contract event 0xf25dcadbf389c152c9f293f7c48c878f24d1985fb3568315e7abe669b6a863b5.
//
// Solidity: event KeyRemovalChargeUnset(uint256 indexed curveId)
func (_Bindings *BindingsFilterer) WatchKeyRemovalChargeUnset(opts *bind.WatchOpts, sink chan<- *BindingsKeyRemovalChargeUnset, curveId []*big.Int) (event.Subscription, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "KeyRemovalChargeUnset", curveIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsKeyRemovalChargeUnset)
				if err := _Bindings.contract.UnpackLog(event, "KeyRemovalChargeUnset", log); err != nil {
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

// ParseKeyRemovalChargeUnset is a log parse operation binding the contract event 0xf25dcadbf389c152c9f293f7c48c878f24d1985fb3568315e7abe669b6a863b5.
//
// Solidity: event KeyRemovalChargeUnset(uint256 indexed curveId)
func (_Bindings *BindingsFilterer) ParseKeyRemovalChargeUnset(log types.Log) (*BindingsKeyRemovalChargeUnset, error) {
	event := new(BindingsKeyRemovalChargeUnset)
	if err := _Bindings.contract.UnpackLog(event, "KeyRemovalChargeUnset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsKeysLimitSetIterator is returned from FilterKeysLimitSet and is used to iterate over the raw logs and unpacked data for KeysLimitSet events raised by the Bindings contract.
type BindingsKeysLimitSetIterator struct {
	Event *BindingsKeysLimitSet // Event containing the contract specifics and raw log

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
func (it *BindingsKeysLimitSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsKeysLimitSet)
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
		it.Event = new(BindingsKeysLimitSet)
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
func (it *BindingsKeysLimitSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsKeysLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsKeysLimitSet represents a KeysLimitSet event raised by the Bindings contract.
type BindingsKeysLimitSet struct {
	CurveId *big.Int
	Limit   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterKeysLimitSet is a free log retrieval operation binding the contract event 0x9afc700402022d6c295c2958cca42433c1d1c0ed17bd245398d287d0a39bf98c.
//
// Solidity: event KeysLimitSet(uint256 indexed curveId, uint256 limit)
func (_Bindings *BindingsFilterer) FilterKeysLimitSet(opts *bind.FilterOpts, curveId []*big.Int) (*BindingsKeysLimitSetIterator, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "KeysLimitSet", curveIdRule)
	if err != nil {
		return nil, err
	}
	return &BindingsKeysLimitSetIterator{contract: _Bindings.contract, event: "KeysLimitSet", logs: logs, sub: sub}, nil
}

// WatchKeysLimitSet is a free log subscription operation binding the contract event 0x9afc700402022d6c295c2958cca42433c1d1c0ed17bd245398d287d0a39bf98c.
//
// Solidity: event KeysLimitSet(uint256 indexed curveId, uint256 limit)
func (_Bindings *BindingsFilterer) WatchKeysLimitSet(opts *bind.WatchOpts, sink chan<- *BindingsKeysLimitSet, curveId []*big.Int) (event.Subscription, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "KeysLimitSet", curveIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsKeysLimitSet)
				if err := _Bindings.contract.UnpackLog(event, "KeysLimitSet", log); err != nil {
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

// ParseKeysLimitSet is a log parse operation binding the contract event 0x9afc700402022d6c295c2958cca42433c1d1c0ed17bd245398d287d0a39bf98c.
//
// Solidity: event KeysLimitSet(uint256 indexed curveId, uint256 limit)
func (_Bindings *BindingsFilterer) ParseKeysLimitSet(log types.Log) (*BindingsKeysLimitSet, error) {
	event := new(BindingsKeysLimitSet)
	if err := _Bindings.contract.UnpackLog(event, "KeysLimitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsKeysLimitUnsetIterator is returned from FilterKeysLimitUnset and is used to iterate over the raw logs and unpacked data for KeysLimitUnset events raised by the Bindings contract.
type BindingsKeysLimitUnsetIterator struct {
	Event *BindingsKeysLimitUnset // Event containing the contract specifics and raw log

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
func (it *BindingsKeysLimitUnsetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsKeysLimitUnset)
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
		it.Event = new(BindingsKeysLimitUnset)
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
func (it *BindingsKeysLimitUnsetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsKeysLimitUnsetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsKeysLimitUnset represents a KeysLimitUnset event raised by the Bindings contract.
type BindingsKeysLimitUnset struct {
	CurveId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterKeysLimitUnset is a free log retrieval operation binding the contract event 0xdf66171f66c39a977630f44442f2fe955a5a8c6805ca3880b77a054ca4adf734.
//
// Solidity: event KeysLimitUnset(uint256 indexed curveId)
func (_Bindings *BindingsFilterer) FilterKeysLimitUnset(opts *bind.FilterOpts, curveId []*big.Int) (*BindingsKeysLimitUnsetIterator, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "KeysLimitUnset", curveIdRule)
	if err != nil {
		return nil, err
	}
	return &BindingsKeysLimitUnsetIterator{contract: _Bindings.contract, event: "KeysLimitUnset", logs: logs, sub: sub}, nil
}

// WatchKeysLimitUnset is a free log subscription operation binding the contract event 0xdf66171f66c39a977630f44442f2fe955a5a8c6805ca3880b77a054ca4adf734.
//
// Solidity: event KeysLimitUnset(uint256 indexed curveId)
func (_Bindings *BindingsFilterer) WatchKeysLimitUnset(opts *bind.WatchOpts, sink chan<- *BindingsKeysLimitUnset, curveId []*big.Int) (event.Subscription, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "KeysLimitUnset", curveIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsKeysLimitUnset)
				if err := _Bindings.contract.UnpackLog(event, "KeysLimitUnset", log); err != nil {
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

// ParseKeysLimitUnset is a log parse operation binding the contract event 0xdf66171f66c39a977630f44442f2fe955a5a8c6805ca3880b77a054ca4adf734.
//
// Solidity: event KeysLimitUnset(uint256 indexed curveId)
func (_Bindings *BindingsFilterer) ParseKeysLimitUnset(log types.Log) (*BindingsKeysLimitUnset, error) {
	event := new(BindingsKeysLimitUnset)
	if err := _Bindings.contract.UnpackLog(event, "KeysLimitUnset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsMaxWithdrawalRequestFeeSetIterator is returned from FilterMaxWithdrawalRequestFeeSet and is used to iterate over the raw logs and unpacked data for MaxWithdrawalRequestFeeSet events raised by the Bindings contract.
type BindingsMaxWithdrawalRequestFeeSetIterator struct {
	Event *BindingsMaxWithdrawalRequestFeeSet // Event containing the contract specifics and raw log

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
func (it *BindingsMaxWithdrawalRequestFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsMaxWithdrawalRequestFeeSet)
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
		it.Event = new(BindingsMaxWithdrawalRequestFeeSet)
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
func (it *BindingsMaxWithdrawalRequestFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsMaxWithdrawalRequestFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsMaxWithdrawalRequestFeeSet represents a MaxWithdrawalRequestFeeSet event raised by the Bindings contract.
type BindingsMaxWithdrawalRequestFeeSet struct {
	CurveId *big.Int
	Fee     *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMaxWithdrawalRequestFeeSet is a free log retrieval operation binding the contract event 0x1b5c70fd9206d548520cfeafa68b28852b05b6035f0408c6956f5ddf3819dbb2.
//
// Solidity: event MaxWithdrawalRequestFeeSet(uint256 indexed curveId, uint256 fee)
func (_Bindings *BindingsFilterer) FilterMaxWithdrawalRequestFeeSet(opts *bind.FilterOpts, curveId []*big.Int) (*BindingsMaxWithdrawalRequestFeeSetIterator, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "MaxWithdrawalRequestFeeSet", curveIdRule)
	if err != nil {
		return nil, err
	}
	return &BindingsMaxWithdrawalRequestFeeSetIterator{contract: _Bindings.contract, event: "MaxWithdrawalRequestFeeSet", logs: logs, sub: sub}, nil
}

// WatchMaxWithdrawalRequestFeeSet is a free log subscription operation binding the contract event 0x1b5c70fd9206d548520cfeafa68b28852b05b6035f0408c6956f5ddf3819dbb2.
//
// Solidity: event MaxWithdrawalRequestFeeSet(uint256 indexed curveId, uint256 fee)
func (_Bindings *BindingsFilterer) WatchMaxWithdrawalRequestFeeSet(opts *bind.WatchOpts, sink chan<- *BindingsMaxWithdrawalRequestFeeSet, curveId []*big.Int) (event.Subscription, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "MaxWithdrawalRequestFeeSet", curveIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsMaxWithdrawalRequestFeeSet)
				if err := _Bindings.contract.UnpackLog(event, "MaxWithdrawalRequestFeeSet", log); err != nil {
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

// ParseMaxWithdrawalRequestFeeSet is a log parse operation binding the contract event 0x1b5c70fd9206d548520cfeafa68b28852b05b6035f0408c6956f5ddf3819dbb2.
//
// Solidity: event MaxWithdrawalRequestFeeSet(uint256 indexed curveId, uint256 fee)
func (_Bindings *BindingsFilterer) ParseMaxWithdrawalRequestFeeSet(log types.Log) (*BindingsMaxWithdrawalRequestFeeSet, error) {
	event := new(BindingsMaxWithdrawalRequestFeeSet)
	if err := _Bindings.contract.UnpackLog(event, "MaxWithdrawalRequestFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsMaxWithdrawalRequestFeeUnsetIterator is returned from FilterMaxWithdrawalRequestFeeUnset and is used to iterate over the raw logs and unpacked data for MaxWithdrawalRequestFeeUnset events raised by the Bindings contract.
type BindingsMaxWithdrawalRequestFeeUnsetIterator struct {
	Event *BindingsMaxWithdrawalRequestFeeUnset // Event containing the contract specifics and raw log

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
func (it *BindingsMaxWithdrawalRequestFeeUnsetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsMaxWithdrawalRequestFeeUnset)
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
		it.Event = new(BindingsMaxWithdrawalRequestFeeUnset)
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
func (it *BindingsMaxWithdrawalRequestFeeUnsetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsMaxWithdrawalRequestFeeUnsetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsMaxWithdrawalRequestFeeUnset represents a MaxWithdrawalRequestFeeUnset event raised by the Bindings contract.
type BindingsMaxWithdrawalRequestFeeUnset struct {
	CurveId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMaxWithdrawalRequestFeeUnset is a free log retrieval operation binding the contract event 0xbd84391fd589d04b1bf2df444b2d1ec15a5ff2e9eb2eb979688a9adc9f777eec.
//
// Solidity: event MaxWithdrawalRequestFeeUnset(uint256 indexed curveId)
func (_Bindings *BindingsFilterer) FilterMaxWithdrawalRequestFeeUnset(opts *bind.FilterOpts, curveId []*big.Int) (*BindingsMaxWithdrawalRequestFeeUnsetIterator, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "MaxWithdrawalRequestFeeUnset", curveIdRule)
	if err != nil {
		return nil, err
	}
	return &BindingsMaxWithdrawalRequestFeeUnsetIterator{contract: _Bindings.contract, event: "MaxWithdrawalRequestFeeUnset", logs: logs, sub: sub}, nil
}

// WatchMaxWithdrawalRequestFeeUnset is a free log subscription operation binding the contract event 0xbd84391fd589d04b1bf2df444b2d1ec15a5ff2e9eb2eb979688a9adc9f777eec.
//
// Solidity: event MaxWithdrawalRequestFeeUnset(uint256 indexed curveId)
func (_Bindings *BindingsFilterer) WatchMaxWithdrawalRequestFeeUnset(opts *bind.WatchOpts, sink chan<- *BindingsMaxWithdrawalRequestFeeUnset, curveId []*big.Int) (event.Subscription, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "MaxWithdrawalRequestFeeUnset", curveIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsMaxWithdrawalRequestFeeUnset)
				if err := _Bindings.contract.UnpackLog(event, "MaxWithdrawalRequestFeeUnset", log); err != nil {
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

// ParseMaxWithdrawalRequestFeeUnset is a log parse operation binding the contract event 0xbd84391fd589d04b1bf2df444b2d1ec15a5ff2e9eb2eb979688a9adc9f777eec.
//
// Solidity: event MaxWithdrawalRequestFeeUnset(uint256 indexed curveId)
func (_Bindings *BindingsFilterer) ParseMaxWithdrawalRequestFeeUnset(log types.Log) (*BindingsMaxWithdrawalRequestFeeUnset, error) {
	event := new(BindingsMaxWithdrawalRequestFeeUnset)
	if err := _Bindings.contract.UnpackLog(event, "MaxWithdrawalRequestFeeUnset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsPerformanceCoefficientsSetIterator is returned from FilterPerformanceCoefficientsSet and is used to iterate over the raw logs and unpacked data for PerformanceCoefficientsSet events raised by the Bindings contract.
type BindingsPerformanceCoefficientsSetIterator struct {
	Event *BindingsPerformanceCoefficientsSet // Event containing the contract specifics and raw log

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
func (it *BindingsPerformanceCoefficientsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsPerformanceCoefficientsSet)
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
		it.Event = new(BindingsPerformanceCoefficientsSet)
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
func (it *BindingsPerformanceCoefficientsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsPerformanceCoefficientsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsPerformanceCoefficientsSet represents a PerformanceCoefficientsSet event raised by the Bindings contract.
type BindingsPerformanceCoefficientsSet struct {
	CurveId            *big.Int
	AttestationsWeight *big.Int
	BlocksWeight       *big.Int
	SyncWeight         *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterPerformanceCoefficientsSet is a free log retrieval operation binding the contract event 0xc561186e15cc08e3740a07971a4560b2fc9ca49d06a1da53ccc755c53ffad6b5.
//
// Solidity: event PerformanceCoefficientsSet(uint256 indexed curveId, uint256 attestationsWeight, uint256 blocksWeight, uint256 syncWeight)
func (_Bindings *BindingsFilterer) FilterPerformanceCoefficientsSet(opts *bind.FilterOpts, curveId []*big.Int) (*BindingsPerformanceCoefficientsSetIterator, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "PerformanceCoefficientsSet", curveIdRule)
	if err != nil {
		return nil, err
	}
	return &BindingsPerformanceCoefficientsSetIterator{contract: _Bindings.contract, event: "PerformanceCoefficientsSet", logs: logs, sub: sub}, nil
}

// WatchPerformanceCoefficientsSet is a free log subscription operation binding the contract event 0xc561186e15cc08e3740a07971a4560b2fc9ca49d06a1da53ccc755c53ffad6b5.
//
// Solidity: event PerformanceCoefficientsSet(uint256 indexed curveId, uint256 attestationsWeight, uint256 blocksWeight, uint256 syncWeight)
func (_Bindings *BindingsFilterer) WatchPerformanceCoefficientsSet(opts *bind.WatchOpts, sink chan<- *BindingsPerformanceCoefficientsSet, curveId []*big.Int) (event.Subscription, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "PerformanceCoefficientsSet", curveIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsPerformanceCoefficientsSet)
				if err := _Bindings.contract.UnpackLog(event, "PerformanceCoefficientsSet", log); err != nil {
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

// ParsePerformanceCoefficientsSet is a log parse operation binding the contract event 0xc561186e15cc08e3740a07971a4560b2fc9ca49d06a1da53ccc755c53ffad6b5.
//
// Solidity: event PerformanceCoefficientsSet(uint256 indexed curveId, uint256 attestationsWeight, uint256 blocksWeight, uint256 syncWeight)
func (_Bindings *BindingsFilterer) ParsePerformanceCoefficientsSet(log types.Log) (*BindingsPerformanceCoefficientsSet, error) {
	event := new(BindingsPerformanceCoefficientsSet)
	if err := _Bindings.contract.UnpackLog(event, "PerformanceCoefficientsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsPerformanceCoefficientsUnsetIterator is returned from FilterPerformanceCoefficientsUnset and is used to iterate over the raw logs and unpacked data for PerformanceCoefficientsUnset events raised by the Bindings contract.
type BindingsPerformanceCoefficientsUnsetIterator struct {
	Event *BindingsPerformanceCoefficientsUnset // Event containing the contract specifics and raw log

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
func (it *BindingsPerformanceCoefficientsUnsetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsPerformanceCoefficientsUnset)
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
		it.Event = new(BindingsPerformanceCoefficientsUnset)
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
func (it *BindingsPerformanceCoefficientsUnsetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsPerformanceCoefficientsUnsetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsPerformanceCoefficientsUnset represents a PerformanceCoefficientsUnset event raised by the Bindings contract.
type BindingsPerformanceCoefficientsUnset struct {
	CurveId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPerformanceCoefficientsUnset is a free log retrieval operation binding the contract event 0x85fce40bc4dd2050ec84e91faf44c326631d96d8687b104c80a3eeb24885b613.
//
// Solidity: event PerformanceCoefficientsUnset(uint256 indexed curveId)
func (_Bindings *BindingsFilterer) FilterPerformanceCoefficientsUnset(opts *bind.FilterOpts, curveId []*big.Int) (*BindingsPerformanceCoefficientsUnsetIterator, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "PerformanceCoefficientsUnset", curveIdRule)
	if err != nil {
		return nil, err
	}
	return &BindingsPerformanceCoefficientsUnsetIterator{contract: _Bindings.contract, event: "PerformanceCoefficientsUnset", logs: logs, sub: sub}, nil
}

// WatchPerformanceCoefficientsUnset is a free log subscription operation binding the contract event 0x85fce40bc4dd2050ec84e91faf44c326631d96d8687b104c80a3eeb24885b613.
//
// Solidity: event PerformanceCoefficientsUnset(uint256 indexed curveId)
func (_Bindings *BindingsFilterer) WatchPerformanceCoefficientsUnset(opts *bind.WatchOpts, sink chan<- *BindingsPerformanceCoefficientsUnset, curveId []*big.Int) (event.Subscription, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "PerformanceCoefficientsUnset", curveIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsPerformanceCoefficientsUnset)
				if err := _Bindings.contract.UnpackLog(event, "PerformanceCoefficientsUnset", log); err != nil {
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

// ParsePerformanceCoefficientsUnset is a log parse operation binding the contract event 0x85fce40bc4dd2050ec84e91faf44c326631d96d8687b104c80a3eeb24885b613.
//
// Solidity: event PerformanceCoefficientsUnset(uint256 indexed curveId)
func (_Bindings *BindingsFilterer) ParsePerformanceCoefficientsUnset(log types.Log) (*BindingsPerformanceCoefficientsUnset, error) {
	event := new(BindingsPerformanceCoefficientsUnset)
	if err := _Bindings.contract.UnpackLog(event, "PerformanceCoefficientsUnset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsPerformanceLeewayDataSetIterator is returned from FilterPerformanceLeewayDataSet and is used to iterate over the raw logs and unpacked data for PerformanceLeewayDataSet events raised by the Bindings contract.
type BindingsPerformanceLeewayDataSetIterator struct {
	Event *BindingsPerformanceLeewayDataSet // Event containing the contract specifics and raw log

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
func (it *BindingsPerformanceLeewayDataSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsPerformanceLeewayDataSet)
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
		it.Event = new(BindingsPerformanceLeewayDataSet)
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
func (it *BindingsPerformanceLeewayDataSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsPerformanceLeewayDataSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsPerformanceLeewayDataSet represents a PerformanceLeewayDataSet event raised by the Bindings contract.
type BindingsPerformanceLeewayDataSet struct {
	CurveId *big.Int
	Data    []ICSParametersRegistryKeyNumberValueInterval
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPerformanceLeewayDataSet is a free log retrieval operation binding the contract event 0x095ebf90f03ce7e04c184cbc20f7e79674aa2a457e6e09e9082e0dc0350365a8.
//
// Solidity: event PerformanceLeewayDataSet(uint256 indexed curveId, (uint256,uint256)[] data)
func (_Bindings *BindingsFilterer) FilterPerformanceLeewayDataSet(opts *bind.FilterOpts, curveId []*big.Int) (*BindingsPerformanceLeewayDataSetIterator, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "PerformanceLeewayDataSet", curveIdRule)
	if err != nil {
		return nil, err
	}
	return &BindingsPerformanceLeewayDataSetIterator{contract: _Bindings.contract, event: "PerformanceLeewayDataSet", logs: logs, sub: sub}, nil
}

// WatchPerformanceLeewayDataSet is a free log subscription operation binding the contract event 0x095ebf90f03ce7e04c184cbc20f7e79674aa2a457e6e09e9082e0dc0350365a8.
//
// Solidity: event PerformanceLeewayDataSet(uint256 indexed curveId, (uint256,uint256)[] data)
func (_Bindings *BindingsFilterer) WatchPerformanceLeewayDataSet(opts *bind.WatchOpts, sink chan<- *BindingsPerformanceLeewayDataSet, curveId []*big.Int) (event.Subscription, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "PerformanceLeewayDataSet", curveIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsPerformanceLeewayDataSet)
				if err := _Bindings.contract.UnpackLog(event, "PerformanceLeewayDataSet", log); err != nil {
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

// ParsePerformanceLeewayDataSet is a log parse operation binding the contract event 0x095ebf90f03ce7e04c184cbc20f7e79674aa2a457e6e09e9082e0dc0350365a8.
//
// Solidity: event PerformanceLeewayDataSet(uint256 indexed curveId, (uint256,uint256)[] data)
func (_Bindings *BindingsFilterer) ParsePerformanceLeewayDataSet(log types.Log) (*BindingsPerformanceLeewayDataSet, error) {
	event := new(BindingsPerformanceLeewayDataSet)
	if err := _Bindings.contract.UnpackLog(event, "PerformanceLeewayDataSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsPerformanceLeewayDataUnsetIterator is returned from FilterPerformanceLeewayDataUnset and is used to iterate over the raw logs and unpacked data for PerformanceLeewayDataUnset events raised by the Bindings contract.
type BindingsPerformanceLeewayDataUnsetIterator struct {
	Event *BindingsPerformanceLeewayDataUnset // Event containing the contract specifics and raw log

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
func (it *BindingsPerformanceLeewayDataUnsetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsPerformanceLeewayDataUnset)
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
		it.Event = new(BindingsPerformanceLeewayDataUnset)
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
func (it *BindingsPerformanceLeewayDataUnsetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsPerformanceLeewayDataUnsetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsPerformanceLeewayDataUnset represents a PerformanceLeewayDataUnset event raised by the Bindings contract.
type BindingsPerformanceLeewayDataUnset struct {
	CurveId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPerformanceLeewayDataUnset is a free log retrieval operation binding the contract event 0x7cf549cc2bf55a4f33490a260e27f58e9cbad2b10b5cdcdc9f93dc951fdc68fd.
//
// Solidity: event PerformanceLeewayDataUnset(uint256 indexed curveId)
func (_Bindings *BindingsFilterer) FilterPerformanceLeewayDataUnset(opts *bind.FilterOpts, curveId []*big.Int) (*BindingsPerformanceLeewayDataUnsetIterator, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "PerformanceLeewayDataUnset", curveIdRule)
	if err != nil {
		return nil, err
	}
	return &BindingsPerformanceLeewayDataUnsetIterator{contract: _Bindings.contract, event: "PerformanceLeewayDataUnset", logs: logs, sub: sub}, nil
}

// WatchPerformanceLeewayDataUnset is a free log subscription operation binding the contract event 0x7cf549cc2bf55a4f33490a260e27f58e9cbad2b10b5cdcdc9f93dc951fdc68fd.
//
// Solidity: event PerformanceLeewayDataUnset(uint256 indexed curveId)
func (_Bindings *BindingsFilterer) WatchPerformanceLeewayDataUnset(opts *bind.WatchOpts, sink chan<- *BindingsPerformanceLeewayDataUnset, curveId []*big.Int) (event.Subscription, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "PerformanceLeewayDataUnset", curveIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsPerformanceLeewayDataUnset)
				if err := _Bindings.contract.UnpackLog(event, "PerformanceLeewayDataUnset", log); err != nil {
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

// ParsePerformanceLeewayDataUnset is a log parse operation binding the contract event 0x7cf549cc2bf55a4f33490a260e27f58e9cbad2b10b5cdcdc9f93dc951fdc68fd.
//
// Solidity: event PerformanceLeewayDataUnset(uint256 indexed curveId)
func (_Bindings *BindingsFilterer) ParsePerformanceLeewayDataUnset(log types.Log) (*BindingsPerformanceLeewayDataUnset, error) {
	event := new(BindingsPerformanceLeewayDataUnset)
	if err := _Bindings.contract.UnpackLog(event, "PerformanceLeewayDataUnset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsQueueConfigSetIterator is returned from FilterQueueConfigSet and is used to iterate over the raw logs and unpacked data for QueueConfigSet events raised by the Bindings contract.
type BindingsQueueConfigSetIterator struct {
	Event *BindingsQueueConfigSet // Event containing the contract specifics and raw log

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
func (it *BindingsQueueConfigSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsQueueConfigSet)
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
		it.Event = new(BindingsQueueConfigSet)
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
func (it *BindingsQueueConfigSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsQueueConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsQueueConfigSet represents a QueueConfigSet event raised by the Bindings contract.
type BindingsQueueConfigSet struct {
	CurveId     *big.Int
	Priority    *big.Int
	MaxDeposits *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterQueueConfigSet is a free log retrieval operation binding the contract event 0x66461c5e3c113da34b5f1035b801700a20e2f1a26c87d45b4346fc68e3a8cb5c.
//
// Solidity: event QueueConfigSet(uint256 indexed curveId, uint256 priority, uint256 maxDeposits)
func (_Bindings *BindingsFilterer) FilterQueueConfigSet(opts *bind.FilterOpts, curveId []*big.Int) (*BindingsQueueConfigSetIterator, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "QueueConfigSet", curveIdRule)
	if err != nil {
		return nil, err
	}
	return &BindingsQueueConfigSetIterator{contract: _Bindings.contract, event: "QueueConfigSet", logs: logs, sub: sub}, nil
}

// WatchQueueConfigSet is a free log subscription operation binding the contract event 0x66461c5e3c113da34b5f1035b801700a20e2f1a26c87d45b4346fc68e3a8cb5c.
//
// Solidity: event QueueConfigSet(uint256 indexed curveId, uint256 priority, uint256 maxDeposits)
func (_Bindings *BindingsFilterer) WatchQueueConfigSet(opts *bind.WatchOpts, sink chan<- *BindingsQueueConfigSet, curveId []*big.Int) (event.Subscription, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "QueueConfigSet", curveIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsQueueConfigSet)
				if err := _Bindings.contract.UnpackLog(event, "QueueConfigSet", log); err != nil {
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

// ParseQueueConfigSet is a log parse operation binding the contract event 0x66461c5e3c113da34b5f1035b801700a20e2f1a26c87d45b4346fc68e3a8cb5c.
//
// Solidity: event QueueConfigSet(uint256 indexed curveId, uint256 priority, uint256 maxDeposits)
func (_Bindings *BindingsFilterer) ParseQueueConfigSet(log types.Log) (*BindingsQueueConfigSet, error) {
	event := new(BindingsQueueConfigSet)
	if err := _Bindings.contract.UnpackLog(event, "QueueConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsQueueConfigUnsetIterator is returned from FilterQueueConfigUnset and is used to iterate over the raw logs and unpacked data for QueueConfigUnset events raised by the Bindings contract.
type BindingsQueueConfigUnsetIterator struct {
	Event *BindingsQueueConfigUnset // Event containing the contract specifics and raw log

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
func (it *BindingsQueueConfigUnsetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsQueueConfigUnset)
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
		it.Event = new(BindingsQueueConfigUnset)
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
func (it *BindingsQueueConfigUnsetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsQueueConfigUnsetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsQueueConfigUnset represents a QueueConfigUnset event raised by the Bindings contract.
type BindingsQueueConfigUnset struct {
	CurveId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterQueueConfigUnset is a free log retrieval operation binding the contract event 0x6aa2be35a03b2495cc2e2d7c23c395d6294eac0f08b8975d5b61ecfe858c433b.
//
// Solidity: event QueueConfigUnset(uint256 indexed curveId)
func (_Bindings *BindingsFilterer) FilterQueueConfigUnset(opts *bind.FilterOpts, curveId []*big.Int) (*BindingsQueueConfigUnsetIterator, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "QueueConfigUnset", curveIdRule)
	if err != nil {
		return nil, err
	}
	return &BindingsQueueConfigUnsetIterator{contract: _Bindings.contract, event: "QueueConfigUnset", logs: logs, sub: sub}, nil
}

// WatchQueueConfigUnset is a free log subscription operation binding the contract event 0x6aa2be35a03b2495cc2e2d7c23c395d6294eac0f08b8975d5b61ecfe858c433b.
//
// Solidity: event QueueConfigUnset(uint256 indexed curveId)
func (_Bindings *BindingsFilterer) WatchQueueConfigUnset(opts *bind.WatchOpts, sink chan<- *BindingsQueueConfigUnset, curveId []*big.Int) (event.Subscription, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "QueueConfigUnset", curveIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsQueueConfigUnset)
				if err := _Bindings.contract.UnpackLog(event, "QueueConfigUnset", log); err != nil {
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

// ParseQueueConfigUnset is a log parse operation binding the contract event 0x6aa2be35a03b2495cc2e2d7c23c395d6294eac0f08b8975d5b61ecfe858c433b.
//
// Solidity: event QueueConfigUnset(uint256 indexed curveId)
func (_Bindings *BindingsFilterer) ParseQueueConfigUnset(log types.Log) (*BindingsQueueConfigUnset, error) {
	event := new(BindingsQueueConfigUnset)
	if err := _Bindings.contract.UnpackLog(event, "QueueConfigUnset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsRewardShareDataSetIterator is returned from FilterRewardShareDataSet and is used to iterate over the raw logs and unpacked data for RewardShareDataSet events raised by the Bindings contract.
type BindingsRewardShareDataSetIterator struct {
	Event *BindingsRewardShareDataSet // Event containing the contract specifics and raw log

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
func (it *BindingsRewardShareDataSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsRewardShareDataSet)
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
		it.Event = new(BindingsRewardShareDataSet)
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
func (it *BindingsRewardShareDataSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsRewardShareDataSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsRewardShareDataSet represents a RewardShareDataSet event raised by the Bindings contract.
type BindingsRewardShareDataSet struct {
	CurveId *big.Int
	Data    []ICSParametersRegistryKeyNumberValueInterval
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRewardShareDataSet is a free log retrieval operation binding the contract event 0x1e4643269db812a4814fc9923a5cb5b9f8f212193add82e9b525d9e20c110410.
//
// Solidity: event RewardShareDataSet(uint256 indexed curveId, (uint256,uint256)[] data)
func (_Bindings *BindingsFilterer) FilterRewardShareDataSet(opts *bind.FilterOpts, curveId []*big.Int) (*BindingsRewardShareDataSetIterator, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "RewardShareDataSet", curveIdRule)
	if err != nil {
		return nil, err
	}
	return &BindingsRewardShareDataSetIterator{contract: _Bindings.contract, event: "RewardShareDataSet", logs: logs, sub: sub}, nil
}

// WatchRewardShareDataSet is a free log subscription operation binding the contract event 0x1e4643269db812a4814fc9923a5cb5b9f8f212193add82e9b525d9e20c110410.
//
// Solidity: event RewardShareDataSet(uint256 indexed curveId, (uint256,uint256)[] data)
func (_Bindings *BindingsFilterer) WatchRewardShareDataSet(opts *bind.WatchOpts, sink chan<- *BindingsRewardShareDataSet, curveId []*big.Int) (event.Subscription, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "RewardShareDataSet", curveIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsRewardShareDataSet)
				if err := _Bindings.contract.UnpackLog(event, "RewardShareDataSet", log); err != nil {
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

// ParseRewardShareDataSet is a log parse operation binding the contract event 0x1e4643269db812a4814fc9923a5cb5b9f8f212193add82e9b525d9e20c110410.
//
// Solidity: event RewardShareDataSet(uint256 indexed curveId, (uint256,uint256)[] data)
func (_Bindings *BindingsFilterer) ParseRewardShareDataSet(log types.Log) (*BindingsRewardShareDataSet, error) {
	event := new(BindingsRewardShareDataSet)
	if err := _Bindings.contract.UnpackLog(event, "RewardShareDataSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsRewardShareDataUnsetIterator is returned from FilterRewardShareDataUnset and is used to iterate over the raw logs and unpacked data for RewardShareDataUnset events raised by the Bindings contract.
type BindingsRewardShareDataUnsetIterator struct {
	Event *BindingsRewardShareDataUnset // Event containing the contract specifics and raw log

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
func (it *BindingsRewardShareDataUnsetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsRewardShareDataUnset)
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
		it.Event = new(BindingsRewardShareDataUnset)
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
func (it *BindingsRewardShareDataUnsetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsRewardShareDataUnsetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsRewardShareDataUnset represents a RewardShareDataUnset event raised by the Bindings contract.
type BindingsRewardShareDataUnset struct {
	CurveId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRewardShareDataUnset is a free log retrieval operation binding the contract event 0x15c3f807b29a6a5da4ae9068d9c7f556c9ae59f1c5301e85e9a6f966bfd2e84d.
//
// Solidity: event RewardShareDataUnset(uint256 indexed curveId)
func (_Bindings *BindingsFilterer) FilterRewardShareDataUnset(opts *bind.FilterOpts, curveId []*big.Int) (*BindingsRewardShareDataUnsetIterator, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "RewardShareDataUnset", curveIdRule)
	if err != nil {
		return nil, err
	}
	return &BindingsRewardShareDataUnsetIterator{contract: _Bindings.contract, event: "RewardShareDataUnset", logs: logs, sub: sub}, nil
}

// WatchRewardShareDataUnset is a free log subscription operation binding the contract event 0x15c3f807b29a6a5da4ae9068d9c7f556c9ae59f1c5301e85e9a6f966bfd2e84d.
//
// Solidity: event RewardShareDataUnset(uint256 indexed curveId)
func (_Bindings *BindingsFilterer) WatchRewardShareDataUnset(opts *bind.WatchOpts, sink chan<- *BindingsRewardShareDataUnset, curveId []*big.Int) (event.Subscription, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "RewardShareDataUnset", curveIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsRewardShareDataUnset)
				if err := _Bindings.contract.UnpackLog(event, "RewardShareDataUnset", log); err != nil {
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

// ParseRewardShareDataUnset is a log parse operation binding the contract event 0x15c3f807b29a6a5da4ae9068d9c7f556c9ae59f1c5301e85e9a6f966bfd2e84d.
//
// Solidity: event RewardShareDataUnset(uint256 indexed curveId)
func (_Bindings *BindingsFilterer) ParseRewardShareDataUnset(log types.Log) (*BindingsRewardShareDataUnset, error) {
	event := new(BindingsRewardShareDataUnset)
	if err := _Bindings.contract.UnpackLog(event, "RewardShareDataUnset", log); err != nil {
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

// BindingsStrikesParamsSetIterator is returned from FilterStrikesParamsSet and is used to iterate over the raw logs and unpacked data for StrikesParamsSet events raised by the Bindings contract.
type BindingsStrikesParamsSetIterator struct {
	Event *BindingsStrikesParamsSet // Event containing the contract specifics and raw log

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
func (it *BindingsStrikesParamsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsStrikesParamsSet)
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
		it.Event = new(BindingsStrikesParamsSet)
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
func (it *BindingsStrikesParamsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsStrikesParamsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsStrikesParamsSet represents a StrikesParamsSet event raised by the Bindings contract.
type BindingsStrikesParamsSet struct {
	CurveId   *big.Int
	Lifetime  *big.Int
	Threshold *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStrikesParamsSet is a free log retrieval operation binding the contract event 0xd95a69e8c1c04bb8dfa68176be1fad93af3f5f260cb6ee5fa0792bb6af5b0966.
//
// Solidity: event StrikesParamsSet(uint256 indexed curveId, uint256 lifetime, uint256 threshold)
func (_Bindings *BindingsFilterer) FilterStrikesParamsSet(opts *bind.FilterOpts, curveId []*big.Int) (*BindingsStrikesParamsSetIterator, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "StrikesParamsSet", curveIdRule)
	if err != nil {
		return nil, err
	}
	return &BindingsStrikesParamsSetIterator{contract: _Bindings.contract, event: "StrikesParamsSet", logs: logs, sub: sub}, nil
}

// WatchStrikesParamsSet is a free log subscription operation binding the contract event 0xd95a69e8c1c04bb8dfa68176be1fad93af3f5f260cb6ee5fa0792bb6af5b0966.
//
// Solidity: event StrikesParamsSet(uint256 indexed curveId, uint256 lifetime, uint256 threshold)
func (_Bindings *BindingsFilterer) WatchStrikesParamsSet(opts *bind.WatchOpts, sink chan<- *BindingsStrikesParamsSet, curveId []*big.Int) (event.Subscription, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "StrikesParamsSet", curveIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsStrikesParamsSet)
				if err := _Bindings.contract.UnpackLog(event, "StrikesParamsSet", log); err != nil {
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

// ParseStrikesParamsSet is a log parse operation binding the contract event 0xd95a69e8c1c04bb8dfa68176be1fad93af3f5f260cb6ee5fa0792bb6af5b0966.
//
// Solidity: event StrikesParamsSet(uint256 indexed curveId, uint256 lifetime, uint256 threshold)
func (_Bindings *BindingsFilterer) ParseStrikesParamsSet(log types.Log) (*BindingsStrikesParamsSet, error) {
	event := new(BindingsStrikesParamsSet)
	if err := _Bindings.contract.UnpackLog(event, "StrikesParamsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsStrikesParamsUnsetIterator is returned from FilterStrikesParamsUnset and is used to iterate over the raw logs and unpacked data for StrikesParamsUnset events raised by the Bindings contract.
type BindingsStrikesParamsUnsetIterator struct {
	Event *BindingsStrikesParamsUnset // Event containing the contract specifics and raw log

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
func (it *BindingsStrikesParamsUnsetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsStrikesParamsUnset)
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
		it.Event = new(BindingsStrikesParamsUnset)
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
func (it *BindingsStrikesParamsUnsetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsStrikesParamsUnsetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsStrikesParamsUnset represents a StrikesParamsUnset event raised by the Bindings contract.
type BindingsStrikesParamsUnset struct {
	CurveId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStrikesParamsUnset is a free log retrieval operation binding the contract event 0xf1f344546b4e8f595b2ee9e9d26ed7baf40dc1fe8ef739b03280cb814c8e5df3.
//
// Solidity: event StrikesParamsUnset(uint256 indexed curveId)
func (_Bindings *BindingsFilterer) FilterStrikesParamsUnset(opts *bind.FilterOpts, curveId []*big.Int) (*BindingsStrikesParamsUnsetIterator, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "StrikesParamsUnset", curveIdRule)
	if err != nil {
		return nil, err
	}
	return &BindingsStrikesParamsUnsetIterator{contract: _Bindings.contract, event: "StrikesParamsUnset", logs: logs, sub: sub}, nil
}

// WatchStrikesParamsUnset is a free log subscription operation binding the contract event 0xf1f344546b4e8f595b2ee9e9d26ed7baf40dc1fe8ef739b03280cb814c8e5df3.
//
// Solidity: event StrikesParamsUnset(uint256 indexed curveId)
func (_Bindings *BindingsFilterer) WatchStrikesParamsUnset(opts *bind.WatchOpts, sink chan<- *BindingsStrikesParamsUnset, curveId []*big.Int) (event.Subscription, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "StrikesParamsUnset", curveIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsStrikesParamsUnset)
				if err := _Bindings.contract.UnpackLog(event, "StrikesParamsUnset", log); err != nil {
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

// ParseStrikesParamsUnset is a log parse operation binding the contract event 0xf1f344546b4e8f595b2ee9e9d26ed7baf40dc1fe8ef739b03280cb814c8e5df3.
//
// Solidity: event StrikesParamsUnset(uint256 indexed curveId)
func (_Bindings *BindingsFilterer) ParseStrikesParamsUnset(log types.Log) (*BindingsStrikesParamsUnset, error) {
	event := new(BindingsStrikesParamsUnset)
	if err := _Bindings.contract.UnpackLog(event, "StrikesParamsUnset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
