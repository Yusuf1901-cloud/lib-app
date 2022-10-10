package service

import (
	"github.com/Yusuf1901-cloud/lib-app/models"
	"github.com/Yusuf1901-cloud/lib-app/pkg/repository"
)

type librarianBookService struct {
	repo repository.LibrarianBooks
}

func NewLibrarianBookService(repo repository.LibrarianBooks) *librarianBookService {
	return &librarianBookService{
		repo: repo,
	}
}

func (s *librarianBookService) Create(librarianId int, book models.Book) (int, error) {
	return s.repo.Create(librarianId, book)
}

func (s *librarianBookService) GetAllLibrarianBooks(librarianId int) ([]models.Book, error) {
	return s.repo.GetAllLibrarianBooks(librarianId)
}

func (s *librarianBookService) GetLibrarianBookById(librarianId, bookId int) (models.Book, error) {
	return s.repo.GetLibrairianBookById(librarianId, bookId)
}

func (s *librarianBookService) DeleteLibrarianBook(librarianId, bookId int) error {
	return s.repo.DeleteLibrarianBook(librarianId, bookId)
}

func (s *librarianBookService) UpdateLibrarianBook(librarianId, bookId int, input models.UpdateBookInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateLibrarianBook(librarianId, bookId, input)
}
