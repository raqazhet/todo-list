package repository

import (
	"todolist"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type ListRepos struct {
	DB *sqlx.DB
}

func NewListRepo(db *sqlx.DB) *ListRepos {
	return &ListRepos{DB: db}
}

func (r *ListRepos) CreateList(userId int, list todolist.TodoList) (int, error) {
	tx, err := r.DB.Begin()
	if err != nil {
		return 0, err
	}
	query := `INSERT INTO todo_lists (title,description)
	VALUES($1,$2) RETURNING id`
	id := 0
	if err := tx.QueryRow(query, list.Title, list.Description).Scan(&id); err != nil {
		tx.Rollback()
		logrus.Printf("createList err: %v", err)
		return 0, err
	}
	stmt := `INSERT INTO users_lists (user_id,list_id)
	VALUES($1,$2)`
	_, err = tx.Exec(stmt, userId, id)
	if err != nil {
		tx.Rollback()
		logrus.Printf("user_Lists insert err: %v", err)
		return 0, err
	}
	return id, tx.Commit()
}

func (r *ListRepos) GetAllLists(userId int) ([]todolist.TodoList, error) {
	query := `SELECT title,description 
	FROM todo_lists 
	INNER JOIN users_lists ON todo_lists.id = users_lists.list_id 
	WHERE user_id = $1`
	tx, err := r.DB.Begin()
	if err != nil {
		logrus.Printf("tx begin err: %v", err)
		return nil, err
	}
	lists := []todolist.TodoList{}
	rows, err := tx.Query(query, userId)
	if err != nil {
		tx.Rollback()
		logrus.Printf("tx.Query() GetAllLists method err: %v", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		list := todolist.TodoList{}
		if err := rows.Scan(&list.Title, &list.Description); err != nil {
			tx.Rollback()
			logrus.Printf("err in rows.Scan GetALLLists method: %v", err)
			return nil, err
		}
		lists = append(lists, list)
	}
	if err := rows.Err(); err != nil {
		tx.Rollback()
		logrus.Printf("rows.Err: %v", err)
		return nil, err
	}
	return lists, tx.Commit()
}

func (r *ListRepos) GetListById(userId, listId int) (todolist.TodoList, error) {
	query := `SELECT title,description FROM todo_lists INNER JOIN users_lists ON todo_lists.id =users_lists.list_id
	where users_lists.user_id = $1 AND users_lists.list_id=$2`
	tx, err := r.DB.Begin()
	if err != nil {
		return todolist.TodoList{}, err
	}
	list := todolist.TodoList{}
	if err := tx.QueryRow(query, userId, listId).Scan(&list.Title, &list.Description); err != nil {
		tx.Rollback()
		logrus.Printf("GetListById err: %v", err)
		return todolist.TodoList{}, err
	}
	return list, nil
}
