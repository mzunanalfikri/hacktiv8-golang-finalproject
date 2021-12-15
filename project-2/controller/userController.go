package controller

import (
	"encoding/json"
	"net/http"
	"project-2/config"
	"project-2/model"
	"project-2/tool"
	"time"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user model.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(result)
}

func CreateUser(user model.User) (*model.User, error) {
	db := config.GetDB()

	hashedPassword := tool.HashPassword(user.Password)
	user.Password = hashedPassword
	user.CreatedAt = time.Now()

	err := db.Model(&model.User{}).Create(&user).Error

	return &user, err
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginParam model.LoginParam

	err := json.NewDecoder(r.Body).Decode(&loginParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if isAllowed, user := IsLoginAllowed(loginParam); isAllowed {
		token := config.CreateToken(user.ID)

		json.NewEncoder(w).Encode(token)
	}
}

func IsLoginAllowed(param model.LoginParam) (bool, *model.User) {
	user, err := GetUserByEmail(param.Email)
	if err != nil {
		panic(err)
	}

	return tool.CheckPasswordHash(param.Password, user.Password), user
}

func GetUserByEmail(email string) (*model.User, error) {
	var user model.User

	db := config.GetDB()

	err := db.Model(&model.User{}).Where("email = ?", email).First(&user).Error

	return &user, err
}
