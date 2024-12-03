package relaysused

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type RelaysUsedAdapter struct {
	dappmanagerUrl string
	mevBootDnpName string
}

// NewRelaysUsedAdapter creates a new instance of DappmanagerAdapter with provided dappmanagerUrl and mevBoostDnpName
func NewRelaysUsedAdapter(dappmanagerUrl, mevBoostDnpName string) *RelaysUsedAdapter {
	return &RelaysUsedAdapter{
		dappmanagerUrl,
		mevBoostDnpName,
	}
}

// GetRelaysUsed fetches env RELAYS from MEV BOOST pkg
func (da *RelaysUsedAdapter) GetRelaysUsed(ctx context.Context) ([]string, error) {
	url := fmt.Sprintf("%s/env/%s?envName=RELAYS", da.dappmanagerUrl, da.mevBootDnpName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create GET request for URL %s: %w", url, err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request to Dappmanager at %s: %w", url, err)
	}

	defer resp.Body.Close()

	var result []string
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response for GetRelaysUsed: %w", err)
	}

	return result, nil
}
