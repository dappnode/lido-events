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

	"github.com/stretchr/testify/mock"
)

func TestHandleValidatorExitRequestEvent(t *testing.T) {
	mockStorage := new(mocks.MockStoragePort)
	mockNotifier := new(mocks.MockNotifierPort)
	mockBeaconchain := new(mocks.MockBeaconchain)

	event := &domain.VeboValidatorExitRequest{
		NodeOperatorId:  big.NewInt(1),
		ValidatorIndex:  big.NewInt(101),
		ValidatorPubkey: []byte("validator_pubkey"),
		Timestamp:       big.NewInt(1234567890),
	}

	exitRequest := domain.ExitRequest{
		Event:  *event,
		Status: domain.StatusActiveOngoing,
	}

	// Set up mock expectations
	mockBeaconchain.On("GetValidatorStatus", "validator_pubkey").Return(domain.StatusActiveOngoing, nil)
	mockStorage.On("SaveExitRequest", event.NodeOperatorId, event.ValidatorIndex.String(), exitRequest).Return(nil)
	mockNotifier.On("SendNotification", mock.Anything).Return(nil)

	scanner := services.NewValidatorExitRequestEventScanner(
		mockStorage, mockNotifier, nil, nil, mockBeaconchain, 0,
	)

	// Call the method
	err := scanner.HandleValidatorExitRequestEvent(event)

	// Assertions
	if err != nil {
		t.Errorf("HandleValidatorExitRequestEvent failed: %v", err)
	}

	mockBeaconchain.AssertExpectations(t)
	mockStorage.AssertExpectations(t)
	mockNotifier.AssertExpectations(t)
}

func TestScanValidatorExitRequestEventsCron(t *testing.T) {
	mockStorage := new(mocks.MockStoragePort)
	mockNotifier := new(mocks.MockNotifierPort)
	mockExecution := new(mocks.MockExecutionPort)
	mockVebo := new(mocks.MockVeboPort)
	mockBeaconchain := new(mocks.MockBeaconchain)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := &sync.WaitGroup{}

	start := uint64(10)
	end := uint64(20)

	event := &domain.VeboValidatorExitRequest{
		NodeOperatorId:  big.NewInt(1),
		ValidatorIndex:  big.NewInt(101),
		ValidatorPubkey: []byte("validator_pubkey"),
		Timestamp:       big.NewInt(1234567890),
	}

	exitRequest := domain.ExitRequest{
		Event:  *event,
		Status: domain.StatusActiveOngoing,
	}

	// Set up mock expectations
	mockStorage.On("GetValidatorExitRequestLastProcessedBlock").Return(start, nil)
	mockExecution.On("GetMostRecentBlockNumber").Return(end, nil)
	mockVebo.On("ScanVeboValidatorExitRequestEvent", mock.Anything, start, &end, mock.Anything).
		Run(func(args mock.Arguments) {
			handleFunc := args.Get(3).(func(*domain.VeboValidatorExitRequest) error)
			handleFunc(event)
		}).Return(nil)
	mockBeaconchain.On("GetValidatorStatus", "validator_pubkey").Return(domain.StatusActiveOngoing, nil)
	mockStorage.On("SaveExitRequest", event.NodeOperatorId, event.ValidatorIndex.String(), exitRequest).Return(nil)
	mockNotifier.On("SendNotification", mock.Anything).Return(nil)
	mockStorage.On("SaveValidatorExitRequestLastProcessedBlock", end).Return(nil)

	scanner := services.NewValidatorExitRequestEventScanner(
		mockStorage, mockNotifier, mockVebo, mockExecution, mockBeaconchain, 0,
	)

	// Run the method
	go scanner.ScanValidatorExitRequestEventsCron(ctx, 1*time.Second, wg)
	time.Sleep(2 * time.Second) // Let the cron execute at least once
	cancel()                    // Stop the cron

	wg.Wait() // Wait for the goroutine to finish

	mockStorage.AssertExpectations(t)
	mockExecution.AssertExpectations(t)
	mockVebo.AssertExpectations(t)
	mockBeaconchain.AssertExpectations(t)
	mockNotifier.AssertExpectations(t)
}
