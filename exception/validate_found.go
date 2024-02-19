package exception

type ValidateFound struct {
	Error []error
}

func NewValidateFound(errCatch []error) ValidateFound {
	return ValidateFound{
		Error: errCatch,
	}
}
