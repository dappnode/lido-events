package services

import (
	"context"
	"math/big"
	"testing"
	"time"

	"lido-events/internal/application/domain"
	"lido-events/internal/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	secondsPerSlotTest          = uint64(12)
	defaultAllowedExitDelayTest = uint64(345600) // 4 days in seconds
	exitDelayMultiplierTest     = uint64(2)
)

func newTestExitRequestEventScanner() (*ExitRequestEventScanner, *mocks.MockExitsStorage, *mocks.MockNotifierPort, *mocks.MockVeboPort, *mocks.MockExecutionPort, *mocks.MockBeaconchain, *mocks.MockCsParametersPort) {
	storage := new(mocks.MockExitsStorage)
	notifier := new(mocks.MockNotifierPort)
	vebo := new(mocks.MockVeboPort)
	execution := new(mocks.MockExecutionPort)
	beacon := new(mocks.MockBeaconchain)
	csParams := new(mocks.MockCsParametersPort)

	scanner := NewExitRequestEventScanner(
		storage,
		notifier,
		vebo,
		execution,
		beacon,
		csParams,
		secondsPerSlotTest,
		defaultAllowedExitDelayTest,
		exitDelayMultiplierTest,
	)

	return scanner, storage, notifier, vebo, execution, beacon, csParams
}

// RunScan tests

func TestRunScan_SkipsWhenNodesSyncing(t *testing.T) {
	t.Parallel()

	scanner, _, _, _, execution, beacon, _ := newTestExitRequestEventScanner()

	// execution or beaconchain syncing should cause RunScan to return error and not proceed
	execution.On("IsSyncing").Return(true, nil).Once()

	ctx := context.Background()
	err := scanner.RunScan(ctx)
	assert.Error(t, err)

	// ensure we did not even query beacon when execution is syncing
	beacon.AssertNotCalled(t, "GetSyncingStatus")
}

func TestRunScan_ProceedsWhenNodesNotSyncing(t *testing.T) {
	t.Parallel()

	scanner, storage, _, vebo, execution, beacon, _ := newTestExitRequestEventScanner()

	ctx := context.Background()

	// nodes not syncing
	execution.On("IsSyncing").Return(false, nil).Once()
	beacon.On("GetSyncingStatus").Return(false, nil).Once()

	// storage has one operator and a non-zero last processed block
	operatorID := big.NewInt(1)
	storage.On("GetOperatorIds").Return([]*big.Int{operatorID}, nil).Once()
	storage.On("GetValidatorExitRequestLastProcessedBlock", operatorID).Return(uint64(100), nil).Once()

	// execution provides most recent block
	execution.On("GetMostRecentBlockNumber").Return(uint64(200), nil).Once()

	// vebo scan succeeds
	vebo.On("ScanVeboValidatorExitRequestEvent", ctx, operatorID, uint64(100), mock.AnythingOfType("*uint64"), mock.AnythingOfType("func(*domain.VeboValidatorExitRequest) error")).Return(nil).Once()

	// saving last processed block succeeds
	storage.On("SaveValidatorExitRequestLastProcessedBlock", operatorID, uint64(200)).Return(nil).Once()

	err := scanner.RunScan(ctx)
	assert.NoError(t, err)

	storage.AssertExpectations(t)
	vebo.AssertExpectations(t)
	execution.AssertExpectations(t)
	beacon.AssertExpectations(t)
}

func TestRunScan_UsesAllowedExitDelayBlockWhenNoLastProcessed(t *testing.T) {
	t.Parallel()

	scanner, storage, _, vebo, execution, beacon, csParams := newTestExitRequestEventScanner()

	ctx := context.Background()

	// nodes not syncing
	execution.On("IsSyncing").Return(false, nil).Once()
	beacon.On("GetSyncingStatus").Return(false, nil).Once()

	// one operator with no last processed block (0)
	operatorID := big.NewInt(1)
	storage.On("GetOperatorIds").Return([]*big.Int{operatorID}, nil).Once()
	storage.On("GetValidatorExitRequestLastProcessedBlock", operatorID).Return(uint64(0), nil).Once()

	// computeAllowedExitDelayBlock dependencies
	// csParameters: return some allowed delay (use defaultAllowedExitDelayTest here)
	csParams.On("GetDefaultAllowedExitDelay", mock.Anything).Return(defaultAllowedExitDelayTest, nil).Once()

	// beacon genesis time: choose a fixed value
	genesisTime := uint64(time.Now().Unix()) - 10*defaultAllowedExitDelayTest
	beacon.On("GetGenesisTime").Return(genesisTime, nil).Once()

	// for simplicity, return a fixed block number from GetBlockNumber
	// we don't assert exact math here, just that it's used as start block
	beacon.On("GetBlockNumber", mock.AnythingOfType("string")).Return(uint64(150), nil).Once()

	// ensureBlockHasReceipts succeeds
	execution.On("GetBlockHasReceipts", mock.AnythingOfType("string")).Return(true, nil).Once()

	// latest block
	execution.On("GetMostRecentBlockNumber").Return(uint64(300), nil).Once()

	// vebo scan should start from the computed allowed block (150)
	vebo.On("ScanVeboValidatorExitRequestEvent", ctx, operatorID, uint64(150), mock.AnythingOfType("*uint64"), mock.AnythingOfType("func(*domain.VeboValidatorExitRequest) error")).Return(nil).Once()

	storage.On("SaveValidatorExitRequestLastProcessedBlock", operatorID, uint64(300)).Return(nil).Once()

	err := scanner.RunScan(ctx)
	assert.NoError(t, err)
}

