package service

import (
	rest_api_auth "github.com/mazhaboy/rest-api-auth"
	"github.com/mazhaboy/rest-api-auth/pkg/repository"
)

type Authorization interface {
	CreateUser(user rest_api_auth.User) (id uint64, err error)
}

type ToDoList interface {
}

type ToDoItem interface {
}

type Service struct {
	Authorization
	ToDoList
	ToDoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		ToDoList:      nil,
		ToDoItem:      nil,
	}
}
