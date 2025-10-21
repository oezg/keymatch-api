package main

import (
	"github.com/gin-gonic/gin"
	"github.com/oezg/keymatch-api/db"
	"github.com/oezg/keymatch-api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)
	server.Run(":3434")
}
