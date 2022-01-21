package service

import (
	"project-3/config"
	"project-3/model"
	"time"
)

func GetAllCategories() ([]*model.Category, error) {
	var categories []*model.Category

	db := config.GetDB()

	err := db.Model(&model.Category{}).Find(&categories).Error

	return categories, err
}

func AddCategory(category model.Category) (*model.Category, error) {
	db := config.GetDB()

	err := db.Model(&model.Category{}).Create(&category).Error

	return &category, err
}

func DeleteCategoryById(id int) (interface{}, error) {
	db := config.GetDB()

	err := db.Model(&model.Category{}).Where("id = ?", id).Delete(&model.Category{}).Error

	return map[string]string{"message": "Catergory has been successfully deleted"}, err
}

func UpdateCategoryById(category model.Category) (*model.Category, error) {

	var (
		currentTime = time.Now()
		db          = config.GetDB()
	)

	category.UpdatedAt = &currentTime
	err := db.Model(model.Category{}).Where("id = ?", category.ID).Updates(&category).Error
	if err != nil {
		return nil, err
	}

	return &category, nil
}
