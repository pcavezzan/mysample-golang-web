package repositories

import "github.com/pcavezzan/sample-apiserver/app/models"

type UserRepository interface {
	GetAll() []models.User
}

func NewUserRepository() UserRepository {
	r := InMemoryUserRepository{}
	r.Data = []models.User{
		{LastName: "lname", FirstName: "fname"},
		{LastName: "bar", FirstName: "foo"},
	}
	return r
}

type InMemoryUserRepository struct {
	Data []models.User
}

func (r InMemoryUserRepository) GetAll() []models.User {
	return r.Data
}