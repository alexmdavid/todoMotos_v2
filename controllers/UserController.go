package controllers
import (
	"net/http"
	"todoMotos/config"
	"todoMotos/models"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
		var users []models.User

		result := config.DB.Find(&users)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, users)
	}