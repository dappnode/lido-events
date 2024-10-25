// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package csaccounting

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

// ICSAccountingPermitInput is an auto generated low-level Go binding around an user-defined struct.
type ICSAccountingPermitInput struct {
	Value    *big.Int
	Deadline *big.Int
	V        uint8
	R        [32]byte
	S        [32]byte
}

// ICSBondCurveBondCurve is an auto generated low-level Go binding around an user-defined struct.
type ICSBondCurveBondCurve struct {
	Points []*big.Int
	Trend  *big.Int
}

// ICSBondLockBondLock is an auto generated low-level Go binding around an user-defined struct.
type ICSBondLockBondLock struct {
	Amount         *big.Int
	RetentionUntil *big.Int
}

// CsaccountingMetaData contains all meta data concerning the Csaccounting contract.
var CsaccountingMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"lidoLocator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"communityStakingModule\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxCurveLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minBondLockRetentionPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxBondLockRetentionPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ACCOUNTING_MANAGER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CSM\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractICSModule\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_BOND_CURVE_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"LIDO\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractILido\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"LIDO_LOCATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractILidoLocator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MANAGE_BOND_CURVES_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_BOND_LOCK_RETENTION_PERIOD\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_CURVE_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_BOND_LOCK_RETENTION_PERIOD\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_CURVE_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSE_INFINITELY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSE_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RECOVERER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RESET_BOND_CURVE_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RESUME_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SET_BOND_CURVE_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WITHDRAWAL_QUEUE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIWithdrawalQueue\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WSTETH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIWstETH\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addBondCurve\",\"inputs\":[{\"name\":\"bondCurve\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"chargeFee\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"chargePenaltyRecipient\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimRewardsStETH\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"stETHAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rewardAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cumulativeFeeShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rewardsProof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimRewardsUnstETH\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"stEthAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rewardAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cumulativeFeeShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rewardsProof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimRewardsWstETH\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"wstETHAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rewardAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cumulativeFeeShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rewardsProof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"compensateLockedBondETH\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"depositETH\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"depositStETH\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"stETHAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"permit\",\"type\":\"tuple\",\"internalType\":\"structICSAccounting.PermitInput\",\"components\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositWstETH\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"wstETHAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"permit\",\"type\":\"tuple\",\"internalType\":\"structICSAccounting.PermitInput\",\"components\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"feeDistributor\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractICSFeeDistributor\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getActualLockedBond\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBond\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBondAmountByKeysCount\",\"inputs\":[{\"name\":\"keys\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"curve\",\"type\":\"tuple\",\"internalType\":\"structICSBondCurve.BondCurve\",\"components\":[{\"name\":\"points\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"trend\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getBondAmountByKeysCount\",\"inputs\":[{\"name\":\"keys\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"curveId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBondAmountByKeysCountWstETH\",\"inputs\":[{\"name\":\"keysCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"curveId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBondAmountByKeysCountWstETH\",\"inputs\":[{\"name\":\"keysCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"curve\",\"type\":\"tuple\",\"internalType\":\"structICSBondCurve.BondCurve\",\"components\":[{\"name\":\"points\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"trend\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBondCurve\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structICSBondCurve.BondCurve\",\"components\":[{\"name\":\"points\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"trend\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBondCurveId\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBondLockRetentionPeriod\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBondShares\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBondSummary\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"current\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"required\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBondSummaryShares\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"current\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"required\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurveInfo\",\"inputs\":[{\"name\":\"curveId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structICSBondCurve.BondCurve\",\"components\":[{\"name\":\"points\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"trend\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getKeysCountByBondAmount\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"curve\",\"type\":\"tuple\",\"internalType\":\"structICSBondCurve.BondCurve\",\"components\":[{\"name\":\"points\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"trend\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getKeysCountByBondAmount\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"curveId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLockedBondInfo\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structICSBondLock.BondLock\",\"components\":[{\"name\":\"amount\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"retentionUntil\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredBondForNextKeys\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"additionalKeys\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredBondForNextKeysWstETH\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"additionalKeys\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getResumeSinceTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMember\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMemberCount\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUnbondedKeysCount\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUnbondedKeysCountToEject\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"bondCurve\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_feeDistributor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bondLockRetentionPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_chargePenaltyRecipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isPaused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockBondETH\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseFor\",\"inputs\":[{\"name\":\"duration\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"penalize\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pullFeeRewards\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"cumulativeFeeShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rewardsProof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"recoverERC1155\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"recoverERC20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"recoverERC721\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"recoverEther\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"recoverStETHShares\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseLockedBondETH\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renewBurnerAllowance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resetBondCurve\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resume\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBondCurve\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"curveId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChargePenaltyRecipient\",\"inputs\":[{\"name\":\"_chargePenaltyRecipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setLockedBondRetentionPeriod\",\"inputs\":[{\"name\":\"retention\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"settleLockedBondETH\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalBondShares\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateBondCurve\",\"inputs\":[{\"name\":\"curveId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"bondCurve\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BondBurned\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"toBurnAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"burnedAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BondCharged\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"toChargeAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"chargedAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BondClaimedStETH\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BondClaimedUnstETH\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BondClaimedWstETH\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BondCurveAdded\",\"inputs\":[{\"name\":\"bondCurve\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BondCurveSet\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"curveId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BondCurveUpdated\",\"inputs\":[{\"name\":\"curveId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"bondCurve\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BondDepositedETH\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BondDepositedStETH\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BondDepositedWstETH\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BondLockChanged\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"newAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"retentionUntil\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BondLockCompensated\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BondLockRemoved\",\"inputs\":[{\"name\":\"nodeOperatorId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BondLockRetentionPeriodChanged\",\"inputs\":[{\"name\":\"retentionPeriod\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChargePenaltyRecipientSet\",\"inputs\":[{\"name\":\"chargePenaltyRecipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ERC1155Recovered\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ERC20Recovered\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ERC721Recovered\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EtherRecovered\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"duration\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Resumed\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StETHSharesRecovered\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ElRewardsVaultReceiveFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedToSendEther\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidBondCurveId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidBondCurveLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidBondCurveMaxLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidBondCurveValues\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidBondLockAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidBondLockRetentionPeriod\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialisationCurveId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NodeOperatorDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllowedToRecover\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NothingToClaim\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PauseUntilMustBeInFuture\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PausedExpected\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ResumedExpected\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintDowncast\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SenderIsNotCSM\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAdminAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroChargePenaltyRecipientAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroFeeDistributorAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroLocatorAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroModuleAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroPauseDuration\",\"inputs\":[]}]",
}

// CsaccountingABI is the input ABI used to generate the binding from.
// Deprecated: Use CsaccountingMetaData.ABI instead.
var CsaccountingABI = CsaccountingMetaData.ABI

// Csaccounting is an auto generated Go binding around an Ethereum contract.
type Csaccounting struct {
	CsaccountingCaller     // Read-only binding to the contract
	CsaccountingTransactor // Write-only binding to the contract
	CsaccountingFilterer   // Log filterer for contract events
}

// CsaccountingCaller is an auto generated read-only Go binding around an Ethereum contract.
type CsaccountingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CsaccountingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CsaccountingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CsaccountingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CsaccountingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CsaccountingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CsaccountingSession struct {
	Contract     *Csaccounting     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CsaccountingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CsaccountingCallerSession struct {
	Contract *CsaccountingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CsaccountingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CsaccountingTransactorSession struct {
	Contract     *CsaccountingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CsaccountingRaw is an auto generated low-level Go binding around an Ethereum contract.
type CsaccountingRaw struct {
	Contract *Csaccounting // Generic contract binding to access the raw methods on
}

// CsaccountingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CsaccountingCallerRaw struct {
	Contract *CsaccountingCaller // Generic read-only contract binding to access the raw methods on
}

// CsaccountingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CsaccountingTransactorRaw struct {
	Contract *CsaccountingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCsaccounting creates a new instance of Csaccounting, bound to a specific deployed contract.
func NewCsaccounting(address common.Address, backend bind.ContractBackend) (*Csaccounting, error) {
	contract, err := bindCsaccounting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Csaccounting{CsaccountingCaller: CsaccountingCaller{contract: contract}, CsaccountingTransactor: CsaccountingTransactor{contract: contract}, CsaccountingFilterer: CsaccountingFilterer{contract: contract}}, nil
}

// NewCsaccountingCaller creates a new read-only instance of Csaccounting, bound to a specific deployed contract.
func NewCsaccountingCaller(address common.Address, caller bind.ContractCaller) (*CsaccountingCaller, error) {
	contract, err := bindCsaccounting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CsaccountingCaller{contract: contract}, nil
}

// NewCsaccountingTransactor creates a new write-only instance of Csaccounting, bound to a specific deployed contract.
func NewCsaccountingTransactor(address common.Address, transactor bind.ContractTransactor) (*CsaccountingTransactor, error) {
	contract, err := bindCsaccounting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CsaccountingTransactor{contract: contract}, nil
}

// NewCsaccountingFilterer creates a new log filterer instance of Csaccounting, bound to a specific deployed contract.
func NewCsaccountingFilterer(address common.Address, filterer bind.ContractFilterer) (*CsaccountingFilterer, error) {
	contract, err := bindCsaccounting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CsaccountingFilterer{contract: contract}, nil
}

// bindCsaccounting binds a generic wrapper to an already deployed contract.
func bindCsaccounting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CsaccountingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Csaccounting *CsaccountingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Csaccounting.Contract.CsaccountingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Csaccounting *CsaccountingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Csaccounting.Contract.CsaccountingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Csaccounting *CsaccountingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Csaccounting.Contract.CsaccountingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Csaccounting *CsaccountingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Csaccounting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Csaccounting *CsaccountingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Csaccounting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Csaccounting *CsaccountingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Csaccounting.Contract.contract.Transact(opts, method, params...)
}

// ACCOUNTINGMANAGERROLE is a free data retrieval call binding the contract method 0x8ed5c5d7.
//
// Solidity: function ACCOUNTING_MANAGER_ROLE() view returns(bytes32)
func (_Csaccounting *CsaccountingCaller) ACCOUNTINGMANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "ACCOUNTING_MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ACCOUNTINGMANAGERROLE is a free data retrieval call binding the contract method 0x8ed5c5d7.
//
// Solidity: function ACCOUNTING_MANAGER_ROLE() view returns(bytes32)
func (_Csaccounting *CsaccountingSession) ACCOUNTINGMANAGERROLE() ([32]byte, error) {
	return _Csaccounting.Contract.ACCOUNTINGMANAGERROLE(&_Csaccounting.CallOpts)
}

// ACCOUNTINGMANAGERROLE is a free data retrieval call binding the contract method 0x8ed5c5d7.
//
// Solidity: function ACCOUNTING_MANAGER_ROLE() view returns(bytes32)
func (_Csaccounting *CsaccountingCallerSession) ACCOUNTINGMANAGERROLE() ([32]byte, error) {
	return _Csaccounting.Contract.ACCOUNTINGMANAGERROLE(&_Csaccounting.CallOpts)
}

// CSM is a free data retrieval call binding the contract method 0x8de2b272.
//
// Solidity: function CSM() view returns(address)
func (_Csaccounting *CsaccountingCaller) CSM(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "CSM")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CSM is a free data retrieval call binding the contract method 0x8de2b272.
//
// Solidity: function CSM() view returns(address)
func (_Csaccounting *CsaccountingSession) CSM() (common.Address, error) {
	return _Csaccounting.Contract.CSM(&_Csaccounting.CallOpts)
}

// CSM is a free data retrieval call binding the contract method 0x8de2b272.
//
// Solidity: function CSM() view returns(address)
func (_Csaccounting *CsaccountingCallerSession) CSM() (common.Address, error) {
	return _Csaccounting.Contract.CSM(&_Csaccounting.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Csaccounting *CsaccountingCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Csaccounting *CsaccountingSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Csaccounting.Contract.DEFAULTADMINROLE(&_Csaccounting.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Csaccounting *CsaccountingCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Csaccounting.Contract.DEFAULTADMINROLE(&_Csaccounting.CallOpts)
}

// DEFAULTBONDCURVEID is a free data retrieval call binding the contract method 0x443fbfef.
//
// Solidity: function DEFAULT_BOND_CURVE_ID() view returns(uint256)
func (_Csaccounting *CsaccountingCaller) DEFAULTBONDCURVEID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "DEFAULT_BOND_CURVE_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEFAULTBONDCURVEID is a free data retrieval call binding the contract method 0x443fbfef.
//
// Solidity: function DEFAULT_BOND_CURVE_ID() view returns(uint256)
func (_Csaccounting *CsaccountingSession) DEFAULTBONDCURVEID() (*big.Int, error) {
	return _Csaccounting.Contract.DEFAULTBONDCURVEID(&_Csaccounting.CallOpts)
}

// DEFAULTBONDCURVEID is a free data retrieval call binding the contract method 0x443fbfef.
//
// Solidity: function DEFAULT_BOND_CURVE_ID() view returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) DEFAULTBONDCURVEID() (*big.Int, error) {
	return _Csaccounting.Contract.DEFAULTBONDCURVEID(&_Csaccounting.CallOpts)
}

// LIDO is a free data retrieval call binding the contract method 0x8b21f170.
//
// Solidity: function LIDO() view returns(address)
func (_Csaccounting *CsaccountingCaller) LIDO(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "LIDO")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LIDO is a free data retrieval call binding the contract method 0x8b21f170.
//
// Solidity: function LIDO() view returns(address)
func (_Csaccounting *CsaccountingSession) LIDO() (common.Address, error) {
	return _Csaccounting.Contract.LIDO(&_Csaccounting.CallOpts)
}

// LIDO is a free data retrieval call binding the contract method 0x8b21f170.
//
// Solidity: function LIDO() view returns(address)
func (_Csaccounting *CsaccountingCallerSession) LIDO() (common.Address, error) {
	return _Csaccounting.Contract.LIDO(&_Csaccounting.CallOpts)
}

// LIDOLOCATOR is a free data retrieval call binding the contract method 0xdbba4b48.
//
// Solidity: function LIDO_LOCATOR() view returns(address)
func (_Csaccounting *CsaccountingCaller) LIDOLOCATOR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "LIDO_LOCATOR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LIDOLOCATOR is a free data retrieval call binding the contract method 0xdbba4b48.
//
// Solidity: function LIDO_LOCATOR() view returns(address)
func (_Csaccounting *CsaccountingSession) LIDOLOCATOR() (common.Address, error) {
	return _Csaccounting.Contract.LIDOLOCATOR(&_Csaccounting.CallOpts)
}

// LIDOLOCATOR is a free data retrieval call binding the contract method 0xdbba4b48.
//
// Solidity: function LIDO_LOCATOR() view returns(address)
func (_Csaccounting *CsaccountingCallerSession) LIDOLOCATOR() (common.Address, error) {
	return _Csaccounting.Contract.LIDOLOCATOR(&_Csaccounting.CallOpts)
}

// MANAGEBONDCURVESROLE is a free data retrieval call binding the contract method 0xfee63805.
//
// Solidity: function MANAGE_BOND_CURVES_ROLE() view returns(bytes32)
func (_Csaccounting *CsaccountingCaller) MANAGEBONDCURVESROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "MANAGE_BOND_CURVES_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGEBONDCURVESROLE is a free data retrieval call binding the contract method 0xfee63805.
//
// Solidity: function MANAGE_BOND_CURVES_ROLE() view returns(bytes32)
func (_Csaccounting *CsaccountingSession) MANAGEBONDCURVESROLE() ([32]byte, error) {
	return _Csaccounting.Contract.MANAGEBONDCURVESROLE(&_Csaccounting.CallOpts)
}

// MANAGEBONDCURVESROLE is a free data retrieval call binding the contract method 0xfee63805.
//
// Solidity: function MANAGE_BOND_CURVES_ROLE() view returns(bytes32)
func (_Csaccounting *CsaccountingCallerSession) MANAGEBONDCURVESROLE() ([32]byte, error) {
	return _Csaccounting.Contract.MANAGEBONDCURVESROLE(&_Csaccounting.CallOpts)
}

// MAXBONDLOCKRETENTIONPERIOD is a free data retrieval call binding the contract method 0x526352fc.
//
// Solidity: function MAX_BOND_LOCK_RETENTION_PERIOD() view returns(uint256)
func (_Csaccounting *CsaccountingCaller) MAXBONDLOCKRETENTIONPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "MAX_BOND_LOCK_RETENTION_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBONDLOCKRETENTIONPERIOD is a free data retrieval call binding the contract method 0x526352fc.
//
// Solidity: function MAX_BOND_LOCK_RETENTION_PERIOD() view returns(uint256)
func (_Csaccounting *CsaccountingSession) MAXBONDLOCKRETENTIONPERIOD() (*big.Int, error) {
	return _Csaccounting.Contract.MAXBONDLOCKRETENTIONPERIOD(&_Csaccounting.CallOpts)
}

// MAXBONDLOCKRETENTIONPERIOD is a free data retrieval call binding the contract method 0x526352fc.
//
// Solidity: function MAX_BOND_LOCK_RETENTION_PERIOD() view returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) MAXBONDLOCKRETENTIONPERIOD() (*big.Int, error) {
	return _Csaccounting.Contract.MAXBONDLOCKRETENTIONPERIOD(&_Csaccounting.CallOpts)
}

// MAXCURVELENGTH is a free data retrieval call binding the contract method 0x4b2ce9fe.
//
// Solidity: function MAX_CURVE_LENGTH() view returns(uint256)
func (_Csaccounting *CsaccountingCaller) MAXCURVELENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "MAX_CURVE_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXCURVELENGTH is a free data retrieval call binding the contract method 0x4b2ce9fe.
//
// Solidity: function MAX_CURVE_LENGTH() view returns(uint256)
func (_Csaccounting *CsaccountingSession) MAXCURVELENGTH() (*big.Int, error) {
	return _Csaccounting.Contract.MAXCURVELENGTH(&_Csaccounting.CallOpts)
}

// MAXCURVELENGTH is a free data retrieval call binding the contract method 0x4b2ce9fe.
//
// Solidity: function MAX_CURVE_LENGTH() view returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) MAXCURVELENGTH() (*big.Int, error) {
	return _Csaccounting.Contract.MAXCURVELENGTH(&_Csaccounting.CallOpts)
}

// MINBONDLOCKRETENTIONPERIOD is a free data retrieval call binding the contract method 0xae849756.
//
// Solidity: function MIN_BOND_LOCK_RETENTION_PERIOD() view returns(uint256)
func (_Csaccounting *CsaccountingCaller) MINBONDLOCKRETENTIONPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "MIN_BOND_LOCK_RETENTION_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINBONDLOCKRETENTIONPERIOD is a free data retrieval call binding the contract method 0xae849756.
//
// Solidity: function MIN_BOND_LOCK_RETENTION_PERIOD() view returns(uint256)
func (_Csaccounting *CsaccountingSession) MINBONDLOCKRETENTIONPERIOD() (*big.Int, error) {
	return _Csaccounting.Contract.MINBONDLOCKRETENTIONPERIOD(&_Csaccounting.CallOpts)
}

