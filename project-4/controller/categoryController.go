package controller

import (
	"fmt"
	"net/http"
	"project-4/middleware"
	"project-4/model"
	"project-4/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var (
		category model.Category
		claim    = middleware.AuthContext(c)
	)

	if claim == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Errorf("authorization header needed"))
		return
	}

	err := c.ShouldBind(&category)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	result, err := service.CreateCategory(category)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":                  result.ID,
		"type":                result.Type,
		"sold_product_amount": result.SoldProductAmount,
		"created_at":          result.CreatedAt,
	})
}

func GetCategories(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var (
		claim = middleware.AuthContext(c)
	)

	if claim == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Errorf("authorization header needed"))
		return
	}

	categories, err := service.GetAllCategories()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, categories)
}

func UpdateCategory(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var (
		category model.Category
		claim    = middleware.AuthContext(c)
	)

	if claim == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Errorf("authorization header needed"))
		return
	}

	err := c.ShouldBind(&category)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	id, _ := strconv.Atoi(c.Param("categoryId"))
	category.ID = id

	result, err := service.UpdateCategory(category)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":                  result.ID,
		"type":                result.Type,
		"sold_product_amount": result.SoldProductAmount,
		"updated_at":          result.UpdatedAt,
	})
}

func DeleteCategory(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var (
		claim = middleware.AuthContext(c)
	)

	if claim == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Errorf("authorization header needed"))
		return
	}

	id, _ := strconv.Atoi(c.Param("categoryId"))

	err := service.DeleteCategory(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Category has been successfully deleted",
	})
}
