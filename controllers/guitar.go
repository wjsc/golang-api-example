package controllers

import (
	"net/http"
	"strconv"

	models "app-v1/models"

	"github.com/gin-gonic/gin"
)

var myFirstGuitar = models.Guitar{Id: 1, Year: 1967, Model: "Telecaster"}
var mySecondGuitar = models.Guitar{2, 1985, "Stratocaster"}
var myGuitars = []models.Guitar{myFirstGuitar, mySecondGuitar}

type GuitarController struct{}

func (guitar GuitarController) Get(context *gin.Context) {
	context.JSON(http.StatusOK, myGuitars)
}

func (guitar GuitarController) GetById(context *gin.Context) {

	for _, myGuitar := range myGuitars {
		var id, err = strconv.Atoi(context.Param("id"))
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Guitar id must be an integer"})
			return
		}

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

	myGuitars = append(myGuitars, json)
	context.JSON(http.StatusCreated, gin.H{})
	return
}

func (guitar GuitarController) Delete(context *gin.Context) {
	for index, myGuitar := range myGuitars {
		var id, err = strconv.Atoi(context.Param("id"))
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Guitar id must be an integer"})
			return
		}

		if myGuitar.Id == id {
			myGuitars = append(myGuitars[:index], myGuitars[index+1:]...)
			context.JSON(http.StatusOK, gin.H{})
			return
		}
	}
	context.JSON(http.StatusNotFound, gin.H{"error": "Guitar not found"})
	return
}
