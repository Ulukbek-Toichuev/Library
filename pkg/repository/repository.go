package repository

import (
	"github.com/Ulukbek-Toychuev/book_shop/internal/entity"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GetUser(nickname, userPasswd string) (entity.User, error)
	GetUserByName(id int) (entity.User, error)
}

type Book interface {
	AddBook(userId int, book entity.Book) (int, entity.User, error)
}

type Repository struct {
	Book
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Book:          NewBookPostgres(db),
	}
}
