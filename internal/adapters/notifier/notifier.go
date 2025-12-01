package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Notifier struct {
	Network     string
	Category    Category
	LidoDnpName string
	HTTPClient  *http.Client
}

func NewNotifier(network, lidoDnpName string) *Notifier {
	category := Category(strings.ToLower(network))
	if network == "mainnet" {
		category = Ethereum
	}
	return &Notifier{
		Network:     network,
		Category:    category,
		LidoDnpName: lidoDnpName,
		HTTPClient:  &http.Client{Timeout: 3 * time.Second},
	}
}

type CallToAction struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type Category string

const (
	Ethereum Category = "ethereum"
	Hoodi    Category = "hoodi"
	Holesky  Category = "holesky"
	Gnosis   Category = "gnosis"
	Lukso    Category = "lukso"
)

type Priority string

const (
	Low      Priority = "low"
	Medium   Priority = "medium"
	High     Priority = "high"
	Critical Priority = "critical"
)

type Status string

const (
	Triggered Status = "triggered"
	Resolved  Status = "resolved"
)

type NotificationPayload struct {
	Title         string        `json:"title"`
	Body          string        `json:"body"`
	Category      *Category     `json:"category,omitempty"`
	Status        *Status       `json:"status,omitempty"`
	IsBanner      *bool         `json:"isBanner,omitempty"`
	Priority      *Priority     `json:"priority,omitempty"`
	CorrelationId *string       `json:"correlationId,omitempty"`
	DnpName       *string       `json:"dnpName,omitempty"`
	CallToAction  *CallToAction `json:"callToAction,omitempty"`
}

func (n *Notifier) sendNotification(payload NotificationPayload) error {
	url := fmt.Sprintf("%s/api/v1/notifications", "http://notifier.notifications.dappnode:8080")
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := n.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send notification: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("notification failed with status: %s", resp.Status)
	}
	return nil
}

// SendMisingLogReceiptsNotification sends a notification about missing log receipts.
func (n *Notifier) SendMissingLogReceiptsNotification(message string) error {
	payload := NotificationPayload{
		Title:    "⚠️ Missing Log Receipts Detected",
		Body:     message,
		Category: &n.Category,
		Priority: func() *Priority { p := High; return &p }(),
		DnpName:  &n.LidoDnpName,
	}
	return n.sendNotification(payload)
}

// SendValidatorExitRequestedNotification sends a notification about a validator exit request.
func (n *Notifier) SendValidatorExitRequestedNotification(message string) error {
	payload := NotificationPayload{
		Title:    "🚪 Validator Exit Requested",
		Body:     message,
		Category: &n.Category,
		Priority: func() *Priority { p := Medium; return &p }(),
		DnpName:  &n.LidoDnpName,
	}
	return n.sendNotification(payload)
}

// SendValidatorExecutingExitNotification sends a notification about a validator executing exit.
func (n *Notifier) SendValidatorExecutingExitNotification(message string) error {
	payload := NotificationPayload{
		Title:    "🏃 Validator Executing Exit",
		Body:     message,
		Category: &n.Category,
		Priority: func() *Priority { p := Medium; return &p }(),
		DnpName:  &n.LidoDnpName,
	}
	return n.sendNotification(payload)
}

// SendValidatorFailedExitNotification sends a notification about a validator failed exit.
func (n *Notifier) SendValidatorFailedExitNotification(message string) error {
	payload := NotificationPayload{
		Title:    "❌ Validator Exit Failed",
		Body:     message,
		Category: &n.Category,
		Priority: func() *Priority { p := High; return &p }(),
		DnpName:  &n.LidoDnpName,
	}
	return n.sendNotification(payload)
}

// SendValidatorSucceedExitNotification sends a notification about a validator successful exit.
func (n *Notifier) SendValidatorSucceedExitNotification(message string) error {
	payload := NotificationPayload{
		Title:    "✅ Validator Exit Succeeded",
		Body:     message,
		Category: &n.Category,
		Priority: func() *Priority { p := Medium; return &p }(),
		DnpName:  &n.LidoDnpName,
	}
	return n.sendNotification(payload)
}

// SendBlackListedNotification sends a notification about blacklisted relays.
func (n *Notifier) SendBlackListedNotification(message string) error {
	payload := NotificationPayload{
		Title:    "🚨 Blacklisted Relays Detected",
		Body:     message,
		Category: &n.Category,
		Priority: func() *Priority { p := High; return &p }(),
		DnpName:  &n.LidoDnpName,
	}
	return n.sendNotification(payload)
}

// SendMissingRelayNotification sends a notification about missing mandatory relays.
func (n *Notifier) SendMissingRelayNotification(message string) error {
	payload := NotificationPayload{
		Title:    "⚠️ No Mandatory Relays in Use",
		Body:     message,
		Category: &n.Category,
		Priority: func() *Priority { p := Medium; return &p }(),
		DnpName:  &n.LidoDnpName,
	}
	return n.sendNotification(payload)
}
