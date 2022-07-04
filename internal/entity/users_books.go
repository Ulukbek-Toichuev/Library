package entity

type Book struct {
	Id         int    `json:"id" db:"id"`
	AuthorName string `json:"author" db:"author_name" binding:"required"`
	BookTitle  string `json:"title" db:"book_title" binding:"required"`
	ISBN       string `json:"isbn" db:"isbn" binding:"required"`
}

type UsersBooks struct {
	id     int `db:"id"`
	UserID int `db:"book_id"`
	BookID int `db:"user_id"`
}

type BookID struct {
	Id int `uri:"id" binding:"required"`
}
