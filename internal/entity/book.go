package entity

type Book struct {
	Id         int    `json:"-"`
	AuthorName string `json:"author"`
	BookTitle  string `json:"title"`
	ISBN       string `json:"isbn"`
}
