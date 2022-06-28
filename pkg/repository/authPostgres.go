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

//Query for creating user
func (ap *AuthPostgres) CreateUser(user entity.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (user_name, user_surname, nickname, user_password) values ($1, $2, $3, $4) RETURNING id", UsersTable)

	row := ap.db.QueryRow(query, user.UserName, user.UserSurName, user.Nickname, user.UserPassword)
	if err := row.Scan(&id); err != nil {
		logrus.Println(err.Error())
		return 0, err
	}

	return id, nil
}

//Query for get users id
func (ap *AuthPostgres) GetUser(nickname, userPasswd string) (entity.User, error) {
	var user entity.User

	query := fmt.Sprintf("SELECT id FROM %s where nickname=$1 AND user_password=$2", UsersTable)

	err := ap.db.Get(&user, query, nickname, userPasswd)
	return user, err
}

//Query for get users nickname
func (ap *AuthPostgres) GetUserByName(id int) (entity.User, error) {
	var user entity.User

	query := fmt.Sprintf("SELECT nickname FROM %s where id=$1", UsersTable)

	err := ap.db.Get(&user, query, id)
	return user, err
}