func TestRunScan_PropagatesErrorsFromDependencies(t *testing.T) {
	t.Parallel()

	// We will check three error paths: GetMostRecentBlockNumber, ScanVeboValidatorExitRequestEvent, SaveValidatorExitRequestLastProcessedBlock
	t.Run("error from GetMostRecentBlockNumber", func(t *testing.T) {
		scanner, storage, _, vebo, execution, beacon, _ := newTestExitRequestEventScanner()

		ctx := context.Background()
		execution.On("IsSyncing").Return(false, nil).Once()
		beacon.On("GetSyncingStatus").Return(false, nil).Once()

		operatorID := big.NewInt(1)
		storage.On("GetOperatorIds").Return([]*big.Int{operatorID}, nil).Once()
		storage.On("GetValidatorExitRequestLastProcessedBlock", operatorID).Return(uint64(100), nil).Once()

		execution.On("GetMostRecentBlockNumber").Return(uint64(0), assert.AnError).Once()

		err := scanner.RunScan(ctx)
		assert.Error(t, err)

		vebo.AssertNotCalled(t, "ScanVeboValidatorExitRequestEvent", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
	})

	t.Run("error from ScanVeboValidatorExitRequestEvent", func(t *testing.T) {
		scanner, storage, _, vebo, execution, beacon, _ := newTestExitRequestEventScanner()

		ctx := context.Background()
		execution.On("IsSyncing").Return(false, nil).Once()
		beacon.On("GetSyncingStatus").Return(false, nil).Once()

		operatorID := big.NewInt(1)
		storage.On("GetOperatorIds").Return([]*big.Int{operatorID}, nil).Once()
		storage.On("GetValidatorExitRequestLastProcessedBlock", operatorID).Return(uint64(100), nil).Once()

		execution.On("GetMostRecentBlockNumber").Return(uint64(200), nil).Once()

		vebo.On("ScanVeboValidatorExitRequestEvent", ctx, operatorID, uint64(100), mock.AnythingOfType("*uint64"), mock.AnythingOfType("func(*domain.VeboValidatorExitRequest) error")).Return(assert.AnError).Once()

		err := scanner.RunScan(ctx)
		assert.Error(t, err)
	})

	t.Run("error from SaveValidatorExitRequestLastProcessedBlock", func(t *testing.T) {
		scanner, storage, _, vebo, execution, beacon, _ := newTestExitRequestEventScanner()

		ctx := context.Background()
		execution.On("IsSyncing").Return(false, nil).Once()
		beacon.On("GetSyncingStatus").Return(false, nil).Once()

		operatorID := big.NewInt(1)
		storage.On("GetOperatorIds").Return([]*big.Int{operatorID}, nil).Once()
		storage.On("GetValidatorExitRequestLastProcessedBlock", operatorID).Return(uint64(100), nil).Once()

		execution.On("GetMostRecentBlockNumber").Return(uint64(200), nil).Once()

		vebo.On("ScanVeboValidatorExitRequestEvent", ctx, operatorID, uint64(100), mock.AnythingOfType("*uint64"), mock.AnythingOfType("func(*domain.VeboValidatorExitRequest) error")).Return(nil).Once()

		storage.On("SaveValidatorExitRequestLastProcessedBlock", operatorID, uint64(200)).Return(assert.AnError).Once()

		err := scanner.RunScan(ctx)
		assert.Error(t, err)
	})
}

// computeAllowedExitDelayBlock tests

func TestComputeAllowedExitDelayBlock_FallbackToDefaultDelayOnError(t *testing.T) {
	t.Parallel()

	scanner, _, _, _, execution, beacon, csParams := newTestExitRequestEventScanner()

	ctx := context.Background()

	// simulate error from csParameters; scanner should fallback to defaultAllowedExitDelayTest
	csParams.On("GetDefaultAllowedExitDelay", mock.Anything).Return(uint64(0), assert.AnError).Once()

	// fixed now for deterministic behaviour
	now := uint64(time.Now().Unix())
	// we cannot inject time.Now(), but we can still ensure the rest of the flow works

	// beacon genesis time before now so that slots are positive
	genesisTime := now - 10*defaultAllowedExitDelayTest
	beacon.On("GetGenesisTime").Return(genesisTime, nil).Once()

	// any block number for computed slot
	beacon.On("GetBlockNumber", mock.AnythingOfType("string")).Return(uint64(123), nil).Once()

	// ensureBlockHasReceipts: receipts exist
	execution.On("GetBlockHasReceipts", mock.AnythingOfType("string")).Return(true, nil).Once()

	block, err := scanner.computeAllowedExitDelayBlock(ctx)
	assert.NoError(t, err)
	assert.Equal(t, uint64(123), block)
}

func TestComputeAllowedExitDelayBlock_ErrorWhenEnsureBlockHasNoReceipts(t *testing.T) {
	t.Parallel()

	scanner, _, notifier, _, execution, beacon, csParams := newTestExitRequestEventScanner()

	ctx := context.Background()

	// normal delay
	csParams.On("GetDefaultAllowedExitDelay", mock.Anything).Return(defaultAllowedExitDelayTest, nil).Once()

	now := uint64(time.Now().Unix())
	genesisTime := now - 10*defaultAllowedExitDelayTest
	beacon.On("GetGenesisTime").Return(genesisTime, nil).Once()

	beacon.On("GetBlockNumber", mock.AnythingOfType("string")).Return(uint64(456), nil).Once()

	// receipts missing for that block
	execution.On("GetBlockHasReceipts", mock.AnythingOfType("string")).Return(false, nil).Once()

	// expect notification about missing log receipts
	notifier.On("SendMissingLogReceiptsNotification", mock.AnythingOfType("string")).Return(nil).Once()

	block, err := scanner.computeAllowedExitDelayBlock(ctx)
	assert.Error(t, err)
	assert.Equal(t, uint64(0), block)
}

// ensureBlockHasReceipts tests

func TestEnsureBlockHasReceipts_SendsNotificationWhenMissing(t *testing.T) {
	t.Parallel()

	scanner, _, notifier, _, execution, _, _ := newTestExitRequestEventScanner()

	// block without receipts
	execution.On("GetBlockHasReceipts", mock.AnythingOfType("string")).Return(false, nil).Once()

	notifier.On("SendMissingLogReceiptsNotification", mock.AnythingOfType("string")).Return(nil).Once()

	err := scanner.ensureBlockHasReceipts(123)
	assert.Error(t, err)

	notifier.AssertExpectations(t)
}

// HandleValidatorExitRequestEvent tests

func TestHandleValidatorExitRequestEvent_SendsNotificationForActiveStatuses(t *testing.T) {
	t.Parallel()

	statuses := []domain.ValidatorStatus{domain.StatusActiveOngoing, domain.StatusActiveSlashed}

	for _, st := range statuses {
		st := st
		t.Run(string(st), func(t *testing.T) {
			t.Parallel()

			scanner, storage, notifier, _, _, beacon, _ := newTestExitRequestEventScanner()

			// validator status returned from beacon
			beacon.On("GetValidatorStatus", mock.AnythingOfType("string")).Return(st, nil).Once()

			// expect notification
			notifier.On("SendValidatorExitRequestedNotification", mock.AnythingOfType("string")).Return(nil).Once()

			// expect SaveExitRequest
			operatorID := big.NewInt(1)
			validatorIndex := "123"
			storage.On("SaveExitRequest", operatorID, validatorIndex, mock.AnythingOfType("domain.ExitRequest")).Return(nil).Once()

			// build event
			event := &domain.VeboValidatorExitRequest{
				NodeOperatorId:  operatorID,
				ValidatorIndex:  big.NewInt(123),
				ValidatorPubkey: []byte{0x01, 0x02},
			}

			err := scanner.HandleValidatorExitRequestEvent(event)
			assert.NoError(t, err)

			storage.AssertExpectations(t)
			notifier.AssertExpectations(t)
			beacon.AssertExpectations(t)
		})
	}
}

func TestHandleValidatorExitRequestEvent_SavesExitRequestForNonActiveStatusWithoutNotification(t *testing.T) {
	t.Parallel()

	scanner, storage, notifier, _, _, beacon, _ := newTestExitRequestEventScanner()

	status := domain.StatusExitedUnslashed
	beacon.On("GetValidatorStatus", mock.AnythingOfType("string")).Return(status, nil).Once()

	// no notification should be sent
	notifier.AssertNotCalled(t, "SendValidatorExitRequestedNotification", mock.Anything)

	operatorID := big.NewInt(1)
	validatorIndex := "123"
	storage.On("SaveExitRequest", operatorID, validatorIndex, mock.AnythingOfType("domain.ExitRequest")).Return(nil).Once()

	event := &domain.VeboValidatorExitRequest{
		NodeOperatorId:  operatorID,
		ValidatorIndex:  big.NewInt(123),
		ValidatorPubkey: []byte{0x01, 0x02},
	}

	err := scanner.HandleValidatorExitRequestEvent(event)
	assert.NoError(t, err)

	storage.AssertExpectations(t)
	beacon.AssertExpectations(t)
}
