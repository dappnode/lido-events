package ports

type NotifierPort interface {
	SendNotification(message string) error
	UpdateBotConfig() error
}
