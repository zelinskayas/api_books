package repository

import (
	"api_books"
	"github.com/jmoiron/sqlx"
)

type Author interface {
	Create(author api_books.Authors) (int, error)
	GetAll() ([]api_books.Authors, error)
	GetById(authorId int) (api_books.Authors, error)
	Delete(authorId int) error
	Update(authorId int, input api_books.UpdateAuthorInput) error
}

type Book interface {
	Create(book api_books.Books) (int, error)
	GetAll() ([]api_books.Books, error)
	GetById(bookId int) (api_books.Books, error)
	Delete(bookId int) error
	Update(bookId int, input api_books.UpdateBookInput) error
}

type BookWithAuthor interface {
	Update(bookId, authorId int, input api_books.UpdateBooksWithAuthorsInput) error
}

type Repository struct {
	Author
	Book
	BookWithAuthor
}

func NewService() *Repository {
	return &Repository{}
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Author:         NewAuthorsPostgres(db),
		Book:           NewBooksPostgres(db),
		BookWithAuthor: NewBooksWithAuthorsPostgres(db),
	}
}
