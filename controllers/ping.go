package controllers

import (
	"github.com/deeptik2/simple-rest-service/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PingControllers(c *gin.Context) {
	result, err := services.PingServiceVar.PingService()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.String(http.StatusOK, result)
	}
}
