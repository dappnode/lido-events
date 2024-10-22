package domain

type IndexerRepository interface {
	UpdateTelegramToken(token string) error
	UpdateOperatorID(operatorID string) error
	GetLidoReport(start, end int) (map[string]Report, error)
	GetExitRequests() (map[string]string, error)
}
