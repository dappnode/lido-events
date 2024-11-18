package exitvalidator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// VoluntaryExitRequest holds parameters needed to sign a voluntary exit
type VoluntaryExitRequest struct {
	PubKey string
	// SigningRoot           string
	GenesisValidatorsRoot string
	ForkInfo              ForkInfo
	VoluntaryExit         VoluntaryExit
}

// Eth2SignVoluntaryExit sends a request to the Web3Signer URL to sign a voluntary exit message
// and returns the signature as a string.
// API docs: https://consensys.github.io/web3signer/web3signer-eth2.html#tag/Signing/operation/ETH2_SIGN
func Eth2SignVoluntaryExit(signerUrl string, req VoluntaryExitRequest) (string, error) {
	identifier := req.PubKey // Identifier is the BLS public key in hex format.

	// Define the request body structure
	body := map[string]interface{}{
		"type": "VOLUNTARY_EXIT",
		//"signingRoot": req.SigningRoot,
		"fork_info": map[string]interface{}{
			"fork": map[string]interface{}{
				"previous_version": req.ForkInfo.PreviousVersion,
				"current_version":  req.ForkInfo.CurrentVersion,
				"epoch":            req.ForkInfo.Epoch,
			},
			"genesis_validators_root": req.GenesisValidatorsRoot,
		},
		"voluntary_exit": map[string]interface{}{
			"epoch":           req.VoluntaryExit.Epoch,
			"validator_index": req.VoluntaryExit.ValidatorIndex,
		},
	}

	// Marshal the body to JSON
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request body: %w", err)
	}

	// Make the HTTP request
	url := fmt.Sprintf("%s/api/v1/eth2/sign/%s", signerUrl, identifier)
	reqHttp, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	reqHttp.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(reqHttp)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Decode the response body to extract the signature
	var result struct {
		Signature string `json:"signature"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	return result.Signature, nil
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
