package repository

import (
	"github.com/Yusuf1901-cloud/lib-app/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	CreateLibrarian(librarian models.Librarian) (int, error)
	GetUser(username, password string) (models.User, error)
	GetLibrarian(username, password string) (models.Librarian, error)
}

type UserBooks interface {
}

type LibrarianBooks interface {
	Create(librarianId int, book models.Book) (int, error)
	GetAllLibrarianBooks(librarianId int) ([]models.Book, error)
	GetLibrairianBookById(librarianId, bookId int) (models.Book, error)
	DeleteLibrarianBook(librarianId, bookId int) error
	UpdateLibrarianBook(librarianId, bookId int, input models.UpdateBookInput) error
}

type Repository struct {
	Authorization
	UserBooks
	LibrarianBooks
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization:  NewAuthPostgres(db),
		LibrarianBooks: NewLibrarianBookPostgres(db),
	}
}
