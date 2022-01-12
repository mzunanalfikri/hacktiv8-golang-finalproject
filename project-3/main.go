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
	r.POST("/users/login", controller.RegisterUser)
	r.POST("/users/update-account", controller.UpdateUser)
	r.POST("/users/delete-account", controller.DeleteUser)

	r.Run(":8080")
}
