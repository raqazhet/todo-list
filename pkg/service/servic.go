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
	TodoList interface{}
	TodoItem interface{}
)

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo repository.Repository) *Service {
	return &Service{Authorization: NewAuthService(repo)}
}
