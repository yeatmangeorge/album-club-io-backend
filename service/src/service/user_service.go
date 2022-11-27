package service

import "albumclubio.com/domain/src/model"

type UserService interface {
	Save(user model.User) error
}
