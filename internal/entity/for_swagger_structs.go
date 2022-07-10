package entity

type ForSwaggerUserStruct struct {
	UserName     string `json:"username" binding:"required"`
	UserSurName  string `json:"usersurname" binding:"required"`
	Nickname     string `json:"nickname" binding:"required"`
	UserPassword string `json:"password" binding:"required"`
}

type ForSwaggerBookStruct struct {
	AuthorName string `json:"author" binding:"required"`
	BookTitle  string `json:"title" binding:"required"`
	ISBN       string `json:"isbn" binding:"required"`
}
