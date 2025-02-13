package models

import "gorm.io/gorm"

type Auth struct {
	gorm.Model
	LoginToken        string
	RegistrationToken string
}
