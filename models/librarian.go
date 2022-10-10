package models

type Librarian struct {
	Id        int    `json:"-" db:"id"`
	Full_name string `json:"full_name" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type LibrarianBooks struct {
	Id          int
	Book_id     int
	LibrarianId int
	Updated_at  string
	Deleted_at  int
}
