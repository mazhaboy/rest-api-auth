package repository

import (
	"github.com/jmoiron/sqlx"
	rest_api_auth "github.com/mazhaboy/rest-api-auth"
)

type Authorization interface {
	CreateUser(user rest_api_auth.User) (id uint64, err error)
}

type ToDoList interface {
}

type ToDoItem interface {
}

type Repository struct {
	Authorization
	ToDoList
	ToDoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		ToDoList:      nil,
		ToDoItem:      nil,
	}
}
