package main

import (
	controllers "app-v1/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	var guitarController controllers.GuitarController
	app.GET("/guitar", guitarController.Get)
	app.GET("/guitar/:id", guitarController.GetById)
	app.POST("/guitar", guitarController.Post)
	app.PUT("/guitar/:id", guitarController.Put)
	app.DELETE("/guitar/:id", guitarController.Delete)
	app.Run()
}
