package service

import (
	"project-4/config"
	"project-4/model"
	"time"
)

func CreateCategory(category model.Category) (*model.Category, error) {
	var db = config.GetDB()

	category.CreatedAt = time.Now()

	err := db.Model(&model.Category{}).Create(&category).Error

	return &category, err
}

func GetCategoryDetail(id int) (*model.Category, error) {
	var (
		category model.Category
		db       = config.GetDB()
	)

	err := db.Model(&model.Category{}).Where("id = ?", id).First(&category).Error

	return &category, err
}

func GetAllCategories() (*[]model.Category, error) {
	var (
		category []model.Category
		db       = config.GetDB()
	)

	err := db.Model(&model.Category{}).Preload("Product").Find(&category).Error

	return &category, err
}

func UpdateCategory(category model.Category) (*model.Category, error) {
	var (
		db  = config.GetDB()
		now = time.Now()
	)

	category.UpdatedAt = &now

	err := db.Model(&model.Category{}).Where("id = ?", category.ID).Update("type", category.Type).Error
	if err != nil {
		return nil, err
	}

	return GetCategoryDetail(category.ID)
}

func DeleteCategory(id int) error {
	var db = config.GetDB()

	err := db.Model(&model.Category{}).Where("id = ?", id).Delete(&model.Category{}).Error

	return err
}
