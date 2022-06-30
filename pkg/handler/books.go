package handler

import (
	"net/http"

	"github.com/Ulukbek-Toychuev/book_shop/internal/entity"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type User struct {
	Nickname     string `json:"nickname" binding:"required"`
	UserPassword string `json:"password" binding:"required"`
}

type Book struct {
	AuthorName string `json:"author" binding:"required"`
	BookTitle  string `json:"title" binding:"required"`
}

type getAllBookStruct struct {
	Data []entity.Book `json:"all books"`
}

func (h *Handler) GetAllBook(c *gin.Context) {

	lists, err := h.services.GetAllBook()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllBookStruct{
		Data: lists,
	})

}

func (h *Handler) GetBookByID(c *gin.Context) {

}

func (h *Handler) AddBook(c *gin.Context) {

	var input entity.Book

	UserId, err := getUserID(c)
	if err != nil {
		return
	}

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Check your request body")
		return
	}

	id, user, err := h.services.AddBook(UserId, input)
	if err != nil {
		logrus.Println(err.Error())
		newErrorResponse(c, http.StatusInternalServerError, "Check your request body, server dont understand your shit")
		return
	}

	if user.Nickname != "admin" {
		newErrorResponse(c, http.StatusForbidden, "Permission denied, only admin can add books")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":        id,
		"user nick": user.Nickname,
		"Status":    "You can add books",
	})
}

func (h *Handler) DeleteBook(c *gin.Context) {

}

func (h *Handler) UpdateteBok(c *gin.Context) {

}
