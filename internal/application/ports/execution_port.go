package ports

import "github.com/ethereum/go-ethereum/common"

type ExecutionPort interface {
	GetMostRecentBlockNumber() (uint64, error)
	GetBlockTimestampByNumber(blockNumber uint64) (uint64, error)
	IsSyncing() (bool, error)
	GetTransactionReceipt(txHash common.Hash) (map[string]interface{}, error)
}
