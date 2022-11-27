package model

import (
	"errors"
)

type User struct {
	id    string
	Email string
}

func (user *User) Id() string {
	return user.id
}

func (user *User) Validate() error {
	if len([]rune(user.Email)) > 3 {
		return errors.New("email is too short")
	}
	return nil
}

// func NewUser(email string) User {
// 	return User{
// 		id:    uuid.NewString(),
// 		email: email,
// 	}
// }

// func (u User) GetID() string {
// 	return u.id
// }

// func (u User) GetEmail() string {
// 	return u.email
// }
