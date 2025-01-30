package storage

import (
	"lido-events/internal/application/domain"
)

// GetElRewardsStealingPenaltyReportedLastProcessedBlock retrieves the last processed block for the ELRewardsStealingPenaltyReported event
func (fs *Storage) GetElRewardsStealingPenaltiesReportedLastProcessedBlock() (uint64, error) {
	db, err := fs.LoadDatabase()
	if err != nil {
		return 0, err
	}

	return db.ElRewardsStealingPenaltiesReported.LastProcessedBlock, nil
}

// SaveElRewardsStealingPenaltiesReportedLastProcessedBlock updates the last processed block for the ELRewardsStealingPenaltyReported event
func (fs *Storage) SaveElRewardsStealingPenaltiesReportedLastProcessedBlock(block uint64) error {
	db, err := fs.LoadDatabase()
	if err != nil {
		return err
	}

	db.ElRewardsStealingPenaltiesReported.LastProcessedBlock = block
	return fs.SaveDatabase(db)
}

// GetElRewardsStealingPenaltiesReported retrieves all ELRewardsStealingPenaltyReported events
func (fs *Storage) GetElRewardsStealingPenaltiesReported() ([]domain.CsmoduleELRewardsStealingPenaltyReported, error) {
	db, err := fs.LoadDatabase()
	if err != nil {
		return nil, err
	}

	return db.ElRewardsStealingPenaltiesReported.Penalties, nil
}

// SaveElRewardsStealingPenaltyReported saves an ELRewardsStealingPenaltyReported event
func (fs *Storage) SaveElRewardsStealingPenaltyReported(penalty domain.CsmoduleELRewardsStealingPenaltyReported) error {
	db, err := fs.LoadDatabase()
	if err != nil {
		return err
	}

	db.ElRewardsStealingPenaltiesReported.Penalties = append(db.ElRewardsStealingPenaltiesReported.Penalties, penalty)
	return fs.SaveDatabase(db)
}
