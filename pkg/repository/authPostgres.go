package repository

import (
	"fmt"

	"github.com/Ulukbek-Toychuev/book_shop/internal/entity"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (ap *AuthPostgres) CreateUser(user entity.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (user_name, user_surname, user_password) values ($1, $2, $3) RETURNING id", UsersTable)
	row := ap.db.QueryRow(query, user.UserName, user.UserSurName, user.UserPassword)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
