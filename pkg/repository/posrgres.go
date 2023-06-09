package repository

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// const (
// 	usersTable      = "users"
// 	todoListsTable  = "todo_lists"
// 	usersListsTable = "users_lists"
// 	todoItemsTable  = "todo_items"
// 	listsItemsTable = "lists_items"
// )

type Config struct {
	Host     string
	Port     string
	UserName string
	Password string
	DBname   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.UserName, cfg.DBname, cfg.Password, cfg.SSLMode))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return db, nil
}
