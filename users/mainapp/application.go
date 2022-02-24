package mainapp

import (
	"github.com/DiptoChakrabarty/go-mvcs/users/controllers"
	"github.com/DiptoChakrabarty/go-mvcs/users/models"
	"github.com/DiptoChakrabarty/go-mvcs/users/services"
	"github.com/gin-gonic/gin"
)

var (
	router                        = gin.New()
	UserDataBase models.UserModel = models.NewModelDB()

	UserService services.UserOperationService = services.NewUserService(UserDataBase)

	UserController controllers.UserOperationController = controllers.NewUserController(UserService)
)

func StartUserApplication() {
	router.Use(gin.Recovery())
	router.GET("/health", controllers.Health)
	router.GET("/users/:user_id", UserController.GetUser)
	router.POST("/users", UserController.CreateUser)
	router.DELETE("/users/:user_id", UserController.DeleteUser)
	router.Run(":5000")
}
