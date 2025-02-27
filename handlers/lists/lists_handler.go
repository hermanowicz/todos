package handlers

import "github.com/jmoiron/sqlx"

type ListsHandler struct {
	DB sqlx.DB
}

func NewListsHandler(db sqlx.DB) *ListsHandler {
	return &ListsHandler{
		DB: db,
	}
}
