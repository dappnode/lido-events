package services

import (
	"errors"
	"lido-events/internal/aplication/ports"
)

type NotifierService struct {
	notifierAdapter ports.NotifierPort
}

func NewNotifierService(notifierAdapter ports.NotifierPort) *NotifierService {
	return &NotifierService{
		notifierAdapter,
	}
}

func (ns *NotifierService) SendNotification(message string) error {
	if ns.notifierAdapter == nil {
		return errors.New("no notifier available")
	}
	return ns.notifierAdapter.SendNotification(message)
}
