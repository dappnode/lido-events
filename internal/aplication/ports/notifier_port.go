package ports

type NotifierPort interface {
	Send(message string) error
}
