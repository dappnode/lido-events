package csmodule

import (
	"context"
	"lido-events/internal/adapters/csModule/bindings"
	"lido-events/internal/application/domain"
	"lido-events/internal/application/ports"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type CsModuleAdapter struct {
	client          *ethclient.Client
	csModuleAddress common.Address
	storageAdapter  ports.StoragePort
	nodeOperatorIds []*big.Int
	mu              sync.RWMutex
}

func NewCsModuleAdapter(
	wsURL string,
	csModuleAddress common.Address,
	storageAdapter ports.StoragePort,
) (*CsModuleAdapter, error) {
	client, err := ethclient.Dial(wsURL)
	if err != nil {
		return nil, err
	}

	initialOperatorIds, err := storageAdapter.GetOperatorIds()
	if err != nil {
		return nil, err
	}

	adapter := &CsModuleAdapter{
		client:          client,
		csModuleAddress: csModuleAddress,
		storageAdapter:  storageAdapter,
		nodeOperatorIds: initialOperatorIds,
	}

	operatorIdUpdates := storageAdapter.RegisterOperatorIdListener()
	go func() {
		for newOperatorIds := range operatorIdUpdates {
			adapter.mu.Lock()
			adapter.nodeOperatorIds = newOperatorIds
			adapter.mu.Unlock()
		}
	}()

	return adapter, nil
}

// WatchCsModuleEvents watches for events emitted by the CsModule contract and calls the appropriate handler functions. Not required to log errors since it will be initialized from main
func (csma *CsModuleAdapter) WatchCsModuleEvents(ctx context.Context, handlers ports.CsModuleHandlers) error {
	csModuleContract, err := bindings.NewCsmodule(csma.csModuleAddress, csma.client)
	if err != nil {
		return err
	}

	// Watch for DepositedSigningKeysCountChanged
	depositedSigningKeysChangedChan := make(chan *domain.CsmoduleDepositedSigningKeysCountChanged)
	subDepositedSigningKeys, err := csModuleContract.WatchDepositedSigningKeysCountChanged(&bind.WatchOpts{Context: ctx}, depositedSigningKeysChangedChan, csma.nodeOperatorIds)
	if err != nil {
		return err
	}

	// Watch for ELRewardsStealingPenaltyReported
	elRewardsStealingPenaltyReportedChan := make(chan *domain.CsmoduleELRewardsStealingPenaltyReported)
	subELRewardsStealingPenaltyReported, err := csModuleContract.WatchELRewardsStealingPenaltyReported(&bind.WatchOpts{Context: ctx}, elRewardsStealingPenaltyReportedChan, csma.nodeOperatorIds)
	if err != nil {
		return err
	}

	// Watch for ELRewardsStealingPenaltySettled
	elRewardsStealingPenaltySettledChan := make(chan *domain.CsmoduleELRewardsStealingPenaltySettled)
	subELRewardsStealingPenaltySettled, err := csModuleContract.WatchELRewardsStealingPenaltySettled(&bind.WatchOpts{Context: ctx}, elRewardsStealingPenaltySettledChan, csma.nodeOperatorIds)
	if err != nil {
		return err
	}

	// Watch for ELRewardsStealingPenaltyCancelled
	elRewardsStealingPenaltyCancelledChan := make(chan *domain.CsmoduleELRewardsStealingPenaltyCancelled)
	subELRewardsStealingPenaltyCancelled, err := csModuleContract.WatchELRewardsStealingPenaltyCancelled(&bind.WatchOpts{Context: ctx}, elRewardsStealingPenaltyCancelledChan, csma.nodeOperatorIds)
	if err != nil {
		return err
	}

	// Watch for InitialSlashingSubmitted
	initialSlashingSubmittedChan := make(chan *domain.CsmoduleInitialSlashingSubmitted)
	subInitialSlashingSubmitted, err := csModuleContract.WatchInitialSlashingSubmitted(&bind.WatchOpts{Context: ctx}, initialSlashingSubmittedChan, csma.nodeOperatorIds)
	if err != nil {
		return err
	}

	// Watch for KeyRemovalChargeApplied
	keyRemovalChargeAppliedChan := make(chan *domain.CsmoduleKeyRemovalChargeApplied)
	subKeyRemovalChargeApplied, err := csModuleContract.WatchKeyRemovalChargeApplied(&bind.WatchOpts{Context: ctx}, keyRemovalChargeAppliedChan, csma.nodeOperatorIds)
	if err != nil {
		return err
	}

	// Watch for NodeOperatorManagerAddressChangeProposed
	nodeOperatorManagerAddressChangeProposedChan := make(chan *domain.CsmoduleNodeOperatorManagerAddressChangeProposed)
	subNodeOperatorManagerAddressChangeProposed, err := csModuleContract.WatchNodeOperatorManagerAddressChangeProposed(&bind.WatchOpts{Context: ctx}, nodeOperatorManagerAddressChangeProposedChan, csma.nodeOperatorIds, []common.Address{}, []common.Address{})
	if err != nil {
		return err
	}

	// Watch for NodeOperatorManagerAddressChanged
	nodeOperatorManagerAddressChangedChan := make(chan *domain.CsmoduleNodeOperatorManagerAddressChanged)
	subNodeOperatorManagerAddressChanged, err := csModuleContract.WatchNodeOperatorManagerAddressChanged(&bind.WatchOpts{Context: ctx}, nodeOperatorManagerAddressChangedChan, csma.nodeOperatorIds, []common.Address{}, []common.Address{})
	if err != nil {
		return err
	}

	// Watch for NodeOperatorRewardAddressChangeProposed
	nodeOperatorRewardAddressChangeProposedChan := make(chan *domain.CsmoduleNodeOperatorRewardAddressChangeProposed)
	subNodeOperatorRewardAddressChangeProposed, err := csModuleContract.WatchNodeOperatorRewardAddressChangeProposed(&bind.WatchOpts{Context: ctx}, nodeOperatorRewardAddressChangeProposedChan, csma.nodeOperatorIds, []common.Address{}, []common.Address{})
	if err != nil {
		return err
	}

	// Watch for NodeOperatorRewardAddressChanged
	nodeOperatorRewardAddressChangedChan := make(chan *domain.CsmoduleNodeOperatorRewardAddressChanged)
	subNodeOperatorRewardAddressChanged, err := csModuleContract.WatchNodeOperatorRewardAddressChanged(&bind.WatchOpts{Context: ctx}, nodeOperatorRewardAddressChangedChan, csma.nodeOperatorIds, []common.Address{}, []common.Address{})
	if err != nil {
		return err
	}

	// Watch for StuckSigningKeysCountChanged
	stuckSigningKeysCountChangedChan := make(chan *domain.CsmoduleStuckSigningKeysCountChanged)
	subStuckSigningKeysCountChanged, err := csModuleContract.WatchStuckSigningKeysCountChanged(&bind.WatchOpts{Context: ctx}, stuckSigningKeysCountChangedChan, csma.nodeOperatorIds)
	if err != nil {
		return err
	}

	// Watch for VettedSigningKeysCountDecreased
	vettedSigningKeysCountDecreasedChan := make(chan *domain.CsmoduleVettedSigningKeysCountDecreased)
	subVettedSigningKeysCountDecreased, err := csModuleContract.WatchVettedSigningKeysCountDecreased(&bind.WatchOpts{Context: ctx}, vettedSigningKeysCountDecreasedChan, csma.nodeOperatorIds)
	if err != nil {
		return err
	}

	// Watch for WithdrawalSubmitted
	withdrawalSubmittedChan := make(chan *domain.CsmoduleWithdrawalSubmitted)
	subWithdrawalSubmitted, err := csModuleContract.WatchWithdrawalSubmitted(&bind.WatchOpts{Context: ctx}, withdrawalSubmittedChan, csma.nodeOperatorIds)
	if err != nil {
		return err
	}

	// Watch for TotalSigningKeysCountChanged
	totalSigningKeysCountChangedChan := make(chan *domain.CsmoduleTotalSigningKeysCountChanged)
	subTotalSigningKeysCountChanged, err := csModuleContract.WatchTotalSigningKeysCountChanged(&bind.WatchOpts{Context: ctx}, totalSigningKeysCountChangedChan, csma.nodeOperatorIds)
	if err != nil {
		return err
	}

	// Watch for PublicRelease
	publicReleaseChan := make(chan *domain.CsmodulePublicRelease)
	subPublicRelease, err := csModuleContract.WatchPublicRelease(&bind.WatchOpts{Context: ctx}, publicReleaseChan)
	if err != nil {
		return err
	}

	go func() {
		defer subDepositedSigningKeys.Unsubscribe()
		defer subELRewardsStealingPenaltyReported.Unsubscribe()
		defer subELRewardsStealingPenaltySettled.Unsubscribe()
		defer subELRewardsStealingPenaltyCancelled.Unsubscribe()
		defer subInitialSlashingSubmitted.Unsubscribe()
		defer subKeyRemovalChargeApplied.Unsubscribe()
		defer subNodeOperatorManagerAddressChangeProposed.Unsubscribe()
		defer subNodeOperatorManagerAddressChanged.Unsubscribe()
		defer subNodeOperatorRewardAddressChangeProposed.Unsubscribe()
		defer subNodeOperatorRewardAddressChanged.Unsubscribe()
		defer subStuckSigningKeysCountChanged.Unsubscribe()
		defer subVettedSigningKeysCountDecreased.Unsubscribe()
		defer subWithdrawalSubmitted.Unsubscribe()
		defer subTotalSigningKeysCountChanged.Unsubscribe()
		defer subPublicRelease.Unsubscribe()

		for {
			select {
			case event := <-depositedSigningKeysChangedChan:
				handlers.HandleDepositedSigningKeysCountChanged(event)
				return
			case event := <-elRewardsStealingPenaltyReportedChan:
				handlers.HandleElRewardsStealingPenaltyReported(event)
				return
			case event := <-elRewardsStealingPenaltySettledChan:
				handlers.HandleElRewardsStealingPenaltySettled(event)
				return
			case event := <-elRewardsStealingPenaltyCancelledChan:
				handlers.HandleElRewardsStealingPenaltyCancelled(event)
				return
			case event := <-initialSlashingSubmittedChan:
				handlers.HandleInitialSlashingSubmitted(event)
				return
			case event := <-keyRemovalChargeAppliedChan:
				handlers.HandleKeyRemovalChargeApplied(event)
				return
			case event := <-nodeOperatorManagerAddressChangeProposedChan:
				handlers.HandleNodeOperatorManagerAddressChangeProposed(event)
				return
			case event := <-nodeOperatorManagerAddressChangedChan:
				handlers.HandleNodeOperatorManagerAddressChanged(event)
				return
			case event := <-nodeOperatorRewardAddressChangeProposedChan:
				handlers.HandleNodeOperatorRewardAddressChangeProposed(event)
				return
			case event := <-nodeOperatorRewardAddressChangedChan:
				handlers.HandleNodeOperatorRewardAddressChanged(event)
				return
			case event := <-stuckSigningKeysCountChangedChan:
				handlers.HandleStuckSigningKeysCountChanged(event)
				return
			case event := <-vettedSigningKeysCountDecreasedChan:
				handlers.HandleVettedSigningKeysCountDecreased(event)
				return
			case event := <-withdrawalSubmittedChan:
				handlers.HandleWithdrawalSubmitted(event)
				return
			case event := <-totalSigningKeysCountChangedChan:
				handlers.HandleTotalSigningKeysCountChanged(event)
				return
			case event := <-publicReleaseChan:
				handlers.HandlePublicRelease(event)
				return

			case <-ctx.Done():
				return
			}
		}
	}()

	// Add additional channels and watch setups for other CSModule events
	return nil
}
