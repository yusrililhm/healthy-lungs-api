package exception

import "net/http"

type Error struct {
	ErrStatus  int    `json:"status"`
	ErrMessage string `json:"message"`
	ErrError   string `json:"error"`
}

// Error implements Exception.
func (e *Error) Error() string {
	return e.ErrError
}

// Message implements Exception.
func (e *Error) Message() string {
	return e.ErrMessage
}

// Status implements Exception.
func (e *Error) Status() int {
	return e.ErrStatus
}

type Exception interface {
	Status() int
	Message() string
	Error() string
}

func NewBadRequestError(message string) Exception {
	return &Error{
		ErrStatus:  http.StatusBadRequest,
		ErrMessage: message,
		ErrError:   "BAD_REQUEST",
	}
}

func NewInternalServerError(message string) Exception {
	return &Error{
		ErrStatus:  http.StatusInternalServerError,
		ErrMessage: message,
		ErrError:   "INTERNAL_SERVER_ERROR",
	}
}

func NewNotFoundError(message string) Exception {
	return &Error{
		ErrStatus:  http.StatusNotFound,
		ErrMessage: message,
		ErrError:   "NOT_FOUND",
	}
}

func NewUnprocessableEntityError(message string) Exception {
	return &Error{
		ErrStatus:  http.StatusUnprocessableEntity,
		ErrMessage: message,
		ErrError:   "UNPROCESSABLE_ENTITY",
	}
}

func NewConflictError(message string) Exception {
	return &Error{
		ErrStatus:  http.StatusConflict,
		ErrMessage: message,
		ErrError:   "CONFLICT_ERROR",
	}
}

func NewUnauthenticatedError(message string) Exception {
	return &Error{
		ErrStatus:  http.StatusUnauthorized,
		ErrMessage: message,
		ErrError:   "UNAUTHENTICATED_ERROR",
	}
}
