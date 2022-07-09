package service

import (
	"crypto/sha1"
	"fmt"
	rest_api_auth "github.com/mazhaboy/rest-api-auth"
	"github.com/mazhaboy/rest-api-auth/pkg/repository"
)

const salt = "2543ghgh"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user rest_api_auth.User) (id uint64, err error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
