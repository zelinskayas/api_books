package service

import (
	"api_books"
	"api_books/pkg/repository"
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

type Service struct {
	Author
	Book
	BookWithAuthor
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Author:         NewAuthorsService(repos.Author),
		Book:           NewBooksService(repos.Book),
		BookWithAuthor: NewBooksWithAuthorsService(repos.BookWithAuthor),
	}
}
