package csmodule

import (
	"context"
	"fmt"
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
	client            *ethclient.Client
	csModuleAddress   common.Address
	storageAdapter    ports.StoragePort
	nodeOperatorIds   []*big.Int
	mu                sync.RWMutex
	resubscribeSignal chan struct{}
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
		client:            client,
		csModuleAddress:   csModuleAddress,
		storageAdapter:    storageAdapter,
		nodeOperatorIds:   initialOperatorIds,
		resubscribeSignal: make(chan struct{}, 1),
	}

	operatorIdUpdates := storageAdapter.RegisterOperatorIdListener()
	go adapter.handleOperatorIdUpdates(operatorIdUpdates)

	return adapter, nil
}

func (csma *CsModuleAdapter) handleOperatorIdUpdates(updates <-chan []*big.Int) {
	for newOperatorIds := range updates {
		csma.mu.Lock()
		csma.nodeOperatorIds = newOperatorIds
		csma.mu.Unlock()

		// Notify the resubscribe signal channel
		select {
		case csma.resubscribeSignal <- struct{}{}:
		default: // Prevent blocking if there's already a pending signal
		}
	}
}

// Expose a method to retrieve the resubscribe signal channel
func (csma *CsModuleAdapter) ResubscribeSignal() <-chan struct{} {
	return csma.resubscribeSignal
}

// ScanNodeOperatorEvents scans the CsModule contract for NodeOperator events. The CsModule events tracked are: NodeOperatorAdded, NodeOperatorManagerAddressChanged, NodeOperatorRewardAddressChanged
func (csma *CsModuleAdapter) ScanNodeOperatorEvents(ctx context.Context, start uint64, end *uint64, handleNodeOperatorAddedEvent func(*domain.CsmoduleNodeOperatorAdded) error, handleNodeOperatorManagerAddressChangedEvent func(*domain.CsmoduleNodeOperatorManagerAddressChanged) error, handleNodeOperatorRewardAddressChangedEvent func(*domain.CsmoduleNodeOperatorRewardAddressChanged) error) error {
	csModuleContract, err := bindings.NewCsmodule(csma.csModuleAddress, csma.client)
	if err != nil {
		return fmt.Errorf("failed to create CsModule contract instance: %w", err)
	}

	// Get operator ids
	operatorIds, err := csma.storageAdapter.GetOperatorIds()
	if err != nil {
		return fmt.Errorf("failed to retrieve operator IDs from storage: %w", err)
	}

	// Filter for NodeOperatorAdded events
	nodeOperatorAddedEvents, err := csModuleContract.FilterNodeOperatorAdded(&bind.FilterOpts{Context: ctx, Start: start, End: end}, operatorIds, []common.Address{}, []common.Address{})
	if err != nil {
		return fmt.Errorf("failed to filter NodeOperatorAdded events for block range %d to %v: %w", start, *end, err)
	}
	for nodeOperatorAddedEvents.Next() {
		if err := nodeOperatorAddedEvents.Error(); err != nil {
			return err
		}

		if err := handleNodeOperatorAddedEvent(nodeOperatorAddedEvents.Event); err != nil {
			return err
		}
	}

	// Filter for NodeOperatorManagerAddressChanged events
	nodeOperatorManagerAddressChangedEvents, err := csModuleContract.FilterNodeOperatorManagerAddressChanged(&bind.FilterOpts{Context: ctx, Start: start, End: end}, operatorIds, []common.Address{}, []common.Address{})
	if err != nil {
		return fmt.Errorf("failed to filter NodeOperatorManagerAddressChanged events for block range %d to %v: %w", start, *end, err)
	}
	for nodeOperatorManagerAddressChangedEvents.Next() {
		if err := nodeOperatorManagerAddressChangedEvents.Error(); err != nil {
			return err
		}

		if err := handleNodeOperatorManagerAddressChangedEvent(nodeOperatorManagerAddressChangedEvents.Event); err != nil {
			return err
		}
	}

	// Filter for NodeOperatorRewardAddressChanged events
	nodeOperatorRewardAddressChangedEvents, err := csModuleContract.FilterNodeOperatorRewardAddressChanged(&bind.FilterOpts{Context: ctx, Start: start, End: end}, operatorIds, []common.Address{}, []common.Address{})
	if err != nil {
		return fmt.Errorf("failed to filter NodeOperatorRewardAddressChanged events for block range %d to %v: %w", start, *end, err)
	}
	for nodeOperatorRewardAddressChangedEvents.Next() {
		if err := nodeOperatorRewardAddressChangedEvents.Error(); err != nil {
			return err
		}

		if err := handleNodeOperatorRewardAddressChangedEvent(nodeOperatorRewardAddressChangedEvents.Event); err != nil {
			return err
		}
	}

	return nil
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
			case event := <-elRewardsStealingPenaltyReportedChan:
				handlers.HandleElRewardsStealingPenaltyReported(event)
			case event := <-elRewardsStealingPenaltySettledChan:
				handlers.HandleElRewardsStealingPenaltySettled(event)
			case event := <-elRewardsStealingPenaltyCancelledChan:
				handlers.HandleElRewardsStealingPenaltyCancelled(event)
			case event := <-initialSlashingSubmittedChan:
				handlers.HandleInitialSlashingSubmitted(event)
			case event := <-keyRemovalChargeAppliedChan:
				handlers.HandleKeyRemovalChargeApplied(event)
			case event := <-nodeOperatorManagerAddressChangeProposedChan:
				handlers.HandleNodeOperatorManagerAddressChangeProposed(event)
			case event := <-nodeOperatorManagerAddressChangedChan:
				handlers.HandleNodeOperatorManagerAddressChanged(event)
			case event := <-nodeOperatorRewardAddressChangeProposedChan:
				handlers.HandleNodeOperatorRewardAddressChangeProposed(event)
			case event := <-nodeOperatorRewardAddressChangedChan:
				handlers.HandleNodeOperatorRewardAddressChanged(event)
			case event := <-stuckSigningKeysCountChangedChan:
				handlers.HandleStuckSigningKeysCountChanged(event)
			case event := <-vettedSigningKeysCountDecreasedChan:
				handlers.HandleVettedSigningKeysCountDecreased(event)
			case event := <-withdrawalSubmittedChan:
				handlers.HandleWithdrawalSubmitted(event)
			case event := <-totalSigningKeysCountChangedChan:
				handlers.HandleTotalSigningKeysCountChanged(event)
			case event := <-publicReleaseChan:
				handlers.HandlePublicRelease(event)

			case <-ctx.Done():
				return
			}
		}
	}()

	// Add additional channels and watch setups for other CSModule events
	return nil
}
