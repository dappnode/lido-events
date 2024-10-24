package services

import (
	"errors"
	"lido-events/internal/domain/services"
)

type NotifierService struct {
	notifier services.Notifier
}

func NewNotifierService(notifier services.Notifier) *NotifierService {
	return &NotifierService{
		notifier: notifier,
	}
}

func (ns *NotifierService) Send(message string) error {
	if ns.notifier == nil {
		return errors.New("no notifier available")
	}
	return ns.notifier.Send(message)
}
