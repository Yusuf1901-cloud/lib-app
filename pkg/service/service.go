package service

import (
	"github.com/Yusuf1901-cloud/lib-app/models"
	"github.com/Yusuf1901-cloud/lib-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	CreateLibrarian(librarian models.Librarian) (int, error)
	GenerateUserToken(username, password string) (string, error)
	GenerateLibrarianToken(username, password string) (string, error)
	ParseUserToken(token string) (int, error)
	ParseLibrarianToken(token string) (int, error)
}

type UserBooks interface {
	GetAllUsersBooks(userId int) ([]models.Book, error)
	GetUserBookById(userId, bookId int) (models.Book, error)
}

type LibrarianBooks interface {
	Create(librarianId int, book models.Book) (int, error)
	GetAllLibrarianBooks(librarianId int) ([]models.Book, error)
	GetLibrarianBookById(librarianId, bookId int) (models.Book, error)
	DeleteLibrarianBook(librarianId, bookId int) error
	UpdateLibrarianBook(librarianId, bookId int, input models.UpdateBookInput) error
	GiveBook(request *models.GiveBookRequest) error
}

type Service struct {
	Authorization
	UserBooks
	LibrarianBooks
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization:  NewAuthService(repos.Authorization),
		LibrarianBooks: NewLibrarianBookService(repos.LibrarianBooks),
		UserBooks:      NewUsersBooksService(repos.UserBooks),
	}
}
