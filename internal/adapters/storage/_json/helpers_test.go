//go:build integration
// +build integration

package _json_test

import (
	"encoding/json"

	"os"
	"testing"

	"lido-events/internal/adapters/storage/_json"

	"github.com/stretchr/testify/require"
)

// CreateTempDatabaseFile creates a temporary database file with optional initial content for tests.
func CreateTempDatabaseFile(t *testing.T, content *_json.Database) *os.File {
	tmpFile, err := os.CreateTemp("", "db.json")
	require.NoError(t, err)

	if content != nil {
		data, err := json.Marshal(content)
		require.NoError(t, err)
		_, err = tmpFile.Write(data)
		require.NoError(t, err)
	}

	err = tmpFile.Close()
	require.NoError(t, err)
	return tmpFile
}
