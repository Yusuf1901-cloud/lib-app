package repository

import (
	"fmt"

	"github.com/Yusuf1901-cloud/lib-app/models"
	"github.com/jmoiron/sqlx"
)

type userPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *userPostgres {
	return &userPostgres{
		db: db,
	}
}

func (r *userPostgres) GetAllUsersBooks(userId int) ([]models.Book, error) {
	var results []models.Book
	query := fmt.Sprintf(`SELECT b.id, b.title, b.author, b.created_at, b.published_date, b.is_lent, ub.user_id FROM %s b
						INNER JOIN %s ub ON b.id = ub.book_id 
						WHERE ub.user_id = $1`, booksTable, usersBooksTable)
	err := r.db.Select(&results, query, userId)

	return results, err
}

func (r *userPostgres) GetUserBookById(userId, bookId int) (models.Book, error) {
	var book models.Book
	query := fmt.Sprintf(`SELECT b.id, b.author, b.created_at, b.publidhed_date, b.user_id FROM %s b
						INNER JOIN %s ub ON b.id = ub.book_id 
							WHERE ub.usr_id = $1 AND ub.book_id = $2 `, booksTable, usersBooksTable)
	err := r.db.Select(&book, query, userId, bookId)

	return book, err
}
