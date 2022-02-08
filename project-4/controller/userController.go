package controller

import (
	"fmt"
	"net/http"
	"project-4/middleware"
	"project-4/model"
	"project-4/service"
	"project-4/tool"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var (
		user = model.User{
			Role:    "customer",
			Balance: 0,
		}
		password string
	)

	err := c.ShouldBind(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	password = user.Password

	result, err := service.CreateUser(user)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":         result.ID,
		"full_name":  result.Fullname,
		"email":      result.Email,
		"password":   password,
		"balance":    result.Balance,
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

func TopupUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var (
		user  model.User
		claim = middleware.AuthContext(c)
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

	result, err := service.UpdateUserBalance(user)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": fmt.Sprintf("Your balance has been successfully updated to Rp %d", result.Balance),
	})
}
