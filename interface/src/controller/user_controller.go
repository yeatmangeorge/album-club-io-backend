package controller

import "github.com/yeatmangeorge/album-club-io-backend/interface/src/request"

type UserController interface {
	Save(request request.NewUserRequest) (string, error)
}
