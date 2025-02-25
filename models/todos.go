package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type Todos struct {
	gorm.Model
	TodoList  string
	TodoTitle string
	TodoBody  string
	User      sql.NullString
	Status    bool
}
