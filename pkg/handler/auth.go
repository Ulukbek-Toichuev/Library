package handler

import (
	"github.com/Ulukbek-Toychuev/book_shop/internal/entity"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) SignIn(c *gin.Context) {
	var input entity.User

	if err := c.BindJSON(&input); err != nil {
		logrus.Errorf("Cant binding json:")
	}
}

func (h *Handler) SignUp(c *gin.Context) {

}
