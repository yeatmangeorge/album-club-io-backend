package service

import (
	"github.com/yeatmangeorge/album-club-io-backend/domain/src/model"
	"github.com/yeatmangeorge/album-club-io-backend/service/src/repository"
)

type StandardUserService struct {
	userRepo repository.UserRepository
}

func NewStandardUserService(
	userRepo repository.UserRepository,
) *StandardUserService {
	return &StandardUserService{
		userRepo: userRepo,
	}
}

func (userService StandardUserService) Save(user model.User) error {
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
