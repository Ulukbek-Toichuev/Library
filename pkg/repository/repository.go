package repository

type Book interface {
}

type Repository struct {
	Book
}

func NewRepository() *Repository {
	return &Repository{}
}
