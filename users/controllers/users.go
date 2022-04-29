package controllers

import (
	"net/http"
	"strconv"

	"github.com/DiptoChakrabarty/go-mvcs/users/models"
	"github.com/DiptoChakrabarty/go-mvcs/users/services"
	"github.com/DiptoChakrabarty/go-mvcs/users/utils/resterrors"
	"github.com/gin-gonic/gin"
)

type UserOperationController interface {
	CreateUser(ctx *gin.Context)
	GetUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

type UserController struct {
	svc services.UserOperationService
}

func NewUserController(svc services.UserOperationService) UserOperationController {
	return &UserController{
		svc: svc,
	}
}

func (user *UserController) CreateUser(ctx *gin.Context) {
	var usr models.User
	err := ctx.ShouldBindJSON(&usr)
	//fmt.Println(usr)
	if err != nil {
		restErr := resterrors.BadRequestError("Invalid Values Provided")
		ctx.JSON(restErr.Status, restErr)
	}
	result, saveErr := user.svc.AddUser(usr)
	if err != nil {
		ctx.JSON(saveErr.Status, saveErr)
	}
	//fmt.Println(&result)
	ctx.JSON(http.StatusOK, result.Marshall(ctx.GetHeader("X-Public") == "true"))
}

func (user *UserController) GetUser(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 0, 0)
	result, err := user.svc.GetUser(id)
	if err != nil {
		return
	}
	//fmt.Println(result)
	ctx.JSON(http.StatusOK, result.Marshall(ctx.GetHeader("X-Public") == "true"))
}

func (user *UserController) UpdateUser(ctx *gin.Context) {
	var usr models.User
	err := ctx.ShouldBindJSON(&usr)
	if err != nil {
		restErr := resterrors.BadRequestError("Invalid Values Provided")
		ctx.JSON(restErr.Status, restErr)
	}
	result, saveErr := user.svc.UpdateUser(usr)
	if err != nil {
		ctx.JSON(saveErr.Status, saveErr)
	}
	ctx.JSON(http.StatusOK, result.Marshall(ctx.GetHeader("X-Public") == "true"))
}

func (user *UserController) DeleteUser(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 0, 0)
	err := user.svc.DeleteUser(id)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "User Deleted",
	})
}
