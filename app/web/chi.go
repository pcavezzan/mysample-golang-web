package web

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/pcavezzan/sample-apiserver/app"
	"github.com/pcavezzan/sample-apiserver/app/web/handlers/api"
	"github.com/pcavezzan/sample-apiserver/log"
	"net/http"
)

func startChi(factory *app.AppFactory, config *app.Config) {
	uh := api.UserHandler{UserService: factory.UserService()}

	r := chi.NewRouter()
	r.Get("/api/users", uh.Users)

	addr := fmt.Sprintf(":%d", config.Port)
	log.Infof("Listening to %v", addr)

	http.ListenAndServe(addr, r)
}

