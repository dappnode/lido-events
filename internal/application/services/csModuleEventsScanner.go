package services

import (
	"context"
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

func (cs *CsModuleEventsScanner) ScanAddressEvents(ctx context.Context, address common.Address) {
	isSyncing, err := cs.executionPort.IsSyncing()
	if err != nil {
		logger.ErrorWithPrefix(cs.servicePrefix, "Error checking if node is syncing: %v", err)
		return
	}

	if isSyncing {
		logger.InfoWithPrefix(cs.servicePrefix, "Node is syncing, skipping CsModule events scan")
		return
	}

	// Skip if tx receipt not found. This means that the node does not store log receipts and there are no logs at all
	receiptExists, err := cs.executionPort.GetTransactionReceiptExists(cs.csModuleTxReceipt)
	if err != nil {
		logger.ErrorWithPrefix(cs.servicePrefix, "Error checking if transaction receipt exists: %v", err)
		return
	}
	if !receiptExists {
		logger.WarnWithPrefix(cs.servicePrefix, "Transaction receipt for csModule deployment not found. This probably means your node does not store log receipts, check out the official documentation of your node and configure the node to store them. Skipping CsModule events scan")
		return
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
		logger.ErrorWithPrefix(cs.servicePrefix, "Failed to get most recent block number: %v", err)
		return
	}

	// return if start block is greater than end block
	if start > end {
		logger.WarnWithPrefix(cs.servicePrefix, "Start block is greater than end block, skipping CsModule events scan")
		return
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
		logger.ErrorWithPrefix(cs.servicePrefix, "Error scanning NodeOperator events: %v", err)
		return
	}

	if err := cs.storagePort.SaveAddressLastProcessedBlock(address, end); err != nil {
		logger.ErrorWithPrefix(cs.servicePrefix, "Error saving last processed block: %v", err)
		return
	}

	logger.DebugWithPrefix(cs.servicePrefix, "Address events scan completed")
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
