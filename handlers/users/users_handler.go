package handlers

import "github.com/jmoiron/sqlx"

type UsersHandler struct {
	DB sqlx.DB
}

func NewUsersHandler(db sqlx.DB) *UsersHandler {
	return &UsersHandler{
		DB: db,
	}
}
