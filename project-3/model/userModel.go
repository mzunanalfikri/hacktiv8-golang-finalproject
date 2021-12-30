package model

import "time"

type User struct {
	ID        int       `json:"id"`
	Fullname  string    `json:"full_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
