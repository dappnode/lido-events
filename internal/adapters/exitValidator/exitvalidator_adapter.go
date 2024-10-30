package exitvalidator

import (
	"fmt"
)

// ExitValidatorAdapter orchestrates the exit validator process using Signer and BeaconChain adapters
type ExitValidatorAdapter struct {
	beaconChainUrl string
	signerUrl      string
}

// NewExitValidatorAdapter creates a new instance of ExitValidatorAdapter with provided beaconChain and signer adapters
func NewExitValidatorAdapter(beaconChainUrl string, signerUrl string) *ExitValidatorAdapter {
	return &ExitValidatorAdapter{
		beaconChainUrl,
		signerUrl,
	}
}

// ExitValidator orchestrates the voluntary exit process for a validator
func (e *ExitValidatorAdapter) ExitValidator(pubKey, validatorIndex string) error {
	// Step 1: Get the fork information
	forkInfo, err := GetStateFork(e.beaconChainUrl, "head")
	if err != nil {
		return fmt.Errorf("failed to get state fork: %w", err)
	}

	// Step 2: Get the genesis information
	genesisInfo, err := GetGenesis(e.beaconChainUrl)
	if err != nil {
		return fmt.Errorf("failed to get genesis info: %w", err)
	}

	// Step 3: Get the current epoch (assumed to be head’s slot converted to epoch)
	header, err := GetBlockHeader(e.beaconChainUrl, "head")
	if err != nil {
		return fmt.Errorf("failed to get block header: %w", err)
	}
	currentEpoch, err := GetEpochHeader(e.beaconChainUrl, header.Data.Header.Message.Slot)
	if err != nil {
		return fmt.Errorf("failed to get epoch header: %w", err)
	}

	// Step 4: Prepare the signing root (dummy value here) and voluntary exit data
	voluntaryExitReq := VoluntaryExitRequest{
		PubKey:                pubKey,
		GenesisValidatorsRoot: genesisInfo.Data.GenesisValidatorsRoot,
		ForkInfo: ForkInfo{
			PreviousVersion: forkInfo.Data.PreviousVersion,
			CurrentVersion:  forkInfo.Data.CurrentVersion,
			Epoch:           forkInfo.Data.Epoch,
		},
		VoluntaryExit: VoluntaryExit{
			Epoch:          fmt.Sprintf("%d", currentEpoch),
			ValidatorIndex: validatorIndex,
		},
	}

	// Step 5: Sign the voluntary exit using the Signer adapter
	signature, err := Eth2SignVoluntaryExit(e.signerUrl, voluntaryExitReq)
	if err != nil {
		return fmt.Errorf("failed to sign voluntary exit: %w", err)
	}

	// Step 6: Submit the signed voluntary exit to the beacon chain pool
	if err := SubmitPoolVoluntaryExit(e.beaconChainUrl, voluntaryExitReq.VoluntaryExit.Epoch, validatorIndex, signature); err != nil {
		return fmt.Errorf("failed to submit pool voluntary exit: %w", err)
	}

	return nil
}
