package main

import (
	"app-v1/controllers"
	"app-v1/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	var guitarMemoryStorage storage.GuitarMemoryStorage

	var guitarController = controllers.GuitarController{
		&guitarMemoryStorage,
	}

	router.GET("/guitar", guitarController.Get)
	router.GET("/guitar/:id", guitarController.GetById)
	router.POST("/guitar", guitarController.Post)
	router.PUT("/guitar/:id", guitarController.Put)
	router.DELETE("/guitar/:id", guitarController.Delete)
	router.Run()
}
