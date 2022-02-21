package mainapp

import (
	"github.com/DiptoChakrabarty/go-mvcs/users/controllers"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.New()
)

func StartUserApplication() {
	router.Use(gin.Recovery())
	router.GET("/health", controllers.Health)
	router.GET("/users/:user_id", controllers.GetUser)
	router.POST("/users", controllers.CreateUser)
	router.DELETE("/users/:user_id", controllers.DeleteUser)
	router.Run(":5000")
}
