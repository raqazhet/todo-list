package repository

import (
	"todolist"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type ItemRepos struct {
	DB *sqlx.DB
}

func NewItemRepo(db *sqlx.DB) *ItemRepos {
	return &ItemRepos{DB: db}
}

func (r *ItemRepos) CreateItem(listId int, input todolist.TodoItem) (int, error) {
	tx, err := r.DB.Begin()
	if err != nil {
		return 0, err
	}
	query := `Insert into todo_items (title,description)
	VALUES($1,$2) RETURNING id`
	var itemID int
	if err := tx.QueryRow(query, input.Title, input.Description).Scan(&itemID); err != nil {
		tx.Rollback()
		logrus.Printf("createItem err: %v", err)
		return 0, err
	}
	stmt := `Insert into lists_items (item_id,list_id)
	Values($1,$2)`
	_, err = tx.Exec(stmt, itemID, listId)
	if err != nil {
		tx.Rollback()
		logrus.Printf("insertInto list_item err: %v", err)
		return 0, err
	}
	return itemID, tx.Commit()
}

func (r *ItemRepos) GetAllItems(listId int) ([]todolist.TodoItem, error) {
	tx, err := r.DB.Begin()
	if err != nil {
		return nil, err
	}
	query := `SELECT todo_items.id,title, description, done
	FROM todo_items
	JOIN lists_items ON todo_items.id = lists_items.item_id
	WHERE lists_items.list_id = $1;`
	items := []todolist.TodoItem{}
	rows, err := tx.Query(query, listId)
	if err != nil {
		logrus.Printf("getAllItems err: %v", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		item := todolist.TodoItem{}
		if err := rows.Scan(&item.Id, &item.Title, &item.Description, &item.Done); err != nil {
			logrus.Printf("getAllItems scan err: %v", err)
			return nil, err
		}
		items = append(items, item)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func (r *ItemRepos) GetItemByID(userID, itemId int) (todolist.TodoItem, error) {
	query := `SELECT ti.id, ti.title, ti.description, ti.done
	FROM todo_items ti
	INNER JOIN lists_items li ON ti.id = li.item_id
	INNER JOIN users_lists ul ON ul.list_id = li.list_id
	WHERE ti.id = $1 AND ul.user_id = $2;`

	tx, err := r.DB.Begin()
	if err != nil {
		return todolist.TodoItem{}, err
	}

	item := todolist.TodoItem{}
	if err := tx.QueryRow(query, itemId, userID).Scan(&item.Id, &item.Title, &item.Description, &item.Done); err != nil {
		logrus.Printf("getAllItems err: %v", err)
		return todolist.TodoItem{}, err
	}

	return item, nil
}

func (r *ItemRepos) DeleteItem(userid, itemid int) error {
	query := `DELETE FROM todo_items ti
	USING users_lists ul,lists_items li
	WHERE ti.id = li.item_id ANd ul.list_id = li.list_id AND ul.user_id=$1 And ti.id=$2`
	_, err := r.DB.Exec(query, userid, itemid)
	if err != nil {
		logrus.Printf("DelteItem query err: %v", err)
		return err
	}
	return nil
}

func (r *ItemRepos) UpdateItem(userid int, itemid int, input todolist.UpdateItemInput) error {
	query := `UPDATE todo_items ti
	SET SET title=$1, description=$2,done=$3
	WHERE id IN (
		SELECT ti.id
		FROM users_lists
		WHERE user_id = $3 AND list_id = $4
	)`
	_, err := r.DB.Exec(query, input.Title, input.Description, input.Done, userid, itemid)
	if err != nil {
		logrus.Printf("updateItem err: %v", err)
		return err
	}
	return nil
}
