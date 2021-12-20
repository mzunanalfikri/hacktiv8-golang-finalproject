package controller

import (
	"encoding/json"
	"net/http"
	"project-2/config"
	"project-2/model"
	"project-2/service"
	"strconv"

	"github.com/gorilla/mux"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user model.User

	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := service.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(result)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginParam model.LoginParam

	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&loginParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if isAllowed, user := service.IsLoginAllowed(loginParam); isAllowed {
		token := config.CreateToken(user.ID)

		json.NewEncoder(w).Encode(map[string]string{
			"token": token,
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "email and password not match",
	})
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var (
		vars            = mux.Vars(r)
		id              = vars["id"]
		updateUserParam model.UpdateUserParam
	)

	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&updateUserParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	intID, _ := strconv.Atoi(id)
	user, err := service.UpdateUser(updateUserParam, intID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var (
		vars = mux.Vars(r)
		id   = vars["id"]
	)

	w.Header().Set("Content-Type", "application/json")

	intID, _ := strconv.Atoi(id)
	_, err := service.DeleteUser(intID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Your account has been successfully deleted",
	})
}
