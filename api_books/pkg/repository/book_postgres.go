package repository

import (
	"api_books"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
)

type BooksPostgres struct {
	db *sqlx.DB
}

func NewBooksPostgres(db *sqlx.DB) *BooksPostgres {
	return &BooksPostgres{db: db}
}

func (r *BooksPostgres) Create(book api_books.Books) (int, error) {
	var id int

	createBookQuery := fmt.Sprintf("INSERT INTO %s (title, author_id, publication_year, isbn) VALUES($1, $2, $3, $4) RETURNING id", booksTable)
	row := r.db.QueryRow(createBookQuery, book.Title, book.AuthorId, book.PublicationYear, book.Isbn)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *BooksPostgres) GetAll() ([]api_books.Books, error) {
	var books []api_books.Books

	query := fmt.Sprintf("SELECT id, title, author_id, publication_year, isbn FROM %s", booksTable)
	err := r.db.Select(&books, query)

	return books, err
}

func (r *BooksPostgres) GetById(bookId int) (api_books.Books, error) {
	var book api_books.Books

	query := fmt.Sprintf("SELECT id, title, author_id, publication_year, isbn FROM %s where id = $1", booksTable)
	err := r.db.Get(&book, query, bookId)

	return book, err
}

func (r *BooksPostgres) Delete(bookId int) error {
	query := fmt.Sprintf("DELETE FROM %s where id = $1", booksTable)
	_, err := r.db.Exec(query, bookId)

	return err
}

func (r *BooksPostgres) Update(bookId int, input api_books.UpdateBookInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.AuthorId != nil {
		setValues = append(setValues, fmt.Sprintf("author_id=$%d", argId))
		args = append(args, *input.AuthorId)
		argId++
	}

	if input.PublicationYear != nil {
		setValues = append(setValues, fmt.Sprintf("publication_year=$%d", argId))
		args = append(args, *input.PublicationYear)
		argId++
	}

	if input.Isbn != nil {
		setValues = append(setValues, fmt.Sprintf("isbn=$%d", argId))
		args = append(args, *input.Isbn)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", booksTable, setQuery, argId)
	args = append(args, bookId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
