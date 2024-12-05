package relaysused

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
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
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create GET request for URL %s: %w", url, err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request to Dappmanager at %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response status: %d", resp.StatusCode)
	}

	var responseString string
	if err := json.NewDecoder(resp.Body).Decode(&responseString); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	// Split the comma-separated string into a slice of strings
	relays := []string{}
	if responseString != "" {
		relays = splitAndTrim(responseString, ",")
	}

	return relays, nil
}

// Helper function to split and trim strings
func splitAndTrim(input string, sep string) []string {
	parts := strings.Split(input, sep)
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}
