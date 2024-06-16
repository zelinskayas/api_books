package repository

import (
	"api_books"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

type AuthorsPostgres struct {
	db *sqlx.DB
}

func NewAuthorsPostgres(db *sqlx.DB) *AuthorsPostgres {
	return &AuthorsPostgres{db: db}
}

func ParseDateTime(dateString string) (time.Time, error) {
	dateParse, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		return time.Time{}, errors.New("invalid parse DateOfBirth: " + err.Error())
	}
	return dateParse, nil
}

func (r *AuthorsPostgres) Create(author api_books.Authors) (int, error) {
	var id int
	var dateBirth time.Time

	dateBirth, err := ParseDateTime(author.DateOfBirth)
	if err != nil {
		return 0, err
	}

	createAuthorQuery := fmt.Sprintf("INSERT INTO %s (first_Name, last_Name, biography, date_of_birth) VALUES($1, $2, $3, $4) RETURNING id", authorsTable)
	row := r.db.QueryRow(createAuthorQuery, author.FirstName, author.LastName, author.Biography, dateBirth)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthorsPostgres) GetAll() ([]api_books.Authors, error) {
	var authors []api_books.Authors

	query := fmt.Sprintf("SELECT id, first_name, last_name, biography, to_char(date_of_birth, 'YYYY-MM-DD') as date_of_birth FROM %s", authorsTable)
	err := r.db.Select(&authors, query)

	return authors, err
}

func (r *AuthorsPostgres) GetById(authorId int) (api_books.Authors, error) {
	var author api_books.Authors

	query := fmt.Sprintf("SELECT id, first_name, last_name, biography, to_char(date_of_birth, 'YYYY-MM-DD') as date_of_birth FROM %s where id = $1", authorsTable)
	err := r.db.Get(&author, query, authorId)

	return author, err
}

func (r *AuthorsPostgres) Delete(authorId int) error {
	query := fmt.Sprintf("DELETE FROM %s where id = $1", authorsTable)
	_, err := r.db.Exec(query, authorId)

	return err
}

func (r *AuthorsPostgres) Update(authorId int, input api_books.UpdateAuthorInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	var dateBirth time.Time
	var err error

	if input.DateOfBirth != nil {
		dateBirth, err = ParseDateTime(*input.DateOfBirth)
		if err != nil {
			return err
		}
	}

	if input.FirstName != nil {
		setValues = append(setValues, fmt.Sprintf("first_name=$%d", argId))
		args = append(args, *input.FirstName)
		argId++
	}

	if input.LastName != nil {
		setValues = append(setValues, fmt.Sprintf("last_name=$%d", argId))
		args = append(args, *input.LastName)
		argId++
	}

	if input.Biography != nil {
		setValues = append(setValues, fmt.Sprintf("biography=$%d", argId))
		args = append(args, *input.Biography)
		argId++
	}

	if input.DateOfBirth != nil {
		setValues = append(setValues, fmt.Sprintf("date_of_birth=$%d", argId))
		args = append(args, dateBirth)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", authorsTable, setQuery, argId)
	args = append(args, authorId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err = r.db.Exec(query, args...)
	return err
}
