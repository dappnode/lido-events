package ipfs

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"lido-events/internal/application/domain"
	"net/http"
	"time"
)

type IPFSAdapter struct {
	GatewayURL string
}

func NewIPFSAdapter(gatewayURL string) *IPFSAdapter {
	return &IPFSAdapter{
		GatewayURL: gatewayURL,
	}
}

// FetchAndParseIpfs fetches data from IPFS using a CID and parses the JSON content as a Report.
// If it is called with a timeout, it will set a timeout context for the request. If the request
// times out, it will return a boolean indicating a timeout error.
func (ia *IPFSAdapter) FetchAndParseIpfs(cid string, timeout *time.Duration) (domain.OriginalReport, bool, error) {
	isTimeoutError := false

	// Construct the request URL
	url := fmt.Sprintf("%s/ipfs/%s", ia.GatewayURL, cid)

	var req *http.Request
	var err error

	// Create request with or without timeout context
	if timeout != nil {
		ctx, cancel := context.WithTimeout(context.Background(), *timeout)
		defer cancel()
		req, err = http.NewRequestWithContext(ctx, "GET", url, nil)
	} else {
		req, err = http.NewRequest("GET", url, nil)
	}

	if err != nil {
		return domain.OriginalReport{}, false, fmt.Errorf("failed to create request for URL %s: %w", url, err)
	}

	// Set the Accept header to explicitly request JSON
	req.Header.Set("Accept", "application/json")

	// Use the default HTTP client
	client := &http.Client{}

	// Execute the request
	resp, err := client.Do(req)
	if err != nil {
		if timeout != nil && errors.Is(err, context.DeadlineExceeded) {
			isTimeoutError = true
			return domain.OriginalReport{}, isTimeoutError, fmt.Errorf("request to IPFS timed out after %v: %w", *timeout, err)
		}
		return domain.OriginalReport{}, false, fmt.Errorf("failed to fetch data from IPFS at %s: %w", url, err)
	}
	defer resp.Body.Close()

	// Check for timeout responses from the upstream server
	if resp.StatusCode == http.StatusRequestTimeout || resp.StatusCode == http.StatusGatewayTimeout {
		return domain.OriginalReport{}, true, fmt.Errorf("timeout error fetching from IPFS gateway with status %s", resp.Status)
	}

	// Ensure response status is OK
	if resp.StatusCode != http.StatusOK {
		return domain.OriginalReport{}, false, fmt.Errorf("unexpected response status %s fetching from IPFS gateway", resp.Status)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return domain.OriginalReport{}, false, fmt.Errorf("failed to read response body from IPFS gateway: %w", err)
	}

	// Parse the JSON data into the Report struct
	var report domain.OriginalReport
	if err := json.Unmarshal(body, &report); err != nil {
		return domain.OriginalReport{}, false, fmt.Errorf("failed to unmarshal JSON data from IPFS response: %w", err)
	}

	return report, false, nil
}
