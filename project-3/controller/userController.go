package controller

import (
	"fmt"
	"net/http"
	"project-3/middleware"
	"project-3/model"
	"project-3/service"
	"project-3/tool"

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

func LoginUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var loginParam model.LoginParam

	err := c.ShouldBind(&loginParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	if isAllowed, user := service.IsLoginAllowed(loginParam); isAllowed {
		token := tool.TokenCreate(user.ID)

		c.JSON(http.StatusOK, map[string]string{
			"token": token,
		})
		return
	}

	c.JSON(http.StatusForbidden, map[string]interface{}{
		"message": "email and password not match",
	})
}

func UpdateUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var (
		claim = middleware.AuthContext(c)
		user  = model.User{}
	)

	if claim == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Errorf("authorization header needed"))
		return
	}

	user.ID = claim.ID

	err := c.ShouldBind(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	result, err := service.UpdateUser(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":         result.ID,
		"full_name":  result.Fullname,
		"email":      result.Email,
		"updated_at": result.UpdatedAt,
	})
}

func DeleteUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var (
		claim = middleware.AuthContext(c)
	)

	_, err := service.DeleteUser(claim.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Your account has been successfully deleted",
	})
}
