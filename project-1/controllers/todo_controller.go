package controllers

import (
	"encoding/json"
	"net/http"
	"project-1/database"
	"project-1/models"
)

// GetAllTodos godoc
// @Summary Get all todos
// @Description Get details of all To Do list
// @Tags todos
// @Accept json
// @Produce json
// @Success 200 {array} models.Todo
// @Router /todos [get]
func GetAllTodos(rw http.ResponseWriter, r *http.Request) {
	var todos []models.Todo

	db := database.GetDB()

	db.Find(&todos)
	json.NewEncoder(rw).Encode(todos)
}
