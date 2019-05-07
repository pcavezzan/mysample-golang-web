package api

import (
	"github.com/pcavezzan/sample-apiserver/app/services"
	"net/http"
)

type ParkingHandler struct {
	ParkingService *services.ParkingService
}

func (ph ParkingHandler) Users(w http.ResponseWriter, r *http.Request) {
	u := ph.ParkingService.GetAll()
	writeJson(w, r, u)
}