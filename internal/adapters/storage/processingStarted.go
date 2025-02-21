package storage

import (
	"lido-events/internal/application/domain"
)

// GetProcessingStartedLastProcessedBlock retrieves the last processed block for the ProcessingStarted event.
func (s *Storage) GetProcessingStartedLastProcessedBlock() (uint64, error) {
	db, err := s.LoadDatabase()
	if err != nil {
		return 0, err
	}

	return db.ProcessingStarted.LastProcessedBlock, nil
}

// SaveProcessingStartedLastProcessedBlock updates the last processed block for the ProcessingStarted event.
func (s *Storage) SaveProcessingStartedLastProcessedBlock(block uint64) error {
	db, err := s.LoadDatabase()
	if err != nil {
		return err
	}

	db.ProcessingStarted.LastProcessedBlock = block
	return s.SaveDatabase(db)
}

// GetProcessingStartedEvents retrieves all ProcessingStarted events.
func (s *Storage) GetProcessingStartedEvents() ([]domain.BindingsProcessingStarted, error) {
	db, err := s.LoadDatabase()
	if err != nil {
		return nil, err
	}

	return db.ProcessingStarted.Events, nil
}

// SaveProcessingStartedEvent saves a ProcessingStarted event.
func (s *Storage) SaveProcessingStartedEvent(event domain.BindingsProcessingStarted) error {
	db, err := s.LoadDatabase()
	if err != nil {
		return err
	}

	db.ProcessingStarted.Events = append(db.ProcessingStarted.Events, event)
	return s.SaveDatabase(db)
}
