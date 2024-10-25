package contractbindings

import (
	"fmt"

	csaccounting "lido-events/internal/adapters/subscriber/bindings/CSAccounting"
	csfeedistributor "lido-events/internal/adapters/subscriber/bindings/CSFeeDistributor"
	csmodule "lido-events/internal/adapters/subscriber/bindings/CSModule"
	vebo "lido-events/internal/adapters/subscriber/bindings/VEBO"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ContractBindings struct {
	CSAccounting     *csaccounting.Csaccounting
	CSFeeDistributor *csfeedistributor.Csfeedistributor
	CSModule         *csmodule.Csmodule
	VEBO             *vebo.Vebo

	// Store addresses for each contract
	CSAccountingAddress     common.Address
	CSFeeDistributorAddress common.Address
	CSModuleAddress         common.Address
	VEBOAddress             common.Address
}

// InitializeContracts initializes all contract bindings with the provided addresses.
func InitializeContracts(
	client *ethclient.Client,
	csAccountingAddress, csFeeDistributorAddress, csModuleAddress, veboAddress string,
) (*ContractBindings, error) {

	// Initialize CSAccounting
	csAccountingInstance, err := csaccounting.NewCsaccounting(common.HexToAddress(csAccountingAddress), client)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize CSAccounting: %w", err)
	}

	// Initialize CSFeeDistributor
	csFeeDistributorInstance, err := csfeedistributor.NewCsfeedistributor(common.HexToAddress(csFeeDistributorAddress), client)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize CSFeeDistributor: %w", err)
	}

	// Initialize CSModule
	csModuleInstance, err := csmodule.NewCsmodule(common.HexToAddress(csModuleAddress), client)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize CSModule: %w", err)
	}

	// Initialize VEBO
	veboInstance, err := vebo.NewVebo(common.HexToAddress(veboAddress), client)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize VEBO: %w", err)
	}

	return &ContractBindings{
		CSAccounting:            csAccountingInstance,
		CSFeeDistributor:        csFeeDistributorInstance,
		CSModule:                csModuleInstance,
		VEBO:                    veboInstance,
		CSAccountingAddress:     common.HexToAddress(csAccountingAddress),
		CSFeeDistributorAddress: common.HexToAddress(csFeeDistributorAddress),
		CSModuleAddress:         common.HexToAddress(csModuleAddress),
		VEBOAddress:             common.HexToAddress(veboAddress),
	}, nil
}
