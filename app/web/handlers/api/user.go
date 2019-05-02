package api

import (
	"github.com/pcavezzan/sample-apiserver/app/services"
	"net/http"
)

type UserHandler struct {
	UserService services.UserService
}

func (uh UserHandler) Users(w http.ResponseWriter, r *http.Request) {
	u := uh.UserService.GetAll()
	writeJson(w, r, u)
}