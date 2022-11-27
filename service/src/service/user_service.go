package service

import "album-club-io-backend/domain/src/model"

type UserService interface {
	Save(user model.User) error
}
