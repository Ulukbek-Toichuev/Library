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

func (b *BookPostgres) AddBook(userId int, book entity.Book) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (author_name, book_title, isbn) VALUES ($1, $2, $3) RETURNING id", BookTable)

	row := b.db.QueryRow(query, book.AuthorName, book.BookTitle, book.ISBN)
	if err := row.Scan(&id); err != nil {
		logrus.Println(err.Error())

	}
	return id, nil
}
