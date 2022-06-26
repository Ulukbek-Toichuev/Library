package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllBook(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"message": "Succesfully",
	})
}

func (h *Handler) GetBookByID(c *gin.Context) {

}

func (h *Handler) AddBook(c *gin.Context) {

}

func (h *Handler) DeleteBook(c *gin.Context) {

}

func (h *Handler) UpdateteBok(c *gin.Context) {

}
