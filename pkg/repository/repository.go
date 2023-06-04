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
		GetListById(userId, listId int) (todolist.TodoList, error)
		UpdateListInout(userid, listID int, input todolist.UpdateListInput) error
		Delete(userId, listId int) error
	}
	TodoItem interface {
		CreateItem(listId int, input todolist.TodoItem) (int, error)
		GetAllItems(listId int) ([]todolist.TodoItem, error)
		GetItemByID(userID, itemID int) (todolist.TodoItem, error)
		DeleteItem(userID, itemid int) error
		UpdateItem(userID, itemid int, input todolist.UpdateItemInput) error
	}
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
		TodoItem:      NewItemRepo(db),
	}
}
