package repository

import "github.com/yeatmangeorge/album-club-io-backend/domain/src/model"

type UserRepository interface {
	FindAll() ([]model.User, error)
	FindByEmail(email string) (model.User, error)
	Save(model.User) error
}
