package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/mazhaboy/rest-api-auth/domain"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user domain.User) (uint64, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return uint64(id), nil
}

func (r *AuthPostgres) GetUser(username string, password string) (domain.User, error) {
	var user domain.User
	query := fmt.Sprintf("SELECT id from %s where username=$1 and password_hash=$2", usersTable)

	err := r.db.Get(&user, query, username, password)

	return user, err
}
