package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	User string
	Body string
	List string
	Done bool
}
