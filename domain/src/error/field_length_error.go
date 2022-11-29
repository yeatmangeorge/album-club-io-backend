package domainerror

import (
	"errors"

	"github.com/yeatmangeorge/album-club-io-backend/domain/src/util"
)

// NewFieldLengthError should be thrown when a field is not the desired length.
// This will automatically provide a nicely formatted message.
func NewFieldLengthError(
	fieldName string,
	minLength int,
	maxLength int,
) error {
	return errors.New(
		"field " + fieldName + " is not a valid length. Min length is " +
			util.IntToString(minLength) + ", and max length is " +
			util.IntToString(maxLength),
	)
}
