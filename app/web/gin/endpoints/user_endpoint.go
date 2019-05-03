package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/pcavezzan/sample-apiserver/app/services"
	"net/http"
)

type User struct {
	services.UserService
}

func (u User) Users(c *gin.Context) {
	c.JSON(http.StatusOK, u.UserService.GetAll())
}