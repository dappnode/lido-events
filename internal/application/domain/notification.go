package domain

type LidoNotificationsEnabled map[LidoNotification]bool

type LidoNotification string

type LidoNotifications struct {
	Exit        LidoNotification
	Relay       LidoNotification
	Performance LidoNotification
}

var Notifications LidoNotifications

func InitNotifications(network string) {
	Notifications = LidoNotifications{
		Exit:        LidoNotification(network + "exit"),
		Relay:       LidoNotification(network + "relay"),
		Performance: LidoNotification(network + "performance"),
	}
}
