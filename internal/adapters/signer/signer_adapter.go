package signer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// TODO: add return type

type SignerAdapter struct {
	signerUrl string
}

// NewSignerAdapter initializes a new SignerAdapter with the given signerUrl
func NewSignerAdapter(signerUrl string) *SignerAdapter {
	return &SignerAdapter{
		signerUrl: signerUrl,
	}
}

// Eth2SignVoluntaryExit sends a request to the Web3Signer URL to sign a voluntary exit message.
// API docs: https://consensys.github.io/web3signer/web3signer-eth2.html#tag/Signing/operation/ETH2_SIGN
func (sa *SignerAdapter) Eth2SignVoluntaryExit(pubKey, signingRoot, genesisValidatorsRoot string, forkInfo ForkInfo, voluntaryExit VoluntaryExit) error {
	identifier := pubKey // Identifier is the BLS public key in hex format.

	// Define the request body structure
	body := map[string]interface{}{
		"type":        "VOLUNTARY_EXIT",
		"signingRoot": signingRoot,
		"fork_info": map[string]interface{}{
			"fork": map[string]interface{}{
				"previous_version": forkInfo.PreviousVersion,
				"current_version":  forkInfo.CurrentVersion,
				"epoch":            forkInfo.Epoch,
			},
			"genesis_validators_root": genesisValidatorsRoot,
		},
		"voluntary_exit": map[string]interface{}{
			"epoch":           voluntaryExit.Epoch,
			"validator_index": voluntaryExit.ValidatorIndex,
		},
	}

	// Marshal the body to JSON
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %w", err)
	}

	// Make the HTTP request
	url := fmt.Sprintf("%s/api/v1/eth2/sign/%s", sa.signerUrl, identifier)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

// ForkInfo holds the fork information
type ForkInfo struct {
	PreviousVersion string
	CurrentVersion  string
	Epoch           string
}

// VoluntaryExit holds the voluntary exit information
type VoluntaryExit struct {
	Epoch          string
	ValidatorIndex string
}
