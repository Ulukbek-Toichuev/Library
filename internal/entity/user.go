package entity

type User struct {
	Id           string `json:"-"`
	UserName     string `json:"username" binding:"required"`
	UserSurName  string `json:"usersurname" binding:"required"`
	UserPassword string `json:"password" binding:"required"`
}
