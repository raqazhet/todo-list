package service

import (
	"todolist"
	"todolist/pkg/repository"
)

type (
	Authorization interface {
		CreateUser(user todolist.User) (int, error)
		GenerateToken(username, password string) (string, error)
		ParseToken(token string) (int, error)
	}
	TodoList interface {
		CreateList(userId int, list todolist.TodoList) (int, error)
		GetAllLists(userId int) ([]todolist.TodoList, error)
		GetListById(userId, listId int) (todolist.TodoList, error)
		Delete(userId, listId int) error
	}
	TodoItem interface{}
)

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo),
		TodoList:      NewListService(repo),
	}
}
