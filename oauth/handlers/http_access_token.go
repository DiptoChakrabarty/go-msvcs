package handlers

import (
	"net/http"
	"strings"

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

func (hdr accessTokenHandler) GetById(ctx *gin.Context) {
	accessTokenID := strings.TrimSpace(ctx.Param("access_token_id"))
	accessToken, err := hdr.service.GetById(accessTokenID)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, accessToken)
}