// MINBONDLOCKRETENTIONPERIOD is a free data retrieval call binding the contract method 0xae849756.
//
// Solidity: function MIN_BOND_LOCK_RETENTION_PERIOD() view returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) MINBONDLOCKRETENTIONPERIOD() (*big.Int, error) {
	return _Csaccounting.Contract.MINBONDLOCKRETENTIONPERIOD(&_Csaccounting.CallOpts)
}

// MINCURVELENGTH is a free data retrieval call binding the contract method 0xcb11c527.
//
// Solidity: function MIN_CURVE_LENGTH() view returns(uint256)
func (_Csaccounting *CsaccountingCaller) MINCURVELENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "MIN_CURVE_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINCURVELENGTH is a free data retrieval call binding the contract method 0xcb11c527.
//
// Solidity: function MIN_CURVE_LENGTH() view returns(uint256)
func (_Csaccounting *CsaccountingSession) MINCURVELENGTH() (*big.Int, error) {
	return _Csaccounting.Contract.MINCURVELENGTH(&_Csaccounting.CallOpts)
}

// MINCURVELENGTH is a free data retrieval call binding the contract method 0xcb11c527.
//
// Solidity: function MIN_CURVE_LENGTH() view returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) MINCURVELENGTH() (*big.Int, error) {
	return _Csaccounting.Contract.MINCURVELENGTH(&_Csaccounting.CallOpts)
}

// PAUSEINFINITELY is a free data retrieval call binding the contract method 0xa302ee38.
//
// Solidity: function PAUSE_INFINITELY() view returns(uint256)
func (_Csaccounting *CsaccountingCaller) PAUSEINFINITELY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "PAUSE_INFINITELY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PAUSEINFINITELY is a free data retrieval call binding the contract method 0xa302ee38.
//
// Solidity: function PAUSE_INFINITELY() view returns(uint256)
func (_Csaccounting *CsaccountingSession) PAUSEINFINITELY() (*big.Int, error) {
	return _Csaccounting.Contract.PAUSEINFINITELY(&_Csaccounting.CallOpts)
}

// PAUSEINFINITELY is a free data retrieval call binding the contract method 0xa302ee38.
//
// Solidity: function PAUSE_INFINITELY() view returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) PAUSEINFINITELY() (*big.Int, error) {
	return _Csaccounting.Contract.PAUSEINFINITELY(&_Csaccounting.CallOpts)
}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_Csaccounting *CsaccountingCaller) PAUSEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "PAUSE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_Csaccounting *CsaccountingSession) PAUSEROLE() ([32]byte, error) {
	return _Csaccounting.Contract.PAUSEROLE(&_Csaccounting.CallOpts)
}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_Csaccounting *CsaccountingCallerSession) PAUSEROLE() ([32]byte, error) {
	return _Csaccounting.Contract.PAUSEROLE(&_Csaccounting.CallOpts)
}

// RECOVERERROLE is a free data retrieval call binding the contract method 0xacf1c948.
//
// Solidity: function RECOVERER_ROLE() view returns(bytes32)
func (_Csaccounting *CsaccountingCaller) RECOVERERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "RECOVERER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RECOVERERROLE is a free data retrieval call binding the contract method 0xacf1c948.
//
// Solidity: function RECOVERER_ROLE() view returns(bytes32)
func (_Csaccounting *CsaccountingSession) RECOVERERROLE() ([32]byte, error) {
	return _Csaccounting.Contract.RECOVERERROLE(&_Csaccounting.CallOpts)
}

// RECOVERERROLE is a free data retrieval call binding the contract method 0xacf1c948.
//
// Solidity: function RECOVERER_ROLE() view returns(bytes32)
func (_Csaccounting *CsaccountingCallerSession) RECOVERERROLE() ([32]byte, error) {
	return _Csaccounting.Contract.RECOVERERROLE(&_Csaccounting.CallOpts)
}

// RESETBONDCURVEROLE is a free data retrieval call binding the contract method 0x21d439d5.
//
// Solidity: function RESET_BOND_CURVE_ROLE() view returns(bytes32)
func (_Csaccounting *CsaccountingCaller) RESETBONDCURVEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "RESET_BOND_CURVE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RESETBONDCURVEROLE is a free data retrieval call binding the contract method 0x21d439d5.
//
// Solidity: function RESET_BOND_CURVE_ROLE() view returns(bytes32)
func (_Csaccounting *CsaccountingSession) RESETBONDCURVEROLE() ([32]byte, error) {
	return _Csaccounting.Contract.RESETBONDCURVEROLE(&_Csaccounting.CallOpts)
}

// RESETBONDCURVEROLE is a free data retrieval call binding the contract method 0x21d439d5.
//
// Solidity: function RESET_BOND_CURVE_ROLE() view returns(bytes32)
func (_Csaccounting *CsaccountingCallerSession) RESETBONDCURVEROLE() ([32]byte, error) {
	return _Csaccounting.Contract.RESETBONDCURVEROLE(&_Csaccounting.CallOpts)
}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_Csaccounting *CsaccountingCaller) RESUMEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "RESUME_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_Csaccounting *CsaccountingSession) RESUMEROLE() ([32]byte, error) {
	return _Csaccounting.Contract.RESUMEROLE(&_Csaccounting.CallOpts)
}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_Csaccounting *CsaccountingCallerSession) RESUMEROLE() ([32]byte, error) {
	return _Csaccounting.Contract.RESUMEROLE(&_Csaccounting.CallOpts)
}

// SETBONDCURVEROLE is a free data retrieval call binding the contract method 0x4342b3c1.
//
// Solidity: function SET_BOND_CURVE_ROLE() view returns(bytes32)
func (_Csaccounting *CsaccountingCaller) SETBONDCURVEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "SET_BOND_CURVE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SETBONDCURVEROLE is a free data retrieval call binding the contract method 0x4342b3c1.
//
// Solidity: function SET_BOND_CURVE_ROLE() view returns(bytes32)
func (_Csaccounting *CsaccountingSession) SETBONDCURVEROLE() ([32]byte, error) {
	return _Csaccounting.Contract.SETBONDCURVEROLE(&_Csaccounting.CallOpts)
}

// SETBONDCURVEROLE is a free data retrieval call binding the contract method 0x4342b3c1.
//
// Solidity: function SET_BOND_CURVE_ROLE() view returns(bytes32)
func (_Csaccounting *CsaccountingCallerSession) SETBONDCURVEROLE() ([32]byte, error) {
	return _Csaccounting.Contract.SETBONDCURVEROLE(&_Csaccounting.CallOpts)
}

// WITHDRAWALQUEUE is a free data retrieval call binding the contract method 0x699340f4.
//
// Solidity: function WITHDRAWAL_QUEUE() view returns(address)
func (_Csaccounting *CsaccountingCaller) WITHDRAWALQUEUE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "WITHDRAWAL_QUEUE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WITHDRAWALQUEUE is a free data retrieval call binding the contract method 0x699340f4.
//
// Solidity: function WITHDRAWAL_QUEUE() view returns(address)
func (_Csaccounting *CsaccountingSession) WITHDRAWALQUEUE() (common.Address, error) {
	return _Csaccounting.Contract.WITHDRAWALQUEUE(&_Csaccounting.CallOpts)
}

// WITHDRAWALQUEUE is a free data retrieval call binding the contract method 0x699340f4.
//
// Solidity: function WITHDRAWAL_QUEUE() view returns(address)
func (_Csaccounting *CsaccountingCallerSession) WITHDRAWALQUEUE() (common.Address, error) {
	return _Csaccounting.Contract.WITHDRAWALQUEUE(&_Csaccounting.CallOpts)
}

// WSTETH is a free data retrieval call binding the contract method 0xd9fb643a.
//
// Solidity: function WSTETH() view returns(address)
func (_Csaccounting *CsaccountingCaller) WSTETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "WSTETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WSTETH is a free data retrieval call binding the contract method 0xd9fb643a.
//
// Solidity: function WSTETH() view returns(address)
func (_Csaccounting *CsaccountingSession) WSTETH() (common.Address, error) {
	return _Csaccounting.Contract.WSTETH(&_Csaccounting.CallOpts)
}

// WSTETH is a free data retrieval call binding the contract method 0xd9fb643a.
//
// Solidity: function WSTETH() view returns(address)
func (_Csaccounting *CsaccountingCallerSession) WSTETH() (common.Address, error) {
	return _Csaccounting.Contract.WSTETH(&_Csaccounting.CallOpts)
}

// ChargePenaltyRecipient is a free data retrieval call binding the contract method 0x165123dd.
//
// Solidity: function chargePenaltyRecipient() view returns(address)
func (_Csaccounting *CsaccountingCaller) ChargePenaltyRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "chargePenaltyRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChargePenaltyRecipient is a free data retrieval call binding the contract method 0x165123dd.
//
// Solidity: function chargePenaltyRecipient() view returns(address)
func (_Csaccounting *CsaccountingSession) ChargePenaltyRecipient() (common.Address, error) {
	return _Csaccounting.Contract.ChargePenaltyRecipient(&_Csaccounting.CallOpts)
}

// ChargePenaltyRecipient is a free data retrieval call binding the contract method 0x165123dd.
//
// Solidity: function chargePenaltyRecipient() view returns(address)
func (_Csaccounting *CsaccountingCallerSession) ChargePenaltyRecipient() (common.Address, error) {
	return _Csaccounting.Contract.ChargePenaltyRecipient(&_Csaccounting.CallOpts)
}

// FeeDistributor is a free data retrieval call binding the contract method 0x0d43e8ad.
//
// Solidity: function feeDistributor() view returns(address)
func (_Csaccounting *CsaccountingCaller) FeeDistributor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "feeDistributor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeDistributor is a free data retrieval call binding the contract method 0x0d43e8ad.
//
// Solidity: function feeDistributor() view returns(address)
func (_Csaccounting *CsaccountingSession) FeeDistributor() (common.Address, error) {
	return _Csaccounting.Contract.FeeDistributor(&_Csaccounting.CallOpts)
}

// FeeDistributor is a free data retrieval call binding the contract method 0x0d43e8ad.
//
// Solidity: function feeDistributor() view returns(address)
func (_Csaccounting *CsaccountingCallerSession) FeeDistributor() (common.Address, error) {
	return _Csaccounting.Contract.FeeDistributor(&_Csaccounting.CallOpts)
}

// GetActualLockedBond is a free data retrieval call binding the contract method 0xead42a69.
//
// Solidity: function getActualLockedBond(uint256 nodeOperatorId) view returns(uint256)
func (_Csaccounting *CsaccountingCaller) GetActualLockedBond(opts *bind.CallOpts, nodeOperatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getActualLockedBond", nodeOperatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActualLockedBond is a free data retrieval call binding the contract method 0xead42a69.
//
// Solidity: function getActualLockedBond(uint256 nodeOperatorId) view returns(uint256)
func (_Csaccounting *CsaccountingSession) GetActualLockedBond(nodeOperatorId *big.Int) (*big.Int, error) {
	return _Csaccounting.Contract.GetActualLockedBond(&_Csaccounting.CallOpts, nodeOperatorId)
}

// GetActualLockedBond is a free data retrieval call binding the contract method 0xead42a69.
//
// Solidity: function getActualLockedBond(uint256 nodeOperatorId) view returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) GetActualLockedBond(nodeOperatorId *big.Int) (*big.Int, error) {
	return _Csaccounting.Contract.GetActualLockedBond(&_Csaccounting.CallOpts, nodeOperatorId)
}

// GetBond is a free data retrieval call binding the contract method 0xd8fe7642.
//
// Solidity: function getBond(uint256 nodeOperatorId) view returns(uint256)
func (_Csaccounting *CsaccountingCaller) GetBond(opts *bind.CallOpts, nodeOperatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getBond", nodeOperatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBond is a free data retrieval call binding the contract method 0xd8fe7642.
//
// Solidity: function getBond(uint256 nodeOperatorId) view returns(uint256)
func (_Csaccounting *CsaccountingSession) GetBond(nodeOperatorId *big.Int) (*big.Int, error) {
	return _Csaccounting.Contract.GetBond(&_Csaccounting.CallOpts, nodeOperatorId)
}

// GetBond is a free data retrieval call binding the contract method 0xd8fe7642.
//
// Solidity: function getBond(uint256 nodeOperatorId) view returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) GetBond(nodeOperatorId *big.Int) (*big.Int, error) {
	return _Csaccounting.Contract.GetBond(&_Csaccounting.CallOpts, nodeOperatorId)
}

// GetBondAmountByKeysCount is a free data retrieval call binding the contract method 0x0f23e742.
//
// Solidity: function getBondAmountByKeysCount(uint256 keys, (uint256[],uint256) curve) pure returns(uint256)
func (_Csaccounting *CsaccountingCaller) GetBondAmountByKeysCount(opts *bind.CallOpts, keys *big.Int, curve ICSBondCurveBondCurve) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getBondAmountByKeysCount", keys, curve)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBondAmountByKeysCount is a free data retrieval call binding the contract method 0x0f23e742.
//
// Solidity: function getBondAmountByKeysCount(uint256 keys, (uint256[],uint256) curve) pure returns(uint256)
func (_Csaccounting *CsaccountingSession) GetBondAmountByKeysCount(keys *big.Int, curve ICSBondCurveBondCurve) (*big.Int, error) {
	return _Csaccounting.Contract.GetBondAmountByKeysCount(&_Csaccounting.CallOpts, keys, curve)
}

// GetBondAmountByKeysCount is a free data retrieval call binding the contract method 0x0f23e742.
//
// Solidity: function getBondAmountByKeysCount(uint256 keys, (uint256[],uint256) curve) pure returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) GetBondAmountByKeysCount(keys *big.Int, curve ICSBondCurveBondCurve) (*big.Int, error) {
	return _Csaccounting.Contract.GetBondAmountByKeysCount(&_Csaccounting.CallOpts, keys, curve)
}

// GetBondAmountByKeysCount0 is a free data retrieval call binding the contract method 0x546da24f.
//
// Solidity: function getBondAmountByKeysCount(uint256 keys, uint256 curveId) view returns(uint256)
func (_Csaccounting *CsaccountingCaller) GetBondAmountByKeysCount0(opts *bind.CallOpts, keys *big.Int, curveId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getBondAmountByKeysCount0", keys, curveId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBondAmountByKeysCount0 is a free data retrieval call binding the contract method 0x546da24f.
//
// Solidity: function getBondAmountByKeysCount(uint256 keys, uint256 curveId) view returns(uint256)
func (_Csaccounting *CsaccountingSession) GetBondAmountByKeysCount0(keys *big.Int, curveId *big.Int) (*big.Int, error) {
	return _Csaccounting.Contract.GetBondAmountByKeysCount0(&_Csaccounting.CallOpts, keys, curveId)
}

// GetBondAmountByKeysCount0 is a free data retrieval call binding the contract method 0x546da24f.
//
// Solidity: function getBondAmountByKeysCount(uint256 keys, uint256 curveId) view returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) GetBondAmountByKeysCount0(keys *big.Int, curveId *big.Int) (*big.Int, error) {
	return _Csaccounting.Contract.GetBondAmountByKeysCount0(&_Csaccounting.CallOpts, keys, curveId)
}

// GetBondAmountByKeysCountWstETH is a free data retrieval call binding the contract method 0x13d1234b.
//
// Solidity: function getBondAmountByKeysCountWstETH(uint256 keysCount, uint256 curveId) view returns(uint256)
func (_Csaccounting *CsaccountingCaller) GetBondAmountByKeysCountWstETH(opts *bind.CallOpts, keysCount *big.Int, curveId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getBondAmountByKeysCountWstETH", keysCount, curveId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBondAmountByKeysCountWstETH is a free data retrieval call binding the contract method 0x13d1234b.
//
// Solidity: function getBondAmountByKeysCountWstETH(uint256 keysCount, uint256 curveId) view returns(uint256)
func (_Csaccounting *CsaccountingSession) GetBondAmountByKeysCountWstETH(keysCount *big.Int, curveId *big.Int) (*big.Int, error) {
	return _Csaccounting.Contract.GetBondAmountByKeysCountWstETH(&_Csaccounting.CallOpts, keysCount, curveId)
}

// GetBondAmountByKeysCountWstETH is a free data retrieval call binding the contract method 0x13d1234b.
//
// Solidity: function getBondAmountByKeysCountWstETH(uint256 keysCount, uint256 curveId) view returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) GetBondAmountByKeysCountWstETH(keysCount *big.Int, curveId *big.Int) (*big.Int, error) {
	return _Csaccounting.Contract.GetBondAmountByKeysCountWstETH(&_Csaccounting.CallOpts, keysCount, curveId)
}

// GetBondAmountByKeysCountWstETH0 is a free data retrieval call binding the contract method 0x9a4df8f0.
//
// Solidity: function getBondAmountByKeysCountWstETH(uint256 keysCount, (uint256[],uint256) curve) view returns(uint256)
func (_Csaccounting *CsaccountingCaller) GetBondAmountByKeysCountWstETH0(opts *bind.CallOpts, keysCount *big.Int, curve ICSBondCurveBondCurve) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getBondAmountByKeysCountWstETH0", keysCount, curve)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBondAmountByKeysCountWstETH0 is a free data retrieval call binding the contract method 0x9a4df8f0.
//
// Solidity: function getBondAmountByKeysCountWstETH(uint256 keysCount, (uint256[],uint256) curve) view returns(uint256)
func (_Csaccounting *CsaccountingSession) GetBondAmountByKeysCountWstETH0(keysCount *big.Int, curve ICSBondCurveBondCurve) (*big.Int, error) {
	return _Csaccounting.Contract.GetBondAmountByKeysCountWstETH0(&_Csaccounting.CallOpts, keysCount, curve)
}

// GetBondAmountByKeysCountWstETH0 is a free data retrieval call binding the contract method 0x9a4df8f0.
//
// Solidity: function getBondAmountByKeysCountWstETH(uint256 keysCount, (uint256[],uint256) curve) view returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) GetBondAmountByKeysCountWstETH0(keysCount *big.Int, curve ICSBondCurveBondCurve) (*big.Int, error) {
	return _Csaccounting.Contract.GetBondAmountByKeysCountWstETH0(&_Csaccounting.CallOpts, keysCount, curve)
}

// GetBondCurve is a free data retrieval call binding the contract method 0x6e13f099.
//
// Solidity: function getBondCurve(uint256 nodeOperatorId) view returns((uint256[],uint256))
func (_Csaccounting *CsaccountingCaller) GetBondCurve(opts *bind.CallOpts, nodeOperatorId *big.Int) (ICSBondCurveBondCurve, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getBondCurve", nodeOperatorId)

	if err != nil {
		return *new(ICSBondCurveBondCurve), err
	}

	out0 := *abi.ConvertType(out[0], new(ICSBondCurveBondCurve)).(*ICSBondCurveBondCurve)

	return out0, err

}

