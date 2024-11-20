//go:build integration
// +build integration

package beaconchain_test

import (
	"os"
	"testing"

	"lido-events/internal/adapters/beaconchain"
	"lido-events/internal/application/domain"

	"github.com/stretchr/testify/assert"
)

// setupBeaconchainAdapter initializes the BeaconchainAdapter
func setupBeaconchainAdapter(t *testing.T) *beaconchain.BeaconchainAdapter {
	beaconchainURL := os.Getenv("BEACONCHAIN_URL")
	if beaconchainURL == "" {
		t.Fatal("BEACONCHAIN_URL environment variable not set")
	}

	return beaconchain.NewBeaconchainAdapter(beaconchainURL)
}

func TestGetValidatorStatusIntegration(t *testing.T) {
	adapter := setupBeaconchainAdapter(t)

	// Test an active validator
	pubkeyActive := "0x800000b3884235f70b06fec68c19642fc9e81e34fbe7f1c0ae156b8b45860dfe5ac71037ae561c2a759ba83401488e18"
	status, err := adapter.GetValidatorStatus(pubkeyActive)
	assert.NoError(t, err)
	assert.Equal(t, domain.StatusActiveOngoing, status)

	// Test a slashed validator
	pubkeySlashed := "0x8d6f381707c822288503112982628ff01cf9d6fa7753bd6b8a9909db5bd7d0b44b36709e4142e137dbed4b0c1bd23b2c"
	status, err = adapter.GetValidatorStatus(pubkeySlashed)
	assert.NoError(t, err)
	assert.Equal(t, domain.StatusExitedSlashed, status)
}

func TestPostStateValidatorsIntegration(t *testing.T) {
	adapter := setupBeaconchainAdapter(t)

	// Test with active and slashed validators
	ids := []string{
		"0x800000b3884235f70b06fec68c19642fc9e81e34fbe7f1c0ae156b8b45860dfe5ac71037ae561c2a759ba83401488e18",
		"0x8d6f381707c822288503112982628ff01cf9d6fa7753bd6b8a9909db5bd7d0b44b36709e4142e137dbed4b0c1bd23b2c",
	}
	response, err := adapter.PostStateValidators("finalized", ids, nil)
	assert.NoError(t, err)
	assert.Len(t, response.Data, 2)

	// Validate the statuses
	assert.Equal(t, domain.StatusActiveOngoing, domain.ValidatorStatus(response.Data[0].Status))
	assert.Equal(t, domain.StatusExitedSlashed, domain.ValidatorStatus(response.Data[1].Status))
}

func TestSubmitPoolVoluntaryExitIntegration(t *testing.T) {
	adapter := setupBeaconchainAdapter(t)

	// Example voluntary exit submission (modify epoch and validatorIndex accordingly)
	err := adapter.SubmitPoolVoluntaryExit("10", "886680", "0x...")
	assert.Error(t, err) // Should error since we're not providing a valid signature
}

func TestGetStateForkIntegration(t *testing.T) {
	adapter := setupBeaconchainAdapter(t)

	// Retrieve fork for "head"
	response, err := adapter.GetStateFork("head")
	assert.NoError(t, err)
	assert.NotEmpty(t, response.Data.CurrentVersion)
	t.Logf("State fork version: %s", response.Data.CurrentVersion)
}

func TestGetGenesisIntegration(t *testing.T) {
	adapter := setupBeaconchainAdapter(t)

	// Retrieve genesis data
	response, err := adapter.GetGenesis()
	assert.NoError(t, err)
	assert.NotEmpty(t, response.Data.GenesisTime)
	t.Logf("Genesis time: %s", response.Data.GenesisTime)
}

func TestGetBlockHeaderIntegration(t *testing.T) {
	adapter := setupBeaconchainAdapter(t)

	// Retrieve block header for "finalized"
	response, err := adapter.GetBlockHeader("finalized")
	assert.NoError(t, err)
	assert.NotEmpty(t, response.Data.Header.Message.Slot)
	t.Logf("Finalized block slot: %s", response.Data.Header.Message.Slot)
}

func TestGetEpochHeaderIntegration(t *testing.T) {
	adapter := setupBeaconchainAdapter(t)

	// Retrieve epoch header for "finalized"
	epoch, err := adapter.GetEpochHeader("finalized")
	assert.NoError(t, err)
	assert.Greater(t, epoch, uint64(0))
	t.Logf("Epoch for finalized block: %d", epoch)
}
