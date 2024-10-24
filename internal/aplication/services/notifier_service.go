package services

import (
	"errors"
	"lido-events/internal/aplication/ports"
)

type NotifierService struct {
	notifier ports.NotifierPort
}

func NewNotifierService(notifier ports.NotifierPort) *NotifierService {
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