// GetBondCurve is a free data retrieval call binding the contract method 0x6e13f099.
//
// Solidity: function getBondCurve(uint256 nodeOperatorId) view returns((uint256[],uint256))
func (_Csaccounting *CsaccountingSession) GetBondCurve(nodeOperatorId *big.Int) (ICSBondCurveBondCurve, error) {
	return _Csaccounting.Contract.GetBondCurve(&_Csaccounting.CallOpts, nodeOperatorId)
}

// GetBondCurve is a free data retrieval call binding the contract method 0x6e13f099.
//
// Solidity: function getBondCurve(uint256 nodeOperatorId) view returns((uint256[],uint256))
func (_Csaccounting *CsaccountingCallerSession) GetBondCurve(nodeOperatorId *big.Int) (ICSBondCurveBondCurve, error) {
	return _Csaccounting.Contract.GetBondCurve(&_Csaccounting.CallOpts, nodeOperatorId)
}

// GetBondCurveId is a free data retrieval call binding the contract method 0x0569b947.
//
// Solidity: function getBondCurveId(uint256 nodeOperatorId) view returns(uint256)
func (_Csaccounting *CsaccountingCaller) GetBondCurveId(opts *bind.CallOpts, nodeOperatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getBondCurveId", nodeOperatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBondCurveId is a free data retrieval call binding the contract method 0x0569b947.
//
// Solidity: function getBondCurveId(uint256 nodeOperatorId) view returns(uint256)
func (_Csaccounting *CsaccountingSession) GetBondCurveId(nodeOperatorId *big.Int) (*big.Int, error) {
	return _Csaccounting.Contract.GetBondCurveId(&_Csaccounting.CallOpts, nodeOperatorId)
}

// GetBondCurveId is a free data retrieval call binding the contract method 0x0569b947.
//
// Solidity: function getBondCurveId(uint256 nodeOperatorId) view returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) GetBondCurveId(nodeOperatorId *big.Int) (*big.Int, error) {
	return _Csaccounting.Contract.GetBondCurveId(&_Csaccounting.CallOpts, nodeOperatorId)
}

// GetBondLockRetentionPeriod is a free data retrieval call binding the contract method 0xdef82d02.
//
// Solidity: function getBondLockRetentionPeriod() view returns(uint256)
func (_Csaccounting *CsaccountingCaller) GetBondLockRetentionPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getBondLockRetentionPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBondLockRetentionPeriod is a free data retrieval call binding the contract method 0xdef82d02.
//
// Solidity: function getBondLockRetentionPeriod() view returns(uint256)
func (_Csaccounting *CsaccountingSession) GetBondLockRetentionPeriod() (*big.Int, error) {
	return _Csaccounting.Contract.GetBondLockRetentionPeriod(&_Csaccounting.CallOpts)
}

// GetBondLockRetentionPeriod is a free data retrieval call binding the contract method 0xdef82d02.
//
// Solidity: function getBondLockRetentionPeriod() view returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) GetBondLockRetentionPeriod() (*big.Int, error) {
	return _Csaccounting.Contract.GetBondLockRetentionPeriod(&_Csaccounting.CallOpts)
}

// GetBondShares is a free data retrieval call binding the contract method 0x06cd0e90.
//
// Solidity: function getBondShares(uint256 nodeOperatorId) view returns(uint256)
func (_Csaccounting *CsaccountingCaller) GetBondShares(opts *bind.CallOpts, nodeOperatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getBondShares", nodeOperatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBondShares is a free data retrieval call binding the contract method 0x06cd0e90.
//
// Solidity: function getBondShares(uint256 nodeOperatorId) view returns(uint256)
func (_Csaccounting *CsaccountingSession) GetBondShares(nodeOperatorId *big.Int) (*big.Int, error) {
	return _Csaccounting.Contract.GetBondShares(&_Csaccounting.CallOpts, nodeOperatorId)
}

// GetBondShares is a free data retrieval call binding the contract method 0x06cd0e90.
//
// Solidity: function getBondShares(uint256 nodeOperatorId) view returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) GetBondShares(nodeOperatorId *big.Int) (*big.Int, error) {
	return _Csaccounting.Contract.GetBondShares(&_Csaccounting.CallOpts, nodeOperatorId)
}

