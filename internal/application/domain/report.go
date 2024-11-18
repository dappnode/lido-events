package domain

// Define the Reports type as a map indexed by frame
type Reports map[string]Report // indexed by "frame" in the format "start-end"

// Define the new Report structure
type Performance struct {
	Assigned int `json:"assigned"`
	Included int `json:"included"`
}

type Validator struct {
	Perf    Performance `json:"perf"`
	Slashed bool        `json:"slashed"`
}

type Data struct {
	Distributed int64                `json:"distributed"`
	Stuck       bool                 `json:"stuck"`
	Validators  map[string]Validator `json:"validators"`
}

type Report struct {
	Frame     [2]int  `json:"frame"`
	Data      Data    `json:"data"`
	Threshold float64 `json:"threshold"`
}

type OriginalReport struct {
	Frame     [2]int          `json:"frame"`
	Operators map[string]Data `json:"operators"` // indexed by operator ID
	Threshold float64         `json:"threshold"`
}
