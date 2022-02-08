package model

import "time"

type Product struct {
	ID         int        `json:"id"`
	Title      string     `json:"title" binding:"required"`
	Price      int        `json:"price" binding:"gte=0,lte=50000000,required"`
	Stock      int        `json:"stock" binding:"gte=5,required"`
	CategoryID int        `json:"category_id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}
