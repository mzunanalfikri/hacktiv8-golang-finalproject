package main

import (
	"project-3/config"
	"project-3/controller"
	"project-3/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	config.StartDB()
	r := gin.Default()

	r.POST("/users/register", controller.RegisterUser)
	r.POST("/users/login", controller.LoginUser)

	// r.Use(middleware.Auth())
	authGroup := r.Group("")
	authGroup.Use(middleware.Auth())
	authGroup.POST("/users/update-account", controller.UpdateUser)
	authGroup.POST("/users/delete-account", controller.DeleteUser)

	authGroup.GET("/categories", controller.GetCategories)
	authGroup.POST("/categories", controller.AddCategory)
	authGroup.PATCH("/categories/:id", controller.UpdateCategory)
	authGroup.DELETE("/categories/:id", controller.DeleteCategory)

	r.Run(":8080")
}
