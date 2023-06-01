package service

import "todolist/pkg/repository"

type (
	Authorization interface{}
	TodoList      interface{}
	TodoItem      interface{}
)

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo repository.Repository) *Service {
	return &Service{}
}
