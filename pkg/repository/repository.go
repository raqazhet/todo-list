package repository

import "github.com/jmoiron/sqlx"

type (
	Authorization interface{}
	TodoList      interface{}
	TodoItem      interface{}
)

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
