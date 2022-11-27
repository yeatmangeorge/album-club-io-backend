package service_test

import (
	"testing"

	"albumclubio.com/domain/src/model"
	"albumclubio.com/domain/src/repository/mocks"
	"albumclubio.com/domain/src/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserServiceSave(t *testing.T) {
	userRepo := &mocks.UserRepository{}
	userRepo.On("Save", mock.AnythingOfType("model.User")).Return(nil)
	userService := service.NewUserService(userRepo)

	shouldPassErr := userService.Save(model.User{
		Email: "toolong@gmail.com",
	})
	assert.Nil(t, shouldPassErr)

	shouldNotPassErr := userService.Save(model.User{
		Email: "123",
	})
	assert.NotNil(t, shouldNotPassErr)
}
