package defs

import "database/sql"

type Todo struct {
	TodoList  sql.NullString `json:"todo_list_name" omit:"empty"`
	TodoTitle string         `json:"todo_title"`
	TodoBody  string         `json:"todo_body"`
	User      sql.NullString `json:"user"`
	Status    bool           `json:"status"`
}
