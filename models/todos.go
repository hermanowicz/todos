package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type Todos struct {
	gorm.Model
	TodoList  sql.NullString `json:"todo_list_name"`
	TodoTitle string         `json:"todo_title"`
	TodoBody  string         `json:"todo_body"`
	User      sql.NullString `json:"user"`
	Status    bool           `json:"status"`
}
