package repository

import (
	"fmt"

	"github.com/Yusuf1901-cloud/lib-app/models"
	"github.com/jmoiron/sqlx"
)

type authPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *authPostgres {
	return &authPostgres{
		db: db,
	}
}

func (r *authPostgres) CreateUser(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (first_name, last_name, username, email, password_hash) VALUES ($1, $2, $3, $4, $5) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.First_name, user.Last_name, user.UserName, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *authPostgres) CreateLibrarian(librarian models.Librarian) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (full_name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", librariansTable)
	row := r.db.QueryRow(query, librarian.Full_name, librarian.Username, librarian.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *authPostgres) GetUser(username, password string) (models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT id FROM  %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)
	return user, err
}

func (r *authPostgres) GetLibrarian(username, password string) (models.Librarian, error) {
	var librarian models.Librarian
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", librariansTable)

	err := r.db.Get(&librarian, query, username, password)
	return librarian, err
}
