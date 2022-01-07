package service

import (
	"project-3/config"
	"project-3/model"
	"project-3/tool"
	"time"
)

func CreateUser(user model.User) (*model.User, error) {
	db := config.GetDB()

	hashedPassword := tool.HashPassword(user.Password)
	user.Password = hashedPassword
	user.CreatedAt = time.Now()

	err := db.Model(&model.User{}).Create(&user).Error

	return &user, err
}
