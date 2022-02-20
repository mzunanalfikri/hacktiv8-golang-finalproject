package model

import "time"

type TransactionHistory struct {
	ID         int        `json:"id"`
	ProductID  int        `json:"product_id"`
	Product    Product    `json:"product" gorm:"foreignKey:ProductID"`
	UserID     int        `json:"user_id"`
	User       User       `json:"user" gorm:"foreignkey:UserID"`
	Quantity   int        `json:"quantity" binding:"required"`
	TotalPrice int        `json:"total_price" binding:"required"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}

type TransactionParam struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}
