//go:build integration
// +build integration

package ipfs_test

import (
	"lido-events/internal/adapters/ipfs"
	"os"
	"testing"
)

func TestFetchAndParseIpfs(t *testing.T) {
	// Retrieve the IPFS gateway URL from the environment variable
	ipfsURL := os.Getenv("IPFS_URL")
	if ipfsURL == "" {
		t.Fatal("IPFS_URL environment variable not set")
	}

	// Initialize the IPFS adapter with the gateway URL
	adapter := ipfs.NewIPFSAdapter(ipfsURL)

	// The CID to fetch from IPFS
	cid := "QmTN9oYsjcJjGMpRrT2PZD4iY6aJpy5aCSBvGKU2a9EMQF"

	// Fetch and parse the report from IPFS
	report, err := adapter.FetchAndParseIpfs(cid)
	if err != nil {
		t.Fatalf("failed to fetch and parse IPFS data: %v", err)
	}

	// Print the fetched report for manual inspection (optional)
	t.Logf("Fetched report: %+v", report)

	// Assertions to validate fields based on expected structure

	// Check if `Frame` has been populated
	if len(report.Frame) != 2 || report.Frame[0] == 0 || report.Frame[1] == 0 {
		t.Error("expected non-zero values in the frame array")
	}

	// Check if `Operators` map has been populated
	if len(report.Operators) == 0 {
		t.Error("expected non-empty operators map")
	}

	// Check one operator (assuming we know at least one operator key, e.g., "0")
	if operator, exists := report.Operators["0"]; exists {
		// Validate the Distributed field
		if operator.Distributed == 0 {
			t.Error("expected non-zero distributed amount for operator '0'")
		}

		// Check if the Validators map has been populated for this operator
		if len(operator.Validators) == 0 {
			t.Error("expected non-empty validators map for operator '0'")
		}

		// Check one validator within operator "0" (assuming a known validator ID, e.g., "1735661")
		if validator, exists := operator.Validators["1735661"]; exists {
			// Validate the Assigned and Included fields in `Perf`
			if validator.Perf.Assigned == 0 {
				t.Error("expected non-zero assigned value in validator '1735661' performance")
			}
			if validator.Perf.Included == 0 {
				t.Error("expected non-zero included value in validator '1735661' performance")
			}

			// Validate the `Slashed` field
			if validator.Slashed {
				t.Error("expected slashed to be false for validator '1735661'")
			}
		} else {
			t.Error("expected validator '1735661' to be present in operator '0'")
		}
	} else {
		t.Error("expected operator '0' to be present in operators map")
	}

	// Validate the `Threshold` field
	if report.Threshold == 0 {
		t.Error("expected non-zero threshold value")
	}
}
