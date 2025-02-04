package ports

import (
	"lido-events/internal/application/domain"
	"time"
)

type IpfsPort interface {
	FetchAndParseIpfs(cid string, timeout *time.Duration) (domain.OriginalReport, bool, error)
}
