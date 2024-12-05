package ports

type ExecutionPort interface {
	GetMostRecentBlockNumber() (uint64, error)
	GetBlockTimestampByNumber(blockNumber uint64) (uint64, error)
	IsSyncing() (bool, error)
}
