package handlers

import (
	"net/http"

	"github.com/DiptoChakrabarty/go-mvcs/oauth/access_token_service"
	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetById(ctx *gin.Context)
}

type accessTokenHandler struct {
	service access_token_service.AccessTokenService
}

func NewHandler(svc access_token_service.AccessTokenService) AccessTokenHandler {
	return accessTokenHandler{
		service: svc,
	}
}

func (h accessTokenHandler) GetById(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, "Implement Me")
}
