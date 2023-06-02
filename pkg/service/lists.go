package service

import (
	"todolist"
	"todolist/pkg/repository"
)

type ListsService struct {
	repo repository.Repository
}

func NewListService(repo repository.Repository) *ListsService {
	return &ListsService{repo: repo}
}

func (s *ListsService) CreateList(userId int, list todolist.TodoList) (int, error) {
	listId, err := s.repo.TodoList.CreateList(userId, list)
	if err != nil {
		return 0, err
	}
	return listId, nil
}

func (s *ListsService) GetAllLists(userId int) ([]todolist.TodoList, error) {
	return s.repo.GetAllLists(userId)
}
