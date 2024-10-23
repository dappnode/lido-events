package service

import (
	"errors"
	"lido-events/internal/ports"
)

type NotifierService struct {
	notifier ports.Notifier
}

func NewNotifierService(notifier ports.Notifier) *NotifierService {
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
