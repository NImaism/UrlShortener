package Service

import (
	"Shorterism/Model"
	"Shorterism/Repository"
)

type userService interface {
	RegisterUser(user Model.User) error
	Login(user Model.LoginModel) (bool, error)
}

type UserService struct{}

func NewUserService() userService {
	return UserService{}
}

func (UserService) RegisterUser(User Model.User) error {
	Srv := Repository.NewUserRepository()
	return Srv.RegisterUser(User)
}

func (UserService) Login(User Model.LoginModel) (bool, error) {
	Srv := Repository.NewUserRepository()
	return Srv.Login(User)
}
