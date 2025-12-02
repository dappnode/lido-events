package performance

import (
	"database/sql"
	"fmt"
	"lido-events/internal/application/domain"

	_ "github.com/mattn/go-sqlite3"
)

// Adapter wraps a sqlite DB connection used to store performance reports.
type Adapter struct {
	db *sql.DB
}

// NewAdapter opens (or creates) the sqlite database at dbPath and ensures the
// required schema exists.
func NewAdapter(dbPath string) (*Adapter, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open sqlite database: %w", err)
	}

	adapter := &Adapter{db: db}
	if err := adapter.migrate(); err != nil {
		_ = db.Close()
		return nil, err
	}

	return adapter, nil
}

// migrate creates the performance table if it does not exist.
func (a *Adapter) migrate() error {
	const createTable = `
CREATE TABLE IF NOT EXISTS performance (
	log_cid TEXT NOT NULL,
	frame_start INTEGER NOT NULL,
	frame_end INTEGER NOT NULL,
	no_id TEXT NOT NULL,
	validator_index TEXT NOT NULL,
	performance REAL NOT NULL,
	slashed INTEGER NOT NULL,
	strikes INTEGER NOT NULL,
	threshold REAL NOT NULL
);`

	if _, err := a.db.Exec(createTable); err != nil {
		return fmt.Errorf("failed to create performance table: %w", err)
	}

	return nil
}

// Close closes the underlying database connection.
func (a *Adapter) Close() error {
	if a.db == nil {
		return nil
	}
	return a.db.Close()
}

// SaveReport stores all operator/validator rows for a given logCid and report.
// Existing rows for the same logCid will be removed before inserting.
func (a *Adapter) SaveReport(logCid string, report *domain.Report) error {
	tx, err := a.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	// Remove any existing rows for this logCid to keep data consistent.
	if _, err = tx.Exec(`DELETE FROM performance WHERE log_cid = ?`, logCid); err != nil {
		return fmt.Errorf("failed to delete existing rows for log_cid %s: %w", logCid, err)
	}

	stmt, err := tx.Prepare(`
INSERT INTO performance (
	log_cid, frame_start, frame_end, no_id, validator_index,
	performance, slashed, strikes, threshold
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);
`)
	if err != nil {
		return fmt.Errorf("failed to prepare insert statement: %w", err)
	}
	defer stmt.Close()

	frameStart := 0
	frameEnd := 0
	if len(report.Frame) == 2 {
		frameStart = report.Frame[0]
		frameEnd = report.Frame[1]
	}

	for noID, data := range report.Operators {
		for validatorIndex, v := range data.Validators {
			_, err = stmt.Exec(
				logCid,
				frameStart,
				frameEnd,
				noID,
				validatorIndex,
				v.Performance,
				boolToInt(v.Slashed),
				v.Strikes,
				v.Threshold,
			)
			if err != nil {
				return fmt.Errorf("failed to insert performance row: %w", err)
			}
		}
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

// GetNoPerformance retrieves all performance rows for a given node operator ID (noID)
// and returns them as a slice of domain.Report structs (one per log CID),
// omitting the log CID from the response type.
func (a *Adapter) GetNoPerformance(noID string) ([]*domain.Report, error) {
	rows, err := a.db.Query(`
SELECT log_cid, frame_start, frame_end, validator_index,
	performance, slashed, strikes, threshold
FROM performance
WHERE no_id = ?
ORDER BY frame_start, log_cid;
`, noID)
	if err != nil {
		return nil, fmt.Errorf("failed to query performance for no_id %s: %w", noID, err)
	}
	defer rows.Close()

	reportsByCid := make(map[string]*domain.Report)

	for rows.Next() {
		var (
			logCid         string
			frameStart     int
			frameEnd       int
			validatorIndex string
			performance    float64
			slashedInt     int
			strikes        int
			threshold      float64
		)

		if err := rows.Scan(
			&logCid,
			&frameStart,
			&frameEnd,
			&validatorIndex,
			&performance,
			&slashedInt,
			&strikes,
			&threshold,
		); err != nil {
			return nil, fmt.Errorf("failed to scan performance row: %w", err)
		}

		report, ok := reportsByCid[logCid]
		if !ok {
			report = &domain.Report{
				Frame:     [2]int{frameStart, frameEnd},
				Operators: map[string]domain.Data{},
			}
			reportsByCid[logCid] = report
		}

		data, ok := report.Operators[noID]
		if !ok {
			data = domain.Data{Validators: map[string]domain.Validator{}}
		}

		validator := domain.Validator{
			Performance: performance,
			Slashed:     slashedInt != 0,
			Strikes:     strikes,
			Threshold:   threshold,
		}

		if data.Validators == nil {
			data.Validators = make(map[string]domain.Validator)
		}
		data.Validators[validatorIndex] = validator
		report.Operators[noID] = data
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iteration error while querying performance: %w", err)
	}

	reports := make([]*domain.Report, 0, len(reportsByCid))
	for _, r := range reportsByCid {
		reports = append(reports, r)
	}

	return reports, nil
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
