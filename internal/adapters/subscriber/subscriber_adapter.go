package subscriber

import (
	"context"
	"log"
	"math/big"

	"lido-events/internal/adapters/subscriber/bindings/csfeedistributor"
	"lido-events/internal/adapters/subscriber/bindings/csmodule"
	"lido-events/internal/adapters/subscriber/bindings/vebo"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type SubscriberAdapter struct {
	client                  *ethclient.Client
	CsFeeDistributorAddress common.Address
	CsModuleAddress         common.Address
	VeboAddress             common.Address
	StakingModuleId         []*big.Int
	NodeOperatorId          []*big.Int
	ValidatorIndex          []*big.Int // TODO: where to get it?
	RefSlot                 []*big.Int // TODO: where to get it?
}

func NewSubscriberAdapter(
	wsURL string,
	csFeeDistributorAddress, csModuleAddress, veboAddress common.Address,
) (*SubscriberAdapter, error) {
	client, err := ethclient.Dial(wsURL)
	if err != nil {
		return nil, err
	}

	return &SubscriberAdapter{
		client:                  client,
		CsFeeDistributorAddress: csFeeDistributorAddress,
		CsModuleAddress:         csModuleAddress,
		VeboAddress:             veboAddress,
	}, nil
}

func (sa *SubscriberAdapter) SubscribeToEvents(ctx context.Context, handleLog func(logData interface{}) error) error {
	if err := sa.watchCsModuleEvents(ctx, handleLog); err != nil {
		return err
	}
	if err := sa.watchVeboEvents(ctx, handleLog); err != nil {
		return err
	}
	if err := sa.watchCsFeeDistributorEvents(ctx, handleLog); err != nil {
		return err
	}
	return nil
}

func (sa *SubscriberAdapter) watchCsModuleEvents(ctx context.Context, handleLog func(logData interface{}) error) error {
	csModuleContract, err := csmodule.NewCsmodule(sa.CsModuleAddress, sa.client)
	if err != nil {
		return err
	}

	// Watch for DepositedSigningKeysCountChanged
	depositedSigningKeysChangedChan := make(chan *csmodule.CsmoduleDepositedSigningKeysCountChanged)
	sub, err := csModuleContract.WatchDepositedSigningKeysCountChanged(&bind.WatchOpts{Context: ctx}, depositedSigningKeysChangedChan, sa.NodeOperatorId)
	if err != nil {
		return err
	}

	// Watch for ELRewardsStealingPenaltyReported
	elRewardsStealingPenaltyReportedChan := make(chan *csmodule.CsmoduleELRewardsStealingPenaltyReported)
	subELRewardsStealingPenaltyReported, err := csModuleContract.WatchELRewardsStealingPenaltyReported(&bind.WatchOpts{Context: ctx}, elRewardsStealingPenaltyReportedChan, sa.NodeOperatorId)
	if err != nil {
		return err
	}

	// Watch for ELRewardsStealingPenaltySettled
	elRewardsStealingPenaltySettledChan := make(chan *csmodule.CsmoduleELRewardsStealingPenaltySettled)
	subELRewardsStealingPenaltySettled, err := csModuleContract.WatchELRewardsStealingPenaltySettled(&bind.WatchOpts{Context: ctx}, elRewardsStealingPenaltySettledChan, sa.NodeOperatorId)
	if err != nil {
		return err
	}

	// Watch for ELRewardsStealingPenaltyCancelled
	elRewardsStealingPenaltyCancelledChan := make(chan *csmodule.CsmoduleELRewardsStealingPenaltyCancelled)
	subELRewardsStealingPenaltyCancelled, err := csModuleContract.WatchELRewardsStealingPenaltyCancelled(&bind.WatchOpts{Context: ctx}, elRewardsStealingPenaltyCancelledChan, sa.NodeOperatorId)
	if err != nil {
		return err
	}

	// Watch for InitialSlashingSubmitted
	initialSlashingSubmittedChan := make(chan *csmodule.CsmoduleInitialSlashingSubmitted)
	subInitialSlashingSubmitted, err := csModuleContract.WatchInitialSlashingSubmitted(&bind.WatchOpts{Context: ctx}, initialSlashingSubmittedChan, sa.NodeOperatorId)
	if err != nil {
		return err
	}

	// Watch for KeyRemovalChargeApplied
	keyRemovalChargeAppliedChan := make(chan *csmodule.CsmoduleKeyRemovalChargeApplied)
	subKeyRemovalChargeApplied, err := csModuleContract.WatchKeyRemovalChargeApplied(&bind.WatchOpts{Context: ctx}, keyRemovalChargeAppliedChan, sa.NodeOperatorId)
	if err != nil {
		return err
	}

	// Watch for NodeOperatorManagerAddressChangeProposed
	nodeOperatorManagerAddressChangeProposedChan := make(chan *csmodule.CsmoduleNodeOperatorManagerAddressChangeProposed)
	subNodeOperatorManagerAddressChangeProposed, err := csModuleContract.WatchNodeOperatorManagerAddressChangeProposed(&bind.WatchOpts{Context: ctx}, nodeOperatorManagerAddressChangeProposedChan, sa.NodeOperatorId, "", "") // TODO: add missing args
	if err != nil {
		return err
	}

	// Watch for NodeOperatorManagerAddressChanged
	nodeOperatorManagerAddressChangedChan := make(chan *csmodule.CsmoduleNodeOperatorManagerAddressChanged)
	subNodeOperatorManagerAddressChanged, err := csModuleContract.WatchNodeOperatorManagerAddressChanged(&bind.WatchOpts{Context: ctx}, nodeOperatorManagerAddressChangedChan, sa.NodeOperatorId, "", "") // TODO: add missing args
	if err != nil {
		return err
	}

	// Watch for NodeOperatorRewardAddressChangeProposed
	nodeOperatorRewardAddressChangeProposedChan := make(chan *csmodule.CsmoduleNodeOperatorRewardAddressChangeProposed)
	subNodeOperatorRewardAddressChangeProposed, err := csModuleContract.WatchNodeOperatorRewardAddressChangeProposed(&bind.WatchOpts{Context: ctx}, nodeOperatorRewardAddressChangeProposedChan, sa.NodeOperatorId, "", "") // TODO: add missing args
	if err != nil {
		return err
	}

	// Watch for NodeOperatorRewardAddressChanged
	nodeOperatorRewardAddressChangedChan := make(chan *csmodule.CsmoduleNodeOperatorRewardAddressChanged)
	subNodeOperatorRewardAddressChanged, err := csModuleContract.WatchNodeOperatorRewardAddressChanged(&bind.WatchOpts{Context: ctx}, nodeOperatorRewardAddressChangedChan, sa.NodeOperatorId, "", "") // TODO: add missing args
	if err != nil {
		return err
	}

	// Watch for StuckSigningKeysCountChanged
	stuckSigningKeysCountChangedChan := make(chan *csmodule.CsmoduleStuckSigningKeysCountChanged)
	subStuckSigningKeysCountChanged, err := csModuleContract.WatchStuckSigningKeysCountChanged(&bind.WatchOpts{Context: ctx}, stuckSigningKeysCountChangedChan, sa.NodeOperatorId)
	if err != nil {
		return err
	}

	// Watch for VettedSigningKeysCountDecreased
	vettedSigningKeysCountDecreasedChan := make(chan *csmodule.CsmoduleVettedSigningKeysCountDecreased)
	subVettedSigningKeysCountDecreased, err := csModuleContract.WatchVettedSigningKeysCountDecreased(&bind.WatchOpts{Context: ctx}, vettedSigningKeysCountDecreasedChan, sa.NodeOperatorId)
	if err != nil {
		return err
	}

	// Watch for WithdrawalSubmitted
	withdrawalSubmittedChan := make(chan *csmodule.CsmoduleWithdrawalSubmitted)
	subWithdrawalSubmitted, err := csModuleContract.WatchWithdrawalSubmitted(&bind.WatchOpts{Context: ctx}, withdrawalSubmittedChan, sa.NodeOperatorId)
	if err != nil {
		return err
	}

	// Watch for TotalSigningKeysCountChanged
	totalSigningKeysCountChangedChan := make(chan *csmodule.CsmoduleTotalSigningKeysCountChanged)
	subTotalSigningKeysCountChanged, err := csModuleContract.WatchTotalSigningKeysCountChanged(&bind.WatchOpts{Context: ctx}, totalSigningKeysCountChangedChan, sa.NodeOperatorId)
	if err != nil {
		return err
	}

	// Watch for PublicRelease
	publicReleaseChan := make(chan *csmodule.CsmodulePublicRelease)
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
				if err := handleLog(event); err != nil {
					log.Printf("Error handling DepositedSigningKeysCountChanged: %v", err)
				}
			case event := <-elRewardsStealingPenaltyReportedChan:
				if err := handleLog(event); err != nil {
					log.Printf("Error handling ELRewardsStealingPenaltyReported: %v", err)
				}
			case event := <-elRewardsStealingPenaltySettledChan:
				if err := handleLog(event); err != nil {
					log.Printf("Error handling ELRewardsStealingPenaltySettled: %v", err)
				}
			case event := <-elRewardsStealingPenaltyCancelledChan:
				if err := handleLog(event); err != nil {
					log.Printf("Error handling ELRewardsStealingPenaltyCancelled: %v", err)
				}
			case event := <-initialSlashingSubmittedChan:
				if err := handleLog(event); err != nil {
					log.Printf("Error handling InitialSlashingSubmitted: %v", err)
				}
			case event := <-keyRemovalChargeAppliedChan:
				if err := handleLog(event); err != nil {
					log.Printf("Error handling KeyRemovalChargeApplied: %v", err)
				}
			case event := <-nodeOperatorManagerAddressChangeProposedChan:
				if err := handleLog(event); err != nil {
					log.Printf("Error handling NodeOperatorManagerAddressChangeProposed: %v", err)
				}
			case event := <-nodeOperatorManagerAddressChangedChan:
				if err := handleLog(event); err != nil {
					log.Printf("Error handling NodeOperatorManagerAddressChanged: %v", err)
				}
			case event := <-nodeOperatorRewardAddressChangeProposedChan:
				if err := handleLog(event); err != nil {
					log.Printf("Error handling NodeOperatorRewardAddressChangeProposed: %v", err)
				}
			case event := <-nodeOperatorRewardAddressChangedChan:
				if err := handleLog(event); err != nil {
					log.Printf("Error handling NodeOperatorRewardAddressChanged: %v", err)
				}
			case event := <-stuckSigningKeysCountChangedChan:
				if err := handleLog(event); err != nil {
					log.Printf("Error handling StuckSigningKeysCountChanged: %v", err)
				}
			case event := <-vettedSigningKeysCountDecreasedChan:
				if err := handleLog(event); err != nil {
					log.Printf("Error handling VettedSigningKeysCountDecreased: %v", err)
				}
			case event := <-withdrawalSubmittedChan:
				if err := handleLog(event); err != nil {
					log.Printf("Error handling WithdrawalSubmitted: %v", err)
				}
			case event := <-totalSigningKeysCountChangedChan:
				if err := handleLog(event); err != nil {
					log.Printf("Error handling TotalSigningKeysCountChanged: %v", err)
				}
			case event := <-publicReleaseChan:
				if err := handleLog(event); err != nil {
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

func (sa *SubscriberAdapter) watchCsFeeDistributorEvents(ctx context.Context, handleLog func(logData interface{}) error) error {
	csFeeDistributorContract, err := csfeedistributor.NewCsfeedistributor(sa.CsFeeDistributorAddress, sa.client)
	if err != nil {
		return err
	}

	distributionDataUpdatedChan := make(chan *csfeedistributor.CsfeedistributorDistributionDataUpdated)
	sub, err := csFeeDistributorContract.WatchDistributionDataUpdated(&bind.WatchOpts{Context: ctx}, distributionDataUpdatedChan)
	if err != nil {
		return err
	}

	go func() {
		defer sub.Unsubscribe()
		for {
			select {
			case event := <-distributionDataUpdatedChan:
				if err := handleLog(event); err != nil {
					log.Printf("Error handling log: %v", err)
				}
			case err := <-sub.Err():
				log.Printf("Subscription error: %v", err)
				return
			case <-ctx.Done():
				return
			}
		}
	}()

	return nil
}

func (sa *SubscriberAdapter) watchVeboEvents(ctx context.Context, handleLog func(logData interface{}) error) error {
	veboContract, err := vebo.NewVebo(sa.VeboAddress, sa.client)
	if err != nil {
		return err
	}

	// Watch for ValidatorExitRequest
	validatorExitRequestChan := make(chan *vebo.VeboValidatorExitRequest)
	subExit, err := veboContract.WatchValidatorExitRequest(&bind.WatchOpts{Context: ctx}, validatorExitRequestChan, sa.StakingModuleId, sa.NodeOperatorId, sa.ValidatorIndex)
	if err != nil {
		return err
	}

	// Watch for ReportSubmitted
	reportSubmittedChan := make(chan *vebo.VeboReportSubmitted)
	subReport, err := veboContract.WatchReportSubmitted(&bind.WatchOpts{Context: ctx}, reportSubmittedChan, sa.RefSlot)
	if err != nil {
		return err
	}

	go func() {
		defer subExit.Unsubscribe()
		defer subReport.Unsubscribe()
		for {
			select {
			case event := <-validatorExitRequestChan:
				if err := handleLog(event); err != nil {
					log.Printf("Error handling ValidatorExitRequest: %v", err)
				}
			case event := <-reportSubmittedChan:
				if err := handleLog(event); err != nil {
					log.Printf("Error handling ReportSubmitted: %v", err)
				}
			case err := <-subExit.Err():
				log.Printf("Subscription error (ValidatorExitRequest): %v", err)
				return
			case err := <-subReport.Err():
				log.Printf("Subscription error (ReportSubmitted): %v", err)
				return
			case <-ctx.Done():
				return
			}
		}
	}()

	return nil
}
