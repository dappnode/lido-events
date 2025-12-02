package csfeedistributor

import (
	"context"
	"fmt"
	"lido-events/internal/adapters/csFeeDistributor/bindings"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type CsFeeDistributor struct {
	rpcClient               *ethclient.Client
	csFeeDistributorAddress common.Address
}

func NewCsFeeDistributor(
	rpcClient *ethclient.Client,
	csFeeDistributorAddress common.Address,
) (*CsFeeDistributor, error) {
	return &CsFeeDistributor{
		rpcClient:               rpcClient,
		csFeeDistributorAddress: csFeeDistributorAddress,
	}, nil
}

// GetAllLogCids retrieves all historical log CIDs from the CsFeeDistributor contract.
func (cs *CsFeeDistributor) GetAllLogCids(ctx context.Context) ([]string, error) {
	contract, err := bindings.NewBindings(cs.csFeeDistributorAddress, cs.rpcClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create CsFeeDistributor contract instance: %w", err)
	}

	// Get total number of historical distribution entries
	count, err := contract.DistributionDataHistoryCount(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get DistributionDataHistoryCount: %w", err)
	}

	if count.Sign() == 0 {
		return []string{}, nil
	}

	logCids := make([]string, 0, count.Int64())

	for i := int64(0); i < count.Int64(); i++ {
		index := big.NewInt(i)
		data, err := contract.GetHistoricalDistributionData(nil, index)
		if err != nil {
			return nil, fmt.Errorf("failed to get historical distribution data for index %d: %w", i, err)
		}

		logCids = append(logCids, data.LogCid)
	}

	return logCids, nil
}
