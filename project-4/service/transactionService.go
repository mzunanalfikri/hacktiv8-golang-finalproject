package service

import (
	"errors"
	"project-4/config"
	"project-4/model"
	"time"
)

func BuyProduct(transactionParam model.TransactionParam, userID int) (interface{}, error) {

	var result map[string]interface{}

	// cek product
	product, err := GetProductById(transactionParam.ProductID)
	if err != nil {
		return result, errors.New("Product not found")
	}

	if transactionParam.Quantity < 0 {
		return result, errors.New("Quantity must > 0")
	}

	if transactionParam.Quantity > product.Stock {
		return result, errors.New("Product stock is not enough.")
	}

	totalPrice := product.Price * transactionParam.Quantity
	user, err := GetUserDetail(userID)
	if err != nil {
		return result, err
	}

	if user.Balance < totalPrice {
		return result, errors.New("Your balance is not enough.")
	}

	category, err := GetCategoryDetail(product.CategoryID)
	if err != nil {
		return result, err
	}

	// pengecekan sudah selesai
	product.Stock = product.Stock - transactionParam.Quantity
	UpdateProduct(*product)
	user.Balance = user.Balance - totalPrice
	UpdateUserBalance(*user)
	category.SoldProductAmount = category.SoldProductAmount + transactionParam.Quantity

	transactionHistory := model.TransactionHistory{
		ProductID:  product.ID,
		UserID:     user.ID,
		Quantity:   transactionParam.Quantity,
		TotalPrice: totalPrice,
		CreatedAt:  time.Now(),
	}
	err = CreateTransactionHistory(transactionHistory)
	if err != nil {
		return result, err
	}

	result = map[string]interface{}{
		"message": "You have succesfully purchased the product.",
		"transaction_bill": map[string]interface{}{
			"total_price":   totalPrice,
			"quantify":      transactionHistory.Quantity,
			"product_title": product.Title,
		},
	}

	return result, nil
}

func CreateTransactionHistory(transactionHistory model.TransactionHistory) error {
	db := config.GetDB()

	err := db.Model(&model.TransactionHistory{}).Create(&transactionHistory).Error

	return err
}

func GetUserTransaction(userId int) (*[]model.TransactionHistory, error) {
	var (
		transactionHistory []model.TransactionHistory
		db                 = config.GetDB()
	)

	err := db.Model(&model.TransactionHistory{}).Preload("Product").Preload("User").Where("user_id = ?", userId).Find(&transactionHistory).Error

	return &transactionHistory, err
}

func GetAllTransaction() (*[]model.TransactionHistory, error) {
	var (
		transactionHistory []model.TransactionHistory
		db                 = config.GetDB()
	)

	err := db.Model(&model.TransactionHistory{}).Preload("Product").Preload("User").Find(&transactionHistory).Error

	return &transactionHistory, err
}
