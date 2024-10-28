package storage

import (
	"encoding/json"
	"lido-events/internal/application/domain"
	"os"
	"strconv"
)

// GetLidoReport loads the Lido report data from the performance JSON file for the given epoch range
func (fs *Storage) GetLidoReport(start, end string) (domain.Reports, error) {
	performanceData, err := loadPerformance(fs)
	if err != nil {
		return nil, err
	}
	reportData := make(map[uint32]domain.Report)
	for epoch, report := range performanceData {
		
		// convert start and end to uint32
		startInt, err := strconv.ParseUint(start, 10, 32)
		if err != nil {
			return nil, err
		}
		endInt, err := strconv.ParseUint(end, 10, 32)
		if err != nil {
			return nil, err
		}

		// Explicitly convert to uint32 since ParseUint returns uint64
		startUint32 := uint32(startInt)
		endUint32 := uint32(endInt)

		if epoch >= startUint32 && epoch <= endUint32 {
			reportData[epoch] = report
		}
	}
	return reportData, nil
}

// SaveLidoReport saves the Lido report data to the performance JSON file
func (fs *Storage) SaveLidoReport(reports domain.Reports) error {

	performanceData, err := loadPerformance(fs)
	if err != nil {
		return err
	}

	for epoch, report := range reports {
		performanceData[epoch] = report
	}

	return savePerformance(fs, performanceData)
}

// loadPerformance loads the validator performance data from the JSON file
func loadPerformance(fs *Storage) (map[uint32]domain.Report, error) {
	file, err := os.ReadFile(fs.PerformanceFile)
	if err != nil {
		return nil, err
	}
	var data map[uint32]domain.Report
	err = json.Unmarshal(file, &data)
	return data, err
}

// savePerformance saves the validator performance data to the JSON file
func savePerformance(fs *Storage, data map[uint32]domain.Report) error {
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fs.PerformanceFile, file, 0644)
}
