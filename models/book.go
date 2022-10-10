package models

import "errors"

type Book struct {
	Id             int    `json:"id" db:"id"`
	Title          string `json:"title" db:"title" binding:"required"`
	Author         string `json:"author" db:"author" binding:"required"`
	Created_at     string `json:"created_at" db:"created_at"`
	Count          int    `json:"count" db:"count"`
	IsLent         bool   `json:"is_lent"`
	Published_date string `json:"published_date" db:"published_date" binding:"required"`
}

type UpdateBookInput struct {
	Title  *string `json:"title"`
	Author *string `json:"author"`
	Count  *int    `json:"count"`
	IsLent *bool   `json:"is_lent"`
}

type GiveBookRequest struct {
	BookId      int
	StudentId   int
	LibrarianId int
}

func (i UpdateBookInput) Validate() error {
	if i.Title == nil && i.Author == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
