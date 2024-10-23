package storage

// Storage handles file-based storage for validator performance and exit requests
type Storage struct {
	PerformanceFile string
	ExitRequestFile string
	ConfigFile      string
}

func NewStorageAdapter() *Storage {
	return &Storage{
		PerformanceFile: "validators-performance.json",
		ExitRequestFile: "exit-requests.json",
		ConfigFile:      "config.json",
	}
}
