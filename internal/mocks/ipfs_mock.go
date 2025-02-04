package mocks

import (
	"lido-events/internal/application/domain"

	"github.com/stretchr/testify/mock"
)

// MockIpfsPort is a mock implementation of the IpfsPort interface using testify/mock
type MockIpfsPort struct {
	mock.Mock
}

// FetchAndParseIpfs simulates fetching and parsing an IPFS content by CID
func (m *MockIpfsPort) FetchAndParseIpfs(cid string) (domain.OriginalReport, bool, error) {
	args := m.Called(cid)
	return args.Get(0).(domain.OriginalReport), false, args.Error(1)
}
