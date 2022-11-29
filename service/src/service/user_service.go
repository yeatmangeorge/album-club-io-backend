package service

import "github.com/yeatmangeorge/album-club-io-backend/domain/src/model"

type UserService interface {
	Save(user model.User) error
}
