package services

import "lido-events/internal/domain"

type Service struct {
	Storage  domain.Storage  // Core Storage interface for data management
	Notifier domain.Notifier // Separate Notifier interface for notifications
	Config   domain.Config
}

func (s *Service) UpdateTelegramToken(token string) error {
	err := s.Storage.UpdateTelegramToken(token)
	if err != nil {
		return err
	}

	if s.Notifier != nil {
		return s.Notifier.Notify("Telegram token updated.")
	}
	return nil
}

func (s *Service) UpdateOperatorID(operatorID string) error {
	err := s.Storage.UpdateOperatorID(operatorID)
	if err != nil {
		return err
	}

	if s.Notifier != nil {
		return s.Notifier.Notify("Operator ID updated.")
	}
	return nil
}

func (s *Service) GetLidoReport(start, end int) (map[string]domain.Report, error) {
	return s.Storage.GetLidoReport(start, end)
}

func (s *Service) GetExitRequests() (map[string]string, error) {
	return s.Storage.GetExitRequests()
}
