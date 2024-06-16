package api_books

import "errors"

type Authors struct {
	Id          int    `json:"id" db:"id"`
	FirstName   string `json:"firstName" db:"first_name" binding:"required"`
	LastName    string `json:"lastName" db:"last_name" binding:"required"`
	Biography   string `json:"biography" db:"biography"`
	DateOfBirth string `json:"dateOfBirth" db:"date_of_birth"`
}

type Books struct {
	Id              int    `json:"id" db:"id"`
	Title           string `json:"title" db:"title" binding:"required"`
	AuthorId        int    `json:"authorId" db:"author_id" binding:"required"`
	PublicationYear int    `json:"publicationYear" db:"publication_year"`
	Isbn            string `json:"isbn" db:"isbn" binding:"required"`
}

type BooksWithAuthors struct {
	Books   Books   `json:"book"`
	Authors Authors `json:"authors"`
}

type UpdateAuthorInput struct {
	FirstName   *string `json:"firstName"`
	LastName    *string `json:"lastName"`
	Biography   *string `json:"biography"`
	DateOfBirth *string `json:"dateOfBirth"`
}

func (i UpdateAuthorInput) Validate() error {
	if i.FirstName == nil && i.LastName == nil && i.Biography == nil && i.DateOfBirth == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

type UpdateBookInput struct {
	Title           *string `json:"title"`
	AuthorId        *int    `json:"authorId"`
	PublicationYear *int    `json:"publicationYear"`
	Isbn            *string `json:"isbn"`
}

func (i UpdateBookInput) Validate() error {
	if i.Title == nil && i.AuthorId == nil && i.PublicationYear == nil && i.Isbn == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

type UpdateBooksWithAuthorsInput struct {
	UpdateBookInput   `json:"book"`
	UpdateAuthorInput `json:"author"`
}
