package model

import (
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
)

type Photo struct {
	ID        int
	Title     string
	Caption   string
	PhotoUrl  string
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Photo) Validate() error {
	if govalidator.IsNull(p.Title) {
		return fmt.Errorf("title is null")
	}

	if govalidator.IsNull(p.PhotoUrl) {
		return fmt.Errorf("photo url is null")
	}

	return nil
}
