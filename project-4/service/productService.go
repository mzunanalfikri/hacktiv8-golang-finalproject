package service

import (
	"project-4/config"
	"project-4/model"
	"time"
)

func CreateProduct(product model.Product) (*model.Product, error) {
	db := config.GetDB()

	product.CreatedAt = time.Now()

	err := db.Model(&model.Product{}).Create(&product).Error

	return &product, err
}

func GetAllProducts() (*[]model.Product, error) {
	var (
		products []model.Product
		db       = config.GetDB()
	)

	err := db.Model(&model.Product{}).Find(&products).Error

	return &products, err
}

func GetProductById(id int) (*model.Product, error) {
	var (
		product model.Product
		db      = config.GetDB()
	)

	err := db.Model(&model.Product{}).Where("id = ?", id).First(&product).Error

	return &product, err
}

func UpdateProduct(product model.Product) (*model.Product, error) {
	var (
		db  = config.GetDB()
		now = time.Now()
	)

	product.UpdatedAt = &now

	err := db.Model(&model.Product{}).Where("id = ?", product.ID).Updates(product).Error
	if err != nil {
		return nil, err
	}
	return GetProductById(product.ID)
}

func DeleteProduct(id int) error {
	var db = config.GetDB()

	err := db.Model(&model.Product{}).Where("id = ?", id).Delete(&model.Product{}).Error

	return err
}
