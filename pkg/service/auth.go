package service

import (
	"crypto/sha1"
	"errors"
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
	tokenTTL   = 10 * time.Hour
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

func (auth *AuthService) GenerateTokem(nickname, userPasswd string) (string, error) {
	user, err := auth.repo.GetUser(nickname, generateUserPasswd(userPasswd))
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

func (auth *AuthService) ParseToken(accessToken string) (int, error) {

	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, nil
	}

	claims, ok := token.Claims.(*tokenClaims)

	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func generateUserPasswd(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (auth *AuthService) GetUserByName(id int) (entity.User, error) {
	var user entity.User

	user, err := auth.repo.GetUserByName(id)
	if err != nil {
		logrus.Fatal(err.Error(), " func GetUserByName")
	}
	return user, nil
}
