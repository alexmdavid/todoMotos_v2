package main

import (
	"todoMotos/config"
	"todoMotos/routes"

	
)

func main() {
	
	config.ConnectDatabase()

	router := routes.SetupRouter()

	

	router.Run(":8080")
}
