package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/Yusuf1901-cloud/lib-app/models"
	"github.com/Yusuf1901-cloud/lib-app/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "apsuf234l;aknd"
	signingKey = "aweruhbeh;#9asdfQWSF"
	tokenTTl   = 12 * time.Hour
)

type userTokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type librarianTokenClaims struct {
	jwt.StandardClaims
	LIbrarianId int `json:"librarian_id"`
}

type authService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) Authorization {
	return &authService{
		repo: repo,
	}
}

func (s *authService) CreateUser(user models.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *authService) CreateLibrarian(librarian models.Librarian) (int, error) {
	librarian.Password = s.generatePasswordHash(librarian.Password)
	return s.repo.CreateLibrarian(librarian)
}

func (s *authService) GenerateUserToken(username, password string) (string, error) {
	// get user from DB
	user, err := s.repo.GetUser(username, s.generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &userTokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTl).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *authService) ParseUserToken(accesstoken string) (int, error) {
	token, err := jwt.ParseWithClaims(accesstoken, &userTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*userTokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *UserTokenClaims")
	}

	return claims.UserId, nil
}

func (s *authService) ParseLibrarianToken(accesstoken string) (int, error) {
	token, err := jwt.ParseWithClaims(accesstoken, &librarianTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*librarianTokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *LibrarianTokenClaims")
	}

	return claims.LIbrarianId, nil
}

func (s *authService) GenerateLibrarianToken(username, password string) (string, error) {
	// get user from DB
	librarian, err := s.repo.GetLibrarian(username, s.generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &librarianTokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTl).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		librarian.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *authService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
