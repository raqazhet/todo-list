package service

import (
	"todolist"
	"todolist/pkg/repository"
)

type ItemService struct {
	repo     repository.TodoItem
	listrepo repository.TodoList
}

func NewItemService(repo repository.TodoList, itemrpo repository.TodoItem) *ItemService {
	return &ItemService{
		repo:     itemrpo,
		listrepo: repo,
	}
}

func (s *ItemService) CreateItem(userId, listId int, input todolist.TodoItem) (int, error) {
	_, err := s.listrepo.GetListById(userId, listId)
	if err != nil {
		return 0, err
	}
	return s.repo.CreateItem(listId, input)
}

func (s *ItemService) GetAllItems(userId, listId int) ([]todolist.TodoItem, error) {
	_, err := s.listrepo.GetListById(userId, listId)
	if err != nil {
		return nil, err
	}
	return s.repo.GetAllItems(listId)
}

func (s *ItemService) GetItemByID(userID, itemID int) (todolist.TodoItem, error) {
	return s.repo.GetItemByID(userID, itemID)
}

func (s *ItemService) DeleteItem(userId, itemId int) error {
	return s.repo.DeleteItem(userId, itemId)
}

func (s *ItemService) UpdateItem(userID, itemid int, input todolist.UpdateItemInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateItem(userID, itemid, input)
}
