package domain

// i.e https://ipfs.io/ipfs/QmPQxGAFtFz2whzcPK1a2vm3ewPt2iVk9Q58etFVqwyiWM

type Validator struct {
	Performance float64 `json:"performance"` // 0.846349206349206
	Slashed     bool    `json:"slashed"`     // false
	Strikes     int     `json:"strikes"`     // 1
	Threshold   float64 `json:"threshold"`   // 0.895871468500919
}

type Data struct {
	DistributedRewards int64                `json:"distributed_rewards"` // 672741727797680
	Validators         map[string]Validator `json:"validators"`          // indexed by validator index
}

type Report struct {
	Frame     [2]int          `json:"frame"`
	Operators map[string]Data `json:"operators"` // indexed by operator ID
}
