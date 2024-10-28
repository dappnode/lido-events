package domain

type Reports map[uint32]Report // indexed by epoch

type Report struct {
	Threshold float64
	Validators Validators
	RefSlot uint32 // lido report returns `ref_slot`, which can be converted to epoch
}

type Validators map[string]ValidatorPerformance // indexed by validator pubkey

type ValidatorPerformance struct {
	Included uint
	Assigned uint
}
