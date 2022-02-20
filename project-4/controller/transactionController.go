package controller

import (
	"net/http"
	"project-4/middleware"
	"project-4/model"
	"project-4/service"

	"github.com/gin-gonic/gin"
)

func BuyProduct(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var transactionParam model.TransactionParam
	var claim = middleware.AuthContext(c)

	err := c.ShouldBind(&transactionParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	result, err := service.BuyProduct(transactionParam, claim.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}

func GetUserTransaction(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var (
		claim = middleware.AuthContext(c)
	)

	result, err := service.GetUserTransaction(claim.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}

func GetAllTransaction(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	result, err := service.GetAllTransaction()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}
