package ipfs

import (
	"fmt"
	"io"
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

// FetchAndParseIpfs fetches data from IPFS using a CID and returns the content in CAR format.
func (ia *IPFSAdapter) FetchAndParseIpfs(cid string) ([]byte, error) {
	url := fmt.Sprintf("%s/ipfs/%s?format=car", ia.GatewayURL, cid)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set the Accept header to request CAR format
	req.Header.Set("Accept", "application/vnd.ipld.car")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data from IPFS: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	// Use io.ReadAll instead of ioutil.ReadAll
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, nil
}
