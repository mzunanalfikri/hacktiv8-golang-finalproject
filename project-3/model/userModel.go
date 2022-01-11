package model

import "time"

type User struct {
	ID        int        `json:"id"`
	Fullname  string     `json:"full_name" binding:"required"`
	Email     string     `json:"email" binding:"email,required"`
	Password  string     `json:"password" binding:"required,gte=6"`
	Role      string     `json:"role" binding:"required,oneof='admin' 'member'"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type LoginParam struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
