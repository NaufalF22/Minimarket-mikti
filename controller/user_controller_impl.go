package controller

import "minimarket_mikti/service"

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(service service.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		UserService: service,
	}
}
