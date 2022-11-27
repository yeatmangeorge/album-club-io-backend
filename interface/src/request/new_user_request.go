package request

import "album-club-io-backend/domain/src/model"

type NewUserRequest struct {
	Email string `json:"email"`
}

func (newUserRequest *NewUserRequest) ToUser() (model.User, error) {
	user := model.NewUser(
		newUserRequest.Email,
	)
	err := user.Validate()

	return user, err
}
