package controller_test

import (
	"testing"

	"albumclubio.com/interface/src/controller"
	"albumclubio.com/service/src/repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserControllerSave(t *testing.T) {
	userRepo := &mocks.MockUserRepository{}
	userRepo.On("Save", mock.AnythingOfType("model.User")).Return(nil)
	userController := controller.NewStandardUserController(userRepo)

	shouldNotPassRequest1 := []byte("sdflkjasfdlkjaskjbeb")
	_, err1 := userController.Save(shouldNotPassRequest1)
	assert.NotNil(t, err1)

	shouldNotPassRequest2 := []byte(`{"email": ""}`)
	_, err2 := userController.Save(shouldNotPassRequest2)
	assert.NotNil(t, err2)

	shouldPassRequest := []byte(`{"email": "test@gmail.com"}`)
	_, err3 := userController.Save(shouldPassRequest)
	assert.Nil(t, err3)
}
