package model

import "time"

type User struct {
	ID        int        `json:"id"`
	Fullname  string     `json:"full_name" binding:"required"`
	Email     string     `json:"email" binding:"email,required"`
	Password  string     `json:"password" binding:"gte=6,required"`
	Role      string     `json:"role" binding:"oneof='admin' 'customer',required"`
	Balance   int        `json:"balance" binding:"gte=0,lte=1000000000"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type LoginParam struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
