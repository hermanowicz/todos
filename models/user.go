package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Mail    string
	Active  bool
	Blocked bool
}
