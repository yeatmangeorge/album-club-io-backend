package controller

import "albumclubio.com/interface/src/request"

type UserController interface {
	Save(request request.NewUserRequest) (string, error)
}
