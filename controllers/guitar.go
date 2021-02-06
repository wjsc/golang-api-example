package controllers

import (
	"net/http"
	"strconv"

	models "app-v1/models"
	storage "app-v1/storage"
	"github.com/gin-gonic/gin"
)

type GuitarController struct{}

func (guitar GuitarController) Get(context *gin.Context) {
	context.JSON(http.StatusOK, storage.GuitarStorage)
}

func (guitar GuitarController) GetById(context *gin.Context) {

	var id, err = strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Guitar id must be an integer"})
		return
	}

	for _, myGuitar := range storage.GuitarStorage {
		if myGuitar.Id == id {
			context.JSON(http.StatusOK, gin.H{
				"id":    myGuitar.Id,
				"year":  myGuitar.Year,
				"model": myGuitar.Model,
			})
			return
		}
	}

	context.JSON(http.StatusNotFound, gin.H{"error": "Guitar not found"})
	return
}

func (guitar GuitarController) Post(context *gin.Context) {
	var json models.Guitar

	if err := context.ShouldBindJSON(&json); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	storage.GuitarStorage = append(storage.GuitarStorage, json)
	context.JSON(http.StatusCreated, gin.H{})
	return
}

func (guitar GuitarController) Put(context *gin.Context) {

	var json models.Guitar

	if err := context.ShouldBindJSON(&json); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var id, err = strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Guitar id must be an integer"})
		return
	}

	for index, myGuitar := range storage.GuitarStorage {
		if myGuitar.Id == id {
			json.Id = id
			storage.GuitarStorage[index] = json
			context.JSON(http.StatusOK, gin.H{})
			return
		}
	}

	context.JSON(http.StatusNotFound, gin.H{"error": "Guitar not found"})
	return
}


func (guitar GuitarController) Delete(context *gin.Context) {
	for index, myGuitar := range storage.GuitarStorage {
		var id, err = strconv.Atoi(context.Param("id"))
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Guitar id must be an integer"})
			return
		}

		if myGuitar.Id == id {
			storage.GuitarStorage = append(storage.GuitarStorage[:index], storage.GuitarStorage[index+1:]...)
			context.JSON(http.StatusOK, gin.H{})
			return
		}
	}
	context.JSON(http.StatusNotFound, gin.H{"error": "Guitar not found"})
	return
}
