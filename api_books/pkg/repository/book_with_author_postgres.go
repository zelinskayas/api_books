package repository

import (
	"api_books"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

type BooksWithAuthorsPostgres struct {
	db *sqlx.DB
}

func NewBooksWithAuthorsPostgres(db *sqlx.DB) *BooksWithAuthorsPostgres {
	return &BooksWithAuthorsPostgres{db: db}
}

func (r *BooksWithAuthorsPostgres) Update(bookId, authorId int, input api_books.UpdateBooksWithAuthorsInput) error {
	setValuesBook := make([]string, 0)
	argsBook := make([]interface{}, 0)
	argIdBook := 1

	if input.Title != nil {
		setValuesBook = append(setValuesBook, fmt.Sprintf("title=$%d", argIdBook))
		argsBook = append(argsBook, *input.Title)
		argIdBook++
	}

	if input.AuthorId != nil {
		setValuesBook = append(setValuesBook, fmt.Sprintf("author_id=$%d", argIdBook))
		argsBook = append(argsBook, *input.AuthorId)
		argIdBook++
	}

	if input.PublicationYear != nil {
		setValuesBook = append(setValuesBook, fmt.Sprintf("publication_year=$%d", argIdBook))
		argsBook = append(argsBook, *input.PublicationYear)
		argIdBook++
	}

	if input.Isbn != nil {
		setValuesBook = append(setValuesBook, fmt.Sprintf("isbn=$%d", argIdBook))
		argsBook = append(argsBook, *input.Isbn)
		argIdBook++
	}

	setQueryBook := strings.Join(setValuesBook, ", ")
	queryBook := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", booksTable, setQueryBook, argIdBook)
	argsBook = append(argsBook, bookId)

	logrus.Debugf("updateQueryBook: %s", queryBook)
	logrus.Debugf("argsBook: %s", argsBook)

	//authors
	setValuesAuthor := make([]string, 0)
	argsAuthor := make([]interface{}, 0)
	argIdAuthor := 1
	var dateBirthAuthor time.Time
	var err error

	if input.DateOfBirth != nil {
		dateBirthAuthor, err = ParseDateTime(*input.DateOfBirth)
		if err != nil {
			return err
		}
	}

	if input.FirstName != nil {
		setValuesAuthor = append(setValuesAuthor, fmt.Sprintf("first_name=$%d", argIdAuthor))
		argsAuthor = append(argsAuthor, *input.FirstName)
		argIdAuthor++
	}

	if input.LastName != nil {
		setValuesAuthor = append(setValuesAuthor, fmt.Sprintf("last_name=$%d", argIdAuthor))
		argsAuthor = append(argsAuthor, *input.LastName)
		argIdAuthor++
	}

	if input.Biography != nil {
		setValuesAuthor = append(setValuesAuthor, fmt.Sprintf("biography=$%d", argIdAuthor))
		argsAuthor = append(argsAuthor, *input.Biography)
		argIdAuthor++
	}

	if input.DateOfBirth != nil {
		setValuesAuthor = append(setValuesAuthor, fmt.Sprintf("date_of_birth=$%d", argIdAuthor))
		argsAuthor = append(argsAuthor, dateBirthAuthor)
		argIdAuthor++
	}

	setQueryAuthor := strings.Join(setValuesAuthor, ", ")
	queryAuthor := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", authorsTable, setQueryAuthor, argIdAuthor)
	argsAuthor = append(argsAuthor, authorId)

	logrus.Debugf("updateAuthorQuery: %s", queryAuthor)
	logrus.Debugf("argsAuthor: %s", argsAuthor)

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(queryBook, argsBook...)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(queryAuthor, argsAuthor...)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
