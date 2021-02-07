package controllers

import (
	"net/http"
	"strconv"
	models "app-v1/models"
	storage "app-v1/storage"
	"github.com/gin-gonic/gin"
)

type GuitarController struct{
	Storage *storage.GuitarStorageType
}

func (controller GuitarController) Get(context *gin.Context) {
	var guitars, err = controller.Storage.Get()
	if err!=nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})	
	}
	context.JSON(http.StatusOK, guitars)
}

func (controller GuitarController) GetById(context *gin.Context) {

	var id, err = strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Guitar id must be an integer"})
		return
	}

	guitar, err := controller.Storage.GetById(id)
	if err!= nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Guitar not found"})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"id":    guitar.Id,
		"year":  guitar.Year,
		"model": guitar.Model,
	})
	return
}

func (controller GuitarController) Post(context *gin.Context) {
	var json models.GuitarModel

	if err := context.ShouldBindJSON(&json); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	guitar, err := controller.Storage.Create(json)
	if err!=nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusCreated, guitar)
	return
}

func (controller GuitarController) Put(context *gin.Context) {
	var json models.GuitarModel

	if err := context.ShouldBindJSON(&json); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var id, err = strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Guitar id must be an integer"})
		return
	}

	json.Id = id
	guitar, err := controller.Storage.Update(json)
	if err!=nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusOK, guitar)
	return
}

func (controller GuitarController) Delete(context *gin.Context) {
	
	var id, err = strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Guitar id must be an integer"})
		return
	}

	err = controller.Storage.Delete(id)
	if err!=nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Guitar not found"})
	}
	context.JSON(http.StatusOK, gin.H{})
	return
}
