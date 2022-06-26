package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/Ulukbek-Toychuev/book_shop/internal/entity"
	"github.com/Ulukbek-Toychuev/book_shop/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

const (
	salt       = "kjdfa548521593asdtgnxcvbu"
	signingKey = "vdznfgjawguqpasd3984oreoyqaw21657fqwd"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (auth *AuthService) CreateUser(user entity.User) (int, error) {
	user.UserPassword = generateUserPasswd(user.UserPassword)

	return auth.repo.CreateUser(user)
}

func (auth *AuthService) GenerateTokem(userName, userSurname, userPasswd string) (string, error) {
	user, err := auth.repo.GetUser(userName, userSurname, generateUserPasswd(userPasswd))
	if err != nil {
		logrus.Println("Error in function \"GenerateTokem\"")
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func generateUserPasswd(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
