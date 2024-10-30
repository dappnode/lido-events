package beaconchain

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// BeaconChainAdapter holds the URL for the Beacon chain API
type BeaconChainAdapter struct {
	beaconChainUrl string
}

// NewBeaconChainAdapter initializes a new BeaconChainAdapter with the given beaconChainUrl
func NewBeaconChainAdapter(beaconChainUrl string) *BeaconChainAdapter {
	return &BeaconChainAdapter{
		beaconChainUrl: beaconChainUrl,
	}
}

// SubmitPoolVoluntaryExit submits a voluntary exit message to the Beacon chain pool.
// It takes epoch, validatorIndex, and signature as arguments.
func (bca *BeaconChainAdapter) SubmitPoolVoluntaryExit(epoch, validatorIndex, signature string) error {
	// Define the request body structure
	body := map[string]interface{}{
		"message": map[string]string{
			"epoch":           epoch,
			"validator_index": validatorIndex,
		},
		"signature": signature,
	}

	// Marshal the body to JSON
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %w", err)
	}

	// Make the HTTP request
	url := fmt.Sprintf("%s/eth/v1/beacon/pool/voluntary_exits", bca.beaconChainUrl)
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


// GetStateFork retrieves the state fork for a specified state_id.
// API docs: https://ethereum.github.io/beacon-APIs/#/Beacon/getStateFork
func (bca *BeaconChainAdapter) GetStateFork(stateID string) (*StateForkResponse, error) {
	url := fmt.Sprintf("%s/eth/v1/beacon/states/%s/fork", bca.beaconChainUrl, stateID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	var result StateForkResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return &result, nil
}

// GetGenesis retrieves the genesis data.
// API docs: https://ethereum.github.io/beacon-APIs/#/Beacon/getGenesis
func (bca *BeaconChainAdapter) GetGenesis() (*GenesisResponse, error) {
	url := fmt.Sprintf("%s/eth/v1/beacon/genesis", bca.beaconChainUrl)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	var result GenesisResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return &result, nil
}

// GetBlockHeader retrieves the block header for a given block ID.
// API docs: https://ethereum.github.io/beacon-APIs/#/Beacon/getBlockHeader
func (bca *BeaconChainAdapter) GetBlockHeader(blockID string) (*BlockHeaderResponse, error) {
	url := fmt.Sprintf("%s/eth/v1/beacon/headers/%s", bca.beaconChainUrl, blockID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	var result BlockHeaderResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return &result, nil
}

// GetEpochHeader retrieves the epoch header for a specified block ID.
func (bca *BeaconChainAdapter) GetEpochHeader(blockID string) (int, error) {
	header, err := bca.GetBlockHeader(blockID)
	if err != nil {
		return 0, fmt.Errorf("failed to get block header: %w", err)
	}
	slot := header.Data.Header.Message.Slot
	epoch := bca.getEpochFromSlot(slot)
	return epoch, nil
}

// Helper function to convert slot to epoch.
func (bca *BeaconChainAdapter) getEpochFromSlot(slot string) int {
	// Convert slot to int and calculate epoch
	// Assume slot 0 is epoch 0, with slots per epoch set to 32
	const slotsPerEpoch = 32
	slotInt := parseInt(slot)
	return slotInt / slotsPerEpoch
}

// parseInt converts a string slot to an integer.
func parseInt(slot string) int {
	var result int
	fmt.Sscanf(slot, "%d", &result)
	return result
}

// Response structs for deserializing JSON responses

type StateForkResponse struct {
	ExecutionOptimistic bool `json:"execution_optimistic"`
	Finalized           bool `json:"finalized"`
	Data                struct {
		PreviousVersion string `json:"previous_version"`
		CurrentVersion  string `json:"current_version"`
		Epoch           string `json:"epoch"`
	} `json:"data"`
}

type GenesisResponse struct {
	Data struct {
		GenesisTime           string `json:"genesis_time"`
		GenesisValidatorsRoot string `json:"genesis_validators_root"`
		GenesisForkVersion    string `json:"genesis_fork_version"`
	} `json:"data"`
}

type BlockHeaderResponse struct {
	ExecutionOptimistic bool `json:"execution_optimistic"`
	Finalized           bool `json:"finalized"`
	Data                struct {
		Root      string `json:"root"`
		Canonical bool   `json:"canonical"`
		Header    struct {
			Message struct {
				Slot          string `json:"slot"`
				ProposerIndex string `json:"proposer_index"`
				ParentRoot    string `json:"parent_root"`
				StateRoot     string `json:"state_root"`
				BodyRoot      string `json:"body_root"`
			} `json:"message"`
			Signature string `json:"signature"`
		} `json:"header"`
	} `json:"data"`
}
