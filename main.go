package main

import (
	"crudapi/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	route := gin.Default()

	models.ConnectDB()

	route.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "HI WORLD"})
	})

	route.Run()
}