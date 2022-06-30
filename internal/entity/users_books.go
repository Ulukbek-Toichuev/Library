package entity

type Book struct {
	//Id         int    `json:"id" db:"id"`
	Id         int    `json:"-"`
	AuthorName string `json:"author" db:"author_name" binding:"required"`
	BookTitle  string `json:"title" db:"book_title" binding:"required"`
	ISBN       string `json:"isbn" db:"isbn" binding:"required"`
}

type UsersBooks struct {
	id     int
	UserID int
	BookID int
}
