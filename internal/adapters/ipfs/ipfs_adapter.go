package ipfs

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"lido-events/internal/application/domain"
	"net/http"
	"os"
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
func (ia *IPFSAdapter) FetchAndParseIpfs(cid string, timeout *time.Duration) (domain.OriginalReport, bool, error) {
	isTimeoutError := false
	// Construct the URL without any format query parameter
	url := fmt.Sprintf("%s/ipfs/%s", ia.GatewayURL, cid)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return domain.OriginalReport{}, isTimeoutError, fmt.Errorf("failed to create request for URL %s: %w", url, err)
	}

	// Set the Accept header to explicitly request JSON
	req.Header.Set("Accept", "application/json")

	// Set up the HTTP client with timeout if provided
	client := &http.Client{}
	if timeout != nil {
		client.Timeout = *timeout
	}

	resp, err := client.Do(req)
	if err != nil {
		// Check if error is due to timeout
		if os.IsTimeout(err) || errors.Is(err, context.DeadlineExceeded) {
			isTimeoutError = true
		}
		return domain.OriginalReport{}, isTimeoutError, fmt.Errorf("failed to fetch data from IPFS at %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusRequestTimeout || resp.StatusCode == http.StatusGatewayTimeout {
		isTimeoutError = true
		return domain.OriginalReport{}, isTimeoutError, fmt.Errorf("timeout error fetching from IPFS gateway with status %s", resp.Status)
	}

	if resp.StatusCode != http.StatusOK {
		return domain.OriginalReport{}, isTimeoutError, fmt.Errorf("unexpected response status %s fetching from IPFS gateway", resp.Status)
	}

	// Read the JSON data directly from the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return domain.OriginalReport{}, isTimeoutError, fmt.Errorf("failed to read response body from IPFS gateway: %w", err)
	}

	// Parse the JSON data into the Report struct
	var report domain.OriginalReport
	if err := json.Unmarshal(body, &report); err != nil {
		return domain.OriginalReport{}, isTimeoutError, fmt.Errorf("failed to unmarshal JSON data from IPFS response: %w", err)
	}

	return report, isTimeoutError, nil
}
