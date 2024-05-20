package controller

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(service service.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		UserService: service,
	}
}
