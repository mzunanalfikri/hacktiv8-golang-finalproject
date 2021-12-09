package models

import (
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

func (u *User) Validate() {
	if !govalidator.IsEmail(u.Email) {
		// not email
	}

	if govalidator.IsNull(u.Email) {
		// email is null
	}

	if govalidator.IsNull(u.Username) {
		// username is null
	}

	if len(u.Password) < 6 {
		// password too short
	}

	if u.Age < 8 {
		// under age
	}
}
