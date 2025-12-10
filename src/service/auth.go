package service

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lunasky-hy/dialy-note-app/src/authorization"
	"github.com/lunasky-hy/dialy-note-app/src/model"
	"github.com/lunasky-hy/dialy-note-app/src/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repos repository.DiaryRepository
	authHandler authorization.AuthHandler
	secret string
}

func (s AuthService) Register(user model.User) (string, error) {
	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	newUser := model.User{Name: user.Name, Password: hashedPass}
	create_err := s.repos.UserCreate(newUser)

	if create_err != nil {
		return "", create_err
	}

	token, err := s.AuthorizeUser(user)
	return token, err
}

func (s AuthService) AuthorizeUser(user model.User) (string, error) {
	auth_err := s.authHandler.AuthUser(&user)

	if auth_err != nil {
		return "", auth_err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Name,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})
	tokenStr, jwt_err := token.SignedString([]byte(s.secret))

	return tokenStr, jwt_err
}

func CreateAuthService(repos repository.DiaryRepository) AuthService {
	secret := os.Getenv("AUTH_SECRET");
	authHander := authorization.CreateAuthHandler(repos)
	s := AuthService{repos: repos, secret: secret, authHandler: authHander}
	return s
}