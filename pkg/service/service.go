package service

import (
	"github.com/Ulukbek-Toychuev/book_shop/internal/entity"
	"github.com/Ulukbek-Toychuev/book_shop/pkg/repository"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GenerateTokem(nickname, userPasswd string) (string, error)
	ParseToken(token string) (int, error)
	GetUserByName(id int) (entity.User, error)
}

type Book interface {
	AddBook(userId int, book entity.Book) (int, entity.User, error)
	GetAllBook() ([]entity.Book, error)
}

type Service struct {
	Authorization
	Book
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		Book:          NewBookService(repo.Book),
	}
}
