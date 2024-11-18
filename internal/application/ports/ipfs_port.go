package ports

import "lido-events/internal/application/domain"

type IpfsPort interface {
	FetchAndParseIpfs(cid string) (domain.OriginalReport, error)
}
