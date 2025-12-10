package ports

type ExecutionPort interface {
	GetMostRecentBlockNumber() (uint64, error)
	IsSyncing() (bool, error)
	GetBlockHasReceipts(blockID string) (bool, error)
}
