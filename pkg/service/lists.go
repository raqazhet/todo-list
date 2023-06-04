package service

import (
	"todolist"
	"todolist/pkg/repository"
)

type ListsService struct {
	repo repository.TodoList
}

func NewListService(repo repository.TodoList) *ListsService {
	return &ListsService{repo: repo}
}

func (s *ListsService) CreateList(userId int, list todolist.TodoList) (int, error) {
	listId, err := s.repo.CreateList(userId, list)
	if err != nil {
		return 0, err
	}
	return listId, nil
}

func (s *ListsService) GetAllLists(userId int) ([]todolist.TodoList, error) {
	return s.repo.GetAllLists(userId)
}

func (s *ListsService) GetListById(userId, listId int) (todolist.TodoList, error) {
	return s.repo.GetListById(userId, listId)
}

func (s *ListsService) Delete(userID, lisId int) error {
	return s.repo.Delete(userID, lisId)
}

func (s *ListsService) UpdateListInout(userid, listID int, input todolist.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateListInout(userid, listID, input)
}
