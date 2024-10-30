package ports

type ExitValidator interface {
	ExitValidator(pubKey, validatorIndex string) error
}
