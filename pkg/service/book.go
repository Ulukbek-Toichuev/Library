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

func (b *BookService) AddBook(userId int, book entity.Book) (int, error) {

	return b.repo.AddBook(userId, book)
}
