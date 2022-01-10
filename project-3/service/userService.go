package service

import (
	"project-3/config"
	"project-3/model"
	"project-3/tool"
	"time"
)

func IsLoginAllowed(param model.LoginParam) (bool, *model.User) {
	user, err := GetUserByEmail(param.Email)
	if err != nil {
		return false, nil
	}

	return tool.CheckPasswordHash(param.Password, user.Password), user
}

func GetUserByEmail(email string) (*model.User, error) {
	var user model.User

	db := config.GetDB()

	err := db.Model(&model.User{}).Where("email = ?", email).First(&user).Error

	return &user, err
}

func CreateUser(user model.User) (*model.User, error) {
	db := config.GetDB()

	hashedPassword := tool.HashPassword(user.Password)
	user.Password = hashedPassword
	user.CreatedAt = time.Now()

	err := db.Model(&model.User{}).Create(&user).Error

	return &user, err
}
