package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"todoMotos/config"
	"todoMotos/models"
)

func main() {
	
	config.ConnectDatabase()

	router := gin.Default()

	router.GET("/Bikes", func(c *gin.Context) {
		var bikes []models.Bike

		result := config.DB.Find(&bikes)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, bikes)
	})

	router.Run(":8080")
}