// GetBondSummary is a free data retrieval call binding the contract method 0xce19793f.
//
// Solidity: function getBondSummary(uint256 nodeOperatorId) view returns(uint256 current, uint256 required)
func (_Csaccounting *CsaccountingCaller) GetBondSummary(opts *bind.CallOpts, nodeOperatorId *big.Int) (struct {
	Current  *big.Int
	Required *big.Int
}, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getBondSummary", nodeOperatorId)

	outstruct := new(struct {
		Current  *big.Int
		Required *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Current = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Required = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetBondSummary is a free data retrieval call binding the contract method 0xce19793f.
//
// Solidity: function getBondSummary(uint256 nodeOperatorId) view returns(uint256 current, uint256 required)
func (_Csaccounting *CsaccountingSession) GetBondSummary(nodeOperatorId *big.Int) (struct {
	Current  *big.Int
	Required *big.Int
}, error) {
	return _Csaccounting.Contract.GetBondSummary(&_Csaccounting.CallOpts, nodeOperatorId)
}

// GetBondSummary is a free data retrieval call binding the contract method 0xce19793f.
//
// Solidity: function getBondSummary(uint256 nodeOperatorId) view returns(uint256 current, uint256 required)
func (_Csaccounting *CsaccountingCallerSession) GetBondSummary(nodeOperatorId *big.Int) (struct {
	Current  *big.Int
	Required *big.Int
}, error) {
	return _Csaccounting.Contract.GetBondSummary(&_Csaccounting.CallOpts, nodeOperatorId)
}

// GetBondSummaryShares is a free data retrieval call binding the contract method 0x8f6549ae.
//
// Solidity: function getBondSummaryShares(uint256 nodeOperatorId) view returns(uint256 current, uint256 required)
func (_Csaccounting *CsaccountingCaller) GetBondSummaryShares(opts *bind.CallOpts, nodeOperatorId *big.Int) (struct {
	Current  *big.Int
	Required *big.Int
}, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getBondSummaryShares", nodeOperatorId)

	outstruct := new(struct {
		Current  *big.Int
		Required *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Current = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Required = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetBondSummaryShares is a free data retrieval call binding the contract method 0x8f6549ae.
//
// Solidity: function getBondSummaryShares(uint256 nodeOperatorId) view returns(uint256 current, uint256 required)
func (_Csaccounting *CsaccountingSession) GetBondSummaryShares(nodeOperatorId *big.Int) (struct {
	Current  *big.Int
	Required *big.Int
}, error) {
	return _Csaccounting.Contract.GetBondSummaryShares(&_Csaccounting.CallOpts, nodeOperatorId)
}

// GetBondSummaryShares is a free data retrieval call binding the contract method 0x8f6549ae.
//
// Solidity: function getBondSummaryShares(uint256 nodeOperatorId) view returns(uint256 current, uint256 required)
func (_Csaccounting *CsaccountingCallerSession) GetBondSummaryShares(nodeOperatorId *big.Int) (struct {
	Current  *big.Int
	Required *big.Int
}, error) {
	return _Csaccounting.Contract.GetBondSummaryShares(&_Csaccounting.CallOpts, nodeOperatorId)
}

// GetCurveInfo is a free data retrieval call binding the contract method 0xb5b624bf.
//
// Solidity: function getCurveInfo(uint256 curveId) view returns((uint256[],uint256))
func (_Csaccounting *CsaccountingCaller) GetCurveInfo(opts *bind.CallOpts, curveId *big.Int) (ICSBondCurveBondCurve, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getCurveInfo", curveId)

	if err != nil {
		return *new(ICSBondCurveBondCurve), err
	}

	out0 := *abi.ConvertType(out[0], new(ICSBondCurveBondCurve)).(*ICSBondCurveBondCurve)

	return out0, err

}

// GetCurveInfo is a free data retrieval call binding the contract method 0xb5b624bf.
//
// Solidity: function getCurveInfo(uint256 curveId) view returns((uint256[],uint256))
func (_Csaccounting *CsaccountingSession) GetCurveInfo(curveId *big.Int) (ICSBondCurveBondCurve, error) {
	return _Csaccounting.Contract.GetCurveInfo(&_Csaccounting.CallOpts, curveId)
}

// GetCurveInfo is a free data retrieval call binding the contract method 0xb5b624bf.
//
// Solidity: function getCurveInfo(uint256 curveId) view returns((uint256[],uint256))
func (_Csaccounting *CsaccountingCallerSession) GetCurveInfo(curveId *big.Int) (ICSBondCurveBondCurve, error) {
	return _Csaccounting.Contract.GetCurveInfo(&_Csaccounting.CallOpts, curveId)
}

// GetKeysCountByBondAmount is a free data retrieval call binding the contract method 0xd2fa16a6.
//
// Solidity: function getKeysCountByBondAmount(uint256 amount, (uint256[],uint256) curve) pure returns(uint256)
func (_Csaccounting *CsaccountingCaller) GetKeysCountByBondAmount(opts *bind.CallOpts, amount *big.Int, curve ICSBondCurveBondCurve) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getKeysCountByBondAmount", amount, curve)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetKeysCountByBondAmount is a free data retrieval call binding the contract method 0xd2fa16a6.
//
// Solidity: function getKeysCountByBondAmount(uint256 amount, (uint256[],uint256) curve) pure returns(uint256)
func (_Csaccounting *CsaccountingSession) GetKeysCountByBondAmount(amount *big.Int, curve ICSBondCurveBondCurve) (*big.Int, error) {
	return _Csaccounting.Contract.GetKeysCountByBondAmount(&_Csaccounting.CallOpts, amount, curve)
}

// GetKeysCountByBondAmount is a free data retrieval call binding the contract method 0xd2fa16a6.
//
// Solidity: function getKeysCountByBondAmount(uint256 amount, (uint256[],uint256) curve) pure returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) GetKeysCountByBondAmount(amount *big.Int, curve ICSBondCurveBondCurve) (*big.Int, error) {
	return _Csaccounting.Contract.GetKeysCountByBondAmount(&_Csaccounting.CallOpts, amount, curve)
}

// GetKeysCountByBondAmount0 is a free data retrieval call binding the contract method 0xdc38ea3d.
//
// Solidity: function getKeysCountByBondAmount(uint256 amount, uint256 curveId) view returns(uint256)
func (_Csaccounting *CsaccountingCaller) GetKeysCountByBondAmount0(opts *bind.CallOpts, amount *big.Int, curveId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getKeysCountByBondAmount0", amount, curveId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetKeysCountByBondAmount0 is a free data retrieval call binding the contract method 0xdc38ea3d.
//
// Solidity: function getKeysCountByBondAmount(uint256 amount, uint256 curveId) view returns(uint256)
func (_Csaccounting *CsaccountingSession) GetKeysCountByBondAmount0(amount *big.Int, curveId *big.Int) (*big.Int, error) {
	return _Csaccounting.Contract.GetKeysCountByBondAmount0(&_Csaccounting.CallOpts, amount, curveId)
}

// GetKeysCountByBondAmount0 is a free data retrieval call binding the contract method 0xdc38ea3d.
//
// Solidity: function getKeysCountByBondAmount(uint256 amount, uint256 curveId) view returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) GetKeysCountByBondAmount0(amount *big.Int, curveId *big.Int) (*big.Int, error) {
	return _Csaccounting.Contract.GetKeysCountByBondAmount0(&_Csaccounting.CallOpts, amount, curveId)
}

// GetLockedBondInfo is a free data retrieval call binding the contract method 0x83316184.
//
// Solidity: function getLockedBondInfo(uint256 nodeOperatorId) view returns((uint128,uint128))
func (_Csaccounting *CsaccountingCaller) GetLockedBondInfo(opts *bind.CallOpts, nodeOperatorId *big.Int) (ICSBondLockBondLock, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getLockedBondInfo", nodeOperatorId)

	if err != nil {
		return *new(ICSBondLockBondLock), err
	}

	out0 := *abi.ConvertType(out[0], new(ICSBondLockBondLock)).(*ICSBondLockBondLock)

	return out0, err

}

// GetLockedBondInfo is a free data retrieval call binding the contract method 0x83316184.
//
// Solidity: function getLockedBondInfo(uint256 nodeOperatorId) view returns((uint128,uint128))
func (_Csaccounting *CsaccountingSession) GetLockedBondInfo(nodeOperatorId *big.Int) (ICSBondLockBondLock, error) {
	return _Csaccounting.Contract.GetLockedBondInfo(&_Csaccounting.CallOpts, nodeOperatorId)
}

// GetLockedBondInfo is a free data retrieval call binding the contract method 0x83316184.
//
// Solidity: function getLockedBondInfo(uint256 nodeOperatorId) view returns((uint128,uint128))
func (_Csaccounting *CsaccountingCallerSession) GetLockedBondInfo(nodeOperatorId *big.Int) (ICSBondLockBondLock, error) {
	return _Csaccounting.Contract.GetLockedBondInfo(&_Csaccounting.CallOpts, nodeOperatorId)
}

// GetRequiredBondForNextKeys is a free data retrieval call binding the contract method 0xb148db6a.
//
// Solidity: function getRequiredBondForNextKeys(uint256 nodeOperatorId, uint256 additionalKeys) view returns(uint256)
func (_Csaccounting *CsaccountingCaller) GetRequiredBondForNextKeys(opts *bind.CallOpts, nodeOperatorId *big.Int, additionalKeys *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getRequiredBondForNextKeys", nodeOperatorId, additionalKeys)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequiredBondForNextKeys is a free data retrieval call binding the contract method 0xb148db6a.
//
// Solidity: function getRequiredBondForNextKeys(uint256 nodeOperatorId, uint256 additionalKeys) view returns(uint256)
func (_Csaccounting *CsaccountingSession) GetRequiredBondForNextKeys(nodeOperatorId *big.Int, additionalKeys *big.Int) (*big.Int, error) {
	return _Csaccounting.Contract.GetRequiredBondForNextKeys(&_Csaccounting.CallOpts, nodeOperatorId, additionalKeys)
}

// GetRequiredBondForNextKeys is a free data retrieval call binding the contract method 0xb148db6a.
//
// Solidity: function getRequiredBondForNextKeys(uint256 nodeOperatorId, uint256 additionalKeys) view returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) GetRequiredBondForNextKeys(nodeOperatorId *big.Int, additionalKeys *big.Int) (*big.Int, error) {
	return _Csaccounting.Contract.GetRequiredBondForNextKeys(&_Csaccounting.CallOpts, nodeOperatorId, additionalKeys)
}

// GetRequiredBondForNextKeysWstETH is a free data retrieval call binding the contract method 0x28846981.
//
// Solidity: function getRequiredBondForNextKeysWstETH(uint256 nodeOperatorId, uint256 additionalKeys) view returns(uint256)
func (_Csaccounting *CsaccountingCaller) GetRequiredBondForNextKeysWstETH(opts *bind.CallOpts, nodeOperatorId *big.Int, additionalKeys *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getRequiredBondForNextKeysWstETH", nodeOperatorId, additionalKeys)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequiredBondForNextKeysWstETH is a free data retrieval call binding the contract method 0x28846981.
//
// Solidity: function getRequiredBondForNextKeysWstETH(uint256 nodeOperatorId, uint256 additionalKeys) view returns(uint256)
func (_Csaccounting *CsaccountingSession) GetRequiredBondForNextKeysWstETH(nodeOperatorId *big.Int, additionalKeys *big.Int) (*big.Int, error) {
	return _Csaccounting.Contract.GetRequiredBondForNextKeysWstETH(&_Csaccounting.CallOpts, nodeOperatorId, additionalKeys)
}

// GetRequiredBondForNextKeysWstETH is a free data retrieval call binding the contract method 0x28846981.
//
// Solidity: function getRequiredBondForNextKeysWstETH(uint256 nodeOperatorId, uint256 additionalKeys) view returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) GetRequiredBondForNextKeysWstETH(nodeOperatorId *big.Int, additionalKeys *big.Int) (*big.Int, error) {
	return _Csaccounting.Contract.GetRequiredBondForNextKeysWstETH(&_Csaccounting.CallOpts, nodeOperatorId, additionalKeys)
}

// GetResumeSinceTimestamp is a free data retrieval call binding the contract method 0x589ff76c.
//
// Solidity: function getResumeSinceTimestamp() view returns(uint256)
func (_Csaccounting *CsaccountingCaller) GetResumeSinceTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getResumeSinceTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetResumeSinceTimestamp is a free data retrieval call binding the contract method 0x589ff76c.
//
// Solidity: function getResumeSinceTimestamp() view returns(uint256)
func (_Csaccounting *CsaccountingSession) GetResumeSinceTimestamp() (*big.Int, error) {
	return _Csaccounting.Contract.GetResumeSinceTimestamp(&_Csaccounting.CallOpts)
}

// GetResumeSinceTimestamp is a free data retrieval call binding the contract method 0x589ff76c.
//
// Solidity: function getResumeSinceTimestamp() view returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) GetResumeSinceTimestamp() (*big.Int, error) {
	return _Csaccounting.Contract.GetResumeSinceTimestamp(&_Csaccounting.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Csaccounting *CsaccountingCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Csaccounting *CsaccountingSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Csaccounting.Contract.GetRoleAdmin(&_Csaccounting.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Csaccounting *CsaccountingCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Csaccounting.Contract.GetRoleAdmin(&_Csaccounting.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Csaccounting *CsaccountingCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Csaccounting *CsaccountingSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Csaccounting.Contract.GetRoleMember(&_Csaccounting.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Csaccounting *CsaccountingCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Csaccounting.Contract.GetRoleMember(&_Csaccounting.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Csaccounting *CsaccountingCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Csaccounting *CsaccountingSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Csaccounting.Contract.GetRoleMemberCount(&_Csaccounting.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Csaccounting.Contract.GetRoleMemberCount(&_Csaccounting.CallOpts, role)
}

// GetUnbondedKeysCount is a free data retrieval call binding the contract method 0x01a5e9e3.
//
// Solidity: function getUnbondedKeysCount(uint256 nodeOperatorId) view returns(uint256)
func (_Csaccounting *CsaccountingCaller) GetUnbondedKeysCount(opts *bind.CallOpts, nodeOperatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getUnbondedKeysCount", nodeOperatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnbondedKeysCount is a free data retrieval call binding the contract method 0x01a5e9e3.
//
// Solidity: function getUnbondedKeysCount(uint256 nodeOperatorId) view returns(uint256)
func (_Csaccounting *CsaccountingSession) GetUnbondedKeysCount(nodeOperatorId *big.Int) (*big.Int, error) {
	return _Csaccounting.Contract.GetUnbondedKeysCount(&_Csaccounting.CallOpts, nodeOperatorId)
}

// GetUnbondedKeysCount is a free data retrieval call binding the contract method 0x01a5e9e3.
//
// Solidity: function getUnbondedKeysCount(uint256 nodeOperatorId) view returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) GetUnbondedKeysCount(nodeOperatorId *big.Int) (*big.Int, error) {
	return _Csaccounting.Contract.GetUnbondedKeysCount(&_Csaccounting.CallOpts, nodeOperatorId)
}

// GetUnbondedKeysCountToEject is a free data retrieval call binding the contract method 0x9c516102.
//
// Solidity: function getUnbondedKeysCountToEject(uint256 nodeOperatorId) view returns(uint256)
func (_Csaccounting *CsaccountingCaller) GetUnbondedKeysCountToEject(opts *bind.CallOpts, nodeOperatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getUnbondedKeysCountToEject", nodeOperatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnbondedKeysCountToEject is a free data retrieval call binding the contract method 0x9c516102.
//
// Solidity: function getUnbondedKeysCountToEject(uint256 nodeOperatorId) view returns(uint256)
func (_Csaccounting *CsaccountingSession) GetUnbondedKeysCountToEject(nodeOperatorId *big.Int) (*big.Int, error) {
	return _Csaccounting.Contract.GetUnbondedKeysCountToEject(&_Csaccounting.CallOpts, nodeOperatorId)
}

// GetUnbondedKeysCountToEject is a free data retrieval call binding the contract method 0x9c516102.
//
// Solidity: function getUnbondedKeysCountToEject(uint256 nodeOperatorId) view returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) GetUnbondedKeysCountToEject(nodeOperatorId *big.Int) (*big.Int, error) {
	return _Csaccounting.Contract.GetUnbondedKeysCountToEject(&_Csaccounting.CallOpts, nodeOperatorId)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Csaccounting *CsaccountingCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Csaccounting *CsaccountingSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Csaccounting.Contract.HasRole(&_Csaccounting.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Csaccounting *CsaccountingCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Csaccounting.Contract.HasRole(&_Csaccounting.CallOpts, role, account)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Csaccounting *CsaccountingCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Csaccounting *CsaccountingSession) IsPaused() (bool, error) {
	return _Csaccounting.Contract.IsPaused(&_Csaccounting.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Csaccounting *CsaccountingCallerSession) IsPaused() (bool, error) {
	return _Csaccounting.Contract.IsPaused(&_Csaccounting.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Csaccounting *CsaccountingCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Csaccounting *CsaccountingSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Csaccounting.Contract.SupportsInterface(&_Csaccounting.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Csaccounting *CsaccountingCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Csaccounting.Contract.SupportsInterface(&_Csaccounting.CallOpts, interfaceId)
}

// TotalBondShares is a free data retrieval call binding the contract method 0x74d70aea.
//
// Solidity: function totalBondShares() view returns(uint256)
func (_Csaccounting *CsaccountingCaller) TotalBondShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "totalBondShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBondShares is a free data retrieval call binding the contract method 0x74d70aea.
//
// Solidity: function totalBondShares() view returns(uint256)
func (_Csaccounting *CsaccountingSession) TotalBondShares() (*big.Int, error) {
	return _Csaccounting.Contract.TotalBondShares(&_Csaccounting.CallOpts)
}

// TotalBondShares is a free data retrieval call binding the contract method 0x74d70aea.
//
// Solidity: function totalBondShares() view returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) TotalBondShares() (*big.Int, error) {
	return _Csaccounting.Contract.TotalBondShares(&_Csaccounting.CallOpts)
}

// AddBondCurve is a paid mutator transaction binding the contract method 0x573b6245.
//
// Solidity: function addBondCurve(uint256[] bondCurve) returns(uint256 id)
func (_Csaccounting *CsaccountingTransactor) AddBondCurve(opts *bind.TransactOpts, bondCurve []*big.Int) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "addBondCurve", bondCurve)
}

// AddBondCurve is a paid mutator transaction binding the contract method 0x573b6245.
//
// Solidity: function addBondCurve(uint256[] bondCurve) returns(uint256 id)
func (_Csaccounting *CsaccountingSession) AddBondCurve(bondCurve []*big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.AddBondCurve(&_Csaccounting.TransactOpts, bondCurve)
}

// AddBondCurve is a paid mutator transaction binding the contract method 0x573b6245.
//
// Solidity: function addBondCurve(uint256[] bondCurve) returns(uint256 id)
func (_Csaccounting *CsaccountingTransactorSession) AddBondCurve(bondCurve []*big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.AddBondCurve(&_Csaccounting.TransactOpts, bondCurve)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x881fa03c.
//
// Solidity: function chargeFee(uint256 nodeOperatorId, uint256 amount) returns()
func (_Csaccounting *CsaccountingTransactor) ChargeFee(opts *bind.TransactOpts, nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "chargeFee", nodeOperatorId, amount)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x881fa03c.
//
// Solidity: function chargeFee(uint256 nodeOperatorId, uint256 amount) returns()
func (_Csaccounting *CsaccountingSession) ChargeFee(nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.ChargeFee(&_Csaccounting.TransactOpts, nodeOperatorId, amount)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x881fa03c.
//
// Solidity: function chargeFee(uint256 nodeOperatorId, uint256 amount) returns()
func (_Csaccounting *CsaccountingTransactorSession) ChargeFee(nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.ChargeFee(&_Csaccounting.TransactOpts, nodeOperatorId, amount)
}

// ClaimRewardsStETH is a paid mutator transaction binding the contract method 0xf9391223.
//
// Solidity: function claimRewardsStETH(uint256 nodeOperatorId, uint256 stETHAmount, address rewardAddress, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns()
func (_Csaccounting *CsaccountingTransactor) ClaimRewardsStETH(opts *bind.TransactOpts, nodeOperatorId *big.Int, stETHAmount *big.Int, rewardAddress common.Address, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "claimRewardsStETH", nodeOperatorId, stETHAmount, rewardAddress, cumulativeFeeShares, rewardsProof)
}

// ClaimRewardsStETH is a paid mutator transaction binding the contract method 0xf9391223.
//
// Solidity: function claimRewardsStETH(uint256 nodeOperatorId, uint256 stETHAmount, address rewardAddress, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns()
func (_Csaccounting *CsaccountingSession) ClaimRewardsStETH(nodeOperatorId *big.Int, stETHAmount *big.Int, rewardAddress common.Address, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _Csaccounting.Contract.ClaimRewardsStETH(&_Csaccounting.TransactOpts, nodeOperatorId, stETHAmount, rewardAddress, cumulativeFeeShares, rewardsProof)
}

// ClaimRewardsStETH is a paid mutator transaction binding the contract method 0xf9391223.
//
// Solidity: function claimRewardsStETH(uint256 nodeOperatorId, uint256 stETHAmount, address rewardAddress, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns()
func (_Csaccounting *CsaccountingTransactorSession) ClaimRewardsStETH(nodeOperatorId *big.Int, stETHAmount *big.Int, rewardAddress common.Address, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _Csaccounting.Contract.ClaimRewardsStETH(&_Csaccounting.TransactOpts, nodeOperatorId, stETHAmount, rewardAddress, cumulativeFeeShares, rewardsProof)
}

// ClaimRewardsUnstETH is a paid mutator transaction binding the contract method 0xcc810cb9.
//
// Solidity: function claimRewardsUnstETH(uint256 nodeOperatorId, uint256 stEthAmount, address rewardAddress, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns()
func (_Csaccounting *CsaccountingTransactor) ClaimRewardsUnstETH(opts *bind.TransactOpts, nodeOperatorId *big.Int, stEthAmount *big.Int, rewardAddress common.Address, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "claimRewardsUnstETH", nodeOperatorId, stEthAmount, rewardAddress, cumulativeFeeShares, rewardsProof)
}

// ClaimRewardsUnstETH is a paid mutator transaction binding the contract method 0xcc810cb9.
//
// Solidity: function claimRewardsUnstETH(uint256 nodeOperatorId, uint256 stEthAmount, address rewardAddress, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns()
func (_Csaccounting *CsaccountingSession) ClaimRewardsUnstETH(nodeOperatorId *big.Int, stEthAmount *big.Int, rewardAddress common.Address, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _Csaccounting.Contract.ClaimRewardsUnstETH(&_Csaccounting.TransactOpts, nodeOperatorId, stEthAmount, rewardAddress, cumulativeFeeShares, rewardsProof)
}

// ClaimRewardsUnstETH is a paid mutator transaction binding the contract method 0xcc810cb9.
//
// Solidity: function claimRewardsUnstETH(uint256 nodeOperatorId, uint256 stEthAmount, address rewardAddress, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns()
func (_Csaccounting *CsaccountingTransactorSession) ClaimRewardsUnstETH(nodeOperatorId *big.Int, stEthAmount *big.Int, rewardAddress common.Address, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _Csaccounting.Contract.ClaimRewardsUnstETH(&_Csaccounting.TransactOpts, nodeOperatorId, stEthAmount, rewardAddress, cumulativeFeeShares, rewardsProof)
}

// ClaimRewardsWstETH is a paid mutator transaction binding the contract method 0x70903eb9.
//
// Solidity: function claimRewardsWstETH(uint256 nodeOperatorId, uint256 wstETHAmount, address rewardAddress, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns()
func (_Csaccounting *CsaccountingTransactor) ClaimRewardsWstETH(opts *bind.TransactOpts, nodeOperatorId *big.Int, wstETHAmount *big.Int, rewardAddress common.Address, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "claimRewardsWstETH", nodeOperatorId, wstETHAmount, rewardAddress, cumulativeFeeShares, rewardsProof)
}

// ClaimRewardsWstETH is a paid mutator transaction binding the contract method 0x70903eb9.
//
// Solidity: function claimRewardsWstETH(uint256 nodeOperatorId, uint256 wstETHAmount, address rewardAddress, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns()
func (_Csaccounting *CsaccountingSession) ClaimRewardsWstETH(nodeOperatorId *big.Int, wstETHAmount *big.Int, rewardAddress common.Address, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _Csaccounting.Contract.ClaimRewardsWstETH(&_Csaccounting.TransactOpts, nodeOperatorId, wstETHAmount, rewardAddress, cumulativeFeeShares, rewardsProof)
}

// ClaimRewardsWstETH is a paid mutator transaction binding the contract method 0x70903eb9.
//
// Solidity: function claimRewardsWstETH(uint256 nodeOperatorId, uint256 wstETHAmount, address rewardAddress, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns()
func (_Csaccounting *CsaccountingTransactorSession) ClaimRewardsWstETH(nodeOperatorId *big.Int, wstETHAmount *big.Int, rewardAddress common.Address, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _Csaccounting.Contract.ClaimRewardsWstETH(&_Csaccounting.TransactOpts, nodeOperatorId, wstETHAmount, rewardAddress, cumulativeFeeShares, rewardsProof)
}

// CompensateLockedBondETH is a paid mutator transaction binding the contract method 0x15b5c477.
//
// Solidity: function compensateLockedBondETH(uint256 nodeOperatorId) payable returns()
func (_Csaccounting *CsaccountingTransactor) CompensateLockedBondETH(opts *bind.TransactOpts, nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "compensateLockedBondETH", nodeOperatorId)
}

// CompensateLockedBondETH is a paid mutator transaction binding the contract method 0x15b5c477.
//
// Solidity: function compensateLockedBondETH(uint256 nodeOperatorId) payable returns()
func (_Csaccounting *CsaccountingSession) CompensateLockedBondETH(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.CompensateLockedBondETH(&_Csaccounting.TransactOpts, nodeOperatorId)
}

// CompensateLockedBondETH is a paid mutator transaction binding the contract method 0x15b5c477.
//
// Solidity: function compensateLockedBondETH(uint256 nodeOperatorId) payable returns()
func (_Csaccounting *CsaccountingTransactorSession) CompensateLockedBondETH(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.CompensateLockedBondETH(&_Csaccounting.TransactOpts, nodeOperatorId)
}

// DepositETH is a paid mutator transaction binding the contract method 0x2e599054.
//
// Solidity: function depositETH(address from, uint256 nodeOperatorId) payable returns()
func (_Csaccounting *CsaccountingTransactor) DepositETH(opts *bind.TransactOpts, from common.Address, nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "depositETH", from, nodeOperatorId)
}

// DepositETH is a paid mutator transaction binding the contract method 0x2e599054.
//
// Solidity: function depositETH(address from, uint256 nodeOperatorId) payable returns()
func (_Csaccounting *CsaccountingSession) DepositETH(from common.Address, nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.DepositETH(&_Csaccounting.TransactOpts, from, nodeOperatorId)
}

// DepositETH is a paid mutator transaction binding the contract method 0x2e599054.
//
// Solidity: function depositETH(address from, uint256 nodeOperatorId) payable returns()
func (_Csaccounting *CsaccountingTransactorSession) DepositETH(from common.Address, nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.DepositETH(&_Csaccounting.TransactOpts, from, nodeOperatorId)
}

// DepositStETH is a paid mutator transaction binding the contract method 0x4c7ed3d2.
//
// Solidity: function depositStETH(address from, uint256 nodeOperatorId, uint256 stETHAmount, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_Csaccounting *CsaccountingTransactor) DepositStETH(opts *bind.TransactOpts, from common.Address, nodeOperatorId *big.Int, stETHAmount *big.Int, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "depositStETH", from, nodeOperatorId, stETHAmount, permit)
}

// DepositStETH is a paid mutator transaction binding the contract method 0x4c7ed3d2.
//
// Solidity: function depositStETH(address from, uint256 nodeOperatorId, uint256 stETHAmount, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_Csaccounting *CsaccountingSession) DepositStETH(from common.Address, nodeOperatorId *big.Int, stETHAmount *big.Int, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _Csaccounting.Contract.DepositStETH(&_Csaccounting.TransactOpts, from, nodeOperatorId, stETHAmount, permit)
}

// DepositStETH is a paid mutator transaction binding the contract method 0x4c7ed3d2.
//
// Solidity: function depositStETH(address from, uint256 nodeOperatorId, uint256 stETHAmount, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_Csaccounting *CsaccountingTransactorSession) DepositStETH(from common.Address, nodeOperatorId *big.Int, stETHAmount *big.Int, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _Csaccounting.Contract.DepositStETH(&_Csaccounting.TransactOpts, from, nodeOperatorId, stETHAmount, permit)
}

// DepositWstETH is a paid mutator transaction binding the contract method 0xf7966efe.
//
// Solidity: function depositWstETH(address from, uint256 nodeOperatorId, uint256 wstETHAmount, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_Csaccounting *CsaccountingTransactor) DepositWstETH(opts *bind.TransactOpts, from common.Address, nodeOperatorId *big.Int, wstETHAmount *big.Int, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "depositWstETH", from, nodeOperatorId, wstETHAmount, permit)
}

// DepositWstETH is a paid mutator transaction binding the contract method 0xf7966efe.
//
// Solidity: function depositWstETH(address from, uint256 nodeOperatorId, uint256 wstETHAmount, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_Csaccounting *CsaccountingSession) DepositWstETH(from common.Address, nodeOperatorId *big.Int, wstETHAmount *big.Int, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _Csaccounting.Contract.DepositWstETH(&_Csaccounting.TransactOpts, from, nodeOperatorId, wstETHAmount, permit)
}

// DepositWstETH is a paid mutator transaction binding the contract method 0xf7966efe.
//
// Solidity: function depositWstETH(address from, uint256 nodeOperatorId, uint256 wstETHAmount, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_Csaccounting *CsaccountingTransactorSession) DepositWstETH(from common.Address, nodeOperatorId *big.Int, wstETHAmount *big.Int, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _Csaccounting.Contract.DepositWstETH(&_Csaccounting.TransactOpts, from, nodeOperatorId, wstETHAmount, permit)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Csaccounting *CsaccountingTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Csaccounting *CsaccountingSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Csaccounting.Contract.GrantRole(&_Csaccounting.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Csaccounting *CsaccountingTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Csaccounting.Contract.GrantRole(&_Csaccounting.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xfab382f1.
//
// Solidity: function initialize(uint256[] bondCurve, address admin, address _feeDistributor, uint256 bondLockRetentionPeriod, address _chargePenaltyRecipient) returns()
func (_Csaccounting *CsaccountingTransactor) Initialize(opts *bind.TransactOpts, bondCurve []*big.Int, admin common.Address, _feeDistributor common.Address, bondLockRetentionPeriod *big.Int, _chargePenaltyRecipient common.Address) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "initialize", bondCurve, admin, _feeDistributor, bondLockRetentionPeriod, _chargePenaltyRecipient)
}

// Initialize is a paid mutator transaction binding the contract method 0xfab382f1.
//
// Solidity: function initialize(uint256[] bondCurve, address admin, address _feeDistributor, uint256 bondLockRetentionPeriod, address _chargePenaltyRecipient) returns()
func (_Csaccounting *CsaccountingSession) Initialize(bondCurve []*big.Int, admin common.Address, _feeDistributor common.Address, bondLockRetentionPeriod *big.Int, _chargePenaltyRecipient common.Address) (*types.Transaction, error) {
	return _Csaccounting.Contract.Initialize(&_Csaccounting.TransactOpts, bondCurve, admin, _feeDistributor, bondLockRetentionPeriod, _chargePenaltyRecipient)
}

// Initialize is a paid mutator transaction binding the contract method 0xfab382f1.
//
// Solidity: function initialize(uint256[] bondCurve, address admin, address _feeDistributor, uint256 bondLockRetentionPeriod, address _chargePenaltyRecipient) returns()
func (_Csaccounting *CsaccountingTransactorSession) Initialize(bondCurve []*big.Int, admin common.Address, _feeDistributor common.Address, bondLockRetentionPeriod *big.Int, _chargePenaltyRecipient common.Address) (*types.Transaction, error) {
	return _Csaccounting.Contract.Initialize(&_Csaccounting.TransactOpts, bondCurve, admin, _feeDistributor, bondLockRetentionPeriod, _chargePenaltyRecipient)
}

// LockBondETH is a paid mutator transaction binding the contract method 0xdcab7f83.
//
// Solidity: function lockBondETH(uint256 nodeOperatorId, uint256 amount) returns()
func (_Csaccounting *CsaccountingTransactor) LockBondETH(opts *bind.TransactOpts, nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "lockBondETH", nodeOperatorId, amount)
}

// LockBondETH is a paid mutator transaction binding the contract method 0xdcab7f83.
//
// Solidity: function lockBondETH(uint256 nodeOperatorId, uint256 amount) returns()
func (_Csaccounting *CsaccountingSession) LockBondETH(nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.LockBondETH(&_Csaccounting.TransactOpts, nodeOperatorId, amount)
}

// LockBondETH is a paid mutator transaction binding the contract method 0xdcab7f83.
//
// Solidity: function lockBondETH(uint256 nodeOperatorId, uint256 amount) returns()
func (_Csaccounting *CsaccountingTransactorSession) LockBondETH(nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.LockBondETH(&_Csaccounting.TransactOpts, nodeOperatorId, amount)
}

// PauseFor is a paid mutator transaction binding the contract method 0xf3f449c7.
//
// Solidity: function pauseFor(uint256 duration) returns()
func (_Csaccounting *CsaccountingTransactor) PauseFor(opts *bind.TransactOpts, duration *big.Int) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "pauseFor", duration)
}

