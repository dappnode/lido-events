package csparameters

import (
	"context"
	"fmt"

	"lido-events/internal/adapters/csparameters/bindings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// CsParameters is a thin adapter around the CSParametersRegistry contract bindings.
type CsParameters struct {
	rpcClient           *ethclient.Client
	csParametersAddress common.Address
}

// NewCsParameters creates a new CsParameters adapter instance.
func NewCsParameters(
	rpcClient *ethclient.Client,
	csParametersAddress common.Address,
) (*CsParameters, error) {
	return &CsParameters{
		rpcClient:           rpcClient,
		csParametersAddress: csParametersAddress,
	}, nil
}

// GetDefaultAllowedExitDelay retrieves the defaultAllowedExitDelay value from the
// CSParametersRegistry contract.
func (cs *CsParameters) GetDefaultAllowedExitDelay(ctx context.Context) (uint64, error) {
	contract, err := bindings.NewBindings(cs.csParametersAddress, cs.rpcClient)
	if err != nil {
		return 0, fmt.Errorf("failed to create CsParameters contract instance: %w", err)
	}

	delay, err := contract.DefaultAllowedExitDelay(nil)
	if err != nil {
		return 0, fmt.Errorf("failed to get defaultAllowedExitDelay: %w", err)
	}

	return delay.Uint64(), nil
}
