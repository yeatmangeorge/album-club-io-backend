package error

import (
	"albumclubio.com/domain/src/error"
	"net/http"
)

func NewUnprocessibleEntityError(message string) error.MessageErr {
	return &error.StandardMessageErr{
		ErrMessage: message,
		ErrStatus:  http.StatusUnprocessableEntity,
		ErrError:   "invalid_request",
	}
}
