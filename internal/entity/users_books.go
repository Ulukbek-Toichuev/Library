package entity

type Book struct {
	Id         int    `json:"-"`
	AuthorName string `json:"author" binding:"required"`
	BookTitle  string `json:"title" binding:"required"`
	ISBN       string `json:"isbn" binding:"required"`
}

type UsersBooks struct {
	id     int
	UserID int
	BookID int
}
