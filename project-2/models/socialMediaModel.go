package models

import (
	"fmt"

	"github.com/asaskevich/govalidator"
)

type SocialMedia struct {
	ID             int
	Name           string
	SocialMediaUrl string
	UserID         int
}

func (s *SocialMedia) Validate() error {
	if govalidator.IsNull(s.Name) {
		return fmt.Errorf("name is null")
	}

	if govalidator.IsNull(s.SocialMediaUrl) {
		return fmt.Errorf("social media url is null")
	}

	return nil
}
