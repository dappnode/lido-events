package services

import (
	"context"
	"fmt"
	"sync"
	"time"

	"lido-events/internal/application/domain"
	"lido-events/internal/application/ports"
	"lido-events/internal/logger"
)

type RelayCronService struct {
	relaysAllowedPort ports.RelaysAllowedPort
	relaysUsedPort    ports.RelaysUsedPort
	notifierPort      ports.NotifierPort
	servicePrefix     string
}

func NewRelayCronService(
	relaysAllowedPort ports.RelaysAllowedPort,
	relaysUsedPort ports.RelaysUsedPort,
	notifierPort ports.NotifierPort,
) *RelayCronService {
	return &RelayCronService{
		relaysAllowedPort: relaysAllowedPort,
		relaysUsedPort:    relaysUsedPort,
		notifierPort:      notifierPort,
		servicePrefix:     "RelaysChecker",
	}
}

func (rcs *RelayCronService) StartRelayMonitoringCron(ctx context.Context, interval time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(1)

	// Execute immediately on startup
	if err := rcs.monitorRelays(ctx); err != nil {
		logger.ErrorWithPrefix(rcs.servicePrefix, "Error monitoring relays: %v", err)
	}

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := rcs.monitorRelays(ctx); err != nil {
				logger.ErrorWithPrefix(rcs.servicePrefix, "Error monitoring relays: %v", err)
			}
		case <-ctx.Done():
			logger.InfoWithPrefix(rcs.servicePrefix, "Stopping periodic cron for monitoring relays")
			return
		}
	}
}

func (rcs *RelayCronService) monitorRelays(ctx context.Context) error {
	// Fetch allowed relays
	allowedRelays, err := rcs.relaysAllowedPort.GetRelaysAllowList(ctx)
	if err != nil {
		return fmt.Errorf("failed to fetch allowed relays: %w", err)
	}

	// Fetch used relays
	usedRelays, err := rcs.relaysUsedPort.GetRelaysUsed(ctx)
	if err != nil {
		return fmt.Errorf("failed to fetch used relays: %w", err)
	}

	// Analyze relays
	blacklistedRelays, missingMandatoryRelays := rcs.analyzeRelays(allowedRelays, usedRelays)

	// Send notifications if issues are found
	if len(blacklistedRelays) > 0 {
		message := rcs.buildBlacklistNotification(blacklistedRelays)
		if err := rcs.notifierPort.SendNotification(message); err != nil {
			return fmt.Errorf("failed to send blacklist notification: %w", err)
		}
	}

	if len(missingMandatoryRelays) > 0 {
		message := rcs.buildMissingMandatoryNotification(missingMandatoryRelays)
		if err := rcs.notifierPort.SendNotification(message); err != nil {
			return fmt.Errorf("failed to send missing mandatory relays notification: %w", err)
		}
	}

	return nil
}

func (rcs *RelayCronService) analyzeRelays(allowedRelays []domain.RelayAllowed, usedRelays []string) ([]string, []domain.RelayAllowed) {
	allowedMap := make(map[string]domain.RelayAllowed)
	for _, relay := range allowedRelays {
		allowedMap[relay.Uri] = relay
	}

	var blacklistedRelays []string
	var missingMandatoryRelays []domain.RelayAllowed

	// Check for blacklisted relays
	for _, used := range usedRelays {
		if _, exists := allowedMap[used]; !exists {
			blacklistedRelays = append(blacklistedRelays, used)
		}
	}

	// Check for missing mandatory relays
	for _, allowed := range allowedRelays {
		if allowed.IsMandatory {
			found := false
			for _, used := range usedRelays {
				if allowed.Uri == used {
					found = true
					break
				}
			}
			if !found {
				missingMandatoryRelays = append(missingMandatoryRelays, allowed)
			}
		}
	}

	return blacklistedRelays, missingMandatoryRelays
}

func (rcs *RelayCronService) buildBlacklistNotification(blacklistedRelays []string) string {
	message := "🚨 *Relay Blacklist Warning*\n\nThe following relays are being used but are not allowed:\n"
	for _, relay := range blacklistedRelays {
		message += fmt.Sprintf("- `%s`\n", relay)
	}
	message += "\nPlease review and address these issues promptly."
	return message
}

func (rcs *RelayCronService) buildMissingMandatoryNotification(missingMandatoryRelays []domain.RelayAllowed) string {
	message := "⚠️ *Mandatory Relay Missing*\n\nThe following mandatory relays are not being used:\n"
	for _, relay := range missingMandatoryRelays {
		message += fmt.Sprintf("- `%s` (Operator: %s, Description: %s)\n", relay.Uri, relay.Operator, relay.Description)
	}
	message += "\nPlease ensure these relays are used as required."
	return message
}
