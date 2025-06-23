package controllers
import (
	"net/http"
	"todoMotos/config"
	"todoMotos/models"
	"github.com/gin-gonic/gin"
	"strconv"
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


func GetUserByID(c *gin.Context) {
	idStr := c.Param("id")
	idUint, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var user models.User
	result := config.DB.First(&user, uint(idUint))
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})

}
