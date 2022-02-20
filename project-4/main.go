package main

import (
	"os"
	"project-4/config"
	"project-4/controller"
	"project-4/middleware"

	"github.com/gin-gonic/gin"
)

const defaultPort = "8080"

func main() {
	config.StartDB()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r := gin.Default()
	r.POST("/users/register", controller.RegisterUser)
	r.POST("/users/login", controller.LoginUser)

	authGroup := r.Group("")
	authGroup.Use(middleware.Auth())
	authGroup.PATCH("/users/topup", controller.TopupUser)

	adminGroup := r.Group("")
	adminGroup.Use(middleware.Auth())
	adminGroup.Use(middleware.IsAdmin())
	adminGroup.POST("/categories", controller.CreateCategory)
	adminGroup.GET("/categories", controller.GetCategories)
	adminGroup.PATCH("/categories/:categoryId", controller.UpdateCategory)
	adminGroup.DELETE("/categories/:categoryId", controller.DeleteCategory)

	r.Run(":" + port)
}
