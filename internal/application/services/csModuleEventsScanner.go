package services

import (
	"context"
	"fmt"
	"lido-events/internal/application/domain"
	"lido-events/internal/application/ports"
	"lido-events/internal/logger"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type CsModuleEventsScanner struct {
	storagePort                     ports.StoragePort
	executionPort                   ports.ExecutionPort
	csModulePort                    ports.CsModulePort
	csFeeDistributorBlockDeployment uint64
	csModuleTxReceipt               common.Hash
	BlockScannerMinDistance         uint64
	servicePrefix                   string
}

func NewCsModuleEventsScanner(storagePort ports.StoragePort, executionPort ports.ExecutionPort, csModulePort ports.CsModulePort, csFeeDistributorBlockDeployment uint64, csModuleTxReceipt common.Hash, blockScannerMinDistance uint64) *CsModuleEventsScanner {
	return &CsModuleEventsScanner{
		storagePort:                     storagePort,
		executionPort:                   executionPort,
		csModulePort:                    csModulePort,
		csFeeDistributorBlockDeployment: csFeeDistributorBlockDeployment,
		csModuleTxReceipt:               csModuleTxReceipt,
		BlockScannerMinDistance:         blockScannerMinDistance,
		servicePrefix:                   "CsModuleEventsScanner",
	}
}

func (cs *CsModuleEventsScanner) ScanWithdrawalsSubmittedEvents(ctx context.Context, operatorId *big.Int) error {
	isSyncing, err := cs.executionPort.IsSyncing()
	if err != nil {
		return fmt.Errorf("error checking if node is syncing: %w", err)
	}

	if isSyncing {
		return fmt.Errorf("node is syncing")
	}

	// Skip if tx receipt not found. This means that the node does not store log receipts and there are no logs at all
	receiptExists, err := cs.executionPort.GetTransactionReceiptExists(cs.csModuleTxReceipt)
	if err != nil {
		return fmt.Errorf("error checking if transaction receipt exists: %w", err)
	}
	if !receiptExists {
		return fmt.Errorf("transaction receipt for csModule deployment not found")
	}

	start, err := cs.storagePort.GetWithdrawalsSubmittedLastProcessedBlock(operatorId)
	if err != nil {
		logger.WarnWithPrefix(cs.servicePrefix, "Failed to get last processed block, using deployment block: %v", err)
		start = cs.csFeeDistributorBlockDeployment
	}

	if start == 0 {
		logger.InfoWithPrefix(cs.servicePrefix, "No last processed block found, using deployment block")
		start = cs.csFeeDistributorBlockDeployment
	}

	end, err := cs.executionPort.GetMostRecentBlockNumber()
	if err != nil {
		return fmt.Errorf("error getting most recent block number: %w", err)
	}

	// return if distance is less than 10 epoch = 10 * 32 blocks
	if end-start < cs.BlockScannerMinDistance {
		logger.InfoWithPrefix(cs.servicePrefix, "Block distance is less than 320 blocks (10 epochs), skipping WithdrawalsSubmittedEvents scan")
		return nil
	}

	logger.DebugWithPrefix(cs.servicePrefix, "Scanning WithdrawalsSubmitted events from block %d to %d", start, end)
	if err := cs.csModulePort.ScanWithdrawalSubmitted(
		ctx,
		operatorId,
		start,
		&end,
		cs.HandleWithdrawalsSubmittedEvent,
	); err != nil {
		return fmt.Errorf("error scanning WithdrawalsSubmitted events: %w", err)
	}

	if err := cs.storagePort.SaveWithdrawalsSubmittedLastProcessedBlock(operatorId, end); err != nil {
		return fmt.Errorf("error saving last processed block: %w", err)
	}

	logger.DebugWithPrefix(cs.servicePrefix, "WithdrawalsSubmitted events scan completed")
	return nil
}

// ScanElRewardsStealingPenaltyReported
func (cs *CsModuleEventsScanner) ScanElRewardsStealingPenaltyReported(ctx context.Context) error {
	isSyncing, err := cs.executionPort.IsSyncing()
	if err != nil {
		return fmt.Errorf("error checking if node is syncing: %w", err)
	}

	if isSyncing {
		return fmt.Errorf("node is syncing")
	}

	// Skip if tx receipt not found. This means that the node does not store log receipts and there are no logs at all
	receiptExists, err := cs.executionPort.GetTransactionReceiptExists(cs.csModuleTxReceipt)
	if err != nil {
		return fmt.Errorf("error checking if transaction receipt exists: %w", err)
	}

	if !receiptExists {
		return fmt.Errorf("transaction receipt for csModule deployment not found")
	}

	start, err := cs.storagePort.GetElRewardsStealingPenaltiesReportedLastProcessedBlock()
	if err != nil {
		logger.WarnWithPrefix(cs.servicePrefix, "Failed to get last processed block, using deployment block: %v", err)
		start = cs.csFeeDistributorBlockDeployment
	}

	if start == 0 {
		logger.InfoWithPrefix(cs.servicePrefix, "No last processed block found, using deployment block")
		start = cs.csFeeDistributorBlockDeployment
	}

	end, err := cs.executionPort.GetMostRecentBlockNumber()
	if err != nil {
		return fmt.Errorf("error getting most recent block number: %w", err)
	}

	// return if distance is less than 10 epoch = 10 * 32 blocks
	if end-start < cs.BlockScannerMinDistance {
		logger.InfoWithPrefix(cs.servicePrefix, "Block distance is less than 320 blocks (10 epochs), skipping ElRewardsStealingPenaltyReported scan")
		return nil
	}

	logger.DebugWithPrefix(cs.servicePrefix, "Scanning ElRewardsStealingPenaltyReported events from block %d to %d", start, end)
	if err := cs.csModulePort.ScanElRewardsStealingPenaltyReported(
		ctx,
		start,
		&end,
		cs.HandleElRewardsStealingPenaltyReported,
	); err != nil {
		return fmt.Errorf("error scanning ElRewardsStealingPenaltyReported events: %w", err)
	}

	if err := cs.storagePort.SaveElRewardsStealingPenaltiesReportedLastProcessedBlock(end); err != nil {
		return fmt.Errorf("error saving last processed block: %w", err)
	}

	logger.DebugWithPrefix(cs.servicePrefix, "ElRewardsStealingPenaltyReported events scan completed")
	return nil
}

func (cs *CsModuleEventsScanner) ScanAddressEvents(ctx context.Context, address common.Address, force bool) error {
	isSyncing, err := cs.executionPort.IsSyncing()
	if err != nil {
		return fmt.Errorf("error checking if node is syncing: %w", err)
	}

	if isSyncing {
		return fmt.Errorf("node is syncing")
	}

	// Skip if tx receipt not found. This means that the node does not store log receipts and there are no logs at all
	receiptExists, err := cs.executionPort.GetTransactionReceiptExists(cs.csModuleTxReceipt)
	if err != nil {
		return fmt.Errorf("error checking if transaction receipt exists: %w", err)
	}
	if !receiptExists {
		return fmt.Errorf("transaction receipt for csModule deployment not found")
	}

	start, err := cs.storagePort.GetAddressLastProcessedBlock(address)
	if err != nil {
		logger.WarnWithPrefix(cs.servicePrefix, "Failed to get last processed block, using deployment block: %v", err)
		start = cs.csFeeDistributorBlockDeployment
	}

	if start == 0 {
		logger.InfoWithPrefix(cs.servicePrefix, "No last processed block found, using deployment block")
		start = cs.csFeeDistributorBlockDeployment
	}

	end, err := cs.executionPort.GetMostRecentBlockNumber()
	if err != nil {
		return fmt.Errorf("error getting most recent block number: %w", err)
	}

	// return if distance is less than min distance AND not forced
	if !force && end-start < cs.BlockScannerMinDistance {
		logger.InfoWithPrefix(cs.servicePrefix, "Block distance is less than 320 blocks (10 epochs), skipping aaddress events scan")
		return nil
	}

	logger.DebugWithPrefix(cs.servicePrefix, "Scanning events for address %s from block %d to %d", address.Hex(), start, end)
	if err := cs.csModulePort.ScanNodeOperatorEvents(
		ctx,
		address,
		start,
		&end,
		cs.HandleNodeOperatorAddedEvent,
		cs.HandleNodeOperatorManagerAddressChangedEvent,
		cs.HandleNodeOperatorRewardAddressChangedEvent,
		cs.HandleNodeOperatorRewardAddressChangeProposedEvent,
		cs.HandleNodeOperatorManagerAddressChangeProposedEvent,
	); err != nil {
		return fmt.Errorf("error scanning NodeOperator events: %w", err)
	}

	if err := cs.storagePort.SaveAddressLastProcessedBlock(address, end); err != nil {
		return fmt.Errorf("error saving last processed block: %w", err)
	}

	logger.DebugWithPrefix(cs.servicePrefix, "Address events scan completed")
	return nil
}

func (cs *CsModuleEventsScanner) HandleNodeOperatorAddedEvent(event *domain.CsmoduleNodeOperatorAdded, address common.Address) error {
	logger.DebugWithPrefix(cs.servicePrefix, "Handling NodeOperatorAdded event: %+v", event)

	if err := cs.storagePort.SetNodeOperatorAdded(address, *event); err != nil {
		logger.ErrorWithPrefix(cs.servicePrefix, "Error saving NodeOperatorAdded event: %v", err)
		return err
	}
	return nil
}

func (cs *CsModuleEventsScanner) HandleNodeOperatorManagerAddressChangedEvent(event *domain.CsmoduleNodeOperatorManagerAddressChanged, address common.Address) error {
	logger.DebugWithPrefix(cs.servicePrefix, "Handling NodeOperatorManagerAddressChanged event: %+v", event)

	if err := cs.storagePort.SetNodeOperatorManagerAddressChanged(address, *event); err != nil {
		logger.ErrorWithPrefix(cs.servicePrefix, "Error saving NodeOperatorManagerAddressChanged event: %v", err)
		return err
	}
	return nil
}

func (cs *CsModuleEventsScanner) HandleNodeOperatorRewardAddressChangedEvent(event *domain.CsmoduleNodeOperatorRewardAddressChanged, address common.Address) error {
	logger.DebugWithPrefix(cs.servicePrefix, "Handling NodeOperatorRewardAddressChanged event: %+v", event)

	if err := cs.storagePort.SetNodeOperatorRewardAddressChanged(address, *event); err != nil {
		logger.ErrorWithPrefix(cs.servicePrefix, "Error saving NodeOperatorRewardAddressChanged event: %v", err)
		return err
	}
	return nil
}

func (cs *CsModuleEventsScanner) HandleNodeOperatorRewardAddressChangeProposedEvent(event *domain.CsmoduleNodeOperatorRewardAddressChangeProposed, address common.Address) error {
	logger.DebugWithPrefix(cs.servicePrefix, "Handling NodeOperatorRewardAddressChangeProposed event: %+v", event)

	if err := cs.storagePort.SetNodeOperatorRewardAddressChangeProposed(address, *event); err != nil {
		logger.ErrorWithPrefix(cs.servicePrefix, "Error saving NodeOperatorRewardAddressChangeProposed event: %v", err)
		return err
	}
	return nil
}

func (cs *CsModuleEventsScanner) HandleNodeOperatorManagerAddressChangeProposedEvent(event *domain.CsmoduleNodeOperatorManagerAddressChangeProposed, address common.Address) error {
	logger.DebugWithPrefix(cs.servicePrefix, "Handling NodeOperatorManagerAddressChangeProposed event: %+v", event)

	if err := cs.storagePort.SetNodeOperatorManagerAddressChangeProposed(address, *event); err != nil {
		logger.ErrorWithPrefix(cs.servicePrefix, "Error saving NodeOperatorManagerAddressChangeProposed event: %v", err)
		return err
	}
	return nil
}

func (cs *CsModuleEventsScanner) HandleWithdrawalsSubmittedEvent(withdrawalSubmitted *domain.CsmoduleWithdrawalSubmitted, operatorId *big.Int) error {
	logger.DebugWithPrefix(cs.servicePrefix, "Found WithdrawalsSubmitted event")

	if err := cs.storagePort.SaveWithdrawal(operatorId, *withdrawalSubmitted); err != nil {
		logger.ErrorWithPrefix(cs.servicePrefix, "Error saving WithdrawalsSubmitted event: %v", err)
		return err
	}
	return nil
}

func (cs *CsModuleEventsScanner) HandleElRewardsStealingPenaltyReported(elRewardsStealingPenaltyReported *domain.CsmoduleELRewardsStealingPenaltyReported) error {
	logger.DebugWithPrefix(cs.servicePrefix, "Found ElRewardsStealingPenaltyReported event")

	if err := cs.storagePort.SaveElRewardsStealingPenaltyReported(*elRewardsStealingPenaltyReported); err != nil {
		logger.ErrorWithPrefix(cs.servicePrefix, "Error saving ElRewardsStealingPenaltyReported event: %v", err)
		return err
	}
	return nil
}
