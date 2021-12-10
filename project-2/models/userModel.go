package models

import (
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
)

type User struct {
	ID        int
	Username  string
	Email     string
	Password  string
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) Validate() error {
	if !govalidator.IsEmail(u.Email) {
		return fmt.Errorf("email not valid")
	}

	if govalidator.IsNull(u.Email) {
		return fmt.Errorf("email is null")
	}

	if govalidator.IsNull(u.Username) {
		return fmt.Errorf("username is null")
	}

	if len(u.Password) < 6 {
		return fmt.Errorf("password too short")
	}

	if u.Age < 8 {
		return fmt.Errorf("under age")
	}

	return nil
}
