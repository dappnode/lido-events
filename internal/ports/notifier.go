package ports

// Notifier defines the methods for sending notifications
type Notifier interface {
	Send(message string) error
}
