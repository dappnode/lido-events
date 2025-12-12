package exitvalidator

import (
	"fmt"
	"lido-events/internal/adapters/beaconchain" // Import the BeaconchainAdapter package
)

// ExitValidator orchestrates the exit validator process using Signer and BeaconChain adapters
type ExitValidator struct {
	beaconChainAdapter *beaconchain.Beaconchain // Dependency injection for the BeaconchainAdapter
	signerUrl          string
}

// NewExitValidator creates a new instance of ExitValidator with provided Beaconchain and Signer adapters
func NewExitValidator(beaconChainAdapter *beaconchain.Beaconchain, signerUrl string) *ExitValidator {
	return &ExitValidator{
		beaconChainAdapter: beaconChainAdapter,
		signerUrl:          signerUrl,
	}
}

// ExitValidator orchestrates the voluntary exit process for a validator
func (e *ExitValidator) ExitValidator(pubKey, validatorIndex string) error {
	// Step 1: Get the fork information
	forkInfo, err := e.beaconChainAdapter.GetStateFork("head")
	if err != nil {
		return fmt.Errorf("failed to get state fork: %w", err)
	}

	// Step 2: Get the genesis information
	genesisInfo, err := e.beaconChainAdapter.GetGenesis()
	if err != nil {
		return fmt.Errorf("failed to get genesis info: %w", err)
	}

	// Step 3: Get the current epoch (assumed to be head’s slot converted to epoch)
	currentEpoch, err := e.beaconChainAdapter.GetEpochHeader("head")
	if err != nil {
		return fmt.Errorf("failed to get epoch header: %w", err)
	}

	// Step 4: Prepare the signing root and voluntary exit data
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
	if err := e.beaconChainAdapter.SubmitPoolVoluntaryExit(voluntaryExitReq.VoluntaryExit.Epoch, validatorIndex, signature); err != nil {
		return fmt.Errorf("failed to submit pool voluntary exit: %w", err)
	}

	return nil
}
