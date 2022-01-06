package main

import (
	"project-3/middleware"
	"project-3/model"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middleware.Auth())

	r.POST("/", func(c *gin.Context) {
		var user model.User
		err := c.ShouldBind(&user)
		if err != nil {
			panic(err)
		}

		c.JSON(200, user)
	})

	r.Run(":8080")
}
