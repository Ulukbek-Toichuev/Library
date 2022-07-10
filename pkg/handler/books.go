package handler

import (
	"fmt"
	"net/http"
	"strconv"

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

// @Summary Get user books
// @Security ApiKeyAuth
// @Tags books
// @Description Get all user books in library
// @ID get-user-books
// @Accept json
// @Producer json
// @Succes 200 {string} string
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/book/my [get]
func (h *Handler) GetAllMyBooks(c *gin.Context) {
	UserId, err := getUserID(c)
	if err != nil {
		return
	}

	lists, err := h.services.GetAllMyBooks(UserId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllBookStruct{
		Data: lists,
	})

}

// @Summary Get all books
// @Security ApiKeyAuth
// @Tags books
// @Description Get all books in library
// @ID get-all-books
// @Accept json
// @Producer json
// @Succes 200 {string} string
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/book/all [get]
func (h *Handler) GetAllBooks(c *gin.Context) {

	lists, err := h.services.GetAllBook()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllBookStruct{
		Data: lists,
	})

}

// @Summary Get book by ID
// @Security ApiKeyAuth
// @Tags books
// @Description Get book by ID in library
// @ID get-book-by-ID
// @Accept json
// @Producer json
// @Param id path int true "Book ID"
// @Succes 200 {string} string
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/book/{id} [get]
func (h *Handler) GetBookByID(c *gin.Context) {
	var input entity.BookID

	if err := c.ShouldBindUri(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Internal server Error")
		return
	}

	id, err := h.services.GetBookByID(input.Id)
	if err != nil {
		logrus.Println(err.Error())
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":     id.Id,
		"title":  id.BookTitle,
		"author": id.AuthorName,
		"isbn":   id.ISBN,
	})
}

// @Summary Add book
// @Security ApiKeyAuth
// @Tags books
// @Description Adding book to the library
// @ID add-book
// @Accept json
// @Producer json
// @Param input body entity.ForSwaggerBookStruct true "Book"
// @Succes 200 {string} string
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/book/admin [post]
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
		"id":        fmt.Sprint(id),
		"user nick": user.Nickname,
		"Status":    "Book is succesfully added to the library!",
	})
}

// @Summary Delete user book
// @Security ApiKeyAuth
// @Tags books
// @Description Delete user book by ID in library
// @ID delete-user-book
// @Accept json
// @Producer json
// @Param id path int true "Book ID"
// @Succes 200 {string} string
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/book/{id} [delete]
func (h *Handler) DeleteBook(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Println(err.Error(), " Error in handler")
		newErrorResponse(c, http.StatusBadRequest, "Invalid id in uri")
		return
	}

	err = h.services.DeleteBook(id)
	if err != nil {
		logrus.Println(err.Error(), " Error in handler")
		newErrorResponse(c, http.StatusInternalServerError, "Something went wrong")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "ok",
	})

}

func (h *Handler) UpdateteBok(c *gin.Context) {

}

// @Summary Add book for user list
// @Security ApiKeyAuth
// @Tags books
// @Description Adding book for user list
// @ID add-book-for-user-list
// @Accept json
// @Producer json
// @Param input body entity.BookID true "book ID"
// @Succes 200 {string} string
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/book/ [post]
func (h *Handler) AddingUsersBook(c *gin.Context) {
	var input entity.BookID

	UserId, err := getUserID(c)
	if err != nil {
		return
	}
	logrus.Println(UserId)

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Check your request body")
		return
	}

	id, err := h.services.AddingUsersBook(UserId, input.Id)
	if err != nil {
		logrus.Println(err.Error())
		newErrorResponse(c, http.StatusInternalServerError, "Error on the server")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":     fmt.Sprint(id),
		"Status": "Book is succesfully added to the library!",
	})

}
