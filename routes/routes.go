package routes
import (
	"github.com/gin-gonic/gin"
	"todoMotos/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()


	router.GET("/Bikes", controllers.GetBikes)
	router.GET("/Users", controllers.GetUsers)

	return router
}

