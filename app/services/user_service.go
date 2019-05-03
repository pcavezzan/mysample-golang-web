package services

import (
	"github.com/pcavezzan/sample-apiserver/app/models"
	"github.com/pcavezzan/sample-apiserver/app/repositories"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return UserService{userRepository}
}

func (u UserService) GetAll() []models.User {
	return u.userRepository.GetAll()
}
