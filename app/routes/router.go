package routes

import (
	"fin-dashboard-api/app/controllers"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/status", CheckStatus)
	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)
	router.GET("/validate", controllers.Validate)
	return router
}
