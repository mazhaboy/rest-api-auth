package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/mazhaboy/rest-api-auth/domain"
)

type AuthorizationRepo interface {
	CreateUser(user domain.User) (id uint64, err error)
	GetUser(username string, password string) (domain.User, error)
}

type ToDoListRepo interface {
}

type ToDoItemRepo interface {
}

type Repository struct {
	AuthorizationRepo
	ToDoListRepo
	ToDoItemRepo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		AuthorizationRepo: NewAuthPostgres(db),
		ToDoListRepo:      nil,
		ToDoItemRepo:      nil,
	}
}
