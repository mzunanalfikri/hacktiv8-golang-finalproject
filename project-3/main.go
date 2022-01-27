package main

import (
	"os"
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

	authGroup.GET("/tasks", controller.GetTasks)
	authGroup.POST("/tasks", controller.AddTask)
	authGroup.POST("/tasks/:id", controller.UpdateTask)
	authGroup.PATCH("/tasks/update-status/:id", controller.UpdateStatusTask)
	authGroup.PATCH("/tasks/update-category/:id", controller.UpdateCategoryTask)
	authGroup.DELETE("/tasks/:id", controller.DeleteTask)

	r.Run(":" + os.Getenv("PORT"))
}
