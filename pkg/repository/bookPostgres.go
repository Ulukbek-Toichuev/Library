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

	return bookList, nil
}

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

func (b *BookPostgres) AddingUsersBook(userId int, book entity.Book) error {
	var id int

	/*

				select
			     (select book_title from book b where b.id = ub.book_id),
			     (select author_name from book b where b.id = ub.book_id)
				from users_book ub
				where ub.user_id = 1;

			insert into users_book
			(user_id, book_id)
			values
			(

			(select users.id from users where users.user_name = 'sergey'), (select book.id from book where book.book_title = '%%%%%%%%%%%%2222')
		);
	*/

	query := fmt.Sprintf("INSERT INTO %s (user_id, book_id) VALUES ( (SELECT users.id FROM %s WHERE users.id = $1), (SELECT book.id FROM %s WHERE book.book_title = $2) ) RETURNING id;", UsersBookTable, UsersTable, BookTable)

	rowFromInsert := b.db.QueryRow(query, userId, book.BookTitle)
	if err := rowFromInsert.Scan(&id); err != nil {
		logrus.Println(err.Error(), "AddingUsersBook()")
		return err
	}
	return nil
}

func (b *BookPostgres) GetBookByID(bookID int) (entity.Book, error) {
	var book entity.Book

	query := fmt.Sprintf("SELECT id, author_name, book_title, isbn FROM %s WHERE id=$1", BookTable)

	if err := b.db.Get(&book, query, bookID); err != nil {
		logrus.Println(err.Error(), " GetBookByID()")
		return entity.Book{}, err
	}
	return book, nil
}

func (b *BookPostgres) DeleteBook(bookID int) error {

	/*trn, err := b.db.Begin()
	if err != nil {
		return entity.Book{}, err
	}

	querySelect := fmt.Sprintf("SELECT id FROM %s WHERE id=$1", BookTable)
	_, err = trn.Query(querySelect, bookID)
	if err != nil {
		trn.Rollback()
		logrus.Println(err.Error(), " ROLLBACK ")
		return entity.Book{}, err
	}*/

	queryDelete := fmt.Sprintf("DELETE FROM %s WHERE id=$1", UsersBookTable)

	_, err := b.db.Exec(queryDelete, bookID)
	if err != nil {
		logrus.Println(err.Error(), " DeleteBook ")
		return err
	}

	return nil
}
