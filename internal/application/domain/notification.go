package domain

type LidoNotificationsEnabled map[LidoNotification]bool

type LidoNotification string

type LidoNotifications struct {
	ValidatorExitRequest LidoNotification
	ValidatorExit        LidoNotification
	RelaysBlacklist      LidoNotification
	RleaysMissing        LidoNotification
	MissingLogReceipts   LidoNotification
	NewPerformanceReport LidoNotification
}

var Notifications LidoNotifications

func InitNotifications(network string) {
	prefix := network + "-"
	Notifications = LidoNotifications{
		ValidatorExitRequest: LidoNotification(prefix + "validator-exit-request"),
		ValidatorExit:        LidoNotification(prefix + "validator-exit"),
		RelaysBlacklist:      LidoNotification(prefix + "relays-blacklist"),
		RleaysMissing:        LidoNotification(prefix + "relays-missing"),
		MissingLogReceipts:   LidoNotification(prefix + "missing-log-receipts"),
		NewPerformanceReport: LidoNotification(prefix + "new-performance-report"),
	}
}
