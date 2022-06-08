package handlers

import (
	"net/http"
	"strings"

	"github.com/DiptoChakrabarty/go-mvcs/oauth/access_token"
	"github.com/DiptoChakrabarty/go-mvcs/oauth/access_token_service"
	"github.com/DiptoChakrabarty/go-mvcs/resterrors"
	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetById(ctx *gin.Context)
	Create(ctx *gin.Context)
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
		ctx.JSON(err.Status(), err)
		return
	}
	ctx.JSON(http.StatusOK, accessToken)
}

func (hdr accessTokenHandler) Create(ctx *gin.Context) {
	var at access_token.AccessToken
	if err := ctx.ShouldBindJSON(&at); err != nil {
		restErr := resterrors.BadRequestError("invalid json body")
		ctx.JSON(restErr.Status(), restErr)
		return
	}

	if err := hdr.service.Create(at); err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	ctx.JSON(http.StatusCreated, at)
}
