package controller

import (
	"net/http"
	"project-3/model"
	"project-3/service"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var user = model.User{
		Role: "member",
	}

	err := c.ShouldBind(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	result, err := service.CreateUser(user)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":         result.ID,
		"full_name":  result.Fullname,
		"email":      result.Email,
		"created_at": result.CreatedAt,
	})
}
