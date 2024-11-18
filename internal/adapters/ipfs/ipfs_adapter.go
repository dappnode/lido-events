package ipfs

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"lido-events/internal/application/domain"
	"log"
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
		log.Printf("Error creating request: %v", err)
		return domain.OriginalReport{}, err
	}

	// Set the Accept header to explicitly request JSON
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to fetch data from IPFSt: %v", err)
		return domain.OriginalReport{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errorMessage := fmt.Sprintf("unexpected response status fetching IPFS gateway: %s", resp.Status)
		log.Println(errorMessage)
		return domain.OriginalReport{}, errors.New(errorMessage)
	}

	// Read the JSON data directly from the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return domain.OriginalReport{}, err
	}

	// Parse the JSON data into the Report struct
	var report domain.OriginalReport
	err = json.Unmarshal(body, &report)
	if err != nil {
		log.Printf("Error unmarshalling JSON data to Report: %v", err)
		return domain.OriginalReport{}, err
	}

	return report, nil
}
