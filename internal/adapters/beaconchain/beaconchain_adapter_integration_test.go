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
	t.Skip()
	adapter := setupBeaconchainAdapter(t)

	// Test a slashed validator
	pubkeyExited := "0x972255d9a5085d082d485f1e17999b38967e022057aba66a477cd93bce5cfa980bc42df82b208987ed46b9cdbc7b5fcb"
	status, err := adapter.GetValidatorStatus(pubkeyExited)
	assert.NoError(t, err)
	assert.Equal(t, domain.StatusExitedUnslashed, status)
}

func TestPostStateValidatorsIntegration(t *testing.T) {
	t.Skip()
	adapter := setupBeaconchainAdapter(t)

	// Test with active and slashed validators
	ids := []string{
		"1802081",
	}
	response, err := adapter.PostStateValidators("finalized", ids, nil)
	assert.NoError(t, err)
	assert.Len(t, response.Data, 1)

	// Validate the statuses
	assert.Equal(t, domain.StatusExitedUnslashed, domain.ValidatorStatus(response.Data[0].Status))
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
