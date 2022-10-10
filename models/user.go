package models

type User struct {
	Id         int    `json:"-" db:"id"`
	First_name string `json:"first_name" binding:"required"`
	Last_name  string `json:"last_name" binding:"required"`
	Email      string `json:"email" binding:"required"`
	UserName   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

type UserBooks struct {
	Id          int
	BookId      int
	UserId      int
	LibrarianId int
	Status      bool // false ==> qaytarmagan ,, true ==> qaytargan
}
