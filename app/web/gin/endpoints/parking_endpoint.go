package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/pcavezzan/sample-apiserver/app/services"
	"net/http"
)

type ParkingEndpoint struct {
	*services.ParkingService
}

func (p ParkingEndpoint) Users(c *gin.Context) {
	c.JSON(http.StatusOK, p.ParkingService.GetAll())
}