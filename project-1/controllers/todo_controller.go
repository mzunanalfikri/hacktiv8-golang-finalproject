package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"project-1/database"
	"project-1/models"

	"github.com/gorilla/mux"
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

// GetTodoDetail godoc
// @Summary Get todo detail
// @Description Get details of todo by id
// @Tags todos
// @Accept json
// @Produce json
// @Success 200 {object} models.Todo
// @Router /todo [get]
func GetTodoDetail(rw http.ResponseWriter, r *http.Request) {
	var (
		vars = mux.Vars(r)
		id   = vars["id"]
		todo models.Todo
	)

	db := database.GetDB()

	db.Where("id = ?", id).First(todo)
	json.NewEncoder(rw).Encode(todo)
}

// CreateTodo godoc
// @Summary Create todo
// @Description create todo from parameter
// @Tags todos
// @Accept json
// @Produce json
// @Success 200 {object} models.Todo
// @Router /create-todo [post]
func CreateTodo(rw http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	c, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	if err := json.Unmarshal(c, &todo); err != nil {
		panic(err)
	}	

	db := database.GetDB()

	db.Create(todo)
	json.NewEncoder(rw).Encode(todo)
}
