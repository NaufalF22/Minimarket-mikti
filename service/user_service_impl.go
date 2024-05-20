package service

type UserServiceImpl struct {
	repository   repository.UserRepository
	tokenUseCase helper.TokenUseCase
}

func NewUserService(repository repository.UserRepository, token helper.TokenUseCase) *UserServiceImpl {
	return &UserServiceImpl{
		repository:   repository,
		tokenUseCase: token,
	}
}
