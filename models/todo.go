package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	User sql.NullString
	Body string
	List sql.NullString
	Done bool
}
