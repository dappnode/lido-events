package ports

type NotifierPort interface {
	SendMissingLogReceiptsNotification(message string) error
	SendValidatorExitRequestedNotification(message string) error
	SendValidatorFailedExitNotification(message string) error
	SendValidatorSucceedExitNotification(message string) error
	SendBlackListedNotification(message string) error
	SendMissingRelayNotification(message string) error
}
