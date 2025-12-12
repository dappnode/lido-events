package notifier

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strings"
	"time"

	"lido-events/internal/adapters/notifier/dappmanager"
	"lido-events/internal/application/domain"
)

type Notifier struct {
	ctx          context.Context
	Network      string
	Category     Category
	LidoDnpName  string
	BrainUrl     string
	StakersUiUrl string
	BeaconchaUrl string
	HTTPClient   *http.Client
	DappManager  *dappmanager.DappManager
}

func NewNotifier(ctx context.Context, network, dappmanagerUrl, lidoDnpName, brainUrl, stakersUiUrl, beaconchaUrl string) *Notifier {
	category := Category(strings.ToLower(network))
	if network == "mainnet" {
		category = Ethereum
	}

	dm := dappmanager.NewDappManager(dappmanagerUrl, lidoDnpName)

	return &Notifier{
		ctx:          ctx,
		Network:      network,
		Category:     category,
		LidoDnpName:  lidoDnpName,
		BrainUrl:     brainUrl,
		StakersUiUrl: stakersUiUrl,
		BeaconchaUrl: beaconchaUrl,
		HTTPClient:   &http.Client{Timeout: 3 * time.Second},
		DappManager:  dm,
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
	req, err := http.NewRequestWithContext(n.ctx, "POST", url, bytes.NewBuffer(body))
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
	enabled, err := n.DappManager.IsMissingLogReceiptsEnabled(n.ctx)
	if err != nil || !enabled {
		return err
	}

	payload := NotificationPayload{
		Title:         "⚠️ Lido CSM: Execution client missing Log Receipts",
		Body:          message,
		Category:      &n.Category,
		Priority:      func() *Priority { p := High; return &p }(),
		DnpName:       &n.LidoDnpName,
		CorrelationId: func() *string { s := string(domain.Notifications.MissingLogReceipts); return &s }(),
		CallToAction: &CallToAction{
			Title: "Switch execution",
			URL:   n.StakersUiUrl,
		},
	}
	return n.sendNotification(payload)
}

// SendValidatorExitRequestedNotification sends a notification about a validator exit request.
func (n *Notifier) SendValidatorExitRequestedNotification(message string) error {
	enabled, err := n.DappManager.IsValidatorExitRequestEnabled(n.ctx)
	if err != nil || !enabled {
		return err
	}

	payload := NotificationPayload{
		Title:         "🚪 Lido CSM: Validator Exit Requested",
		Body:          message,
		Category:      &n.Category,
		Priority:      func() *Priority { p := Medium; return &p }(),
		DnpName:       &n.LidoDnpName,
		CorrelationId: func() *string { s := string(domain.Notifications.ValidatorExitRequest); return &s }(),
	}
	return n.sendNotification(payload)
}

// SendValidatorFailedExitNotification sends a notification about a validator failed exit.
func (n *Notifier) SendValidatorFailedExitNotification(message string) error {
	enabled, err := n.DappManager.IsValidatorExitEnabled(n.ctx)
	if err != nil || !enabled {
		return err
	}

	payload := NotificationPayload{
		Title:         "❌ Lido CSM: Validator Exit Failed",
		Body:          message,
		Category:      &n.Category,
		Priority:      func() *Priority { p := Critical; return &p }(),
		DnpName:       &n.LidoDnpName,
		CorrelationId: func() *string { s := string(domain.Notifications.ValidatorExit); return &s }(),
		CallToAction: &CallToAction{
			Title: "Exit validator manually",
			URL:   n.BrainUrl,
		},
	}
	return n.sendNotification(payload)
}

// SendValidatorSucceedExitNotification sends a notification about a validator successful exit.
func (n *Notifier) SendValidatorSucceedExitNotification(message string, validatorIndex *big.Int) error {
	enabled, err := n.DappManager.IsValidatorExitEnabled(n.ctx)
	if err != nil || !enabled {
		return err
	}

	payload := NotificationPayload{
		Title:         "✅  Lido CSM: Validator Exit Succeeded",
		Body:          message,
		Category:      &n.Category,
		Priority:      func() *Priority { p := Medium; return &p }(),
		DnpName:       &n.LidoDnpName,
		CorrelationId: func() *string { s := string(domain.Notifications.ValidatorExit); return &s }(),
		CallToAction: &CallToAction{
			Title: "View validator",
			URL:   fmt.Sprintf("%s/validator/%s", n.BeaconchaUrl, validatorIndex.String()),
		},
	}
	return n.sendNotification(payload)
}

// SendBlackListedNotification sends a notification about blacklisted relays.
func (n *Notifier) SendBlackListedNotification(message string) error {
	enabled, err := n.DappManager.IsRelaysBlacklistEnabled(n.ctx)
	if err != nil || !enabled {
		return err
	}

	payload := NotificationPayload{
		Title:         "🚨 Lido CSM: Blacklisted Relays Detected",
		Body:          message,
		Category:      &n.Category,
		Priority:      func() *Priority { p := High; return &p }(),
		DnpName:       &n.LidoDnpName,
		CorrelationId: func() *string { s := string(domain.Notifications.RelaysBlacklist); return &s }(),
		CallToAction: &CallToAction{
			Title: "Review relays",
			URL:   n.StakersUiUrl,
		},
	}
	return n.sendNotification(payload)
}

// SendMissingRelayNotification sends a notification about missing mandatory relays.
func (n *Notifier) SendMissingRelayNotification(message string) error {
	enabled, err := n.DappManager.IsRelaysMissingEnabled(n.ctx)
	if err != nil || !enabled {
		return err
	}

	payload := NotificationPayload{
		Title:         "⚠️ Lido CSM: No Mandatory Relays in Use",
		Body:          message,
		Category:      &n.Category,
		Priority:      func() *Priority { p := High; return &p }(),
		DnpName:       &n.LidoDnpName,
		CorrelationId: func() *string { s := string(domain.Notifications.RleaysMissing); return &s }(),
		CallToAction: &CallToAction{
			Title: "Review relays",
			URL:   n.StakersUiUrl,
		},
	}
	return n.sendNotification(payload)
}
