package model

import (
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
)

type Comment struct {
	ID        int
	UserID    int
	PhotoID   int
	Message   string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

func (c *Comment) Validate() error {
	if govalidator.IsNull(c.Message) {
		return fmt.Errorf("message is null")
	}

	return nil
}

func (Comment) TableName() string {
	return "comment"
}
