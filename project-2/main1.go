package main

import (
	"project-2/models"
)

func main() {
	user := models.User{
		Email:    "adf@gmail.com",
		Password: "afa",
	}

	user.Validate()
}
