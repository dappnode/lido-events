package services

type Notifier interface {
	Send(message string) error
}
