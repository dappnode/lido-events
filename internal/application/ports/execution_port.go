package ports

type ExecutionPort interface {
	GetMostRecentBlockNumber() (uint64, error)
}
