//go:build unit
// +build unit

package services_test

import (
	"context"
	"lido-events/internal/application/domain"
	"lido-events/internal/application/services"
	"lido-events/internal/mocks"
	"math/big"
	"sync"
	"testing"
	"time"
)

func TestLoadPendingHashes_WithComplexTypes(t *testing.T) {
	mockStorage := new(mocks.MockStoragePort)
	mockIpfs := new(mocks.MockIpfsPort)

	operatorIDs := []*big.Int{
		big.NewInt(1),
		big.NewInt(2),
	}

	pendingHashes := []string{
		"bafybeigdyrztf3zbd72ztfb7udai7n7yoqgzft3u34ud2ytlmqf27vnwae",
	}

	originalReport := domain.OriginalReport{
		Frame:     [2]int{100, 200},
		Threshold: 10.5,
		Operators: map[string]domain.Data{
			"1": {
				Distributed: 500,
				Stuck:       false,
				Validators: map[string]domain.Validator{
					"validator1": {
						Perf:    domain.Performance{Assigned: 10, Included: 9},
						Slashed: false,
					},
					"validator2": {
						Perf:    domain.Performance{Assigned: 8, Included: 8},
						Slashed: true,
					},
				},
			},
			"2": {
				Distributed: 1000,
				Stuck:       true,
				Validators: map[string]domain.Validator{
					"validator3": {
						Perf:    domain.Performance{Assigned: 15, Included: 14},
						Slashed: false,
					},
				},
			},
		},
	}

	report1 := domain.Report{
		Frame:     originalReport.Frame,
		Threshold: originalReport.Threshold,
		Data:      originalReport.Operators["1"],
	}

	report2 := domain.Report{
		Frame:     originalReport.Frame,
		Threshold: originalReport.Threshold,
		Data:      originalReport.Operators["2"],
	}

	// Set up mock expectations
	mockStorage.On("GetOperatorIds").Return(operatorIDs, nil)
	mockStorage.On("GetPendingHashes").Return(pendingHashes, nil)
	mockIpfs.On("FetchAndParseIpfs", pendingHashes[0]).Return(originalReport, nil)

	mockStorage.On("SaveReport", operatorIDs[0], report1).Return(nil)
	mockStorage.On("SaveReport", operatorIDs[1], report2).Return(nil)
	mockStorage.On("DeletePendingHash", pendingHashes[0]).Return(nil)

	// Create the PendingHashesLoader instance
	loader := services.NewPendingHashesLoader(mockStorage, mockIpfs)

	// Call the loadPendingHashes method
	err := loader.LoadPendingHashes()

	// Assertions
	if err != nil {
		t.Errorf("LoadPendingHashes failed: %v", err)
	}

	mockStorage.AssertExpectations(t)
	mockIpfs.AssertExpectations(t)
}

func TestLoadPendingHashesCron_WithComplexTypes(t *testing.T) {
	mockStorage := new(mocks.MockStoragePort)
	mockIpfs := new(mocks.MockIpfsPort)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := &sync.WaitGroup{}
	channel := make(chan struct{})

	operatorIDs := []*big.Int{
		big.NewInt(1),
	}

	pendingHashes := []string{
		"bafybeigdyrztf3zbd72ztfb7udai7n7yoqgzft3u34ud2ytlmqf27vnwae",
	}

	originalReport := domain.OriginalReport{
		Frame:     [2]int{100, 200},
		Threshold: 10.5,
		Operators: map[string]domain.Data{
			"1": {
				Distributed: 500,
				Stuck:       false,
				Validators: map[string]domain.Validator{
					"validator1": {
						Perf:    domain.Performance{Assigned: 10, Included: 9},
						Slashed: false,
					},
				},
			},
		},
	}

	report := domain.Report{
		Frame:     originalReport.Frame,
		Threshold: originalReport.Threshold,
		Data:      originalReport.Operators["1"],
	}

	// Set up mock expectations
	mockStorage.On("GetOperatorIds").Return(operatorIDs, nil)
	mockStorage.On("GetPendingHashes").Return(pendingHashes, nil)
	mockIpfs.On("FetchAndParseIpfs", pendingHashes[0]).Return(originalReport, nil)
	mockStorage.On("SaveReport", operatorIDs[0], report).Return(nil)
	mockStorage.On("DeletePendingHash", pendingHashes[0]).Return(nil)

	// Create the PendingHashesLoader instance
	loader := services.NewPendingHashesLoader(mockStorage, mockIpfs)

	// Run the LoadPendingHashesCron method
	go loader.LoadPendingHashesCron(ctx, 1*time.Second, wg, channel)
	time.Sleep(2 * time.Second) // Let the cron execute at least once
	cancel()                    // Stop the cron

	wg.Wait() // Wait for the goroutine to finish

	mockStorage.AssertExpectations(t)
	mockIpfs.AssertExpectations(t)
}
