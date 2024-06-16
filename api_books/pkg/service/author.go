package service

import (
	"api_books"
	"api_books/pkg/repository"
)

type AuthorsService struct {
	repo repository.Author
}

func NewAuthorsService(repo repository.Author) *AuthorsService {
	return &AuthorsService{repo: repo}
}

func (s *AuthorsService) Create(author api_books.Authors) (int, error) {
	return s.repo.Create(author)
}

func (s *AuthorsService) GetAll() ([]api_books.Authors, error) {
	return s.repo.GetAll()
}

func (s *AuthorsService) GetById(authorId int) (api_books.Authors, error) {
	return s.repo.GetById(authorId)
}

func (s *AuthorsService) Delete(authorId int) error {
	return s.repo.Delete(authorId)
}

func (s *AuthorsService) Update(authorId int, input api_books.UpdateAuthorInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(authorId, input)
}
