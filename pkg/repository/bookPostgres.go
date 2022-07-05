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

//SQL запрос на добавление книги
func (b *BookPostgres) AddBook(userId int, book entity.Book) (int, entity.User, error) {
	var (
		id   int
		user entity.User
	)

	queryForGetUserNickName := fmt.Sprintf("SELECT nickname FROM %s where id=$1", UsersTable)

	if err := b.db.Get(&user, queryForGetUserNickName, userId); err != nil {
		return 1, entity.User{}, err
	} else if user.Nickname != "admin" {
		return 0, entity.User{}, nil
	}

	query := fmt.Sprintf("INSERT INTO %s (author_name, book_title, isbn) VALUES ($1, $2, $3) RETURNING id", BookTable)

	rowFromInsert := b.db.QueryRow(query, book.AuthorName, book.BookTitle, book.ISBN)
	if err := rowFromInsert.Scan(&id); err != nil {

		return 0, entity.User{}, err

	}
	return id, user, nil
}

//SQL запрос получение всех книг
func (b *BookPostgres) GetAllBook() ([]entity.Book, error) {
	var bookList []entity.Book

	query := fmt.Sprintf("SELECT id, author_name, book_title, isbn FROM %s", BookTable)

	err := b.db.Select(&bookList, query)
	if err != nil {
		return nil, err
	}

	return bookList, nil
}

//SQL запрос на получение книг пользователя
func (b *BookPostgres) GetAllMyBooks(userId int) ([]entity.Book, error) {
	var bookList []entity.Book

	titleQuery := fmt.Sprintf("SELECT book_title FROM %s b where b.id = ub.book_id", BookTable)
	authorQuery := fmt.Sprintf("SELECT author_name FROM %s b where b.id = ub.book_id", BookTable)
	isbnQuery := fmt.Sprintf("SELECT isbn FROM %s b where b.id = ub.book_id", BookTable)

	query := fmt.Sprintf("SELECT id, (%s), (%s), (%s) FROM %s ub WHERE ub.user_id = $1", titleQuery, authorQuery, isbnQuery, UsersBookTable)

	if err := b.db.Select(&bookList, query, userId); err != nil {
		logrus.Println(err.Error(), "func GetAllMyBooks")
		return nil, err
	}

	return bookList, nil
}

//SQL запрос на добавление книг в список пользователя
func (b *BookPostgres) AddingUsersBook(userId int, bookID int) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (user_id, book_id) VALUES ( (SELECT users.id FROM %s WHERE users.id = $1), (SELECT book.id FROM %s WHERE book.id = $2) ) RETURNING id;",
		UsersBookTable, UsersTable, BookTable)

	rowFromInsert := b.db.QueryRow(query, userId, bookID)
	if err := rowFromInsert.Scan(&id); err != nil {
		return 1, err
	}
	return id, nil
}

//SQL запрос на получение книги по ID
func (b *BookPostgres) GetBookByID(bookID int) (entity.Book, error) {
	var book entity.Book

	query := fmt.Sprintf("SELECT id, author_name, book_title, isbn FROM %s WHERE id=$1", BookTable)

	if err := b.db.Get(&book, query, bookID); err != nil {
		return entity.Book{}, err
	}
	return book, nil
}

//SQL запрос на удаление книги из списка пользователя
func (b *BookPostgres) DeleteBook(bookID int) error {
	queryDelete := fmt.Sprintf("DELETE FROM %s WHERE id=$1", UsersBookTable)

	_, err := b.db.Exec(queryDelete, bookID)
	if err != nil {
		return err
	}

	return nil
}
