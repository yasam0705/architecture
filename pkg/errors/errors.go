package errors

import "github/architecture/internal/entity"

type Error struct {
	message string
	code    int
	details []string
}

func (e *Error) Error() string {
	return e.message
}

func (e *Error) GetCode() int {
	return e.code
}

func (e *Error) GetDetails() []string {
	return e.details
}

func ErrorStatus(code int, details ...string) *Error {
	return &Error{
		message: entity.Codes(code),
		code:    code,
		details: details,
	}
}
