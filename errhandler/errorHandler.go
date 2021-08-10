package errhandler

type CommonError struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

type BadRequestError CommonError

func (e *BadRequestError) Error() string {
	return e.Message
}

func NewBadRequestError(message string) error {
	return &BadRequestError{
		Message: message,
		Code:    "invalid.request",
	}
}
