package models

import "time"

type Photo struct {
	ID        int
	Title     string
	Caption   string
	PhotoUrl  string
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
