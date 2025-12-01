package ports

import (
	"lido-events/internal/application/domain"
)

type PerformanceStorage interface {
	SaveReport(logCid string, report *domain.Report) error
	GetNoPerformance(noID string) ([]*domain.Report, error)
}
