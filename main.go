package main

import (
	"github.com/gin-gonic/gin"

	controllers "app-v1/controllers"
)

func main() {
	app := gin.Default()
	var guitarController controllers.GuitarController
	app.GET("/guitar", guitarController.Get)
	app.GET("/guitar/:id", guitarController.GetById)
	app.POST("/guitar", guitarController.Post)
	app.DELETE("/guitar/:id", guitarController.Delete)
	app.Run()
}
