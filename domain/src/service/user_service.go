package service

import (
	"albumclubio.com/domain/src/model"
	"albumclubio.com/domain/src/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (userService UserService) Save(user model.User) error {
	validationError := user.Validate()
	if validationError != nil {
		return validationError
	}

	return userService.userRepo.Save(user)
}

// func (s UserService) Duplicated(email string) error {
// 	user, err := s.repo.FindByEmail(email)
// 	if user != nil {
// 		return fmt.Errorf("%s already exists", email)
// 	}
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
