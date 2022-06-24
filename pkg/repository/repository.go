package repository

import "github.com/jmoiron/sqlx"

type Book interface {
}

type Repository struct {
	Book
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
