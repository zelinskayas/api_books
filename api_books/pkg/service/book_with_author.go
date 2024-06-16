package service

import (
	"api_books"
	"api_books/pkg/repository"
)

type BooksWithAuthorsService struct {
	repo repository.BookWithAuthor
}

func NewBooksWithAuthorsService(repo repository.BookWithAuthor) *BooksWithAuthorsService {
	return &BooksWithAuthorsService{repo: repo}
}

func (s *BooksWithAuthorsService) Update(bookId, authorId int, input api_books.UpdateBooksWithAuthorsInput) error {
	if err := input.UpdateBookInput.Validate(); err != nil {
		return err
	}

	if err := input.UpdateAuthorInput.Validate(); err != nil {
		return err
	}

	return s.repo.Update(bookId, authorId, input)
}
