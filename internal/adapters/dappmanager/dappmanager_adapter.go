package dappmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type DappmanagerAdapter struct {
	dappmanagerUrl string
	mevBootDnpName string
}

// NewDappmanagerAdapter creates a new instance of DappmanagerAdapter with provided dappmanagerUrl and mevBoostDnpName
func NewDappmanagerAdapter(dappmanagerUrl, mevBoostDnpName string) *DappmanagerAdapter {
	return &DappmanagerAdapter{
		dappmanagerUrl,
		mevBoostDnpName,
	}
}

// GetRelaysUsed fetches env RELAYS from MEV BOOST pkg
func (da *DappmanagerAdapter) GetRelaysUsed(ctx context.Context) ([]string, error) {
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
