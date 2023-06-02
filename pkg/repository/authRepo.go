package repository

import (
	"todolist"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type AuthRepo struct {
	db *sqlx.DB
}

func NewAuthRepo(db *sqlx.DB) *AuthRepo {
	return &AuthRepo{
		db: db,
	}
}

func (r *AuthRepo) CreateUser(user todolist.User) (int, error) {
	query := `INSERT INTO users (name, username,password_hash)
	VALUES ($1,$2,$3) RETURNING id`
	args := []any{user.Name, user.Username, user.Password}
	rows := r.db.QueryRow(query, args...)
	id := 0
	if err := rows.Scan(&id); err != nil {
		logrus.Printf("scan err: %v", err)
		return 0, err
	}
	return id, nil
}

func (r *AuthRepo) GetUser(username, password string) (todolist.User, error) {
	query := `SELECT * FROM users WHERE username=$1 and password_hash=$2`
	users := todolist.User{}
	if err := r.db.QueryRow(query, username, password).Scan(&users.Id, &users.Name, &users.Username, &users.Password); err != nil {
		logrus.Printf("GetUser err: %v", err)
		return todolist.User{}, err
	}
	return users, nil
}
