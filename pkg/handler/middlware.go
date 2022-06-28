package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	authorizationHeader = "Authorization"
	userIDCTX           = "userId"
	userNameCTX         = "username"
	userSurnameCTX      = "usersurname"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "Empty Authorization header!")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "Invalid Authorization header!")
		return
	}

	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
	}

	c.Set(userIDCTX, userId)
}

func getUserID(c *gin.Context) (int, error) {

	id, ok := c.Get(userIDCTX)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "User id not found")
		return 0, errors.New("User id not found")
	}

	intID, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "User id is of invalid type")
		return 0, errors.New("User id is of invalid type")
	}

	return intID, nil

}

func (h *Handler) checkUser(c *gin.Context, id int) (string, error) {
	user, err := h.services.GetUserByName(id)
	if err != nil {
		logrus.Println(err.Error())
	}

	if user.Nickname != "admin" {
		return user.Nickname, errors.New("Not admin")
	}

	return user.Nickname, nil
}
