package services

import (
	"github.com/pcavezzan/sample-apiserver/app/models"
	"github.com/pcavezzan/sample-apiserver/app/repositories"
)

type ParkingService struct {
	parkingRepository repositories.ParkingRepository
}

func NewParkingService(parkingRepository repositories.ParkingRepository) ParkingService {
	return ParkingService{parkingRepository}
}

func (p ParkingService) GetAll() []models.Parking {
	return p.parkingRepository.GetAll()
}
