//go:build unit
// +build unit

package services_test

import (
	"context"
	"fmt"
	"lido-events/internal/application/domain"
	"lido-events/internal/application/services"
	"lido-events/internal/mocks"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
)

func TestHandleDistributionLogUpdatedEvent(t *testing.T) {
	mockStorage := new(mocks.MockStoragePort)
	mockNotifier := new(mocks.MockNotifierPort)

	event := &domain.BindingsDistributionLogUpdated{LogCid: "test-log-cid"}

	// Set up mock expectations
	mockStorage.On("AddPendingHash", event.LogCid).Return(nil)
	mockNotifier.On("SendNotification", mock.Anything).Return(nil)

	scanner := services.NewDistributionLogUpdatedEventScanner(
		mockStorage, mockNotifier, nil, nil, 0,
	)

	// Call the method
	err := scanner.HandleDistributionLogUpdatedEvent(event)

	// Assertions
	if err != nil {
		t.Errorf("handleDistributionLogUpdatedEvent failed: %v", err)
	}

	mockStorage.AssertExpectations(t)
	mockNotifier.AssertExpectations(t)
}

func TestHandleDistributionLogUpdatedEvent_AddPendingHashFails(t *testing.T) {
	mockStorage := new(mocks.MockStoragePort)
	mockNotifier := new(mocks.MockNotifierPort)

	event := &domain.BindingsDistributionLogUpdated{LogCid: "test-log-cid"}

	// Set up mock expectations
	mockStorage.On("AddPendingHash", event.LogCid).Return(fmt.Errorf("mock error"))

	scanner := services.NewDistributionLogUpdatedEventScanner(
		mockStorage, mockNotifier, nil, nil, 0,
	)

	// Call the method
	err := scanner.HandleDistributionLogUpdatedEvent(event)

	// Assertions
	if err == nil {
		t.Error("Expected error, got nil")
	}

	mockStorage.AssertExpectations(t)
}

func TestScanDistributionLogUpdatedEventsCron(t *testing.T) {
	mockStorage := new(mocks.MockStoragePort)
	mockNotifier := new(mocks.MockNotifierPort)
	mockExecution := new(mocks.MockExecutionPort)
	mockDistributor := new(mocks.MockCsFeeDistributorImplPort)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := &sync.WaitGroup{}
	channel := make(chan struct{})

	start := uint64(100)
	end := uint64(200)

	// Set up mock expectations
	mockStorage.On("GetDistributionLogLastProcessedBlock").Return(start, nil)
	mockExecution.On("GetMostRecentBlockNumber").Return(end, nil)
	mockDistributor.On("ScanDistributionLogUpdatedEvents", mock.Anything, start, &end, mock.Anything).
		Run(func(args mock.Arguments) {
			handleFunc := args.Get(3).(func(*domain.BindingsDistributionLogUpdated) error)
			handleFunc(&domain.BindingsDistributionLogUpdated{LogCid: "test-log-cid"})
		}).
		Return(nil)
	mockStorage.On("SaveDistributionLogLastProcessedBlock", end).Return(nil)
	mockStorage.On("AddPendingHash", "test-log-cid").Return(nil)
	mockNotifier.On("SendNotification", mock.Anything).Return(nil)

	scanner := services.NewDistributionLogUpdatedEventScanner(
		mockStorage, mockNotifier, mockExecution, mockDistributor, 0,
	)

	// Run the method
	go scanner.ScanDistributionLogUpdatedEventsCron(ctx, 1*time.Second, wg, channel)
	time.Sleep(2 * time.Second) // Let the cron execute at least once
	cancel()                    // Stop the cron

	wg.Wait() // Wait for the goroutine to finish

	mockStorage.AssertExpectations(t)
	mockExecution.AssertExpectations(t)
	mockDistributor.AssertExpectations(t)
	mockNotifier.AssertExpectations(t)
}

func TestScanDistributionLogUpdatedEventsCron_NoLastProcessedBlock(t *testing.T) {
	mockStorage := new(mocks.MockStoragePort)
	mockNotifier := new(mocks.MockNotifierPort)
	mockExecution := new(mocks.MockExecutionPort)
	mockDistributor := new(mocks.MockCsFeeDistributorImplPort)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := &sync.WaitGroup{}
	channel := make(chan struct{})

	deploymentBlock := uint64(50)
	end := uint64(200)

	// Set up mock expectations
	mockStorage.On("GetDistributionLogLastProcessedBlock").Return(uint64(0), nil)
	mockExecution.On("GetMostRecentBlockNumber").Return(end, nil)
	mockDistributor.On("ScanDistributionLogUpdatedEvents", mock.Anything, deploymentBlock, &end, mock.Anything).
		Run(func(args mock.Arguments) {
			handleFunc := args.Get(3).(func(*domain.BindingsDistributionLogUpdated) error)
			handleFunc(&domain.BindingsDistributionLogUpdated{LogCid: "test-log-cid"})
		}).
		Return(nil)
	mockStorage.On("SaveDistributionLogLastProcessedBlock", end).Return(nil)
	mockStorage.On("AddPendingHash", "test-log-cid").Return(nil)
	mockNotifier.On("SendNotification", mock.Anything).Return(nil)

	scanner := services.NewDistributionLogUpdatedEventScanner(
		mockStorage, mockNotifier, mockExecution, mockDistributor, deploymentBlock,
	)

	// Run the method
	go scanner.ScanDistributionLogUpdatedEventsCron(ctx, 1*time.Second, wg, channel)
	time.Sleep(2 * time.Second) // Let the cron execute at least once
	cancel()                    // Stop the cron

	wg.Wait() // Wait for the goroutine to finish

	mockStorage.AssertExpectations(t)
	mockExecution.AssertExpectations(t)
	mockDistributor.AssertExpectations(t)
	mockNotifier.AssertExpectations(t)
}
