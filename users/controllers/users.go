package controllers

import (
	"net/http"
	"strconv"

	"github.com/DiptoChakrabarty/go-mvcs/users/services"
	"github.com/DiptoChakrabarty/go-mvcs/users/types/users"
	"github.com/DiptoChakrabarty/go-mvcs/users/utils/resterrors"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	svc services.UserService
}

func (user *UserController) CreateUser(ctx *gin.Context) {
	var usr users.User
	err := ctx.ShouldBindJSON(&usr)
	if err != nil {
		restErr := resterrors.BadRequestError("Invlaid Values Provided")
		ctx.JSON(restErr.Status, restErr)
	}
	result, saveErr := user.svc.AddUser(usr)
	if err != nil {
		ctx.JSON(saveErr.Status, saveErr)
	}
	ctx.JSON(http.StatusCreated, result)
}

func (user *UserController) GetUser(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 0, 0)
	result, err := user.svc.GetUser(id)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusCreated, result)
}

func (user *UserController) UpdateUser(ctx *gin.Context) {
	var usr users.User
	err := ctx.ShouldBindJSON(&usr)
	if err != nil {
		restErr := resterrors.BadRequestError("Invlaid Values Provided")
		ctx.JSON(restErr.Status, restErr)
	}
	result, saveErr := user.svc.UpdateUser(usr)
	if err != nil {
		ctx.JSON(saveErr.Status, saveErr)
	}
	ctx.JSON(http.StatusCreated, result)
}

func (user *UserController) DeleteUser(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 0, 0)
	result, err := user.svc.DeleteUser(id)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusCreated, result)
}
