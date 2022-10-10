package repository

import (
	"fmt"
	"strings"

	"github.com/Yusuf1901-cloud/lib-app/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type librarianPostgres struct {
	db *sqlx.DB
}

func NewLibrarianBookPostgres(db *sqlx.DB) *librarianPostgres {
	return &librarianPostgres{
		db: db,
	}
}

func (r *librarianPostgres) Create(librarianId int, book models.Book) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createBookQuery := fmt.Sprintf("INSERT INTO %s (title, author, count, published_date) VALUES ($1, $2, $3, $4) RETURNING id", booksTable)
	row := tx.QueryRow(createBookQuery, book.Title, book.Author, book.Count, book.Published_date)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createLibrariansBooksQuery := fmt.Sprintf("INSERT INTO %s (librarian_id, book_id) VALUES ($1, $2)", librariansBooksTable)
	_, err = tx.Exec(createLibrariansBooksQuery, librarianId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *librarianPostgres) GetAllLibrarianBooks(librarianId int) ([]models.Book, error) {
	var books []models.Book
	query := fmt.Sprintf("SELECT b.id, b.title, b.author, b.created_at, b.published_date, b.is_lent FROM %s b INNER JOIN %s lb ON b.id = lb.book_id WHERE lb.librarian_id = $1", booksTable, librariansBooksTable)
	err := r.db.Select(&books, query, librarianId)
	return books, err
}

func (r *librarianPostgres) GetLibrairianBookById(librarianId, bookId int) (models.Book, error) {
	var book models.Book
	query := fmt.Sprintf(`SELECT b.id, b.title, b.author, b.created_at, b.published_date, b.is_lent FROM %s b INNER JOIN %s lb 
						ON b.id = lb.book_id 
						WHERE lb.librarian_id = $1 AND b.id = $2`, booksTable, librariansBooksTable)
	err := r.db.Get(&book, query, librarianId, bookId)
	return book, err
}

func (r *librarianPostgres) DeleteLibrarianBook(librarianId, bookId int) error {
	query := fmt.Sprintf("DELETE FROM %s b USING %s lb WHERE b.id = lb.book_id AND lb.librarian_id = $1 AND lb.book_id = $2", booksTable, librariansBooksTable)
	_, err := r.db.Exec(query, librarianId, bookId)
	return err
}

func (r *librarianPostgres) UpdateLibrarianBook(librarianId, bookId int, input models.UpdateBookInput) error {
	var newBookInput models.Book
	newBookInput.IsLent = false // default
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, input.Title)
		argId++
	}
	if input.Author != nil {
		setValues = append(setValues, fmt.Sprintf("author=$%d", argId))
		args = append(args, input.Author)
		argId++
	}
	if *input.Count != 1 {
		setValues = append(setValues, fmt.Sprintf("count=$%d", argId))
		args = append(args, input.Count)
		argId++
	}
	if *input.IsLent {
		setValues = append(setValues, fmt.Sprintf("is_lent=$%d", argId))
		args = append(args, input.IsLent)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE TABLE %s b SET %s FROM %s lb WHERE b.id = lb.book_id AND lb.book_id=$1 AND lb.librarian_id = $2",
		booksTable, setQuery, librariansBooksTable, argId, argId+1)
	args = append(args, bookId, librarianId)
	logrus.Debug("updateQuery : %s", query)
	logrus.Debug("args : %s", args)
	_, err := r.db.Exec(query, args...)
	return err
}