// PauseFor is a paid mutator transaction binding the contract method 0xf3f449c7.
//
// Solidity: function pauseFor(uint256 duration) returns()
func (_Csaccounting *CsaccountingSession) PauseFor(duration *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.PauseFor(&_Csaccounting.TransactOpts, duration)
}

// PauseFor is a paid mutator transaction binding the contract method 0xf3f449c7.
//
// Solidity: function pauseFor(uint256 duration) returns()
func (_Csaccounting *CsaccountingTransactorSession) PauseFor(duration *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.PauseFor(&_Csaccounting.TransactOpts, duration)
}

// Penalize is a paid mutator transaction binding the contract method 0xe5220e3f.
//
// Solidity: function penalize(uint256 nodeOperatorId, uint256 amount) returns()
func (_Csaccounting *CsaccountingTransactor) Penalize(opts *bind.TransactOpts, nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "penalize", nodeOperatorId, amount)
}

// Penalize is a paid mutator transaction binding the contract method 0xe5220e3f.
//
// Solidity: function penalize(uint256 nodeOperatorId, uint256 amount) returns()
func (_Csaccounting *CsaccountingSession) Penalize(nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.Penalize(&_Csaccounting.TransactOpts, nodeOperatorId, amount)
}

// Penalize is a paid mutator transaction binding the contract method 0xe5220e3f.
//
// Solidity: function penalize(uint256 nodeOperatorId, uint256 amount) returns()
func (_Csaccounting *CsaccountingTransactorSession) Penalize(nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.Penalize(&_Csaccounting.TransactOpts, nodeOperatorId, amount)
}

// PullFeeRewards is a paid mutator transaction binding the contract method 0x9b4c6c27.
//
// Solidity: function pullFeeRewards(uint256 nodeOperatorId, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns()
func (_Csaccounting *CsaccountingTransactor) PullFeeRewards(opts *bind.TransactOpts, nodeOperatorId *big.Int, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "pullFeeRewards", nodeOperatorId, cumulativeFeeShares, rewardsProof)
}

// PullFeeRewards is a paid mutator transaction binding the contract method 0x9b4c6c27.
//
// Solidity: function pullFeeRewards(uint256 nodeOperatorId, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns()
func (_Csaccounting *CsaccountingSession) PullFeeRewards(nodeOperatorId *big.Int, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _Csaccounting.Contract.PullFeeRewards(&_Csaccounting.TransactOpts, nodeOperatorId, cumulativeFeeShares, rewardsProof)
}

// PullFeeRewards is a paid mutator transaction binding the contract method 0x9b4c6c27.
//
// Solidity: function pullFeeRewards(uint256 nodeOperatorId, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns()
func (_Csaccounting *CsaccountingTransactorSession) PullFeeRewards(nodeOperatorId *big.Int, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _Csaccounting.Contract.PullFeeRewards(&_Csaccounting.TransactOpts, nodeOperatorId, cumulativeFeeShares, rewardsProof)
}

// RecoverERC1155 is a paid mutator transaction binding the contract method 0x5c654ad9.
//
// Solidity: function recoverERC1155(address token, uint256 tokenId) returns()
func (_Csaccounting *CsaccountingTransactor) RecoverERC1155(opts *bind.TransactOpts, token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "recoverERC1155", token, tokenId)
}

// RecoverERC1155 is a paid mutator transaction binding the contract method 0x5c654ad9.
//
// Solidity: function recoverERC1155(address token, uint256 tokenId) returns()
func (_Csaccounting *CsaccountingSession) RecoverERC1155(token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.RecoverERC1155(&_Csaccounting.TransactOpts, token, tokenId)
}

// RecoverERC1155 is a paid mutator transaction binding the contract method 0x5c654ad9.
//
// Solidity: function recoverERC1155(address token, uint256 tokenId) returns()
func (_Csaccounting *CsaccountingTransactorSession) RecoverERC1155(token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.RecoverERC1155(&_Csaccounting.TransactOpts, token, tokenId)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address token, uint256 amount) returns()
func (_Csaccounting *CsaccountingTransactor) RecoverERC20(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "recoverERC20", token, amount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address token, uint256 amount) returns()
func (_Csaccounting *CsaccountingSession) RecoverERC20(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.RecoverERC20(&_Csaccounting.TransactOpts, token, amount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address token, uint256 amount) returns()
func (_Csaccounting *CsaccountingTransactorSession) RecoverERC20(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.RecoverERC20(&_Csaccounting.TransactOpts, token, amount)
}

// RecoverERC721 is a paid mutator transaction binding the contract method 0x819d4cc6.
//
// Solidity: function recoverERC721(address token, uint256 tokenId) returns()
func (_Csaccounting *CsaccountingTransactor) RecoverERC721(opts *bind.TransactOpts, token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "recoverERC721", token, tokenId)
}

// RecoverERC721 is a paid mutator transaction binding the contract method 0x819d4cc6.
//
// Solidity: function recoverERC721(address token, uint256 tokenId) returns()
func (_Csaccounting *CsaccountingSession) RecoverERC721(token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.RecoverERC721(&_Csaccounting.TransactOpts, token, tokenId)
}

// RecoverERC721 is a paid mutator transaction binding the contract method 0x819d4cc6.
//
// Solidity: function recoverERC721(address token, uint256 tokenId) returns()
func (_Csaccounting *CsaccountingTransactorSession) RecoverERC721(token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.RecoverERC721(&_Csaccounting.TransactOpts, token, tokenId)
}

// RecoverEther is a paid mutator transaction binding the contract method 0x52d8bfc2.
//
// Solidity: function recoverEther() returns()
func (_Csaccounting *CsaccountingTransactor) RecoverEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "recoverEther")
}

// RecoverEther is a paid mutator transaction binding the contract method 0x52d8bfc2.
//
// Solidity: function recoverEther() returns()
func (_Csaccounting *CsaccountingSession) RecoverEther() (*types.Transaction, error) {
	return _Csaccounting.Contract.RecoverEther(&_Csaccounting.TransactOpts)
}

// RecoverEther is a paid mutator transaction binding the contract method 0x52d8bfc2.
//
// Solidity: function recoverEther() returns()
func (_Csaccounting *CsaccountingTransactorSession) RecoverEther() (*types.Transaction, error) {
	return _Csaccounting.Contract.RecoverEther(&_Csaccounting.TransactOpts)
}

// RecoverStETHShares is a paid mutator transaction binding the contract method 0x5a73bdc8.
//
// Solidity: function recoverStETHShares() returns()
func (_Csaccounting *CsaccountingTransactor) RecoverStETHShares(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "recoverStETHShares")
}

// RecoverStETHShares is a paid mutator transaction binding the contract method 0x5a73bdc8.
//
// Solidity: function recoverStETHShares() returns()
func (_Csaccounting *CsaccountingSession) RecoverStETHShares() (*types.Transaction, error) {
	return _Csaccounting.Contract.RecoverStETHShares(&_Csaccounting.TransactOpts)
}

// RecoverStETHShares is a paid mutator transaction binding the contract method 0x5a73bdc8.
//
// Solidity: function recoverStETHShares() returns()
func (_Csaccounting *CsaccountingTransactorSession) RecoverStETHShares() (*types.Transaction, error) {
	return _Csaccounting.Contract.RecoverStETHShares(&_Csaccounting.TransactOpts)
}

// ReleaseLockedBondETH is a paid mutator transaction binding the contract method 0xd963ae55.
//
// Solidity: function releaseLockedBondETH(uint256 nodeOperatorId, uint256 amount) returns()
func (_Csaccounting *CsaccountingTransactor) ReleaseLockedBondETH(opts *bind.TransactOpts, nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "releaseLockedBondETH", nodeOperatorId, amount)
}

// ReleaseLockedBondETH is a paid mutator transaction binding the contract method 0xd963ae55.
//
// Solidity: function releaseLockedBondETH(uint256 nodeOperatorId, uint256 amount) returns()
func (_Csaccounting *CsaccountingSession) ReleaseLockedBondETH(nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.ReleaseLockedBondETH(&_Csaccounting.TransactOpts, nodeOperatorId, amount)
}

// ReleaseLockedBondETH is a paid mutator transaction binding the contract method 0xd963ae55.
//
// Solidity: function releaseLockedBondETH(uint256 nodeOperatorId, uint256 amount) returns()
func (_Csaccounting *CsaccountingTransactorSession) ReleaseLockedBondETH(nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.ReleaseLockedBondETH(&_Csaccounting.TransactOpts, nodeOperatorId, amount)
}

// RenewBurnerAllowance is a paid mutator transaction binding the contract method 0xf3efecc4.
//
// Solidity: function renewBurnerAllowance() returns()
func (_Csaccounting *CsaccountingTransactor) RenewBurnerAllowance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "renewBurnerAllowance")
}

// RenewBurnerAllowance is a paid mutator transaction binding the contract method 0xf3efecc4.
//
// Solidity: function renewBurnerAllowance() returns()
func (_Csaccounting *CsaccountingSession) RenewBurnerAllowance() (*types.Transaction, error) {
	return _Csaccounting.Contract.RenewBurnerAllowance(&_Csaccounting.TransactOpts)
}

// RenewBurnerAllowance is a paid mutator transaction binding the contract method 0xf3efecc4.
//
// Solidity: function renewBurnerAllowance() returns()
func (_Csaccounting *CsaccountingTransactorSession) RenewBurnerAllowance() (*types.Transaction, error) {
	return _Csaccounting.Contract.RenewBurnerAllowance(&_Csaccounting.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Csaccounting *CsaccountingTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Csaccounting *CsaccountingSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Csaccounting.Contract.RenounceRole(&_Csaccounting.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Csaccounting *CsaccountingTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Csaccounting.Contract.RenounceRole(&_Csaccounting.TransactOpts, role, callerConfirmation)
}

// ResetBondCurve is a paid mutator transaction binding the contract method 0x449add1b.
//
// Solidity: function resetBondCurve(uint256 nodeOperatorId) returns()
func (_Csaccounting *CsaccountingTransactor) ResetBondCurve(opts *bind.TransactOpts, nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "resetBondCurve", nodeOperatorId)
}

// ResetBondCurve is a paid mutator transaction binding the contract method 0x449add1b.
//
// Solidity: function resetBondCurve(uint256 nodeOperatorId) returns()
func (_Csaccounting *CsaccountingSession) ResetBondCurve(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.ResetBondCurve(&_Csaccounting.TransactOpts, nodeOperatorId)
}

// ResetBondCurve is a paid mutator transaction binding the contract method 0x449add1b.
//
// Solidity: function resetBondCurve(uint256 nodeOperatorId) returns()
func (_Csaccounting *CsaccountingTransactorSession) ResetBondCurve(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.ResetBondCurve(&_Csaccounting.TransactOpts, nodeOperatorId)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Csaccounting *CsaccountingTransactor) Resume(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "resume")
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Csaccounting *CsaccountingSession) Resume() (*types.Transaction, error) {
	return _Csaccounting.Contract.Resume(&_Csaccounting.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Csaccounting *CsaccountingTransactorSession) Resume() (*types.Transaction, error) {
	return _Csaccounting.Contract.Resume(&_Csaccounting.TransactOpts)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Csaccounting *CsaccountingTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Csaccounting *CsaccountingSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Csaccounting.Contract.RevokeRole(&_Csaccounting.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Csaccounting *CsaccountingTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Csaccounting.Contract.RevokeRole(&_Csaccounting.TransactOpts, role, account)
}

// SetBondCurve is a paid mutator transaction binding the contract method 0xb2d03e4d.
//
// Solidity: function setBondCurve(uint256 nodeOperatorId, uint256 curveId) returns()
func (_Csaccounting *CsaccountingTransactor) SetBondCurve(opts *bind.TransactOpts, nodeOperatorId *big.Int, curveId *big.Int) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "setBondCurve", nodeOperatorId, curveId)
}

// SetBondCurve is a paid mutator transaction binding the contract method 0xb2d03e4d.
//
// Solidity: function setBondCurve(uint256 nodeOperatorId, uint256 curveId) returns()
func (_Csaccounting *CsaccountingSession) SetBondCurve(nodeOperatorId *big.Int, curveId *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.SetBondCurve(&_Csaccounting.TransactOpts, nodeOperatorId, curveId)
}

// SetBondCurve is a paid mutator transaction binding the contract method 0xb2d03e4d.
//
// Solidity: function setBondCurve(uint256 nodeOperatorId, uint256 curveId) returns()
func (_Csaccounting *CsaccountingTransactorSession) SetBondCurve(nodeOperatorId *big.Int, curveId *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.SetBondCurve(&_Csaccounting.TransactOpts, nodeOperatorId, curveId)
}

// SetChargePenaltyRecipient is a paid mutator transaction binding the contract method 0x433cd6c3.
//
// Solidity: function setChargePenaltyRecipient(address _chargePenaltyRecipient) returns()
func (_Csaccounting *CsaccountingTransactor) SetChargePenaltyRecipient(opts *bind.TransactOpts, _chargePenaltyRecipient common.Address) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "setChargePenaltyRecipient", _chargePenaltyRecipient)
}

// SetChargePenaltyRecipient is a paid mutator transaction binding the contract method 0x433cd6c3.
//
// Solidity: function setChargePenaltyRecipient(address _chargePenaltyRecipient) returns()
func (_Csaccounting *CsaccountingSession) SetChargePenaltyRecipient(_chargePenaltyRecipient common.Address) (*types.Transaction, error) {
	return _Csaccounting.Contract.SetChargePenaltyRecipient(&_Csaccounting.TransactOpts, _chargePenaltyRecipient)
}

// SetChargePenaltyRecipient is a paid mutator transaction binding the contract method 0x433cd6c3.
//
// Solidity: function setChargePenaltyRecipient(address _chargePenaltyRecipient) returns()
func (_Csaccounting *CsaccountingTransactorSession) SetChargePenaltyRecipient(_chargePenaltyRecipient common.Address) (*types.Transaction, error) {
	return _Csaccounting.Contract.SetChargePenaltyRecipient(&_Csaccounting.TransactOpts, _chargePenaltyRecipient)
}

// SetLockedBondRetentionPeriod is a paid mutator transaction binding the contract method 0x99965225.
//
// Solidity: function setLockedBondRetentionPeriod(uint256 retention) returns()
func (_Csaccounting *CsaccountingTransactor) SetLockedBondRetentionPeriod(opts *bind.TransactOpts, retention *big.Int) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "setLockedBondRetentionPeriod", retention)
}

// SetLockedBondRetentionPeriod is a paid mutator transaction binding the contract method 0x99965225.
//
// Solidity: function setLockedBondRetentionPeriod(uint256 retention) returns()
func (_Csaccounting *CsaccountingSession) SetLockedBondRetentionPeriod(retention *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.SetLockedBondRetentionPeriod(&_Csaccounting.TransactOpts, retention)
}

// SetLockedBondRetentionPeriod is a paid mutator transaction binding the contract method 0x99965225.
//
// Solidity: function setLockedBondRetentionPeriod(uint256 retention) returns()
func (_Csaccounting *CsaccountingTransactorSession) SetLockedBondRetentionPeriod(retention *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.SetLockedBondRetentionPeriod(&_Csaccounting.TransactOpts, retention)
}

// SettleLockedBondETH is a paid mutator transaction binding the contract method 0x4bb22a72.
//
// Solidity: function settleLockedBondETH(uint256 nodeOperatorId) returns()
func (_Csaccounting *CsaccountingTransactor) SettleLockedBondETH(opts *bind.TransactOpts, nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "settleLockedBondETH", nodeOperatorId)
}

// SettleLockedBondETH is a paid mutator transaction binding the contract method 0x4bb22a72.
//
// Solidity: function settleLockedBondETH(uint256 nodeOperatorId) returns()
func (_Csaccounting *CsaccountingSession) SettleLockedBondETH(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.SettleLockedBondETH(&_Csaccounting.TransactOpts, nodeOperatorId)
}

// SettleLockedBondETH is a paid mutator transaction binding the contract method 0x4bb22a72.
//
// Solidity: function settleLockedBondETH(uint256 nodeOperatorId) returns()
func (_Csaccounting *CsaccountingTransactorSession) SettleLockedBondETH(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.SettleLockedBondETH(&_Csaccounting.TransactOpts, nodeOperatorId)
}

// UpdateBondCurve is a paid mutator transaction binding the contract method 0x019c1a4f.
//
// Solidity: function updateBondCurve(uint256 curveId, uint256[] bondCurve) returns()
func (_Csaccounting *CsaccountingTransactor) UpdateBondCurve(opts *bind.TransactOpts, curveId *big.Int, bondCurve []*big.Int) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "updateBondCurve", curveId, bondCurve)
}

// UpdateBondCurve is a paid mutator transaction binding the contract method 0x019c1a4f.
//
// Solidity: function updateBondCurve(uint256 curveId, uint256[] bondCurve) returns()
func (_Csaccounting *CsaccountingSession) UpdateBondCurve(curveId *big.Int, bondCurve []*big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.UpdateBondCurve(&_Csaccounting.TransactOpts, curveId, bondCurve)
}

// UpdateBondCurve is a paid mutator transaction binding the contract method 0x019c1a4f.
//
// Solidity: function updateBondCurve(uint256 curveId, uint256[] bondCurve) returns()
func (_Csaccounting *CsaccountingTransactorSession) UpdateBondCurve(curveId *big.Int, bondCurve []*big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.UpdateBondCurve(&_Csaccounting.TransactOpts, curveId, bondCurve)
}

