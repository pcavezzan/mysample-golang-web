package app

import (
	"github.com/pcavezzan/sample-apiserver/app/repositories"
	"github.com/pcavezzan/sample-apiserver/app/services"
)

type App interface {
	Start(config *Config)
}

type AppFactory struct {
	userRepository *repositories.ParkingRepository
	userService *services.ParkingService
}

func NewAppFactory(config *Config) *AppFactory {
	userRepository := repositories.NewParkingRepository()
	userService := services.NewParkingService(userRepository)
	return &AppFactory{
		userRepository: &userRepository,
		userService: &userService,
	}
}

func (f *AppFactory) ParkingService() *services.ParkingService {
	return f.userService
}