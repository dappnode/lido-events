package ipfs

import (
	"encoding/json"
	"fmt"
	"io"
	"lido-events/internal/application/domain"
	"net/http"
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
func (ia *IPFSAdapter) FetchAndParseIpfs(cid string) (domain.Report, error) {
	// Construct the URL without any format query parameter
	url := fmt.Sprintf("%s/ipfs/%s", ia.GatewayURL, cid)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return domain.Report{}, fmt.Errorf("failed to create request for URL %s: %w", url, err)
	}

	// Set the Accept header to explicitly request JSON
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return domain.Report{}, fmt.Errorf("failed to fetch data from IPFS at %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return domain.Report{}, fmt.Errorf("unexpected response status %s fetching from IPFS gateway", resp.Status)
	}

	// Read the JSON data directly from the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return domain.Report{}, fmt.Errorf("failed to read response body from IPFS gateway: %w", err)
	}

	// First, try to parse the JSON data into a single Report
	var report domain.Report
	if err := json.Unmarshal(body, &report); err == nil {
		return report, nil
	}

	// If that fails, try to parse it as an array of Reports and return the first element
	var reports []domain.Report
	if err := json.Unmarshal(body, &reports); err != nil {
		return domain.Report{}, fmt.Errorf("failed to unmarshal JSON data from IPFS response: %w", err)
	}
	if len(reports) == 0 {
		return domain.Report{}, fmt.Errorf("IPFS response contained an empty array")
	}

	return reports[0], nil
}
