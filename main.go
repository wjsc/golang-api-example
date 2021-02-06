package main

import (
	controllers "app-v1/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	var guitarController controllers.GuitarController
	router.GET("/guitar", guitarController.Get)
	router.GET("/guitar/:id", guitarController.GetById)
	router.POST("/guitar", guitarController.Post)
	router.PUT("/guitar/:id", guitarController.Put)
	router.DELETE("/guitar/:id", guitarController.Delete)
	router.Run()
}
