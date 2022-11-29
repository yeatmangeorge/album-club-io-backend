package controller

import (
	"encoding/json"
	"log"

	"github.com/yeatmangeorge/album-club-io-backend/interface/src/request"
	"github.com/yeatmangeorge/album-club-io-backend/service/src/repository"
	"github.com/yeatmangeorge/album-club-io-backend/service/src/service"
)

type StandardUserController struct {
	userService service.UserService
	//todo logger dependancy injection use zap
}

func NewStandardUserController(
	userRepo repository.UserRepository,
) *StandardUserController {
	userService := service.NewStandardUserService(userRepo)

	return &StandardUserController{
		userService: userService,
	}
}

//todo proper return type and should probably pass in something better
func (userController *StandardUserController) Save(
	jsonRequest []byte,
) (string, error) {
	log.Println("Unmarshelling JSON...")
	var request request.NewUserRequest
	unmarshellingErr := json.Unmarshal(jsonRequest, &request)
	if unmarshellingErr != nil {
		return "", unmarshellingErr
	}

	log.Println("Creating user...")
	user, userCreationError := request.ToUser()
	if userCreationError != nil {
		return "", userCreationError
	}

	log.Println("Saving user with email: " + user.Email + " to repository...")
	saveToRepositoryError := userController.userService.Save(user)
	if saveToRepositoryError != nil {
		return "", saveToRepositoryError
	}

	return "Success!!!!!", nil
}
