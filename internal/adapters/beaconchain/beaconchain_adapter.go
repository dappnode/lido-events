package beaconchain

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"lido-events/internal/application/domain"
	"net/http"
)

// BeaconchainAdapter orchestrates the exit validator process using Signer and BeaconChain adapters.
type BeaconchainAdapter struct {
	beaconChainUrl string
}

// NewBeaconchainAdapter creates a new instance of ExitValidatorAdapter with provided beaconChain and signer adapters.
func NewBeaconchainAdapter(beaconChainUrl string) *BeaconchainAdapter {
	return &BeaconchainAdapter{
		beaconChainUrl: beaconChainUrl,
	}
}

func (b *BeaconchainAdapter) GetValidatorStatus(pubkey string) (domain.ValidatorStatus, error) {
	validatorData, err := b.PostStateValidators("finalized", []string{pubkey}, nil)
	if err != nil {
		return "", fmt.Errorf("failed to get validator data for pubkey %s: %w", pubkey, err)
	}

	if len(validatorData.Data) == 0 {
		return "", errors.New("validator not found")
	}

	return domain.ValidatorStatus(validatorData.Data[0].Status), nil
}

// PostStateValidators sends a POST request to fetch validator data based on provided ids and statuses.
// API docs: https://ethereum.github.io/beacon-APIs/#/Beacon/postStateValidators
func (b *BeaconchainAdapter) PostStateValidators(stateID string, ids []string, statuses []domain.ValidatorStatus) (*ValidatorsResponse, error) {
	body := map[string]interface{}{
		"ids":      ids,
		"statuses": statuses,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body for PostStateValidators: %w", err)
	}

	reqUrl := fmt.Sprintf("%s/eth/v1/beacon/states/%s/validators", b.beaconChainUrl, stateID)
	req, err := http.NewRequest("POST", reqUrl, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create POST request for URL %s: %w", reqUrl, err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request to Beaconchain at %s: %w", reqUrl, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d fetching validator data", resp.StatusCode)
	}

	var result ValidatorsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response from Beaconchain: %w", err)
	}

	return &result, nil
}

// SubmitPoolVoluntaryExit submits a voluntary exit message to the Beacon chain pool.
// It takes epoch, validatorIndex, and signature as arguments.
func (b *BeaconchainAdapter) SubmitPoolVoluntaryExit(epoch, validatorIndex, signature string) error {
	body := map[string]interface{}{
		"message": map[string]string{
			"epoch":           epoch,
			"validator_index": validatorIndex,
		},
		"signature": signature,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("failed to marshal request body for SubmitPoolVoluntaryExit: %w", err)
	}

	url := fmt.Sprintf("%s/eth/v1/beacon/pool/voluntary_exits", b.beaconChainUrl)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("failed to create POST request for URL %s: %w", url, err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request to Beaconchain at %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code %d submitting voluntary exit", resp.StatusCode)
	}

	return nil
}

// GetStateFork retrieves the state fork for a specified state_id.
// API docs: https://ethereum.github.io/beacon-APIs/#/Beacon/getStateFork
func (b *BeaconchainAdapter) GetStateFork(stateID string) (*StateForkResponse, error) {
	url := fmt.Sprintf("%s/eth/v1/beacon/states/%s/fork", b.beaconChainUrl, stateID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to send request to Beaconchain at %s: %w", url, err)
	}
	defer resp.Body.Close()

	var result StateForkResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response for GetStateFork: %w", err)
	}

	return &result, nil
}

// GetGenesis retrieves the genesis data.
// API docs: https://ethereum.github.io/beacon-APIs/#/Beacon/getGenesis
func (b *BeaconchainAdapter) GetGenesis() (*GenesisResponse, error) {
	url := fmt.Sprintf("%s/eth/v1/beacon/genesis", b.beaconChainUrl)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to send request to Beaconchain at %s: %w", url, err)
	}
	defer resp.Body.Close()

	var result GenesisResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response for GetGenesis: %w", err)
	}

	return &result, nil
}

// GetBlockHeader retrieves the block header for a given block ID.
// Block identifier. Can be one of: "head" (canonical head in node's view), "genesis", "finalized", <slot>, <hex encoded blockRoot with 0x prefix>.
// API docs: https://ethereum.github.io/beacon-APIs/#/Beacon/getBlockHeader
func (b *BeaconchainAdapter) GetBlockHeader(blockID string) (*BlockHeaderResponse, error) {
	url := fmt.Sprintf("%s/eth/v1/beacon/headers/%s", b.beaconChainUrl, blockID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to send request to Beaconchain at %s: %w", url, err)
	}
	defer resp.Body.Close()

	var result BlockHeaderResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response for GetBlockHeader: %w", err)
	}

	return &result, nil
}

func (b *BeaconchainAdapter) GetEpochHeader(blockID string) (uint64, error) {
	header, err := b.GetBlockHeader(blockID)
	if err != nil {
		return 0, fmt.Errorf("failed to get block header for blockID %s: %w", blockID, err)
	}

	slot := header.Data.Header.Message.Slot
	epoch := getEpochFromSlot(slot)
	return epoch, nil
}

func getEpochFromSlot(slot string) uint64 {
	const slotsPerEpoch = 32
	slotInt := parseInt(slot)
	return slotInt / slotsPerEpoch
}

func parseInt(slot string) uint64 {
	var result uint64
	fmt.Sscanf(slot, "%d", &result)
	return result
}

// Response structs for deserializing JSON responses

// ValidatorsResponse holds the structure for the GetStateValidators response.
type ValidatorsResponse struct {
	ExecutionOptimistic bool `json:"execution_optimistic"`
	Finalized           bool `json:"finalized"`
	Data                []struct {
		Index     string `json:"index"`
		Balance   string `json:"balance"`
		Status    string `json:"status"`
		Validator struct {
			Pubkey                     string `json:"pubkey"`
			WithdrawalCredentials      string `json:"withdrawal_credentials"`
			EffectiveBalance           string `json:"effective_balance"`
			Slashed                    bool   `json:"slashed"`
			ActivationEligibilityEpoch string `json:"activation_eligibility_epoch"`
			ActivationEpoch            string `json:"activation_epoch"`
			ExitEpoch                  string `json:"exit_epoch"`
			WithdrawableEpoch          string `json:"withdrawable_epoch"`
		} `json:"validator"`
	} `json:"data"`
}

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
