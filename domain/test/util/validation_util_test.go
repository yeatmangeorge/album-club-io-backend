package util

import (
	"testing"

	"album-club-io-backend/domain/src/util/validate"

	"github.com/stretchr/testify/assert"
)

func TestValidEmail(t *testing.T) {
	email := "test@gmail.com"
	validationErr := validate.Email(email)
	assert.Nil(t, validationErr)
}
