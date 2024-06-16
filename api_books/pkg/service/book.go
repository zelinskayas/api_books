package service

import (
	"api_books"
	"api_books/pkg/repository"
)

type BooksService struct {
	repo repository.Book
}

func NewBooksService(repo repository.Book) *BooksService {
	return &BooksService{repo: repo}
}

func (s *BooksService) Create(book api_books.Books) (int, error) {
	return s.repo.Create(book)
}

func (s *BooksService) GetAll() ([]api_books.Books, error) {
	return s.repo.GetAll()
}

func (s *BooksService) GetById(bookId int) (api_books.Books, error) {
	return s.repo.GetById(bookId)
}

func (s *BooksService) Delete(bookId int) error {
	return s.repo.Delete(bookId)
}

func (s *BooksService) Update(bookId int, input api_books.UpdateBookInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(bookId, input)
}
