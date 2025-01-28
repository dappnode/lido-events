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
	websocketClient   *ethclient.Client
	rpcClient         *ethclient.Client
	csModuleAddress   common.Address
	storageAdapter    ports.StoragePort
	nodeOperatorIds   []*big.Int
	mu                sync.RWMutex
	resubscribeSignal chan struct{}
	blockChunkSize    uint64
}

func NewCsModuleAdapter(
	websocketClient *ethclient.Client,
	rpcClient *ethclient.Client,
	csModuleAddress common.Address,
	storageAdapter ports.StoragePort,
	blockChunkSize uint64,
) (*CsModuleAdapter, error) {
	initialOperatorIds, err := storageAdapter.GetOperatorIds()
	if err != nil {
		return nil, err
	}

	adapter := &CsModuleAdapter{
		websocketClient:   websocketClient,
		rpcClient:         rpcClient,
		csModuleAddress:   csModuleAddress,
		storageAdapter:    storageAdapter,
		nodeOperatorIds:   initialOperatorIds,
		resubscribeSignal: make(chan struct{}, 1),
		blockChunkSize:    blockChunkSize,
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

// ScanNodeOperatorEvents scans the CsModule contract for NodeOperator events in chunks of blocks.
func (csma *CsModuleAdapter) ScanNodeOperatorEvents(
	ctx context.Context,
	address common.Address,
	start uint64,
	end *uint64,
	handleNodeOperatorAddedEvent func(*domain.CsmoduleNodeOperatorAdded, common.Address) error,
	handleNodeOperatorManagerAddressChangedEvent func(*domain.CsmoduleNodeOperatorManagerAddressChanged, common.Address) error,
	handleNodeOperatorRewardAddressChangedEvent func(*domain.CsmoduleNodeOperatorRewardAddressChanged, common.Address) error,
	handleNodeOperatorRewardAddressChangeProposedEvent func(*domain.CsmoduleNodeOperatorRewardAddressChangeProposed, common.Address) error,
	handleNodeOperatorManagerAddressChangeProposedEvent func(*domain.CsmoduleNodeOperatorManagerAddressChangeProposed, common.Address) error,
) error {
	if end == nil {
		return fmt.Errorf("end block cannot be nil")
	}

	csModuleContract, err := bindings.NewCsmodule(csma.csModuleAddress, csma.rpcClient)
	if err != nil {
		return fmt.Errorf("failed to create CsModule contract instance: %w", err)
	}

	// Helper function to scan a chunk
	scanChunk := func(chunkStart, chunkEnd uint64) error {
		// Filter for NodeOperatorAdded events with the address as manager
		operatorAddedManagerEvents, err := csModuleContract.FilterNodeOperatorAdded(
			&bind.FilterOpts{Context: ctx, Start: chunkStart, End: &chunkEnd},
			[]*big.Int{},              // No filter for `nodeOperatorId`
			[]common.Address{address}, // Address as manager
			[]common.Address{},        // No filter for rewardAddress
		)
		if err != nil {
			return fmt.Errorf("failed to filter NodeOperatorAdded (manager) events for block range %d to %d: %w", chunkStart, chunkEnd, err)
		}
		for operatorAddedManagerEvents.Next() {
			if err := operatorAddedManagerEvents.Error(); err != nil {
				return err
			}
			if err := handleNodeOperatorAddedEvent(operatorAddedManagerEvents.Event, address); err != nil {
				return err
			}
		}

		// Filter for NodeOperatorAdded events with the address as reward
		operatorAddedRewardEvents, err := csModuleContract.FilterNodeOperatorAdded(
			&bind.FilterOpts{Context: ctx, Start: chunkStart, End: &chunkEnd},
			[]*big.Int{},              // No filter for `nodeOperatorId`
			[]common.Address{},        // No filter for managerAddress
			[]common.Address{address}, // Address as reward
		)
		if err != nil {
			return fmt.Errorf("failed to filter NodeOperatorAdded (reward) events for block range %d to %d: %w", chunkStart, chunkEnd, err)
		}
		for operatorAddedRewardEvents.Next() {
			if err := operatorAddedRewardEvents.Error(); err != nil {
				return err
			}
			if err := handleNodeOperatorAddedEvent(operatorAddedRewardEvents.Event, address); err != nil {
				return err
			}
		}

		// Filter for NodeOperatorManagerAddressChanged events with the address as manager
		operatorManagerChangedManagerEvents, err := csModuleContract.FilterNodeOperatorManagerAddressChanged(
			&bind.FilterOpts{Context: ctx, Start: chunkStart, End: &chunkEnd},
			[]*big.Int{},              // No filter for `nodeOperatorId`
			[]common.Address{address}, // Address as manager
			[]common.Address{},        // No filter for rewardAddress
		)
		if err != nil {
			return fmt.Errorf("failed to filter NodeOperatorManagerAddressChanged (manager) events for block range %d to %d: %w", chunkStart, chunkEnd, err)
		}
		for operatorManagerChangedManagerEvents.Next() {
			if err := operatorManagerChangedManagerEvents.Error(); err != nil {
				return err
			}
			if err := handleNodeOperatorManagerAddressChangedEvent(operatorManagerChangedManagerEvents.Event, address); err != nil {
				return err
			}
		}

		// Filter for NodeOperatorManagerAddressChanged events with the address as reward
		operatorManagerChangedRewardEvents, err := csModuleContract.FilterNodeOperatorManagerAddressChanged(
			&bind.FilterOpts{Context: ctx, Start: chunkStart, End: &chunkEnd},
			[]*big.Int{},              // No filter for `nodeOperatorId`
			[]common.Address{},        // No filter for managerAddress
			[]common.Address{address}, // Address as reward
		)
		if err != nil {
			return fmt.Errorf("failed to filter NodeOperatorManagerAddressChanged (reward) events for block range %d to %d: %w", chunkStart, chunkEnd, err)
		}
		for operatorManagerChangedRewardEvents.Next() {
			if err := operatorManagerChangedRewardEvents.Error(); err != nil {
				return err
			}
			if err := handleNodeOperatorManagerAddressChangedEvent(operatorManagerChangedRewardEvents.Event, address); err != nil {
				return err
			}
		}

		// Filter for NodeOperatorRewardAddressChanged events with the address as manager
		operatorRewardChangedManagerEvents, err := csModuleContract.FilterNodeOperatorRewardAddressChanged(
			&bind.FilterOpts{Context: ctx, Start: chunkStart, End: &chunkEnd},
			[]*big.Int{},              // No filter for `nodeOperatorId`
			[]common.Address{address}, // Address as manager
			[]common.Address{},        // No filter for rewardAddress
		)
		if err != nil {
			return fmt.Errorf("failed to filter NodeOperatorRewardAddressChanged (manager) events for block range %d to %d: %w", chunkStart, chunkEnd, err)
		}
		for operatorRewardChangedManagerEvents.Next() {
			if err := operatorRewardChangedManagerEvents.Error(); err != nil {
				return err
			}
			if err := handleNodeOperatorRewardAddressChangedEvent(operatorRewardChangedManagerEvents.Event, address); err != nil {
				return err
			}
		}

		// Filter for NodeOperatorRewardAddressChanged events with the address as reward
		operatorRewardChangedRewardEvents, err := csModuleContract.FilterNodeOperatorRewardAddressChanged(
			&bind.FilterOpts{Context: ctx, Start: chunkStart, End: &chunkEnd},
			[]*big.Int{},              // No filter for `nodeOperatorId`
			[]common.Address{},        // No filter for managerAddress
			[]common.Address{address}, // Address as reward
		)
		if err != nil {
			return fmt.Errorf("failed to filter NodeOperatorRewardAddressChanged (reward) events for block range %d to %d: %w", chunkStart, chunkEnd, err)
		}
		for operatorRewardChangedRewardEvents.Next() {
			if err := operatorRewardChangedRewardEvents.Error(); err != nil {
				return err
			}
			if err := handleNodeOperatorRewardAddressChangedEvent(operatorRewardChangedRewardEvents.Event, address); err != nil {
				return err
			}
		}

		// Filter NodeOperatorRewardAddressChangeProposed events with the address as old address
		operatorRewardChangeProposedManagerEvents, err := csModuleContract.FilterNodeOperatorRewardAddressChangeProposed(
			&bind.FilterOpts{Context: ctx, Start: chunkStart, End: &chunkEnd},
			[]*big.Int{},              // No filter for `nodeOperatorId`
			[]common.Address{address}, // Address as old address
			[]common.Address{},        // No filter for rewardAddress
		)
		if err != nil {
			return fmt.Errorf("failed to filter NodeOperatorRewardAddressChangeProposed (manager) events for block range %d to %d: %w", chunkStart, chunkEnd, err)
		}
		for operatorRewardChangeProposedManagerEvents.Next() {
			if err := operatorRewardChangeProposedManagerEvents.Error(); err != nil {
				return err
			}
			if err := handleNodeOperatorRewardAddressChangeProposedEvent(operatorRewardChangeProposedManagerEvents.Event, address); err != nil {
				return err
			}
		}

		// Filter NodeOperatorRewardAddressChangeProposed events with the address as new address
		operatorRewardChangeProposedRewardEvents, err := csModuleContract.FilterNodeOperatorRewardAddressChangeProposed(
			&bind.FilterOpts{Context: ctx, Start: chunkStart, End: &chunkEnd},
			[]*big.Int{},              // No filter for `nodeOperatorId`
			[]common.Address{},        // No filter for managerAddress
			[]common.Address{address}, // Address as new address
		)
		if err != nil {
			return fmt.Errorf("failed to filter NodeOperatorRewardAddressChangeProposed (reward) events for block range %d to %d: %w", chunkStart, chunkEnd, err)
		}
		for operatorRewardChangeProposedRewardEvents.Next() {
			if err := operatorRewardChangeProposedRewardEvents.Error(); err != nil {
				return err
			}
			if err := handleNodeOperatorRewardAddressChangeProposedEvent(operatorRewardChangeProposedRewardEvents.Event, address); err != nil {
				return err
			}
		}

		// Filter NodeOperatorManagerAddressChangeProposed events with the address as old address
		operatorManagerChangeProposedManagerEvents, err := csModuleContract.FilterNodeOperatorManagerAddressChangeProposed(
			&bind.FilterOpts{Context: ctx, Start: chunkStart, End: &chunkEnd},
			[]*big.Int{},              // No filter for `nodeOperatorId`
			[]common.Address{address}, // Address as old address
			[]common.Address{},        // No filter for rewardAddress
		)
		if err != nil {
			return fmt.Errorf("failed to filter NodeOperatorManagerAddressChangeProposed (manager) events for block range %d to %d: %w", chunkStart, chunkEnd, err)
		}
		for operatorManagerChangeProposedManagerEvents.Next() {
			if err := operatorManagerChangeProposedManagerEvents.Error(); err != nil {
				return err
			}
			// Handle the event
		}

		// Filter NodeOperatorManagerAddressChangeProposed events with the address as new address
		operatorManagerChangeProposedRewardEvents, err := csModuleContract.FilterNodeOperatorManagerAddressChangeProposed(
			&bind.FilterOpts{Context: ctx, Start: chunkStart, End: &chunkEnd},
			[]*big.Int{},              // No filter for `nodeOperatorId`
			[]common.Address{},        // No filter for managerAddress
			[]common.Address{address}, // Address as new address
		)
		if err != nil {
			return fmt.Errorf("failed to filter NodeOperatorManagerAddressChangeProposed (reward) events for block range %d to %d: %w", chunkStart, chunkEnd, err)
		}
		for operatorManagerChangeProposedRewardEvents.Next() {
			if err := operatorManagerChangeProposedRewardEvents.Error(); err != nil {
				return err
			}
			// Handle the event
		}

		return nil
	}

	// Iterate through block ranges in chunks
	for currentStart := start; currentStart <= *end; currentStart += csma.blockChunkSize {
		// Determine the end of the current chunk
		currentEnd := currentStart + csma.blockChunkSize - 1
		if currentEnd > *end {
			currentEnd = *end
		}

		// Scan the current chunk
		if err := scanChunk(currentStart, currentEnd); err != nil {
			return fmt.Errorf("error scanning block range %d to %d: %w", currentStart, currentEnd, err)
		}
	}

	return nil
}

// WatchCsModuleEvents watches for events emitted by the CsModule contract and calls the appropriate handler functions. Not required to log errors since it will be initialized from main
func (csma *CsModuleAdapter) WatchCsModuleEvents(ctx context.Context, handlers ports.CsModuleHandlers) error {
	csModuleContract, err := bindings.NewCsmodule(csma.csModuleAddress, csma.websocketClient)
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
