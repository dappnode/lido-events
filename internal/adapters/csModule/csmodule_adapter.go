package csmodule

import (
	"context"
	"lido-events/internal/adapters/csModule/bindings"
	"lido-events/internal/aplication/domain"
	"lido-events/internal/aplication/ports"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type CsModuleAdapter struct {
	client             *ethclient.Client
	CsModuleAddress    common.Address
	NodeOperatorId     []*big.Int
	OldAddress         []common.Address // TODO: where to get this
	NewAddress         []common.Address // TODO: where to get this
	OldProposedAddress []common.Address // TODO: where to get this
	NewProposedAddress []common.Address // TODO: where to get this
}

func NewCsModuleAdapter(
	wsURL string,
	csModuleAddress common.Address,
	nodeOperatorId []*big.Int,
	oldAddress []common.Address,
	newAddress []common.Address,
	oldProposedAddress []common.Address,
	newProposedAddress []common.Address,
) (*CsModuleAdapter, error) {
	client, err := ethclient.Dial(wsURL)
	if err != nil {
		return nil, err
	}

	return &CsModuleAdapter{
		client:             client,
		CsModuleAddress:    csModuleAddress,
		NodeOperatorId:     nodeOperatorId,
		OldAddress:         oldAddress,
		NewAddress:         newAddress,
		OldProposedAddress: oldProposedAddress,
		NewProposedAddress: newProposedAddress,
	}, nil
}

func (csma *CsModuleAdapter) WatchCsModuleEvents(ctx context.Context, handlers ports.CsModuleHandlers) error {
	csModuleContract, err := bindings.NewCsmodule(csma.CsModuleAddress, csma.client)
	if err != nil {
		return err
	}

	// Watch for DepositedSigningKeysCountChanged
	depositedSigningKeysChangedChan := make(chan *domain.CsmoduleDepositedSigningKeysCountChanged)
	sub, err := csModuleContract.WatchDepositedSigningKeysCountChanged(&bind.WatchOpts{Context: ctx}, depositedSigningKeysChangedChan, csma.NodeOperatorId)
	if err != nil {
		return err
	}

	// Watch for ELRewardsStealingPenaltyReported
	elRewardsStealingPenaltyReportedChan := make(chan *domain.CsmoduleELRewardsStealingPenaltyReported)
	subELRewardsStealingPenaltyReported, err := csModuleContract.WatchELRewardsStealingPenaltyReported(&bind.WatchOpts{Context: ctx}, elRewardsStealingPenaltyReportedChan, csma.NodeOperatorId)
	if err != nil {
		return err
	}

	// Watch for ELRewardsStealingPenaltySettled
	elRewardsStealingPenaltySettledChan := make(chan *domain.CsmoduleELRewardsStealingPenaltySettled)
	subELRewardsStealingPenaltySettled, err := csModuleContract.WatchELRewardsStealingPenaltySettled(&bind.WatchOpts{Context: ctx}, elRewardsStealingPenaltySettledChan, csma.NodeOperatorId)
	if err != nil {
		return err
	}

	// Watch for ELRewardsStealingPenaltyCancelled
	elRewardsStealingPenaltyCancelledChan := make(chan *domain.CsmoduleELRewardsStealingPenaltyCancelled)
	subELRewardsStealingPenaltyCancelled, err := csModuleContract.WatchELRewardsStealingPenaltyCancelled(&bind.WatchOpts{Context: ctx}, elRewardsStealingPenaltyCancelledChan, csma.NodeOperatorId)
	if err != nil {
		return err
	}

	// Watch for InitialSlashingSubmitted
	initialSlashingSubmittedChan := make(chan *domain.CsmoduleInitialSlashingSubmitted)
	subInitialSlashingSubmitted, err := csModuleContract.WatchInitialSlashingSubmitted(&bind.WatchOpts{Context: ctx}, initialSlashingSubmittedChan, csma.NodeOperatorId)
	if err != nil {
		return err
	}

	// Watch for KeyRemovalChargeApplied
	keyRemovalChargeAppliedChan := make(chan *domain.CsmoduleKeyRemovalChargeApplied)
	subKeyRemovalChargeApplied, err := csModuleContract.WatchKeyRemovalChargeApplied(&bind.WatchOpts{Context: ctx}, keyRemovalChargeAppliedChan, csma.NodeOperatorId)
	if err != nil {
		return err
	}

	// Watch for NodeOperatorManagerAddressChangeProposed
	nodeOperatorManagerAddressChangeProposedChan := make(chan *domain.CsmoduleNodeOperatorManagerAddressChangeProposed)
	subNodeOperatorManagerAddressChangeProposed, err := csModuleContract.WatchNodeOperatorManagerAddressChangeProposed(&bind.WatchOpts{Context: ctx}, nodeOperatorManagerAddressChangeProposedChan, csma.NodeOperatorId, csma.OldProposedAddress, csma.NewProposedAddress)
	if err != nil {
		return err
	}

	// Watch for NodeOperatorManagerAddressChanged
	nodeOperatorManagerAddressChangedChan := make(chan *domain.CsmoduleNodeOperatorManagerAddressChanged)
	subNodeOperatorManagerAddressChanged, err := csModuleContract.WatchNodeOperatorManagerAddressChanged(&bind.WatchOpts{Context: ctx}, nodeOperatorManagerAddressChangedChan, csma.NodeOperatorId, csma.OldAddress, csma.NewAddress)
	if err != nil {
		return err
	}

	// Watch for NodeOperatorRewardAddressChangeProposed
	nodeOperatorRewardAddressChangeProposedChan := make(chan *domain.CsmoduleNodeOperatorRewardAddressChangeProposed)
	subNodeOperatorRewardAddressChangeProposed, err := csModuleContract.WatchNodeOperatorRewardAddressChangeProposed(&bind.WatchOpts{Context: ctx}, nodeOperatorRewardAddressChangeProposedChan, csma.NodeOperatorId, csma.OldProposedAddress, csma.NewProposedAddress)
	if err != nil {
		return err
	}

	// Watch for NodeOperatorRewardAddressChanged
	nodeOperatorRewardAddressChangedChan := make(chan *domain.CsmoduleNodeOperatorRewardAddressChanged)
	subNodeOperatorRewardAddressChanged, err := csModuleContract.WatchNodeOperatorRewardAddressChanged(&bind.WatchOpts{Context: ctx}, nodeOperatorRewardAddressChangedChan, csma.NodeOperatorId, csma.OldAddress, csma.NewAddress)
	if err != nil {
		return err
	}

	// Watch for StuckSigningKeysCountChanged
	stuckSigningKeysCountChangedChan := make(chan *domain.CsmoduleStuckSigningKeysCountChanged)
	subStuckSigningKeysCountChanged, err := csModuleContract.WatchStuckSigningKeysCountChanged(&bind.WatchOpts{Context: ctx}, stuckSigningKeysCountChangedChan, csma.NodeOperatorId)
	if err != nil {
		return err
	}

	// Watch for VettedSigningKeysCountDecreased
	vettedSigningKeysCountDecreasedChan := make(chan *domain.CsmoduleVettedSigningKeysCountDecreased)
	subVettedSigningKeysCountDecreased, err := csModuleContract.WatchVettedSigningKeysCountDecreased(&bind.WatchOpts{Context: ctx}, vettedSigningKeysCountDecreasedChan, csma.NodeOperatorId)
	if err != nil {
		return err
	}

	// Watch for WithdrawalSubmitted
	withdrawalSubmittedChan := make(chan *domain.CsmoduleWithdrawalSubmitted)
	subWithdrawalSubmitted, err := csModuleContract.WatchWithdrawalSubmitted(&bind.WatchOpts{Context: ctx}, withdrawalSubmittedChan, csma.NodeOperatorId)
	if err != nil {
		return err
	}

	// Watch for TotalSigningKeysCountChanged
	totalSigningKeysCountChangedChan := make(chan *domain.CsmoduleTotalSigningKeysCountChanged)
	subTotalSigningKeysCountChanged, err := csModuleContract.WatchTotalSigningKeysCountChanged(&bind.WatchOpts{Context: ctx}, totalSigningKeysCountChangedChan, csma.NodeOperatorId)
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
		defer sub.Unsubscribe()
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
				if err := handlers.HandleDepositedSigningKeysCountChanged(event); err != nil {
					log.Printf("Error handling DepositedSigningKeysCountChanged: %v", err)
				}
			case event := <-elRewardsStealingPenaltyReportedChan:
				if err := handlers.HandleElRewardsStealingPenaltyReported(event); err != nil {
					log.Printf("Error handling ELRewardsStealingPenaltyReported: %v", err)
				}
			case event := <-elRewardsStealingPenaltySettledChan:
				if err := handlers.HandleElRewardsStealingPenaltySettled(event); err != nil {
					log.Printf("Error handling ELRewardsStealingPenaltySettled: %v", err)
				}
			case event := <-elRewardsStealingPenaltyCancelledChan:
				if err := handlers.HandleElRewardsStealingPenaltyCancelled(event); err != nil {
					log.Printf("Error handling ELRewardsStealingPenaltyCancelled: %v", err)
				}
			case event := <-initialSlashingSubmittedChan:
				if err := handlers.HandleInitialSlashingSubmitted(event); err != nil {
					log.Printf("Error handling InitialSlashingSubmitted: %v", err)
				}
			case event := <-keyRemovalChargeAppliedChan:
				if err := handlers.HandleKeyRemovalChargeApplied(event); err != nil {
					log.Printf("Error handling KeyRemovalChargeApplied: %v", err)
				}
			case event := <-nodeOperatorManagerAddressChangeProposedChan:
				if err := handlers.HandleNodeOperatorManagerAddressChangeProposed(event); err != nil {
					log.Printf("Error handling NodeOperatorManagerAddressChangeProposed: %v", err)
				}
			case event := <-nodeOperatorManagerAddressChangedChan:
				if err := handlers.HandleNodeOperatorManagerAddressChanged(event); err != nil {
					log.Printf("Error handling NodeOperatorManagerAddressChanged: %v", err)
				}
			case event := <-nodeOperatorRewardAddressChangeProposedChan:
				if err := handlers.HandleNodeOperatorRewardAddressChangeProposed(event); err != nil {
					log.Printf("Error handling NodeOperatorRewardAddressChangeProposed: %v", err)
				}
			case event := <-nodeOperatorRewardAddressChangedChan:
				if err := handlers.HandleNodeOperatorRewardAddressChanged(event); err != nil {
					log.Printf("Error handling NodeOperatorRewardAddressChanged: %v", err)
				}
			case event := <-stuckSigningKeysCountChangedChan:
				if err := handlers.HandleStuckSigningKeysCountChanged(event); err != nil {
					log.Printf("Error handling StuckSigningKeysCountChanged: %v", err)
				}
			case event := <-vettedSigningKeysCountDecreasedChan:
				if err := handlers.HandleVettedSigningKeysCountDecreased(event); err != nil {
					log.Printf("Error handling VettedSigningKeysCountDecreased: %v", err)
				}
			case event := <-withdrawalSubmittedChan:
				if err := handlers.HandleWithdrawalSubmitted(event); err != nil {
					log.Printf("Error handling WithdrawalSubmitted: %v", err)
				}
			case event := <-totalSigningKeysCountChangedChan:
				if err := handlers.HandleTotalSigningKeysCountChanged(event); err != nil {
					log.Printf("Error handling TotalSigningKeysCountChanged: %v", err)
				}
			case event := <-publicReleaseChan:
				if err := handlers.HandlePublicRelease(event); err != nil {
					log.Printf("Error handling PublicRelease: %v", err)
				}
			case err := <-sub.Err():
				log.Printf("Subscription error: %v", err)
				return
			case <-ctx.Done():
				return
			}
		}
	}()

	// Add additional channels and watch setups for other CSModule events
	return nil
}
