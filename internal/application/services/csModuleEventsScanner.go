package services

import (
	"context"
	"fmt"
	"lido-events/internal/application/domain"
	"lido-events/internal/application/ports"
	"lido-events/internal/logger"

	"github.com/ethereum/go-ethereum/common"
)

type CsModuleEventsScanner struct {
	storagePort                     ports.StoragePort
	executionPort                   ports.ExecutionPort
	csModulePort                    ports.CsModulePort
	csFeeDistributorBlockDeployment uint64
	csModuleTxReceipt               common.Hash
	servicePrefix                   string
}

func NewCsModuleEventsScanner(storagePort ports.StoragePort, executionPort ports.ExecutionPort, csModulePort ports.CsModulePort, csFeeDistributorBlockDeployment uint64, csModuleTxReceipt common.Hash) *CsModuleEventsScanner {
	return &CsModuleEventsScanner{
		storagePort:                     storagePort,
		executionPort:                   executionPort,
		csModulePort:                    csModulePort,
		csFeeDistributorBlockDeployment: csFeeDistributorBlockDeployment,
		csModuleTxReceipt:               csModuleTxReceipt,
		servicePrefix:                   "CsModuleEventsScanner",
	}
}

func (cs *CsModuleEventsScanner) ScanAddressEvents(ctx context.Context, address common.Address) error {
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

	// return if start block is greater than end block
	if start > end {
		return fmt.Errorf("start block is greater than end block")
	}

	// to avoid spamming the node with requests, only scan if the distance is less than 500 blocks | 100 minutes | 1.67 hours
	if end-start > 500 {
		logger.InfoWithPrefix(cs.servicePrefix, "Block distance is greater than 500, skipping scan to avoid spamming the node")
		return nil
	}

	if err := cs.csModulePort.ScanNodeOperatorEvents(
		ctx,
		address,
		start,
		&end,
		cs.HandleNodeOperatorAddedEvent,
		cs.HandleNodeOperatorManagerAddressChangedEvent,
		cs.HandleNodeOperatorRewardAddressChangedEvent,
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
