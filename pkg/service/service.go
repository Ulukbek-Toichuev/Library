package service

import "github.com/Ulukbek-Toychuev/book_shop/pkg/repository"

type Book interface {
}

type Users interface {
}

type Service struct {
	Book
	Users
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
