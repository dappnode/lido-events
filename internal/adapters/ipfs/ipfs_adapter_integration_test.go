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

	// Add assertions here to validate fields, if you have known expected values
	// Example (replace with actual expected values if known):
	if report.Distributable == 0 {
		t.Error("expected non-zero distributable amount")
	}

	// Add more assertions as needed based on the expected data structure.
}
