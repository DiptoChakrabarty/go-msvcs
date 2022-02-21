package controllers

import (
	"net/http"
	"strconv"

	"github.com/DiptoChakrabarty/go-mvcs/users/services"
	"github.com/DiptoChakrabarty/go-mvcs/users/types/users"
	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	var usr users.User
	err := ctx.ShouldBindJSON(&usr)
	if err != nil {
		return
	}
	result, err := services.AddUser(usr)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusCreated, result)
}

func GetUser(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 0, 0)
	result, err := services.GetUser(id)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusCreated, result)
}

func DeleteUser(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 0, 0)
	result, err := services.DeleteUser(id)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusCreated, result)
}
