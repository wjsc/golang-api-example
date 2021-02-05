package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// import "github.com/gin-gonic/gin"

// const url = "https://swapi.dev/api/people/1/"

// type Response struct {
// 	Name string `json:"name"`
// }
// resp, err := http.Get(url)
// if err != nil {
// 	fmt.Println("Error making HTTP GET Request")
// }

// defer resp.Body.Close()
// body, err := ioutil.ReadAll(resp.Body)
// sb := string(body)
// data := Response{}
// json.Unmarshal([]byte(sb), &data)

// fmt.Printf("Name: %s", data.Name)

var myFirstGuitar = Guitar{Id: 1, Year: 1967, Model: "Telecaster"}
var mySecondGuitar = Guitar{2, 1985, "Stratocaster"}
var myGuitars = []Guitar{myFirstGuitar, mySecondGuitar}

func getController(context *gin.Context) {

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
		}
	}

	context.JSON(http.StatusNotFound, gin.H{"error": "Guitar not found"})
	return
}

func postController(context *gin.Context) {
	var json Guitar

	if err := context.ShouldBindJSON(&json); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	myGuitars = append(myGuitars, json)
	fmt.Println(json)
	fmt.Println(myGuitars)
	context.JSON(http.StatusCreated, gin.H{})
	return
}

func main() {
	app := gin.Default()
	app.GET("/guitar/:id", getController)
	app.POST("/guitar", postController)
	app.Run()
}
