package controllers

import (
	"net/http"
	"todoMotos/config"
	"todoMotos/models"
	"github.com/gin-gonic/gin"
)

func GetBikes(c *gin.Context) {
	var bikes []models.Bike

	result := config.DB.Find(&bikes)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, bikes)
}
