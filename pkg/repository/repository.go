package repository

import (
	"todolist"

	"github.com/jmoiron/sqlx"
)

type (
	Authorization interface {
		CreateUser(user todolist.User) (int, error)
		GetUser(usrname, password string) (todolist.User, error)
	}
	TodoList interface{}
	TodoItem interface{}
)

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Authorization: NewAuthRepo(db)}
}
