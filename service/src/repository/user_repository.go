package repository

import "albumclubio.com/domain/src/model"

type UserRepository interface {
	FindAll() ([]model.User, error)
	FindByEmail(email string) (model.User, error)
	Save(model.User) error
}
