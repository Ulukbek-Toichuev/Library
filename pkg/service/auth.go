package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/Ulukbek-Toychuev/book_shop/internal/entity"
	"github.com/Ulukbek-Toychuev/book_shop/pkg/repository"
)

const salt = "kjdfgkjjbwehba548521593asdtgnxcvbu"

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

func generateUserPasswd(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
