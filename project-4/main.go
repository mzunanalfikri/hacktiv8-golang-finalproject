package main

import (
	"os"
	"project-4/config"
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
	r.Use(middleware.Auth())

	r.Run(":" + port)
}
