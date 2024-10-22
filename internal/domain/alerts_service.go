package domain

type AlertService interface {
	Notify(message string) error
}
