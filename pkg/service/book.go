package service

import (
	"github.com/Ulukbek-Toychuev/book_shop/internal/entity"
	"github.com/Ulukbek-Toychuev/book_shop/pkg/repository"
)

type BookService struct {
	repo repository.Book
}

func NewBookService(repo repository.Book) *BookService {
	return &BookService{repo: repo}
}

func (b *BookService) AddBook(userId int, book entity.Book) (int, entity.User, error) {

	return b.repo.AddBook(userId, book)
}

func (b *BookService) GetAllBook() ([]entity.Book, error) {
	return b.repo.GetAllBook()
}

func (b *BookService) GetAllMyBooks(userId int) ([]entity.Book, error) {
	return b.repo.GetAllMyBooks(userId)
}

func (b *BookService) AddingUsersBook(userId int, bookID int) (int, error) {
	return b.repo.AddingUsersBook(userId, bookID)
}

func (b *BookService) GetBookByID(bookID int) (entity.Book, error) {
	return b.repo.GetBookByID(bookID)
}

func (b *BookService) DeleteBook(bookID int) error {
	return b.repo.DeleteBook(bookID)
}
