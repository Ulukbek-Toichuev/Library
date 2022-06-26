package repository

import (
	"fmt"

	"github.com/Ulukbek-Toychuev/book_shop/internal/entity"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
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
		logrus.Println("Error in function 'CreateUser'")
		return 0, err
	}

	return id, nil
}

func (ap *AuthPostgres) GetUser(userName, userSurname, userPasswd string) (entity.User, error) {
	var user entity.User

	query := fmt.Sprintf("SELECT id FROM %s where user_name=$1 AND user_surname=$2 AND user_password=$3", UsersTable)

	err := ap.db.Get(&user, query, userName, userSurname, userPasswd)
	if err != nil {
		logrus.Println(err.Error())
	}

	return user, err
}
