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
	adapter := ipfs.NewIPFS(ipfsURL)

	cids := []string{
		"QmTzK9hcbX8VHeUZWX1iQpJ9oKXRR15HsaLoS4oXUZ7MRe",
		"QmPQxGAFtFz2whzcPK1a2vm3ewPt2iVk9Q58etFVqwyiWM",
		"QmcitxE2oUD2yrEwXnRsSdixWUbfmokAjHMh59j2RKSZM7",
	}

	for _, cid := range cids {
		report, err := adapter.FetchAndParseIpfs(cid)
		if err != nil {
			t.Fatalf("failed to fetch and parse IPFS data for cid %s: %v", cid, err)
		}

		// Print the full report for inspection
		t.Logf("CID %s report: %+v", cid, report)
	}
}
