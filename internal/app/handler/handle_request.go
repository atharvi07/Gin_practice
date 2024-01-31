package handler

import (
	"net/http"

	"github.com/atharvi07/gin_practice/internal/app/service"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	GetAllUsers(ctx *gin.Context)
}

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{
		userService: userService,
	}
}

func (userHandler userHandler) GetAllUsers(ctx *gin.Context) {
	data, err := userHandler.userService.GetAllUsers()
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, data)
}
