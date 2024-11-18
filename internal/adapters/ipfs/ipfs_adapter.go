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
func (ia *IPFSAdapter) FetchAndParseIpfs(cid string) (domain.OriginalReport, error) {
	// Construct the URL without any format query parameter
	url := fmt.Sprintf("%s/ipfs/%s", ia.GatewayURL, cid)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return domain.OriginalReport{}, fmt.Errorf("failed to create request: %w", err)
	}

	// Set the Accept header to explicitly request JSON
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error fetching data from IPFS: %v\n", err)
		return domain.OriginalReport{}, fmt.Errorf("failed to fetch data from IPFS: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return domain.OriginalReport{}, fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	// Read the JSON data directly from the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return domain.OriginalReport{}, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse the JSON data into the Report struct
	var report domain.OriginalReport
	err = json.Unmarshal(body, &report)
	if err != nil {
		fmt.Printf("Error unmarshalling JSON data to Report: %v\n", err)
		return domain.OriginalReport{}, fmt.Errorf("failed to unmarshal JSON data: %w", err)
	}

	return report, nil
}
