package service_test

import (
	"testing"

	"albumclubio.com/domain/src/model"
	"albumclubio.com/service/src/repository/mocks"
	"albumclubio.com/service/src/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserServiceSave(t *testing.T) {
	userRepo := &mocks.MockUserRepository{}
	userRepo.On("Save", mock.AnythingOfType("model.User")).Return(nil)
	userService := service.NewStandardUserService(userRepo)

	shouldPassErr := userService.Save(model.User{
		Email: "toolong@gmail.com",
	})
	assert.Nil(t, shouldPassErr)

	shouldNotPassErr := userService.Save(model.User{
		Email: "123",
	})
	assert.NotNil(t, shouldNotPassErr)
}
