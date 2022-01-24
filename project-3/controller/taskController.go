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

func GetTasks(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var (
		claim = middleware.AuthContext(c)
	)

	if claim == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Errorf("authorization header needed"))
		return
	}

	tasks, err := service.GetAllTasks()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func AddTask(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var task model.Task

	isAdmin, err := middleware.IsAdmin(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Errorf("authorization header needed"))
		return
	}

	if !isAdmin {
		c.JSON(http.StatusForbidden, "Admin only")
		return
	}

	err = c.ShouldBind(&task)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	result, err := service.AddTask(task)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}

func UpdateTask(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var (
		claim = middleware.AuthContext(c)
		task  = model.Task{}
	)

	if claim == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Errorf("authorization header needed"))
		return
	}

	task.ID = claim.ID

	err := c.ShouldBind(&task)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	result, err := service.UpdateTask(task)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":          result.ID,
		"title":       result.Title,
		"description": result.Description,
		"status":      result.Status,
		"user_id":     result.UserID,
		"category_id": result.CategoryID,
		"updated_at":  result.UpdatedAt,
	})
}

func UpdateStatusTask(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var task model.Task

	isAdmin, err := middleware.IsAdmin(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Errorf("authorization header needed"))
		return
	}

	if !isAdmin {
		c.JSON(http.StatusForbidden, "Admin only")
		return
	}

	err = c.ShouldBind(&task)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	task.ID = id

	result, err := service.UpdateStatusByTaskId(task)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)

}

func UpdateCategoryTask(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var task model.Task

	isAdmin, err := middleware.IsAdmin(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Errorf("authorization header needed"))
		return
	}

	if !isAdmin {
		c.JSON(http.StatusForbidden, "Admin only")
		return
	}

	err = c.ShouldBind(&task)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	task.ID = id

	result, err := service.UpdateCategoryByTaskId(task)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}

func DeleteTask(c *gin.Context) {
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

	result, err := service.DeleteTaskById(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}
