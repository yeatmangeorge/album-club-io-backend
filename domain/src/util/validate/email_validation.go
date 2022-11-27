package validate

import (
	"net/mail"

	"album-club-io-backend/domain/src/constant"
	domainerror "album-club-io-backend/domain/src/error"
)

func Email(email string) error {
	if len(email) > constant.MAX_EMAIL_LENGTH {
		return domainerror.NewFieldLengthError(
			"Email",
			constant.MIN_EMAIL_LENGTH,
			constant.MAX_EMAIL_LENGTH,
		)
	}
	_, err := mail.ParseAddress(email)
	if err != nil {
		return err
	}

	return nil
}
