package models

import "gorm.io/gorm"

type TodoList struct {
	gorm.Model
	Name     string
	ListName string
}
