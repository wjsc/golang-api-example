package main

import (
	controllers "app-v1/controllers"
	storage "app-v1/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	var guitarStorage storage.GuitarStorageType

	guitarController := controllers.GuitarController{
		&guitarStorage,
	}

	router.GET("/guitar", guitarController.Get)
	router.GET("/guitar/:id", guitarController.GetById)
	router.POST("/guitar", guitarController.Post)
	router.PUT("/guitar/:id", guitarController.Put)
	router.DELETE("/guitar/:id", guitarController.Delete)
	router.Run()
}
