package main

import (
	"github.com/deeptik2/simple-rest-service/controllers"
)
import "github.com/gin-gonic/gin"

var router = gin.Default()

func main() {
	gin.SetMode(gin.ReleaseMode)

	router.GET("/ping", controllers.PingControllers)
	router.GET("/albums", controllers.GetAlbums)
	router.POST("/albums", controllers.PostAlbums)
	router.GET("/albums/:id", controllers.GetAlbumById)

	router.Run("0.0.0.0:8010")
}
