package ports

import "math/big"

type NotifierPort interface {
	SendMissingLogReceiptsNotification(message string) error
	SendValidatorExitRequestedNotification(message string) error
	SendValidatorFailedExitNotification(message string) error
	SendValidatorSucceedExitNotification(message string, validatorIndex *big.Int) error
	SendBlackListedNotification(message string) error
	SendMissingRelayNotification(message string) error
}
