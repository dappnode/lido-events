package vebo

import (
	"context"
	"encoding/json"
	"io"
	"lido-events/internal/adapters/vebo/bindings"
	"lido-events/internal/application/domain"
	"lido-events/internal/application/ports"
	"log"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type VeboAdapter struct {
	client          *ethclient.Client
	VeboAddress     common.Address
	StakingModuleId []*big.Int
	NodeOperatorId  []*big.Int
	ValidatorIndex  []*big.Int // TODO: where to get it?
	RefSlot         []*big.Int // TODO: where to get it?
	ipfsEndpoint    string // Add IPFS endpoint URL field
}

type IPFSData struct {
    Blockstamp struct {
        BlockHash      string `json:"block_hash"`
        BlockNumber    int    `json:"block_number"`
        BlockTimestamp int    `json:"block_timestamp"`
        RefEpoch       int    `json:"ref_epoch"`
        RefSlot        int    `json:"ref_slot"`
        SlotNumber     int    `json:"slot_number"`
        StateRoot      string `json:"state_root"`
    } `json:"blockstamp"`
    Distributable int64 `json:"distributable"`
    Frame         []int `json:"frame"`
    Operators     map[string]struct {
        Distributed int64 `json:"distributed"`
        Stuck       bool  `json:"stuck"`
        Validators  map[string]struct {
            Perf struct {
                Assigned int `json:"assigned"`
                Included int `json:"included"`
            } `json:"perf"`
            Slashed bool `json:"slashed"`
        } `json:"validators"`
    } `json:"operators"`
    Threshold float64 `json:"threshold"`
}


func NewVeboAdapter(
	wsURL string,
	veboAddress common.Address,
	stakingModuleId []*big.Int,
	nodeOperatorId []*big.Int,
	validatorIndex []*big.Int,
	refSlot []*big.Int,
	ipfsEndpoint string, 
) (*VeboAdapter, error) {
	client, err := ethclient.Dial(wsURL)
	if err != nil {
		return nil, err
	}

	return &VeboAdapter{
		client:          client,
		VeboAddress:     veboAddress,
		StakingModuleId: stakingModuleId,
		NodeOperatorId:  nodeOperatorId,
		ValidatorIndex:  validatorIndex,
		RefSlot:         refSlot,
		ipfsEndpoint:    ipfsEndpoint, // Set the IPFS endpoint
	}, nil
}

func (va *VeboAdapter) WatchVeboEvents(ctx context.Context, handlers ports.VeboHandlers) error {
	veboContract, err := bindings.NewVebo(va.VeboAddress, va.client)
	if err != nil {
		return err
	}

	// Watch for ValidatorExitRequest
	validatorExitRequestChan := make(chan *domain.VeboValidatorExitRequest)
	subExit, err := veboContract.WatchValidatorExitRequest(&bind.WatchOpts{Context: ctx}, validatorExitRequestChan, va.StakingModuleId, va.NodeOperatorId, va.ValidatorIndex)
	if err != nil {
		return err
	}

	// Watch for ReportSubmitted
	reportSubmittedChan := make(chan *domain.VeboReportSubmitted)
	subReport, err := veboContract.WatchReportSubmitted(&bind.WatchOpts{Context: ctx}, reportSubmittedChan, va.RefSlot)
	if err != nil {
		return err
	}

	go func() {
		defer subExit.Unsubscribe()
		defer subReport.Unsubscribe()
		for {
			select {
			case event := <-validatorExitRequestChan:
				if err := handlers.HandleValidatorExitRequestEvent(event); err != nil {
					log.Printf("Error handling ValidatorExitRequest: %v", err)
				}
			case event := <-reportSubmittedChan:
				if err := handlers.HandleReportSubmittedEvent(event); err != nil {
					log.Printf("Error handling ReportSubmitted: %v", err)
				}
			case err := <-subExit.Err():
				log.Printf("Subscription error (ValidatorExitRequest): %v", err)
				return
			case err := <-subReport.Err():
				log.Printf("Subscription error (ReportSubmitted): %v", err)
				return
			case <-ctx.Done():
				return
			}
		}
	}()

	return nil
}

// FetchAndParseIPFSContent fetches the IPFS content and parses it into a domain.Report struct.
func (va *VeboAdapter) FetchAndParseIPFSContent(ctx context.Context, ipfsHash string, operatorId domain.OperatorId) (*domain.Report, error) {
    url := va.ipfsEndpoint + "/api/v0/cat?arg=" + ipfsHash
    response, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer response.Body.Close()

    data, err := io.ReadAll(response.Body)
    if err != nil {
        return nil, err
    }

    var ipfsData IPFSData
    err = json.Unmarshal(data, &ipfsData)
    if err != nil {
        return nil, err
    }

    report := &domain.Report{
        Threshold: ipfsData.Threshold,
        Validators: make(domain.Validators),
		RefSlot: uint32(ipfsData.Blockstamp.RefSlot),
    }

	operatorKey := operatorId.String()
    if operator, ok := ipfsData.Operators[operatorKey]; ok {
        for pubKey, validator := range operator.Validators {
            report.Validators[pubKey] = domain.ValidatorPerformance{
                Included: uint(validator.Perf.Included),
                Assigned: uint(validator.Perf.Assigned),
            }
        }
    }

    return report, nil
}

