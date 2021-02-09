package controllers
import (
	"github.com/gin-gonic/gin"
)

type IController interface {
	Get(context *gin.Context)
	GetById(context *gin.Context)
	Post(context *gin.Context)
	Put(context *gin.Context)
	Delete(context *gin.Context)
}