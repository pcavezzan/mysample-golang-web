package app

import (
	"github.com/pcavezzan/sample-apiserver/app/repositories"
	"github.com/pcavezzan/sample-apiserver/app/services"
)

type App interface {
	Start(config *Config)
}

type AppFactory struct {
	userRepository repositories.UserRepository
	userService services.UserService
}

func NewAppFactory(config *Config) *AppFactory {
	userRepository := repositories.NewUserRepository()
	userService := services.NewUserService(userRepository)
	return &AppFactory{
		userRepository: userRepository,
		userService: userService,
	}
}

func (f *AppFactory) UserService() services.UserService {
	return f.userService
}