package service

import (
	"project-4/config"
	"project-4/model"
	"project-4/tool"
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

func UpdateUserBalance(user model.User) (*model.User, error) {
	var (
		db  = config.GetDB()
		now = time.Now()
	)

	user.UpdatedAt = &now

	err := db.Model(&model.User{}).Where("id = ?", user.ID).Update("balance", user.Balance).Error

	return &user, err
}

func GetUserDetail(id int) (*model.User, error) {
	var (
		user model.User
		db   = config.GetDB()
	)

	err := db.Model(&model.User{}).Where("id = ?", id).First(&user).Error

	return &user, err
}
