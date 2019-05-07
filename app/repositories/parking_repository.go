package repositories

import "github.com/pcavezzan/sample-apiserver/app/models"

type ParkingRepository interface {
	GetAll() []models.Parking
}

func NewParkingRepository() ParkingRepository {
	r := InMemoryParkingRepository{}
	r.Data = []models.Parking{
		{Number: 1, Owner: "Foo Bar", Car: "Dodge Viper"},
		{Number: 2, Owner: "James Bond", Car: "Aston Martin"},
	}
	return r
}

type InMemoryParkingRepository struct {
	Data []models.Parking
}

func (r InMemoryParkingRepository) GetAll() []models.Parking {
	return r.Data
}