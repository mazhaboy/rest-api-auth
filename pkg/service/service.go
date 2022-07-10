package service

import (
	"github.com/mazhaboy/rest-api-auth/domain"
	"github.com/mazhaboy/rest-api-auth/pkg/repository"
)

type AuthorizationService interface {
	CreateUser(user domain.User) (id uint64, err error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (uint64, error)
}

type ToDoListService interface {
}

type ToDoItemService interface {
}

type Service struct {
	AuthorizationService
	ToDoListService
	ToDoItemService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		AuthorizationService: NewAuthService(repos.AuthorizationRepo),
		ToDoListService:      nil,
		ToDoItemService:      nil,
	}
}
