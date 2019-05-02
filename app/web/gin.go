package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pcavezzan/sample-apiserver/app"
	"github.com/pcavezzan/sample-apiserver/app/web/gin/endpoints"
	"github.com/pcavezzan/sample-apiserver/log"
)

func startGin(factory *app.AppFactory, config *app.Config) {
	r := gin.Default()

	ue := endpoints.User{factory.UserService()}

	r.GET("/api/users", ue.Users)

	addr := fmt.Sprintf(":%d", config.Port)
	log.Infof("Listening to %v", addr)

	r.Run(addr)
}
