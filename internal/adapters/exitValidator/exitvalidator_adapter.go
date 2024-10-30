package exitvalidator

import (
	"fmt"
	"lido-events/internal/adapters/beaconchain"
	"lido-events/internal/adapters/signer"
)

// ExitValidatorAdapter orchestrates the exit validator process using Signer and BeaconChain adapters
type ExitValidatorAdapter struct {
	beaconChain *beaconchain.BeaconChainAdapter
	signer      *signer.SignerAdapter
}

// NewExitValidatorAdapter creates a new instance of ExitValidatorAdapter with provided beaconChain and signer adapters
func NewExitValidatorAdapter(beaconChain *beaconchain.BeaconChainAdapter, signer *signer.SignerAdapter) *ExitValidatorAdapter {
	return &ExitValidatorAdapter{
		beaconChain: beaconChain,
		signer:      signer,
	}
}

// ExitValidator orchestrates the voluntary exit process for a validator
func (e *ExitValidatorAdapter) ExitValidator(pubKey, validatorIndex string) error {
	// Step 1: Get the fork information
	forkInfo, err := e.beaconChain.GetStateFork("head")
	if err != nil {
		return fmt.Errorf("failed to get state fork: %w", err)
	}

	// Step 2: Get the genesis information
	genesisInfo, err := e.beaconChain.GetGenesis()
	if err != nil {
		return fmt.Errorf("failed to get genesis info: %w", err)
	}

	// Step 3: Get the current epoch (assumed to be head’s slot converted to epoch)
	header, err := e.beaconChain.GetBlockHeader("head")
	if err != nil {
		return fmt.Errorf("failed to get block header: %w", err)
	}
	currentEpoch, err := e.beaconChain.GetEpochHeader(header.Data.Header.Message.Slot)
	if err != nil {
		return fmt.Errorf("failed to get epoch header: %w", err)
	}

	// Step 4: Prepare the signing root (dummy value here) and voluntary exit data
	signingRoot := "0x..." // Assume this value is provided or calculated elsewhere in real implementation
	voluntaryExitReq := signer.VoluntaryExitRequest{
		PubKey:                pubKey,
		SigningRoot:           signingRoot,
		GenesisValidatorsRoot: genesisInfo.Data.GenesisValidatorsRoot,
		ForkInfo: signer.ForkInfo{
			PreviousVersion: forkInfo.Data.PreviousVersion,
			CurrentVersion:  forkInfo.Data.CurrentVersion,
			Epoch:           forkInfo.Data.Epoch,
		},
		VoluntaryExit: signer.VoluntaryExit{
			Epoch:          fmt.Sprintf("%d", currentEpoch),
			ValidatorIndex: validatorIndex,
		},
	}

	// Step 5: Sign the voluntary exit using the Signer adapter
	signature, err := e.signer.Eth2SignVoluntaryExit(voluntaryExitReq)
	if err != nil {
		return fmt.Errorf("failed to sign voluntary exit: %w", err)
	}

	// Step 6: Submit the signed voluntary exit to the beacon chain pool
	if err := e.beaconChain.SubmitPoolVoluntaryExit(voluntaryExitReq.VoluntaryExit.Epoch, validatorIndex, signature); err != nil {
		return fmt.Errorf("failed to submit pool voluntary exit: %w", err)
	}

	return nil
}
