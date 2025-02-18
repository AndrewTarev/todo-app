package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	userTable       = "users"
	todoListsTable  = "todo_lists"
	userListsTable  = "users_lists"
	todoItemTable   = "todo_items"
	listsItemsTable = "lists_items"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBname   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBname, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
