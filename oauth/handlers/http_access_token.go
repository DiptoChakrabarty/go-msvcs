package handlers

import (
	"github.com/DiptoChakrabarty/go-mvcs/oauth/access_token_service"
	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
}

type accessTokenHandler struct {
	service access_token_service.AccessTokenService
}

func NewHandler(svc access_token_service.AccessTokenService) AccessTokenHandler {
	return accessTokenHandler{
		service: svc,
	}
}
