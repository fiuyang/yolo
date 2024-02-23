package exception

type ValidationError struct {
	err         error
	requestType interface{}
}

func NewValidateError(
	err error,
	requestType interface{},
) *ValidationError {
	return &ValidationError{
		err:         err,
		requestType: requestType,
	}
}

func (err *ValidationError) Error() string {
	return err.err.Error()
}
