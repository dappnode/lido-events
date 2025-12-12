package services

import (
	"context"
	"errors"
	"testing"

	"lido-events/internal/application/domain"
	"lido-events/internal/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// 1. If there is an existing logcid already on db it does not get written.
func TestLoadHashes_SkipsExistingCid(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	mockPerf := new(mocks.MockPerformanceStorage)
	mockNotifier := new(mocks.MockNotifierPort)
	mockCsFee := new(mocks.MockCsFeeDistributor)
	mockIpfs := new(mocks.MockIpfsPort)

	// Existing CID in DB
	existing := []string{"cid1"}
	mockPerf.On("GetUniqueLogCids").Return(existing, nil).Once()

	// CsFeeDistributor returns the same CID
	mockCsFee.On("GetAllLogCids", ctx).Return([]string{"cid1"}, nil).Once()

	loader := NewAllHashesLoader(mockPerf, mockNotifier, mockCsFee, mockIpfs)

	err := loader.loadHashes(ctx)
	assert.NoError(t, err)

	// Ensure GetUniqueLogCids and GetAllLogCids were called
	mockPerf.AssertCalled(t, "GetUniqueLogCids")
	mockCsFee.AssertCalled(t, "GetAllLogCids", ctx)

	// Ensure SaveReport and IPFS are not called because CID already exists
	mockPerf.AssertNotCalled(t, "SaveReport", "cid1", mock.Anything)
	mockIpfs.AssertNotCalled(t, "FetchAndParseIpfs", "cid1")
}

// 2. If there is no CID in DB then it should be written to DB.
func TestLoadHashes_WritesNewCid(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	mockPerf := new(mocks.MockPerformanceStorage)
	mockNotifier := new(mocks.MockNotifierPort)
	mockCsFee := new(mocks.MockCsFeeDistributor)
	mockIpfs := new(mocks.MockIpfsPort)

	// No existing CIDs in DB
	mockPerf.On("GetUniqueLogCids").Return([]string{}, nil).Once()

	// CsFeeDistributor returns one new CID
	mockCsFee.On("GetAllLogCids", ctx).Return([]string{"cid1"}, nil).Once()

	report := domain.Report{}
	mockIpfs.On("FetchAndParseIpfs", "cid1").Return(report, nil).Once()

	// Expect SaveReport to be called once with cid1
	mockPerf.On("SaveReport", "cid1", mock.AnythingOfType("*domain.Report")).Return(nil).Once()

	loader := NewAllHashesLoader(mockPerf, mockNotifier, mockCsFee, mockIpfs)

	err := loader.loadHashes(ctx)
	assert.NoError(t, err)

	mockPerf.AssertExpectations(t)
	mockCsFee.AssertExpectations(t)
	mockIpfs.AssertExpectations(t)
}

// 3. If there are no log CIDs from CsFeeDistributor, then DB should not be written.
func TestLoadHashes_NoLogCids_NoWrites(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	mockPerf := new(mocks.MockPerformanceStorage)
	mockNotifier := new(mocks.MockNotifierPort)
	mockCsFee := new(mocks.MockCsFeeDistributor)
	mockIpfs := new(mocks.MockIpfsPort)

	// No existing CIDs in DB
	mockPerf.On("GetUniqueLogCids").Return([]string{}, nil).Once()

	// CsFeeDistributor returns no CIDs
	mockCsFee.On("GetAllLogCids", ctx).Return([]string{}, nil).Once()

	loader := NewAllHashesLoader(mockPerf, mockNotifier, mockCsFee, mockIpfs)

	err := loader.loadHashes(ctx)
	assert.NoError(t, err)

	// No writes or IPFS fetches should occur
	mockPerf.AssertNotCalled(t, "SaveReport", mock.Anything, mock.Anything)
	mockIpfs.AssertNotCalled(t, "FetchAndParseIpfs", mock.Anything)
}

// 4. Additional: when fetching existing CIDs from performance storage fails,
// loadHashes should return the error and not proceed.
func TestLoadHashes_ErrorOnGetUniqueLogCids(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	mockPerf := new(mocks.MockPerformanceStorage)
	mockNotifier := new(mocks.MockNotifierPort)
	mockCsFee := new(mocks.MockCsFeeDistributor)
	mockIpfs := new(mocks.MockIpfsPort)

	expectedErr := errors.New("db failure")
	// Return a typed nil slice to satisfy the mocked signature
	mockPerf.On("GetUniqueLogCids").Return([]string(nil), expectedErr).Once()

	loader := NewAllHashesLoader(mockPerf, mockNotifier, mockCsFee, mockIpfs)

	err := loader.loadHashes(ctx)
	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)

	// Ensure we did not attempt to call CsFeeDistributor or IPFS
	mockCsFee.AssertNotCalled(t, "GetAllLogCids", mock.Anything)
	mockIpfs.AssertNotCalled(t, "FetchAndParseIpfs", mock.Anything)
}
