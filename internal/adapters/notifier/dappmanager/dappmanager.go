package dappmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"lido-events/internal/application/domain"
)

// DappManager is the adapter to interact with the DappManager API
type DappManager struct {
	baseURL     string
	lidoDnpName string
	client      *http.Client
}

// NewDappManager creates a new DappManager instance
func NewDappManager(baseURL string, dnpName string) *DappManager {
	return &DappManager{
		baseURL:     baseURL,
		lidoDnpName: dnpName,
		client:      &http.Client{},
	}
}

// Manifest represents the manifest of a package
type manifest struct {
	Notifications struct {
		CustomEndpoints []CustomEndpoint `json:"customEndpoints"`
	} `json:"notifications"`
}

type CustomEndpoint struct {
	Name          string `json:"name"`
	Enabled       bool   `json:"enabled"`
	Description   string `json:"description"`
	IsBanner      bool   `json:"isBanner"`
	CorrelationId string `json:"correlationId"`
	Metric        *struct {
		Treshold float64 `json:"treshold"`
		Min      float64 `json:"min"`
		Max      float64 `json:"max"`
		Unit     string  `json:"unit"`
	} `json:"metric,omitempty"`
}

// getNotificationsEnabled retrieves the notifications from the DappManager API
func (d *DappManager) getNotificationsEnabled(ctx context.Context) (domain.LidoNotificationsEnabled, error) {
	customEndpoints, err := d.getLidoManifestNotifications(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get notifications from signer manifest: %w", err)
	}

	notifications := make(domain.LidoNotificationsEnabled)
	for _, endpoint := range customEndpoints {
		notifications[domain.LidoNotification(endpoint.CorrelationId)] = endpoint.Enabled
	}

	return notifications, nil
}

// isNotificationEnabled is a generic helper that checks whether a given
// notification (identified by its LidoNotification key) is enabled according
// to the signer manifest configuration retrieved from DappManager.
func (d *DappManager) isNotificationEnabled(ctx context.Context, notification domain.LidoNotification) (bool, error) {
	notifications, err := d.getNotificationsEnabled(ctx)
	if err != nil {
		return false, err
	}

	enabled, ok := notifications[notification]
	if !ok {
		return false, nil
	}

	return enabled, nil
}

// IsValidatorExitRequestEnabled reports whether the validator exit request
// notification (correlationId like "<network>-validator-exit-request") is
// enabled in the signer manifest.
func (d *DappManager) IsValidatorExitRequestEnabled(ctx context.Context) (bool, error) {
	return d.isNotificationEnabled(ctx, domain.Notifications.ValidatorExitRequest)
}

// IsValidatorExitEnabled reports whether the validator exit notification
// (correlationId like "<network>-validator-exit") is enabled.
func (d *DappManager) IsValidatorExitEnabled(ctx context.Context) (bool, error) {
	return d.isNotificationEnabled(ctx, domain.Notifications.ValidatorExit)
}

// IsRelaysBlacklistEnabled reports whether the relays blacklist notification
// (correlationId like "<network>-relays-blacklist") is enabled.
func (d *DappManager) IsRelaysBlacklistEnabled(ctx context.Context) (bool, error) {
	return d.isNotificationEnabled(ctx, domain.Notifications.RelaysBlacklist)
}

// IsRelaysMissingEnabled reports whether the relays missing notification
// (correlationId like "<network>-relays-missing") is enabled.
func (d *DappManager) IsRelaysMissingEnabled(ctx context.Context) (bool, error) {
	return d.isNotificationEnabled(ctx, domain.Notifications.RleaysMissing)
}

// IsMissingLogReceiptsEnabled reports whether the missing log receipts
// notification (correlationId like "<network>-missing-log-receipts") is
// enabled.
func (d *DappManager) IsMissingLogReceiptsEnabled(ctx context.Context) (bool, error) {
	return d.isNotificationEnabled(ctx, domain.Notifications.MissingLogReceipts)
}

// getLidoManifestNotifications gets the notifications from the Signer package manifest
func (d *DappManager) getLidoManifestNotifications(ctx context.Context) ([]CustomEndpoint, error) {
	url := d.baseURL + "/package-manifest/" + d.lidoDnpName

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request for package %s: %w", d.lidoDnpName, err)
	}

	resp, err := d.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch manifest for package %s: %w", d.lidoDnpName, err)
	}
	defer resp.Body.Close()

	// This covers all 2xx status codes. If its not 2xx, we dont bother parsing the manifest and return an error
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status code %d for package %s", resp.StatusCode, d.lidoDnpName)
	}

	var manifest manifest
	if err := json.NewDecoder(resp.Body).Decode(&manifest); err != nil {
		return nil, fmt.Errorf("failed to decode manifest for package %s: %w", d.lidoDnpName, err)
	}

	return manifest.Notifications.CustomEndpoints, nil
}
