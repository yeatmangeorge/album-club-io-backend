package model

import (
	"github.com/yeatmangeorge/album-club-io-backend/domain/src/util/validate"

	"github.com/google/uuid"
)

type User struct {
	id    string
	Email string
}

//GETTERS AND SETTERS

func (user *User) GetID() string {
	return user.id
}

func (user *User) SetEmail(email string) error {
	user.Email = email
	return user.Validate()
}

//FUNCTIONS

func (user *User) Validate() error {
	emailValidationError := validate.Email(user.Email)
	if emailValidationError != nil {
		return emailValidationError
	}

	return nil
}

func NewUser(email string) User {
	return User{
		id:    uuid.NewString(),
		Email: email,
	}
}

// func (u User) GetID() string {
// 	return u.id
// }

// func (u User) GetEmail() string {
// 	return u.email
// }
