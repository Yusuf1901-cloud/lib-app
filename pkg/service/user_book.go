package service

import (
	"github.com/Yusuf1901-cloud/lib-app/models"
	"github.com/Yusuf1901-cloud/lib-app/pkg/repository"
)

type usersBooksService struct {
	repo repository.UserBooks
}

func NewUsersBooksService(repo repository.UserBooks) *usersBooksService {
	return &usersBooksService{
		repo: repo,
	}
}

func (s *usersBooksService) GetAllUsersBooks(userId int) ([]models.Book, error) {
	return s.repo.GetAllUsersBooks(userId)
}

func (s *usersBooksService) GetUserBookById(userId, bookId int) (models.Book, error) {
	return s.repo.GetUserBookById(userId, bookId)
}
