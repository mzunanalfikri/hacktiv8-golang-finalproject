package model

import "time"

type Category struct {
	ID                int        `json:"id"`
	Type              string     `json:"type" binding:"required"`
	SoldProductAmount int        `json:"sold_product_amount"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         *time.Time `json:"updated_at"`
	Product           []*Product `json:"product" gorm:"foreignKey:CategoryID;references:ID"`
}
