package models

import "gorm.io/gorm"

type TodoLists struct {
	gorm.Model
	User     string
	ListName string `gorm:"index"`
}
