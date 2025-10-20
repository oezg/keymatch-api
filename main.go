package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/hello", helloWorld)

	server.Run(":3434")
}

func helloWorld(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
}
