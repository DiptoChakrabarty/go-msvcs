package oauth

import (
	"github.com/DiptoChakrabarty/go-mvcs/oauth/access_token_service"
	"github.com/DiptoChakrabarty/go-mvcs/oauth/handlers"
	"github.com/DiptoChakrabarty/go-mvcs/oauth/repository/database"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func main() {
	DBRepository := database.NewDBRepository()
	TokenService := access_token_service.NewAccessTokenService(DBRepository)
	TokenHandler := handlers.NewHandler(TokenService)

	router.GET("/oauth/access_token/:access_token_id", TokenHandler.GetById)
	router.POST("/oauth/access_token/", TokenHandler.Create)
	router.Run(":3000")
}
