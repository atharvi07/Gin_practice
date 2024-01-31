package main

import (

	"github.com/atharvi07/gin_practice/internal/app/routes"
	"github.com/gin-gonic/gin"
)

func main()  {
	engine := gin.Default()
	routes.RegisterRoute(engine)
	engine.Run(":8080")
}