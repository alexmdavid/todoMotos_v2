package main

import (
	"todoMotos/config"
	"todoMotos/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	
	config.ConnectDatabase()

	router := gin.Default()

	routes.SetupRouter()

	router.Run(":8080")
}
