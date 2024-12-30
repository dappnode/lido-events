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
	stakersUiUrl      string
	relaysAllowedPort ports.RelaysAllowedPort
	relaysUsedPort    ports.RelaysUsedPort
	notifierPort      ports.NotifierPort
	servicePrefix     string
}

func NewRelayCronService(
	stakersUiUrl string,
	relaysAllowedPort ports.RelaysAllowedPort,
	relaysUsedPort ports.RelaysUsedPort,
	notifierPort ports.NotifierPort,
) *RelayCronService {
	return &RelayCronService{
		stakersUiUrl:      stakersUiUrl,
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
	blacklistedRelays, mandatoryRelaysUsed := rcs.analyzeRelays(allowedRelays, usedRelays)

	// Send notifications if issues are found
	if len(blacklistedRelays) > 0 {
		message := rcs.buildBlacklistNotification(blacklistedRelays)
		if err := rcs.notifierPort.SendNotification(message); err != nil {
			return fmt.Errorf("failed to send blacklist notification: %w", err)
		}
	}

	// Change: Notify if no mandatory relays are being used
	if mandatoryRelaysUsed == 0 {
		message := rcs.buildMissingMandatoryNotification(allowedRelays)
		if err := rcs.notifierPort.SendNotification(message); err != nil {
			return fmt.Errorf("failed to send missing mandatory relays notification: %w", err)
		}
	}

	return nil
}

func (rcs *RelayCronService) analyzeRelays(allowedRelays []domain.RelayAllowed, usedRelays []string) ([]string, int) {
	allowedMap := make(map[string]domain.RelayAllowed)
	for _, relay := range allowedRelays {
		allowedMap[relay.Uri] = relay
	}

	var blacklistedRelays []string
	mandatoryRelaysUsed := 0

	// Check for blacklisted relays
	for _, used := range usedRelays {
		if _, exists := allowedMap[used]; !exists {
			blacklistedRelays = append(blacklistedRelays, used)
		}
	}

	// Check for mandatory relays in use
	for _, allowed := range allowedRelays {
		if allowed.IsMandatory {
			for _, used := range usedRelays {
				if allowed.Uri == used {
					mandatoryRelaysUsed++
					break
				}
			}
		}
	}

	return blacklistedRelays, mandatoryRelaysUsed
}

func (rcs *RelayCronService) buildMissingMandatoryNotification(allowedRelays []domain.RelayAllowed) string {
	message := "- ⚠️ Mandatory Relay Alert\n\nNo mandatory relays are currently in use. At least one mandatory relay must be used. Edit your relays in the stakers UI\n"
	message += "\nMandatory relays:\n"
	for _, relay := range allowedRelays {
		if relay.IsMandatory {
			message += fmt.Sprintf("- `%s` (Operator: %s, Description: %s)\n", relay.Uri, relay.Operator, relay.Description)
		}
	}
	message += "\nPlease ensure at least one mandatory relay is in use. Allowed are:"
	for _, relay := range allowedRelays {
		message += fmt.Sprintf("\n- `%s`", relay.Uri)
	}
	message += "\n\nEdit your relays in the Stakers UI to use at least one mandatory relay." + rcs.stakersUiUrl
	return message
}

func (rcs *RelayCronService) buildBlacklistNotification(blacklistedRelays []string) string {
	message := "- 🚨 Relay Blacklist Warning\n\nThe following relays are being used but are not allowed:\n"
	for _, relay := range blacklistedRelays {
		message += fmt.Sprintf("- `%s`\n", relay)
	}
	message += "\nPlease edit your relays in the Stakers UI to use only allowed relays." + rcs.stakersUiUrl
	return message
}
