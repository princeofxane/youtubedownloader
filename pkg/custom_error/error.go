package custom_error

type CustomError struct {
	Message    string
	StatusCode int
}

func (e *CustomError) Error() string {
	return e.Message
}

func (e *CustomError) GetCode() string {
	return e.Message
}

func CreateErr(msg string, code int) *CustomError {
	return &CustomError{
		Message:    msg,
		StatusCode: code,
	}
}