// CsaccountingBondBurnedIterator is returned from FilterBondBurned and is used to iterate over the raw logs and unpacked data for BondBurned events raised by the Csaccounting contract.
type CsaccountingBondBurnedIterator struct {
	Event *CsaccountingBondBurned // Event containing the contract specifics and raw log

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
func (it *CsaccountingBondBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingBondBurned)
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
		it.Event = new(CsaccountingBondBurned)
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
func (it *CsaccountingBondBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingBondBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingBondBurned represents a BondBurned event raised by the Csaccounting contract.
type CsaccountingBondBurned struct {
	NodeOperatorId *big.Int
	ToBurnAmount   *big.Int
	BurnedAmount   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBondBurned is a free log retrieval operation binding the contract event 0x4da924ae7845fe96897faab524b536685b8bbc4d82fbb45c10d941e0f86ade0f.
//
// Solidity: event BondBurned(uint256 indexed nodeOperatorId, uint256 toBurnAmount, uint256 burnedAmount)
func (_Csaccounting *CsaccountingFilterer) FilterBondBurned(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsaccountingBondBurnedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "BondBurned", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsaccountingBondBurnedIterator{contract: _Csaccounting.contract, event: "BondBurned", logs: logs, sub: sub}, nil
}

// WatchBondBurned is a free log subscription operation binding the contract event 0x4da924ae7845fe96897faab524b536685b8bbc4d82fbb45c10d941e0f86ade0f.
//
// Solidity: event BondBurned(uint256 indexed nodeOperatorId, uint256 toBurnAmount, uint256 burnedAmount)
func (_Csaccounting *CsaccountingFilterer) WatchBondBurned(opts *bind.WatchOpts, sink chan<- *CsaccountingBondBurned, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "BondBurned", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingBondBurned)
				if err := _Csaccounting.contract.UnpackLog(event, "BondBurned", log); err != nil {
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

// ParseBondBurned is a log parse operation binding the contract event 0x4da924ae7845fe96897faab524b536685b8bbc4d82fbb45c10d941e0f86ade0f.
//
// Solidity: event BondBurned(uint256 indexed nodeOperatorId, uint256 toBurnAmount, uint256 burnedAmount)
func (_Csaccounting *CsaccountingFilterer) ParseBondBurned(log types.Log) (*CsaccountingBondBurned, error) {
	event := new(CsaccountingBondBurned)
	if err := _Csaccounting.contract.UnpackLog(event, "BondBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingBondChargedIterator is returned from FilterBondCharged and is used to iterate over the raw logs and unpacked data for BondCharged events raised by the Csaccounting contract.
type CsaccountingBondChargedIterator struct {
	Event *CsaccountingBondCharged // Event containing the contract specifics and raw log

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
func (it *CsaccountingBondChargedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingBondCharged)
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
		it.Event = new(CsaccountingBondCharged)
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
func (it *CsaccountingBondChargedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingBondChargedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingBondCharged represents a BondCharged event raised by the Csaccounting contract.
type CsaccountingBondCharged struct {
	NodeOperatorId *big.Int
	ToChargeAmount *big.Int
	ChargedAmount  *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBondCharged is a free log retrieval operation binding the contract event 0x8615528474a7bb3a28d9971535d956b79242b8e8fcfb27f3e331270fff088afd.
//
// Solidity: event BondCharged(uint256 indexed nodeOperatorId, uint256 toChargeAmount, uint256 chargedAmount)
func (_Csaccounting *CsaccountingFilterer) FilterBondCharged(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsaccountingBondChargedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "BondCharged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsaccountingBondChargedIterator{contract: _Csaccounting.contract, event: "BondCharged", logs: logs, sub: sub}, nil
}

// WatchBondCharged is a free log subscription operation binding the contract event 0x8615528474a7bb3a28d9971535d956b79242b8e8fcfb27f3e331270fff088afd.
//
// Solidity: event BondCharged(uint256 indexed nodeOperatorId, uint256 toChargeAmount, uint256 chargedAmount)
func (_Csaccounting *CsaccountingFilterer) WatchBondCharged(opts *bind.WatchOpts, sink chan<- *CsaccountingBondCharged, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "BondCharged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingBondCharged)
				if err := _Csaccounting.contract.UnpackLog(event, "BondCharged", log); err != nil {
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

// ParseBondCharged is a log parse operation binding the contract event 0x8615528474a7bb3a28d9971535d956b79242b8e8fcfb27f3e331270fff088afd.
//
// Solidity: event BondCharged(uint256 indexed nodeOperatorId, uint256 toChargeAmount, uint256 chargedAmount)
func (_Csaccounting *CsaccountingFilterer) ParseBondCharged(log types.Log) (*CsaccountingBondCharged, error) {
	event := new(CsaccountingBondCharged)
	if err := _Csaccounting.contract.UnpackLog(event, "BondCharged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingBondClaimedStETHIterator is returned from FilterBondClaimedStETH and is used to iterate over the raw logs and unpacked data for BondClaimedStETH events raised by the Csaccounting contract.
type CsaccountingBondClaimedStETHIterator struct {
	Event *CsaccountingBondClaimedStETH // Event containing the contract specifics and raw log

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
func (it *CsaccountingBondClaimedStETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingBondClaimedStETH)
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
		it.Event = new(CsaccountingBondClaimedStETH)
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
func (it *CsaccountingBondClaimedStETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingBondClaimedStETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingBondClaimedStETH represents a BondClaimedStETH event raised by the Csaccounting contract.
type CsaccountingBondClaimedStETH struct {
	NodeOperatorId *big.Int
	To             common.Address
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBondClaimedStETH is a free log retrieval operation binding the contract event 0x3e3a1398fe71575ed0c17a80cd9d46ad684c2c75c2fad7b0e7dde15e78ab22d3.
//
// Solidity: event BondClaimedStETH(uint256 indexed nodeOperatorId, address to, uint256 amount)
func (_Csaccounting *CsaccountingFilterer) FilterBondClaimedStETH(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsaccountingBondClaimedStETHIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "BondClaimedStETH", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsaccountingBondClaimedStETHIterator{contract: _Csaccounting.contract, event: "BondClaimedStETH", logs: logs, sub: sub}, nil
}

// WatchBondClaimedStETH is a free log subscription operation binding the contract event 0x3e3a1398fe71575ed0c17a80cd9d46ad684c2c75c2fad7b0e7dde15e78ab22d3.
//
// Solidity: event BondClaimedStETH(uint256 indexed nodeOperatorId, address to, uint256 amount)
func (_Csaccounting *CsaccountingFilterer) WatchBondClaimedStETH(opts *bind.WatchOpts, sink chan<- *CsaccountingBondClaimedStETH, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "BondClaimedStETH", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingBondClaimedStETH)
				if err := _Csaccounting.contract.UnpackLog(event, "BondClaimedStETH", log); err != nil {
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

// ParseBondClaimedStETH is a log parse operation binding the contract event 0x3e3a1398fe71575ed0c17a80cd9d46ad684c2c75c2fad7b0e7dde15e78ab22d3.
//
// Solidity: event BondClaimedStETH(uint256 indexed nodeOperatorId, address to, uint256 amount)
func (_Csaccounting *CsaccountingFilterer) ParseBondClaimedStETH(log types.Log) (*CsaccountingBondClaimedStETH, error) {
	event := new(CsaccountingBondClaimedStETH)
	if err := _Csaccounting.contract.UnpackLog(event, "BondClaimedStETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingBondClaimedUnstETHIterator is returned from FilterBondClaimedUnstETH and is used to iterate over the raw logs and unpacked data for BondClaimedUnstETH events raised by the Csaccounting contract.
type CsaccountingBondClaimedUnstETHIterator struct {
	Event *CsaccountingBondClaimedUnstETH // Event containing the contract specifics and raw log

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
func (it *CsaccountingBondClaimedUnstETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingBondClaimedUnstETH)
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
		it.Event = new(CsaccountingBondClaimedUnstETH)
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
func (it *CsaccountingBondClaimedUnstETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingBondClaimedUnstETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingBondClaimedUnstETH represents a BondClaimedUnstETH event raised by the Csaccounting contract.
type CsaccountingBondClaimedUnstETH struct {
	NodeOperatorId *big.Int
	To             common.Address
	Amount         *big.Int
	RequestId      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBondClaimedUnstETH is a free log retrieval operation binding the contract event 0x26673a9d018b21192d08ee14377b798f11b9e5b15ea1559c110265716b8985b5.
//
// Solidity: event BondClaimedUnstETH(uint256 indexed nodeOperatorId, address to, uint256 amount, uint256 requestId)
func (_Csaccounting *CsaccountingFilterer) FilterBondClaimedUnstETH(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsaccountingBondClaimedUnstETHIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "BondClaimedUnstETH", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsaccountingBondClaimedUnstETHIterator{contract: _Csaccounting.contract, event: "BondClaimedUnstETH", logs: logs, sub: sub}, nil
}

// WatchBondClaimedUnstETH is a free log subscription operation binding the contract event 0x26673a9d018b21192d08ee14377b798f11b9e5b15ea1559c110265716b8985b5.
//
// Solidity: event BondClaimedUnstETH(uint256 indexed nodeOperatorId, address to, uint256 amount, uint256 requestId)
func (_Csaccounting *CsaccountingFilterer) WatchBondClaimedUnstETH(opts *bind.WatchOpts, sink chan<- *CsaccountingBondClaimedUnstETH, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "BondClaimedUnstETH", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingBondClaimedUnstETH)
				if err := _Csaccounting.contract.UnpackLog(event, "BondClaimedUnstETH", log); err != nil {
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

// ParseBondClaimedUnstETH is a log parse operation binding the contract event 0x26673a9d018b21192d08ee14377b798f11b9e5b15ea1559c110265716b8985b5.
//
// Solidity: event BondClaimedUnstETH(uint256 indexed nodeOperatorId, address to, uint256 amount, uint256 requestId)
func (_Csaccounting *CsaccountingFilterer) ParseBondClaimedUnstETH(log types.Log) (*CsaccountingBondClaimedUnstETH, error) {
	event := new(CsaccountingBondClaimedUnstETH)
	if err := _Csaccounting.contract.UnpackLog(event, "BondClaimedUnstETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingBondClaimedWstETHIterator is returned from FilterBondClaimedWstETH and is used to iterate over the raw logs and unpacked data for BondClaimedWstETH events raised by the Csaccounting contract.
type CsaccountingBondClaimedWstETHIterator struct {
	Event *CsaccountingBondClaimedWstETH // Event containing the contract specifics and raw log

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
func (it *CsaccountingBondClaimedWstETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingBondClaimedWstETH)
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
		it.Event = new(CsaccountingBondClaimedWstETH)
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
func (it *CsaccountingBondClaimedWstETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingBondClaimedWstETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingBondClaimedWstETH represents a BondClaimedWstETH event raised by the Csaccounting contract.
type CsaccountingBondClaimedWstETH struct {
	NodeOperatorId *big.Int
	To             common.Address
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBondClaimedWstETH is a free log retrieval operation binding the contract event 0xe6a8c06447e05a412e5e9581e088941f3994db3d8a9bfd3275b38d77acacafac.
//
// Solidity: event BondClaimedWstETH(uint256 indexed nodeOperatorId, address to, uint256 amount)
func (_Csaccounting *CsaccountingFilterer) FilterBondClaimedWstETH(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsaccountingBondClaimedWstETHIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "BondClaimedWstETH", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsaccountingBondClaimedWstETHIterator{contract: _Csaccounting.contract, event: "BondClaimedWstETH", logs: logs, sub: sub}, nil
}

// WatchBondClaimedWstETH is a free log subscription operation binding the contract event 0xe6a8c06447e05a412e5e9581e088941f3994db3d8a9bfd3275b38d77acacafac.
//
// Solidity: event BondClaimedWstETH(uint256 indexed nodeOperatorId, address to, uint256 amount)
func (_Csaccounting *CsaccountingFilterer) WatchBondClaimedWstETH(opts *bind.WatchOpts, sink chan<- *CsaccountingBondClaimedWstETH, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "BondClaimedWstETH", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingBondClaimedWstETH)
				if err := _Csaccounting.contract.UnpackLog(event, "BondClaimedWstETH", log); err != nil {
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

// ParseBondClaimedWstETH is a log parse operation binding the contract event 0xe6a8c06447e05a412e5e9581e088941f3994db3d8a9bfd3275b38d77acacafac.
//
// Solidity: event BondClaimedWstETH(uint256 indexed nodeOperatorId, address to, uint256 amount)
func (_Csaccounting *CsaccountingFilterer) ParseBondClaimedWstETH(log types.Log) (*CsaccountingBondClaimedWstETH, error) {
	event := new(CsaccountingBondClaimedWstETH)
	if err := _Csaccounting.contract.UnpackLog(event, "BondClaimedWstETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingBondCurveAddedIterator is returned from FilterBondCurveAdded and is used to iterate over the raw logs and unpacked data for BondCurveAdded events raised by the Csaccounting contract.
type CsaccountingBondCurveAddedIterator struct {
	Event *CsaccountingBondCurveAdded // Event containing the contract specifics and raw log

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
func (it *CsaccountingBondCurveAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingBondCurveAdded)
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
		it.Event = new(CsaccountingBondCurveAdded)
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
func (it *CsaccountingBondCurveAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingBondCurveAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingBondCurveAdded represents a BondCurveAdded event raised by the Csaccounting contract.
type CsaccountingBondCurveAdded struct {
	BondCurve []*big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBondCurveAdded is a free log retrieval operation binding the contract event 0x1fb1d9b944dd7015e95b7b7a9623c45792e4532badcf9c6e7a284d7d4d0570f0.
//
// Solidity: event BondCurveAdded(uint256[] bondCurve)
func (_Csaccounting *CsaccountingFilterer) FilterBondCurveAdded(opts *bind.FilterOpts) (*CsaccountingBondCurveAddedIterator, error) {

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "BondCurveAdded")
	if err != nil {
		return nil, err
	}
	return &CsaccountingBondCurveAddedIterator{contract: _Csaccounting.contract, event: "BondCurveAdded", logs: logs, sub: sub}, nil
}

// WatchBondCurveAdded is a free log subscription operation binding the contract event 0x1fb1d9b944dd7015e95b7b7a9623c45792e4532badcf9c6e7a284d7d4d0570f0.
//
// Solidity: event BondCurveAdded(uint256[] bondCurve)
func (_Csaccounting *CsaccountingFilterer) WatchBondCurveAdded(opts *bind.WatchOpts, sink chan<- *CsaccountingBondCurveAdded) (event.Subscription, error) {

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "BondCurveAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingBondCurveAdded)
				if err := _Csaccounting.contract.UnpackLog(event, "BondCurveAdded", log); err != nil {
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

// ParseBondCurveAdded is a log parse operation binding the contract event 0x1fb1d9b944dd7015e95b7b7a9623c45792e4532badcf9c6e7a284d7d4d0570f0.
//
// Solidity: event BondCurveAdded(uint256[] bondCurve)
func (_Csaccounting *CsaccountingFilterer) ParseBondCurveAdded(log types.Log) (*CsaccountingBondCurveAdded, error) {
	event := new(CsaccountingBondCurveAdded)
	if err := _Csaccounting.contract.UnpackLog(event, "BondCurveAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingBondCurveSetIterator is returned from FilterBondCurveSet and is used to iterate over the raw logs and unpacked data for BondCurveSet events raised by the Csaccounting contract.
type CsaccountingBondCurveSetIterator struct {
	Event *CsaccountingBondCurveSet // Event containing the contract specifics and raw log

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
func (it *CsaccountingBondCurveSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingBondCurveSet)
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
		it.Event = new(CsaccountingBondCurveSet)
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
func (it *CsaccountingBondCurveSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingBondCurveSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingBondCurveSet represents a BondCurveSet event raised by the Csaccounting contract.
type CsaccountingBondCurveSet struct {
	NodeOperatorId *big.Int
	CurveId        *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBondCurveSet is a free log retrieval operation binding the contract event 0x4642db1736894887bc907d721f20af84d3e585a0a3cea90f41b78b2aa906541b.
//
// Solidity: event BondCurveSet(uint256 indexed nodeOperatorId, uint256 curveId)
func (_Csaccounting *CsaccountingFilterer) FilterBondCurveSet(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsaccountingBondCurveSetIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "BondCurveSet", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsaccountingBondCurveSetIterator{contract: _Csaccounting.contract, event: "BondCurveSet", logs: logs, sub: sub}, nil
}

// WatchBondCurveSet is a free log subscription operation binding the contract event 0x4642db1736894887bc907d721f20af84d3e585a0a3cea90f41b78b2aa906541b.
//
// Solidity: event BondCurveSet(uint256 indexed nodeOperatorId, uint256 curveId)
func (_Csaccounting *CsaccountingFilterer) WatchBondCurveSet(opts *bind.WatchOpts, sink chan<- *CsaccountingBondCurveSet, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "BondCurveSet", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingBondCurveSet)
				if err := _Csaccounting.contract.UnpackLog(event, "BondCurveSet", log); err != nil {
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

// ParseBondCurveSet is a log parse operation binding the contract event 0x4642db1736894887bc907d721f20af84d3e585a0a3cea90f41b78b2aa906541b.
//
// Solidity: event BondCurveSet(uint256 indexed nodeOperatorId, uint256 curveId)
func (_Csaccounting *CsaccountingFilterer) ParseBondCurveSet(log types.Log) (*CsaccountingBondCurveSet, error) {
	event := new(CsaccountingBondCurveSet)
	if err := _Csaccounting.contract.UnpackLog(event, "BondCurveSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingBondCurveUpdatedIterator is returned from FilterBondCurveUpdated and is used to iterate over the raw logs and unpacked data for BondCurveUpdated events raised by the Csaccounting contract.
type CsaccountingBondCurveUpdatedIterator struct {
	Event *CsaccountingBondCurveUpdated // Event containing the contract specifics and raw log

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
func (it *CsaccountingBondCurveUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingBondCurveUpdated)
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
		it.Event = new(CsaccountingBondCurveUpdated)
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
func (it *CsaccountingBondCurveUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingBondCurveUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingBondCurveUpdated represents a BondCurveUpdated event raised by the Csaccounting contract.
type CsaccountingBondCurveUpdated struct {
	CurveId   *big.Int
	BondCurve []*big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBondCurveUpdated is a free log retrieval operation binding the contract event 0x53da7af401538204fd91f2946f2fe85d05224d2cc766fd7aa9fbd8bf4fb4ce9f.
//
// Solidity: event BondCurveUpdated(uint256 indexed curveId, uint256[] bondCurve)
func (_Csaccounting *CsaccountingFilterer) FilterBondCurveUpdated(opts *bind.FilterOpts, curveId []*big.Int) (*CsaccountingBondCurveUpdatedIterator, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "BondCurveUpdated", curveIdRule)
	if err != nil {
		return nil, err
	}
	return &CsaccountingBondCurveUpdatedIterator{contract: _Csaccounting.contract, event: "BondCurveUpdated", logs: logs, sub: sub}, nil
}

// WatchBondCurveUpdated is a free log subscription operation binding the contract event 0x53da7af401538204fd91f2946f2fe85d05224d2cc766fd7aa9fbd8bf4fb4ce9f.
//
// Solidity: event BondCurveUpdated(uint256 indexed curveId, uint256[] bondCurve)
func (_Csaccounting *CsaccountingFilterer) WatchBondCurveUpdated(opts *bind.WatchOpts, sink chan<- *CsaccountingBondCurveUpdated, curveId []*big.Int) (event.Subscription, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "BondCurveUpdated", curveIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingBondCurveUpdated)
				if err := _Csaccounting.contract.UnpackLog(event, "BondCurveUpdated", log); err != nil {
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

// ParseBondCurveUpdated is a log parse operation binding the contract event 0x53da7af401538204fd91f2946f2fe85d05224d2cc766fd7aa9fbd8bf4fb4ce9f.
//
// Solidity: event BondCurveUpdated(uint256 indexed curveId, uint256[] bondCurve)
func (_Csaccounting *CsaccountingFilterer) ParseBondCurveUpdated(log types.Log) (*CsaccountingBondCurveUpdated, error) {
	event := new(CsaccountingBondCurveUpdated)
	if err := _Csaccounting.contract.UnpackLog(event, "BondCurveUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingBondDepositedETHIterator is returned from FilterBondDepositedETH and is used to iterate over the raw logs and unpacked data for BondDepositedETH events raised by the Csaccounting contract.
type CsaccountingBondDepositedETHIterator struct {
	Event *CsaccountingBondDepositedETH // Event containing the contract specifics and raw log

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
func (it *CsaccountingBondDepositedETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingBondDepositedETH)
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
		it.Event = new(CsaccountingBondDepositedETH)
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
func (it *CsaccountingBondDepositedETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingBondDepositedETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingBondDepositedETH represents a BondDepositedETH event raised by the Csaccounting contract.
type CsaccountingBondDepositedETH struct {
	NodeOperatorId *big.Int
	From           common.Address
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBondDepositedETH is a free log retrieval operation binding the contract event 0x16ec5116295424dec7fd52c87d9971a963ea7f59f741ad9ad468f0312055dc49.
//
// Solidity: event BondDepositedETH(uint256 indexed nodeOperatorId, address from, uint256 amount)
func (_Csaccounting *CsaccountingFilterer) FilterBondDepositedETH(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsaccountingBondDepositedETHIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "BondDepositedETH", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsaccountingBondDepositedETHIterator{contract: _Csaccounting.contract, event: "BondDepositedETH", logs: logs, sub: sub}, nil
}

// WatchBondDepositedETH is a free log subscription operation binding the contract event 0x16ec5116295424dec7fd52c87d9971a963ea7f59f741ad9ad468f0312055dc49.
//
// Solidity: event BondDepositedETH(uint256 indexed nodeOperatorId, address from, uint256 amount)
func (_Csaccounting *CsaccountingFilterer) WatchBondDepositedETH(opts *bind.WatchOpts, sink chan<- *CsaccountingBondDepositedETH, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "BondDepositedETH", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingBondDepositedETH)
				if err := _Csaccounting.contract.UnpackLog(event, "BondDepositedETH", log); err != nil {
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

// ParseBondDepositedETH is a log parse operation binding the contract event 0x16ec5116295424dec7fd52c87d9971a963ea7f59f741ad9ad468f0312055dc49.
//
// Solidity: event BondDepositedETH(uint256 indexed nodeOperatorId, address from, uint256 amount)
func (_Csaccounting *CsaccountingFilterer) ParseBondDepositedETH(log types.Log) (*CsaccountingBondDepositedETH, error) {
	event := new(CsaccountingBondDepositedETH)
	if err := _Csaccounting.contract.UnpackLog(event, "BondDepositedETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingBondDepositedStETHIterator is returned from FilterBondDepositedStETH and is used to iterate over the raw logs and unpacked data for BondDepositedStETH events raised by the Csaccounting contract.
type CsaccountingBondDepositedStETHIterator struct {
	Event *CsaccountingBondDepositedStETH // Event containing the contract specifics and raw log

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
func (it *CsaccountingBondDepositedStETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingBondDepositedStETH)
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
		it.Event = new(CsaccountingBondDepositedStETH)
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
func (it *CsaccountingBondDepositedStETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingBondDepositedStETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingBondDepositedStETH represents a BondDepositedStETH event raised by the Csaccounting contract.
type CsaccountingBondDepositedStETH struct {
	NodeOperatorId *big.Int
	From           common.Address
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBondDepositedStETH is a free log retrieval operation binding the contract event 0xee31ebba29fd5471227e42fd8ca621a892d689901892cb8febb03fe802c3214b.
//
// Solidity: event BondDepositedStETH(uint256 indexed nodeOperatorId, address from, uint256 amount)
func (_Csaccounting *CsaccountingFilterer) FilterBondDepositedStETH(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsaccountingBondDepositedStETHIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "BondDepositedStETH", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsaccountingBondDepositedStETHIterator{contract: _Csaccounting.contract, event: "BondDepositedStETH", logs: logs, sub: sub}, nil
}

// WatchBondDepositedStETH is a free log subscription operation binding the contract event 0xee31ebba29fd5471227e42fd8ca621a892d689901892cb8febb03fe802c3214b.
//
// Solidity: event BondDepositedStETH(uint256 indexed nodeOperatorId, address from, uint256 amount)
func (_Csaccounting *CsaccountingFilterer) WatchBondDepositedStETH(opts *bind.WatchOpts, sink chan<- *CsaccountingBondDepositedStETH, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "BondDepositedStETH", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingBondDepositedStETH)
				if err := _Csaccounting.contract.UnpackLog(event, "BondDepositedStETH", log); err != nil {
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

// ParseBondDepositedStETH is a log parse operation binding the contract event 0xee31ebba29fd5471227e42fd8ca621a892d689901892cb8febb03fe802c3214b.
//
// Solidity: event BondDepositedStETH(uint256 indexed nodeOperatorId, address from, uint256 amount)
func (_Csaccounting *CsaccountingFilterer) ParseBondDepositedStETH(log types.Log) (*CsaccountingBondDepositedStETH, error) {
	event := new(CsaccountingBondDepositedStETH)
	if err := _Csaccounting.contract.UnpackLog(event, "BondDepositedStETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingBondDepositedWstETHIterator is returned from FilterBondDepositedWstETH and is used to iterate over the raw logs and unpacked data for BondDepositedWstETH events raised by the Csaccounting contract.
type CsaccountingBondDepositedWstETHIterator struct {
	Event *CsaccountingBondDepositedWstETH // Event containing the contract specifics and raw log

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
func (it *CsaccountingBondDepositedWstETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingBondDepositedWstETH)
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
		it.Event = new(CsaccountingBondDepositedWstETH)
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
func (it *CsaccountingBondDepositedWstETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingBondDepositedWstETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingBondDepositedWstETH represents a BondDepositedWstETH event raised by the Csaccounting contract.
type CsaccountingBondDepositedWstETH struct {
	NodeOperatorId *big.Int
	From           common.Address
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBondDepositedWstETH is a free log retrieval operation binding the contract event 0x6576bbc9c5b478bf9717dc3d2bcb485e5ff0727df77c72558727597f3609d3f1.
//
// Solidity: event BondDepositedWstETH(uint256 indexed nodeOperatorId, address from, uint256 amount)
func (_Csaccounting *CsaccountingFilterer) FilterBondDepositedWstETH(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsaccountingBondDepositedWstETHIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "BondDepositedWstETH", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsaccountingBondDepositedWstETHIterator{contract: _Csaccounting.contract, event: "BondDepositedWstETH", logs: logs, sub: sub}, nil
}

// WatchBondDepositedWstETH is a free log subscription operation binding the contract event 0x6576bbc9c5b478bf9717dc3d2bcb485e5ff0727df77c72558727597f3609d3f1.
//
// Solidity: event BondDepositedWstETH(uint256 indexed nodeOperatorId, address from, uint256 amount)
func (_Csaccounting *CsaccountingFilterer) WatchBondDepositedWstETH(opts *bind.WatchOpts, sink chan<- *CsaccountingBondDepositedWstETH, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "BondDepositedWstETH", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingBondDepositedWstETH)
				if err := _Csaccounting.contract.UnpackLog(event, "BondDepositedWstETH", log); err != nil {
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

// ParseBondDepositedWstETH is a log parse operation binding the contract event 0x6576bbc9c5b478bf9717dc3d2bcb485e5ff0727df77c72558727597f3609d3f1.
//
// Solidity: event BondDepositedWstETH(uint256 indexed nodeOperatorId, address from, uint256 amount)
func (_Csaccounting *CsaccountingFilterer) ParseBondDepositedWstETH(log types.Log) (*CsaccountingBondDepositedWstETH, error) {
	event := new(CsaccountingBondDepositedWstETH)
	if err := _Csaccounting.contract.UnpackLog(event, "BondDepositedWstETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingBondLockChangedIterator is returned from FilterBondLockChanged and is used to iterate over the raw logs and unpacked data for BondLockChanged events raised by the Csaccounting contract.
type CsaccountingBondLockChangedIterator struct {
	Event *CsaccountingBondLockChanged // Event containing the contract specifics and raw log

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
func (it *CsaccountingBondLockChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingBondLockChanged)
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
		it.Event = new(CsaccountingBondLockChanged)
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
func (it *CsaccountingBondLockChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingBondLockChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingBondLockChanged represents a BondLockChanged event raised by the Csaccounting contract.
type CsaccountingBondLockChanged struct {
	NodeOperatorId *big.Int
	NewAmount      *big.Int
	RetentionUntil *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBondLockChanged is a free log retrieval operation binding the contract event 0x69a153d448f54b17f05cf3b268a2efab87c94a4727d108c4ca4aa3e5d65113de.
//
// Solidity: event BondLockChanged(uint256 indexed nodeOperatorId, uint256 newAmount, uint256 retentionUntil)
func (_Csaccounting *CsaccountingFilterer) FilterBondLockChanged(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsaccountingBondLockChangedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "BondLockChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsaccountingBondLockChangedIterator{contract: _Csaccounting.contract, event: "BondLockChanged", logs: logs, sub: sub}, nil
}

// WatchBondLockChanged is a free log subscription operation binding the contract event 0x69a153d448f54b17f05cf3b268a2efab87c94a4727d108c4ca4aa3e5d65113de.
//
// Solidity: event BondLockChanged(uint256 indexed nodeOperatorId, uint256 newAmount, uint256 retentionUntil)
func (_Csaccounting *CsaccountingFilterer) WatchBondLockChanged(opts *bind.WatchOpts, sink chan<- *CsaccountingBondLockChanged, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "BondLockChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingBondLockChanged)
				if err := _Csaccounting.contract.UnpackLog(event, "BondLockChanged", log); err != nil {
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

// ParseBondLockChanged is a log parse operation binding the contract event 0x69a153d448f54b17f05cf3b268a2efab87c94a4727d108c4ca4aa3e5d65113de.
//
// Solidity: event BondLockChanged(uint256 indexed nodeOperatorId, uint256 newAmount, uint256 retentionUntil)
func (_Csaccounting *CsaccountingFilterer) ParseBondLockChanged(log types.Log) (*CsaccountingBondLockChanged, error) {
	event := new(CsaccountingBondLockChanged)
	if err := _Csaccounting.contract.UnpackLog(event, "BondLockChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingBondLockCompensatedIterator is returned from FilterBondLockCompensated and is used to iterate over the raw logs and unpacked data for BondLockCompensated events raised by the Csaccounting contract.
type CsaccountingBondLockCompensatedIterator struct {
	Event *CsaccountingBondLockCompensated // Event containing the contract specifics and raw log

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
func (it *CsaccountingBondLockCompensatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingBondLockCompensated)
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
		it.Event = new(CsaccountingBondLockCompensated)
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
func (it *CsaccountingBondLockCompensatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingBondLockCompensatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingBondLockCompensated represents a BondLockCompensated event raised by the Csaccounting contract.
type CsaccountingBondLockCompensated struct {
	NodeOperatorId *big.Int
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBondLockCompensated is a free log retrieval operation binding the contract event 0xb6ee6e3aae6776519627b46786a622b642c38cabfe4c97cb34054fd63fc11a23.
//
// Solidity: event BondLockCompensated(uint256 indexed nodeOperatorId, uint256 amount)
func (_Csaccounting *CsaccountingFilterer) FilterBondLockCompensated(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsaccountingBondLockCompensatedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "BondLockCompensated", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsaccountingBondLockCompensatedIterator{contract: _Csaccounting.contract, event: "BondLockCompensated", logs: logs, sub: sub}, nil
}

// WatchBondLockCompensated is a free log subscription operation binding the contract event 0xb6ee6e3aae6776519627b46786a622b642c38cabfe4c97cb34054fd63fc11a23.
//
// Solidity: event BondLockCompensated(uint256 indexed nodeOperatorId, uint256 amount)
func (_Csaccounting *CsaccountingFilterer) WatchBondLockCompensated(opts *bind.WatchOpts, sink chan<- *CsaccountingBondLockCompensated, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "BondLockCompensated", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingBondLockCompensated)
				if err := _Csaccounting.contract.UnpackLog(event, "BondLockCompensated", log); err != nil {
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

// ParseBondLockCompensated is a log parse operation binding the contract event 0xb6ee6e3aae6776519627b46786a622b642c38cabfe4c97cb34054fd63fc11a23.
//
// Solidity: event BondLockCompensated(uint256 indexed nodeOperatorId, uint256 amount)
func (_Csaccounting *CsaccountingFilterer) ParseBondLockCompensated(log types.Log) (*CsaccountingBondLockCompensated, error) {
	event := new(CsaccountingBondLockCompensated)
	if err := _Csaccounting.contract.UnpackLog(event, "BondLockCompensated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingBondLockRemovedIterator is returned from FilterBondLockRemoved and is used to iterate over the raw logs and unpacked data for BondLockRemoved events raised by the Csaccounting contract.
type CsaccountingBondLockRemovedIterator struct {
	Event *CsaccountingBondLockRemoved // Event containing the contract specifics and raw log

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
func (it *CsaccountingBondLockRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingBondLockRemoved)
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
		it.Event = new(CsaccountingBondLockRemoved)
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
func (it *CsaccountingBondLockRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingBondLockRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingBondLockRemoved represents a BondLockRemoved event raised by the Csaccounting contract.
type CsaccountingBondLockRemoved struct {
	NodeOperatorId *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBondLockRemoved is a free log retrieval operation binding the contract event 0x844ae6b00e8a437dcdde1a634feab3273e08bb5c274a4be3b195b308ae0ba20a.
//
// Solidity: event BondLockRemoved(uint256 indexed nodeOperatorId)
func (_Csaccounting *CsaccountingFilterer) FilterBondLockRemoved(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsaccountingBondLockRemovedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "BondLockRemoved", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsaccountingBondLockRemovedIterator{contract: _Csaccounting.contract, event: "BondLockRemoved", logs: logs, sub: sub}, nil
}

// WatchBondLockRemoved is a free log subscription operation binding the contract event 0x844ae6b00e8a437dcdde1a634feab3273e08bb5c274a4be3b195b308ae0ba20a.
//
// Solidity: event BondLockRemoved(uint256 indexed nodeOperatorId)
func (_Csaccounting *CsaccountingFilterer) WatchBondLockRemoved(opts *bind.WatchOpts, sink chan<- *CsaccountingBondLockRemoved, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "BondLockRemoved", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingBondLockRemoved)
				if err := _Csaccounting.contract.UnpackLog(event, "BondLockRemoved", log); err != nil {
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

// ParseBondLockRemoved is a log parse operation binding the contract event 0x844ae6b00e8a437dcdde1a634feab3273e08bb5c274a4be3b195b308ae0ba20a.
//
// Solidity: event BondLockRemoved(uint256 indexed nodeOperatorId)
func (_Csaccounting *CsaccountingFilterer) ParseBondLockRemoved(log types.Log) (*CsaccountingBondLockRemoved, error) {
	event := new(CsaccountingBondLockRemoved)
	if err := _Csaccounting.contract.UnpackLog(event, "BondLockRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingBondLockRetentionPeriodChangedIterator is returned from FilterBondLockRetentionPeriodChanged and is used to iterate over the raw logs and unpacked data for BondLockRetentionPeriodChanged events raised by the Csaccounting contract.
type CsaccountingBondLockRetentionPeriodChangedIterator struct {
	Event *CsaccountingBondLockRetentionPeriodChanged // Event containing the contract specifics and raw log

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
func (it *CsaccountingBondLockRetentionPeriodChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingBondLockRetentionPeriodChanged)
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
		it.Event = new(CsaccountingBondLockRetentionPeriodChanged)
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
func (it *CsaccountingBondLockRetentionPeriodChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingBondLockRetentionPeriodChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingBondLockRetentionPeriodChanged represents a BondLockRetentionPeriodChanged event raised by the Csaccounting contract.
type CsaccountingBondLockRetentionPeriodChanged struct {
	RetentionPeriod *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBondLockRetentionPeriodChanged is a free log retrieval operation binding the contract event 0xdaf5eddbe9ed0768e54cc8f739a9cb86c57fc70da07eff01d9ba886f21a7a4b3.
//
// Solidity: event BondLockRetentionPeriodChanged(uint256 retentionPeriod)
func (_Csaccounting *CsaccountingFilterer) FilterBondLockRetentionPeriodChanged(opts *bind.FilterOpts) (*CsaccountingBondLockRetentionPeriodChangedIterator, error) {

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "BondLockRetentionPeriodChanged")
	if err != nil {
		return nil, err
	}
	return &CsaccountingBondLockRetentionPeriodChangedIterator{contract: _Csaccounting.contract, event: "BondLockRetentionPeriodChanged", logs: logs, sub: sub}, nil
}

// WatchBondLockRetentionPeriodChanged is a free log subscription operation binding the contract event 0xdaf5eddbe9ed0768e54cc8f739a9cb86c57fc70da07eff01d9ba886f21a7a4b3.
//
// Solidity: event BondLockRetentionPeriodChanged(uint256 retentionPeriod)
func (_Csaccounting *CsaccountingFilterer) WatchBondLockRetentionPeriodChanged(opts *bind.WatchOpts, sink chan<- *CsaccountingBondLockRetentionPeriodChanged) (event.Subscription, error) {

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "BondLockRetentionPeriodChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingBondLockRetentionPeriodChanged)
				if err := _Csaccounting.contract.UnpackLog(event, "BondLockRetentionPeriodChanged", log); err != nil {
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

// ParseBondLockRetentionPeriodChanged is a log parse operation binding the contract event 0xdaf5eddbe9ed0768e54cc8f739a9cb86c57fc70da07eff01d9ba886f21a7a4b3.
//
// Solidity: event BondLockRetentionPeriodChanged(uint256 retentionPeriod)
func (_Csaccounting *CsaccountingFilterer) ParseBondLockRetentionPeriodChanged(log types.Log) (*CsaccountingBondLockRetentionPeriodChanged, error) {
	event := new(CsaccountingBondLockRetentionPeriodChanged)
	if err := _Csaccounting.contract.UnpackLog(event, "BondLockRetentionPeriodChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingChargePenaltyRecipientSetIterator is returned from FilterChargePenaltyRecipientSet and is used to iterate over the raw logs and unpacked data for ChargePenaltyRecipientSet events raised by the Csaccounting contract.
type CsaccountingChargePenaltyRecipientSetIterator struct {
	Event *CsaccountingChargePenaltyRecipientSet // Event containing the contract specifics and raw log

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
func (it *CsaccountingChargePenaltyRecipientSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingChargePenaltyRecipientSet)
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
		it.Event = new(CsaccountingChargePenaltyRecipientSet)
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
func (it *CsaccountingChargePenaltyRecipientSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingChargePenaltyRecipientSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingChargePenaltyRecipientSet represents a ChargePenaltyRecipientSet event raised by the Csaccounting contract.
type CsaccountingChargePenaltyRecipientSet struct {
	ChargePenaltyRecipient common.Address
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterChargePenaltyRecipientSet is a free log retrieval operation binding the contract event 0x4beaaee83871b066b675515d6a53567e76411f60266703cef934a01905a4d832.
//
// Solidity: event ChargePenaltyRecipientSet(address chargePenaltyRecipient)
func (_Csaccounting *CsaccountingFilterer) FilterChargePenaltyRecipientSet(opts *bind.FilterOpts) (*CsaccountingChargePenaltyRecipientSetIterator, error) {

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "ChargePenaltyRecipientSet")
	if err != nil {
		return nil, err
	}
	return &CsaccountingChargePenaltyRecipientSetIterator{contract: _Csaccounting.contract, event: "ChargePenaltyRecipientSet", logs: logs, sub: sub}, nil
}

// WatchChargePenaltyRecipientSet is a free log subscription operation binding the contract event 0x4beaaee83871b066b675515d6a53567e76411f60266703cef934a01905a4d832.
//
// Solidity: event ChargePenaltyRecipientSet(address chargePenaltyRecipient)
func (_Csaccounting *CsaccountingFilterer) WatchChargePenaltyRecipientSet(opts *bind.WatchOpts, sink chan<- *CsaccountingChargePenaltyRecipientSet) (event.Subscription, error) {

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "ChargePenaltyRecipientSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingChargePenaltyRecipientSet)
				if err := _Csaccounting.contract.UnpackLog(event, "ChargePenaltyRecipientSet", log); err != nil {
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

// ParseChargePenaltyRecipientSet is a log parse operation binding the contract event 0x4beaaee83871b066b675515d6a53567e76411f60266703cef934a01905a4d832.
//
// Solidity: event ChargePenaltyRecipientSet(address chargePenaltyRecipient)
func (_Csaccounting *CsaccountingFilterer) ParseChargePenaltyRecipientSet(log types.Log) (*CsaccountingChargePenaltyRecipientSet, error) {
	event := new(CsaccountingChargePenaltyRecipientSet)
	if err := _Csaccounting.contract.UnpackLog(event, "ChargePenaltyRecipientSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingERC1155RecoveredIterator is returned from FilterERC1155Recovered and is used to iterate over the raw logs and unpacked data for ERC1155Recovered events raised by the Csaccounting contract.
type CsaccountingERC1155RecoveredIterator struct {
	Event *CsaccountingERC1155Recovered // Event containing the contract specifics and raw log

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
func (it *CsaccountingERC1155RecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingERC1155Recovered)
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
		it.Event = new(CsaccountingERC1155Recovered)
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
func (it *CsaccountingERC1155RecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingERC1155RecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingERC1155Recovered represents a ERC1155Recovered event raised by the Csaccounting contract.
type CsaccountingERC1155Recovered struct {
	Token     common.Address
	TokenId   *big.Int
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERC1155Recovered is a free log retrieval operation binding the contract event 0x5cf02e753b3eb0f4bee4460a72817d8e5e3c75cd4d65c1d0b06dca88b8032936.
//
// Solidity: event ERC1155Recovered(address indexed token, uint256 tokenId, address indexed recipient, uint256 amount)
func (_Csaccounting *CsaccountingFilterer) FilterERC1155Recovered(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*CsaccountingERC1155RecoveredIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "ERC1155Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &CsaccountingERC1155RecoveredIterator{contract: _Csaccounting.contract, event: "ERC1155Recovered", logs: logs, sub: sub}, nil
}

// WatchERC1155Recovered is a free log subscription operation binding the contract event 0x5cf02e753b3eb0f4bee4460a72817d8e5e3c75cd4d65c1d0b06dca88b8032936.
//
// Solidity: event ERC1155Recovered(address indexed token, uint256 tokenId, address indexed recipient, uint256 amount)
func (_Csaccounting *CsaccountingFilterer) WatchERC1155Recovered(opts *bind.WatchOpts, sink chan<- *CsaccountingERC1155Recovered, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "ERC1155Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingERC1155Recovered)
				if err := _Csaccounting.contract.UnpackLog(event, "ERC1155Recovered", log); err != nil {
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
func (_Csaccounting *CsaccountingFilterer) ParseERC1155Recovered(log types.Log) (*CsaccountingERC1155Recovered, error) {
	event := new(CsaccountingERC1155Recovered)
	if err := _Csaccounting.contract.UnpackLog(event, "ERC1155Recovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingERC20RecoveredIterator is returned from FilterERC20Recovered and is used to iterate over the raw logs and unpacked data for ERC20Recovered events raised by the Csaccounting contract.
type CsaccountingERC20RecoveredIterator struct {
	Event *CsaccountingERC20Recovered // Event containing the contract specifics and raw log

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
func (it *CsaccountingERC20RecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingERC20Recovered)
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
		it.Event = new(CsaccountingERC20Recovered)
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
func (it *CsaccountingERC20RecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingERC20RecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingERC20Recovered represents a ERC20Recovered event raised by the Csaccounting contract.
type CsaccountingERC20Recovered struct {
	Token     common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERC20Recovered is a free log retrieval operation binding the contract event 0xaca8fb252cde442184e5f10e0f2e6e4029e8cd7717cae63559079610702436aa.
//
// Solidity: event ERC20Recovered(address indexed token, address indexed recipient, uint256 amount)
func (_Csaccounting *CsaccountingFilterer) FilterERC20Recovered(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*CsaccountingERC20RecoveredIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "ERC20Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &CsaccountingERC20RecoveredIterator{contract: _Csaccounting.contract, event: "ERC20Recovered", logs: logs, sub: sub}, nil
}

// WatchERC20Recovered is a free log subscription operation binding the contract event 0xaca8fb252cde442184e5f10e0f2e6e4029e8cd7717cae63559079610702436aa.
//
// Solidity: event ERC20Recovered(address indexed token, address indexed recipient, uint256 amount)
func (_Csaccounting *CsaccountingFilterer) WatchERC20Recovered(opts *bind.WatchOpts, sink chan<- *CsaccountingERC20Recovered, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "ERC20Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingERC20Recovered)
				if err := _Csaccounting.contract.UnpackLog(event, "ERC20Recovered", log); err != nil {
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
func (_Csaccounting *CsaccountingFilterer) ParseERC20Recovered(log types.Log) (*CsaccountingERC20Recovered, error) {
	event := new(CsaccountingERC20Recovered)
	if err := _Csaccounting.contract.UnpackLog(event, "ERC20Recovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingERC721RecoveredIterator is returned from FilterERC721Recovered and is used to iterate over the raw logs and unpacked data for ERC721Recovered events raised by the Csaccounting contract.
type CsaccountingERC721RecoveredIterator struct {
	Event *CsaccountingERC721Recovered // Event containing the contract specifics and raw log

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
func (it *CsaccountingERC721RecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingERC721Recovered)
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
		it.Event = new(CsaccountingERC721Recovered)
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
func (it *CsaccountingERC721RecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingERC721RecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingERC721Recovered represents a ERC721Recovered event raised by the Csaccounting contract.
type CsaccountingERC721Recovered struct {
	Token     common.Address
	TokenId   *big.Int
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERC721Recovered is a free log retrieval operation binding the contract event 0x8166bf75d2ff2fa3c8f3c44410540bf42e9a5359b48409e8d660291dc9f788c8.
//
// Solidity: event ERC721Recovered(address indexed token, uint256 tokenId, address indexed recipient)
func (_Csaccounting *CsaccountingFilterer) FilterERC721Recovered(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*CsaccountingERC721RecoveredIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "ERC721Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &CsaccountingERC721RecoveredIterator{contract: _Csaccounting.contract, event: "ERC721Recovered", logs: logs, sub: sub}, nil
}

// WatchERC721Recovered is a free log subscription operation binding the contract event 0x8166bf75d2ff2fa3c8f3c44410540bf42e9a5359b48409e8d660291dc9f788c8.
//
// Solidity: event ERC721Recovered(address indexed token, uint256 tokenId, address indexed recipient)
func (_Csaccounting *CsaccountingFilterer) WatchERC721Recovered(opts *bind.WatchOpts, sink chan<- *CsaccountingERC721Recovered, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "ERC721Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingERC721Recovered)
				if err := _Csaccounting.contract.UnpackLog(event, "ERC721Recovered", log); err != nil {
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
func (_Csaccounting *CsaccountingFilterer) ParseERC721Recovered(log types.Log) (*CsaccountingERC721Recovered, error) {
	event := new(CsaccountingERC721Recovered)
	if err := _Csaccounting.contract.UnpackLog(event, "ERC721Recovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingEtherRecoveredIterator is returned from FilterEtherRecovered and is used to iterate over the raw logs and unpacked data for EtherRecovered events raised by the Csaccounting contract.
type CsaccountingEtherRecoveredIterator struct {
	Event *CsaccountingEtherRecovered // Event containing the contract specifics and raw log

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
func (it *CsaccountingEtherRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingEtherRecovered)
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
		it.Event = new(CsaccountingEtherRecovered)
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
func (it *CsaccountingEtherRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingEtherRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingEtherRecovered represents a EtherRecovered event raised by the Csaccounting contract.
type CsaccountingEtherRecovered struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEtherRecovered is a free log retrieval operation binding the contract event 0x8e274e42262a7f013b700b35c2b4629ccce1702f8fe83f8dfb7eacbb26a4382c.
//
// Solidity: event EtherRecovered(address indexed recipient, uint256 amount)
func (_Csaccounting *CsaccountingFilterer) FilterEtherRecovered(opts *bind.FilterOpts, recipient []common.Address) (*CsaccountingEtherRecoveredIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "EtherRecovered", recipientRule)
	if err != nil {
		return nil, err
	}
	return &CsaccountingEtherRecoveredIterator{contract: _Csaccounting.contract, event: "EtherRecovered", logs: logs, sub: sub}, nil
}

// WatchEtherRecovered is a free log subscription operation binding the contract event 0x8e274e42262a7f013b700b35c2b4629ccce1702f8fe83f8dfb7eacbb26a4382c.
//
// Solidity: event EtherRecovered(address indexed recipient, uint256 amount)
func (_Csaccounting *CsaccountingFilterer) WatchEtherRecovered(opts *bind.WatchOpts, sink chan<- *CsaccountingEtherRecovered, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "EtherRecovered", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingEtherRecovered)
				if err := _Csaccounting.contract.UnpackLog(event, "EtherRecovered", log); err != nil {
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
func (_Csaccounting *CsaccountingFilterer) ParseEtherRecovered(log types.Log) (*CsaccountingEtherRecovered, error) {
	event := new(CsaccountingEtherRecovered)
	if err := _Csaccounting.contract.UnpackLog(event, "EtherRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Csaccounting contract.
type CsaccountingInitializedIterator struct {
	Event *CsaccountingInitialized // Event containing the contract specifics and raw log

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
func (it *CsaccountingInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingInitialized)
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
		it.Event = new(CsaccountingInitialized)
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
func (it *CsaccountingInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingInitialized represents a Initialized event raised by the Csaccounting contract.
type CsaccountingInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Csaccounting *CsaccountingFilterer) FilterInitialized(opts *bind.FilterOpts) (*CsaccountingInitializedIterator, error) {

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CsaccountingInitializedIterator{contract: _Csaccounting.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Csaccounting *CsaccountingFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CsaccountingInitialized) (event.Subscription, error) {

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingInitialized)
				if err := _Csaccounting.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Csaccounting *CsaccountingFilterer) ParseInitialized(log types.Log) (*CsaccountingInitialized, error) {
	event := new(CsaccountingInitialized)
	if err := _Csaccounting.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Csaccounting contract.
type CsaccountingPausedIterator struct {
	Event *CsaccountingPaused // Event containing the contract specifics and raw log

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
func (it *CsaccountingPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingPaused)
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
		it.Event = new(CsaccountingPaused)
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
func (it *CsaccountingPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingPaused represents a Paused event raised by the Csaccounting contract.
type CsaccountingPaused struct {
	Duration *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e.
//
// Solidity: event Paused(uint256 duration)
func (_Csaccounting *CsaccountingFilterer) FilterPaused(opts *bind.FilterOpts) (*CsaccountingPausedIterator, error) {

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CsaccountingPausedIterator{contract: _Csaccounting.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e.
//
// Solidity: event Paused(uint256 duration)
func (_Csaccounting *CsaccountingFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CsaccountingPaused) (event.Subscription, error) {

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingPaused)
				if err := _Csaccounting.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Csaccounting *CsaccountingFilterer) ParsePaused(log types.Log) (*CsaccountingPaused, error) {
	event := new(CsaccountingPaused)
	if err := _Csaccounting.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingResumedIterator is returned from FilterResumed and is used to iterate over the raw logs and unpacked data for Resumed events raised by the Csaccounting contract.
type CsaccountingResumedIterator struct {
	Event *CsaccountingResumed // Event containing the contract specifics and raw log

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
func (it *CsaccountingResumedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingResumed)
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
		it.Event = new(CsaccountingResumed)
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
func (it *CsaccountingResumedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingResumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingResumed represents a Resumed event raised by the Csaccounting contract.
type CsaccountingResumed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterResumed is a free log retrieval operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_Csaccounting *CsaccountingFilterer) FilterResumed(opts *bind.FilterOpts) (*CsaccountingResumedIterator, error) {

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "Resumed")
	if err != nil {
		return nil, err
	}
	return &CsaccountingResumedIterator{contract: _Csaccounting.contract, event: "Resumed", logs: logs, sub: sub}, nil
}

// WatchResumed is a free log subscription operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_Csaccounting *CsaccountingFilterer) WatchResumed(opts *bind.WatchOpts, sink chan<- *CsaccountingResumed) (event.Subscription, error) {

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "Resumed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingResumed)
				if err := _Csaccounting.contract.UnpackLog(event, "Resumed", log); err != nil {
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
func (_Csaccounting *CsaccountingFilterer) ParseResumed(log types.Log) (*CsaccountingResumed, error) {
	event := new(CsaccountingResumed)
	if err := _Csaccounting.contract.UnpackLog(event, "Resumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Csaccounting contract.
type CsaccountingRoleAdminChangedIterator struct {
	Event *CsaccountingRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *CsaccountingRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingRoleAdminChanged)
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
		it.Event = new(CsaccountingRoleAdminChanged)
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
func (it *CsaccountingRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingRoleAdminChanged represents a RoleAdminChanged event raised by the Csaccounting contract.
type CsaccountingRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Csaccounting *CsaccountingFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*CsaccountingRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &CsaccountingRoleAdminChangedIterator{contract: _Csaccounting.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Csaccounting *CsaccountingFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *CsaccountingRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingRoleAdminChanged)
				if err := _Csaccounting.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Csaccounting *CsaccountingFilterer) ParseRoleAdminChanged(log types.Log) (*CsaccountingRoleAdminChanged, error) {
	event := new(CsaccountingRoleAdminChanged)
	if err := _Csaccounting.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Csaccounting contract.
type CsaccountingRoleGrantedIterator struct {
	Event *CsaccountingRoleGranted // Event containing the contract specifics and raw log

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
func (it *CsaccountingRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingRoleGranted)
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
		it.Event = new(CsaccountingRoleGranted)
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
func (it *CsaccountingRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingRoleGranted represents a RoleGranted event raised by the Csaccounting contract.
type CsaccountingRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Csaccounting *CsaccountingFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CsaccountingRoleGrantedIterator, error) {

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

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CsaccountingRoleGrantedIterator{contract: _Csaccounting.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Csaccounting *CsaccountingFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *CsaccountingRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingRoleGranted)
				if err := _Csaccounting.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Csaccounting *CsaccountingFilterer) ParseRoleGranted(log types.Log) (*CsaccountingRoleGranted, error) {
	event := new(CsaccountingRoleGranted)
	if err := _Csaccounting.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Csaccounting contract.
type CsaccountingRoleRevokedIterator struct {
	Event *CsaccountingRoleRevoked // Event containing the contract specifics and raw log

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
func (it *CsaccountingRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingRoleRevoked)
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
		it.Event = new(CsaccountingRoleRevoked)
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
func (it *CsaccountingRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingRoleRevoked represents a RoleRevoked event raised by the Csaccounting contract.
type CsaccountingRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Csaccounting *CsaccountingFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CsaccountingRoleRevokedIterator, error) {

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

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CsaccountingRoleRevokedIterator{contract: _Csaccounting.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Csaccounting *CsaccountingFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *CsaccountingRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingRoleRevoked)
				if err := _Csaccounting.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Csaccounting *CsaccountingFilterer) ParseRoleRevoked(log types.Log) (*CsaccountingRoleRevoked, error) {
	event := new(CsaccountingRoleRevoked)
	if err := _Csaccounting.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingStETHSharesRecoveredIterator is returned from FilterStETHSharesRecovered and is used to iterate over the raw logs and unpacked data for StETHSharesRecovered events raised by the Csaccounting contract.
type CsaccountingStETHSharesRecoveredIterator struct {
	Event *CsaccountingStETHSharesRecovered // Event containing the contract specifics and raw log

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
func (it *CsaccountingStETHSharesRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingStETHSharesRecovered)
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
		it.Event = new(CsaccountingStETHSharesRecovered)
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
func (it *CsaccountingStETHSharesRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingStETHSharesRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingStETHSharesRecovered represents a StETHSharesRecovered event raised by the Csaccounting contract.
type CsaccountingStETHSharesRecovered struct {
	Recipient common.Address
	Shares    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStETHSharesRecovered is a free log retrieval operation binding the contract event 0x426e7e0100db57255d4af4a46cd49552ef74f5f002bbdc8d4ebb6371c0070a02.
//
// Solidity: event StETHSharesRecovered(address indexed recipient, uint256 shares)
func (_Csaccounting *CsaccountingFilterer) FilterStETHSharesRecovered(opts *bind.FilterOpts, recipient []common.Address) (*CsaccountingStETHSharesRecoveredIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "StETHSharesRecovered", recipientRule)
	if err != nil {
		return nil, err
	}
	return &CsaccountingStETHSharesRecoveredIterator{contract: _Csaccounting.contract, event: "StETHSharesRecovered", logs: logs, sub: sub}, nil
}

// WatchStETHSharesRecovered is a free log subscription operation binding the contract event 0x426e7e0100db57255d4af4a46cd49552ef74f5f002bbdc8d4ebb6371c0070a02.
//
// Solidity: event StETHSharesRecovered(address indexed recipient, uint256 shares)
func (_Csaccounting *CsaccountingFilterer) WatchStETHSharesRecovered(opts *bind.WatchOpts, sink chan<- *CsaccountingStETHSharesRecovered, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "StETHSharesRecovered", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingStETHSharesRecovered)
				if err := _Csaccounting.contract.UnpackLog(event, "StETHSharesRecovered", log); err != nil {
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
func (_Csaccounting *CsaccountingFilterer) ParseStETHSharesRecovered(log types.Log) (*CsaccountingStETHSharesRecovered, error) {
	event := new(CsaccountingStETHSharesRecovered)
	if err := _Csaccounting.contract.UnpackLog(event, "StETHSharesRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
