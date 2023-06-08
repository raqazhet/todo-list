package service

import (
	"todolist"
	"todolist/pkg/redisC"
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
		UpdateListInout(userid, listID int, input todolist.UpdateListInput) error
		Delete(userId, listId int) error
	}
	TodoItem interface {
		CreateItem(userId, listId int, input todolist.TodoItem) (int, error)
		GetAllItems(userId, listId int) ([]todolist.TodoItem, error)
		GetItemByID(userID, itemID int) (todolist.TodoItem, error)
		DeleteItem(userID, itemID int) error
		UpdateItem(userID, itemid int, input todolist.UpdateItemInput) error
	}
)

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo repository.Repository, cashe *redisC.RedisCashe) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization, cashe),
		TodoList:      NewListService(repo.TodoList, cashe),
		TodoItem:      NewItemService(repo.TodoList, repo.TodoItem),
	}
}
