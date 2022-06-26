package handler

import (
	"net/http"

	"github.com/Ulukbek-Toychuev/book_shop/internal/entity"
	"github.com/gin-gonic/gin"
)

//Два обработчика для регистрации и аутентификации пользователя

type signInInput struct {
	UserName     string `json:"username" binding:"required"`
	UserSurName  string `json:"usersurname" binding:"required"`
	UserPassword string `json:"password" binding:"required"`
}

//Обработчик для аутентификации

func (h *Handler) SignIn(c *gin.Context) {
	//Для записи данных из JSON от пользователей
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.GenerateTokem(input.UserName, input.UserSurName, input.UserPassword)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": token,
	})

}

//Обработчик для регистрации

func (h *Handler) SignUp(c *gin.Context) {
	//Для записи данных из JSON от пользователей
	var input entity.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
