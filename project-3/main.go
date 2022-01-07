package main

import (
	"project-3/controller"
	"project-3/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middleware.Auth())

	r.POST("/users/register", controller.RegisterUser)

	r.Run(":8080")
}
