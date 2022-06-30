package repository

import (
	"fmt"

	"github.com/Ulukbek-Toychuev/book_shop/internal/entity"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type BookPostgres struct {
	db *sqlx.DB
}

func NewBookPostgres(db *sqlx.DB) *BookPostgres {
	return &BookPostgres{db: db}
}

func (b *BookPostgres) AddBook(userId int, book entity.Book) (int, entity.User, error) {
	var (
		id   int
		user entity.User
	)

	queryForGetUserNickName := fmt.Sprintf("SELECT nickname FROM %s where id=$1", UsersTable)

	if err := b.db.Get(&user, queryForGetUserNickName, userId); err != nil {
		logrus.Println(err.Error(), "Send SELECT QUERY")
		return 1, user, err
	} else if user.Nickname != "admin" {
		return 0, user, nil
	}

	query := fmt.Sprintf("INSERT INTO %s (author_name, book_title, isbn) VALUES ($1, $2, $3) RETURNING id", BookTable)

	rowFromInsert := b.db.QueryRow(query, book.AuthorName, book.BookTitle, book.ISBN)
	if err := rowFromInsert.Scan(&id); err != nil {

		return 0, user, err

	}
	return id, user, nil
}

func (b *BookPostgres) GetAllBook() ([]entity.Book, error) {
	var bookList []entity.Book

	query := fmt.Sprintf("SELECT author_name, book_title, isbn FROM %s", BookTable)

	err := b.db.Select(&bookList, query)
	if err != nil {
		logrus.Println(err.Error(), "FUNC GetAllBook")
		return nil, err
	}

	return bookList, err
}
