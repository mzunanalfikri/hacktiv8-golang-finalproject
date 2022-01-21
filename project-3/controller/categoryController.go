package controller

import (
	"fmt"
	"net/http"
	"project-3/middleware"
	"project-3/model"
	"project-3/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

func AddCategory(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var category model.Category

	isAdmin, err := middleware.IsAdmin(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Errorf("authorization header needed"))
		return
	}

	if !isAdmin {
		c.JSON(http.StatusForbidden, "Admin only")
		return
	}

	err = c.ShouldBind(&category)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	result, err := service.AddCategory(category)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}

func UpdateCategory(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var category model.Category

	isAdmin, err := middleware.IsAdmin(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Errorf("authorization header needed"))
		return
	}

	if !isAdmin {
		c.JSON(http.StatusForbidden, "Admin only")
		return
	}

	err = c.ShouldBind(&category)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	category.ID = id

	result, err := service.UpdateCategoryById(category)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)

}

func DeleteCategory(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	isAdmin, err := middleware.IsAdmin(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Errorf("authorization header needed"))
		return
	}

	if !isAdmin {
		c.JSON(http.StatusForbidden, "Admin only")
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))

	result, err := service.DeleteCategoryById(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}
