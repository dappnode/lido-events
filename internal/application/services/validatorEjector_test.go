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

func TestValidatorEjector_EjectValidator(t *testing.T) {
	mockStorage := new(mocks.MockStoragePort)
	mockNotifier := new(mocks.MockNotifierPort)
	mockExitValidator := new(mocks.MockExitValidator)
	mockBeaconchain := new(mocks.MockBeaconchain)

	operatorID := big.NewInt(1)
	validatorPubKey := "0x1234567890abcdef"
	validatorIndex := "101"

	exitRequest := domain.ExitRequest{
		Event: domain.VeboValidatorExitRequest{
			NodeOperatorId:  operatorID,
			ValidatorIndex:  big.NewInt(101),
			ValidatorPubkey: []byte(validatorPubKey),
			Timestamp:       big.NewInt(1234567890),
		},
		Status: domain.StatusActiveOngoing,
	}

	exitRequests := domain.ExitRequests{
		validatorIndex: exitRequest,
	}

	// Set up mock expectations
	mockStorage.On("GetOperatorIds").Return([]*big.Int{operatorID}, nil)
	mockStorage.On("GetExitRequests", operatorID.String()).Return(exitRequests, nil)
	mockStorage.On("UpdateExitRequestStatus", operatorID.String(), validatorPubKey, domain.StatusActiveExiting).Return(nil)

	mockNotifier.On("SendNotification", mock.Anything).Return(nil)
	mockExitValidator.On("ExitValidator", validatorPubKey, validatorIndex).Return(nil)
	mockBeaconchain.On("GetValidatorStatus", validatorPubKey).Return(domain.StatusActiveExiting, nil)

	// Create the ValidatorEjector instance
	ejector := services.NewValidatorEjectorService(mockStorage, mockNotifier, mockExitValidator, mockBeaconchain)

	// Call the ejectValidator method
	err := ejector.EjectValidator()

	// Assertions
	if err != nil {
		t.Errorf("EjectValidator failed: %v", err)
	}

	mockStorage.AssertExpectations(t)
	mockNotifier.AssertExpectations(t)
	mockExitValidator.AssertExpectations(t)
	mockBeaconchain.AssertExpectations(t)
}

func TestValidatorEjectorCron(t *testing.T) {
	mockStorage := new(mocks.MockStoragePort)
	mockNotifier := new(mocks.MockNotifierPort)
	mockExitValidator := new(mocks.MockExitValidator)
	mockBeaconchain := new(mocks.MockBeaconchain)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := &sync.WaitGroup{}

	operatorID := big.NewInt(1)
	validatorPubKey := "0x1234567890abcdef"
	validatorIndex := "101"

	exitRequest := domain.ExitRequest{
		Event: domain.VeboValidatorExitRequest{
			NodeOperatorId:  operatorID,
			ValidatorIndex:  big.NewInt(101),
			ValidatorPubkey: []byte(validatorPubKey),
			Timestamp:       big.NewInt(1234567890),
		},
		Status: domain.StatusActiveOngoing,
	}

	exitRequests := domain.ExitRequests{
		validatorIndex: exitRequest,
	}

	// Set up mock expectations
	mockStorage.On("GetOperatorIds").Return([]*big.Int{operatorID}, nil)
	mockStorage.On("GetExitRequests", operatorID.String()).Return(exitRequests, nil)
	mockStorage.On("UpdateExitRequestStatus", operatorID.String(), validatorPubKey, domain.StatusActiveExiting).Return(nil)

	mockNotifier.On("SendNotification", mock.Anything).Return(nil)
	mockExitValidator.On("ExitValidator", validatorPubKey, validatorIndex).Return(nil)
	mockBeaconchain.On("GetValidatorStatus", validatorPubKey).Return(domain.StatusActiveExiting, nil)

	// Create the ValidatorEjector instance
	ejector := services.NewValidatorEjectorService(mockStorage, mockNotifier, mockExitValidator, mockBeaconchain)

	// Run the ValidatorEjectorCron method
	go ejector.ValidatorEjectorCron(ctx, 1*time.Second, wg)
	time.Sleep(2 * time.Second) // Let the cron execute at least once
	cancel()                    // Stop the cron

	wg.Wait() // Wait for the goroutine to finish

	mockStorage.AssertExpectations(t)
	mockNotifier.AssertExpectations(t)
	mockExitValidator.AssertExpectations(t)
	mockBeaconchain.AssertExpectations(t)
}
