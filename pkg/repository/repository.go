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
	TodoList interface {
		CreateList(userId int, list todolist.TodoList) (int, error)
		GetAllLists(userId int) ([]todolist.TodoList, error)
	}
	TodoItem interface{}
)

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepo(db),
		TodoList:      NewListRepo(db),
	}
}
