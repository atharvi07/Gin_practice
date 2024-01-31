package routes

import (
	"github.com/atharvi07/gin_practice/internal/app/handler"
	"github.com/atharvi07/gin_practice/internal/app/repository"
	"github.com/atharvi07/gin_practice/internal/app/service"
	"github.com/gin-gonic/gin"
)

func RegisterRoute(engine *gin.Engine) {
	userRepository := repository.NewUserRepository("https://reqres.in")
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	engine.GET("/api/users",userHandler.GetAllUsers)
}